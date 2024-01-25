package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"qnc/biz/dal/db"
	mimg "qnc/biz/model/image"
	"qnc/biz/model/order"
	"qnc/biz/model/user"
	service "qnc/biz/service/user"
	"qnc/pkg/errno"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ImageService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewImageService create user service
func NewImageService(ctx context.Context, c *app.RequestContext) *ImageService {
	return &ImageService{ctx: ctx, c: c}
}

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

var SdBaseUrl = "http://127.0.0.1:7860"

func (s *ImageService) ProcessImageUd(req *mimg.ImageUdRequest) (resp *mimg.ImageUdResponse, err error) {
	cuid, exists := s.c.Get("current_user_id")
	if !exists {
		return nil, errno.AuthorizationFailedErr
	}
	req.UserId = cuid.(int64)

	//prepare param
	base64Content, width, height, coordinates, err := prepareParam(req)
	if err != nil {
		return nil, err
	}

	//TODO get price config
	point := 2.0
	pname := "UD image"
	var pid int64 = 1

	ts := time.Now().Unix()
	// add order record
	orderId, err := db.CreateOrder(&db.Order{
		UserId:     req.UserId,
		ProdName:   pname,
		ProdId:     pid,
		RealCost:   point,
		BaseCost:   point,
		Status:     order.STATUS_INIT,
		Ip:         s.c.ClientIP(),
		CreateTime: ts,
		UpdateTime: ts,
	})
	if err != nil {
		return nil, err
	}

	// decrease account
	var freq user.FundsRequest
	freq.UserId = req.UserId
	freq.Amount = point
	freq.EventType = user.TYPE_PAYMENT
	freq.OrderId = orderId
	balance, err := service.NewAccountService(s.ctx, s.c).Decrease(&freq)
	if err != nil {
		return nil, err
	}

	// process image
	res, err := processImage(base64Content, width, height, coordinates)
	if err != nil {
		hlog.Error("process image err:", err)
		return nil, errors.New("process image failed")
	}
	// hlog.Debug((res))
	if res != nil && res["image"] != nil {
		db.UpdateOrder(&db.Order{
			ID:         orderId,
			Status:     order.STATUS_SUCCESS,
			UpdateTime: ts,
		})

		var resp = new(mimg.ImageUdResponse)
		resp.ProcessedImage = res["image"].(string)
		resp.Balance = balance
		return resp, nil
	} else {
		ts := time.Now().Unix()
		db.UpdateOrder(&db.Order{
			ID:         orderId,
			Status:     order.STATUS_FALID,
			Remark:     res["msg"].(string),
			UpdateTime: ts,
		})
		//TODO refund

		return nil, errors.New(res["msg"].(string))
	}
}

func prepareParam(req *mimg.ImageUdRequest) (base64Content string, width, height int, coordinates []Coordinate, err error) {
	f, err := req.FileHeader.Open()
	if err != nil {
		return "", 0, 0, nil, err
		// panic(err)
	}
	defer f.Close()

	// 读取文件到字节数组
	fileRaw, err := io.ReadAll(f)
	if err != nil {
		return "", 0, 0, nil, err
		// panic(err)
	}
	contentType := http.DetectContentType(fileRaw)

	base64Content = base64.StdEncoding.EncodeToString(fileRaw)
	base64Content = "data:" + contentType + ";base64," + base64Content

	f, err = req.FileHeader.Open()
	if err != nil {
		return "", 0, 0, nil, err
		// panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		hlog.Errorf("err = ", err)
		return "", 0, 0, nil, err
	}

	b := img.Bounds()
	width = b.Max.X
	height = b.Max.Y

	// var coordinates []Coordinate
	err = json.Unmarshal([]byte(req.Pos), &coordinates)
	if err != nil {
		return "", 0, 0, nil, err
		// panic(err)
	}
	return
}

func processImage(inputImgStr string, w, h int, cords []Coordinate) (map[string]interface{}, error) {

	var pos [][]int
	if cords != nil {
		pos = make([][]int, len(cords))
		for i, pp := range cords {
			x := int(float64(w) * pp.X)
			y := int(float64(h) * pp.Y)
			pos[i] = []int{x, y}
		}
	}

	// fmt.Println(inputImgStr)
	maskStr, err := detectMask(inputImgStr, w, h, pos)
	if err != nil {
		return nil, err
	}
	var res = make(map[string]interface{})
	if maskStr == "" {
		res["msg"] = "No detected available object, Click on the image to add annotations"
		res["image"] = nil
		return res, nil
	}
	maskStr, err = expandMask(inputImgStr, maskStr, 15)
	if err != nil {
		return nil, err
	}
	if maskStr == "" {
		res["msg"] = "No available mask, Click on the image to add annotations"
		res["image"] = nil
		return res, nil
	}
	processedImg, err := inpainting(inputImgStr, maskStr)
	if err != nil {
		return nil, err
	}
	if processedImg == "" {
		res["msg"] = "Process image failded, try again later"
		res["image"] = nil
		return res, nil
	} else {
		res["msg"] = "success"
		res["image"] = processedImg
		return res, nil
	}
}

func detectMask(imgStr string, w, h int, pos [][]int) (string, error) {
	if pos == nil {
		pos = [][]int{{w / 2, 3 * h / 4}}
	}

	payload := map[string]interface{}{
		"input_image":                  imgStr,
		"dino_enabled":                 true,
		"dino_text_prompt":             "dress",
		"dino_preview_checkbox":        false,
		"dino_model_name":              "GroundingDINO_SwinB (938MB)",
		"sam_positive_points":          pos,
		"dino_preview_boxes_selection": []int{1},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		// panic(err)
		return "", err
	}

	hlog.Infof("detect mask request. ")

	// sdURL := "http://127.0.0.1:7860"
	response, err := http.Post(SdBaseUrl+"/sam/sam-predict", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		// panic(err)
		return "", err
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		// panic(err)
		return "", err
	}
	hlog.Infof("detect mask done: ", reply["msg"])

	if masks, ok := reply["masks"]; ok {
		lens := len(masks.([]interface{}))
		if lens > 1 {
			return masks.([]interface{})[1].(string), nil
		}
	}

	return "", errors.New("detect mask failed")
}

func expandMask(imgStr, maskStr string, dilateAmount int) (string, error) {
	if dilateAmount == 0 {
		dilateAmount = 10
	}
	payload := map[string]interface{}{
		"input_image":   imgStr,
		"mask":          maskStr,
		"dilate_amount": dilateAmount,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		// panic(err)
		return "", err
	}

	hlog.Infof("expand mask request. ")
	// sdURL := "http://127.0.0.1:7860"
	response, err := http.Post(SdBaseUrl+"/sam/dilate-mask", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		// panic(err)
		return "", err
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		// panic(err)
		return "", err
	}

	hlog.Infof("expand mask done.")
	if mask, ok := reply["mask"]; ok {
		return mask.(string), nil
	}

	return "", errors.New("expand mask failed")
}

func inpainting(imgStr, maskStr string) (string, error) {
	// model_id = "sd-v1-5-inpainting.ckpt [c6bbc15e32]"
	modelId := "realisticVisionV60B1_v60B1InpaintingVAE.safetensors [346e4b5a73]"
	samplerName := "DPM++ 2M Karras"

	prompt := "dress"
	negativePrompt := "deformed, bad anatomy, disfigured, poorly drawn face, mutation, mutated, extra limb, ugly, poorly drawn hands, missing limb, floating limbs, disconnected limbs, malformed hands, out of focus, long neck, long body, monochrome, feet out of view, head out of view, lowers, ((bad anatomy)), bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, jpeg artifacts, signature, watermark, username, blurry, artist name, extra limb, poorly drawn eyes, (out of frame), black and white, obese, censored, bad legs, bad hands, text, error, missing fingers, extra digit, fewer digits, cropped, worst quality, low quality, normal quality, jpeg artifacts, signature, watermark, username, blurry, (extra legs), (poorly drawn eyes), without hands, bad knees, multiple shoulders, bad neck, ((no head))"

	payload := map[string]interface{}{
		"sampler_ame":     samplerName,
		"prompt":          prompt,
		"negative_prompt": negativePrompt,
		"init_images":     []string{imgStr},
		"mask":            maskStr,
		"inpainting_fill": 0,
		// "inpainting_mask_invert": 1,
		"inpaint_full_res":         1,
		"inpaint_full_res_padding": 32,
		"sampler_name":             samplerName,
		"seed":                     "-1",
		"cfg_scale":                7.0,
		// "width": w,
		// "height": h,
		"steps": 25,
	}

	overrideSettings := map[string]interface{}{
		"sd_model_checkpoint": modelId,
	}
	payload["override_settings"] = overrideSettings

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		// panic(err)
		return "", err
	}

	hlog.Infof("inpaint img2img request. ")

	// sdURL := "http://127.0.0.1:7860"
	response, err := http.Post(SdBaseUrl+"/sdapi/v1/img2img", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		hlog.Errorf("request img2img err: ", err)
		// panic(err)
		return "", err
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		// panic(err)
		return "", err
	}
	hlog.Infof("inpaint img2img done. ")
	hlog.Debug(reply["info"])

	if images, ok := reply["images"]; ok {
		for _, imgStrs := range images.([]interface{}) {
			imgs := strings.Split(imgStrs.(string), ",")
			return imgs[0], nil
		}
	}
	return "", errors.New("impaint img2img failed")
}

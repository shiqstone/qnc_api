package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func Upload(ctx context.Context, c *app.RequestContext) {
	// pos, _ := c.GetPostForm("pos")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	// fmt.Println(fileHeader.Filename)
	f, err := fileHeader.Open()
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// fmt.Println(f)

	// 读取文件到字节数组
	fileRaw, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	contentType := http.DetectContentType(fileRaw)
	// fmt.Printf(string(contentType))

	base64Content := base64.StdEncoding.EncodeToString(fileRaw)
	base64Content = "data:" + contentType + ";base64," + base64Content
	// fmt.Println(base64Content)
	// base64.StdEncoding.EncodeToString(fileRaw)
	// bufstore := make([]byte, 5000000)            //数据缓存
	// base64.StdEncoding.Encode(bufstore, fileRaw) // 文件转base64
	// fmt.Println(f)

	f, err = fileHeader.Open()
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		hlog.Errorf("err = ", err)
		return
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	// fmt.Println("width = ", width)
	// fmt.Println("height = ", height)

	pos := c.FormValue("pos")
	var coordinates []Coordinate
	err = json.Unmarshal([]byte(pos), &coordinates)
	if err != nil {
		c.JSON(http.StatusOK, utils.H{
			"msg":  err.Error(),
			"code": http.StatusBadRequest,
		})
		return
	}

	hlog.CtxInfof(ctx, "Update request, clientIP: "+c.ClientIP())
	res := processImage(base64Content, width, height, coordinates)
	if res != nil && res["image"] != nil {
		c.JSON(http.StatusOK, utils.H{
			"msg":             "success",
			"processed_image": res["image"],
		})
	} else {
		c.JSON(http.StatusOK, utils.H{
			"msg":             res["msg"],
			"processed_image": nil,
		})
	}
}

func processImage(inputImgStr string, w, h int, cords []Coordinate) map[string]interface{} {

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
	maskStr := detectMask(inputImgStr, w, h, pos)
	var res = make(map[string]interface{})
	if maskStr == "" {
		res["msg"] = "No detected available object, Click on the image to add annotations"
		res["image"] = nil
		return res
	}
	maskStr = expandMask(inputImgStr, maskStr, 15)
	if maskStr == "" {
		res["msg"] = "No available mask, Click on the image to add annotations"
		res["image"] = nil
		return res
	}
	processedImg := inpainting(inputImgStr, maskStr)
	if processedImg == "" {
		res["msg"] = "Process image failded, try again later"
		res["image"] = nil
		return res
	} else {
		res["msg"] = "success"
		res["image"] = processedImg
		return res
	}
}

func detectMask(imgStr string, w, h int, pos [][]int) string {
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
		panic(err)
	}

	hlog.Infof("detect mask request. ")

	sdURL := "http://127.0.0.1:7860"
	response, err := http.Post(sdURL+"/sam/sam-predict", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		panic(err)
	}

	hlog.Infof("detect mask done: ", reply["msg"])
	// fmt.Println(reply["masks"])

	if masks, ok := reply["masks"]; ok {
		lens := len(masks.([]interface{}))
		if lens > 1 {
			return masks.([]interface{})[1].(string)
		}
	}

	return ""
}

func expandMask(imgStr, maskStr string, dilateAmount int) string {
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
		panic(err)
	}

	hlog.Infof("expand mask request. ")
	sdURL := "http://127.0.0.1:7860"
	response, err := http.Post(sdURL+"/sam/dilate-mask", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		panic(err)
	}

	hlog.Infof("expand mask done.")
	if mask, ok := reply["mask"]; ok {
		return mask.(string)
	}

	return ""
}

func inpainting(imgStr, maskStr string) string {
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
		panic(err)
	}

	hlog.Infof("inpaint img2img request. ")

	sdURL := "http://127.0.0.1:7860"
	response, err := http.Post(sdURL+"/sdapi/v1/img2img", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		hlog.Errorf("request img2img err: ", err)
		panic(err)
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		panic(err)
	}
	hlog.Infof("inpaint img2img done. ")
	// fmt.Println(reply["info"])

	if images, ok := reply["images"]; ok {
		for _, imgStrs := range images.([]interface{}) {
			imgs := strings.Split(imgStrs.(string), ",")
			return imgs[0]
		}
	}
	return ""
}

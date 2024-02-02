package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	mimg "qnc/biz/model/image"
	"qnc/biz/mw/viper"
	"strings"

	_ "golang.org/x/image/webp"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ImageService struct {
	ctx context.Context
	c   *app.RequestContext
}

var SdBaseUrl string

func Init() {
	config := viper.Conf.SdService
	SdBaseUrl = config.BaseUrl //"http://127.0.0.1:7860"
}

// NewImageService create image service
func NewImageService(ctx context.Context, c *app.RequestContext) *ImageService {
	return &ImageService{ctx: ctx, c: c}
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
		return "", err
	}

	hlog.Infof("detect mask request. ")

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
		return "", err
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		return "", err
	}

	hlog.Infof("expand mask done.")
	if mask, ok := reply["mask"]; ok {
		return mask.(string), nil
	}

	return "", errors.New("expand mask failed")
}

func inpainting(imgStr, maskStr, prompt string) (string, int64, error) {
	// model_id = "sd-v1-5-inpainting.ckpt [c6bbc15e32]"
	modelId := "realisticVisionV60B1_v60B1InpaintingVAE.safetensors [346e4b5a73]"
	samplerName := "DPM++ 2M Karras"

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
		return "", 0, err
	}

	hlog.Infof("inpaint img2img request. ")

	response, err := http.Post(SdBaseUrl+"/sdapi/v1/img2img", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		hlog.Errorf("request img2img err: ", err)
		return "", 0, err
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	if err != nil {
		return "", 0, err
	}
	hlog.Infof("inpaint img2img done. ")

	var seed int64
	infoStr := reply["info"].(string)
	info := make(map[string]interface{})
	err = json.Unmarshal([]byte(infoStr), &info)
	if err == nil {
		seed = int64(info["seed"].(float64))
	}

	if images, ok := reply["images"]; ok {
		for _, imgStrs := range images.([]interface{}) {
			imgs := strings.Split(imgStrs.(string), ",")
			return imgs[0], seed, nil
		}
	}
	return "", seed, errors.New("impaint img2img failed")
}

func processImage(inputImgStr string, w, h int, cords []mimg.Coordinate, prompt string) (msg string, processedImg string, seed int64, err error) {

	var pos [][]int
	if cords != nil {
		pos = make([][]int, len(cords))
		for i, pp := range cords {
			x := int(float64(w) * pp.X)
			y := int(float64(h) * pp.Y)
			pos[i] = []int{x, y}
		}
	}

	maskStr, err := detectMask(inputImgStr, w, h, pos)
	if err != nil {
		return "", "", 0, err
	}
	if maskStr == "" {
		msg = "No detected available object, Click on the image to add annotations"
		return msg, "", 0, nil
	}
	maskStr, err = expandMask(inputImgStr, maskStr, 15)
	if err != nil {
		return "", "", 0, err
	}
	if maskStr == "" {
		msg = "No available mask, Click on the image to add annotations"
		return msg, "", 0, nil
	}
	processedImg, seed, err = inpainting(inputImgStr, maskStr, prompt)
	if err != nil {
		return "", "", 0, nil
	}
	if processedImg == "" {
		msg = "Process image failded, try again later"
		return msg, "", seed, nil
	} else {
		msg = "success"
		return msg, processedImg, seed, nil
	}
}

func GetProgress() (float64, error) {
	hlog.Infof("api progress request. ")

	response, err := http.Get(SdBaseUrl + "/sdapi/v1/progress")
	hlog.Debug((response))
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	reply := make(map[string]interface{})
	err = json.NewDecoder(response.Body).Decode(&reply)
	hlog.Debug(reply)
	if err != nil {
		return 0, err
	}
	hlog.Infof("api progress request. done: ")

	/*
		{
			"progress": 0.0,
			"eta_relative": 0.0,
			"state": {
				"skipped": false,
				"interrupted": false,
				"job": "",
				"job_count": 0,
				"job_timestamp": "0",
				"job_no": 0,
				"sampling_step": 0,
				"sampling_steps": 0
			},
			"current_image": null,
			"textinfo": null
		}
	*/
	if progress, ok := reply["progress"]; ok {
		return progress.(float64), nil
	}

	return 0, errors.New("api progress requestd")
}

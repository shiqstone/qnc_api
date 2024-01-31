package queue

import (
	"encoding/json"
	"qnc/biz/model/image"
	"qnc/biz/mw/redis"
	"qnc/biz/servers"
	service "qnc/biz/service/image"
	"qnc/pkg/errno"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var sdService = make(chan struct{}, 1) // Semaphore for controlling access to SD service

func Init() {
	go ProcessQueue()
}

func ProcessQueue() {
	sdService <- struct{}{}
	q := new(redis.Queue)
	qkey := "processImage"

	for {
		// Check if queue is not empty
		res, err := q.Dequeue(qkey)
		if err == nil {
			// Check if SD service is available

			select {
			case <-sdService:
				progress, err := service.GetProgress()
				if err != nil {
					hlog.Errorf("get SD service err:", err)
					// Re-enqueue the request
					q.ReEnqueue(qkey, res)
				} else if progress > 0 {
					hlog.Debug("SD service is busy, please wait")
					// Re-enqueue the request
					q.ReEnqueue(qkey, res)
				} else {
					var msg string
					var msgType string
					if res.Type == "ud" {
						msgType = "ud"
						ipres, err := service.PostProcessImageUd(res.InputImgStr, res.W, res.H, res.Cords, res.Prompt, res.OrderId)
						if err != nil {
							ipres = new(image.ImageUdResponse)
							ipres.StatusCode = errno.SdProcessErrCode
							ipres.StatusMsg = errno.SdProcessErr.ErrMsg
							jsonData, _ := json.Marshal(ipres)
							msg = string(jsonData)
						} else {
							ipres.StatusCode = consts.StatusOK
							ipres.StatusMsg = consts.StatusMessage(consts.StatusOK)
							jsonData, _ := json.Marshal(ipres)
							msg = string(jsonData)
						}
					} else if res.Type == "tryon" {
						msgType = "tryon"
						ipres, err := service.PostProcessImageTryOn(res.InputImgStr, res.W, res.H, res.Cords, res.Prompt, res.OrderId)
						if err != nil {
							ipres = new(image.ImageTryOnResponse)
							ipres.StatusCode = errno.SdProcessErrCode
							ipres.StatusMsg = errno.SdProcessErr.ErrMsg
							jsonData, _ := json.Marshal(ipres)
							msg = string(jsonData)
						} else {
							ipres.StatusCode = consts.StatusOK
							ipres.StatusMsg = consts.StatusMessage(consts.StatusOK)
							jsonData, _ := json.Marshal(ipres)
							msg = string(jsonData)
						}
					}

					clientId := strconv.FormatInt(res.UserId, 10)

					// Notify frontend with the processed result
					servers.SendMessage2Client(clientId, "1001", msgType, errno.WS_SUCCESS, "success", &msg)

					hlog.Debug("Notify frontend with the processed result")
				}

				sdService <- struct{}{} // Release SD service
			default:
				hlog.Debug("SD service is busy, please wait")
				// Re-enqueue the request
				q.ReEnqueue(qkey, res)
			}
		} else {
			// If queue is empty, wait for a while before checking again
			time.Sleep(1 * time.Second)
		}
	}
}

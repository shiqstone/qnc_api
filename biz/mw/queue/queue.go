package queue

import (
	"encoding/json"
	"qnc/biz/model/image"
	frual "qnc/biz/mw/frugal"
	"qnc/biz/mw/redis"
	"qnc/biz/mw/viper"
	"qnc/biz/servers"
	ec2 "qnc/biz/service/aws"
	service "qnc/biz/service/image"
	"qnc/pkg/errno"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var sdService = make(chan struct{}, 1) // Semaphore for controlling access to SD service
var autoStartStop bool

func Init() {
	autoStartStop = viper.Conf.Aws.AutoStartStop
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
			select {
			case <-sdService:
				//try start ec2 instance
				if autoStartStop {
					ec2.StartInstance()
				}

				// Check if SD service is available
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
					clientId := strconv.FormatInt(res.UserId, 10)

					var msg string
					var msgType string
					if res.Type == "ud" {
						msgType, msg = postProcessImageUd(res)
					} else if res.Type == "tryon" {
						msgType, msg = postProcessImageTryOn(res)
					}
					// hlog.Debugf("post %s process return: %s", msgType, msg)

					// Notify frontend with the processed result
					servers.SendMessage2Client(clientId, "1001", msgType, errno.WS_SUCCESS, "success", &msg)
					hlog.Debug("Notify frontend with the processed result")

					// Notify to stop ec2 instance
					if autoStartStop {
						frual.Notify()
					}
				}

				sdService <- struct{}{} // Release SD service
			default:
				hlog.Debug("SD service is busy, please wait")
				// Re-enqueue the request
				q.ReEnqueue(qkey, res)
			}
		} else {
			// If queue is empty, wait for a while before checking again
			hlog.Error(err)
			time.Sleep(1 * time.Second)
		}
	}
}

func postProcessImageUd(res redis.ImageProcessRequestData) (string, string) {
	msgType := "ud"
	var msg = ""
	ipres, err := service.PostProcessImageUd(res.InputImgStr, res.W, res.H, res.Cords, res.Prompt, res.OrderId)
	if err != nil {
		ipres = new(image.ImageUdResponse)
		ipres.StatusCode = errno.SdProcessErrCode
		ipres.StatusMsg = err.Error() //errno.SdProcessErr.ErrMsg
		jsonData, _ := json.Marshal(ipres)
		msg = string(jsonData)
	} else {
		ipres.StatusCode = consts.StatusOK
		ipres.StatusMsg = consts.StatusMessage(consts.StatusOK)
		jsonData, _ := json.Marshal(ipres)
		msg = string(jsonData)
	}
	return msgType, msg
}

func postProcessImageTryOn(res redis.ImageProcessRequestData) (string, string) {
	msgType := "tryon"
	var msg = ""
	ipres, err := service.PostProcessImageTryOn(res.InputImgStr, res.W, res.H, res.Cords, res.Prompt, res.OrderId)
	if err != nil {
		ipres = new(image.ImageTryOnResponse)
		ipres.StatusCode = errno.SdProcessErrCode
		ipres.StatusMsg = err.Error() //errno.SdProcessErr.ErrMsg
		jsonData, _ := json.Marshal(ipres)
		msg = string(jsonData)
	} else {
		ipres.StatusCode = consts.StatusOK
		ipres.StatusMsg = consts.StatusMessage(consts.StatusOK)
		jsonData, _ := json.Marshal(ipres)
		msg = string(jsonData)
	}
	return msgType, msg
}

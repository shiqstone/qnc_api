package frual

import (
	"qnc/biz/mw/viper"
	ec2 "qnc/biz/service/aws"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var messages chan bool
var trigger chan bool
var dura time.Duration

func Init() {
	messages = make(chan bool)
	trigger = make(chan bool)
	dura = time.Duration(viper.Conf.Aws.IdleTime) // 2 minutes

	go manageCenter()

	go func() {
		for range trigger {
			ec2.StopInstance()
		}
	}()
}

func manageCenter() {
	timer := time.NewTimer(dura * time.Minute)
	timer.Stop() // stop timer at init, until get first message

	for {
		select {
		case <-messages:
			hlog.Debug("Received a message, reset timer.")
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			timer.Reset(dura * time.Minute)
		case <-timer.C:
			hlog.Debug("Timer ended, calling Service")
			trigger <- true
		}
	}
}

func Notify() {
	hlog.Debug("new request message.")
	messages <- true
}

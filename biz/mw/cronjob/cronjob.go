package cronjob

import (
	"qnc/biz/service/currency"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/robfig/cron/v3"
)

func Init() {
	c := cron.New()

	_, err := c.AddFunc("0 6 * * *", currency.Tracking)
	if err != nil {
		hlog.Errorf("Error scheduling task:", err)
		return
	}

	// start cron scheduler
	c.Start()

	// maintain goroutine running，no quit
	select {}
}

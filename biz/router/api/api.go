package api

import (
	"qnc/biz/handler/api"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func PayNotify(r *server.Hertz) {

	r.POST("/api/paynotify", api.Notify)
}

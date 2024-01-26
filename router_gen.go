package main

import (
	router "qnc/biz/router"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)
	router.GeneratedTopup(r)

	router.ApiPayNotify(r)

	router.GetConf(r)

	customizedRegister(r)
}

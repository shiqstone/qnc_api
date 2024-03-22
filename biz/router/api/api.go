package api

import (
	"qnc/biz/handler/api"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func PayNotify(r *server.Hertz) {

	r.POST("/api/paynotify", api.Notify)
}

func GetDepositConf(r *server.Hertz) {

	r.GET("/api/gettopupconf", api.GetTopupConf)
}

func GetDepositProds(r *server.Hertz) {

	r.GET("/api/getdepositprods", api.GetDepositProds)
}

func GetClothes(r *server.Hertz) {

	r.GET("/api/getclothes", api.GetClothes)
}

// func GetEc2Status(r *server.Hertz) {

// 	r.GET("/api/getec2status", api.GetEC2Status)
// }

// func StartEc2Instance(r *server.Hertz) {

// 	r.GET("/api/startec2", api.StartEc2Instance)
// }

// func StopEc2Instance(r *server.Hertz) {

// 	r.GET("/api/stopec2", api.StopEc2Instance)
// }

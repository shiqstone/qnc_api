// Code generated by hertz generator.

package main

import (
	"qnc/biz/dal"
	"qnc/biz/mw/jwt"
	"qnc/biz/mw/queue"
	"qnc/biz/mw/viper"
	"qnc/biz/mw/ws"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/hertz-contrib/cors"
)

func Init() {
	viper.Init()
	dal.Init()
	jwt.Init()
	queue.Init()
	ws.Init()
}

func main() {
	Init()

	// h := server.Default()
	h := server.Default(
		// server.WithHostPorts("127.0.0.1:8899"),
		// server.WithMaxRequestBodySize(10<<20),
		server.WithHostPorts(viper.Conf.App.HostPorts),
		server.WithMaxRequestBodySize(viper.Conf.App.MaxRequestBodySize),
		server.WithTransport(standard.NewTransporter),
	)

	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                       // Allowed domains, need to bring schema
		AllowMethods:     []string{"PUT", "PATCH"},            // Allowed request methods
		AllowHeaders:     []string{"Origin", "Authorization"}, // Allowed request headers
		ExposeHeaders:    []string{"Content-Length"},          // Request headers allowed in the upload_file
		AllowCredentials: true,                                // Whether cookies are attached
		AllowOriginFunc: func(origin string) bool { // Custom domain detection with lower priority than AllowOrigins
			return origin == viper.Conf.App.AllowOrigins //"https://github.com"
		},
		MaxAge: 12 * time.Hour, // Maximum length of upload_file-side cache preflash requests (seconds)
	}))

	register(h)
	h.Spin()
}

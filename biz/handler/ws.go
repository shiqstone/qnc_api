// Code generated by hertz generator.

package handler

import (
	"context"
	"qnc/biz/servers"

	"github.com/cloudwego/hertz/pkg/app"
)

// Ws .
func Ws(ctx context.Context, c *app.RequestContext) {
	servers.StartServeWs(c)

	// go servers.WriteMessage()
}
package api

import (
	"context"

	user "qnc/biz/model/user"
	service "qnc/biz/service/user"
	"qnc/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Api .
// @router //api/paynotify [POST]
func Notify(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.PayStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//TODO verify sign

	err = service.NewAccountService(ctx, c).UpdatePayStatus(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.PayStatusResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	resp := new(user.PayStatusResponse)
	resp.Result = "ok"

	c.JSON(consts.StatusOK, resp)
}

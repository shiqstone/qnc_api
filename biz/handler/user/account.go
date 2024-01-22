// Code generated by hertz generator.

package user

import (
	"context"

	user "qnc/biz/model/user"
	service "qnc/biz/service/user"
	"qnc/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AccountTopup .
// @router //account/topup/ [POST]
func AccountTopup(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.AccountTopupRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	req.Ip = c.ClientIP()
	hlog.Infof("account topup request. param:", req)
	depositId, err := service.NewAccountService(ctx, c).CreateAccountTopup(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.AccountTopupResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	resp := new(user.AccountTopupResponse)
	resp.DepositId = depositId

	c.JSON(consts.StatusOK, resp)
}

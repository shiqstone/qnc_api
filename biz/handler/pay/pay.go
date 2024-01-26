package pay

import (
	"context"
	"qnc/biz/model/pay"
	service "qnc/biz/service/user"
	"qnc/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Payment .
// @router /payment/pay/ [POST]
func Payment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pay.PaymentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	req.Ip = c.ClientIP()
	hlog.Infof("payment request. param:", req)
	depositId, err := service.NewAccountService(ctx, c).CreateAccountTopup(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, pay.PaymentResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	resp := new(pay.PaymentResponse)
	resp.DepositId = depositId

	c.JSON(consts.StatusOK, resp)
}

// @router /payment/status/ [GET]
func GetStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pay.PaymentStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := service.NewAccountService(ctx, c).GetPayStatus(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, pay.PaymentStatusResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

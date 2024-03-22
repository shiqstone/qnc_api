package api

import (
	"context"

	"qnc/biz/model/kv"
	"qnc/biz/model/pay"
	depositservice "qnc/biz/service/deposit"
	kvservice "qnc/biz/service/kv"
	uservice "qnc/biz/service/user"
	"qnc/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Api .
// @router /api/paynotify [POST]
func Notify(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pay.PayStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	//TODO verify sign

	err = uservice.NewAccountService(ctx, c).UpdatePayStatus(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, pay.PayStatusResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	resp := new(pay.PayStatusResponse)
	resp.Result = "ok"

	c.JSON(consts.StatusOK, resp)
}

// @router /api/gettopupconf [GET]
func GetTopupConf(ctx context.Context, c *app.RequestContext) {
	var err error
	resp, err := kvservice.NewKvService(ctx, c).GetDepositConf()
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, pay.PayStatusResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// @router /api/gettdepositprods [GET]
func GetDepositProds(ctx context.Context, c *app.RequestContext) {
	var err error
	var countryCode string
	countryCode = string(c.FormValue("country_code"))
	if len(countryCode) == 0 {
		countryCode = "ID"
	}

	resp, err := depositservice.NewDepositService(ctx, c).GetDepositProds(countryCode)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, kv.DepositProdsResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// @router /api/clothes [GET]
func GetClothes(ctx context.Context, c *app.RequestContext) {
	var err error
	resp, err := kvservice.NewKvService(ctx, c).GetClothes()
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, pay.PayStatusResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// func GetEC2Status(ctx context.Context, c *app.RequestContext) {
// 	status, err := service.NewEc2Service(ctx, c).GetInstanceStatus()
// 	if err != nil {
// 		resp := utils.BuildBaseResp(err)
// 		c.JSON(consts.StatusOK, pay.PayStatusResponse{
// 			StatusCode: resp.StatusCode,
// 			StatusMsg:  resp.StatusMsg,
// 		})
// 		return
// 	}
// 	c.JSON(consts.StatusOK, status)
// }

// func StartEc2Instance(ctx context.Context, c *app.RequestContext) {
// 	err := service.NewEc2Service(ctx, c).StartInstance()
// 	if err != nil {
// 		resp := utils.BuildBaseResp(err)
// 		c.JSON(consts.StatusOK, pay.PayStatusResponse{
// 			StatusCode: resp.StatusCode,
// 			StatusMsg:  resp.StatusMsg,
// 		})
// 		return
// 	}
// 	c.JSON(consts.StatusOK, nil)
// }

// func StopEc2Instance(ctx context.Context, c *app.RequestContext) {
// 	err := service.NewEc2Service(ctx, c).StopInstance()
// 	if err != nil {
// 		resp := utils.BuildBaseResp(err)
// 		c.JSON(consts.StatusOK, pay.PayStatusResponse{
// 			StatusCode: resp.StatusCode,
// 			StatusMsg:  resp.StatusMsg,
// 		})
// 		return
// 	}
// 	c.JSON(consts.StatusOK, nil)
// }

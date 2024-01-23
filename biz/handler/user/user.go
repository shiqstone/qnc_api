package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	user "qnc/biz/model/user"
	"qnc/biz/mw/jwt"
	service "qnc/biz/service/user"
	"qnc/pkg/errno"
	"qnc/pkg/utils"
)

// UserRegister user registration api
//
// @router /user/register/  [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.RegisterResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	hlog.Debugf("user register param:", &req)

	_, err = service.NewUserService(ctx, c).UserRegister(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.RegisterResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}
	jwt.JwtMiddleware.LoginHandler(ctx, c)
	token := c.GetString("token")
	v, _ := c.Get("user_id")
	userId := v.(int64)
	c.JSON(consts.StatusOK, user.RegisterResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Token:      token,
		UserId:     userId,
	})
}

// UserLogin user login api
//
// @router /user/login/  [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	v, _ := c.Get("user_id")
	userId := v.(int64)
	token := c.GetString("token")
	c.JSON(consts.StatusOK, user.LoginResponse{
		StatusCode: errno.SuccessCode,
		StatusMsg:  errno.SuccessMsg,
		Token:      token,
		UserId:     userId,
	})
}

// User get user info
// @router /user/ [GET]
func User(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.Request
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.Response{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	u, err := service.NewUserService(ctx, c).UserInfo(&req)

	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, user.Response{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		User:       u,
	})
}

// User get user info
// @router /user/identifier [POST]
func UserIdentifier(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.IdentifierRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		resp := utils.BuildBaseResp(err)
		c.JSON(consts.StatusOK, user.IdentifierResponse{
			StatusCode: resp.StatusCode,
			StatusMsg:  resp.StatusMsg,
		})
		return
	}

	res, err := service.NewUserService(ctx, c).CheckUserExisted(req.Email)

	resp := utils.BuildBaseResp(err)
	c.JSON(consts.StatusOK, user.IdentifierResponse{
		StatusCode: resp.StatusCode,
		StatusMsg:  resp.StatusMsg,
		Exsited:    res,
	})
}

package pay

import (
	"github.com/cloudwego/hertz/pkg/app"

	"qnc/biz/mw/jwt"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func __2fMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _payMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _paymentMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		jwt.JwtMiddleware.MiddlewareFunc(),
	}
}

func _payPaymentMw() []app.HandlerFunc {
	// your code...
	return nil
}

package pay

import (
	"qnc/biz/handler/pay"

	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Payment(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_pay := root.Group("/pay", _payMw()...)
		{
			_topup := _pay.Group("/payment", _paymentMw()...)
			_topup.POST("/", append(_payPaymentMw(), pay.Payment)...)
		}
		{
			_topup := _pay.Group("/status", _paymentMw()...)
			_topup.POST("/", append(_payPaymentMw(), pay.GetStatus)...)
		}
	}
}

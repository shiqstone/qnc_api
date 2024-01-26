// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.1
// source: account.proto

package user

import (
	_ "qnc/biz/model/api"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FundsRequest struct {
	UserId    int64   `json:"user_id" form:"user_id" query:"user_id"` // 登录用户名
	Amount    float64 `json:"amount" form:"amount" query:"amount"`    // 交易金额
	OrderId   int64   `json:"orderId" form:"orderId" `                // 订单号
	EventType int16   `json:"type" form:"type" `                      // 类型
	Remark    string  `json:"remark" form:"remark"`                   // 备注
}

type FundsResponse struct {
	UserId    int64   `json:"user_id" form:"user_id" query:"user_id"` // 登录用户名
	Coin      float64 `json:"coin" form:"coin" query:"coin"`          // 账户余额
	Amount    float64 `json:"amount" form:"amount" query:"amount"`    // 交易金额
	OrderId   int64   `json:"orderId" form:"orderId" `                // 订单号
	EventType int     `json:"type" form:"type" `                      // 类型
	Remark    string  `json:"remark" form:"remark"`                   // 备注
}

const (
	//increase
	TYPE_ADMIN_INCREASE   int16 = 2
	TYPE_RECHARGE         int16 = 1
	TYPE_PAYMENT_REVERSAL int16 = 3
	TYPE_UNFREEZE         int16 = 4
	TYPE_REFUND_RETURN    int16 = 5

	//deduct
	TYPE_PAYMENT        int16 = 21
	TYPE_ADMIN_DECREASE int16 = 22
	TYPE_FREEZE         int16 = 23
	TYPE_REFUND         int16 = 24
)

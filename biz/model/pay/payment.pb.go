package pay

// //支付模式
// const PAY_TYPE_GATEWAY           = 'GATEWAY_WAP';
// const PAY_TYPE_SDK               = 'MOBILE_SDK';
// const PAY_TYPE_JSAPI             = 'JSAPI';
// const PAY_TYPE_QRCODE            = 'QRCODE';

// //支付渠道
// const PAY_CHANNEL_ZFB_SDK        = 'ZFBPAY_SDK';
// const PAY_CHANNEL_WEIXIN_JSAPI   = 'WXPAY_JSAPI';

// //支付状态
const (
	PAY_STATUS_INIT    int16 = 0 //支付未处理
	PAY_STATUS_SUCCESS int16 = 1 //支付成功
	PAY_STATUS_FALID   int16 = 2 //支付失败
	PAY_STATUS_UNKNOW  int16 = 3 //未知状态,一般是过期
)

type PaymentRequest struct {
	UserId   int64   `protobuf:"varint,1,opt,name=user_id,json=user_id,proto3" json:"user_id,omitempty" form:"user_id" query:"user_id"` // 用户id
	Paytype  int32   `protobuf:"bytes,2,opt,name=paytype,proto3" json:"paytype" form:"paytype" query:"paytype" vd:"$ > 0"`
	Currency string  `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency" form:"currency" query:"currency"`
	Amount   float64 `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount" form:"amount" query:"amount" vd:"$ > 0"`
	Ip       string  `json:"ip"`
}

func (x *PaymentRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PaymentRequest) GetPaytype() int32 {
	if x != nil {
		return x.Paytype
	}
	return 0
}

func (x *PaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PaymentResponse struct {
	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
	Uid        int64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty" form:"uid" query:"uid"`                            // user id
	Balance    float64 `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty" form:"balance" query:"balance"`           // user balance
	DepositId  string  `protobuf:"bytes,5,opt,name=deposit_id,proto3" json:"deposit_id,omitempty" form:"deposit_id" query:"deposit_id"` // deposit id
}

func (x *PaymentResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PaymentResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *PaymentResponse) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PaymentResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type PayStatusRequest struct {
	DepositId  string  `json:"deposit_id" form:"deposit_id" vd:"len($)>0"`   // 充值记录单号
	Amount     float64 `json:"amount" form:"amount" query:"amount" vd:"$>0"` // 交易金额
	Result     string  `json:"result" form:"result" vd:"len($)>0"`           // 支付结果
	PayId      string  `json:"pay_id" form:"pay_id" vd:"len($)>0"`           // 订单号
	PayChannel string  `json:"pay_channel" form:"pay_channel" `              // 渠道
	Remark     string  `json:"remark" form:"remark"`                         // 备注
}

type PayStatusResponse struct {
	DepositId string `json:"deposit_id"`
	Status    int16  `json:"status"`
	Result    string `json:"result" form:"result"` // 结果

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
}

type PaymentStatusRequest struct {
	DepositId string `json:"deposit_id" form:"deposit_id" vd:"len($)>0"`
}

type PaymentStatusResponse struct {
	DepositId string  `json:"deposit_id"`
	Status    int16   `json:"status"`
	Balance   float64 `json:"balance"`
	Result    string  `json:"result" form:"result"` // 结果

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
}

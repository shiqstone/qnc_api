package order

import (
	_ "qnc/biz/model/api"
)

const (
	STATUS_INIT    int16 = 0 //未处理
	STATUS_PAYED   int16 = 1 //已支付，待处理
	STATUS_SUCCESS int16 = 2 //成功
	STATUS_FALID   int16 = 3 //失败
	STATUS_UNKNOW  int16 = 4 //未知状态
)

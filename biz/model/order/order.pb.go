package order

import (
	_ "qnc/biz/model/api"
)

const (
	STATUS_INIT    int16 = 0 //未处理
	STATUS_SUCCESS int16 = 1 //成功
	STATUS_FALID   int16 = 2 //失败
	STATUS_UNKNOW  int16 = 3 //未知状态
)

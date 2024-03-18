package kv

import (
	_ "qnc/biz/model/api"
)

type DepositConfRequest struct {
	CountryCode string `json:"contry_code,omitempty"`
}

type DepositConfResponse struct {
	StatusCode int32                    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg  string                   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
	Products   []map[string]interface{} `protobuf:"bytes,3,opt,name=products,json=products,proto3" json:"products,omitempty" form:"products" query:"products"`
	Payways    []map[string]interface{} `protobuf:"fixed64,4,opt,name=payways,proto3" json:"payways,omitempty" form:"payways" query:"payways"`
	Tips       string                   `protobuf:"bytes,2,opt,name=tips,json=tips,proto3" json:"tips,omitempty" form:"tips" query:"tips"`
}

type ClothesResponse struct {
	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
	Clothes    []string `json:"clothes,omitempty"`
}

type DepositProdsResponse struct {
	StatusCode int32                    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg  string                   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
	Products   []map[string]interface{} `protobuf:"bytes,3,opt,name=products,json=products,proto3" json:"products,omitempty" form:"products" query:"products"`
}

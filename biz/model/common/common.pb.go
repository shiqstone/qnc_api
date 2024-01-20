package common

import (
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`                            // user id
	Name   string  `protobuf:"bytes,2,opt,name=user_name,proto3" json:"user_name,omitempty" form:"user_name" query:"user_name"` // user name
	Email  string  `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" form:"email" query:"email"`                 // email
	Avatar string  `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty" form:"avatar" query:"avatar_url"`         // user avatar URL
	Coin   float64 `protobuf:"bytes,8,opt,name=coin,proto3" json:"coin" form:"coin" query:"coin"`                               // user coin
}

func (*User) ProtoMessage() {}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

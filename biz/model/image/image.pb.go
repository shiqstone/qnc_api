package image

import (
	"mime/multipart"
	_ "qnc/biz/model/api"
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImageUdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64                `form:"user_id"`
	FileHeader multipart.FileHeader `form:"file"`
	Pos        string               `form:"pos"`
}

func (x *ImageUdRequest) Reset() {
	*x = ImageUdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUdRequest) ProtoMessage() {}

func (x *ImageUdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUdRequest.ProtoReflect.Descriptor instead.
func (*ImageUdRequest) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{0}
}

func (x *ImageUdRequest) GetUserid() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ImageUdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode     int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg      string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
	ProcessedImage string  `protobuf:"bytes,3,opt,name=processed_image,json=processedImage,proto3" json:"processed_image,omitempty" form:"processed_image" query:"processed_image"` // processed image base64
	Balance        float64 `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty" form:"balance" query:"balance"`                                                   // user balance
}

func (x *ImageUdResponse) Reset() {
	*x = ImageUdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageUdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageUdResponse) ProtoMessage() {}

func (x *ImageUdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageUdResponse.ProtoReflect.Descriptor instead.
func (*ImageUdResponse) Descriptor() ([]byte, []int) {
	return file_image_proto_rawDescGZIP(), []int{1}
}

func (x *ImageUdResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ImageUdResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ImageUdResponse) GetProcessedImage() string {
	if x != nil {
		return x.ProcessedImage
	}
	return ""
}

func (x *ImageUdResponse) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_image_proto protoreflect.FileDescriptor

var file_image_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x0e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x0f,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x16, 0xca, 0xf3, 0x18, 0x12, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x51, 0x0a, 0x05, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x64, 0x12, 0x15,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0xd2,
	0xc1, 0x18, 0x0a, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2f, 0x75, 0x64, 0x2f, 0x42, 0x15, 0x5a,
	0x13, 0x71, 0x6e, 0x63, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_proto_rawDescOnce sync.Once
	file_image_proto_rawDescData = file_image_proto_rawDesc
)

func file_image_proto_rawDescGZIP() []byte {
	file_image_proto_rawDescOnce.Do(func() {
		file_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_proto_rawDescData)
	})
	return file_image_proto_rawDescData
}

var file_image_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_image_proto_goTypes = []interface{}{
	(*ImageUdRequest)(nil),  // 0: image.ImageUdRequest
	(*ImageUdResponse)(nil), // 1: image.ImageUdResponse
}
var file_image_proto_depIdxs = []int32{
	0, // 0: image.Image.ImageUd:input_type -> image.ImageUdRequest
	1, // 1: image.Image.ImageUd:output_type -> image.ImageUdResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_image_proto_init() }
func file_image_proto_init() {
	if File_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageUdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_image_proto_goTypes,
		DependencyIndexes: file_image_proto_depIdxs,
		MessageInfos:      file_image_proto_msgTypes,
	}.Build()
	File_image_proto = out.File
	file_image_proto_rawDesc = nil
	file_image_proto_goTypes = nil
	file_image_proto_depIdxs = nil
}

type ImageTryOnRequest struct {
	UserId     int64                `form:"user_id"`
	FileHeader multipart.FileHeader `form:"file"`
	Pos        string               `form:"pos"`
	Cloth      string               `form:"cloth"`
}

type ImageTryOnResponse struct {
	StatusCode     int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" query:"status_code" form:"status_code" json:"status_code"`
	StatusMsg      string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`
	ProcessedImage string  `protobuf:"bytes,3,opt,name=processed_image,json=processedImage,proto3" json:"processed_image,omitempty" form:"processed_image" query:"processed_image"` // processed image base64
	Balance        float64 `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty" form:"balance" query:"balance"`                                                   // user balance
}

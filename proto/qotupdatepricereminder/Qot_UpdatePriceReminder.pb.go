// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Qot_UpdatePriceReminder.proto

package qotupdatepricereminder

import (
	_ "github.com/xfzen/futsdk/proto/common"
	qotcommon "github.com/xfzen/futsdk/proto/qotcommon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MarketStatus int32

const (
	MarketStatus_MarketStatus_Unknow  MarketStatus = 0
	MarketStatus_MarketStatus_Open    MarketStatus = 1 // 盘中
	MarketStatus_MarketStatus_USPre   MarketStatus = 2 // 美股盘前
	MarketStatus_MarketStatus_USAfter MarketStatus = 3 // 美股盘后
)

// Enum value maps for MarketStatus.
var (
	MarketStatus_name = map[int32]string{
		0: "MarketStatus_Unknow",
		1: "MarketStatus_Open",
		2: "MarketStatus_USPre",
		3: "MarketStatus_USAfter",
	}
	MarketStatus_value = map[string]int32{
		"MarketStatus_Unknow":  0,
		"MarketStatus_Open":    1,
		"MarketStatus_USPre":   2,
		"MarketStatus_USAfter": 3,
	}
)

func (x MarketStatus) Enum() *MarketStatus {
	p := new(MarketStatus)
	*p = x
	return p
}

func (x MarketStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MarketStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_UpdatePriceReminder_proto_enumTypes[0].Descriptor()
}

func (MarketStatus) Type() protoreflect.EnumType {
	return &file_Qot_UpdatePriceReminder_proto_enumTypes[0]
}

func (x MarketStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *MarketStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = MarketStatus(num)
	return nil
}

// Deprecated: Use MarketStatus.Descriptor instead.
func (MarketStatus) EnumDescriptor() ([]byte, []int) {
	return file_Qot_UpdatePriceReminder_proto_rawDescGZIP(), []int{0}
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security     *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`          //股票
	Price        *float64            `protobuf:"fixed64,2,req,name=price" json:"price,omitempty"`              //价格
	ChangeRate   *float64            `protobuf:"fixed64,3,req,name=changeRate" json:"changeRate,omitempty"`    //涨跌幅
	MarketStatus *int32              `protobuf:"varint,4,req,name=marketStatus" json:"marketStatus,omitempty"` //市场状态
	Content      *string             `protobuf:"bytes,5,req,name=content" json:"content,omitempty"`            //内容
	Note         *string             `protobuf:"bytes,6,req,name=note" json:"note,omitempty"`                  //备注
	Key          *int64              `protobuf:"varint,7,opt,name=key" json:"key,omitempty"`                   // 到价提醒的标识
	Type         *int32              `protobuf:"varint,8,opt,name=type" json:"type,omitempty"`                 // Qot_Common::PriceReminderType，提醒类型
	SetValue     *float64            `protobuf:"fixed64,9,opt,name=setValue" json:"setValue,omitempty"`        // 设置的提醒值
	CurValue     *float64            `protobuf:"fixed64,10,opt,name=curValue" json:"curValue,omitempty"`       // 设置的提醒类型触发时当前值
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_UpdatePriceReminder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_UpdatePriceReminder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_Qot_UpdatePriceReminder_proto_rawDescGZIP(), []int{0}
}

func (x *S2C) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *S2C) GetPrice() float64 {
	if x != nil && x.Price != nil {
		return *x.Price
	}
	return 0
}

func (x *S2C) GetChangeRate() float64 {
	if x != nil && x.ChangeRate != nil {
		return *x.ChangeRate
	}
	return 0
}

func (x *S2C) GetMarketStatus() int32 {
	if x != nil && x.MarketStatus != nil {
		return *x.MarketStatus
	}
	return 0
}

func (x *S2C) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *S2C) GetNote() string {
	if x != nil && x.Note != nil {
		return *x.Note
	}
	return ""
}

func (x *S2C) GetKey() int64 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *S2C) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *S2C) GetSetValue() float64 {
	if x != nil && x.SetValue != nil {
		return *x.SetValue
	}
	return 0
}

func (x *S2C) GetCurValue() float64 {
	if x != nil && x.CurValue != nil {
		return *x.CurValue
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_UpdatePriceReminder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_UpdatePriceReminder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_Qot_UpdatePriceReminder_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_Qot_UpdatePriceReminder_proto protoreflect.FileDescriptor

var file_Qot_UpdatePriceReminder_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x51, 0x6f, 0x74, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x51, 0x6f, 0x74, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x03, 0x53, 0x32, 0x43,
	0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0c,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x73, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53,
	0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a, 0x70, 0x0a, 0x0c, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x4f, 0x70, 0x65, 0x6e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x55, 0x53, 0x50, 0x72, 0x65, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x55, 0x53, 0x41, 0x66, 0x74, 0x65, 0x72, 0x10, 0x03, 0x42, 0x51, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72,
	0x69, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x6f, 0x74, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
}

var (
	file_Qot_UpdatePriceReminder_proto_rawDescOnce sync.Once
	file_Qot_UpdatePriceReminder_proto_rawDescData = file_Qot_UpdatePriceReminder_proto_rawDesc
)

func file_Qot_UpdatePriceReminder_proto_rawDescGZIP() []byte {
	file_Qot_UpdatePriceReminder_proto_rawDescOnce.Do(func() {
		file_Qot_UpdatePriceReminder_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_UpdatePriceReminder_proto_rawDescData)
	})
	return file_Qot_UpdatePriceReminder_proto_rawDescData
}

var file_Qot_UpdatePriceReminder_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Qot_UpdatePriceReminder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Qot_UpdatePriceReminder_proto_goTypes = []interface{}{
	(MarketStatus)(0),          // 0: Qot_UpdatePriceReminder.MarketStatus
	(*S2C)(nil),                // 1: Qot_UpdatePriceReminder.S2C
	(*Response)(nil),           // 2: Qot_UpdatePriceReminder.Response
	(*qotcommon.Security)(nil), // 3: Qot_Common.Security
}
var file_Qot_UpdatePriceReminder_proto_depIdxs = []int32{
	3, // 0: Qot_UpdatePriceReminder.S2C.security:type_name -> Qot_Common.Security
	1, // 1: Qot_UpdatePriceReminder.Response.s2c:type_name -> Qot_UpdatePriceReminder.S2C
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Qot_UpdatePriceReminder_proto_init() }
func file_Qot_UpdatePriceReminder_proto_init() {
	if File_Qot_UpdatePriceReminder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_UpdatePriceReminder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
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
		file_Qot_UpdatePriceReminder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_Qot_UpdatePriceReminder_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_UpdatePriceReminder_proto_goTypes,
		DependencyIndexes: file_Qot_UpdatePriceReminder_proto_depIdxs,
		EnumInfos:         file_Qot_UpdatePriceReminder_proto_enumTypes,
		MessageInfos:      file_Qot_UpdatePriceReminder_proto_msgTypes,
	}.Build()
	File_Qot_UpdatePriceReminder_proto = out.File
	file_Qot_UpdatePriceReminder_proto_rawDesc = nil
	file_Qot_UpdatePriceReminder_proto_goTypes = nil
	file_Qot_UpdatePriceReminder_proto_depIdxs = nil
}

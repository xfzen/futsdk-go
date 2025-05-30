// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Qot_UpdateOrderBook.proto

package qotupdateorderbook

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

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security                *qotcommon.Security    `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`                                 //股票
	OrderBookAskList        []*qotcommon.OrderBook `protobuf:"bytes,2,rep,name=orderBookAskList" json:"orderBookAskList,omitempty"`                 //卖盘
	OrderBookBidList        []*qotcommon.OrderBook `protobuf:"bytes,3,rep,name=orderBookBidList" json:"orderBookBidList,omitempty"`                 //买盘
	SvrRecvTimeBid          *string                `protobuf:"bytes,4,opt,name=svrRecvTimeBid" json:"svrRecvTimeBid,omitempty"`                     // 富途服务器从交易所收到数据的时间(for bid)部分数据的接收时间为零，例如服务器重启或第一次推送的缓存数据。该字段暂时只支持港股。
	SvrRecvTimeBidTimestamp *float64               `protobuf:"fixed64,5,opt,name=svrRecvTimeBidTimestamp" json:"svrRecvTimeBidTimestamp,omitempty"` // 富途服务器从交易所收到数据的时间戳(for bid)
	SvrRecvTimeAsk          *string                `protobuf:"bytes,6,opt,name=svrRecvTimeAsk" json:"svrRecvTimeAsk,omitempty"`                     // 富途服务器从交易所收到数据的时间(for ask)
	SvrRecvTimeAskTimestamp *float64               `protobuf:"fixed64,7,opt,name=svrRecvTimeAskTimestamp" json:"svrRecvTimeAskTimestamp,omitempty"` // 富途服务器从交易所收到数据的时间戳(for ask)
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_UpdateOrderBook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_UpdateOrderBook_proto_msgTypes[0]
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
	return file_Qot_UpdateOrderBook_proto_rawDescGZIP(), []int{0}
}

func (x *S2C) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *S2C) GetOrderBookAskList() []*qotcommon.OrderBook {
	if x != nil {
		return x.OrderBookAskList
	}
	return nil
}

func (x *S2C) GetOrderBookBidList() []*qotcommon.OrderBook {
	if x != nil {
		return x.OrderBookBidList
	}
	return nil
}

func (x *S2C) GetSvrRecvTimeBid() string {
	if x != nil && x.SvrRecvTimeBid != nil {
		return *x.SvrRecvTimeBid
	}
	return ""
}

func (x *S2C) GetSvrRecvTimeBidTimestamp() float64 {
	if x != nil && x.SvrRecvTimeBidTimestamp != nil {
		return *x.SvrRecvTimeBidTimestamp
	}
	return 0
}

func (x *S2C) GetSvrRecvTimeAsk() string {
	if x != nil && x.SvrRecvTimeAsk != nil {
		return *x.SvrRecvTimeAsk
	}
	return ""
}

func (x *S2C) GetSvrRecvTimeAskTimestamp() float64 {
	if x != nil && x.SvrRecvTimeAskTimestamp != nil {
		return *x.SvrRecvTimeAskTimestamp
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
		mi := &file_Qot_UpdateOrderBook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_UpdateOrderBook_proto_msgTypes[1]
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
	return file_Qot_UpdateOrderBook_proto_rawDescGZIP(), []int{1}
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

var File_Qot_UpdateOrderBook_proto protoreflect.FileDescriptor

var file_Qot_UpdateOrderBook_proto_rawDesc = []byte{
	0x0a, 0x19, 0x51, 0x6f, 0x74, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x51, 0x6f, 0x74,
	0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b,
	0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x81, 0x03, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x10, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x10, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a,
	0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x69, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x10,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x0e, 0x73, 0x76, 0x72, 0x52, 0x65, 0x63, 0x76, 0x54, 0x69, 0x6d, 0x65, 0x42,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x76, 0x72, 0x52, 0x65, 0x63,
	0x76, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x76, 0x72, 0x52,
	0x65, 0x63, 0x76, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x73, 0x76, 0x72, 0x52, 0x65,
	0x63, 0x76, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x76, 0x72, 0x52, 0x65, 0x63, 0x76, 0x54, 0x69, 0x6d,
	0x65, 0x41, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x76, 0x72, 0x52,
	0x65, 0x63, 0x76, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x73, 0x6b, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x76,
	0x72, 0x52, 0x65, 0x63, 0x76, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x73, 0x6b, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x73, 0x76, 0x72,
	0x52, 0x65, 0x63, 0x76, 0x54, 0x69, 0x6d, 0x65, 0x41, 0x73, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x42,
	0x4d, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72, 0x69, 0x73, 0x68, 0x65, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d,
	0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x6f, 0x74, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x6f, 0x6b,
}

var (
	file_Qot_UpdateOrderBook_proto_rawDescOnce sync.Once
	file_Qot_UpdateOrderBook_proto_rawDescData = file_Qot_UpdateOrderBook_proto_rawDesc
)

func file_Qot_UpdateOrderBook_proto_rawDescGZIP() []byte {
	file_Qot_UpdateOrderBook_proto_rawDescOnce.Do(func() {
		file_Qot_UpdateOrderBook_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_UpdateOrderBook_proto_rawDescData)
	})
	return file_Qot_UpdateOrderBook_proto_rawDescData
}

var file_Qot_UpdateOrderBook_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Qot_UpdateOrderBook_proto_goTypes = []interface{}{
	(*S2C)(nil),                 // 0: Qot_UpdateOrderBook.S2C
	(*Response)(nil),            // 1: Qot_UpdateOrderBook.Response
	(*qotcommon.Security)(nil),  // 2: Qot_Common.Security
	(*qotcommon.OrderBook)(nil), // 3: Qot_Common.OrderBook
}
var file_Qot_UpdateOrderBook_proto_depIdxs = []int32{
	2, // 0: Qot_UpdateOrderBook.S2C.security:type_name -> Qot_Common.Security
	3, // 1: Qot_UpdateOrderBook.S2C.orderBookAskList:type_name -> Qot_Common.OrderBook
	3, // 2: Qot_UpdateOrderBook.S2C.orderBookBidList:type_name -> Qot_Common.OrderBook
	0, // 3: Qot_UpdateOrderBook.Response.s2c:type_name -> Qot_UpdateOrderBook.S2C
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_Qot_UpdateOrderBook_proto_init() }
func file_Qot_UpdateOrderBook_proto_init() {
	if File_Qot_UpdateOrderBook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_UpdateOrderBook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_UpdateOrderBook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Qot_UpdateOrderBook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_UpdateOrderBook_proto_goTypes,
		DependencyIndexes: file_Qot_UpdateOrderBook_proto_depIdxs,
		MessageInfos:      file_Qot_UpdateOrderBook_proto_msgTypes,
	}.Build()
	File_Qot_UpdateOrderBook_proto = out.File
	file_Qot_UpdateOrderBook_proto_rawDesc = nil
	file_Qot_UpdateOrderBook_proto_goTypes = nil
	file_Qot_UpdateOrderBook_proto_depIdxs = nil
}

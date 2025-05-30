// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: Qot_GetHistoryKLPoints.proto

package qotgethistoryklpoints

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

//当请求时间点数据为空时，如何返回数据
type NoDataMode int32

const (
	NoDataMode_NoDataMode_Null     NoDataMode = 0 //直接返回空数据
	NoDataMode_NoDataMode_Forward  NoDataMode = 1 //往前取值，返回前一个时间点数据
	NoDataMode_NoDataMode_Backward NoDataMode = 2 //向后取值，返回后一个时间点数据
)

// Enum value maps for NoDataMode.
var (
	NoDataMode_name = map[int32]string{
		0: "NoDataMode_Null",
		1: "NoDataMode_Forward",
		2: "NoDataMode_Backward",
	}
	NoDataMode_value = map[string]int32{
		"NoDataMode_Null":     0,
		"NoDataMode_Forward":  1,
		"NoDataMode_Backward": 2,
	}
)

func (x NoDataMode) Enum() *NoDataMode {
	p := new(NoDataMode)
	*p = x
	return p
}

func (x NoDataMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NoDataMode) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_GetHistoryKLPoints_proto_enumTypes[0].Descriptor()
}

func (NoDataMode) Type() protoreflect.EnumType {
	return &file_Qot_GetHistoryKLPoints_proto_enumTypes[0]
}

func (x NoDataMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *NoDataMode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = NoDataMode(num)
	return nil
}

// Deprecated: Use NoDataMode.Descriptor instead.
func (NoDataMode) EnumDescriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKLPoints_proto_rawDescGZIP(), []int{0}
}

//这个时间点返回数据的状态以及来源
type DataStatus int32

const (
	DataStatus_DataStatus_Null     DataStatus = 0 //空数据
	DataStatus_DataStatus_Current  DataStatus = 1 //当前时间点数据
	DataStatus_DataStatus_Previous DataStatus = 2 //前一个时间点数据
	DataStatus_DataStatus_Back     DataStatus = 3 //后一个时间点数据
)

// Enum value maps for DataStatus.
var (
	DataStatus_name = map[int32]string{
		0: "DataStatus_Null",
		1: "DataStatus_Current",
		2: "DataStatus_Previous",
		3: "DataStatus_Back",
	}
	DataStatus_value = map[string]int32{
		"DataStatus_Null":     0,
		"DataStatus_Current":  1,
		"DataStatus_Previous": 2,
		"DataStatus_Back":     3,
	}
)

func (x DataStatus) Enum() *DataStatus {
	p := new(DataStatus)
	*p = x
	return p
}

func (x DataStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_Qot_GetHistoryKLPoints_proto_enumTypes[1].Descriptor()
}

func (DataStatus) Type() protoreflect.EnumType {
	return &file_Qot_GetHistoryKLPoints_proto_enumTypes[1]
}

func (x DataStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DataStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DataStatus(num)
	return nil
}

// Deprecated: Use DataStatus.Descriptor instead.
func (DataStatus) EnumDescriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKLPoints_proto_rawDescGZIP(), []int{1}
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RehabType         *int32                `protobuf:"varint,1,req,name=rehabType" json:"rehabType,omitempty"`                 //Qot_Common.RehabType,复权类型
	KlType            *int32                `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`                       //Qot_Common.KLType,K线类型
	NoDataMode        *int32                `protobuf:"varint,3,req,name=noDataMode" json:"noDataMode,omitempty"`               //NoDataMode,当请求时间点数据为空时，如何返回数据
	SecurityList      []*qotcommon.Security `protobuf:"bytes,4,rep,name=securityList" json:"securityList,omitempty"`            //股票市场以及股票代码
	TimeList          []string              `protobuf:"bytes,5,rep,name=timeList" json:"timeList,omitempty"`                    //时间字符串
	MaxReqSecurityNum *int32                `protobuf:"varint,6,opt,name=maxReqSecurityNum" json:"maxReqSecurityNum,omitempty"` //最多返回多少只股票的数据，如果未指定表示不限制
	NeedKLFieldsFlag  *int64                `protobuf:"varint,7,opt,name=needKLFieldsFlag" json:"needKLFieldsFlag,omitempty"`   //指定返回K线结构体特定某几项数据，KLFields枚举值或组合，如果未指定返回全部字段
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKLPoints_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetRehabType() int32 {
	if x != nil && x.RehabType != nil {
		return *x.RehabType
	}
	return 0
}

func (x *C2S) GetKlType() int32 {
	if x != nil && x.KlType != nil {
		return *x.KlType
	}
	return 0
}

func (x *C2S) GetNoDataMode() int32 {
	if x != nil && x.NoDataMode != nil {
		return *x.NoDataMode
	}
	return 0
}

func (x *C2S) GetSecurityList() []*qotcommon.Security {
	if x != nil {
		return x.SecurityList
	}
	return nil
}

func (x *C2S) GetTimeList() []string {
	if x != nil {
		return x.TimeList
	}
	return nil
}

func (x *C2S) GetMaxReqSecurityNum() int32 {
	if x != nil && x.MaxReqSecurityNum != nil {
		return *x.MaxReqSecurityNum
	}
	return 0
}

func (x *C2S) GetNeedKLFieldsFlag() int64 {
	if x != nil && x.NeedKLFieldsFlag != nil {
		return *x.NeedKLFieldsFlag
	}
	return 0
}

type HistoryPointsKL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  *int32           `protobuf:"varint,1,req,name=status" json:"status,omitempty"`  //DataStatus,数据状态
	ReqTime *string          `protobuf:"bytes,2,req,name=reqTime" json:"reqTime,omitempty"` //请求的时间
	Kl      *qotcommon.KLine `protobuf:"bytes,3,req,name=kl" json:"kl,omitempty"`           //K线数据
}

func (x *HistoryPointsKL) Reset() {
	*x = HistoryPointsKL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryPointsKL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryPointsKL) ProtoMessage() {}

func (x *HistoryPointsKL) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryPointsKL.ProtoReflect.Descriptor instead.
func (*HistoryPointsKL) Descriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKLPoints_proto_rawDescGZIP(), []int{1}
}

func (x *HistoryPointsKL) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *HistoryPointsKL) GetReqTime() string {
	if x != nil && x.ReqTime != nil {
		return *x.ReqTime
	}
	return ""
}

func (x *HistoryPointsKL) GetKl() *qotcommon.KLine {
	if x != nil {
		return x.Kl
	}
	return nil
}

type SecurityHistoryKLPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Security *qotcommon.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"` //股票
	KlList   []*HistoryPointsKL  `protobuf:"bytes,2,rep,name=klList" json:"klList,omitempty"`     //K线数据
}

func (x *SecurityHistoryKLPoints) Reset() {
	*x = SecurityHistoryKLPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityHistoryKLPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityHistoryKLPoints) ProtoMessage() {}

func (x *SecurityHistoryKLPoints) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityHistoryKLPoints.ProtoReflect.Descriptor instead.
func (*SecurityHistoryKLPoints) Descriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKLPoints_proto_rawDescGZIP(), []int{2}
}

func (x *SecurityHistoryKLPoints) GetSecurity() *qotcommon.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *SecurityHistoryKLPoints) GetKlList() []*HistoryPointsKL {
	if x != nil {
		return x.KlList
	}
	return nil
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KlPointList []*SecurityHistoryKLPoints `protobuf:"bytes,1,rep,name=klPointList" json:"klPointList,omitempty"` //多只股票的多点历史K线点
	HasNext     *bool                      `protobuf:"varint,2,opt,name=hasNext" json:"hasNext,omitempty"`        //如请求不指定maxReqSecurityNum值，则不会返回该字段，该字段表示请求是否还有超过指定限制的数据
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[3]
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
	return file_Qot_GetHistoryKLPoints_proto_rawDescGZIP(), []int{3}
}

func (x *S2C) GetKlPointList() []*SecurityHistoryKLPoints {
	if x != nil {
		return x.KlPointList
	}
	return nil
}

func (x *S2C) GetHasNext() bool {
	if x != nil && x.HasNext != nil {
		return *x.HasNext
	}
	return false
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_Qot_GetHistoryKLPoints_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
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
		mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_GetHistoryKLPoints_proto_msgTypes[5]
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
	return file_Qot_GetHistoryKLPoints_proto_rawDescGZIP(), []int{5}
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

var File_Qot_GetHistoryKLPoints_proto protoreflect.FileDescriptor

var file_Qot_GetHistoryKLPoints_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4b, 0x4c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x68, 0x61, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x09, 0x72, 0x65, 0x68, 0x61, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6b, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x6b, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x6f, 0x44, 0x61, 0x74, 0x61,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61,
	0x78, 0x52, 0x65, 0x71, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x65, 0x65, 0x64,
	0x4b, 0x4c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x6e, 0x65, 0x65, 0x64, 0x4b, 0x4c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x46, 0x6c, 0x61, 0x67, 0x22, 0x66, 0x0a, 0x0f, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x4b, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x02, 0x6b, 0x6c, 0x18,
	0x03, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4b, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x02, 0x6b, 0x6c, 0x22, 0x8c, 0x01, 0x0a,
	0x17, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4b, 0x4c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x06, 0x6b, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x51, 0x6f, 0x74,
	0x5f, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x4b, 0x4c, 0x52, 0x06, 0x6b, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x72, 0x0a, 0x03, 0x53,
	0x32, 0x43, 0x12, 0x51, 0x0a, 0x0b, 0x6b, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x4b, 0x4c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x6b, 0x6c, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x22,
	0x38, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x03, 0x63, 0x32,
	0x73, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x4c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x53,
	0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a, 0x52, 0x0a, 0x0a, 0x4e, 0x6f, 0x44, 0x61, 0x74,
	0x61, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4d,
	0x6f, 0x64, 0x65, 0x5f, 0x4e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x6f,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x5f, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65,
	0x5f, 0x42, 0x61, 0x63, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x10, 0x02, 0x2a, 0x67, 0x0a, 0x0a, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x4e, 0x75, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x16,
	0x0a, 0x12, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x10, 0x02, 0x12,
	0x13, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x42, 0x61,
	0x63, 0x6b, 0x10, 0x03, 0x42, 0x50, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x72, 0x69, 0x73, 0x68, 0x65, 0x6e,
	0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x66, 0x75, 0x74, 0x75, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62,
	0x2f, 0x71, 0x6f, 0x74, 0x67, 0x65, 0x74, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x6b, 0x6c,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
}

var (
	file_Qot_GetHistoryKLPoints_proto_rawDescOnce sync.Once
	file_Qot_GetHistoryKLPoints_proto_rawDescData = file_Qot_GetHistoryKLPoints_proto_rawDesc
)

func file_Qot_GetHistoryKLPoints_proto_rawDescGZIP() []byte {
	file_Qot_GetHistoryKLPoints_proto_rawDescOnce.Do(func() {
		file_Qot_GetHistoryKLPoints_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_GetHistoryKLPoints_proto_rawDescData)
	})
	return file_Qot_GetHistoryKLPoints_proto_rawDescData
}

var file_Qot_GetHistoryKLPoints_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Qot_GetHistoryKLPoints_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Qot_GetHistoryKLPoints_proto_goTypes = []interface{}{
	(NoDataMode)(0),                 // 0: Qot_GetHistoryKLPoints.NoDataMode
	(DataStatus)(0),                 // 1: Qot_GetHistoryKLPoints.DataStatus
	(*C2S)(nil),                     // 2: Qot_GetHistoryKLPoints.C2S
	(*HistoryPointsKL)(nil),         // 3: Qot_GetHistoryKLPoints.HistoryPointsKL
	(*SecurityHistoryKLPoints)(nil), // 4: Qot_GetHistoryKLPoints.SecurityHistoryKLPoints
	(*S2C)(nil),                     // 5: Qot_GetHistoryKLPoints.S2C
	(*Request)(nil),                 // 6: Qot_GetHistoryKLPoints.Request
	(*Response)(nil),                // 7: Qot_GetHistoryKLPoints.Response
	(*qotcommon.Security)(nil),      // 8: Qot_Common.Security
	(*qotcommon.KLine)(nil),         // 9: Qot_Common.KLine
}
var file_Qot_GetHistoryKLPoints_proto_depIdxs = []int32{
	8, // 0: Qot_GetHistoryKLPoints.C2S.securityList:type_name -> Qot_Common.Security
	9, // 1: Qot_GetHistoryKLPoints.HistoryPointsKL.kl:type_name -> Qot_Common.KLine
	8, // 2: Qot_GetHistoryKLPoints.SecurityHistoryKLPoints.security:type_name -> Qot_Common.Security
	3, // 3: Qot_GetHistoryKLPoints.SecurityHistoryKLPoints.klList:type_name -> Qot_GetHistoryKLPoints.HistoryPointsKL
	4, // 4: Qot_GetHistoryKLPoints.S2C.klPointList:type_name -> Qot_GetHistoryKLPoints.SecurityHistoryKLPoints
	2, // 5: Qot_GetHistoryKLPoints.Request.c2s:type_name -> Qot_GetHistoryKLPoints.C2S
	5, // 6: Qot_GetHistoryKLPoints.Response.s2c:type_name -> Qot_GetHistoryKLPoints.S2C
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_Qot_GetHistoryKLPoints_proto_init() }
func file_Qot_GetHistoryKLPoints_proto_init() {
	if File_Qot_GetHistoryKLPoints_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_GetHistoryKLPoints_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
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
		file_Qot_GetHistoryKLPoints_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryPointsKL); i {
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
		file_Qot_GetHistoryKLPoints_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityHistoryKLPoints); i {
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
		file_Qot_GetHistoryKLPoints_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_GetHistoryKLPoints_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_Qot_GetHistoryKLPoints_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Qot_GetHistoryKLPoints_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_GetHistoryKLPoints_proto_goTypes,
		DependencyIndexes: file_Qot_GetHistoryKLPoints_proto_depIdxs,
		EnumInfos:         file_Qot_GetHistoryKLPoints_proto_enumTypes,
		MessageInfos:      file_Qot_GetHistoryKLPoints_proto_msgTypes,
	}.Build()
	File_Qot_GetHistoryKLPoints_proto = out.File
	file_Qot_GetHistoryKLPoints_proto_rawDesc = nil
	file_Qot_GetHistoryKLPoints_proto_goTypes = nil
	file_Qot_GetHistoryKLPoints_proto_depIdxs = nil
}

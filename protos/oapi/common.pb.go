// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package oapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 钉钉开放平台OAPI接口通用响应参数
type OpenApiResponse struct {
	// 返回码
	Errcode int64 `protobuf:"varint,1,opt,name=errcode,proto3" json:"errcode,omitempty"`
	// 对返回码的文本描述内容
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenApiResponse) Reset()         { *m = OpenApiResponse{} }
func (m *OpenApiResponse) String() string { return proto.CompactTextString(m) }
func (*OpenApiResponse) ProtoMessage()    {}
func (*OpenApiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_20a5d84bb07f427a, []int{0}
}
func (m *OpenApiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenApiResponse.Unmarshal(m, b)
}
func (m *OpenApiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenApiResponse.Marshal(b, m, deterministic)
}
func (dst *OpenApiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenApiResponse.Merge(dst, src)
}
func (m *OpenApiResponse) XXX_Size() int {
	return xxx_messageInfo_OpenApiResponse.Size(m)
}
func (m *OpenApiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenApiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenApiResponse proto.InternalMessageInfo

func (m *OpenApiResponse) GetErrcode() int64 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *OpenApiResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

// 获取钉钉配置信息（JSAPI）请求参数
type DingTalkConfigRequest struct {
	// 需要鉴权的地址
	TargetUrl            string   `protobuf:"bytes,1,opt,name=targetUrl,proto3" json:"targetUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DingTalkConfigRequest) Reset()         { *m = DingTalkConfigRequest{} }
func (m *DingTalkConfigRequest) String() string { return proto.CompactTextString(m) }
func (*DingTalkConfigRequest) ProtoMessage()    {}
func (*DingTalkConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_20a5d84bb07f427a, []int{1}
}
func (m *DingTalkConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DingTalkConfigRequest.Unmarshal(m, b)
}
func (m *DingTalkConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DingTalkConfigRequest.Marshal(b, m, deterministic)
}
func (dst *DingTalkConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DingTalkConfigRequest.Merge(dst, src)
}
func (m *DingTalkConfigRequest) XXX_Size() int {
	return xxx_messageInfo_DingTalkConfigRequest.Size(m)
}
func (m *DingTalkConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DingTalkConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DingTalkConfigRequest proto.InternalMessageInfo

func (m *DingTalkConfigRequest) GetTargetUrl() string {
	if m != nil {
		return m.TargetUrl
	}
	return ""
}

// 获取钉钉配置信息（JSAPI）相应参数
type DingTalkConfigResponse struct {
	// 当前页面url
	TargetUrl string `protobuf:"bytes,1,opt,name=targetUrl,proto3" json:"targetUrl,omitempty"`
	// 应用的标识
	AgentId string `protobuf:"bytes,2,opt,name=agentId,proto3" json:"agentId,omitempty"`
	// 企业Id
	CorpId string `protobuf:"bytes,3,opt,name=corpId,proto3" json:"corpId,omitempty"`
	// 时间戳
	TimeStamp string `protobuf:"bytes,4,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	// 随机数
	NonceStr string `protobuf:"bytes,5,opt,name=nonceStr,proto3" json:"nonceStr,omitempty"`
	// jsapi_ticket
	Ticket string `protobuf:"bytes,6,opt,name=ticket,proto3" json:"ticket,omitempty"`
	// 签名
	Signature            string   `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DingTalkConfigResponse) Reset()         { *m = DingTalkConfigResponse{} }
func (m *DingTalkConfigResponse) String() string { return proto.CompactTextString(m) }
func (*DingTalkConfigResponse) ProtoMessage()    {}
func (*DingTalkConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_20a5d84bb07f427a, []int{2}
}
func (m *DingTalkConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DingTalkConfigResponse.Unmarshal(m, b)
}
func (m *DingTalkConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DingTalkConfigResponse.Marshal(b, m, deterministic)
}
func (dst *DingTalkConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DingTalkConfigResponse.Merge(dst, src)
}
func (m *DingTalkConfigResponse) XXX_Size() int {
	return xxx_messageInfo_DingTalkConfigResponse.Size(m)
}
func (m *DingTalkConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DingTalkConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DingTalkConfigResponse proto.InternalMessageInfo

func (m *DingTalkConfigResponse) GetTargetUrl() string {
	if m != nil {
		return m.TargetUrl
	}
	return ""
}

func (m *DingTalkConfigResponse) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *DingTalkConfigResponse) GetCorpId() string {
	if m != nil {
		return m.CorpId
	}
	return ""
}

func (m *DingTalkConfigResponse) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

func (m *DingTalkConfigResponse) GetNonceStr() string {
	if m != nil {
		return m.NonceStr
	}
	return ""
}

func (m *DingTalkConfigResponse) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

func (m *DingTalkConfigResponse) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*OpenApiResponse)(nil), "oapi.OpenApiResponse")
	proto.RegisterType((*DingTalkConfigRequest)(nil), "oapi.DingTalkConfigRequest")
	proto.RegisterType((*DingTalkConfigResponse)(nil), "oapi.DingTalkConfigResponse")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_common_20a5d84bb07f427a) }

var fileDescriptor_common_20a5d84bb07f427a = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x5b, 0x53, 0xb3, 0x2a, 0x42, 0xc0, 0xb2, 0x88, 0x42, 0xc9, 0xa9, 0xa7, 0x7a,
	0x10, 0x1f, 0x40, 0xeb, 0xa5, 0xa7, 0x42, 0xaa, 0x17, 0x6f, 0x6b, 0x32, 0x2e, 0x4b, 0x9b, 0x9d,
	0x75, 0x76, 0xfa, 0x52, 0x3e, 0x91, 0x8f, 0x23, 0xbb, 0x9b, 0x54, 0xf0, 0xe0, 0xf1, 0xfb, 0x87,
	0xf9, 0xf8, 0x67, 0xc4, 0x79, 0x83, 0x5d, 0x87, 0x76, 0xe9, 0x08, 0x19, 0xcb, 0x09, 0x2a, 0x67,
	0xaa, 0x95, 0xb8, 0xdc, 0x38, 0xb0, 0x8f, 0xce, 0xd4, 0xe0, 0x1d, 0x5a, 0x0f, 0xa5, 0x14, 0x53,
	0x20, 0x6a, 0xb0, 0x05, 0x99, 0xcd, 0xb3, 0xc5, 0xb8, 0x1e, 0xb0, 0x9c, 0x89, 0x1c, 0x88, 0x3a,
	0xaf, 0xe5, 0x68, 0x9e, 0x2d, 0x8a, 0xba, 0xa7, 0xea, 0x41, 0x5c, 0x3d, 0x1b, 0xab, 0x5f, 0xd4,
	0x7e, 0xb7, 0x42, 0xfb, 0x61, 0x74, 0x0d, 0x9f, 0x07, 0xf0, 0x5c, 0xde, 0x88, 0x82, 0x15, 0x69,
	0xe0, 0x57, 0xda, 0x47, 0x59, 0x51, 0xff, 0x06, 0xd5, 0x77, 0x26, 0x66, 0x7f, 0xf7, 0xfa, 0x0e,
	0xff, 0x2e, 0x86, 0x86, 0x4a, 0x83, 0xe5, 0x75, 0xdb, 0x17, 0x19, 0x30, 0x34, 0x6c, 0x90, 0xdc,
	0xba, 0x95, 0xe3, 0xd4, 0x30, 0x51, 0xf4, 0x99, 0x0e, 0xb6, 0xac, 0x3a, 0x27, 0x27, 0xbd, 0x6f,
	0x08, 0xca, 0x6b, 0x71, 0x6a, 0xd1, 0x36, 0xb0, 0x65, 0x92, 0x27, 0x71, 0x78, 0xe4, 0x60, 0x64,
	0xd3, 0xec, 0x80, 0x65, 0x9e, 0x8c, 0x89, 0x82, 0xd1, 0x1b, 0x6d, 0x15, 0x1f, 0x08, 0xe4, 0x34,
	0x19, 0x8f, 0xc1, 0xd3, 0xed, 0xdb, 0x59, 0xfc, 0xb2, 0xbf, 0x0b, 0x5f, 0xfe, 0x1a, 0x5d, 0x0c,
	0x77, 0x2e, 0x37, 0xca, 0x99, 0xf7, 0x3c, 0x0e, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb,
	0x65, 0x96, 0xb4, 0x92, 0x01, 0x00, 0x00,
}

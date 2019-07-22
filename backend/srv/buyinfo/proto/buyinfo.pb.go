// Code generated by protoc-gen-go. DO NOT EDIT.
// source: buyinfo.proto

package buyinfo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BuyInfoCreateResponse_Status int32

const (
	BuyInfoCreateResponse_UNKNOWN       BuyInfoCreateResponse_Status = 0
	BuyInfoCreateResponse_INVALID_PARAM BuyInfoCreateResponse_Status = -1
	BuyInfoCreateResponse_SUCCESS       BuyInfoCreateResponse_Status = 1
	BuyInfoCreateResponse_INVALID_TOKEN BuyInfoCreateResponse_Status = 2
)

var BuyInfoCreateResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "INVALID_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_TOKEN",
}

var BuyInfoCreateResponse_Status_value = map[string]int32{
	"UNKNOWN":       0,
	"INVALID_PARAM": -1,
	"SUCCESS":       1,
	"INVALID_TOKEN": 2,
}

func (x BuyInfoCreateResponse_Status) String() string {
	return proto.EnumName(BuyInfoCreateResponse_Status_name, int32(x))
}

func (BuyInfoCreateResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e8250a94d2d8d95, []int{3, 0}
}

type BuyInfoMsg struct {
	BuyInfoId            int32    `protobuf:"varint,1,opt,name=buyInfoId,proto3" json:"buyInfoId,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	ReleaseTime          int64    `protobuf:"varint,3,opt,name=releaseTime,proto3" json:"releaseTime,omitempty"`
	ValidTime            int64    `protobuf:"varint,4,opt,name=validTime,proto3" json:"validTime,omitempty"`
	GoodName             string   `protobuf:"bytes,5,opt,name=goodName,proto3" json:"goodName,omitempty"`
	Price                float64  `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	ContentId            string   `protobuf:"bytes,8,opt,name=contentId,proto3" json:"contentId,omitempty"`
	UserId               int32    `protobuf:"varint,9,opt,name=userId,proto3" json:"userId,omitempty"`
	ClearEmpty           bool     `protobuf:"varint,99,opt,name=clearEmpty,proto3" json:"clearEmpty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyInfoMsg) Reset()         { *m = BuyInfoMsg{} }
func (m *BuyInfoMsg) String() string { return proto.CompactTextString(m) }
func (*BuyInfoMsg) ProtoMessage()    {}
func (*BuyInfoMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8250a94d2d8d95, []int{0}
}

func (m *BuyInfoMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyInfoMsg.Unmarshal(m, b)
}
func (m *BuyInfoMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyInfoMsg.Marshal(b, m, deterministic)
}
func (m *BuyInfoMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyInfoMsg.Merge(m, src)
}
func (m *BuyInfoMsg) XXX_Size() int {
	return xxx_messageInfo_BuyInfoMsg.Size(m)
}
func (m *BuyInfoMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyInfoMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BuyInfoMsg proto.InternalMessageInfo

func (m *BuyInfoMsg) GetBuyInfoId() int32 {
	if m != nil {
		return m.BuyInfoId
	}
	return 0
}

func (m *BuyInfoMsg) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BuyInfoMsg) GetReleaseTime() int64 {
	if m != nil {
		return m.ReleaseTime
	}
	return 0
}

func (m *BuyInfoMsg) GetValidTime() int64 {
	if m != nil {
		return m.ValidTime
	}
	return 0
}

func (m *BuyInfoMsg) GetGoodName() string {
	if m != nil {
		return m.GoodName
	}
	return ""
}

func (m *BuyInfoMsg) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *BuyInfoMsg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BuyInfoMsg) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *BuyInfoMsg) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *BuyInfoMsg) GetClearEmpty() bool {
	if m != nil {
		return m.ClearEmpty
	}
	return false
}

type BuyInfoQueryRequest struct {
	BuyInfoId            int32    `protobuf:"varint,1,opt,name=buyInfoId,proto3" json:"buyInfoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyInfoQueryRequest) Reset()         { *m = BuyInfoQueryRequest{} }
func (m *BuyInfoQueryRequest) String() string { return proto.CompactTextString(m) }
func (*BuyInfoQueryRequest) ProtoMessage()    {}
func (*BuyInfoQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8250a94d2d8d95, []int{1}
}

func (m *BuyInfoQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyInfoQueryRequest.Unmarshal(m, b)
}
func (m *BuyInfoQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyInfoQueryRequest.Marshal(b, m, deterministic)
}
func (m *BuyInfoQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyInfoQueryRequest.Merge(m, src)
}
func (m *BuyInfoQueryRequest) XXX_Size() int {
	return xxx_messageInfo_BuyInfoQueryRequest.Size(m)
}
func (m *BuyInfoQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyInfoQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuyInfoQueryRequest proto.InternalMessageInfo

func (m *BuyInfoQueryRequest) GetBuyInfoId() int32 {
	if m != nil {
		return m.BuyInfoId
	}
	return 0
}

type BuyInfoCreateRequest struct {
	ValidTime            int64    `protobuf:"varint,1,opt,name=validTime,proto3" json:"validTime,omitempty"`
	GoodName             string   `protobuf:"bytes,2,opt,name=goodName,proto3" json:"goodName,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ContentId            string   `protobuf:"bytes,5,opt,name=contentId,proto3" json:"contentId,omitempty"`
	ContentToken         string   `protobuf:"bytes,6,opt,name=contentToken,proto3" json:"contentToken,omitempty"`
	UserId               int32    `protobuf:"varint,7,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyInfoCreateRequest) Reset()         { *m = BuyInfoCreateRequest{} }
func (m *BuyInfoCreateRequest) String() string { return proto.CompactTextString(m) }
func (*BuyInfoCreateRequest) ProtoMessage()    {}
func (*BuyInfoCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8250a94d2d8d95, []int{2}
}

func (m *BuyInfoCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyInfoCreateRequest.Unmarshal(m, b)
}
func (m *BuyInfoCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyInfoCreateRequest.Marshal(b, m, deterministic)
}
func (m *BuyInfoCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyInfoCreateRequest.Merge(m, src)
}
func (m *BuyInfoCreateRequest) XXX_Size() int {
	return xxx_messageInfo_BuyInfoCreateRequest.Size(m)
}
func (m *BuyInfoCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyInfoCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuyInfoCreateRequest proto.InternalMessageInfo

func (m *BuyInfoCreateRequest) GetValidTime() int64 {
	if m != nil {
		return m.ValidTime
	}
	return 0
}

func (m *BuyInfoCreateRequest) GetGoodName() string {
	if m != nil {
		return m.GoodName
	}
	return ""
}

func (m *BuyInfoCreateRequest) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *BuyInfoCreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BuyInfoCreateRequest) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *BuyInfoCreateRequest) GetContentToken() string {
	if m != nil {
		return m.ContentToken
	}
	return ""
}

func (m *BuyInfoCreateRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type BuyInfoCreateResponse struct {
	Status               BuyInfoCreateResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=BuyInfoCreateResponse_Status" json:"status,omitempty"`
	BuyInfoId            int32                        `protobuf:"varint,2,opt,name=buyInfoId,proto3" json:"buyInfoId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BuyInfoCreateResponse) Reset()         { *m = BuyInfoCreateResponse{} }
func (m *BuyInfoCreateResponse) String() string { return proto.CompactTextString(m) }
func (*BuyInfoCreateResponse) ProtoMessage()    {}
func (*BuyInfoCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8250a94d2d8d95, []int{3}
}

func (m *BuyInfoCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyInfoCreateResponse.Unmarshal(m, b)
}
func (m *BuyInfoCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyInfoCreateResponse.Marshal(b, m, deterministic)
}
func (m *BuyInfoCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyInfoCreateResponse.Merge(m, src)
}
func (m *BuyInfoCreateResponse) XXX_Size() int {
	return xxx_messageInfo_BuyInfoCreateResponse.Size(m)
}
func (m *BuyInfoCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyInfoCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuyInfoCreateResponse proto.InternalMessageInfo

func (m *BuyInfoCreateResponse) GetStatus() BuyInfoCreateResponse_Status {
	if m != nil {
		return m.Status
	}
	return BuyInfoCreateResponse_UNKNOWN
}

func (m *BuyInfoCreateResponse) GetBuyInfoId() int32 {
	if m != nil {
		return m.BuyInfoId
	}
	return 0
}

type BuyInfoFindRequest struct {
	UserId               int32    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint32   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyInfoFindRequest) Reset()         { *m = BuyInfoFindRequest{} }
func (m *BuyInfoFindRequest) String() string { return proto.CompactTextString(m) }
func (*BuyInfoFindRequest) ProtoMessage()    {}
func (*BuyInfoFindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8250a94d2d8d95, []int{4}
}

func (m *BuyInfoFindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyInfoFindRequest.Unmarshal(m, b)
}
func (m *BuyInfoFindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyInfoFindRequest.Marshal(b, m, deterministic)
}
func (m *BuyInfoFindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyInfoFindRequest.Merge(m, src)
}
func (m *BuyInfoFindRequest) XXX_Size() int {
	return xxx_messageInfo_BuyInfoFindRequest.Size(m)
}
func (m *BuyInfoFindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyInfoFindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuyInfoFindRequest proto.InternalMessageInfo

func (m *BuyInfoFindRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *BuyInfoFindRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *BuyInfoFindRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type BuyInfoFindResponse struct {
	BuyInfo              []*BuyInfoMsg `protobuf:"bytes,1,rep,name=buyInfo,proto3" json:"buyInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BuyInfoFindResponse) Reset()         { *m = BuyInfoFindResponse{} }
func (m *BuyInfoFindResponse) String() string { return proto.CompactTextString(m) }
func (*BuyInfoFindResponse) ProtoMessage()    {}
func (*BuyInfoFindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e8250a94d2d8d95, []int{5}
}

func (m *BuyInfoFindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyInfoFindResponse.Unmarshal(m, b)
}
func (m *BuyInfoFindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyInfoFindResponse.Marshal(b, m, deterministic)
}
func (m *BuyInfoFindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyInfoFindResponse.Merge(m, src)
}
func (m *BuyInfoFindResponse) XXX_Size() int {
	return xxx_messageInfo_BuyInfoFindResponse.Size(m)
}
func (m *BuyInfoFindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyInfoFindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuyInfoFindResponse proto.InternalMessageInfo

func (m *BuyInfoFindResponse) GetBuyInfo() []*BuyInfoMsg {
	if m != nil {
		return m.BuyInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("BuyInfoCreateResponse_Status", BuyInfoCreateResponse_Status_name, BuyInfoCreateResponse_Status_value)
	proto.RegisterType((*BuyInfoMsg)(nil), "BuyInfoMsg")
	proto.RegisterType((*BuyInfoQueryRequest)(nil), "BuyInfoQueryRequest")
	proto.RegisterType((*BuyInfoCreateRequest)(nil), "BuyInfoCreateRequest")
	proto.RegisterType((*BuyInfoCreateResponse)(nil), "BuyInfoCreateResponse")
	proto.RegisterType((*BuyInfoFindRequest)(nil), "BuyInfoFindRequest")
	proto.RegisterType((*BuyInfoFindResponse)(nil), "BuyInfoFindResponse")
}

func init() { proto.RegisterFile("buyinfo.proto", fileDescriptor_0e8250a94d2d8d95) }

var fileDescriptor_0e8250a94d2d8d95 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x9d, 0x76, 0xd3, 0x3f, 0xa7, 0x56, 0xea, 0x6c, 0x77, 0x19, 0x82, 0x4a, 0x08, 0x08,
	0xc1, 0x8b, 0x80, 0x5d, 0xc4, 0x1b, 0x6f, 0x6a, 0xad, 0x10, 0xd6, 0xcd, 0xba, 0xd3, 0xae, 0x82,
	0x37, 0x92, 0x26, 0xd3, 0x25, 0xd8, 0x66, 0x62, 0x66, 0x22, 0xf4, 0x71, 0xbc, 0xf7, 0x1d, 0x7c,
	0x17, 0x5f, 0x44, 0xc9, 0x24, 0x6d, 0xd2, 0x1a, 0x6b, 0xef, 0xbe, 0xef, 0x9c, 0xf9, 0x73, 0x7e,
	0xf3, 0x35, 0xd0, 0x5f, 0xa4, 0x9b, 0x30, 0x5a, 0x72, 0x3b, 0x4e, 0xb8, 0xe4, 0xe6, 0x8f, 0x06,
	0xc0, 0xeb, 0x74, 0xe3, 0x44, 0x4b, 0x7e, 0x25, 0xee, 0xf0, 0x23, 0xe8, 0x2e, 0x72, 0xe5, 0x04,
	0x04, 0x19, 0xc8, 0xd2, 0x68, 0x69, 0xe0, 0x73, 0x68, 0x09, 0xe9, 0xc9, 0x54, 0x90, 0x86, 0x2a,
	0x15, 0x0a, 0x1b, 0xd0, 0x4b, 0xd8, 0x8a, 0x79, 0x82, 0xcd, 0xc3, 0x35, 0x23, 0x4d, 0x03, 0x59,
	0x4d, 0x5a, 0xb5, 0xb2, 0x7d, 0xbf, 0x79, 0xab, 0x30, 0x50, 0xf5, 0x13, 0x55, 0x2f, 0x0d, 0xac,
	0x43, 0xe7, 0x8e, 0xf3, 0xc0, 0xf5, 0xd6, 0x8c, 0x68, 0x06, 0xb2, 0xba, 0x74, 0xa7, 0xf1, 0x10,
	0xb4, 0x38, 0x09, 0x7d, 0x46, 0x5a, 0x06, 0xb2, 0x10, 0xcd, 0x45, 0x76, 0x62, 0xc0, 0x84, 0x9f,
	0x84, 0xb1, 0x0c, 0x79, 0x44, 0xda, 0x6a, 0x51, 0xd5, 0xca, 0x4e, 0xf4, 0x79, 0x24, 0x59, 0x24,
	0x9d, 0x80, 0x74, 0x54, 0xbd, 0x34, 0xb2, 0x49, 0x52, 0xc1, 0x12, 0x27, 0x20, 0xdd, 0x7c, 0x92,
	0x5c, 0xe1, 0x27, 0x00, 0xfe, 0x8a, 0x79, 0xc9, 0x74, 0x1d, 0xcb, 0x0d, 0xf1, 0x0d, 0x64, 0x75,
	0x68, 0xc5, 0x31, 0x2f, 0xe0, 0xb4, 0xa0, 0x75, 0x93, 0xb2, 0x64, 0x43, 0xd9, 0xd7, 0x94, 0x09,
	0x79, 0x1c, 0x9b, 0xf9, 0x0b, 0xc1, 0xb0, 0x58, 0x35, 0x49, 0x98, 0x27, 0x59, 0x65, 0x59, 0x49,
	0x05, 0x1d, 0xa3, 0xd2, 0xf8, 0x17, 0x95, 0xe6, 0x11, 0x2a, 0x27, 0xff, 0xa1, 0xa2, 0x1d, 0x52,
	0x31, 0xe1, 0x7e, 0x21, 0xe6, 0xfc, 0x0b, 0x8b, 0x14, 0xf2, 0x2e, 0xdd, 0xf3, 0x2a, 0xe4, 0xda,
	0x55, 0x72, 0xe6, 0x4f, 0x04, 0x67, 0x07, 0x43, 0x8a, 0x98, 0x47, 0x82, 0xe1, 0x17, 0xbb, 0xd4,
	0x64, 0x23, 0x3e, 0x18, 0x3d, 0xb6, 0x6b, 0xfb, 0xec, 0x99, 0x6a, 0xda, 0x85, 0x6a, 0x8f, 0x69,
	0xe3, 0x90, 0xe9, 0x0d, 0xb4, 0xf2, 0x7e, 0xdc, 0x83, 0xf6, 0xad, 0x7b, 0xe9, 0x5e, 0x7f, 0x74,
	0x07, 0xf7, 0xb0, 0x0e, 0x7d, 0xc7, 0xfd, 0x30, 0x7e, 0xe7, 0xbc, 0xf9, 0xfc, 0x7e, 0x4c, 0xc7,
	0x57, 0x83, 0xdf, 0xdb, 0x1f, 0xca, 0x1a, 0x67, 0xb7, 0x93, 0xc9, 0x74, 0x36, 0x1b, 0x20, 0xfc,
	0xb0, 0x6c, 0x9c, 0x5f, 0x5f, 0x4e, 0xdd, 0x41, 0xc3, 0xfc, 0x04, 0xb8, 0xb8, 0xd8, 0xdb, 0x30,
	0x0a, 0xb6, 0x6f, 0x54, 0xce, 0x8b, 0xf6, 0x92, 0x32, 0x04, 0x6d, 0x15, 0xae, 0x43, 0xa9, 0xae,
	0xd6, 0xa7, 0xb9, 0xc8, 0xba, 0xf9, 0x72, 0x29, 0x98, 0x54, 0x0f, 0xd3, 0xa7, 0x85, 0x32, 0x5f,
	0xed, 0x72, 0x93, 0xef, 0x5d, 0xa0, 0x79, 0x0a, 0xed, 0x62, 0x24, 0x82, 0x8c, 0xa6, 0xd5, 0x1b,
	0xf5, 0xec, 0xf2, 0xcf, 0x48, 0xb7, 0xb5, 0xd1, 0x77, 0x04, 0xed, 0xc2, 0xc7, 0xcf, 0x40, 0x53,
	0xd1, 0xc3, 0x43, 0xbb, 0x26, 0x89, 0x7a, 0x75, 0x03, 0xfc, 0x12, 0x5a, 0x39, 0x63, 0x7c, 0x66,
	0xd7, 0x05, 0x50, 0x3f, 0xaf, 0x7f, 0x0a, 0xfc, 0x1c, 0x4e, 0xb2, 0x7b, 0xe2, 0x53, 0xfb, 0x6f,
	0x22, 0xfa, 0xd0, 0xae, 0x19, 0x65, 0xd1, 0x52, 0xdf, 0x93, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x31, 0xdc, 0x9e, 0x60, 0x04, 0x00, 0x00,
}
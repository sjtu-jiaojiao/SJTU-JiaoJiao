// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

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

type AuthResponse_Status int32

const (
	AuthResponse_UNKNOWN      AuthResponse_Status = 0
	AuthResponse_EMPTY_PARAM  AuthResponse_Status = -1
	AuthResponse_SUCCESS      AuthResponse_Status = 1
	AuthResponse_INVALID_CODE AuthResponse_Status = 2
)

var AuthResponse_Status_name = map[int32]string{
	0:  "UNKNOWN",
	-1: "EMPTY_PARAM",
	1:  "SUCCESS",
	2:  "INVALID_CODE",
}

var AuthResponse_Status_value = map[string]int32{
	"UNKNOWN":      0,
	"EMPTY_PARAM":  -1,
	"SUCCESS":      1,
	"INVALID_CODE": 2,
}

func (x AuthResponse_Status) String() string {
	return proto.EnumName(AuthResponse_Status_name, int32(x))
}

func (AuthResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1, 0}
}

type AuthRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type AuthResponse struct {
	Status               AuthResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=AuthResponse_Status" json:"status,omitempty"`
	Token                string              `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	StudentId            string              `protobuf:"bytes,3,opt,name=studentId,proto3" json:"studentId,omitempty"`
	StudentName          string              `protobuf:"bytes,4,opt,name=studentName,proto3" json:"studentName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetStatus() AuthResponse_Status {
	if m != nil {
		return m.Status
	}
	return AuthResponse_UNKNOWN
}

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResponse) GetStudentId() string {
	if m != nil {
		return m.StudentId
	}
	return ""
}

func (m *AuthResponse) GetStudentName() string {
	if m != nil {
		return m.StudentName
	}
	return ""
}

func init() {
	proto.RegisterEnum("AuthResponse_Status", AuthResponse_Status_name, AuthResponse_Status_value)
	proto.RegisterType((*AuthRequest)(nil), "AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "AuthResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4, 0xe2, 0x76, 0x2c, 0x2d, 0xc9, 0x08, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x1e, 0x31, 0x72, 0xf1, 0x40, 0xd4, 0x14, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x0a, 0xe9, 0x70, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x83, 0x95, 0xf1,
	0x19, 0x89, 0xe8, 0x21, 0x4b, 0xeb, 0x05, 0x83, 0xe5, 0x82, 0xa0, 0x6a, 0x84, 0x44, 0xb8, 0x58,
	0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0x24, 0x98, 0xc0, 0x66, 0x42, 0x38, 0x42, 0x32, 0x5c, 0x9c, 0xc5,
	0x25, 0xa5, 0x29, 0xa9, 0x79, 0x25, 0x9e, 0x29, 0x12, 0xcc, 0x60, 0x19, 0x84, 0x80, 0x90, 0x02,
	0x17, 0x37, 0x94, 0xe3, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x02, 0x96, 0x47, 0x16, 0x52, 0xf2, 0xe3,
	0x62, 0x83, 0xd8, 0x23, 0xc4, 0xcd, 0xc5, 0x1e, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0,
	0x20, 0x24, 0xc1, 0xc5, 0xed, 0xea, 0x1b, 0x10, 0x12, 0x19, 0x1f, 0xe0, 0x18, 0xe4, 0xe8, 0x2b,
	0xf0, 0x1f, 0x06, 0x18, 0x41, 0xca, 0x82, 0x43, 0x9d, 0x9d, 0x5d, 0x83, 0x83, 0x05, 0x18, 0x85,
	0x04, 0xb8, 0x78, 0x3c, 0xfd, 0xc2, 0x1c, 0x7d, 0x3c, 0x5d, 0xe2, 0x9d, 0xfd, 0x5d, 0x5c, 0x05,
	0x98, 0x8c, 0xb4, 0xb9, 0x58, 0x40, 0x9e, 0x10, 0x52, 0x86, 0xd2, 0x3c, 0x7a, 0x48, 0xc1, 0x22,
	0xc5, 0x8b, 0xe2, 0xc3, 0x24, 0x36, 0x70, 0xd8, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0x6d, 0x71, 0xfe, 0x49, 0x01, 0x00, 0x00,
}

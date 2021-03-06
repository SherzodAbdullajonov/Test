// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/server1.proto

package pb

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

type DownloadPostsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadPostsRequest) Reset()         { *m = DownloadPostsRequest{} }
func (m *DownloadPostsRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadPostsRequest) ProtoMessage()    {}
func (*DownloadPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b20ccc079411be, []int{0}
}

func (m *DownloadPostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadPostsRequest.Unmarshal(m, b)
}
func (m *DownloadPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadPostsRequest.Marshal(b, m, deterministic)
}
func (m *DownloadPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadPostsRequest.Merge(m, src)
}
func (m *DownloadPostsRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadPostsRequest.Size(m)
}
func (m *DownloadPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadPostsRequest proto.InternalMessageInfo

type DownloadPostsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadPostsResponse) Reset()         { *m = DownloadPostsResponse{} }
func (m *DownloadPostsResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadPostsResponse) ProtoMessage()    {}
func (*DownloadPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b20ccc079411be, []int{1}
}

func (m *DownloadPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadPostsResponse.Unmarshal(m, b)
}
func (m *DownloadPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadPostsResponse.Marshal(b, m, deterministic)
}
func (m *DownloadPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadPostsResponse.Merge(m, src)
}
func (m *DownloadPostsResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadPostsResponse.Size(m)
}
func (m *DownloadPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadPostsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DownloadPostsRequest)(nil), "data.DownloadPostsRequest")
	proto.RegisterType((*DownloadPostsResponse)(nil), "data.DownloadPostsResponse")
}

func init() { proto.RegisterFile("pb/server1.proto", fileDescriptor_82b20ccc079411be) }

var fileDescriptor_82b20ccc079411be = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0x2f,
	0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x32, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x49,
	0x2c, 0x49, 0x54, 0x12, 0xe3, 0x12, 0x71, 0xc9, 0x2f, 0xcf, 0xcb, 0xc9, 0x4f, 0x4c, 0x09, 0xc8,
	0x2f, 0x2e, 0x29, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x12, 0xe7, 0x12, 0x45, 0x13,
	0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0x0a, 0xe7, 0xe2, 0x76, 0x49, 0x2c, 0x49, 0x0c, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xf2, 0xe0, 0xe2, 0x45, 0x51, 0x27, 0x24, 0xa5, 0x07, 0x32,
	0x57, 0x0f, 0x9b, 0xa1, 0x52, 0xd2, 0x58, 0xe5, 0x20, 0x06, 0x3b, 0x19, 0x46, 0xe9, 0xa7, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x07, 0x67, 0xa4, 0x16, 0x55, 0xe5, 0xa7,
	0x38, 0x26, 0xa5, 0x94, 0xe6, 0xe4, 0x24, 0x66, 0xe5, 0xe7, 0xe5, 0x97, 0xe9, 0x97, 0xa4, 0x16,
	0x97, 0x80, 0x3d, 0x91, 0x99, 0x9c, 0x6a, 0xa8, 0x5f, 0x90, 0x94, 0xc4, 0x06, 0xf6, 0x89, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xa6, 0xc2, 0xde, 0xdd, 0x00, 0x00, 0x00,
}

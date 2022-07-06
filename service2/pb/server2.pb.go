// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/server2.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type Post struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97a181bc8e5e3bf, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type GetPostRequest struct {
	PostId               int64    `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostRequest) Reset()         { *m = GetPostRequest{} }
func (m *GetPostRequest) String() string { return proto.CompactTextString(m) }
func (*GetPostRequest) ProtoMessage()    {}
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97a181bc8e5e3bf, []int{1}
}

func (m *GetPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPostRequest.Unmarshal(m, b)
}
func (m *GetPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPostRequest.Marshal(b, m, deterministic)
}
func (m *GetPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostRequest.Merge(m, src)
}
func (m *GetPostRequest) XXX_Size() int {
	return xxx_messageInfo_GetPostRequest.Size(m)
}
func (m *GetPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostRequest proto.InternalMessageInfo

func (m *GetPostRequest) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type DeletePostRequest struct {
	PostId               int64    `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRequest) Reset()         { *m = DeletePostRequest{} }
func (m *DeletePostRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePostRequest) ProtoMessage()    {}
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97a181bc8e5e3bf, []int{2}
}

func (m *DeletePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRequest.Unmarshal(m, b)
}
func (m *DeletePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRequest.Marshal(b, m, deterministic)
}
func (m *DeletePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRequest.Merge(m, src)
}
func (m *DeletePostRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePostRequest.Size(m)
}
func (m *DeletePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRequest proto.InternalMessageInfo

func (m *DeletePostRequest) GetPostId() int64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

type ListPostsRequest struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsRequest) Reset()         { *m = ListPostsRequest{} }
func (m *ListPostsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPostsRequest) ProtoMessage()    {}
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97a181bc8e5e3bf, []int{3}
}

func (m *ListPostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsRequest.Unmarshal(m, b)
}
func (m *ListPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsRequest.Marshal(b, m, deterministic)
}
func (m *ListPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsRequest.Merge(m, src)
}
func (m *ListPostsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPostsRequest.Size(m)
}
func (m *ListPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsRequest proto.InternalMessageInfo

func (m *ListPostsRequest) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListPostsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ListPostsResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostsResponse) Reset()         { *m = ListPostsResponse{} }
func (m *ListPostsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPostsResponse) ProtoMessage()    {}
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97a181bc8e5e3bf, []int{4}
}

func (m *ListPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostsResponse.Unmarshal(m, b)
}
func (m *ListPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostsResponse.Marshal(b, m, deterministic)
}
func (m *ListPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostsResponse.Merge(m, src)
}
func (m *ListPostsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPostsResponse.Size(m)
}
func (m *ListPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostsResponse proto.InternalMessageInfo

func (m *ListPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "post.Post")
	proto.RegisterType((*GetPostRequest)(nil), "post.GetPostRequest")
	proto.RegisterType((*DeletePostRequest)(nil), "post.DeletePostRequest")
	proto.RegisterType((*ListPostsRequest)(nil), "post.ListPostsRequest")
	proto.RegisterType((*ListPostsResponse)(nil), "post.ListPostsResponse")
}

func init() { proto.RegisterFile("pb/server2.proto", fileDescriptor_e97a181bc8e5e3bf) }

var fileDescriptor_e97a181bc8e5e3bf = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x6f, 0xda, 0x30,
	0x14, 0x56, 0x20, 0x80, 0x78, 0x48, 0x08, 0x2c, 0x04, 0x51, 0x76, 0x89, 0x72, 0x62, 0x1b, 0x4a,
	0xa6, 0x4c, 0xbb, 0xb1, 0xc3, 0xa6, 0x4d, 0x13, 0xd2, 0x0e, 0x15, 0xa8, 0x87, 0xf6, 0x82, 0x08,
	0x7e, 0x0d, 0xae, 0x02, 0x76, 0x63, 0x07, 0x89, 0xfe, 0xe3, 0xfe, 0x8b, 0xca, 0x76, 0xa0, 0xb4,
	0xa8, 0x52, 0x6f, 0xfe, 0x9e, 0xbf, 0xf7, 0xf9, 0xfb, 0xde, 0x33, 0xf4, 0x44, 0x1a, 0x4b, 0x2c,
	0xf6, 0x58, 0x24, 0x91, 0x28, 0xb8, 0xe2, 0xc4, 0x15, 0x5c, 0x2a, 0xff, 0x53, 0xc6, 0x79, 0x96,
	0x63, 0x6c, 0x6a, 0x69, 0x79, 0x17, 0xe3, 0x56, 0xa8, 0x83, 0xa5, 0x84, 0x37, 0xe0, 0x5e, 0x71,
	0xa9, 0x48, 0x17, 0x6a, 0x8c, 0x7a, 0x4e, 0xe0, 0x8c, 0xeb, 0xf3, 0x1a, 0xa3, 0x64, 0x04, 0xad,
	0x52, 0x62, 0xb1, 0x64, 0xd4, 0xab, 0x99, 0x62, 0x53, 0xc3, 0x19, 0x25, 0x03, 0x68, 0x28, 0xa6,
	0x72, 0xf4, 0xea, 0x81, 0x33, 0x6e, 0xcf, 0x2d, 0x20, 0x04, 0xdc, 0x94, 0xd3, 0x83, 0xe7, 0x9a,
	0xa2, 0x39, 0x87, 0x9f, 0xa1, 0xfb, 0x0f, 0x95, 0x56, 0x9f, 0xe3, 0x43, 0x89, 0x52, 0x69, 0x51,
	0xed, 0x68, 0x79, 0x7a, 0xa9, 0xa9, 0xe1, 0x8c, 0x86, 0x13, 0xe8, 0xff, 0xc1, 0x1c, 0x15, 0x7e,
	0x88, 0x3d, 0x85, 0xde, 0x7f, 0x26, 0x8d, 0xb2, 0x3c, 0x92, 0x09, 0xb8, 0x62, 0x95, 0x61, 0xc5,
	0x34, 0x67, 0x6d, 0x35, 0x67, 0x5b, 0xa6, 0xaa, 0x04, 0x16, 0x84, 0x3f, 0xa0, 0x7f, 0xd6, 0x2d,
	0x05, 0xdf, 0x49, 0x24, 0x01, 0x34, 0xb4, 0xb8, 0xf4, 0x9c, 0xa0, 0x3e, 0xee, 0x24, 0x10, 0x69,
	0x14, 0x19, 0x37, 0xf6, 0x22, 0x79, 0x72, 0xa0, 0xa3, 0xf1, 0x02, 0x8b, 0x3d, 0x5b, 0x23, 0xf9,
	0x0a, 0xad, 0x2a, 0x1d, 0x19, 0x58, 0xf6, 0xeb, 0xb0, 0xfe, 0x99, 0x06, 0x99, 0x42, 0xfb, 0xf4,
	0x26, 0x19, 0xda, 0x8b, 0xb7, 0x11, 0xfc, 0xd1, 0x45, 0xbd, 0x32, 0xf7, 0x0d, 0xe0, 0x5a, 0xd0,
	0x95, 0x9d, 0x0e, 0x39, 0xd3, 0xf5, 0x87, 0x91, 0xdd, 0x6d, 0x74, 0xdc, 0x6d, 0xf4, 0x57, 0xef,
	0x96, 0xfc, 0x04, 0x78, 0x99, 0x27, 0xa9, 0x84, 0x2f, 0x26, 0xfc, 0x5e, 0xfb, 0xef, 0xc9, 0xed,
	0x97, 0x8c, 0xa9, 0x4d, 0x99, 0x46, 0x6b, 0xbe, 0x8d, 0x17, 0x1b, 0x2c, 0x1e, 0x39, 0xfd, 0x95,
	0xd2, 0x32, 0xcf, 0x57, 0xf7, 0x7c, 0xc7, 0xf7, 0xe6, 0x9f, 0xb1, 0x35, 0x26, 0xb1, 0x48, 0xd3,
	0xa6, 0xe9, 0xfe, 0xfe, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x2b, 0x01, 0x0c, 0x80, 0x02, 0x00,
	0x00,
}
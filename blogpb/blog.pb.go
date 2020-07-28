// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/blogpb/blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRequest) Reset()         { *m = CreateBlogRequest{} }
func (m *CreateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRequest) ProtoMessage()    {}
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{1}
}

func (m *CreateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRequest.Unmarshal(m, b)
}
func (m *CreateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRequest.Marshal(b, m, deterministic)
}
func (m *CreateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRequest.Merge(m, src)
}
func (m *CreateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRequest.Size(m)
}
func (m *CreateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRequest proto.InternalMessageInfo

func (m *CreateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogResponse) Reset()         { *m = CreateBlogResponse{} }
func (m *CreateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBlogResponse) ProtoMessage()    {}
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{2}
}

func (m *CreateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogResponse.Unmarshal(m, b)
}
func (m *CreateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogResponse.Marshal(b, m, deterministic)
}
func (m *CreateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogResponse.Merge(m, src)
}
func (m *CreateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBlogResponse.Size(m)
}
func (m *CreateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogResponse proto.InternalMessageInfo

func (m *CreateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRequest) Reset()         { *m = ReadBlogRequest{} }
func (m *ReadBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRequest) ProtoMessage()    {}
func (*ReadBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{3}
}

func (m *ReadBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRequest.Unmarshal(m, b)
}
func (m *ReadBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRequest.Marshal(b, m, deterministic)
}
func (m *ReadBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRequest.Merge(m, src)
}
func (m *ReadBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRequest.Size(m)
}
func (m *ReadBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRequest proto.InternalMessageInfo

func (m *ReadBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type ReadBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogResponse) Reset()         { *m = ReadBlogResponse{} }
func (m *ReadBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ReadBlogResponse) ProtoMessage()    {}
func (*ReadBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{4}
}

func (m *ReadBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogResponse.Unmarshal(m, b)
}
func (m *ReadBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogResponse.Marshal(b, m, deterministic)
}
func (m *ReadBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogResponse.Merge(m, src)
}
func (m *ReadBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ReadBlogResponse.Size(m)
}
func (m *ReadBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogResponse proto.InternalMessageInfo

func (m *ReadBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogRequest) Reset()         { *m = UpdateBlogRequest{} }
func (m *UpdateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogRequest) ProtoMessage()    {}
func (*UpdateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{5}
}

func (m *UpdateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogRequest.Unmarshal(m, b)
}
func (m *UpdateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogRequest.Merge(m, src)
}
func (m *UpdateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogRequest.Size(m)
}
func (m *UpdateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogRequest proto.InternalMessageInfo

func (m *UpdateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogResponse) Reset()         { *m = UpdateBlogResponse{} }
func (m *UpdateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogResponse) ProtoMessage()    {}
func (*UpdateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4b0406114889fe6, []int{6}
}

func (m *UpdateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogResponse.Unmarshal(m, b)
}
func (m *UpdateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogResponse.Marshal(b, m, deterministic)
}
func (m *UpdateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogResponse.Merge(m, src)
}
func (m *UpdateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogResponse.Size(m)
}
func (m *UpdateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogResponse proto.InternalMessageInfo

func (m *UpdateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogRequest)(nil), "blog.CreateBlogRequest")
	proto.RegisterType((*CreateBlogResponse)(nil), "blog.CreateBlogResponse")
	proto.RegisterType((*ReadBlogRequest)(nil), "blog.ReadBlogRequest")
	proto.RegisterType((*ReadBlogResponse)(nil), "blog.ReadBlogResponse")
	proto.RegisterType((*UpdateBlogRequest)(nil), "blog.UpdateBlogRequest")
	proto.RegisterType((*UpdateBlogResponse)(nil), "blog.UpdateBlogResponse")
}

func init() { proto.RegisterFile("blog/blogpb/blog.proto", fileDescriptor_a4b0406114889fe6) }

var fileDescriptor_a4b0406114889fe6 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x6d, 0x11, 0x29, 0x9d, 0x26, 0x6a, 0x27, 0xda, 0x6e, 0x6a, 0x62, 0x0c, 0x27, 0xe3, 0xa1,
	0x26, 0xe0, 0xcd, 0x53, 0xeb, 0xa9, 0x57, 0x8c, 0x17, 0x2f, 0x0d, 0xb0, 0x93, 0xba, 0x09, 0x61,
	0x11, 0xb6, 0x7e, 0xa7, 0x9f, 0x64, 0x76, 0x37, 0x14, 0x52, 0x0e, 0xd6, 0x0b, 0x30, 0x33, 0xef,
	0xf1, 0xe6, 0xbd, 0x0c, 0xcc, 0xd2, 0x5c, 0xee, 0x9e, 0xf4, 0xa3, 0x4c, 0xcd, 0x6b, 0x59, 0x56,
	0x52, 0x49, 0x74, 0xf5, 0x77, 0x90, 0x81, 0xbb, 0xce, 0xe5, 0x0e, 0x2f, 0xc0, 0x11, 0x9c, 0x0d,
	0xef, 0x87, 0x0f, 0xe3, 0xd8, 0x11, 0x1c, 0x6f, 0x61, 0x9c, 0xec, 0xd5, 0xa7, 0xac, 0xb6, 0x82,
	0x33, 0xc7, 0xb4, 0x7d, 0xdb, 0xd8, 0x70, 0xbc, 0x86, 0x73, 0x25, 0x54, 0x4e, 0xec, 0xcc, 0x0c,
	0x6c, 0x81, 0x0c, 0x46, 0x99, 0x2c, 0x14, 0x15, 0x8a, 0xb9, 0xa6, 0xdf, 0x94, 0x41, 0x04, 0xd3,
	0xd7, 0x8a, 0x12, 0x45, 0x5a, 0x2a, 0xa6, 0xaf, 0x3d, 0xd5, 0x0a, 0xef, 0xc0, 0x6c, 0x60, 0x34,
	0x27, 0x21, 0x2c, 0xcd, 0x6a, 0x06, 0x60, 0x37, 0x7b, 0x06, 0xec, 0x92, 0xea, 0x52, 0x16, 0x35,
	0xfd, 0xc9, 0x7a, 0x84, 0xcb, 0x98, 0x12, 0xde, 0x15, 0x9a, 0xc3, 0x48, 0x8f, 0xb6, 0x07, 0x7f,
	0x9e, 0x2e, 0x37, 0x3c, 0x08, 0xe1, 0xaa, 0xc5, 0x9e, 0xf8, 0xff, 0x08, 0xa6, 0xef, 0x25, 0xff,
	0xbf, 0x95, 0x2e, 0xe9, 0x34, 0xa9, 0xf0, 0x67, 0x08, 0x13, 0x5d, 0xbe, 0x51, 0xf5, 0x2d, 0x32,
	0xc2, 0x15, 0x40, 0x1b, 0x08, 0xce, 0x2d, 0xbe, 0x97, 0xeb, 0x82, 0xf5, 0x07, 0x56, 0x30, 0x18,
	0xe0, 0x0b, 0xf8, 0x8d, 0x63, 0xbc, 0xb1, 0xb8, 0xa3, 0xb4, 0x16, 0xb3, 0xe3, 0xf6, 0x81, 0xbc,
	0x02, 0x68, 0x5d, 0x34, 0xfa, 0xbd, 0x30, 0x1a, 0xfd, 0xbe, 0xe1, 0x60, 0xb0, 0xf6, 0x3f, 0x3c,
	0x7b, 0x88, 0xa9, 0x67, 0x8e, 0x30, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xa1, 0xdc, 0xad,
	0x9e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
	ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogResponse, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error) {
	out := new(ReadBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogResponse, error) {
	out := new(UpdateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
	ReadBlog(context.Context, *ReadBlogRequest) (*ReadBlogResponse, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogResponse, error)
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateBlog(ctx context.Context, req *CreateBlogRequest) (*CreateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ReadBlog(ctx context.Context, req *ReadBlogRequest) (*ReadBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (*UnimplementedBlogServiceServer) UpdateBlog(ctx context.Context, req *UpdateBlogRequest) (*UpdateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/blogpb/blog.proto",
}
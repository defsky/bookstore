// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package com_afkplayer_service_user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CreatedTime          uint64   `protobuf:"varint,5,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          uint64   `protobuf:"varint,6,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreatedTime() uint64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *User) GetUpdatedTime() uint64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type Token struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	IsValid              bool     `protobuf:"varint,2,opt,name=isValid,proto3" json:"isValid,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Token) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *Token) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Token                *Token   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Request) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	Token                *Token   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Errors               []*Error `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "com.afkplayer.service.user.user")
	proto.RegisterType((*Token)(nil), "com.afkplayer.service.user.Token")
	proto.RegisterType((*Error)(nil), "com.afkplayer.service.user.Error")
	proto.RegisterType((*Request)(nil), "com.afkplayer.service.user.Request")
	proto.RegisterType((*Response)(nil), "com.afkplayer.service.user.Response")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x26, 0x3f, 0x4e, 0xd3, 0x59, 0x01, 0x92, 0x85, 0x90, 0xb5, 0xa7, 0x10, 0x38, 0xf4, 0x14,
	0xa4, 0x2d, 0xa2, 0xe2, 0x58, 0x21, 0xd4, 0x23, 0xc8, 0x6a, 0x39, 0x70, 0x33, 0xf1, 0x20, 0xa2,
	0x26, 0xeb, 0x60, 0x3b, 0x05, 0x5e, 0x80, 0xe7, 0xe0, 0x11, 0x78, 0x44, 0xe4, 0x71, 0x77, 0xe9,
	0x85, 0xdd, 0x54, 0xda, 0x8b, 0xd7, 0x9f, 0xf7, 0x9b, 0x99, 0x6f, 0xbe, 0xb1, 0x03, 0x8f, 0x27,
	0x87, 0xf6, 0x65, 0x58, 0x9a, 0xd1, 0x1a, 0x6f, 0xf8, 0xb2, 0x35, 0x43, 0xa3, 0xbe, 0x5c, 0x8f,
	0xbd, 0xfa, 0x89, 0xb6, 0x71, 0x68, 0x6f, 0xba, 0x16, 0x9b, 0xc0, 0xa8, 0x7f, 0x27, 0x90, 0x87,
	0x0d, 0x7f, 0x04, 0x69, 0xa7, 0x45, 0x52, 0x25, 0x27, 0xb9, 0x4c, 0x3b, 0xcd, 0x39, 0xe4, 0x6b,
	0x35, 0xa0, 0x48, 0xab, 0xe4, 0xe4, 0x58, 0xd2, 0x9e, 0x2f, 0xa1, 0x1c, 0x95, 0x73, 0xdf, 0x8d,
	0xd5, 0x22, 0xa3, 0xf3, 0x2d, 0xe6, 0x4f, 0x80, 0xe1, 0xa0, 0xba, 0x5e, 0xe4, 0xf4, 0x47, 0x04,
	0xbc, 0x82, 0x45, 0x6b, 0x51, 0x79, 0xd4, 0x97, 0xdd, 0x80, 0x82, 0x51, 0xfa, 0xbb, 0x47, 0x81,
	0x31, 0x8d, 0x7a, 0xcb, 0x28, 0x22, 0xe3, 0xce, 0x51, 0x3d, 0x02, 0xbb, 0x34, 0xd7, 0xb8, 0x0e,
	0x25, 0x6e, 0x54, 0x3f, 0x21, 0xa9, 0x3c, 0x96, 0x11, 0x70, 0x01, 0x47, 0x9d, 0xfb, 0xa8, 0xfa,
	0x4e, 0x93, 0xd6, 0x52, 0x6e, 0x20, 0x3f, 0x03, 0x86, 0xd6, 0x1a, 0x4b, 0x5a, 0x17, 0xab, 0x67,
	0xcd, 0xff, 0x7d, 0x68, 0xde, 0x05, 0xa2, 0x8c, 0xfc, 0xfa, 0x14, 0x18, 0xe1, 0x60, 0x42, 0x6b,
	0x74, 0x2c, 0xc8, 0x24, 0xed, 0xf9, 0x53, 0x28, 0x34, 0xfa, 0xd0, 0x69, 0xb4, 0xe6, 0x16, 0xd5,
	0x3f, 0xe0, 0x48, 0xe2, 0xb7, 0x09, 0x9d, 0xe7, 0xaf, 0xa2, 0xa7, 0x14, 0xb6, 0x58, 0x55, 0xbb,
	0xea, 0x86, 0x45, 0xc6, 0x09, 0x9c, 0x01, 0xf3, 0xa1, 0x4f, 0xca, 0xbb, 0x47, 0x2e, 0x19, 0x22,
	0x23, 0xbf, 0xfe, 0x95, 0x42, 0x29, 0xd1, 0x8d, 0x66, 0xed, 0xc8, 0x0e, 0x37, 0xb5, 0x2d, 0x3a,
	0x47, 0xe5, 0x4b, 0xb9, 0x81, 0x5b, 0x55, 0xe9, 0xbd, 0x54, 0xbd, 0x06, 0x16, 0x7e, 0x9d, 0xc8,
	0xab, 0x6c, 0x56, 0x58, 0xa4, 0xff, 0xeb, 0x26, 0xbb, 0x5f, 0x37, 0xfc, 0x0d, 0x14, 0x34, 0x05,
	0x27, 0x18, 0x55, 0x9c, 0x31, 0xb6, 0xdb, 0x80, 0xd5, 0x9f, 0x0c, 0xf2, 0xab, 0x20, 0x5a, 0x42,
	0xf1, 0x96, 0xee, 0x18, 0xdf, 0xab, 0x77, 0xf9, 0x62, 0x17, 0x63, 0x63, 0x6b, 0xfd, 0x80, 0xbf,
	0x87, 0xec, 0x02, 0xfd, 0x01, 0x13, 0x5e, 0x41, 0x71, 0x81, 0xfe, 0xbc, 0xef, 0xf9, 0xf3, 0xdd,
	0x11, 0x74, 0xa9, 0x66, 0xa7, 0xfd, 0x00, 0xf9, 0xf9, 0xe4, 0xbf, 0x1e, 0x50, 0xe8, 0x27, 0x78,
	0x48, 0x0f, 0x4a, 0x79, 0x8c, 0x0f, 0x71, 0xff, 0x30, 0xe7, 0xe6, 0xfe, 0x5c, 0xd0, 0x27, 0xea,
	0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x64, 0xf2, 0xeb, 0xb5, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	ValidateToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Response, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.afkplayer.service.user.User/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Get(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.afkplayer.service.user.User/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.afkplayer.service.user.User/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Auth(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.afkplayer.service.user.User/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ValidateToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/com.afkplayer.service.user.User/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Create(context.Context, *User) (*Response, error)
	Get(context.Context, *User) (*Response, error)
	GetAll(context.Context, *Request) (*Response, error)
	Auth(context.Context, *User) (*Response, error)
	ValidateToken(context.Context, *Token) (*Response, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.afkplayer.service.user.User/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.afkplayer.service.user.User/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Get(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.afkplayer.service.user.User/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAll(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.afkplayer.service.user.User/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Auth(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.afkplayer.service.user.User/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ValidateToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.afkplayer.service.user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _User_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _User_GetAll_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _User_Auth_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _User_ValidateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

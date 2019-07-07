// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email.proto

package emailpb

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

type ActiveInvalideCodeError struct {
	// Message of error
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// ID of missing user
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActiveInvalideCodeError) Reset()         { *m = ActiveInvalideCodeError{} }
func (m *ActiveInvalideCodeError) String() string { return proto.CompactTextString(m) }
func (*ActiveInvalideCodeError) ProtoMessage()    {}
func (*ActiveInvalideCodeError) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{0}
}

func (m *ActiveInvalideCodeError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveInvalideCodeError.Unmarshal(m, b)
}
func (m *ActiveInvalideCodeError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveInvalideCodeError.Marshal(b, m, deterministic)
}
func (m *ActiveInvalideCodeError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveInvalideCodeError.Merge(m, src)
}
func (m *ActiveInvalideCodeError) XXX_Size() int {
	return xxx_messageInfo_ActiveInvalideCodeError.Size(m)
}
func (m *ActiveInvalideCodeError) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveInvalideCodeError.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveInvalideCodeError proto.InternalMessageInfo

func (m *ActiveInvalideCodeError) GetMessage_() string {
	if m != nil {
		return m.Message_
	}
	return ""
}

func (m *ActiveInvalideCodeError) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ActiveRequest struct {
	// operand
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActiveRequest) Reset()         { *m = ActiveRequest{} }
func (m *ActiveRequest) String() string { return proto.CompactTextString(m) }
func (*ActiveRequest) ProtoMessage()    {}
func (*ActiveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{1}
}

func (m *ActiveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveRequest.Unmarshal(m, b)
}
func (m *ActiveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveRequest.Marshal(b, m, deterministic)
}
func (m *ActiveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveRequest.Merge(m, src)
}
func (m *ActiveRequest) XXX_Size() int {
	return xxx_messageInfo_ActiveRequest.Size(m)
}
func (m *ActiveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveRequest proto.InternalMessageInfo

func (m *ActiveRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ActiveResponse struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActiveResponse) Reset()         { *m = ActiveResponse{} }
func (m *ActiveResponse) String() string { return proto.CompactTextString(m) }
func (*ActiveResponse) ProtoMessage()    {}
func (*ActiveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{2}
}

func (m *ActiveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActiveResponse.Unmarshal(m, b)
}
func (m *ActiveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActiveResponse.Marshal(b, m, deterministic)
}
func (m *ActiveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActiveResponse.Merge(m, src)
}
func (m *ActiveResponse) XXX_Size() int {
	return xxx_messageInfo_ActiveResponse.Size(m)
}
func (m *ActiveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActiveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActiveResponse proto.InternalMessageInfo

func (m *ActiveResponse) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

type SendEmailInvalideEmailError struct {
	// Message of error
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// ID of missing user
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEmailInvalideEmailError) Reset()         { *m = SendEmailInvalideEmailError{} }
func (m *SendEmailInvalideEmailError) String() string { return proto.CompactTextString(m) }
func (*SendEmailInvalideEmailError) ProtoMessage()    {}
func (*SendEmailInvalideEmailError) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{3}
}

func (m *SendEmailInvalideEmailError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailInvalideEmailError.Unmarshal(m, b)
}
func (m *SendEmailInvalideEmailError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailInvalideEmailError.Marshal(b, m, deterministic)
}
func (m *SendEmailInvalideEmailError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailInvalideEmailError.Merge(m, src)
}
func (m *SendEmailInvalideEmailError) XXX_Size() int {
	return xxx_messageInfo_SendEmailInvalideEmailError.Size(m)
}
func (m *SendEmailInvalideEmailError) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailInvalideEmailError.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailInvalideEmailError proto.InternalMessageInfo

func (m *SendEmailInvalideEmailError) GetMessage_() string {
	if m != nil {
		return m.Message_
	}
	return ""
}

func (m *SendEmailInvalideEmailError) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SendEmailEmailNotFoundError struct {
	// Message of error
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// ID of missing user
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEmailEmailNotFoundError) Reset()         { *m = SendEmailEmailNotFoundError{} }
func (m *SendEmailEmailNotFoundError) String() string { return proto.CompactTextString(m) }
func (*SendEmailEmailNotFoundError) ProtoMessage()    {}
func (*SendEmailEmailNotFoundError) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{4}
}

func (m *SendEmailEmailNotFoundError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailEmailNotFoundError.Unmarshal(m, b)
}
func (m *SendEmailEmailNotFoundError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailEmailNotFoundError.Marshal(b, m, deterministic)
}
func (m *SendEmailEmailNotFoundError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailEmailNotFoundError.Merge(m, src)
}
func (m *SendEmailEmailNotFoundError) XXX_Size() int {
	return xxx_messageInfo_SendEmailEmailNotFoundError.Size(m)
}
func (m *SendEmailEmailNotFoundError) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailEmailNotFoundError.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailEmailNotFoundError proto.InternalMessageInfo

func (m *SendEmailEmailNotFoundError) GetMessage_() string {
	if m != nil {
		return m.Message_
	}
	return ""
}

func (m *SendEmailEmailNotFoundError) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SendEmailRequest struct {
	// email of userProfile
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEmailRequest) Reset()         { *m = SendEmailRequest{} }
func (m *SendEmailRequest) String() string { return proto.CompactTextString(m) }
func (*SendEmailRequest) ProtoMessage()    {}
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{5}
}

func (m *SendEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailRequest.Unmarshal(m, b)
}
func (m *SendEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailRequest.Marshal(b, m, deterministic)
}
func (m *SendEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailRequest.Merge(m, src)
}
func (m *SendEmailRequest) XXX_Size() int {
	return xxx_messageInfo_SendEmailRequest.Size(m)
}
func (m *SendEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailRequest proto.InternalMessageInfo

func (m *SendEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type SendEmailResponse struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendEmailResponse) Reset()         { *m = SendEmailResponse{} }
func (m *SendEmailResponse) String() string { return proto.CompactTextString(m) }
func (*SendEmailResponse) ProtoMessage()    {}
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{6}
}

func (m *SendEmailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendEmailResponse.Unmarshal(m, b)
}
func (m *SendEmailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendEmailResponse.Marshal(b, m, deterministic)
}
func (m *SendEmailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendEmailResponse.Merge(m, src)
}
func (m *SendEmailResponse) XXX_Size() int {
	return xxx_messageInfo_SendEmailResponse.Size(m)
}
func (m *SendEmailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendEmailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendEmailResponse proto.InternalMessageInfo

func (m *SendEmailResponse) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func init() {
	proto.RegisterType((*ActiveInvalideCodeError)(nil), "email.ActiveInvalideCodeError")
	proto.RegisterType((*ActiveRequest)(nil), "email.ActiveRequest")
	proto.RegisterType((*ActiveResponse)(nil), "email.ActiveResponse")
	proto.RegisterType((*SendEmailInvalideEmailError)(nil), "email.SendEmailInvalideEmailError")
	proto.RegisterType((*SendEmailEmailNotFoundError)(nil), "email.SendEmailEmailNotFoundError")
	proto.RegisterType((*SendEmailRequest)(nil), "email.SendEmailRequest")
	proto.RegisterType((*SendEmailResponse)(nil), "email.SendEmailResponse")
}

func init() { proto.RegisterFile("email.proto", fileDescriptor_6175298cb4ed6faa) }

var fileDescriptor_6175298cb4ed6faa = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0x4d, 0xcc,
	0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x5c, 0xb8, 0xc4, 0x1d,
	0x93, 0x4b, 0x32, 0xcb, 0x52, 0x3d, 0xf3, 0xca, 0x12, 0x73, 0x32, 0x53, 0x52, 0x9d, 0xf3, 0x53,
	0x52, 0x5d, 0x8b, 0x8a, 0xf2, 0x8b, 0x84, 0x24, 0xb9, 0x38, 0x72, 0x53, 0x8b, 0x8b, 0x13, 0xd3,
	0x53, 0xe3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xd8, 0xa1, 0x7c, 0x21, 0x3e, 0x2e, 0xa6,
	0xcc, 0x14, 0x09, 0x26, 0xb0, 0x20, 0x53, 0x66, 0x8a, 0x92, 0x32, 0x17, 0x2f, 0xc4, 0x94, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54, 0xa8, 0x3e,
	0x30, 0x5b, 0x49, 0x8d, 0x8b, 0x0f, 0xa6, 0xa8, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x84,
	0x8b, 0x35, 0x2d, 0x33, 0x35, 0x27, 0x05, 0xaa, 0x0c, 0xc2, 0x51, 0xf2, 0xe0, 0x92, 0x0e, 0x4e,
	0xcd, 0x4b, 0x71, 0x05, 0xb9, 0x0f, 0xe6, 0x2a, 0x30, 0x87, 0x64, 0x67, 0x21, 0x9b, 0x04, 0x26,
	0xfc, 0xf2, 0x4b, 0xdc, 0xf2, 0x4b, 0xf3, 0x52, 0x48, 0x36, 0x49, 0x83, 0x4b, 0x00, 0x6e, 0x12,
	0xcc, 0x8f, 0x22, 0x5c, 0x90, 0x30, 0x84, 0xb9, 0x1e, 0x12, 0xa0, 0x9a, 0x5c, 0x82, 0x48, 0x2a,
	0xf1, 0x79, 0xd4, 0xa8, 0x8e, 0x8b, 0x15, 0xac, 0x4c, 0xc8, 0x94, 0x8b, 0x0d, 0x12, 0x32, 0x42,
	0x22, 0x7a, 0x90, 0x38, 0x42, 0x09, 0x4d, 0x29, 0x51, 0x34, 0x51, 0xa8, 0xa9, 0x76, 0x5c, 0x9c,
	0x70, 0xab, 0x84, 0xc4, 0xa1, 0x6a, 0xd0, 0x9d, 0x29, 0x25, 0x81, 0x29, 0x01, 0xd1, 0xef, 0xc4,
	0x19, 0xc5, 0x0e, 0x96, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0x27, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0xfd, 0x48, 0xca, 0x23, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailClient interface {
	// Active user by email code
	Active(ctx context.Context, in *ActiveRequest, opts ...grpc.CallOption) (*ActiveResponse, error)
	// Send email to active user
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
}

type emailClient struct {
	cc *grpc.ClientConn
}

func NewEmailClient(cc *grpc.ClientConn) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) Active(ctx context.Context, in *ActiveRequest, opts ...grpc.CallOption) (*ActiveResponse, error) {
	out := new(ActiveResponse)
	err := c.cc.Invoke(ctx, "/email.Email/Active", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	out := new(SendEmailResponse)
	err := c.cc.Invoke(ctx, "/email.Email/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
type EmailServer interface {
	// Active user by email code
	Active(context.Context, *ActiveRequest) (*ActiveResponse, error)
	// Send email to active user
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
}

func RegisterEmailServer(s *grpc.Server, srv EmailServer) {
	s.RegisterService(&_Email_serviceDesc, srv)
}

func _Email_Active_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).Active(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.Email/Active",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).Active(ctx, req.(*ActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Email_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email.Email/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Email_serviceDesc = grpc.ServiceDesc{
	ServiceName: "email.Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Active",
			Handler:    _Email_Active_Handler,
		},
		{
			MethodName: "SendEmail",
			Handler:    _Email_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}
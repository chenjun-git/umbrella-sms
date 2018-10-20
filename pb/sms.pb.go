// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sms.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	sms.proto

It has these top-level messages:
	SmsSendSingleReq
	SmsSendSingleResp
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SmsSendSingleReq struct {
	Uid        string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Mobile     string `protobuf:"bytes,2,opt,name=mobile" json:"mobile,omitempty"`
	VerifyCode string `protobuf:"bytes,3,opt,name=verify_code" json:"verify_code,omitempty"`
}

func (m *SmsSendSingleReq) Reset()                    { *m = SmsSendSingleReq{} }
func (m *SmsSendSingleReq) String() string            { return proto.CompactTextString(m) }
func (*SmsSendSingleReq) ProtoMessage()               {}
func (*SmsSendSingleReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SmsSendSingleReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SmsSendSingleReq) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SmsSendSingleReq) GetVerifyCode() string {
	if m != nil {
		return m.VerifyCode
	}
	return ""
}

type SmsSendSingleResp struct {
}

func (m *SmsSendSingleResp) Reset()                    { *m = SmsSendSingleResp{} }
func (m *SmsSendSingleResp) String() string            { return proto.CompactTextString(m) }
func (*SmsSendSingleResp) ProtoMessage()               {}
func (*SmsSendSingleResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*SmsSendSingleReq)(nil), "pb.SmsSendSingleReq")
	proto.RegisterType((*SmsSendSingleResp)(nil), "pb.SmsSendSingleResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sms service

type SmsClient interface {
	SendSignupVerifyCode(ctx context.Context, in *SmsSendSingleReq, opts ...grpc.CallOption) (*SmsSendSingleResp, error)
}

type smsClient struct {
	cc *grpc.ClientConn
}

func NewSmsClient(cc *grpc.ClientConn) SmsClient {
	return &smsClient{cc}
}

func (c *smsClient) SendSignupVerifyCode(ctx context.Context, in *SmsSendSingleReq, opts ...grpc.CallOption) (*SmsSendSingleResp, error) {
	out := new(SmsSendSingleResp)
	err := grpc.Invoke(ctx, "/pb.Sms/SendSignupVerifyCode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sms service

type SmsServer interface {
	SendSignupVerifyCode(context.Context, *SmsSendSingleReq) (*SmsSendSingleResp, error)
}

func RegisterSmsServer(s *grpc.Server, srv SmsServer) {
	s.RegisterService(&_Sms_serviceDesc, srv)
}

func _Sms_SendSignupVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsSendSingleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServer).SendSignupVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sms/SendSignupVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServer).SendSignupVerifyCode(ctx, req.(*SmsSendSingleReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sms_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Sms",
	HandlerType: (*SmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSignupVerifyCode",
			Handler:    _Sms_SendSignupVerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms.proto",
}

func init() { proto.RegisterFile("sms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xce, 0x2d, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe1, 0x12, 0x08, 0xce, 0x2d,
	0x0e, 0x4e, 0xcd, 0x4b, 0x09, 0xce, 0xcc, 0x4b, 0xcf, 0x49, 0x0d, 0x4a, 0x2d, 0x14, 0xe2, 0xe6,
	0x62, 0x2e, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x14, 0xe2, 0xe3, 0x62, 0xcb, 0xcd,
	0x4f, 0xca, 0xcc, 0x49, 0x95, 0x60, 0x02, 0xf3, 0x85, 0xb9, 0xb8, 0xcb, 0x52, 0x8b, 0x32, 0xd3,
	0x2a, 0xe3, 0x93, 0xf3, 0x53, 0x52, 0x25, 0x98, 0x41, 0x82, 0x4a, 0xc2, 0x5c, 0x82, 0x68, 0xa6,
	0x14, 0x17, 0x18, 0xf9, 0x70, 0x31, 0x07, 0xe7, 0x16, 0x0b, 0xb9, 0x72, 0x89, 0x40, 0x24, 0xd2,
	0xf3, 0x4a, 0x0b, 0xc2, 0xc0, 0x5a, 0x9d, 0xf3, 0x53, 0x52, 0x85, 0x44, 0xf4, 0x0a, 0x92, 0xf4,
	0xd0, 0xed, 0x96, 0x12, 0xc5, 0x22, 0x5a, 0x5c, 0xa0, 0xc4, 0x90, 0xc4, 0x06, 0x76, 0xb3, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x2a, 0x31, 0x83, 0xc0, 0x00, 0x00, 0x00,
}
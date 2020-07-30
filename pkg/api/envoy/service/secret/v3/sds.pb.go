// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/secret/v3/sds.proto

package envoy_service_secret_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/service/discovery/v3"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SdsDummy) Reset()         { *m = SdsDummy{} }
func (m *SdsDummy) String() string { return proto.CompactTextString(m) }
func (*SdsDummy) ProtoMessage()    {}
func (*SdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_879f5c9ef2affb70, []int{0}
}

func (m *SdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SdsDummy.Unmarshal(m, b)
}
func (m *SdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SdsDummy.Marshal(b, m, deterministic)
}
func (m *SdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SdsDummy.Merge(m, src)
}
func (m *SdsDummy) XXX_Size() int {
	return xxx_messageInfo_SdsDummy.Size(m)
}
func (m *SdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SdsDummy)(nil), "envoy.service.secret.v3.SdsDummy")
}

func init() { proto.RegisterFile("envoy/service/secret/v3/sds.proto", fileDescriptor_879f5c9ef2affb70) }

var fileDescriptor_879f5c9ef2affb70 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x41, 0x4b, 0x1b, 0x41,
	0x1c, 0xc5, 0xbb, 0x09, 0xa4, 0x65, 0x48, 0xa1, 0x2c, 0xb4, 0x81, 0xa5, 0x2d, 0x4d, 0x4a, 0x4a,
	0x09, 0xed, 0x6c, 0xba, 0x0b, 0x45, 0x72, 0xf0, 0x10, 0x82, 0xe7, 0xe0, 0x7e, 0x00, 0x19, 0x77,
	0xff, 0xc4, 0xc5, 0x64, 0x66, 0x9d, 0xff, 0xec, 0x92, 0xc5, 0x8b, 0x78, 0x0a, 0x5e, 0x03, 0x1e,
	0x3c, 0xfb, 0x29, 0xbc, 0x0b, 0x5e, 0xc5, 0x93, 0x77, 0xbf, 0x81, 0x5f, 0x40, 0x76, 0x27, 0x1b,
	0x5d, 0x25, 0xa2, 0x17, 0x8f, 0xc3, 0xfb, 0xbd, 0xf7, 0xfe, 0x03, 0x8f, 0x34, 0x81, 0x27, 0x22,
	0xb5, 0x11, 0x64, 0x12, 0xfa, 0x60, 0x23, 0xf8, 0x12, 0x94, 0x9d, 0xb8, 0x36, 0x06, 0x48, 0x23,
	0x29, 0x94, 0x30, 0x1b, 0x39, 0x42, 0x17, 0x08, 0xd5, 0x08, 0x4d, 0x5c, 0xab, 0x53, 0xf6, 0x06,
	0x21, 0xfa, 0x22, 0x01, 0x99, 0x66, 0xf6, 0xe5, 0x43, 0x87, 0x58, 0x5f, 0x47, 0x42, 0x8c, 0xc6,
	0x60, 0xb3, 0x28, 0xb4, 0x19, 0xe7, 0x42, 0x31, 0x15, 0x0a, 0xbe, 0xa8, 0xb0, 0x7e, 0xe8, 0xa4,
	0x07, 0x82, 0x2d, 0x01, 0x45, 0x2c, 0x7d, 0x58, 0x10, 0xdf, 0xe2, 0x20, 0x62, 0x25, 0x00, 0x15,
	0x53, 0x71, 0x11, 0xd0, 0x7c, 0x22, 0x27, 0x20, 0x31, 0x14, 0x3c, 0xe4, 0x23, 0x8d, 0xb4, 0xfe,
	0x93, 0x0f, 0x5e, 0x80, 0x83, 0x78, 0x32, 0x49, 0x7b, 0x9d, 0x93, 0xf3, 0xd9, 0xf7, 0x36, 0xf9,
	0x59, 0xfe, 0xd9, 0xfd, 0xcd, 0x89, 0x43, 0x0b, 0xd6, 0xb9, 0xae, 0x92, 0x2f, 0x5e, 0xfe, 0xe7,
	0x41, 0xa1, 0x7b, 0xda, 0x60, 0xee, 0x93, 0xfa, 0x00, 0xc6, 0x8a, 0x69, 0x19, 0xcd, 0x7f, 0x74,
	0x65, 0xa0, 0x4b, 0x73, 0x72, 0x19, 0xb1, 0x09, 0x7b, 0x31, 0xa0, 0xb2, 0x9c, 0xd7, 0x58, 0x30,
	0x12, 0x1c, 0xa1, 0xf5, 0xee, 0xb7, 0xd1, 0x35, 0x4c, 0x49, 0x3e, 0x7a, 0x4a, 0x02, 0x9b, 0x14,
	0xed, 0x7f, 0x9e, 0x8d, 0x7a, 0x5c, 0xfc, 0xf7, 0x85, 0x74, 0xa9, 0x73, 0x6e, 0x90, 0xfa, 0x06,
	0x28, 0x7f, 0xe7, 0x4d, 0x3a, 0x7f, 0x1d, 0x5e, 0xdd, 0xcc, 0x2b, 0x8d, 0xd6, 0xe7, 0xd2, 0xa0,
	0x7a, 0x7a, 0x80, 0x98, 0x8b, 0xd5, 0x9e, 0xd1, 0xb1, 0xd6, 0x8e, 0x4e, 0x8f, 0x6f, 0xdf, 0x3b,
	0xa4, 0xab, 0xd3, 0x61, 0xaa, 0x80, 0x63, 0xbe, 0x2e, 0x25, 0x19, 0xc7, 0x48, 0x48, 0xb5, 0x85,
	0xc2, 0xdf, 0x05, 0x85, 0x54, 0x8d, 0x31, 0x2b, 0xd3, 0xf7, 0xf7, 0xd7, 0xcf, 0x0e, 0x2e, 0x2e,
	0x6b, 0x95, 0x4f, 0x15, 0xd2, 0x0e, 0x85, 0x3e, 0x2e, 0x92, 0x62, 0x9a, 0xd2, 0x15, 0x93, 0xef,
	0x67, 0x13, 0x1a, 0x66, 0x73, 0x1a, 0x1a, 0x33, 0xc3, 0xd8, 0xae, 0xe5, 0xd3, 0x72, 0xef, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x2d, 0xd2, 0xc4, 0x2f, 0x46, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretDiscoveryServiceClient is the client API for SecretDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretDiscoveryServiceClient interface {
	DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error)
	StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error)
	FetchSecrets(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error)
}

type secretDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretDiscoveryServiceClient(cc *grpc.ClientConn) SecretDiscoveryServiceClient {
	return &secretDiscoveryServiceClient{cc}
}

func (c *secretDiscoveryServiceClient) DeltaSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_DeltaSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[0], "/envoy.service.secret.v3.SecretDiscoveryService/DeltaSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceDeltaSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_DeltaSecretsClient interface {
	Send(*v3.DeltaDiscoveryRequest) error
	Recv() (*v3.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceDeltaSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Send(m *v3.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsClient) Recv() (*v3.DeltaDiscoveryResponse, error) {
	m := new(v3.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[1], "/envoy.service.secret.v3.SecretDiscoveryService/StreamSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceStreamSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_StreamSecretsClient interface {
	Send(*v3.DiscoveryRequest) error
	Recv() (*v3.DiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceStreamSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceStreamSecretsClient) Send(m *v3.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsClient) Recv() (*v3.DiscoveryResponse, error) {
	m := new(v3.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, in *v3.DiscoveryRequest, opts ...grpc.CallOption) (*v3.DiscoveryResponse, error) {
	out := new(v3.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.secret.v3.SecretDiscoveryService/FetchSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretDiscoveryServiceServer is the server API for SecretDiscoveryService service.
type SecretDiscoveryServiceServer interface {
	DeltaSecrets(SecretDiscoveryService_DeltaSecretsServer) error
	StreamSecrets(SecretDiscoveryService_StreamSecretsServer) error
	FetchSecrets(context.Context, *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error)
}

// UnimplementedSecretDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSecretDiscoveryServiceServer struct {
}

func (*UnimplementedSecretDiscoveryServiceServer) DeltaSecrets(srv SecretDiscoveryService_DeltaSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaSecrets not implemented")
}
func (*UnimplementedSecretDiscoveryServiceServer) StreamSecrets(srv SecretDiscoveryService_StreamSecretsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSecrets not implemented")
}
func (*UnimplementedSecretDiscoveryServiceServer) FetchSecrets(ctx context.Context, req *v3.DiscoveryRequest) (*v3.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSecrets not implemented")
}

func RegisterSecretDiscoveryServiceServer(s *grpc.Server, srv SecretDiscoveryServiceServer) {
	s.RegisterService(&_SecretDiscoveryService_serviceDesc, srv)
}

func _SecretDiscoveryService_DeltaSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).DeltaSecrets(&secretDiscoveryServiceDeltaSecretsServer{stream})
}

type SecretDiscoveryService_DeltaSecretsServer interface {
	Send(*v3.DeltaDiscoveryResponse) error
	Recv() (*v3.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceDeltaSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Send(m *v3.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceDeltaSecretsServer) Recv() (*v3.DeltaDiscoveryRequest, error) {
	m := new(v3.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_StreamSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).StreamSecrets(&secretDiscoveryServiceStreamSecretsServer{stream})
}

type SecretDiscoveryService_StreamSecretsServer interface {
	Send(*v3.DiscoveryResponse) error
	Recv() (*v3.DiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceStreamSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceStreamSecretsServer) Send(m *v3.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsServer) Recv() (*v3.DiscoveryRequest, error) {
	m := new(v3.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_FetchSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.secret.v3.SecretDiscoveryService/FetchSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, req.(*v3.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.secret.v3.SecretDiscoveryService",
	HandlerType: (*SecretDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSecrets",
			Handler:    _SecretDiscoveryService_FetchSecrets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeltaSecrets",
			Handler:       _SecretDiscoveryService_DeltaSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamSecrets",
			Handler:       _SecretDiscoveryService_StreamSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/secret/v3/sds.proto",
}

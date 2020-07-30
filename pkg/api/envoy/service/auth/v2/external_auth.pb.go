// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v2/external_auth.proto

package envoy_service_auth_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	_type "github.com/datawire/ambassador/pkg/api/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

type CheckRequest struct {
	Attributes           *AttributeContext `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5257cfee93a30acb, []int{0}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetAttributes() *AttributeContext {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type DeniedHttpResponse struct {
	Status               *_type.HttpStatus         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Headers              []*core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	Body                 string                    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DeniedHttpResponse) Reset()         { *m = DeniedHttpResponse{} }
func (m *DeniedHttpResponse) String() string { return proto.CompactTextString(m) }
func (*DeniedHttpResponse) ProtoMessage()    {}
func (*DeniedHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5257cfee93a30acb, []int{1}
}

func (m *DeniedHttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeniedHttpResponse.Unmarshal(m, b)
}
func (m *DeniedHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeniedHttpResponse.Marshal(b, m, deterministic)
}
func (m *DeniedHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeniedHttpResponse.Merge(m, src)
}
func (m *DeniedHttpResponse) XXX_Size() int {
	return xxx_messageInfo_DeniedHttpResponse.Size(m)
}
func (m *DeniedHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeniedHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeniedHttpResponse proto.InternalMessageInfo

func (m *DeniedHttpResponse) GetStatus() *_type.HttpStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DeniedHttpResponse) GetHeaders() []*core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DeniedHttpResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type OkHttpResponse struct {
	Headers              []*core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *OkHttpResponse) Reset()         { *m = OkHttpResponse{} }
func (m *OkHttpResponse) String() string { return proto.CompactTextString(m) }
func (*OkHttpResponse) ProtoMessage()    {}
func (*OkHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5257cfee93a30acb, []int{2}
}

func (m *OkHttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OkHttpResponse.Unmarshal(m, b)
}
func (m *OkHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OkHttpResponse.Marshal(b, m, deterministic)
}
func (m *OkHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkHttpResponse.Merge(m, src)
}
func (m *OkHttpResponse) XXX_Size() int {
	return xxx_messageInfo_OkHttpResponse.Size(m)
}
func (m *OkHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OkHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OkHttpResponse proto.InternalMessageInfo

func (m *OkHttpResponse) GetHeaders() []*core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

type CheckResponse struct {
	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are valid to be assigned to HttpResponse:
	//	*CheckResponse_DeniedResponse
	//	*CheckResponse_OkResponse
	HttpResponse         isCheckResponse_HttpResponse `protobuf_oneof:"http_response"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5257cfee93a30acb, []int{3}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type isCheckResponse_HttpResponse interface {
	isCheckResponse_HttpResponse()
}

type CheckResponse_DeniedResponse struct {
	DeniedResponse *DeniedHttpResponse `protobuf:"bytes,2,opt,name=denied_response,json=deniedResponse,proto3,oneof"`
}

type CheckResponse_OkResponse struct {
	OkResponse *OkHttpResponse `protobuf:"bytes,3,opt,name=ok_response,json=okResponse,proto3,oneof"`
}

func (*CheckResponse_DeniedResponse) isCheckResponse_HttpResponse() {}

func (*CheckResponse_OkResponse) isCheckResponse_HttpResponse() {}

func (m *CheckResponse) GetHttpResponse() isCheckResponse_HttpResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *CheckResponse) GetDeniedResponse() *DeniedHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_DeniedResponse); ok {
		return x.DeniedResponse
	}
	return nil
}

func (m *CheckResponse) GetOkResponse() *OkHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_OkResponse); ok {
		return x.OkResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CheckResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CheckResponse_DeniedResponse)(nil),
		(*CheckResponse_OkResponse)(nil),
	}
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "envoy.service.auth.v2.CheckRequest")
	proto.RegisterType((*DeniedHttpResponse)(nil), "envoy.service.auth.v2.DeniedHttpResponse")
	proto.RegisterType((*OkHttpResponse)(nil), "envoy.service.auth.v2.OkHttpResponse")
	proto.RegisterType((*CheckResponse)(nil), "envoy.service.auth.v2.CheckResponse")
}

func init() {
	proto.RegisterFile("envoy/service/auth/v2/external_auth.proto", fileDescriptor_5257cfee93a30acb)
}

var fileDescriptor_5257cfee93a30acb = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0x13, 0x3f,
	0x10, 0xc6, 0xe3, 0xe4, 0xdf, 0xfc, 0xc1, 0x21, 0x6d, 0xb1, 0x04, 0x8d, 0x22, 0x90, 0xa2, 0xb4,
	0x88, 0x14, 0x09, 0x5b, 0x0a, 0x17, 0x4e, 0x48, 0x4d, 0xa9, 0xc8, 0xad, 0xd1, 0x52, 0xc1, 0x31,
	0x72, 0x76, 0x47, 0xcd, 0x2a, 0xd1, 0xda, 0xd8, 0xb3, 0xab, 0x84, 0x13, 0x47, 0xc4, 0x63, 0xf0,
	0x18, 0x3c, 0x01, 0x57, 0xde, 0x83, 0x27, 0xe0, 0x84, 0xd6, 0x76, 0x42, 0x03, 0x09, 0x17, 0x6e,
	0xd6, 0xce, 0xf7, 0xfd, 0x3c, 0xdf, 0x78, 0x96, 0x9e, 0x42, 0x56, 0xa8, 0xa5, 0xb0, 0x60, 0x8a,
	0x34, 0x06, 0x21, 0x73, 0x9c, 0x8a, 0xa2, 0x2f, 0x60, 0x81, 0x60, 0x32, 0x39, 0x1f, 0x97, 0x1f,
	0xb8, 0x36, 0x0a, 0x15, 0xbb, 0xe7, 0xa4, 0x3c, 0x48, 0xb9, 0xab, 0x14, 0xfd, 0xf6, 0x03, 0x4f,
	0x90, 0x3a, 0x2d, 0x8d, 0xb1, 0x32, 0x20, 0x26, 0xd2, 0x82, 0x37, 0xb5, 0x9f, 0x6e, 0xe7, 0x4b,
	0x44, 0x93, 0x4e, 0x72, 0x84, 0x71, 0xac, 0x32, 0x84, 0x05, 0x06, 0x79, 0x80, 0xe1, 0x52, 0x83,
	0x98, 0x22, 0xea, 0xb1, 0x45, 0x89, 0xb9, 0x0d, 0xd5, 0xa3, 0x6b, 0xa5, 0xae, 0xe7, 0x20, 0x8c,
	0x8e, 0xc5, 0x46, 0xe1, 0x61, 0x9e, 0x68, 0x29, 0x64, 0x96, 0x29, 0x94, 0x98, 0xaa, 0xcc, 0x6e,
	0x96, 0x8f, 0x0a, 0x39, 0x4f, 0x13, 0x89, 0x20, 0x56, 0x07, 0x5f, 0xe8, 0xbe, 0xa5, 0x77, 0xce,
	0xa7, 0x10, 0xcf, 0x22, 0x78, 0x97, 0x83, 0x45, 0xf6, 0x8a, 0xd2, 0x75, 0x67, 0xb6, 0x45, 0x3a,
	0xa4, 0xd7, 0xe8, 0x3f, 0xe6, 0x5b, 0x73, 0xf3, 0xb3, 0x95, 0xf0, 0xdc, 0x27, 0x88, 0x6e, 0x58,
	0xbb, 0x9f, 0x09, 0x65, 0x2f, 0x21, 0x4b, 0x21, 0x19, 0x22, 0xea, 0x08, 0xac, 0x56, 0x99, 0x05,
	0xf6, 0x9c, 0xd6, 0x7d, 0x63, 0x81, 0x7d, 0x3f, 0xb0, 0xcb, 0xbc, 0xbc, 0x54, 0xbe, 0x76, 0xd5,
	0xc1, 0xad, 0x1f, 0x83, 0xbd, 0x4f, 0xa4, 0x7a, 0x48, 0xa2, 0xa0, 0x67, 0x2f, 0xe8, 0xff, 0x53,
	0x90, 0x09, 0x18, 0xdb, 0xaa, 0x76, 0x6a, 0xbd, 0x46, 0xff, 0x24, 0x58, 0xa5, 0x4e, 0xcb, 0x6e,
	0xca, 0xb9, 0xf3, 0xa1, 0x53, 0xbc, 0x91, 0xf3, 0x1c, 0x2e, 0x75, 0x39, 0x87, 0x68, 0x65, 0x62,
	0x8c, 0xfe, 0x37, 0x51, 0xc9, 0xb2, 0x55, 0xeb, 0x90, 0xde, 0xed, 0xc8, 0x9d, 0xbb, 0x23, 0xba,
	0x7f, 0x39, 0xdb, 0xe8, 0xef, 0x1f, 0x6f, 0xe9, 0x7e, 0x27, 0xb4, 0x19, 0x06, 0x1a, 0x88, 0x4f,
	0x7e, 0x4b, 0xcc, 0xb8, 0x7f, 0x43, 0x6e, 0x74, 0xcc, 0x7d, 0xda, 0x75, 0xc6, 0x2b, 0x7a, 0x90,
	0xb8, 0x99, 0x8d, 0x4d, 0xb0, 0xb7, 0xaa, 0xce, 0x74, 0xba, 0xe3, 0x09, 0xfe, 0x9c, 0xf0, 0xb0,
	0x12, 0xed, 0x7b, 0xc6, 0xba, 0x83, 0x21, 0x6d, 0xa8, 0xd9, 0x2f, 0x62, 0xcd, 0x11, 0x1f, 0xed,
	0x20, 0x6e, 0xce, 0x63, 0x58, 0x89, 0xa8, 0x5a, 0x67, 0x19, 0x1c, 0xd0, 0xa6, 0xdb, 0xc9, 0x15,
	0xab, 0x0f, 0xb4, 0x79, 0x96, 0xe3, 0x54, 0x99, 0xf4, 0xbd, 0x5b, 0x3b, 0x76, 0x45, 0xf7, 0x5c,
	0x7c, 0x76, 0xbc, 0x83, 0x7f, 0x73, 0xdb, 0xda, 0x27, 0x7f, 0x17, 0xf9, 0x4b, 0xba, 0x95, 0xc1,
	0xc5, 0x97, 0x0f, 0x5f, 0xbf, 0xd5, 0xab, 0x87, 0x84, 0x1e, 0xa7, 0xca, 0x7b, 0xb4, 0x51, 0x8b,
	0xe5, 0x76, 0xfb, 0xe0, 0xee, 0x45, 0xf8, 0x79, 0xcb, 0xde, 0x46, 0xe5, 0x9e, 0x8f, 0xc8, 0x47,
	0x42, 0x26, 0x75, 0xb7, 0xf3, 0xcf, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x1b, 0x7b, 0xd5,
	0xf3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.auth.v2.Authorization/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

// UnimplementedAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (*UnimplementedAuthorizationServer) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v2.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v2.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v2/external_auth.proto",
}

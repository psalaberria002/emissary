// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/v4alpha/filter.proto

package envoy_config_cluster_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Filter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TypedConfig          *any.Any `protobuf:"bytes,2,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_d31bb1ef5b95b511, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetTypedConfig() *any.Any {
	if m != nil {
		return m.TypedConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "envoy.config.cluster.v4alpha.Filter")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v4alpha/filter.proto", fileDescriptor_d31bb1ef5b95b511)
}

var fileDescriptor_d31bb1ef5b95b511 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x52, 0x47, 0xc5, 0xd4, 0x83, 0x14, 0xc1, 0x39, 0x75, 0x54, 0x41, 0x98, 0x1e,
	0x12, 0x70, 0x82, 0xe0, 0xcd, 0x0e, 0x3c, 0x8f, 0x7d, 0x01, 0x79, 0x5b, 0xd3, 0x1a, 0xa8, 0xef,
	0x95, 0x34, 0x2d, 0xf6, 0xe6, 0x45, 0xf0, 0x33, 0xf8, 0x51, 0xbc, 0x0b, 0x5e, 0xfd, 0x3a, 0x9e,
	0xc4, 0xa4, 0x3b, 0x29, 0xde, 0x42, 0xfe, 0xbf, 0xc7, 0xfb, 0x27, 0xfc, 0x4c, 0x61, 0x4b, 0x9d,
	0x5c, 0x11, 0xe6, 0xba, 0x90, 0xab, 0xb2, 0xa9, 0xad, 0x32, 0xb2, 0xbd, 0x84, 0xb2, 0xba, 0x07,
	0x99, 0xeb, 0xd2, 0x2a, 0x23, 0x2a, 0x43, 0x96, 0xe2, 0x43, 0x47, 0x85, 0xa7, 0xa2, 0xa7, 0xa2,
	0xa7, 0xa3, 0xfd, 0x82, 0xa8, 0x28, 0x95, 0x74, 0x76, 0xd9, 0xe4, 0x12, 0xb0, 0xf3, 0x83, 0xa3,
	0xa3, 0x26, 0xab, 0x40, 0x02, 0x22, 0x59, 0xb0, 0x9a, 0xb0, 0x96, 0xb5, 0x05, 0xdb, 0xd4, 0x7d,
	0x7c, 0xfc, 0x2b, 0x6e, 0x95, 0xa9, 0x35, 0xa1, 0xc6, 0xa2, 0x27, 0x7b, 0x2d, 0x94, 0x3a, 0x03,
	0xab, 0xe4, 0xfa, 0xe0, 0x83, 0x93, 0x67, 0xc6, 0xc3, 0x5b, 0x57, 0x32, 0x3e, 0xe0, 0x03, 0x84,
	0x07, 0x35, 0x64, 0x09, 0x9b, 0x6c, 0xa5, 0x9b, 0x5f, 0xe9, 0xc0, 0x04, 0x09, 0x5b, 0xb8, 0xcb,
	0xf8, 0x8a, 0x6f, 0xdb, 0xae, 0x52, 0xd9, 0x9d, 0x6f, 0x3f, 0x0c, 0x12, 0x36, 0x89, 0x2e, 0x76,
	0x85, 0x2f, 0x2d, 0xd6, 0xa5, 0xc5, 0x0d, 0x76, 0x8b, 0xc8, 0xc9, 0x99, 0x83, 0xd7, 0xa7, 0xaf,
	0xef, 0x2f, 0xe3, 0x84, 0x8f, 0xff, 0x7e, 0xfb, 0x54, 0xf8, 0xe5, 0xe9, 0xec, 0xed, 0xe9, 0xe3,
	0x33, 0x0c, 0x76, 0x36, 0xf8, 0xb9, 0x26, 0xe1, 0x70, 0x65, 0xe8, 0xb1, 0x13, 0xff, 0xfd, 0x59,
	0x1a, 0xf9, 0xe9, 0xf9, 0xcf, 0xf6, 0x39, 0x5b, 0x86, 0xae, 0xc6, 0xf4, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x39, 0xce, 0x8f, 0x87, 0x94, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/dns/v3/dns_table.proto

package envoy_data_dns_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/datawire/ambassador/pkg/api/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type DnsTable struct {
	ExternalRetryCount   uint32                       `protobuf:"varint,1,opt,name=external_retry_count,json=externalRetryCount,proto3" json:"external_retry_count,omitempty"`
	VirtualDomains       []*DnsTable_DnsVirtualDomain `protobuf:"bytes,2,rep,name=virtual_domains,json=virtualDomains,proto3" json:"virtual_domains,omitempty"`
	KnownSuffixes        []*v3.StringMatcher          `protobuf:"bytes,3,rep,name=known_suffixes,json=knownSuffixes,proto3" json:"known_suffixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DnsTable) Reset()         { *m = DnsTable{} }
func (m *DnsTable) String() string { return proto.CompactTextString(m) }
func (*DnsTable) ProtoMessage()    {}
func (*DnsTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0588926abfe2398, []int{0}
}

func (m *DnsTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable.Unmarshal(m, b)
}
func (m *DnsTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable.Marshal(b, m, deterministic)
}
func (m *DnsTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable.Merge(m, src)
}
func (m *DnsTable) XXX_Size() int {
	return xxx_messageInfo_DnsTable.Size(m)
}
func (m *DnsTable) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable proto.InternalMessageInfo

func (m *DnsTable) GetExternalRetryCount() uint32 {
	if m != nil {
		return m.ExternalRetryCount
	}
	return 0
}

func (m *DnsTable) GetVirtualDomains() []*DnsTable_DnsVirtualDomain {
	if m != nil {
		return m.VirtualDomains
	}
	return nil
}

func (m *DnsTable) GetKnownSuffixes() []*v3.StringMatcher {
	if m != nil {
		return m.KnownSuffixes
	}
	return nil
}

type DnsTable_AddressList struct {
	Address              []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnsTable_AddressList) Reset()         { *m = DnsTable_AddressList{} }
func (m *DnsTable_AddressList) String() string { return proto.CompactTextString(m) }
func (*DnsTable_AddressList) ProtoMessage()    {}
func (*DnsTable_AddressList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0588926abfe2398, []int{0, 0}
}

func (m *DnsTable_AddressList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_AddressList.Unmarshal(m, b)
}
func (m *DnsTable_AddressList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_AddressList.Marshal(b, m, deterministic)
}
func (m *DnsTable_AddressList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_AddressList.Merge(m, src)
}
func (m *DnsTable_AddressList) XXX_Size() int {
	return xxx_messageInfo_DnsTable_AddressList.Size(m)
}
func (m *DnsTable_AddressList) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_AddressList.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_AddressList proto.InternalMessageInfo

func (m *DnsTable_AddressList) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

type DnsTable_DnsEndpoint struct {
	// Types that are valid to be assigned to EndpointConfig:
	//	*DnsTable_DnsEndpoint_AddressList
	EndpointConfig       isDnsTable_DnsEndpoint_EndpointConfig `protobuf_oneof:"endpoint_config"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DnsTable_DnsEndpoint) Reset()         { *m = DnsTable_DnsEndpoint{} }
func (m *DnsTable_DnsEndpoint) String() string { return proto.CompactTextString(m) }
func (*DnsTable_DnsEndpoint) ProtoMessage()    {}
func (*DnsTable_DnsEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0588926abfe2398, []int{0, 1}
}

func (m *DnsTable_DnsEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Unmarshal(m, b)
}
func (m *DnsTable_DnsEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Marshal(b, m, deterministic)
}
func (m *DnsTable_DnsEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_DnsEndpoint.Merge(m, src)
}
func (m *DnsTable_DnsEndpoint) XXX_Size() int {
	return xxx_messageInfo_DnsTable_DnsEndpoint.Size(m)
}
func (m *DnsTable_DnsEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_DnsEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_DnsEndpoint proto.InternalMessageInfo

type isDnsTable_DnsEndpoint_EndpointConfig interface {
	isDnsTable_DnsEndpoint_EndpointConfig()
}

type DnsTable_DnsEndpoint_AddressList struct {
	AddressList *DnsTable_AddressList `protobuf:"bytes,1,opt,name=address_list,json=addressList,proto3,oneof"`
}

func (*DnsTable_DnsEndpoint_AddressList) isDnsTable_DnsEndpoint_EndpointConfig() {}

func (m *DnsTable_DnsEndpoint) GetEndpointConfig() isDnsTable_DnsEndpoint_EndpointConfig {
	if m != nil {
		return m.EndpointConfig
	}
	return nil
}

func (m *DnsTable_DnsEndpoint) GetAddressList() *DnsTable_AddressList {
	if x, ok := m.GetEndpointConfig().(*DnsTable_DnsEndpoint_AddressList); ok {
		return x.AddressList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DnsTable_DnsEndpoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DnsTable_DnsEndpoint_AddressList)(nil),
	}
}

type DnsTable_DnsVirtualDomain struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Endpoint             *DnsTable_DnsEndpoint `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AnswerTtl            *duration.Duration    `protobuf:"bytes,3,opt,name=answer_ttl,json=answerTtl,proto3" json:"answer_ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DnsTable_DnsVirtualDomain) Reset()         { *m = DnsTable_DnsVirtualDomain{} }
func (m *DnsTable_DnsVirtualDomain) String() string { return proto.CompactTextString(m) }
func (*DnsTable_DnsVirtualDomain) ProtoMessage()    {}
func (*DnsTable_DnsVirtualDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0588926abfe2398, []int{0, 2}
}

func (m *DnsTable_DnsVirtualDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Unmarshal(m, b)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Marshal(b, m, deterministic)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsTable_DnsVirtualDomain.Merge(m, src)
}
func (m *DnsTable_DnsVirtualDomain) XXX_Size() int {
	return xxx_messageInfo_DnsTable_DnsVirtualDomain.Size(m)
}
func (m *DnsTable_DnsVirtualDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsTable_DnsVirtualDomain.DiscardUnknown(m)
}

var xxx_messageInfo_DnsTable_DnsVirtualDomain proto.InternalMessageInfo

func (m *DnsTable_DnsVirtualDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsTable_DnsVirtualDomain) GetEndpoint() *DnsTable_DnsEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *DnsTable_DnsVirtualDomain) GetAnswerTtl() *duration.Duration {
	if m != nil {
		return m.AnswerTtl
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsTable)(nil), "envoy.data.dns.v3.DnsTable")
	proto.RegisterType((*DnsTable_AddressList)(nil), "envoy.data.dns.v3.DnsTable.AddressList")
	proto.RegisterType((*DnsTable_DnsEndpoint)(nil), "envoy.data.dns.v3.DnsTable.DnsEndpoint")
	proto.RegisterType((*DnsTable_DnsVirtualDomain)(nil), "envoy.data.dns.v3.DnsTable.DnsVirtualDomain")
}

func init() { proto.RegisterFile("envoy/data/dns/v3/dns_table.proto", fileDescriptor_c0588926abfe2398) }

var fileDescriptor_c0588926abfe2398 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xbf, 0x6f, 0xd3, 0x40,
	0x18, 0xed, 0xd9, 0xfd, 0xe1, 0x9e, 0x69, 0x1b, 0x2c, 0x04, 0x26, 0x12, 0x6d, 0x13, 0x21, 0x88,
	0xf8, 0x61, 0x57, 0xc9, 0x80, 0x94, 0x0d, 0x37, 0x48, 0x48, 0x14, 0xa9, 0x72, 0x2b, 0x56, 0xeb,
	0x12, 0x5f, 0xd2, 0x53, 0x9d, 0x3b, 0xeb, 0xee, 0xec, 0x26, 0x1b, 0x23, 0x73, 0x47, 0x26, 0x36,
	0x24, 0xfe, 0x04, 0x26, 0x16, 0x24, 0x56, 0xfe, 0x17, 0x06, 0x94, 0x09, 0xdd, 0xd9, 0x26, 0xa1,
	0x91, 0x68, 0xa7, 0xdc, 0x7d, 0xef, 0xde, 0xfb, 0xbe, 0xf7, 0xe2, 0x0f, 0x36, 0x30, 0xcd, 0xd9,
	0xd4, 0x8f, 0x91, 0x44, 0x7e, 0x4c, 0x85, 0x9f, 0x77, 0xd4, 0x4f, 0x24, 0x51, 0x3f, 0xc1, 0x5e,
	0xca, 0x99, 0x64, 0xce, 0x6d, 0xfd, 0xc4, 0x53, 0x4f, 0xbc, 0x98, 0x0a, 0x2f, 0xef, 0xd4, 0x9b,
	0x05, 0x4b, 0x4e, 0x53, 0xec, 0x8f, 0x91, 0x1c, 0x9c, 0x61, 0xae, 0x98, 0x42, 0x72, 0x42, 0x47,
	0x05, 0xad, 0xbe, 0x3b, 0x62, 0x6c, 0x94, 0x60, 0x5f, 0xdf, 0xfa, 0xd9, 0xd0, 0x8f, 0x33, 0x8e,
	0x24, 0x61, 0xb4, 0xc4, 0x1f, 0x64, 0x71, 0x8a, 0x7c, 0x44, 0x29, 0x93, 0xba, 0x2c, 0x7c, 0x21,
	0x91, 0xcc, 0x44, 0x09, 0x37, 0x96, 0xe0, 0x1c, 0x73, 0x41, 0x18, 0x9d, 0x77, 0xb8, 0x97, 0xa3,
	0x84, 0xc4, 0x48, 0x62, 0xbf, 0x3a, 0x14, 0x40, 0xf3, 0xd3, 0x3a, 0xb4, 0x7a, 0x54, 0x9c, 0x2a,
	0x13, 0xce, 0x01, 0xbc, 0x83, 0x27, 0x12, 0x73, 0x8a, 0x92, 0x88, 0x63, 0xc9, 0xa7, 0xd1, 0x80,
	0x65, 0x54, 0xba, 0x60, 0x1f, 0xb4, 0xb6, 0x42, 0xa7, 0xc2, 0x42, 0x05, 0x1d, 0x2a, 0xc4, 0x89,
	0xe0, 0x4e, 0x4e, 0xb8, 0xcc, 0x50, 0x12, 0xc5, 0x6c, 0x8c, 0x08, 0x15, 0xae, 0xb1, 0x6f, 0xb6,
	0xec, 0xf6, 0x33, 0x6f, 0x29, 0x0a, 0xaf, 0xea, 0xa3, 0x0e, 0xef, 0x0a, 0x56, 0x4f, 0x93, 0x02,
	0x6b, 0x16, 0xac, 0x5d, 0x02, 0xc3, 0x02, 0xe1, 0x76, 0xbe, 0x08, 0x08, 0xe7, 0x0d, 0xdc, 0x3e,
	0xa7, 0xec, 0x82, 0x46, 0x22, 0x1b, 0x0e, 0xc9, 0x04, 0x0b, 0xd7, 0xd4, 0xfa, 0x0f, 0x4b, 0x7d,
	0x95, 0xab, 0x57, 0xe6, 0xaa, 0x7a, 0x9c, 0xe8, 0x5c, 0xdf, 0x16, 0x85, 0x70, 0x4b, 0x73, 0x4f,
	0x4a, 0x6a, 0xfd, 0x1c, 0xda, 0x2f, 0xe3, 0x98, 0x63, 0x21, 0x8e, 0x88, 0x90, 0x4e, 0x0b, 0x6e,
	0xa0, 0xe2, 0xea, 0x82, 0x7d, 0xb3, 0xb5, 0x19, 0x6c, 0xcf, 0x02, 0xfb, 0x12, 0x58, 0x16, 0x68,
	0xae, 0x72, 0xa3, 0x66, 0x86, 0x15, 0xdc, 0x6d, 0x7f, 0xfc, 0xfe, 0x61, 0xf7, 0x39, 0x7c, 0x7a,
	0xd5, 0x53, 0x1b, 0x25, 0xe9, 0x19, 0x9a, 0x1b, 0x5b, 0x50, 0xaf, 0x7f, 0x06, 0xd0, 0xee, 0x51,
	0xf1, 0x8a, 0xc6, 0x29, 0x23, 0x54, 0x3a, 0x47, 0xf0, 0x56, 0x29, 0x17, 0x25, 0x44, 0x14, 0xa1,
	0xda, 0xed, 0xc7, 0xff, 0xcb, 0x69, 0x41, 0xee, 0xf5, 0x4a, 0x68, 0xa3, 0xf9, 0xf5, 0xc6, 0x13,
	0x2d, 0x4c, 0x10, 0xdc, 0x85, 0x3b, 0xb8, 0x3c, 0x47, 0x03, 0x46, 0x87, 0x64, 0xe4, 0x98, 0xbf,
	0x03, 0x50, 0xff, 0x05, 0x60, 0xed, 0xea, 0x5f, 0xe2, 0xec, 0xc2, 0x55, 0x8a, 0xc6, 0x58, 0x8f,
	0xb9, 0x19, 0xc0, 0x59, 0xb0, 0xc1, 0xd7, 0x6a, 0xc6, 0x37, 0x00, 0x42, 0x5d, 0x77, 0x0e, 0xa1,
	0x55, 0x89, 0xb9, 0xc6, 0xf5, 0x56, 0x16, 0xe6, 0x08, 0xff, 0x12, 0x9d, 0x00, 0x42, 0x44, 0xc5,
	0x05, 0xe6, 0x91, 0x94, 0x89, 0x6b, 0x6a, 0x99, 0xfb, 0x5e, 0xb1, 0x0d, 0x5e, 0xb5, 0x0d, 0x5e,
	0xaf, 0xdc, 0x06, 0xfd, 0x99, 0x7c, 0x01, 0xc6, 0x93, 0x95, 0x70, 0xb3, 0xa0, 0x9d, 0xca, 0xa4,
	0xfb, 0x42, 0x25, 0xd1, 0x86, 0x07, 0x37, 0x48, 0xe2, 0x1f, 0x87, 0xdd, 0x47, 0x8a, 0xd8, 0x80,
	0x7b, 0xd7, 0x10, 0x83, 0xee, 0xd7, 0xf7, 0x3f, 0x7e, 0xae, 0x1b, 0x35, 0x03, 0xee, 0x11, 0x56,
	0x78, 0x4c, 0x39, 0x9b, 0x4c, 0x97, 0xed, 0x06, 0x5b, 0x15, 0xe9, 0x58, 0xcd, 0x7e, 0x0c, 0xfa,
	0xeb, 0xda, 0x44, 0xe7, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x0c, 0xbc, 0x3b, 0x3c, 0x04,
	0x00, 0x00,
}

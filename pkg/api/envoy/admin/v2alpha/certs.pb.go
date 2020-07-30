// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/certs.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Certificates struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{0}
}

func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificates.Unmarshal(m, b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return xxx_messageInfo_Certificates.Size(m)
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	CaCert               []*CertificateDetails `protobuf:"bytes,1,rep,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	CertChain            []*CertificateDetails `protobuf:"bytes,2,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{1}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetCaCert() []*CertificateDetails {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *Certificate) GetCertChain() []*CertificateDetails {
	if m != nil {
		return m.CertChain
	}
	return nil
}

type CertificateDetails struct {
	Path                 string                  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	SerialNumber         string                  `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	SubjectAltNames      []*SubjectAlternateName `protobuf:"bytes,3,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	DaysUntilExpiration  uint64                  `protobuf:"varint,4,opt,name=days_until_expiration,json=daysUntilExpiration,proto3" json:"days_until_expiration,omitempty"`
	ValidFrom            *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	ExpirationTime       *timestamp.Timestamp    `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CertificateDetails) Reset()         { *m = CertificateDetails{} }
func (m *CertificateDetails) String() string { return proto.CompactTextString(m) }
func (*CertificateDetails) ProtoMessage()    {}
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{2}
}

func (m *CertificateDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateDetails.Unmarshal(m, b)
}
func (m *CertificateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateDetails.Marshal(b, m, deterministic)
}
func (m *CertificateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateDetails.Merge(m, src)
}
func (m *CertificateDetails) XXX_Size() int {
	return xxx_messageInfo_CertificateDetails.Size(m)
}
func (m *CertificateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateDetails proto.InternalMessageInfo

func (m *CertificateDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CertificateDetails) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CertificateDetails) GetSubjectAltNames() []*SubjectAlternateName {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *CertificateDetails) GetDaysUntilExpiration() uint64 {
	if m != nil {
		return m.DaysUntilExpiration
	}
	return 0
}

func (m *CertificateDetails) GetValidFrom() *timestamp.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *CertificateDetails) GetExpirationTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

type SubjectAlternateName struct {
	// Types that are valid to be assigned to Name:
	//	*SubjectAlternateName_Dns
	//	*SubjectAlternateName_Uri
	//	*SubjectAlternateName_IpAddress
	Name                 isSubjectAlternateName_Name `protobuf_oneof:"name"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SubjectAlternateName) Reset()         { *m = SubjectAlternateName{} }
func (m *SubjectAlternateName) String() string { return proto.CompactTextString(m) }
func (*SubjectAlternateName) ProtoMessage()    {}
func (*SubjectAlternateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7cd1796e05ff7fa, []int{3}
}

func (m *SubjectAlternateName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubjectAlternateName.Unmarshal(m, b)
}
func (m *SubjectAlternateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubjectAlternateName.Marshal(b, m, deterministic)
}
func (m *SubjectAlternateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectAlternateName.Merge(m, src)
}
func (m *SubjectAlternateName) XXX_Size() int {
	return xxx_messageInfo_SubjectAlternateName.Size(m)
}
func (m *SubjectAlternateName) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectAlternateName.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectAlternateName proto.InternalMessageInfo

type isSubjectAlternateName_Name interface {
	isSubjectAlternateName_Name()
}

type SubjectAlternateName_Dns struct {
	Dns string `protobuf:"bytes,1,opt,name=dns,proto3,oneof"`
}

type SubjectAlternateName_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

type SubjectAlternateName_IpAddress struct {
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3,oneof"`
}

func (*SubjectAlternateName_Dns) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_Uri) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_IpAddress) isSubjectAlternateName_Name() {}

func (m *SubjectAlternateName) GetName() isSubjectAlternateName_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SubjectAlternateName) GetDns() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Dns); ok {
		return x.Dns
	}
	return ""
}

func (m *SubjectAlternateName) GetUri() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Uri); ok {
		return x.Uri
	}
	return ""
}

func (m *SubjectAlternateName) GetIpAddress() string {
	if x, ok := m.GetName().(*SubjectAlternateName_IpAddress); ok {
		return x.IpAddress
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubjectAlternateName) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubjectAlternateName_Dns)(nil),
		(*SubjectAlternateName_Uri)(nil),
		(*SubjectAlternateName_IpAddress)(nil),
	}
}

func init() {
	proto.RegisterType((*Certificates)(nil), "envoy.admin.v2alpha.Certificates")
	proto.RegisterType((*Certificate)(nil), "envoy.admin.v2alpha.Certificate")
	proto.RegisterType((*CertificateDetails)(nil), "envoy.admin.v2alpha.CertificateDetails")
	proto.RegisterType((*SubjectAlternateName)(nil), "envoy.admin.v2alpha.SubjectAlternateName")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/certs.proto", fileDescriptor_c7cd1796e05ff7fa) }

var fileDescriptor_c7cd1796e05ff7fa = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x66, 0x93, 0x10, 0x94, 0x49, 0xa0, 0xe0, 0x82, 0xb4, 0x8a, 0x84, 0xb2, 0x84, 0x03, 0xe1,
	0xb2, 0x2b, 0x85, 0x53, 0x6f, 0x34, 0x29, 0x15, 0xa7, 0xaa, 0x5a, 0xda, 0xb3, 0x35, 0xd9, 0x75,
	0x12, 0xa3, 0x5d, 0xdb, 0xb2, 0xbd, 0x51, 0x73, 0xe3, 0x2d, 0x78, 0x17, 0x9e, 0x80, 0x2b, 0x37,
	0x1e, 0x07, 0xd9, 0x4e, 0xdb, 0x20, 0x22, 0x55, 0xdc, 0x3c, 0xdf, 0xcf, 0x78, 0xfc, 0x79, 0x60,
	0xc4, 0xc4, 0x46, 0x6e, 0x33, 0x2c, 0x6b, 0x2e, 0xb2, 0xcd, 0x14, 0x2b, 0xb5, 0xc6, 0xac, 0x60,
	0xda, 0x9a, 0x54, 0x69, 0x69, 0x25, 0x39, 0xf6, 0x82, 0xd4, 0x0b, 0xd2, 0x9d, 0x60, 0x38, 0x5a,
	0x49, 0xb9, 0xaa, 0x58, 0xe6, 0x25, 0x8b, 0x66, 0x99, 0x59, 0x5e, 0x33, 0x63, 0xb1, 0x56, 0xc1,
	0x35, 0x7c, 0xdd, 0x94, 0x0a, 0x33, 0x14, 0x42, 0x5a, 0xb4, 0x5c, 0x0a, 0x93, 0x19, 0x8b, 0xb6,
	0xd9, 0x35, 0x1d, 0x5f, 0xc1, 0x60, 0xce, 0xb4, 0xe5, 0x4b, 0x5e, 0xa0, 0x65, 0x86, 0x9c, 0xc1,
	0xa0, 0xd8, 0xab, 0xe3, 0x28, 0x69, 0x4f, 0xfa, 0xd3, 0x24, 0x3d, 0x70, 0x77, 0xba, 0x67, 0xcc,
	0xff, 0x72, 0x8d, 0xbf, 0x47, 0xd0, 0xdf, 0x63, 0xc9, 0x47, 0x78, 0x52, 0x20, 0x75, 0x92, 0x5d,
	0xc3, 0x77, 0x0f, 0x35, 0x3c, 0x63, 0x16, 0x79, 0x65, 0xf2, 0x6e, 0x81, 0x0e, 0x25, 0xe7, 0x00,
	0xce, 0x4e, 0x8b, 0x35, 0x72, 0x11, 0xb7, 0xfe, 0xaf, 0x49, 0xcf, 0x59, 0xe7, 0xce, 0x39, 0xfe,
	0xdd, 0x02, 0xf2, 0xaf, 0x82, 0x10, 0xe8, 0x28, 0xb4, 0xeb, 0x38, 0x4a, 0xa2, 0x49, 0x2f, 0xf7,
	0x67, 0xf2, 0x16, 0x9e, 0x1a, 0xa6, 0x39, 0x56, 0x54, 0x34, 0xf5, 0x82, 0xe9, 0xb8, 0xe5, 0xc9,
	0x41, 0x00, 0x2f, 0x3c, 0x46, 0xae, 0xe1, 0x85, 0x69, 0x16, 0x5f, 0x59, 0x61, 0x29, 0x56, 0x96,
	0x0a, 0xac, 0x99, 0x89, 0xdb, 0x7e, 0xbc, 0xf7, 0x07, 0xc7, 0xfb, 0x12, 0xd4, 0xa7, 0x95, 0x65,
	0x5a, 0xa0, 0x65, 0x17, 0x58, 0xb3, 0xfc, 0xc8, 0xdc, 0xa1, 0xae, 0x36, 0x64, 0x0a, 0xaf, 0x4a,
	0xdc, 0x1a, 0xda, 0x08, 0xcb, 0x2b, 0xca, 0x6e, 0x14, 0xd7, 0xfe, 0xf7, 0xe2, 0x4e, 0x12, 0x4d,
	0x3a, 0xf9, 0xb1, 0x23, 0xaf, 0x1d, 0xf7, 0xe9, 0x8e, 0x22, 0x27, 0x00, 0x1b, 0xac, 0x78, 0x49,
	0x97, 0x5a, 0xd6, 0xf1, 0xe3, 0x24, 0x9a, 0xf4, 0xa7, 0xc3, 0x34, 0xec, 0x47, 0x7a, 0xbb, 0x1f,
	0xe9, 0xd5, 0xed, 0x7e, 0xe4, 0x3d, 0xaf, 0x3e, 0xd7, 0xb2, 0x26, 0x73, 0x38, 0xba, 0xbf, 0x83,
	0xba, 0x15, 0x8a, 0xbb, 0x0f, 0xfa, 0x9f, 0xdd, 0x5b, 0x1c, 0x38, 0x5e, 0xc1, 0xcb, 0x43, 0x8f,
	0x23, 0x04, 0xda, 0xa5, 0x30, 0x21, 0xda, 0xcf, 0x8f, 0x72, 0x57, 0x38, 0xac, 0xd1, 0x3c, 0x24,
	0xea, 0xb0, 0x46, 0x73, 0x32, 0x02, 0xe0, 0x8a, 0x62, 0x59, 0x6a, 0x66, 0x5c, 0x86, 0x81, 0xea,
	0x71, 0x75, 0x1a, 0xa0, 0x59, 0x17, 0x3a, 0x2e, 0xdf, 0xd9, 0xc9, 0x8f, 0x6f, 0x3f, 0x7f, 0x75,
	0x5b, 0xcf, 0x23, 0x78, 0xc3, 0x65, 0x08, 0x59, 0x69, 0x79, 0xb3, 0x3d, 0x94, 0xf7, 0x0c, 0xdc,
	0x6f, 0x9b, 0x4b, 0x37, 0xfe, 0x65, 0xb4, 0xe8, 0xfa, 0x77, 0x7c, 0xf8, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x70, 0xd0, 0x3a, 0xc0, 0x6d, 0x03, 0x00, 0x00,
}

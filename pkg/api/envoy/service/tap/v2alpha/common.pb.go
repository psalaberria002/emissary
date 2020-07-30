// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v2alpha/common.proto

package envoy_service_tap_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v2/core"
	route "github.com/datawire/ambassador/pkg/api/envoy/api/v2/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type OutputSink_Format int32

const (
	OutputSink_JSON_BODY_AS_BYTES            OutputSink_Format = 0
	OutputSink_JSON_BODY_AS_STRING           OutputSink_Format = 1
	OutputSink_PROTO_BINARY                  OutputSink_Format = 2
	OutputSink_PROTO_BINARY_LENGTH_DELIMITED OutputSink_Format = 3
	OutputSink_PROTO_TEXT                    OutputSink_Format = 4
)

var OutputSink_Format_name = map[int32]string{
	0: "JSON_BODY_AS_BYTES",
	1: "JSON_BODY_AS_STRING",
	2: "PROTO_BINARY",
	3: "PROTO_BINARY_LENGTH_DELIMITED",
	4: "PROTO_TEXT",
}

var OutputSink_Format_value = map[string]int32{
	"JSON_BODY_AS_BYTES":            0,
	"JSON_BODY_AS_STRING":           1,
	"PROTO_BINARY":                  2,
	"PROTO_BINARY_LENGTH_DELIMITED": 3,
	"PROTO_TEXT":                    4,
}

func (x OutputSink_Format) String() string {
	return proto.EnumName(OutputSink_Format_name, int32(x))
}

func (OutputSink_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{4, 0}
}

type TapConfig struct {
	MatchConfig          *MatchPredicate                `protobuf:"bytes,1,opt,name=match_config,json=matchConfig,proto3" json:"match_config,omitempty"`
	OutputConfig         *OutputConfig                  `protobuf:"bytes,2,opt,name=output_config,json=outputConfig,proto3" json:"output_config,omitempty"`
	TapEnabled           *core.RuntimeFractionalPercent `protobuf:"bytes,3,opt,name=tap_enabled,json=tapEnabled,proto3" json:"tap_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TapConfig) Reset()         { *m = TapConfig{} }
func (m *TapConfig) String() string { return proto.CompactTextString(m) }
func (*TapConfig) ProtoMessage()    {}
func (*TapConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{0}
}

func (m *TapConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapConfig.Unmarshal(m, b)
}
func (m *TapConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapConfig.Marshal(b, m, deterministic)
}
func (m *TapConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapConfig.Merge(m, src)
}
func (m *TapConfig) XXX_Size() int {
	return xxx_messageInfo_TapConfig.Size(m)
}
func (m *TapConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TapConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TapConfig proto.InternalMessageInfo

func (m *TapConfig) GetMatchConfig() *MatchPredicate {
	if m != nil {
		return m.MatchConfig
	}
	return nil
}

func (m *TapConfig) GetOutputConfig() *OutputConfig {
	if m != nil {
		return m.OutputConfig
	}
	return nil
}

func (m *TapConfig) GetTapEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.TapEnabled
	}
	return nil
}

type MatchPredicate struct {
	// Types that are valid to be assigned to Rule:
	//	*MatchPredicate_OrMatch
	//	*MatchPredicate_AndMatch
	//	*MatchPredicate_NotMatch
	//	*MatchPredicate_AnyMatch
	//	*MatchPredicate_HttpRequestHeadersMatch
	//	*MatchPredicate_HttpRequestTrailersMatch
	//	*MatchPredicate_HttpResponseHeadersMatch
	//	*MatchPredicate_HttpResponseTrailersMatch
	Rule                 isMatchPredicate_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MatchPredicate) Reset()         { *m = MatchPredicate{} }
func (m *MatchPredicate) String() string { return proto.CompactTextString(m) }
func (*MatchPredicate) ProtoMessage()    {}
func (*MatchPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{1}
}

func (m *MatchPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchPredicate.Unmarshal(m, b)
}
func (m *MatchPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchPredicate.Marshal(b, m, deterministic)
}
func (m *MatchPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchPredicate.Merge(m, src)
}
func (m *MatchPredicate) XXX_Size() int {
	return xxx_messageInfo_MatchPredicate.Size(m)
}
func (m *MatchPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_MatchPredicate proto.InternalMessageInfo

type isMatchPredicate_Rule interface {
	isMatchPredicate_Rule()
}

type MatchPredicate_OrMatch struct {
	OrMatch *MatchPredicate_MatchSet `protobuf:"bytes,1,opt,name=or_match,json=orMatch,proto3,oneof"`
}

type MatchPredicate_AndMatch struct {
	AndMatch *MatchPredicate_MatchSet `protobuf:"bytes,2,opt,name=and_match,json=andMatch,proto3,oneof"`
}

type MatchPredicate_NotMatch struct {
	NotMatch *MatchPredicate `protobuf:"bytes,3,opt,name=not_match,json=notMatch,proto3,oneof"`
}

type MatchPredicate_AnyMatch struct {
	AnyMatch bool `protobuf:"varint,4,opt,name=any_match,json=anyMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestHeadersMatch struct {
	HttpRequestHeadersMatch *HttpHeadersMatch `protobuf:"bytes,5,opt,name=http_request_headers_match,json=httpRequestHeadersMatch,proto3,oneof"`
}

type MatchPredicate_HttpRequestTrailersMatch struct {
	HttpRequestTrailersMatch *HttpHeadersMatch `protobuf:"bytes,6,opt,name=http_request_trailers_match,json=httpRequestTrailersMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseHeadersMatch struct {
	HttpResponseHeadersMatch *HttpHeadersMatch `protobuf:"bytes,7,opt,name=http_response_headers_match,json=httpResponseHeadersMatch,proto3,oneof"`
}

type MatchPredicate_HttpResponseTrailersMatch struct {
	HttpResponseTrailersMatch *HttpHeadersMatch `protobuf:"bytes,8,opt,name=http_response_trailers_match,json=httpResponseTrailersMatch,proto3,oneof"`
}

func (*MatchPredicate_OrMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_AndMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_NotMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_AnyMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestHeadersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpRequestTrailersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseHeadersMatch) isMatchPredicate_Rule() {}

func (*MatchPredicate_HttpResponseTrailersMatch) isMatchPredicate_Rule() {}

func (m *MatchPredicate) GetRule() isMatchPredicate_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *MatchPredicate) GetOrMatch() *MatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*MatchPredicate_OrMatch); ok {
		return x.OrMatch
	}
	return nil
}

func (m *MatchPredicate) GetAndMatch() *MatchPredicate_MatchSet {
	if x, ok := m.GetRule().(*MatchPredicate_AndMatch); ok {
		return x.AndMatch
	}
	return nil
}

func (m *MatchPredicate) GetNotMatch() *MatchPredicate {
	if x, ok := m.GetRule().(*MatchPredicate_NotMatch); ok {
		return x.NotMatch
	}
	return nil
}

func (m *MatchPredicate) GetAnyMatch() bool {
	if x, ok := m.GetRule().(*MatchPredicate_AnyMatch); ok {
		return x.AnyMatch
	}
	return false
}

func (m *MatchPredicate) GetHttpRequestHeadersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpRequestHeadersMatch); ok {
		return x.HttpRequestHeadersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpRequestTrailersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpRequestTrailersMatch); ok {
		return x.HttpRequestTrailersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpResponseHeadersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpResponseHeadersMatch); ok {
		return x.HttpResponseHeadersMatch
	}
	return nil
}

func (m *MatchPredicate) GetHttpResponseTrailersMatch() *HttpHeadersMatch {
	if x, ok := m.GetRule().(*MatchPredicate_HttpResponseTrailersMatch); ok {
		return x.HttpResponseTrailersMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MatchPredicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MatchPredicate_OrMatch)(nil),
		(*MatchPredicate_AndMatch)(nil),
		(*MatchPredicate_NotMatch)(nil),
		(*MatchPredicate_AnyMatch)(nil),
		(*MatchPredicate_HttpRequestHeadersMatch)(nil),
		(*MatchPredicate_HttpRequestTrailersMatch)(nil),
		(*MatchPredicate_HttpResponseHeadersMatch)(nil),
		(*MatchPredicate_HttpResponseTrailersMatch)(nil),
	}
}

type MatchPredicate_MatchSet struct {
	Rules                []*MatchPredicate `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MatchPredicate_MatchSet) Reset()         { *m = MatchPredicate_MatchSet{} }
func (m *MatchPredicate_MatchSet) String() string { return proto.CompactTextString(m) }
func (*MatchPredicate_MatchSet) ProtoMessage()    {}
func (*MatchPredicate_MatchSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{1, 0}
}

func (m *MatchPredicate_MatchSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchPredicate_MatchSet.Unmarshal(m, b)
}
func (m *MatchPredicate_MatchSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchPredicate_MatchSet.Marshal(b, m, deterministic)
}
func (m *MatchPredicate_MatchSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchPredicate_MatchSet.Merge(m, src)
}
func (m *MatchPredicate_MatchSet) XXX_Size() int {
	return xxx_messageInfo_MatchPredicate_MatchSet.Size(m)
}
func (m *MatchPredicate_MatchSet) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchPredicate_MatchSet.DiscardUnknown(m)
}

var xxx_messageInfo_MatchPredicate_MatchSet proto.InternalMessageInfo

func (m *MatchPredicate_MatchSet) GetRules() []*MatchPredicate {
	if m != nil {
		return m.Rules
	}
	return nil
}

type HttpHeadersMatch struct {
	Headers              []*route.HeaderMatcher `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HttpHeadersMatch) Reset()         { *m = HttpHeadersMatch{} }
func (m *HttpHeadersMatch) String() string { return proto.CompactTextString(m) }
func (*HttpHeadersMatch) ProtoMessage()    {}
func (*HttpHeadersMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{2}
}

func (m *HttpHeadersMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpHeadersMatch.Unmarshal(m, b)
}
func (m *HttpHeadersMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpHeadersMatch.Marshal(b, m, deterministic)
}
func (m *HttpHeadersMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpHeadersMatch.Merge(m, src)
}
func (m *HttpHeadersMatch) XXX_Size() int {
	return xxx_messageInfo_HttpHeadersMatch.Size(m)
}
func (m *HttpHeadersMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpHeadersMatch.DiscardUnknown(m)
}

var xxx_messageInfo_HttpHeadersMatch proto.InternalMessageInfo

func (m *HttpHeadersMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type OutputConfig struct {
	Sinks                []*OutputSink         `protobuf:"bytes,1,rep,name=sinks,proto3" json:"sinks,omitempty"`
	MaxBufferedRxBytes   *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_buffered_rx_bytes,json=maxBufferedRxBytes,proto3" json:"max_buffered_rx_bytes,omitempty"`
	MaxBufferedTxBytes   *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_buffered_tx_bytes,json=maxBufferedTxBytes,proto3" json:"max_buffered_tx_bytes,omitempty"`
	Streaming            bool                  `protobuf:"varint,4,opt,name=streaming,proto3" json:"streaming,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OutputConfig) Reset()         { *m = OutputConfig{} }
func (m *OutputConfig) String() string { return proto.CompactTextString(m) }
func (*OutputConfig) ProtoMessage()    {}
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{3}
}

func (m *OutputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputConfig.Unmarshal(m, b)
}
func (m *OutputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputConfig.Marshal(b, m, deterministic)
}
func (m *OutputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputConfig.Merge(m, src)
}
func (m *OutputConfig) XXX_Size() int {
	return xxx_messageInfo_OutputConfig.Size(m)
}
func (m *OutputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputConfig proto.InternalMessageInfo

func (m *OutputConfig) GetSinks() []*OutputSink {
	if m != nil {
		return m.Sinks
	}
	return nil
}

func (m *OutputConfig) GetMaxBufferedRxBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxBufferedRxBytes
	}
	return nil
}

func (m *OutputConfig) GetMaxBufferedTxBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxBufferedTxBytes
	}
	return nil
}

func (m *OutputConfig) GetStreaming() bool {
	if m != nil {
		return m.Streaming
	}
	return false
}

type OutputSink struct {
	Format OutputSink_Format `protobuf:"varint,1,opt,name=format,proto3,enum=envoy.service.tap.v2alpha.OutputSink_Format" json:"format,omitempty"`
	// Types that are valid to be assigned to OutputSinkType:
	//	*OutputSink_StreamingAdmin
	//	*OutputSink_FilePerTap
	//	*OutputSink_StreamingGrpc
	OutputSinkType       isOutputSink_OutputSinkType `protobuf_oneof:"output_sink_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *OutputSink) Reset()         { *m = OutputSink{} }
func (m *OutputSink) String() string { return proto.CompactTextString(m) }
func (*OutputSink) ProtoMessage()    {}
func (*OutputSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{4}
}

func (m *OutputSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputSink.Unmarshal(m, b)
}
func (m *OutputSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputSink.Marshal(b, m, deterministic)
}
func (m *OutputSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputSink.Merge(m, src)
}
func (m *OutputSink) XXX_Size() int {
	return xxx_messageInfo_OutputSink.Size(m)
}
func (m *OutputSink) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputSink.DiscardUnknown(m)
}

var xxx_messageInfo_OutputSink proto.InternalMessageInfo

func (m *OutputSink) GetFormat() OutputSink_Format {
	if m != nil {
		return m.Format
	}
	return OutputSink_JSON_BODY_AS_BYTES
}

type isOutputSink_OutputSinkType interface {
	isOutputSink_OutputSinkType()
}

type OutputSink_StreamingAdmin struct {
	StreamingAdmin *StreamingAdminSink `protobuf:"bytes,2,opt,name=streaming_admin,json=streamingAdmin,proto3,oneof"`
}

type OutputSink_FilePerTap struct {
	FilePerTap *FilePerTapSink `protobuf:"bytes,3,opt,name=file_per_tap,json=filePerTap,proto3,oneof"`
}

type OutputSink_StreamingGrpc struct {
	StreamingGrpc *StreamingGrpcSink `protobuf:"bytes,4,opt,name=streaming_grpc,json=streamingGrpc,proto3,oneof"`
}

func (*OutputSink_StreamingAdmin) isOutputSink_OutputSinkType() {}

func (*OutputSink_FilePerTap) isOutputSink_OutputSinkType() {}

func (*OutputSink_StreamingGrpc) isOutputSink_OutputSinkType() {}

func (m *OutputSink) GetOutputSinkType() isOutputSink_OutputSinkType {
	if m != nil {
		return m.OutputSinkType
	}
	return nil
}

func (m *OutputSink) GetStreamingAdmin() *StreamingAdminSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_StreamingAdmin); ok {
		return x.StreamingAdmin
	}
	return nil
}

func (m *OutputSink) GetFilePerTap() *FilePerTapSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_FilePerTap); ok {
		return x.FilePerTap
	}
	return nil
}

func (m *OutputSink) GetStreamingGrpc() *StreamingGrpcSink {
	if x, ok := m.GetOutputSinkType().(*OutputSink_StreamingGrpc); ok {
		return x.StreamingGrpc
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutputSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutputSink_StreamingAdmin)(nil),
		(*OutputSink_FilePerTap)(nil),
		(*OutputSink_StreamingGrpc)(nil),
	}
}

type StreamingAdminSink struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingAdminSink) Reset()         { *m = StreamingAdminSink{} }
func (m *StreamingAdminSink) String() string { return proto.CompactTextString(m) }
func (*StreamingAdminSink) ProtoMessage()    {}
func (*StreamingAdminSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{5}
}

func (m *StreamingAdminSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingAdminSink.Unmarshal(m, b)
}
func (m *StreamingAdminSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingAdminSink.Marshal(b, m, deterministic)
}
func (m *StreamingAdminSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingAdminSink.Merge(m, src)
}
func (m *StreamingAdminSink) XXX_Size() int {
	return xxx_messageInfo_StreamingAdminSink.Size(m)
}
func (m *StreamingAdminSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingAdminSink.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingAdminSink proto.InternalMessageInfo

type FilePerTapSink struct {
	PathPrefix           string   `protobuf:"bytes,1,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilePerTapSink) Reset()         { *m = FilePerTapSink{} }
func (m *FilePerTapSink) String() string { return proto.CompactTextString(m) }
func (*FilePerTapSink) ProtoMessage()    {}
func (*FilePerTapSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{6}
}

func (m *FilePerTapSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilePerTapSink.Unmarshal(m, b)
}
func (m *FilePerTapSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilePerTapSink.Marshal(b, m, deterministic)
}
func (m *FilePerTapSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePerTapSink.Merge(m, src)
}
func (m *FilePerTapSink) XXX_Size() int {
	return xxx_messageInfo_FilePerTapSink.Size(m)
}
func (m *FilePerTapSink) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePerTapSink.DiscardUnknown(m)
}

var xxx_messageInfo_FilePerTapSink proto.InternalMessageInfo

func (m *FilePerTapSink) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

type StreamingGrpcSink struct {
	TapId                string            `protobuf:"bytes,1,opt,name=tap_id,json=tapId,proto3" json:"tap_id,omitempty"`
	GrpcService          *core.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StreamingGrpcSink) Reset()         { *m = StreamingGrpcSink{} }
func (m *StreamingGrpcSink) String() string { return proto.CompactTextString(m) }
func (*StreamingGrpcSink) ProtoMessage()    {}
func (*StreamingGrpcSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f6a60461d532a14, []int{7}
}

func (m *StreamingGrpcSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingGrpcSink.Unmarshal(m, b)
}
func (m *StreamingGrpcSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingGrpcSink.Marshal(b, m, deterministic)
}
func (m *StreamingGrpcSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingGrpcSink.Merge(m, src)
}
func (m *StreamingGrpcSink) XXX_Size() int {
	return xxx_messageInfo_StreamingGrpcSink.Size(m)
}
func (m *StreamingGrpcSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingGrpcSink.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingGrpcSink proto.InternalMessageInfo

func (m *StreamingGrpcSink) GetTapId() string {
	if m != nil {
		return m.TapId
	}
	return ""
}

func (m *StreamingGrpcSink) GetGrpcService() *core.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.service.tap.v2alpha.OutputSink_Format", OutputSink_Format_name, OutputSink_Format_value)
	proto.RegisterType((*TapConfig)(nil), "envoy.service.tap.v2alpha.TapConfig")
	proto.RegisterType((*MatchPredicate)(nil), "envoy.service.tap.v2alpha.MatchPredicate")
	proto.RegisterType((*MatchPredicate_MatchSet)(nil), "envoy.service.tap.v2alpha.MatchPredicate.MatchSet")
	proto.RegisterType((*HttpHeadersMatch)(nil), "envoy.service.tap.v2alpha.HttpHeadersMatch")
	proto.RegisterType((*OutputConfig)(nil), "envoy.service.tap.v2alpha.OutputConfig")
	proto.RegisterType((*OutputSink)(nil), "envoy.service.tap.v2alpha.OutputSink")
	proto.RegisterType((*StreamingAdminSink)(nil), "envoy.service.tap.v2alpha.StreamingAdminSink")
	proto.RegisterType((*FilePerTapSink)(nil), "envoy.service.tap.v2alpha.FilePerTapSink")
	proto.RegisterType((*StreamingGrpcSink)(nil), "envoy.service.tap.v2alpha.StreamingGrpcSink")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v2alpha/common.proto", fileDescriptor_4f6a60461d532a14)
}

var fileDescriptor_4f6a60461d532a14 = []byte{
	// 1061 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0xb3, 0x76, 0xec, 0x38, 0xcf, 0x6e, 0x30, 0x53, 0x42, 0xd2, 0x90, 0x56, 0xa9, 0x05,
	0x6d, 0x2a, 0x60, 0x57, 0x72, 0x6e, 0x70, 0xca, 0x36, 0x4e, 0x6c, 0x48, 0x62, 0xb3, 0x76, 0xaa,
	0x46, 0x1c, 0x46, 0x63, 0x7b, 0xec, 0x6c, 0x6b, 0xcf, 0x0c, 0xb3, 0x63, 0xd7, 0xbe, 0x20, 0xc4,
	0x0d, 0x2e, 0x48, 0xbd, 0xc0, 0x67, 0xe1, 0xc8, 0x89, 0x2b, 0x5f, 0x05, 0x09, 0x09, 0xe5, 0x80,
	0xd0, 0xce, 0x8c, 0x13, 0x3b, 0x69, 0xd3, 0xa4, 0xbd, 0x58, 0xde, 0x37, 0xef, 0xff, 0xfb, 0xef,
	0xbc, 0xd9, 0x79, 0x33, 0xf0, 0x80, 0xb2, 0x21, 0x1f, 0x7b, 0x11, 0x95, 0xc3, 0xb0, 0x45, 0x3d,
	0x45, 0x84, 0x37, 0x2c, 0x92, 0x9e, 0x38, 0x21, 0x5e, 0x8b, 0xf7, 0xfb, 0x9c, 0xb9, 0x42, 0x72,
	0xc5, 0xd1, 0x1d, 0x9d, 0xe7, 0xda, 0x3c, 0x57, 0x11, 0xe1, 0xda, 0xbc, 0xb5, 0x75, 0x83, 0x20,
	0x22, 0xf4, 0x86, 0x45, 0xaf, 0xc5, 0x25, 0xf5, 0x9a, 0x24, 0xa2, 0x46, 0xb8, 0xf6, 0xf1, 0xe5,
	0xd1, 0xae, 0x14, 0x2d, 0x3c, 0x21, 0x99, 0xac, 0x47, 0x33, 0x59, 0x92, 0x0f, 0x14, 0x35, 0xbf,
	0xb8, 0xc5, 0xfb, 0x82, 0x33, 0xca, 0x54, 0x64, 0x53, 0xef, 0x75, 0x39, 0xef, 0xf6, 0xa8, 0xa7,
	0x9f, 0x9a, 0x83, 0x8e, 0xf7, 0x42, 0x12, 0x21, 0xa8, 0x3c, 0x1b, 0x1f, 0xb4, 0x05, 0xf1, 0x08,
	0x63, 0x5c, 0x11, 0x15, 0x72, 0x16, 0x79, 0xfd, 0xb0, 0x2b, 0x89, 0x9a, 0x58, 0xdd, 0xbd, 0x34,
	0x1e, 0x29, 0xa2, 0x06, 0x13, 0xf9, 0xca, 0x90, 0xf4, 0xc2, 0x36, 0x51, 0xd4, 0x9b, 0xfc, 0x31,
	0x03, 0x85, 0x9f, 0x12, 0xb0, 0xd8, 0x20, 0xe2, 0x31, 0x67, 0x9d, 0xb0, 0x8b, 0x9e, 0x40, 0xae,
	0x4f, 0x54, 0xeb, 0x04, 0xb7, 0xf4, 0xf3, 0xaa, 0xb3, 0xe1, 0x6c, 0x66, 0x8b, 0x8f, 0xdc, 0xd7,
	0x96, 0xc9, 0x3d, 0x88, 0xd3, 0x6b, 0x92, 0xb6, 0xc3, 0x16, 0x51, 0xd4, 0xcf, 0x9c, 0xfa, 0xa9,
	0x9f, 0x9d, 0x44, 0xde, 0x09, 0xb2, 0x1a, 0x74, 0xc6, 0xbd, 0xc5, 0x07, 0x4a, 0x0c, 0xd4, 0x04,
	0x9c, 0xd0, 0xe0, 0x87, 0x57, 0x80, 0xab, 0x3a, 0xdf, 0xe8, 0xa7, 0xb0, 0x39, 0x3e, 0x15, 0x47,
	0xfb, 0x90, 0x55, 0x44, 0x60, 0xca, 0x48, 0xb3, 0x47, 0xdb, 0xab, 0x49, 0x4d, 0xfd, 0xd4, 0x52,
	0x89, 0x08, 0xdd, 0x61, 0xd1, 0x8d, 0x17, 0xc7, 0x0d, 0x06, 0x4c, 0x85, 0x7d, 0xba, 0x2b, 0x49,
	0x2b, 0xae, 0x0e, 0xe9, 0xd5, 0xa8, 0x6c, 0x51, 0xa6, 0x02, 0x50, 0x44, 0x94, 0x8c, 0xbc, 0xf0,
	0x47, 0x1a, 0x96, 0x66, 0xe7, 0x83, 0xaa, 0x90, 0xe1, 0x12, 0xeb, 0xa9, 0xd8, 0x62, 0x14, 0xaf,
	0x5d, 0x0c, 0xf3, 0x58, 0xa7, 0xaa, 0x3c, 0x17, 0x2c, 0x70, 0xa9, 0x9f, 0xd0, 0x37, 0xb0, 0x48,
	0x58, 0xdb, 0x12, 0x13, 0xef, 0x40, 0xcc, 0x10, 0xd6, 0x36, 0xc8, 0x32, 0x2c, 0x32, 0xae, 0x2c,
	0x32, 0x79, 0xc3, 0x15, 0x8b, 0x49, 0x8c, 0x2b, 0x43, 0x7a, 0x10, 0xbf, 0xdc, 0xd8, 0x92, 0xe6,
	0x37, 0x9c, 0xcd, 0x8c, 0xbf, 0x70, 0xea, 0xcf, 0x3f, 0x4b, 0x64, 0x1c, 0xe3, 0x38, 0x36, 0x79,
	0xcf, 0x60, 0xed, 0x44, 0x29, 0x81, 0x25, 0xfd, 0x6e, 0x40, 0x23, 0x85, 0x4f, 0x28, 0x69, 0x53,
	0x19, 0x59, 0x61, 0x6a, 0x66, 0x15, 0x5e, 0xf5, 0x0a, 0x65, 0xa5, 0x44, 0xd9, 0x68, 0x34, 0xb0,
	0x3c, 0x17, 0xac, 0xc4, 0xc0, 0xc0, 0xf0, 0xa6, 0x87, 0x50, 0x0f, 0x3e, 0x9a, 0xf1, 0x52, 0x92,
	0x84, 0xbd, 0x73, 0xb3, 0xf4, 0xdb, 0x98, 0xad, 0x4e, 0x99, 0x35, 0x2c, 0xef, 0xa2, 0x5b, 0x24,
	0x38, 0x8b, 0xe8, 0x85, 0xa9, 0x2d, 0xbc, 0x83, 0x9b, 0x01, 0xce, 0xcc, 0x8d, 0xc1, 0xfa, 0xac,
	0xdb, 0x85, 0xc9, 0x65, 0xde, 0xc6, 0xee, 0xce, 0xb4, 0xdd, 0xcc, 0xec, 0xd6, 0x8e, 0x20, 0x33,
	0xf9, 0x82, 0x50, 0x05, 0x52, 0x72, 0xd0, 0xa3, 0xd1, 0xaa, 0xb3, 0x91, 0xbc, 0xf9, 0x1e, 0x7f,
	0xe9, 0x24, 0x32, 0x89, 0xc0, 0x10, 0xfc, 0x2c, 0xcc, 0xc7, 0x7f, 0x50, 0xf2, 0x5f, 0xdf, 0x29,
	0x54, 0x21, 0x7f, 0xf1, 0xa5, 0xd0, 0x97, 0xb0, 0x60, 0xeb, 0x68, 0xdd, 0xee, 0xcf, 0x6e, 0x51,
	0xdd, 0x13, 0x5d, 0x23, 0xd1, 0x0a, 0x2a, 0x83, 0x89, 0xa2, 0xf0, 0x6b, 0x02, 0x72, 0xd3, 0xcd,
	0x00, 0xed, 0x41, 0x2a, 0x0a, 0xd9, 0xf3, 0x09, 0xeb, 0x93, 0x37, 0x36, 0x91, 0x7a, 0xc8, 0x9e,
	0xfb, 0x70, 0xea, 0x2f, 0xbc, 0x74, 0xe6, 0x33, 0x4e, 0xde, 0x09, 0x8c, 0x1e, 0x55, 0x61, 0xb9,
	0x4f, 0x46, 0xb8, 0x39, 0xe8, 0x74, 0xa8, 0xa4, 0x6d, 0x2c, 0x47, 0xb8, 0x39, 0x56, 0x34, 0xb2,
	0xfb, 0x72, 0xdd, 0x35, 0x3d, 0xd9, 0x9d, 0xf4, 0x64, 0xf7, 0xa8, 0xc2, 0xd4, 0x56, 0xf1, 0x09,
	0xe9, 0x0d, 0x68, 0x80, 0xfa, 0x64, 0xe4, 0x5b, 0x65, 0x30, 0xf2, 0x63, 0xdd, 0x25, 0xa0, 0x9a,
	0x00, 0x93, 0x37, 0x04, 0x36, 0x2c, 0x70, 0x1d, 0x16, 0x23, 0x25, 0x29, 0xe9, 0x87, 0xac, 0x6b,
	0x36, 0x64, 0x70, 0x1e, 0x28, 0xfc, 0x93, 0x04, 0x38, 0x9f, 0x21, 0x3a, 0x84, 0x74, 0x87, 0xcb,
	0x3e, 0x51, 0xba, 0x53, 0x2d, 0x15, 0x3f, 0xbb, 0x56, 0x61, 0xdc, 0x5d, 0xad, 0xd1, 0xab, 0xfa,
	0xa3, 0x6e, 0xb1, 0x96, 0x82, 0x9e, 0xc2, 0x7b, 0x67, 0x5e, 0x98, 0xb4, 0xfb, 0x21, 0xb3, 0x85,
	0xf9, 0xfc, 0x0a, 0x70, 0x7d, 0xa2, 0xd8, 0x8e, 0x05, 0xb1, 0x41, 0x79, 0x2e, 0x58, 0x8a, 0x66,
	0xa2, 0xe8, 0x00, 0x72, 0x9d, 0xb0, 0x47, 0xb1, 0xa0, 0x12, 0x2b, 0x22, 0xae, 0xd1, 0xb4, 0x76,
	0xc3, 0x1e, 0xad, 0x51, 0xd9, 0x20, 0xc2, 0x22, 0xa1, 0x73, 0x16, 0x41, 0x47, 0x70, 0x6e, 0x80,
	0xe3, 0x63, 0x58, 0x97, 0x2a, 0x7b, 0x65, 0x01, 0xce, 0xde, 0x73, 0x4f, 0x8a, 0x96, 0x65, 0xde,
	0x8a, 0xa6, 0x83, 0x85, 0xef, 0x21, 0x6d, 0x6a, 0x83, 0x3e, 0x04, 0xf4, 0x55, 0xbd, 0x7a, 0x88,
	0xfd, 0xea, 0xce, 0x31, 0xde, 0xae, 0x63, 0xff, 0xb8, 0x51, 0xaa, 0xe7, 0xe7, 0xd0, 0x0a, 0xdc,
	0x9e, 0x89, 0xd7, 0x1b, 0x41, 0xe5, 0x70, 0x2f, 0xef, 0xa0, 0x3c, 0xe4, 0x6a, 0x41, 0xb5, 0x51,
	0xc5, 0x7e, 0xe5, 0x70, 0x3b, 0x38, 0xce, 0x27, 0xd0, 0x7d, 0xb8, 0x3b, 0x1d, 0xc1, 0xfb, 0xa5,
	0xc3, 0xbd, 0x46, 0x19, 0xef, 0x94, 0xf6, 0x2b, 0x07, 0x95, 0x46, 0x69, 0x27, 0x9f, 0x44, 0x4b,
	0x00, 0x26, 0xa5, 0x51, 0x7a, 0xda, 0xc8, 0xcf, 0xfb, 0x2b, 0x90, 0xb7, 0x87, 0x66, 0xfc, 0xb9,
	0x62, 0x35, 0x16, 0x76, 0x8b, 0x7d, 0x00, 0xe8, 0x72, 0x99, 0x0b, 0x5f, 0xc0, 0xd2, 0x6c, 0x95,
	0xd0, 0x26, 0x64, 0x05, 0x51, 0x27, 0x58, 0x48, 0xda, 0x09, 0x47, 0xfa, 0xab, 0x58, 0xd4, 0x0d,
	0x5d, 0x26, 0x36, 0x9c, 0x00, 0xe2, 0xb1, 0x9a, 0x1e, 0x2a, 0xbc, 0x80, 0xf7, 0x2f, 0x15, 0x04,
	0x2d, 0x43, 0x3a, 0x3e, 0x5c, 0xc3, 0xb6, 0x51, 0x06, 0x29, 0x45, 0x44, 0xa5, 0x8d, 0xbe, 0x86,
	0xdc, 0xf4, 0x55, 0xc7, 0x7e, 0x13, 0xf7, 0x5e, 0x71, 0xe8, 0x6a, 0x92, 0xc9, 0x9a, 0xbe, 0x18,
	0x74, 0xa7, 0xc2, 0xdf, 0xfe, 0xfd, 0xdb, 0x7f, 0xbf, 0xa4, 0x96, 0xd1, 0x6d, 0xa3, 0x36, 0xb7,
	0x03, 0xb3, 0x50, 0x5b, 0xbf, 0xff, 0xf0, 0xe7, 0x5f, 0xe9, 0x44, 0xde, 0x81, 0x87, 0x21, 0x37,
	0x74, 0x21, 0xf9, 0x68, 0xfc, 0xfa, 0x45, 0xf5, 0xb3, 0x8f, 0xf5, 0xe5, 0xae, 0x16, 0x6f, 0xaf,
	0x9a, 0xd3, 0x4c, 0xeb, 0x7d, 0xb6, 0xf5, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x41, 0xd8,
	0x44, 0x0e, 0x0a, 0x00, 0x00,
}

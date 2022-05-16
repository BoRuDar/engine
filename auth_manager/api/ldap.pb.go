// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ldap.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// https://datatracker.ietf.org/doc/html/rfc4511#section-4.1.11
type LDAPControl struct {
	ControlType          string   `protobuf:"bytes,1,opt,name=controlType,proto3" json:"controlType,omitempty"`
	Criticality          bool     `protobuf:"varint,2,opt,name=criticality,proto3" json:"criticality,omitempty"`
	ControlValue         string   `protobuf:"bytes,3,opt,name=controlValue,proto3" json:"controlValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LDAPControl) Reset()         { *m = LDAPControl{} }
func (m *LDAPControl) String() string { return proto.CompactTextString(m) }
func (*LDAPControl) ProtoMessage()    {}
func (*LDAPControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1ff43c125b40e6, []int{0}
}

func (m *LDAPControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LDAPControl.Unmarshal(m, b)
}
func (m *LDAPControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LDAPControl.Marshal(b, m, deterministic)
}
func (m *LDAPControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LDAPControl.Merge(m, src)
}
func (m *LDAPControl) XXX_Size() int {
	return xxx_messageInfo_LDAPControl.Size(m)
}
func (m *LDAPControl) XXX_DiscardUnknown() {
	xxx_messageInfo_LDAPControl.DiscardUnknown(m)
}

var xxx_messageInfo_LDAPControl proto.InternalMessageInfo

func (m *LDAPControl) GetControlType() string {
	if m != nil {
		return m.ControlType
	}
	return ""
}

func (m *LDAPControl) GetCriticality() bool {
	if m != nil {
		return m.Criticality
	}
	return false
}

func (m *LDAPControl) GetControlValue() string {
	if m != nil {
		return m.ControlValue
	}
	return ""
}

// https://datatracker.ietf.org/doc/html/rfc4511#section-4.5.1
type LDAPSearchRequest struct {
	// ----- connection -----
	// Optional. ID of the preconfigured LDAP catalog
	CatalogId int64 `protobuf:"varint,1,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	// Optional. URL to establish connection to LDAP catalog
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	// // TLS configuration options
	// message TLSConfig {
	//     // TODO: (!)
	//     bytes cert = 1; // PEM: base64
	//     bytes key = 2; // PEM: base64
	//     bytes ca = 3; // PEM: base64
	// }
	// TLSConfig tls = 6;
	// ----- BIND: Authorization -----
	Bind     string `protobuf:"bytes,7,opt,name=bind,proto3" json:"bind,omitempty"`
	Username string `protobuf:"bytes,8,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	// ----- SearchRequest -----
	// baseObject [D]istinguished[N]ame
	BaseObject string `protobuf:"bytes,10,opt,name=baseObject,proto3" json:"baseObject,omitempty"`
	// baseObject              (0),
	// singleLevel             (1),
	// wholeSubtree            (2)
	Scope int32 `protobuf:"varint,11,opt,name=scope,proto3" json:"scope,omitempty"`
	// neverDerefAliases       (0),
	// derefInSearching        (1),
	// derefFindingBaseObj     (2),
	// derefAlways             (3)
	DerefAliases int32    `protobuf:"varint,12,opt,name=derefAliases,proto3" json:"derefAliases,omitempty"`
	SizeLimit    int64    `protobuf:"varint,13,opt,name=sizeLimit,proto3" json:"sizeLimit,omitempty"`
	TimeLimit    int64    `protobuf:"varint,14,opt,name=timeLimit,proto3" json:"timeLimit,omitempty"`
	TypesOnly    bool     `protobuf:"varint,15,opt,name=typesOnly,proto3" json:"typesOnly,omitempty"`
	Filter       string   `protobuf:"bytes,16,opt,name=filter,proto3" json:"filter,omitempty"`
	Attributes   []string `protobuf:"bytes,17,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// ----- LDAPMessage -----
	// Controls NOT implemented yet !
	Controls             []*LDAPControl `protobuf:"bytes,18,rep,name=controls,proto3" json:"controls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LDAPSearchRequest) Reset()         { *m = LDAPSearchRequest{} }
func (m *LDAPSearchRequest) String() string { return proto.CompactTextString(m) }
func (*LDAPSearchRequest) ProtoMessage()    {}
func (*LDAPSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1ff43c125b40e6, []int{1}
}

func (m *LDAPSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LDAPSearchRequest.Unmarshal(m, b)
}
func (m *LDAPSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LDAPSearchRequest.Marshal(b, m, deterministic)
}
func (m *LDAPSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LDAPSearchRequest.Merge(m, src)
}
func (m *LDAPSearchRequest) XXX_Size() int {
	return xxx_messageInfo_LDAPSearchRequest.Size(m)
}
func (m *LDAPSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LDAPSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LDAPSearchRequest proto.InternalMessageInfo

func (m *LDAPSearchRequest) GetCatalogId() int64 {
	if m != nil {
		return m.CatalogId
	}
	return 0
}

func (m *LDAPSearchRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *LDAPSearchRequest) GetBind() string {
	if m != nil {
		return m.Bind
	}
	return ""
}

func (m *LDAPSearchRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LDAPSearchRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LDAPSearchRequest) GetBaseObject() string {
	if m != nil {
		return m.BaseObject
	}
	return ""
}

func (m *LDAPSearchRequest) GetScope() int32 {
	if m != nil {
		return m.Scope
	}
	return 0
}

func (m *LDAPSearchRequest) GetDerefAliases() int32 {
	if m != nil {
		return m.DerefAliases
	}
	return 0
}

func (m *LDAPSearchRequest) GetSizeLimit() int64 {
	if m != nil {
		return m.SizeLimit
	}
	return 0
}

func (m *LDAPSearchRequest) GetTimeLimit() int64 {
	if m != nil {
		return m.TimeLimit
	}
	return 0
}

func (m *LDAPSearchRequest) GetTypesOnly() bool {
	if m != nil {
		return m.TypesOnly
	}
	return false
}

func (m *LDAPSearchRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *LDAPSearchRequest) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *LDAPSearchRequest) GetControls() []*LDAPControl {
	if m != nil {
		return m.Controls
	}
	return nil
}

// https://datatracker.ietf.org/doc/html/rfc4511#section-4.5.2
type LDAPSearchResponse struct {
	// ----- SearchResult (Entry|Reference) -----
	// repeated LDAPSearchEntry entries = 1;
	Entries []*_struct.Struct `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	// ----- LDAPResult -----
	ResultCode        int32    `protobuf:"varint,2,opt,name=resultCode,proto3" json:"resultCode,omitempty"`
	MatchedDN         string   `protobuf:"bytes,3,opt,name=matchedDN,proto3" json:"matchedDN,omitempty"`
	DiagnosticMessage string   `protobuf:"bytes,4,opt,name=diagnosticMessage,proto3" json:"diagnosticMessage,omitempty"`
	Referral          []string `protobuf:"bytes,5,rep,name=referral,proto3" json:"referral,omitempty"`
	// ----- LDAPMessage -----
	Controls             []*LDAPControl `protobuf:"bytes,6,rep,name=controls,proto3" json:"controls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LDAPSearchResponse) Reset()         { *m = LDAPSearchResponse{} }
func (m *LDAPSearchResponse) String() string { return proto.CompactTextString(m) }
func (*LDAPSearchResponse) ProtoMessage()    {}
func (*LDAPSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc1ff43c125b40e6, []int{2}
}

func (m *LDAPSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LDAPSearchResponse.Unmarshal(m, b)
}
func (m *LDAPSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LDAPSearchResponse.Marshal(b, m, deterministic)
}
func (m *LDAPSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LDAPSearchResponse.Merge(m, src)
}
func (m *LDAPSearchResponse) XXX_Size() int {
	return xxx_messageInfo_LDAPSearchResponse.Size(m)
}
func (m *LDAPSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LDAPSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LDAPSearchResponse proto.InternalMessageInfo

func (m *LDAPSearchResponse) GetEntries() []*_struct.Struct {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *LDAPSearchResponse) GetResultCode() int32 {
	if m != nil {
		return m.ResultCode
	}
	return 0
}

func (m *LDAPSearchResponse) GetMatchedDN() string {
	if m != nil {
		return m.MatchedDN
	}
	return ""
}

func (m *LDAPSearchResponse) GetDiagnosticMessage() string {
	if m != nil {
		return m.DiagnosticMessage
	}
	return ""
}

func (m *LDAPSearchResponse) GetReferral() []string {
	if m != nil {
		return m.Referral
	}
	return nil
}

func (m *LDAPSearchResponse) GetControls() []*LDAPControl {
	if m != nil {
		return m.Controls
	}
	return nil
}

func init() {
	proto.RegisterType((*LDAPControl)(nil), "api.LDAPControl")
	proto.RegisterType((*LDAPSearchRequest)(nil), "api.LDAPSearchRequest")
	proto.RegisterType((*LDAPSearchResponse)(nil), "api.LDAPSearchResponse")
}

func init() { proto.RegisterFile("ldap.proto", fileDescriptor_dc1ff43c125b40e6) }

var fileDescriptor_dc1ff43c125b40e6 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x09, 0x69, 0xbb, 0xf6, 0x74, 0x40, 0x6b, 0xa1, 0xcd, 0xaa, 0x06, 0x8a, 0x7a, 0xd5,
	0x8b, 0x29, 0x13, 0xe3, 0x01, 0xd0, 0xb4, 0x49, 0x08, 0x69, 0x30, 0x94, 0x21, 0x6e, 0x91, 0x93,
	0x9c, 0x66, 0x46, 0x6e, 0x1c, 0x7c, 0x1c, 0xa1, 0xf2, 0x48, 0xbc, 0x23, 0x12, 0xb2, 0x93, 0xa6,
	0xa9, 0x76, 0xb1, 0x3b, 0x9f, 0xdf, 0x77, 0xfc, 0xe7, 0x3b, 0xfe, 0x00, 0x54, 0x2e, 0xaa, 0xb8,
	0x32, 0xda, 0x6a, 0x16, 0x8a, 0x4a, 0x2e, 0xce, 0x0a, 0xad, 0x0b, 0x85, 0x17, 0x1e, 0xa5, 0xf5,
	0xfa, 0x82, 0xac, 0xa9, 0x33, 0xdb, 0xb4, 0x2c, 0x6b, 0x98, 0xde, 0xde, 0x5c, 0x7d, 0xbd, 0xd6,
	0xa5, 0x35, 0x5a, 0xb1, 0x08, 0xa6, 0x59, 0xb3, 0xfc, 0xb6, 0xad, 0x90, 0x07, 0x51, 0xb0, 0x9a,
	0x24, 0x7d, 0xe4, 0x3b, 0x8c, 0xb4, 0x32, 0x13, 0x4a, 0xda, 0x2d, 0x7f, 0x1e, 0x05, 0xab, 0x71,
	0xd2, 0x47, 0x6c, 0x09, 0xc7, 0xed, 0x86, 0xef, 0x42, 0xd5, 0xc8, 0x43, 0x7f, 0xc8, 0x01, 0x5b,
	0xfe, 0x0d, 0x61, 0xee, 0xee, 0xbd, 0x47, 0x61, 0xb2, 0x87, 0x04, 0x7f, 0xd5, 0x48, 0x96, 0xbd,
	0x01, 0xc8, 0x84, 0x15, 0x4a, 0x17, 0x3f, 0x64, 0xee, 0x2f, 0x0f, 0x93, 0x49, 0x4b, 0x3e, 0xe5,
	0x6c, 0x06, 0x61, 0x6d, 0x14, 0x1f, 0xfa, 0xf3, 0xdc, 0x92, 0x31, 0x18, 0xa4, 0xb2, 0xcc, 0xf9,
	0x91, 0x47, 0x7e, 0xcd, 0x16, 0x30, 0xae, 0x09, 0x4d, 0x29, 0x36, 0xc8, 0xc7, 0x9e, 0x77, 0xb5,
	0xd3, 0x2a, 0x41, 0xf4, 0x5b, 0x9b, 0x9c, 0x4f, 0x1a, 0x6d, 0x57, 0xb3, 0xb7, 0x00, 0xa9, 0x20,
	0xbc, 0x4b, 0x7f, 0x62, 0x66, 0x39, 0x78, 0xb5, 0x47, 0xd8, 0x6b, 0x18, 0x52, 0xa6, 0x2b, 0xe4,
	0xd3, 0x28, 0x58, 0x0d, 0x93, 0xa6, 0x70, 0x66, 0x73, 0x34, 0xb8, 0xbe, 0x52, 0x52, 0x10, 0x12,
	0x3f, 0xf6, 0xe2, 0x01, 0x63, 0x67, 0x30, 0x21, 0xf9, 0x07, 0x6f, 0xe5, 0x46, 0x5a, 0xfe, 0xa2,
	0x71, 0xd5, 0x01, 0xa7, 0x5a, 0xb9, 0x69, 0xd5, 0x97, 0x8d, 0xda, 0x01, 0xaf, 0x6e, 0x2b, 0xa4,
	0xbb, 0x52, 0x6d, 0xf9, 0x2b, 0x3f, 0xec, 0x3d, 0x60, 0x27, 0x30, 0x5a, 0x4b, 0x65, 0xd1, 0xf0,
	0x99, 0x7f, 0x6f, 0x5b, 0x39, 0x2f, 0xc2, 0x5a, 0x23, 0xd3, 0xda, 0x22, 0xf1, 0x79, 0x14, 0x3a,
	0x2f, 0x7b, 0xc2, 0xce, 0x61, 0xdc, 0x7e, 0x07, 0x71, 0x16, 0x85, 0xab, 0xe9, 0xe5, 0x2c, 0x16,
	0x95, 0x8c, 0x7b, 0x51, 0x48, 0xba, 0x8e, 0xe5, 0xbf, 0x00, 0x58, 0xff, 0xb3, 0xa8, 0xd2, 0x25,
	0x21, 0x7b, 0x07, 0x47, 0x58, 0x5a, 0x23, 0x91, 0x78, 0xe0, 0xcf, 0x38, 0x8d, 0x9b, 0xa8, 0xc5,
	0xbb, 0xa8, 0xc5, 0xf7, 0x3e, 0x6a, 0xc9, 0xae, 0xcf, 0xbd, 0xcb, 0x20, 0xd5, 0xca, 0x5e, 0xeb,
	0x1c, 0x7d, 0x76, 0x86, 0x49, 0x8f, 0x38, 0xb7, 0x1b, 0x61, 0xb3, 0x07, 0xcc, 0x6f, 0xbe, 0xb4,
	0xb9, 0xd9, 0x03, 0x76, 0x0e, 0xf3, 0x5c, 0x8a, 0xa2, 0xd4, 0x64, 0x65, 0xf6, 0x19, 0x89, 0x44,
	0x81, 0x7c, 0xe0, 0xbb, 0x1e, 0x0b, 0xee, 0xaf, 0x0d, 0xae, 0xd1, 0x18, 0xe1, 0x22, 0xe3, 0x26,
	0xd0, 0xd5, 0x07, 0xfe, 0x47, 0x4f, 0xf9, 0xbf, 0xfc, 0x08, 0x03, 0x27, 0xb0, 0x0f, 0x00, 0xfb,
	0x31, 0xb0, 0x93, 0x6e, 0xc7, 0x41, 0x88, 0x17, 0xa7, 0x8f, 0x78, 0x33, 0xaf, 0xe5, 0xb3, 0x74,
	0xe4, 0x07, 0xf3, 0xfe, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xbd, 0xcb, 0x80, 0xa4, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LDAPClient is the client API for LDAP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LDAPClient interface {
	// LDAPSearch performs LDAP Search operation ...
	LDAPSearch(ctx context.Context, in *LDAPSearchRequest, opts ...grpc.CallOption) (*LDAPSearchResponse, error)
}

type lDAPClient struct {
	cc *grpc.ClientConn
}

func NewLDAPClient(cc *grpc.ClientConn) LDAPClient {
	return &lDAPClient{cc}
}

func (c *lDAPClient) LDAPSearch(ctx context.Context, in *LDAPSearchRequest, opts ...grpc.CallOption) (*LDAPSearchResponse, error) {
	out := new(LDAPSearchResponse)
	err := c.cc.Invoke(ctx, "/api.LDAP/LDAPSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LDAPServer is the server API for LDAP service.
type LDAPServer interface {
	// LDAPSearch performs LDAP Search operation ...
	LDAPSearch(context.Context, *LDAPSearchRequest) (*LDAPSearchResponse, error)
}

// UnimplementedLDAPServer can be embedded to have forward compatible implementations.
type UnimplementedLDAPServer struct {
}

func (*UnimplementedLDAPServer) LDAPSearch(ctx context.Context, req *LDAPSearchRequest) (*LDAPSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LDAPSearch not implemented")
}

func RegisterLDAPServer(s *grpc.Server, srv LDAPServer) {
	s.RegisterService(&_LDAP_serviceDesc, srv)
}

func _LDAP_LDAPSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LDAPSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LDAPServer).LDAPSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LDAP/LDAPSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LDAPServer).LDAPSearch(ctx, req.(*LDAPSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LDAP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.LDAP",
	HandlerType: (*LDAPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LDAPSearch",
			Handler:    _LDAP_LDAPSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ldap.proto",
}

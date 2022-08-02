// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: ldap.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LDAPClient is the client API for LDAP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LDAPClient interface {
	// LDAPSearch performs LDAP Search operation ...
	LDAPSearch(ctx context.Context, in *LDAPSearchRequest, opts ...grpc.CallOption) (*LDAPSearchResponse, error)
}

type lDAPClient struct {
	cc grpc.ClientConnInterface
}

func NewLDAPClient(cc grpc.ClientConnInterface) LDAPClient {
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
// All implementations must embed UnimplementedLDAPServer
// for forward compatibility
type LDAPServer interface {
	// LDAPSearch performs LDAP Search operation ...
	LDAPSearch(context.Context, *LDAPSearchRequest) (*LDAPSearchResponse, error)
	mustEmbedUnimplementedLDAPServer()
}

// UnimplementedLDAPServer must be embedded to have forward compatible implementations.
type UnimplementedLDAPServer struct {
}

func (UnimplementedLDAPServer) LDAPSearch(context.Context, *LDAPSearchRequest) (*LDAPSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LDAPSearch not implemented")
}
func (UnimplementedLDAPServer) mustEmbedUnimplementedLDAPServer() {}

// UnsafeLDAPServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LDAPServer will
// result in compilation errors.
type UnsafeLDAPServer interface {
	mustEmbedUnimplementedLDAPServer()
}

func RegisterLDAPServer(s grpc.ServiceRegistrar, srv LDAPServer) {
	s.RegisterService(&LDAP_ServiceDesc, srv)
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

// LDAP_ServiceDesc is the grpc.ServiceDesc for LDAP service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LDAP_ServiceDesc = grpc.ServiceDesc{
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

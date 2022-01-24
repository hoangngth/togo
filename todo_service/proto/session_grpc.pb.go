// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionServiceClient interface {
	GetAccountIDFromToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*AccountID, error)
	GetAccountTypeFromToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*AccountType, error)
	CreateToken(ctx context.Context, in *AccountInfo, opts ...grpc.CallOption) (*TokenString, error)
	RefreshToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*TokenString, error)
	DeleteToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*Status, error)
	CheckToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*Status, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) GetAccountIDFromToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*AccountID, error) {
	out := new(AccountID)
	err := c.cc.Invoke(ctx, "/proto.SessionService/GetAccountIDFromToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetAccountTypeFromToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*AccountType, error) {
	out := new(AccountType)
	err := c.cc.Invoke(ctx, "/proto.SessionService/GetAccountTypeFromToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CreateToken(ctx context.Context, in *AccountInfo, opts ...grpc.CallOption) (*TokenString, error) {
	out := new(TokenString)
	err := c.cc.Invoke(ctx, "/proto.SessionService/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) RefreshToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*TokenString, error) {
	out := new(TokenString)
	err := c.cc.Invoke(ctx, "/proto.SessionService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.SessionService/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) CheckToken(ctx context.Context, in *TokenString, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.SessionService/CheckToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
// All implementations must embed UnimplementedSessionServiceServer
// for forward compatibility
type SessionServiceServer interface {
	GetAccountIDFromToken(context.Context, *TokenString) (*AccountID, error)
	GetAccountTypeFromToken(context.Context, *TokenString) (*AccountType, error)
	CreateToken(context.Context, *AccountInfo) (*TokenString, error)
	RefreshToken(context.Context, *TokenString) (*TokenString, error)
	DeleteToken(context.Context, *TokenString) (*Status, error)
	CheckToken(context.Context, *TokenString) (*Status, error)
	mustEmbedUnimplementedSessionServiceServer()
}

// UnimplementedSessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (UnimplementedSessionServiceServer) GetAccountIDFromToken(context.Context, *TokenString) (*AccountID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountIDFromToken not implemented")
}
func (UnimplementedSessionServiceServer) GetAccountTypeFromToken(context.Context, *TokenString) (*AccountType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountTypeFromToken not implemented")
}
func (UnimplementedSessionServiceServer) CreateToken(context.Context, *AccountInfo) (*TokenString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedSessionServiceServer) RefreshToken(context.Context, *TokenString) (*TokenString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedSessionServiceServer) DeleteToken(context.Context, *TokenString) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (UnimplementedSessionServiceServer) CheckToken(context.Context, *TokenString) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {}

// UnsafeSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServiceServer will
// result in compilation errors.
type UnsafeSessionServiceServer interface {
	mustEmbedUnimplementedSessionServiceServer()
}

func RegisterSessionServiceServer(s grpc.ServiceRegistrar, srv SessionServiceServer) {
	s.RegisterService(&SessionService_ServiceDesc, srv)
}

func _SessionService_GetAccountIDFromToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetAccountIDFromToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/GetAccountIDFromToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetAccountIDFromToken(ctx, req.(*TokenString))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetAccountTypeFromToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetAccountTypeFromToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/GetAccountTypeFromToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetAccountTypeFromToken(ctx, req.(*TokenString))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CreateToken(ctx, req.(*AccountInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).RefreshToken(ctx, req.(*TokenString))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteToken(ctx, req.(*TokenString))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionService/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).CheckToken(ctx, req.(*TokenString))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionService_ServiceDesc is the grpc.ServiceDesc for SessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountIDFromToken",
			Handler:    _SessionService_GetAccountIDFromToken_Handler,
		},
		{
			MethodName: "GetAccountTypeFromToken",
			Handler:    _SessionService_GetAccountTypeFromToken_Handler,
		},
		{
			MethodName: "CreateToken",
			Handler:    _SessionService_CreateToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _SessionService_RefreshToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _SessionService_DeleteToken_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _SessionService_CheckToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
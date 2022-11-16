// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/model.proto

package __

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

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDClient interface {
	Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Response, error)
	GetUserByUserType(ctx context.Context, in *GetUserByUserTypeRequest, opts ...grpc.CallOption) (*GetUserByUserTypeResponse, error)
	GetBalanceByID(ctx context.Context, in *GetBalanceByIDRequest, opts ...grpc.CallOption) (*GetBalanceByIDResponse, error)
	UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*Response, error)
	Authentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	RefreshMyTokens(ctx context.Context, in *RefreshTokensRequest, opts ...grpc.CallOption) (*RefreshTokensResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Response, error)
}

type cRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDClient(cc grpc.ClientConnInterface) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) Registration(ctx context.Context, in *RegistrationRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/CRUD/Registration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/CRUD/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, "/CRUD/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/CRUD/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/CRUD/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) GetUserByUserType(ctx context.Context, in *GetUserByUserTypeRequest, opts ...grpc.CallOption) (*GetUserByUserTypeResponse, error) {
	out := new(GetUserByUserTypeResponse)
	err := c.cc.Invoke(ctx, "/CRUD/GetUserByUserType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) GetBalanceByID(ctx context.Context, in *GetBalanceByIDRequest, opts ...grpc.CallOption) (*GetBalanceByIDResponse, error) {
	out := new(GetBalanceByIDResponse)
	err := c.cc.Invoke(ctx, "/CRUD/GetBalanceByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) UpdateBalance(ctx context.Context, in *UpdateBalanceRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/CRUD/UpdateBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Authentication(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, "/CRUD/Authentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) RefreshMyTokens(ctx context.Context, in *RefreshTokensRequest, opts ...grpc.CallOption) (*RefreshTokensResponse, error) {
	out := new(RefreshTokensResponse)
	err := c.cc.Invoke(ctx, "/CRUD/RefreshMyTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/CRUD/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
// All implementations must embed UnimplementedCRUDServer
// for forward compatibility
type CRUDServer interface {
	Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*Response, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*Response, error)
	GetUserByUserType(context.Context, *GetUserByUserTypeRequest) (*GetUserByUserTypeResponse, error)
	GetBalanceByID(context.Context, *GetBalanceByIDRequest) (*GetBalanceByIDResponse, error)
	UpdateBalance(context.Context, *UpdateBalanceRequest) (*Response, error)
	Authentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	RefreshMyTokens(context.Context, *RefreshTokensRequest) (*RefreshTokensResponse, error)
	Logout(context.Context, *LogoutRequest) (*Response, error)
	mustEmbedUnimplementedCRUDServer()
}

// UnimplementedCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (UnimplementedCRUDServer) Registration(context.Context, *RegistrationRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registration not implemented")
}
func (UnimplementedCRUDServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedCRUDServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedCRUDServer) DeleteUser(context.Context, *DeleteUserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedCRUDServer) UpdateUser(context.Context, *UpdateUserRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedCRUDServer) GetUserByUserType(context.Context, *GetUserByUserTypeRequest) (*GetUserByUserTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUserType not implemented")
}
func (UnimplementedCRUDServer) GetBalanceByID(context.Context, *GetBalanceByIDRequest) (*GetBalanceByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalanceByID not implemented")
}
func (UnimplementedCRUDServer) UpdateBalance(context.Context, *UpdateBalanceRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBalance not implemented")
}
func (UnimplementedCRUDServer) Authentication(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authentication not implemented")
}
func (UnimplementedCRUDServer) RefreshMyTokens(context.Context, *RefreshTokensRequest) (*RefreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshMyTokens not implemented")
}
func (UnimplementedCRUDServer) Logout(context.Context, *LogoutRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedCRUDServer) mustEmbedUnimplementedCRUDServer() {}

// UnsafeCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServer will
// result in compilation errors.
type UnsafeCRUDServer interface {
	mustEmbedUnimplementedCRUDServer()
}

func RegisterCRUDServer(s grpc.ServiceRegistrar, srv CRUDServer) {
	s.RegisterService(&CRUD_ServiceDesc, srv)
}

func _CRUD_Registration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Registration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/Registration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Registration(ctx, req.(*RegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_GetUserByUserType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUserTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).GetUserByUserType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/GetUserByUserType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).GetUserByUserType(ctx, req.(*GetUserByUserTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_GetBalanceByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).GetBalanceByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/GetBalanceByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).GetBalanceByID(ctx, req.(*GetBalanceByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_UpdateBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).UpdateBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/UpdateBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).UpdateBalance(ctx, req.(*UpdateBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Authentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Authentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/Authentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Authentication(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_RefreshMyTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).RefreshMyTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/RefreshMyTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).RefreshMyTokens(ctx, req.(*RefreshTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CRUD/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CRUD_ServiceDesc is the grpc.ServiceDesc for CRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registration",
			Handler:    _CRUD_Registration_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _CRUD_GetUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _CRUD_GetAllUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _CRUD_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _CRUD_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserByUserType",
			Handler:    _CRUD_GetUserByUserType_Handler,
		},
		{
			MethodName: "GetBalanceByID",
			Handler:    _CRUD_GetBalanceByID_Handler,
		},
		{
			MethodName: "UpdateBalance",
			Handler:    _CRUD_UpdateBalance_Handler,
		},
		{
			MethodName: "Authentication",
			Handler:    _CRUD_Authentication_Handler,
		},
		{
			MethodName: "RefreshMyTokens",
			Handler:    _CRUD_RefreshMyTokens_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _CRUD_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/model.proto",
}

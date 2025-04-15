// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_GetUserById_FullMethodName      = "/pb.UserService/GetUserById"
	UserService_GetUsersByIds_FullMethodName    = "/pb.UserService/GetUsersByIds"
	UserService_CreateUser_FullMethodName       = "/pb.UserService/CreateUser"
	UserService_UpdateUserStatus_FullMethodName = "/pb.UserService/UpdateUserStatus"
	UserService_GetUserStatus_FullMethodName    = "/pb.UserService/GetUserStatus"
	UserService_DeleteUser_FullMethodName       = "/pb.UserService/DeleteUser"
	UserService_GetUserIdByEmail_FullMethodName = "/pb.UserService/GetUserIdByEmail"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*PublicUserInfoResp, error)
	GetUsersByIds(ctx context.Context, in *GetUsersByIdsReq, opts ...grpc.CallOption) (*PublicUsersInfoResp, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*NewUserIdResp, error)
	UpdateUserStatus(ctx context.Context, in *UpdateUserStatusReq, opts ...grpc.CallOption) (*UpdateUserStatusResp, error)
	GetUserStatus(ctx context.Context, in *GetUserStatusReq, opts ...grpc.CallOption) (*GetUserStatusResp, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error)
	GetUserIdByEmail(ctx context.Context, in *GetUserIdByEmailReq, opts ...grpc.CallOption) (*GetUserIdByEmailResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*PublicUserInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicUserInfoResp)
	err := c.cc.Invoke(ctx, UserService_GetUserById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsersByIds(ctx context.Context, in *GetUsersByIdsReq, opts ...grpc.CallOption) (*PublicUsersInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicUsersInfoResp)
	err := c.cc.Invoke(ctx, UserService_GetUsersByIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*NewUserIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NewUserIdResp)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusReq, opts ...grpc.CallOption) (*UpdateUserStatusResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserStatusResp)
	err := c.cc.Invoke(ctx, UserService_UpdateUserStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserStatus(ctx context.Context, in *GetUserStatusReq, opts ...grpc.CallOption) (*GetUserStatusResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserStatusResp)
	err := c.cc.Invoke(ctx, UserService_GetUserStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...grpc.CallOption) (*DeleteUserResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResp)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserIdByEmail(ctx context.Context, in *GetUserIdByEmailReq, opts ...grpc.CallOption) (*GetUserIdByEmailResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserIdByEmailResp)
	err := c.cc.Invoke(ctx, UserService_GetUserIdByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	GetUserById(context.Context, *GetUserByIdReq) (*PublicUserInfoResp, error)
	GetUsersByIds(context.Context, *GetUsersByIdsReq) (*PublicUsersInfoResp, error)
	CreateUser(context.Context, *CreateUserReq) (*NewUserIdResp, error)
	UpdateUserStatus(context.Context, *UpdateUserStatusReq) (*UpdateUserStatusResp, error)
	GetUserStatus(context.Context, *GetUserStatusReq) (*GetUserStatusResp, error)
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserResp, error)
	GetUserIdByEmail(context.Context, *GetUserIdByEmailReq) (*GetUserIdByEmailResp, error)
}

// UnimplementedUserServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetUserById(context.Context, *GetUserByIdReq) (*PublicUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceServer) GetUsersByIds(context.Context, *GetUsersByIdsReq) (*PublicUsersInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByIds not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserReq) (*NewUserIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserStatus(context.Context, *UpdateUserStatusReq) (*UpdateUserStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserStatus not implemented")
}
func (UnimplementedUserServiceServer) GetUserStatus(context.Context, *GetUserStatusReq) (*GetUserStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStatus not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserIdByEmail(context.Context, *GetUserIdByEmailReq) (*GetUserIdByEmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserIdByEmail not implemented")
}
func (UnimplementedUserServiceServer) testEmbeddedByValue() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsersByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsersByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUsersByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsersByIds(ctx, req.(*GetUsersByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserStatus(ctx, req.(*UpdateUserStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserStatus(ctx, req.(*GetUserStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserIdByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserIdByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserIdByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserIdByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserIdByEmail(ctx, req.(*GetUserIdByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _UserService_GetUserById_Handler,
		},
		{
			MethodName: "GetUsersByIds",
			Handler:    _UserService_GetUsersByIds_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUserStatus",
			Handler:    _UserService_UpdateUserStatus_Handler,
		},
		{
			MethodName: "GetUserStatus",
			Handler:    _UserService_GetUserStatus_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "GetUserIdByEmail",
			Handler:    _UserService_GetUserIdByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}

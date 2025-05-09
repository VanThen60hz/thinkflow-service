// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/note.proto

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
	NoteService_DeleteUserNotes_FullMethodName = "/pb.NoteService/DeleteUserNotes"
	NoteService_GetNoteById_FullMethodName     = "/pb.NoteService/GetNoteById"
	NoteService_CountNotes_FullMethodName      = "/pb.NoteService/CountNotes"
)

// NoteServiceClient is the client API for NoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteServiceClient interface {
	DeleteUserNotes(ctx context.Context, in *DeleteUserNotesReq, opts ...grpc.CallOption) (*DeleteUserNotesResp, error)
	GetNoteById(ctx context.Context, in *GetNoteByIdReq, opts ...grpc.CallOption) (*GetNoteByIdResp, error)
	CountNotes(ctx context.Context, in *CountNotesReq, opts ...grpc.CallOption) (*CountNotesResp, error)
}

type noteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoteServiceClient(cc grpc.ClientConnInterface) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) DeleteUserNotes(ctx context.Context, in *DeleteUserNotesReq, opts ...grpc.CallOption) (*DeleteUserNotesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserNotesResp)
	err := c.cc.Invoke(ctx, NoteService_DeleteUserNotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) GetNoteById(ctx context.Context, in *GetNoteByIdReq, opts ...grpc.CallOption) (*GetNoteByIdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNoteByIdResp)
	err := c.cc.Invoke(ctx, NoteService_GetNoteById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) CountNotes(ctx context.Context, in *CountNotesReq, opts ...grpc.CallOption) (*CountNotesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountNotesResp)
	err := c.cc.Invoke(ctx, NoteService_CountNotes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteServiceServer is the server API for NoteService service.
// All implementations should embed UnimplementedNoteServiceServer
// for forward compatibility.
type NoteServiceServer interface {
	DeleteUserNotes(context.Context, *DeleteUserNotesReq) (*DeleteUserNotesResp, error)
	GetNoteById(context.Context, *GetNoteByIdReq) (*GetNoteByIdResp, error)
	CountNotes(context.Context, *CountNotesReq) (*CountNotesResp, error)
}

// UnimplementedNoteServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNoteServiceServer struct{}

func (UnimplementedNoteServiceServer) DeleteUserNotes(context.Context, *DeleteUserNotesReq) (*DeleteUserNotesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserNotes not implemented")
}
func (UnimplementedNoteServiceServer) GetNoteById(context.Context, *GetNoteByIdReq) (*GetNoteByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteById not implemented")
}
func (UnimplementedNoteServiceServer) CountNotes(context.Context, *CountNotesReq) (*CountNotesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountNotes not implemented")
}
func (UnimplementedNoteServiceServer) testEmbeddedByValue() {}

// UnsafeNoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteServiceServer will
// result in compilation errors.
type UnsafeNoteServiceServer interface {
	mustEmbedUnimplementedNoteServiceServer()
}

func RegisterNoteServiceServer(s grpc.ServiceRegistrar, srv NoteServiceServer) {
	// If the following call pancis, it indicates UnimplementedNoteServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NoteService_ServiceDesc, srv)
}

func _NoteService_DeleteUserNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserNotesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).DeleteUserNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_DeleteUserNotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).DeleteUserNotes(ctx, req.(*DeleteUserNotesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_GetNoteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).GetNoteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_GetNoteById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).GetNoteById(ctx, req.(*GetNoteByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_CountNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountNotesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).CountNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_CountNotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).CountNotes(ctx, req.(*CountNotesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteService_ServiceDesc is the grpc.ServiceDesc for NoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteUserNotes",
			Handler:    _NoteService_DeleteUserNotes_Handler,
		},
		{
			MethodName: "GetNoteById",
			Handler:    _NoteService_GetNoteById_Handler,
		},
		{
			MethodName: "CountNotes",
			Handler:    _NoteService_CountNotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/note.proto",
}

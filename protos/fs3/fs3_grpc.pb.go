// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: protos/fs3/fs3.proto

package fs3

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

// Fs3Client is the client API for Fs3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Fs3Client interface {
	Copy(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*CopyReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
}

type fs3Client struct {
	cc grpc.ClientConnInterface
}

func NewFs3Client(cc grpc.ClientConnInterface) Fs3Client {
	return &fs3Client{cc}
}

func (c *fs3Client) Copy(ctx context.Context, in *CopyRequest, opts ...grpc.CallOption) (*CopyReply, error) {
	out := new(CopyReply)
	err := c.cc.Invoke(ctx, "/fs3.fs3/Copy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fs3Client) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := c.cc.Invoke(ctx, "/fs3.fs3/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fs3Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/fs3.fs3/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Fs3Server is the server API for Fs3 service.
// All implementations must embed UnimplementedFs3Server
// for forward compatibility
type Fs3Server interface {
	Copy(context.Context, *CopyRequest) (*CopyReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	mustEmbedUnimplementedFs3Server()
}

// UnimplementedFs3Server must be embedded to have forward compatible implementations.
type UnimplementedFs3Server struct {
}

func (UnimplementedFs3Server) Copy(context.Context, *CopyRequest) (*CopyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Copy not implemented")
}
func (UnimplementedFs3Server) Remove(context.Context, *RemoveRequest) (*RemoveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedFs3Server) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFs3Server) mustEmbedUnimplementedFs3Server() {}

// UnsafeFs3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Fs3Server will
// result in compilation errors.
type UnsafeFs3Server interface {
	mustEmbedUnimplementedFs3Server()
}

func RegisterFs3Server(s grpc.ServiceRegistrar, srv Fs3Server) {
	s.RegisterService(&Fs3_ServiceDesc, srv)
}

func _Fs3_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fs3Server).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs3.fs3/Copy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fs3Server).Copy(ctx, req.(*CopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs3_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fs3Server).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs3.fs3/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fs3Server).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs3_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Fs3Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs3.fs3/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Fs3Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Fs3_ServiceDesc is the grpc.ServiceDesc for Fs3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fs3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fs3.fs3",
	HandlerType: (*Fs3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Copy",
			Handler:    _Fs3_Copy_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Fs3_Remove_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Fs3_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/fs3/fs3.proto",
}

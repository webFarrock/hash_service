// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hash_service

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

// HashesClient is the client API for Hashes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HashesClient interface {
	Generate(ctx context.Context, in *ArrayOfStrings, opts ...grpc.CallOption) (*ArrayOfHash, error)
	Check(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ArrayOfHash, error)
}

type hashesClient struct {
	cc grpc.ClientConnInterface
}

func NewHashesClient(cc grpc.ClientConnInterface) HashesClient {
	return &hashesClient{cc}
}

func (c *hashesClient) Generate(ctx context.Context, in *ArrayOfStrings, opts ...grpc.CallOption) (*ArrayOfHash, error) {
	out := new(ArrayOfHash)
	err := c.cc.Invoke(ctx, "/hash_service.Hashes/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hashesClient) Check(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ArrayOfHash, error) {
	out := new(ArrayOfHash)
	err := c.cc.Invoke(ctx, "/hash_service.Hashes/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HashesServer is the server API for Hashes service.
// All implementations must embed UnimplementedHashesServer
// for forward compatibility
type HashesServer interface {
	Generate(context.Context, *ArrayOfStrings) (*ArrayOfHash, error)
	Check(context.Context, *Filter) (*ArrayOfHash, error)
	mustEmbedUnimplementedHashesServer()
}

// UnimplementedHashesServer must be embedded to have forward compatible implementations.
type UnimplementedHashesServer struct {
}

func (UnimplementedHashesServer) Generate(context.Context, *ArrayOfStrings) (*ArrayOfHash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedHashesServer) Check(context.Context, *Filter) (*ArrayOfHash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedHashesServer) mustEmbedUnimplementedHashesServer() {}

// UnsafeHashesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HashesServer will
// result in compilation errors.
type UnsafeHashesServer interface {
	mustEmbedUnimplementedHashesServer()
}

func RegisterHashesServer(s grpc.ServiceRegistrar, srv HashesServer) {
	s.RegisterService(&Hashes_ServiceDesc, srv)
}

func _Hashes_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArrayOfStrings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashesServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash_service.Hashes/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashesServer).Generate(ctx, req.(*ArrayOfStrings))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hashes_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HashesServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash_service.Hashes/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HashesServer).Check(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// Hashes_ServiceDesc is the grpc.ServiceDesc for Hashes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hashes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hash_service.Hashes",
	HandlerType: (*HashesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Hashes_Generate_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Hashes_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

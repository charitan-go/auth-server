// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: pkg/proto/auth.proto

package proto

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
	ProfileGrpcService_CreateDonorProfile_FullMethodName   = "/ProfileGrpcService/CreateDonorProfile"
	ProfileGrpcService_CreateCharityProfile_FullMethodName = "/ProfileGrpcService/CreateCharityProfile"
	ProfileGrpcService_GetDonorProfile_FullMethodName      = "/ProfileGrpcService/GetDonorProfile"
	ProfileGrpcService_GetCharityProfile_FullMethodName    = "/ProfileGrpcService/GetCharityProfile"
)

// ProfileGrpcServiceClient is the client API for ProfileGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileGrpcServiceClient interface {
	CreateDonorProfile(ctx context.Context, in *CreateDonorProfileRequestDto, opts ...grpc.CallOption) (*CreateDonorProfileResponseDto, error)
	CreateCharityProfile(ctx context.Context, in *CreateCharityProfileRequestDto, opts ...grpc.CallOption) (*CreateCharityProfileResponseDto, error)
	GetDonorProfile(ctx context.Context, in *GetDonorProfileRequestDto, opts ...grpc.CallOption) (*GetDonorProfileResponseDto, error)
	GetCharityProfile(ctx context.Context, in *GetCharityProfileRequestDto, opts ...grpc.CallOption) (*GetCharityProfileResponseDto, error)
}

type profileGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileGrpcServiceClient(cc grpc.ClientConnInterface) ProfileGrpcServiceClient {
	return &profileGrpcServiceClient{cc}
}

func (c *profileGrpcServiceClient) CreateDonorProfile(ctx context.Context, in *CreateDonorProfileRequestDto, opts ...grpc.CallOption) (*CreateDonorProfileResponseDto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDonorProfileResponseDto)
	err := c.cc.Invoke(ctx, ProfileGrpcService_CreateDonorProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileGrpcServiceClient) CreateCharityProfile(ctx context.Context, in *CreateCharityProfileRequestDto, opts ...grpc.CallOption) (*CreateCharityProfileResponseDto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCharityProfileResponseDto)
	err := c.cc.Invoke(ctx, ProfileGrpcService_CreateCharityProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileGrpcServiceClient) GetDonorProfile(ctx context.Context, in *GetDonorProfileRequestDto, opts ...grpc.CallOption) (*GetDonorProfileResponseDto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDonorProfileResponseDto)
	err := c.cc.Invoke(ctx, ProfileGrpcService_GetDonorProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileGrpcServiceClient) GetCharityProfile(ctx context.Context, in *GetCharityProfileRequestDto, opts ...grpc.CallOption) (*GetCharityProfileResponseDto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCharityProfileResponseDto)
	err := c.cc.Invoke(ctx, ProfileGrpcService_GetCharityProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileGrpcServiceServer is the server API for ProfileGrpcService service.
// All implementations must embed UnimplementedProfileGrpcServiceServer
// for forward compatibility.
type ProfileGrpcServiceServer interface {
	CreateDonorProfile(context.Context, *CreateDonorProfileRequestDto) (*CreateDonorProfileResponseDto, error)
	CreateCharityProfile(context.Context, *CreateCharityProfileRequestDto) (*CreateCharityProfileResponseDto, error)
	GetDonorProfile(context.Context, *GetDonorProfileRequestDto) (*GetDonorProfileResponseDto, error)
	GetCharityProfile(context.Context, *GetCharityProfileRequestDto) (*GetCharityProfileResponseDto, error)
	mustEmbedUnimplementedProfileGrpcServiceServer()
}

// UnimplementedProfileGrpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfileGrpcServiceServer struct{}

func (UnimplementedProfileGrpcServiceServer) CreateDonorProfile(context.Context, *CreateDonorProfileRequestDto) (*CreateDonorProfileResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDonorProfile not implemented")
}
func (UnimplementedProfileGrpcServiceServer) CreateCharityProfile(context.Context, *CreateCharityProfileRequestDto) (*CreateCharityProfileResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCharityProfile not implemented")
}
func (UnimplementedProfileGrpcServiceServer) GetDonorProfile(context.Context, *GetDonorProfileRequestDto) (*GetDonorProfileResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonorProfile not implemented")
}
func (UnimplementedProfileGrpcServiceServer) GetCharityProfile(context.Context, *GetCharityProfileRequestDto) (*GetCharityProfileResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharityProfile not implemented")
}
func (UnimplementedProfileGrpcServiceServer) mustEmbedUnimplementedProfileGrpcServiceServer() {}
func (UnimplementedProfileGrpcServiceServer) testEmbeddedByValue()                            {}

// UnsafeProfileGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileGrpcServiceServer will
// result in compilation errors.
type UnsafeProfileGrpcServiceServer interface {
	mustEmbedUnimplementedProfileGrpcServiceServer()
}

func RegisterProfileGrpcServiceServer(s grpc.ServiceRegistrar, srv ProfileGrpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedProfileGrpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProfileGrpcService_ServiceDesc, srv)
}

func _ProfileGrpcService_CreateDonorProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDonorProfileRequestDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileGrpcServiceServer).CreateDonorProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileGrpcService_CreateDonorProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileGrpcServiceServer).CreateDonorProfile(ctx, req.(*CreateDonorProfileRequestDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileGrpcService_CreateCharityProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCharityProfileRequestDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileGrpcServiceServer).CreateCharityProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileGrpcService_CreateCharityProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileGrpcServiceServer).CreateCharityProfile(ctx, req.(*CreateCharityProfileRequestDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileGrpcService_GetDonorProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDonorProfileRequestDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileGrpcServiceServer).GetDonorProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileGrpcService_GetDonorProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileGrpcServiceServer).GetDonorProfile(ctx, req.(*GetDonorProfileRequestDto))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileGrpcService_GetCharityProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCharityProfileRequestDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileGrpcServiceServer).GetCharityProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileGrpcService_GetCharityProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileGrpcServiceServer).GetCharityProfile(ctx, req.(*GetCharityProfileRequestDto))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileGrpcService_ServiceDesc is the grpc.ServiceDesc for ProfileGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProfileGrpcService",
	HandlerType: (*ProfileGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDonorProfile",
			Handler:    _ProfileGrpcService_CreateDonorProfile_Handler,
		},
		{
			MethodName: "CreateCharityProfile",
			Handler:    _ProfileGrpcService_CreateCharityProfile_Handler,
		},
		{
			MethodName: "GetDonorProfile",
			Handler:    _ProfileGrpcService_GetDonorProfile_Handler,
		},
		{
			MethodName: "GetCharityProfile",
			Handler:    _ProfileGrpcService_GetCharityProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/auth.proto",
}

const (
	KeyGrpcService_GetPrivateKey_FullMethodName = "/KeyGrpcService/GetPrivateKey"
)

// KeyGrpcServiceClient is the client API for KeyGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyGrpcServiceClient interface {
	GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequestDto, opts ...grpc.CallOption) (*GetPrivateKeyResponseDto, error)
}

type keyGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyGrpcServiceClient(cc grpc.ClientConnInterface) KeyGrpcServiceClient {
	return &keyGrpcServiceClient{cc}
}

func (c *keyGrpcServiceClient) GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequestDto, opts ...grpc.CallOption) (*GetPrivateKeyResponseDto, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPrivateKeyResponseDto)
	err := c.cc.Invoke(ctx, KeyGrpcService_GetPrivateKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyGrpcServiceServer is the server API for KeyGrpcService service.
// All implementations must embed UnimplementedKeyGrpcServiceServer
// for forward compatibility.
type KeyGrpcServiceServer interface {
	GetPrivateKey(context.Context, *GetPrivateKeyRequestDto) (*GetPrivateKeyResponseDto, error)
	mustEmbedUnimplementedKeyGrpcServiceServer()
}

// UnimplementedKeyGrpcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKeyGrpcServiceServer struct{}

func (UnimplementedKeyGrpcServiceServer) GetPrivateKey(context.Context, *GetPrivateKeyRequestDto) (*GetPrivateKeyResponseDto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}
func (UnimplementedKeyGrpcServiceServer) mustEmbedUnimplementedKeyGrpcServiceServer() {}
func (UnimplementedKeyGrpcServiceServer) testEmbeddedByValue()                        {}

// UnsafeKeyGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyGrpcServiceServer will
// result in compilation errors.
type UnsafeKeyGrpcServiceServer interface {
	mustEmbedUnimplementedKeyGrpcServiceServer()
}

func RegisterKeyGrpcServiceServer(s grpc.ServiceRegistrar, srv KeyGrpcServiceServer) {
	// If the following call pancis, it indicates UnimplementedKeyGrpcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KeyGrpcService_ServiceDesc, srv)
}

func _KeyGrpcService_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateKeyRequestDto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyGrpcServiceServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyGrpcService_GetPrivateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyGrpcServiceServer).GetPrivateKey(ctx, req.(*GetPrivateKeyRequestDto))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyGrpcService_ServiceDesc is the grpc.ServiceDesc for KeyGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KeyGrpcService",
	HandlerType: (*KeyGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrivateKey",
			Handler:    _KeyGrpcService_GetPrivateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/auth.proto",
}

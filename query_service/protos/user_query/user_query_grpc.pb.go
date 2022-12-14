// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: user_query.proto

package queryService

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

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRes, error)
	GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdRes, error)
	SearchUser(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error)
	DeleteUserByID(ctx context.Context, in *DeleteUserByIdReq, opts ...grpc.CallOption) (*DeleteUserByIdRes, error)
}

type queryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryServiceClient(cc grpc.ClientConnInterface) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserRes, error) {
	out := new(CreateUserRes)
	err := c.cc.Invoke(ctx, "/queryService.queryService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserRes, error) {
	out := new(UpdateUserRes)
	err := c.cc.Invoke(ctx, "/queryService.queryService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdRes, error) {
	out := new(GetUserByIdRes)
	err := c.cc.Invoke(ctx, "/queryService.queryService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) SearchUser(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error) {
	out := new(SearchRes)
	err := c.cc.Invoke(ctx, "/queryService.queryService/SearchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) DeleteUserByID(ctx context.Context, in *DeleteUserByIdReq, opts ...grpc.CallOption) (*DeleteUserByIdRes, error) {
	out := new(DeleteUserByIdRes)
	err := c.cc.Invoke(ctx, "/queryService.queryService/DeleteUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
// All implementations should embed UnimplementedQueryServiceServer
// for forward compatibility
type QueryServiceServer interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserRes, error)
	GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdRes, error)
	SearchUser(context.Context, *SearchReq) (*SearchRes, error)
	DeleteUserByID(context.Context, *DeleteUserByIdReq) (*DeleteUserByIdRes, error)
}

// UnimplementedQueryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (UnimplementedQueryServiceServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedQueryServiceServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedQueryServiceServer) GetUserById(context.Context, *GetUserByIdReq) (*GetUserByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedQueryServiceServer) SearchUser(context.Context, *SearchReq) (*SearchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedQueryServiceServer) DeleteUserByID(context.Context, *DeleteUserByIdReq) (*DeleteUserByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserByID not implemented")
}

// UnsafeQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServiceServer will
// result in compilation errors.
type UnsafeQueryServiceServer interface {
	mustEmbedUnimplementedQueryServiceServer()
}

func RegisterQueryServiceServer(s grpc.ServiceRegistrar, srv QueryServiceServer) {
	s.RegisterService(&QueryService_ServiceDesc, srv)
}

func _QueryService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryService.queryService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryService.queryService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryService.queryService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).GetUserById(ctx, req.(*GetUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryService.queryService/SearchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).SearchUser(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_DeleteUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).DeleteUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/queryService.queryService/DeleteUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).DeleteUserByID(ctx, req.(*DeleteUserByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryService_ServiceDesc is the grpc.ServiceDesc for QueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "queryService.queryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _QueryService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _QueryService_UpdateUser_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _QueryService_GetUserById_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _QueryService_SearchUser_Handler,
		},
		{
			MethodName: "DeleteUserByID",
			Handler:    _QueryService_DeleteUserByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_query.proto",
}

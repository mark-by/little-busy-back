// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/pkg/proto/crm/crm.proto

package crm

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

// CrmServiceClient is the client API for CrmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrmServiceClient interface {
	GetUserByPhoneNumber(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	GetFutureEventsForUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Events, error)
	GetLastRecordsForUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Events, error)
	GetTomorrowEvents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Events, error)
}

type crmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrmServiceClient(cc grpc.ClientConnInterface) CrmServiceClient {
	return &crmServiceClient{cc}
}

func (c *crmServiceClient) GetUserByPhoneNumber(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/crm.crmService/GetUserByPhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crmServiceClient) GetFutureEventsForUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/crm.crmService/GetFutureEventsForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crmServiceClient) GetLastRecordsForUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/crm.crmService/GetLastRecordsForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crmServiceClient) GetTomorrowEvents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Events, error) {
	out := new(Events)
	err := c.cc.Invoke(ctx, "/crm.crmService/GetTomorrowEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrmServiceServer is the server API for CrmService service.
// All implementations should embed UnimplementedCrmServiceServer
// for forward compatibility
type CrmServiceServer interface {
	GetUserByPhoneNumber(context.Context, *User) (*User, error)
	GetFutureEventsForUser(context.Context, *User) (*Events, error)
	GetLastRecordsForUser(context.Context, *User) (*Events, error)
	GetTomorrowEvents(context.Context, *Empty) (*Events, error)
}

// UnimplementedCrmServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCrmServiceServer struct {
}

func (UnimplementedCrmServiceServer) GetUserByPhoneNumber(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByPhoneNumber not implemented")
}
func (UnimplementedCrmServiceServer) GetFutureEventsForUser(context.Context, *User) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFutureEventsForUser not implemented")
}
func (UnimplementedCrmServiceServer) GetLastRecordsForUser(context.Context, *User) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastRecordsForUser not implemented")
}
func (UnimplementedCrmServiceServer) GetTomorrowEvents(context.Context, *Empty) (*Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTomorrowEvents not implemented")
}

// UnsafeCrmServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrmServiceServer will
// result in compilation errors.
type UnsafeCrmServiceServer interface {
	mustEmbedUnimplementedCrmServiceServer()
}

func RegisterCrmServiceServer(s grpc.ServiceRegistrar, srv CrmServiceServer) {
	s.RegisterService(&CrmService_ServiceDesc, srv)
}

func _CrmService_GetUserByPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrmServiceServer).GetUserByPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crm.crmService/GetUserByPhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrmServiceServer).GetUserByPhoneNumber(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrmService_GetFutureEventsForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrmServiceServer).GetFutureEventsForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crm.crmService/GetFutureEventsForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrmServiceServer).GetFutureEventsForUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrmService_GetLastRecordsForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrmServiceServer).GetLastRecordsForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crm.crmService/GetLastRecordsForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrmServiceServer).GetLastRecordsForUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrmService_GetTomorrowEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrmServiceServer).GetTomorrowEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crm.crmService/GetTomorrowEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrmServiceServer).GetTomorrowEvents(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CrmService_ServiceDesc is the grpc.ServiceDesc for CrmService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrmService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crm.crmService",
	HandlerType: (*CrmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByPhoneNumber",
			Handler:    _CrmService_GetUserByPhoneNumber_Handler,
		},
		{
			MethodName: "GetFutureEventsForUser",
			Handler:    _CrmService_GetFutureEventsForUser_Handler,
		},
		{
			MethodName: "GetLastRecordsForUser",
			Handler:    _CrmService_GetLastRecordsForUser_Handler,
		},
		{
			MethodName: "GetTomorrowEvents",
			Handler:    _CrmService_GetTomorrowEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/pkg/proto/crm/crm.proto",
}

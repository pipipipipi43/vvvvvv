// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: accesskey.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// AccessKeyServiceClient is the client API for AccessKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessKeyServiceClient interface {
	QueryAccessKeys(ctx context.Context, in *QueryAccessKeysRequest, opts ...grpc.CallOption) (*QueryAccessKeysResponse, error)
	GetAccessKey(ctx context.Context, in *GetAccessKeysRequest, opts ...grpc.CallOption) (*GetAccessKeysResponse, error)
	CreateAccessKeys(ctx context.Context, in *CreateAccessKeysRequest, opts ...grpc.CallOption) (*CreateAccessKeysResponse, error)
	UpdateAccessKeys(ctx context.Context, in *UpdateAccessKeysRequest, opts ...grpc.CallOption) (*UpdateAccessKeysResponse, error)
	DeleteAccessKeys(ctx context.Context, in *DeleteAccessKeysRequest, opts ...grpc.CallOption) (*DeleteAccessKeysResponse, error)
}

type accessKeyServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewAccessKeyServiceClient(cc grpc1.ClientConnInterface) AccessKeyServiceClient {
	return &accessKeyServiceClient{cc}
}

func (c *accessKeyServiceClient) QueryAccessKeys(ctx context.Context, in *QueryAccessKeysRequest, opts ...grpc.CallOption) (*QueryAccessKeysResponse, error) {
	out := new(QueryAccessKeysResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.accesskey.AccessKeyService/QueryAccessKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) GetAccessKey(ctx context.Context, in *GetAccessKeysRequest, opts ...grpc.CallOption) (*GetAccessKeysResponse, error) {
	out := new(GetAccessKeysResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.accesskey.AccessKeyService/GetAccessKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) CreateAccessKeys(ctx context.Context, in *CreateAccessKeysRequest, opts ...grpc.CallOption) (*CreateAccessKeysResponse, error) {
	out := new(CreateAccessKeysResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.accesskey.AccessKeyService/CreateAccessKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) UpdateAccessKeys(ctx context.Context, in *UpdateAccessKeysRequest, opts ...grpc.CallOption) (*UpdateAccessKeysResponse, error) {
	out := new(UpdateAccessKeysResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.accesskey.AccessKeyService/UpdateAccessKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessKeyServiceClient) DeleteAccessKeys(ctx context.Context, in *DeleteAccessKeysRequest, opts ...grpc.CallOption) (*DeleteAccessKeysResponse, error) {
	out := new(DeleteAccessKeysResponse)
	err := c.cc.Invoke(ctx, "/erda.core.services.accesskey.AccessKeyService/DeleteAccessKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessKeyServiceServer is the server API for AccessKeyService service.
// All implementations should embed UnimplementedAccessKeyServiceServer
// for forward compatibility
type AccessKeyServiceServer interface {
	QueryAccessKeys(context.Context, *QueryAccessKeysRequest) (*QueryAccessKeysResponse, error)
	GetAccessKey(context.Context, *GetAccessKeysRequest) (*GetAccessKeysResponse, error)
	CreateAccessKeys(context.Context, *CreateAccessKeysRequest) (*CreateAccessKeysResponse, error)
	UpdateAccessKeys(context.Context, *UpdateAccessKeysRequest) (*UpdateAccessKeysResponse, error)
	DeleteAccessKeys(context.Context, *DeleteAccessKeysRequest) (*DeleteAccessKeysResponse, error)
}

// UnimplementedAccessKeyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccessKeyServiceServer struct {
}

func (*UnimplementedAccessKeyServiceServer) QueryAccessKeys(context.Context, *QueryAccessKeysRequest) (*QueryAccessKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAccessKeys not implemented")
}
func (*UnimplementedAccessKeyServiceServer) GetAccessKey(context.Context, *GetAccessKeysRequest) (*GetAccessKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessKey not implemented")
}
func (*UnimplementedAccessKeyServiceServer) CreateAccessKeys(context.Context, *CreateAccessKeysRequest) (*CreateAccessKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccessKeys not implemented")
}
func (*UnimplementedAccessKeyServiceServer) UpdateAccessKeys(context.Context, *UpdateAccessKeysRequest) (*UpdateAccessKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessKeys not implemented")
}
func (*UnimplementedAccessKeyServiceServer) DeleteAccessKeys(context.Context, *DeleteAccessKeysRequest) (*DeleteAccessKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccessKeys not implemented")
}

func RegisterAccessKeyServiceServer(s grpc1.ServiceRegistrar, srv AccessKeyServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_AccessKeyService_serviceDesc(srv, opts...), srv)
}

var _AccessKeyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.core.services.accesskey.AccessKeyService",
	HandlerType: (*AccessKeyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "accesskey.proto",
}

func _get_AccessKeyService_serviceDesc(srv AccessKeyServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_AccessKeyService_QueryAccessKeys_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.QueryAccessKeys(ctx, req.(*QueryAccessKeysRequest))
	}
	var _AccessKeyService_QueryAccessKeys_info transport.ServiceInfo
	if h.Interceptor != nil {
		_AccessKeyService_QueryAccessKeys_info = transport.NewServiceInfo("erda.core.services.accesskey.AccessKeyService", "QueryAccessKeys", srv)
		_AccessKeyService_QueryAccessKeys_Handler = h.Interceptor(_AccessKeyService_QueryAccessKeys_Handler)
	}

	_AccessKeyService_GetAccessKey_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.GetAccessKey(ctx, req.(*GetAccessKeysRequest))
	}
	var _AccessKeyService_GetAccessKey_info transport.ServiceInfo
	if h.Interceptor != nil {
		_AccessKeyService_GetAccessKey_info = transport.NewServiceInfo("erda.core.services.accesskey.AccessKeyService", "GetAccessKey", srv)
		_AccessKeyService_GetAccessKey_Handler = h.Interceptor(_AccessKeyService_GetAccessKey_Handler)
	}

	_AccessKeyService_CreateAccessKeys_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateAccessKeys(ctx, req.(*CreateAccessKeysRequest))
	}
	var _AccessKeyService_CreateAccessKeys_info transport.ServiceInfo
	if h.Interceptor != nil {
		_AccessKeyService_CreateAccessKeys_info = transport.NewServiceInfo("erda.core.services.accesskey.AccessKeyService", "CreateAccessKeys", srv)
		_AccessKeyService_CreateAccessKeys_Handler = h.Interceptor(_AccessKeyService_CreateAccessKeys_Handler)
	}

	_AccessKeyService_UpdateAccessKeys_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.UpdateAccessKeys(ctx, req.(*UpdateAccessKeysRequest))
	}
	var _AccessKeyService_UpdateAccessKeys_info transport.ServiceInfo
	if h.Interceptor != nil {
		_AccessKeyService_UpdateAccessKeys_info = transport.NewServiceInfo("erda.core.services.accesskey.AccessKeyService", "UpdateAccessKeys", srv)
		_AccessKeyService_UpdateAccessKeys_Handler = h.Interceptor(_AccessKeyService_UpdateAccessKeys_Handler)
	}

	_AccessKeyService_DeleteAccessKeys_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.DeleteAccessKeys(ctx, req.(*DeleteAccessKeysRequest))
	}
	var _AccessKeyService_DeleteAccessKeys_info transport.ServiceInfo
	if h.Interceptor != nil {
		_AccessKeyService_DeleteAccessKeys_info = transport.NewServiceInfo("erda.core.services.accesskey.AccessKeyService", "DeleteAccessKeys", srv)
		_AccessKeyService_DeleteAccessKeys_Handler = h.Interceptor(_AccessKeyService_DeleteAccessKeys_Handler)
	}

	var serviceDesc = _AccessKeyService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "QueryAccessKeys",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(QueryAccessKeysRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(AccessKeyServiceServer).QueryAccessKeys(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _AccessKeyService_QueryAccessKeys_info)
				}
				if interceptor == nil {
					return _AccessKeyService_QueryAccessKeys_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.accesskey.AccessKeyService/QueryAccessKeys",
				}
				return interceptor(ctx, in, info, _AccessKeyService_QueryAccessKeys_Handler)
			},
		},
		{
			MethodName: "GetAccessKey",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GetAccessKeysRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(AccessKeyServiceServer).GetAccessKey(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _AccessKeyService_GetAccessKey_info)
				}
				if interceptor == nil {
					return _AccessKeyService_GetAccessKey_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.accesskey.AccessKeyService/GetAccessKey",
				}
				return interceptor(ctx, in, info, _AccessKeyService_GetAccessKey_Handler)
			},
		},
		{
			MethodName: "CreateAccessKeys",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateAccessKeysRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(AccessKeyServiceServer).CreateAccessKeys(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _AccessKeyService_CreateAccessKeys_info)
				}
				if interceptor == nil {
					return _AccessKeyService_CreateAccessKeys_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.accesskey.AccessKeyService/CreateAccessKeys",
				}
				return interceptor(ctx, in, info, _AccessKeyService_CreateAccessKeys_Handler)
			},
		},
		{
			MethodName: "UpdateAccessKeys",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateAccessKeysRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(AccessKeyServiceServer).UpdateAccessKeys(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _AccessKeyService_UpdateAccessKeys_info)
				}
				if interceptor == nil {
					return _AccessKeyService_UpdateAccessKeys_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.accesskey.AccessKeyService/UpdateAccessKeys",
				}
				return interceptor(ctx, in, info, _AccessKeyService_UpdateAccessKeys_Handler)
			},
		},
		{
			MethodName: "DeleteAccessKeys",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(DeleteAccessKeysRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(AccessKeyServiceServer).DeleteAccessKeys(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _AccessKeyService_DeleteAccessKeys_info)
				}
				if interceptor == nil {
					return _AccessKeyService_DeleteAccessKeys_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.core.services.accesskey.AccessKeyService/DeleteAccessKeys",
				}
				return interceptor(ctx, in, info, _AccessKeyService_DeleteAccessKeys_Handler)
			},
		},
	}
	return &serviceDesc
}
// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: tenant.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/tenant/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// TenantService tenant.proto
	TenantService() pb.TenantServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		tenantService: pb.NewTenantServiceClient(cc),
	}
}

type serviceClients struct {
	tenantService pb.TenantServiceClient
}

func (c *serviceClients) TenantService() pb.TenantServiceClient {
	return c.tenantService
}

type tenantServiceWrapper struct {
	client pb.TenantServiceClient
	opts   []grpc1.CallOption
}

func (s *tenantServiceWrapper) CreateTenant(ctx context.Context, req *pb.CreateTenantRequest) (*pb.CreateTenantResponse, error) {
	return s.client.CreateTenant(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *tenantServiceWrapper) GetTenant(ctx context.Context, req *pb.GetTenantRequest) (*pb.GetTenantResponse, error) {
	return s.client.GetTenant(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *tenantServiceWrapper) DeleteTenant(ctx context.Context, req *pb.DeleteTenantRequest) (*pb.DeleteTenantResponse, error) {
	return s.client.DeleteTenant(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

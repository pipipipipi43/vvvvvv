// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: authentication.proto

package client

import (
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/services/authentication/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// AuthenticationService authentication.proto
	AuthenticationService() pb.AuthenticationServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		authenticationService: pb.NewAuthenticationServiceClient(cc),
	}
}

type serviceClients struct {
	authenticationService pb.AuthenticationServiceClient
}

func (c *serviceClients) AuthenticationService() pb.AuthenticationServiceClient {
	return c.authenticationService
}

type authenticationServiceWrapper struct {
	client pb.AuthenticationServiceClient
	opts   []grpc1.CallOption
}

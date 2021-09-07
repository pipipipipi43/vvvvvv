// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: adapter.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/adapter/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// InstrumentationLibraryService adapter.proto
	InstrumentationLibraryService() pb.InstrumentationLibraryServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		instrumentationLibraryService: pb.NewInstrumentationLibraryServiceClient(cc),
	}
}

type serviceClients struct {
	instrumentationLibraryService pb.InstrumentationLibraryServiceClient
}

func (c *serviceClients) InstrumentationLibraryService() pb.InstrumentationLibraryServiceClient {
	return c.instrumentationLibraryService
}

type instrumentationLibraryServiceWrapper struct {
	client pb.InstrumentationLibraryServiceClient
	opts   []grpc1.CallOption
}

func (s *instrumentationLibraryServiceWrapper) GetInstrumentationLibrary(ctx context.Context, req *pb.GetInstrumentationLibraryRequest) (*pb.GetInstrumentationLibraryResponse, error) {
	return s.client.GetInstrumentationLibrary(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *instrumentationLibraryServiceWrapper) GetInstrumentationLibraryDocs(ctx context.Context, req *pb.GetInstrumentationLibraryDocsRequest) (*pb.GetInstrumentationLibraryDocsResponse, error) {
	return s.client.GetInstrumentationLibraryDocs(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: metric.proto, meta.proto

package client

import (
	context "context"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/metric/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// MetricService metric.proto
	MetricService() pb.MetricServiceClient
	// MetricMetaService meta.proto
	MetricMetaService() pb.MetricMetaServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		metricService:     pb.NewMetricServiceClient(cc),
		metricMetaService: pb.NewMetricMetaServiceClient(cc),
	}
}

type serviceClients struct {
	metricService     pb.MetricServiceClient
	metricMetaService pb.MetricMetaServiceClient
}

func (c *serviceClients) MetricService() pb.MetricServiceClient {
	return c.metricService
}

func (c *serviceClients) MetricMetaService() pb.MetricMetaServiceClient {
	return c.metricMetaService
}

type metricServiceWrapper struct {
	client pb.MetricServiceClient
	opts   []grpc1.CallOption
}

func (s *metricServiceWrapper) QueryWithInfluxFormat(ctx context.Context, req *pb.QueryWithInfluxFormatRequest) (*pb.QueryWithInfluxFormatResponse, error) {
	return s.client.QueryWithInfluxFormat(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *metricServiceWrapper) SearchWithInfluxFormat(ctx context.Context, req *pb.QueryWithInfluxFormatRequest) (*pb.QueryWithInfluxFormatResponse, error) {
	return s.client.SearchWithInfluxFormat(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *metricServiceWrapper) QueryWithTableFormat(ctx context.Context, req *pb.QueryWithTableFormatRequest) (*pb.QueryWithTableFormatResponse, error) {
	return s.client.QueryWithTableFormat(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *metricServiceWrapper) SearchWithTableFormat(ctx context.Context, req *pb.QueryWithTableFormatRequest) (*pb.QueryWithTableFormatResponse, error) {
	return s.client.SearchWithTableFormat(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

type metricMetaServiceWrapper struct {
	client pb.MetricMetaServiceClient
	opts   []grpc1.CallOption
}

func (s *metricMetaServiceWrapper) ListMetricNames(ctx context.Context, req *pb.ListMetricNamesRequest) (*pb.ListMetricNamesResponse, error) {
	return s.client.ListMetricNames(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *metricMetaServiceWrapper) ListMetricMeta(ctx context.Context, req *pb.ListMetricMetaRequest) (*pb.ListMetricMetaResponse, error) {
	return s.client.ListMetricMeta(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *metricMetaServiceWrapper) ListMetricGroups(ctx context.Context, req *pb.ListMetricGroupsRequest) (*pb.ListMetricGroupsResponse, error) {
	return s.client.ListMetricGroups(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *metricMetaServiceWrapper) GetMetricGroup(ctx context.Context, req *pb.GetMetricGroupRequest) (*pb.GetMetricGroupResponse, error) {
	return s.client.GetMetricGroup(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
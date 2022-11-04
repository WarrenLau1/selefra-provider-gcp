package logging

import (
	"context"
	"fmt"
	"net"
	"testing"

	logging "cloud.google.com/go/logging/apiv2"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/logging/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeMetricsServer struct {
	pb.UnimplementedMetricsServiceV2Server
}

func (f *fakeMetricsServer) ListLogMetrics(context.Context, *pb.ListLogMetricsRequest) (*pb.ListLogMetricsResponse, error) {
	resp := pb.ListLogMetricsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createMetrics() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeMetricsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterMetricsServiceV2Server(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := logging.NewMetricsClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		LoggingMetricsClient: svc,
	}, nil
}

func TestMetrics(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpLoggingMetricsGenerator{}), createMetrics, gcp_client.TestOptions{})
}

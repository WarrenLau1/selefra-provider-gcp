package monitoring

import (
	"context"
	"fmt"
	"net"
	"testing"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/monitoring/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeAlertPoliciesServer struct {
	pb.UnimplementedAlertPolicyServiceServer
}

func (f *fakeAlertPoliciesServer) ListAlertPolicies(context.Context, *pb.ListAlertPoliciesRequest) (*pb.ListAlertPoliciesResponse, error) {
	resp := pb.ListAlertPoliciesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createAlertPolicies() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeAlertPoliciesServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterAlertPolicyServiceServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := monitoring.NewAlertPolicyClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		MonitoringAlertPolicyClient: svc,
	}, nil
}

func TestAlertPolicies(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpMonitoringAlertPoliciesGenerator{}), createAlertPolicies, gcp_client.TestOptions{})
}

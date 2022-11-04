package run

import (
	"context"
	"fmt"
	"net"
	"testing"

	run "cloud.google.com/go/run/apiv2"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/run/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeServicesServer struct {
	pb.UnimplementedServicesServer
}

func (f *fakeServicesServer) ListServices(context.Context, *pb.ListServicesRequest) (*pb.ListServicesResponse, error) {
	resp := pb.ListServicesResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createServices() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeServicesServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterServicesServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := run.NewServicesClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		RunServicesClient: svc,
	}, nil
}

func TestServices(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpRunServicesGenerator{}), createServices, gcp_client.TestOptions{})
}

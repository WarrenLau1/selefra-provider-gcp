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

type fakeSinksServer struct {
	pb.UnimplementedConfigServiceV2Server
}

func (f *fakeSinksServer) ListSinks(context.Context, *pb.ListSinksRequest) (*pb.ListSinksResponse, error) {
	resp := pb.ListSinksResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createSinks() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeSinksServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterConfigServiceV2Server(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := logging.NewConfigClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		LoggingConfigClient: svc,
	}, nil
}

func TestSinks(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpLoggingSinksGenerator{}), createSinks, gcp_client.TestOptions{})
}

package container

import (
	"context"
	"fmt"
	"net"
	"testing"

	container "cloud.google.com/go/container/apiv1"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/container/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeClustersServer struct {
	pb.UnimplementedClusterManagerServer
}

func (f *fakeClustersServer) ListClusters(context.Context, *pb.ListClustersRequest) (*pb.ListClustersResponse, error) {
	resp := pb.ListClustersResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	return &resp, nil
}

func createClusters() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeClustersServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterClusterManagerServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := container.NewClusterManagerClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		ContainerClusterManagerClient: svc,
	}, nil
}

func TestClusters(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpContainerClustersGenerator{}), createClusters, gcp_client.TestOptions{})
}

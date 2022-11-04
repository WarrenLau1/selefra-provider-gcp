package resourcemanager

import (
	"context"
	"fmt"
	"net"
	"testing"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeFoldersServer struct {
	pb.UnimplementedFoldersServer
}

func (f *fakeFoldersServer) ListFolders(context.Context, *pb.ListFoldersRequest) (*pb.ListFoldersResponse, error) {
	resp := pb.ListFoldersResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createFolders() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeFoldersServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterFoldersServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := resourcemanager.NewFoldersClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		ResourcemanagerFoldersClient: svc,
	}, nil
}

func TestFolders(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpResourcemanagerFoldersGenerator{}), createFolders, gcp_client.TestOptions{})
}

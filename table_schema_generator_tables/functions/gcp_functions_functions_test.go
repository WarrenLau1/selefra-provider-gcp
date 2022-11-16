package functions

import (
	"context"
	"fmt"
	"net"
	"testing"

	functions "cloud.google.com/go/functions/apiv1"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/functions/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeFunctionsServer struct {
	pb.UnimplementedCloudFunctionsServiceServer
}

func (f *fakeFunctionsServer) ListFunctions(context.Context, *pb.ListFunctionsRequest) (*pb.ListFunctionsResponse, error) {
	resp := pb.ListFunctionsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createFunctions() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeFunctionsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterCloudFunctionsServiceServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := functions.NewCloudFunctionsClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		FunctionsCloudFunctionsClient: svc,
	}, nil
}

func TestFunctions(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpFunctionsFunctionsGenerator{}), createFunctions, gcp_client.TestOptions{})
}

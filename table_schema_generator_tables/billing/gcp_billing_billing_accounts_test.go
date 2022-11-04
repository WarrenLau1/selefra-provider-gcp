package billing

import (
	"context"
	"fmt"
	"net"
	"testing"

	billing "cloud.google.com/go/billing/apiv1"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/billing/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeBillingAccountsServer struct {
	pb.UnimplementedCloudBillingServer
}

func (f *fakeBillingAccountsServer) ListBillingAccounts(context.Context, *pb.ListBillingAccountsRequest) (*pb.ListBillingAccountsResponse, error) {
	resp := pb.ListBillingAccountsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createBillingAccounts() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeBillingAccountsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterCloudBillingServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := billing.NewCloudBillingClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		BillingCloudBillingClient: svc,
	}, nil
}

func TestBillingAccounts(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpBillingBillingAccountsGenerator{}), createBillingAccounts, gcp_client.TestOptions{})
}

package domains

import (
	"context"
	"fmt"
	"net"
	"testing"

	domains "cloud.google.com/go/domains/apiv1beta1"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/domains/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeRegistrationsServer struct {
	pb.UnimplementedDomainsServer
}

func (f *fakeRegistrationsServer) ListRegistrations(context.Context, *pb.ListRegistrationsRequest) (*pb.ListRegistrationsResponse, error) {
	resp := pb.ListRegistrationsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createRegistrations() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeRegistrationsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterDomainsServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := domains.NewClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		DomainsClient: svc,
	}, nil
}

func TestRegistrations(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpDomainsRegistrationsGenerator{}), createRegistrations, gcp_client.TestOptions{})
}

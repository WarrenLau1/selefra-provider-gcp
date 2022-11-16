package secretmanager

import (
	"context"
	"fmt"
	"net"
	"testing"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type fakeSecretsServer struct {
	pb.UnimplementedSecretManagerServiceServer
}

func (f *fakeSecretsServer) ListSecrets(context.Context, *pb.ListSecretsRequest) (*pb.ListSecretsResponse, error) {
	resp := pb.ListSecretsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func createSecrets() (*gcp_client.GcpServices, error) {
	fakeServer := &fakeSecretsServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterSecretManagerServiceServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	svc, err := secretmanager.NewClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &gcp_client.GcpServices{
		SecretmanagerClient: svc,
	}, nil
}

func TestSecrets(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpSecretmanagerSecretsGenerator{}), createSecrets, gcp_client.TestOptions{})
}

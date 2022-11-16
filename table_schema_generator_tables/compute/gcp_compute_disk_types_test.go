package compute

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	compute "cloud.google.com/go/compute/apiv1"
	"github.com/julienschmidt/httprouter"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func createDiskTypes() (*gcp_client.GcpServices, error) {
	var item pb.DiskTypeAggregatedList
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}
	emptyStr := ""
	item.NextPageToken = &emptyStr
	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	svc, err := compute.NewDiskTypesRESTClient(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &gcp_client.GcpServices{
		ComputeDiskTypesClient: svc,
	}, nil
}

func TestDiskTypes(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpComputeDiskTypesGenerator{}), createDiskTypes, gcp_client.TestOptions{})
}

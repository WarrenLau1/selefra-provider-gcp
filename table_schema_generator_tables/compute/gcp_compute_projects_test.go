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

func createProjects() (*gcp_client.GcpServices, error) {
	var item pb.Project
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}
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
	svc, err := compute.NewProjectsRESTClient(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &gcp_client.GcpServices{
		ComputeProjectsClient: svc,
	}, nil
}

func TestProjects(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpComputeProjectsGenerator{}), createProjects, gcp_client.TestOptions{})
}

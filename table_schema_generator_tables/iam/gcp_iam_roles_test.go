package iam

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

type MockRolesResult struct {
	Roles []*iam.Role `json:"roles,omitempty"`
}

func createRoles() (*gcp_client.GcpServices, error) {
	var item iam.Role
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}

	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockRolesResult{
			Roles: []*iam.Role{&item},
		}
		b, err := json.Marshal(resp)
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
	svc, err := iam.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &gcp_client.GcpServices{
		Iam: svc,
	}, nil
}

func TestRoles(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpIamRolesGenerator{}), createRoles, gcp_client.TestOptions{})
}

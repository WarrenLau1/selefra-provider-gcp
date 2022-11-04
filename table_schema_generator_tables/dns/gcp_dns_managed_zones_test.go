package dns

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
	"google.golang.org/api/dns/v1"
	"google.golang.org/api/option"
)

type MockManagedZonesResult struct {
	ManagedZones []*dns.ManagedZone `json:"managedzones,omitempty"`
}

func createManagedZones() (*gcp_client.GcpServices, error) {
	var item dns.ManagedZone
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}

	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockManagedZonesResult{
			ManagedZones: []*dns.ManagedZone{&item},
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
	svc, err := dns.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &gcp_client.GcpServices{
		Dns: svc,
	}, nil
}

func TestManagedZones(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpDnsManagedZonesGenerator{}), createManagedZones, gcp_client.TestOptions{})
}

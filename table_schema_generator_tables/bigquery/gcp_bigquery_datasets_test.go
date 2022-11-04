package bigquery

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/selefra/selefra-provider-gcp/faker"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"google.golang.org/api/bigquery/v2"
	"google.golang.org/api/option"
)

func createBigqueryDatasets() (*gcp_client.GcpServices, error) {
	id := "testDataset"
	mux := httprouter.New()
	var dataset bigquery.Dataset
	if err := faker.FakeObject(&dataset); err != nil {
		return nil, err
	}
	dataset.Id = id
	dataset.DatasetReference = &bigquery.DatasetReference{
		DatasetId: id,
	}
	mux.GET("/projects/testProject/datasets", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &bigquery.DatasetList{
			Datasets: []*bigquery.DatasetListDatasets{
				{
					DatasetReference: dataset.DatasetReference,
				},
			},
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

	mux.GET("/projects/testProject/datasets/testDataset", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&dataset)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	mux.GET("/projects/testProject/datasets/testDataset/tables", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &bigquery.TableList{
			Tables: []*bigquery.TableListTables{
				{
					Id: id,
					TableReference: &bigquery.TableReference{
						TableId: id,
					},
				},
			},
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

	var table bigquery.Table
	if err := faker.FakeObject(&table); err != nil {
		return nil, err
	}
	table.Id = id
	table.TableReference = &bigquery.TableReference{
		TableId: id,
	}
	schema := bigquery.TableSchema{
		Fields: []*bigquery.TableFieldSchema{{
			Name: "test",
			Type: "test",
		},
		},
	}
	table.Schema = &schema

	table.ExternalDataConfiguration = &bigquery.ExternalDataConfiguration{
		Autodetect: true,
		Schema:     &schema,
		SourceUris: []string{"test"},
	}
	table.Labels = map[string]string{
		"test": "test",
	}
	table.Clustering = &bigquery.Clustering{
		Fields: []string{"test"},
	}
	if err := faker.FakeObject(&table.Description); err != nil {
		return nil, err
	}
	if err := faker.FakeObject(&table.EncryptionConfiguration); err != nil {
		return nil, err
	}

	mux.GET("/projects/testProject/datasets/testDataset/tables/:table", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("what")
		fmt.Println(r.URL)
		b, err := json.Marshal(&table)
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
	svc, err := bigquery.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &gcp_client.GcpServices{
		BigqueryService: svc,
	}, nil
}

func TestBigqueryDatasets(t *testing.T) {
	gcp_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableGcpBigqueryDatasetsGenerator{}), createBigqueryDatasets, gcp_client.TestOptions{})
}

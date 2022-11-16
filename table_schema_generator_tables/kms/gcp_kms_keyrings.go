package kms

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/cloudkms/v1"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/cloud/kms/v1"
)

type TableGcpKmsKeyringsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpKmsKeyringsGenerator{}

func (x *TableGcpKmsKeyringsGenerator) GetTableName() string {
	return "gcp_kms_keyrings"
}

func (x *TableGcpKmsKeyringsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpKmsKeyringsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpKmsKeyringsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpKmsKeyringsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			locations, err := getAllKmsLocations(c)
			if err != nil {
				maybeError := errors.WithStack(fmt.Errorf("failed to get kms locations. %w", err))
				return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
			}
			for _, l := range locations {
				it := c.GcpServices.KmsKeyManagementClient.ListKeyRings(ctx, &pb.ListKeyRingsRequest{
					Parent: l.Name,
				})
				for {
					resp, err := it.Next()
					if err == iterator.Done {
						break
					}
					if err != nil {
						maybeError := errors.WithStack(err)
						return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
					}

					resultChannel <- resp
				}
			}
			return nil
		},
	}
}

func getAllKmsLocations(c *gcp_client.Client) ([]*cloudkms.Location, error) {
	var locations []*cloudkms.Location
	call := c.GcpServices.KmsoldService.Projects.Locations.List("projects/" + c.ProjectId)
	nextPageToken := ""
	for {
		resp, err := call.PageToken(nextPageToken).Do()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		locations = append(locations, resp.Locations...)

		if resp.NextPageToken == "" {
			break
		}
		nextPageToken = resp.NextPageToken
	}
	return locations, nil
}

func (x *TableGcpKmsKeyringsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpKmsKeyringsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpKmsKeyringsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpKmsCryptoKeysGenerator{}),
	}
}

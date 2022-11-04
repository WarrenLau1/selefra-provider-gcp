package bigquery

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/bigquery/v2"
)

type TableGcpBigqueryDatasetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBigqueryDatasetsGenerator{}

func (x *TableGcpBigqueryDatasetsGenerator) GetTableName() string {
	return "gcp_bigquery_datasets"
}

func (x *TableGcpBigqueryDatasetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBigqueryDatasetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBigqueryDatasetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpBigqueryDatasetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""
			for {
				output, err := c.GcpServices.BigqueryService.Datasets.List(c.ProjectId).PageToken(nextPageToken).Do()
				if err != nil {
					maybeError := errors.WithStack(err)
					return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
				}
				gcp_client.SendResults(resultChannel, output.Datasets, func(result any) (any, error) {
					c := client.(*gcp_client.Client)
					datasetListDataset := result.(*bigquery.DatasetListDatasets)
					item, err := c.GcpServices.BigqueryService.Datasets.Get(c.ProjectId, datasetListDataset.DatasetReference.DatasetId).Do()
					if err != nil {
						return nil, errors.WithStack(err)
					}
					return item, nil

				})
				if output.NextPageToken == "" {
					break
				}
				nextPageToken = output.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableGcpBigqueryDatasetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpBigqueryDatasetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dataset_reference").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_encryption_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_time_travel_hours").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_partition_expiration_ms").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("friendly_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfies_pzs").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_collation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_table_expiration_ms").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProtoEtag()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_case_insensitive").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableGcpBigqueryDatasetsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpBigqueryTablesGenerator{}),
	}
}

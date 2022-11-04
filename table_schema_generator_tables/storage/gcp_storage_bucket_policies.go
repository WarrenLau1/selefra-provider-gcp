package storage

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
)

type TableGcpStorageBucketPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpStorageBucketPoliciesGenerator{}

func (x *TableGcpStorageBucketPoliciesGenerator) GetTableName() string {
	return "gcp_storage_bucket_policies"
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"bucket_name",
		},
	}
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			bkt := task.ParentRow.GetOrDefault("name", nil).(*string)
			output, err := c.GcpServices.StorageClient.Bucket(*bkt).IAM().V3().Policy(ctx)
			if err != nil {
				maybeError := errors.WithStack(err)
				return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
			}
			resultChannel <- output
			return nil
		},
	}
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bindings").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_storage_buckets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_storage_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}

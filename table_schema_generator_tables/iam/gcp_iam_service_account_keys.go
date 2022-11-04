package iam

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iam/v1"
)

type TableGcpIamServiceAccountKeysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpIamServiceAccountKeysGenerator{}

func (x *TableGcpIamServiceAccountKeysGenerator) GetTableName() string {
	return "gcp_iam_service_account_keys"
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"service_account_unique_id",
			"name",
		},
	}
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := task.ParentRawResult.(*iam.ServiceAccount)

			output, err := c.GcpServices.Iam.Projects.ServiceAccounts.Keys.List(p.Name).Context(ctx).Do()
			if err != nil {
				maybeError := errors.WithStack(err)
				return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
			}

			resultChannel <- output.Keys
			return nil
		},
	}
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("key_algorithm").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_key_data").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_after_time").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_before_time").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account_unique_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("unique_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_origin").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_key_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_iam_service_accounts_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_iam_service_accounts.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetSubTables() []*schema.Table {
	return nil
}

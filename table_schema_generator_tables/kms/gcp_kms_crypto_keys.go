package kms

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/genproto/googleapis/cloud/kms/v1"
)

type TableGcpKmsCryptoKeysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpKmsCryptoKeysGenerator{}

func (x *TableGcpKmsCryptoKeysGenerator) GetTableName() string {
	return "gcp_kms_crypto_keys"
}

func (x *TableGcpKmsCryptoKeysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpKmsCryptoKeysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpKmsCryptoKeysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpKmsCryptoKeysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := task.ParentRawResult.(*kms.KeyRing)

			nextPageToken := ""
			call := c.GcpServices.KmsoldService.Projects.Locations.KeyRings.CryptoKeys.List(p.Name).Context(ctx)
			for {
				call.PageToken(nextPageToken)
				resp, err := call.Do()
				if err != nil {
					maybeError := errors.WithStack(err)
					return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
				}
				resultChannel <- resp.CryptoKeys

				if resp.NextPageToken == "" {
					break
				}
				nextPageToken = resp.NextPageToken
			}

			return nil
		},
	}
}

func (x *TableGcpKmsCryptoKeysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpKmsCryptoKeysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("primary").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purpose").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_kms_keyrings_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_kms_keyrings.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_only").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_period").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_template").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crypto_key_backend").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_rotation_time").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("destroy_scheduled_duration").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableGcpKmsCryptoKeysGenerator) GetSubTables() []*schema.Table {
	return nil
}

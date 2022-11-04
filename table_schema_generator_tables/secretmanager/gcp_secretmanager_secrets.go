package secretmanager

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type TableGcpSecretmanagerSecretsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpSecretmanagerSecretsGenerator{}

func (x *TableGcpSecretmanagerSecretsGenerator) GetTableName() string {
	return "gcp_secretmanager_secrets"
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"name",
		},
	}
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListSecretsRequest{
				Parent: "projects/" + c.ProjectId,
			}
			it := c.GcpServices.SecretmanagerClient.ListSecrets(ctx, req)
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
			return nil
		},
	}
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("topics").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_aliases").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProtoEtag()).Build(),
	}
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetSubTables() []*schema.Table {
	return nil
}

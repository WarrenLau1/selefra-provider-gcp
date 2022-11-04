package run

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/cloud/run/v2"
)

type TableGcpRunServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpRunServicesGenerator{}

func (x *TableGcpRunServicesGenerator) GetTableName() string {
	return "gcp_run_services"
}

func (x *TableGcpRunServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpRunServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpRunServicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpRunServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListServicesRequest{
				Parent: "projects/" + c.ProjectId + "locations/-",
			}
			it := c.GcpServices.RunServicesClient.ListServices(ctx, req)
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

func (x *TableGcpRunServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpRunServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("launch_stage").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reconciling").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_ready_revision").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_created_revision").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("DeleteTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modifier").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("binary_authorization").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("terminal_condition").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProtoEtag()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("ExpireTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("traffic").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conditions").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("traffic_statuses").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uri").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingress").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("observed_generation").ColumnType(schema.ColumnTypeBigInt).Build(),
	}
}

func (x *TableGcpRunServicesGenerator) GetSubTables() []*schema.Table {
	return nil
}

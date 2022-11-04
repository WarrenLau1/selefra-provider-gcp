package resourcemanager

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
)

type TableGcpResourcemanagerProjectsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpResourcemanagerProjectsGenerator{}

func (x *TableGcpResourcemanagerProjectsGenerator) GetTableName() string {
	return "gcp_resourcemanager_projects"
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.GetProjectRequest{
				Name: "projects/" + c.ProjectId,
			}
			output, err := c.GcpServices.ResourcemanagerProjectsClient.GetProject(ctx, req)
			if err != nil {
				maybeError := errors.WithStack(err)
				return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
			}
			resultChannel <- output
			return nil
		},
	}
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProtoEtag()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("DeleteTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("UpdateTime")).Build(),
	}
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetSubTables() []*schema.Table {
	return nil
}

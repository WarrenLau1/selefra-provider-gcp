package compute

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

type TableGcpComputeAutoscalersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeAutoscalersGenerator{}

func (x *TableGcpComputeAutoscalersGenerator) GetTableName() string {
	return "gcp_compute_autoscalers"
}

func (x *TableGcpComputeAutoscalersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeAutoscalersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeAutoscalersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"self_link",
		},
	}
}

func (x *TableGcpComputeAutoscalersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListAutoscalersRequest{
				Project: c.ProjectId,
			}
			it := c.GcpServices.ComputeAutoscalersClient.AggregatedList(ctx, req)
			for {
				resp, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					maybeError := errors.WithStack(err)
					return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
				}

				resultChannel <- resp.Value.Autoscalers

			}
			return nil
		},
	}
}

func (x *TableGcpComputeAutoscalersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpComputeAutoscalersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_schedule_status").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recommended_size").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_details").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("autoscaling_policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableGcpComputeAutoscalersGenerator) GetSubTables() []*schema.Table {
	return nil
}

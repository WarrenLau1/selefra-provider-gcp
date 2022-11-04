package redis

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	pb "google.golang.org/genproto/googleapis/cloud/redis/v1"
)

type TableGcpRedisInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpRedisInstancesGenerator{}

func (x *TableGcpRedisInstancesGenerator) GetTableName() string {
	return "gcp_redis_instances"
}

func (x *TableGcpRedisInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpRedisInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpRedisInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"name",
		},
	}
}

func (x *TableGcpRedisInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListInstancesRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			it := c.GcpServices.RedisCloudRedisClient.ListInstances(ctx, req)
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

func (x *TableGcpRedisInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpRedisInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("memory_size_gb").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("persistence_iam_identity").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(gcp_client.ExtractorProtoTimestamp("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redis_configs").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_count").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redis_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorized_network").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("server_ca_certs").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nodes").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_ip_range").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tier").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_schedule").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_replicas_mode").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alternative_location_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reserved_ip_range").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_location_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_encryption_mode").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_endpoint").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_endpoint_port").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connect_mode").ColumnType(schema.ColumnTypeBigInt).Build(),
	}
}

func (x *TableGcpRedisInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}

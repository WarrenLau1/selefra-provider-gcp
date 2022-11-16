package storage

import (
	"context"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpStorageBucketsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpStorageBucketsGenerator{}

func (x *TableGcpStorageBucketsGenerator) GetTableName() string {
	return "gcp_storage_buckets"
}

func (x *TableGcpStorageBucketsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpStorageBucketsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpStorageBucketsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"name",
		},
	}
}

func (x *TableGcpStorageBucketsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			it := c.GcpServices.StorageClient.Buckets(ctx, c.ProjectId)
			for {
				bucket, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					maybeError := errors.WithStack(err)
					return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
				}
				resultChannel <- bucket
			}
			return nil
		},
	}
}

func (x *TableGcpStorageBucketsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpStorageBucketsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acl").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ACL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("website").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProtoEtag()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_policy_only").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_number").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uniform_bucket_level_access").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_object_acl").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultObjectACL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_placement_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versioning_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_class").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CORS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_event_based_hold").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("predefined_default_object_acl").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PredefinedDefaultObjectACL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("meta_generation").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requester_pays").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_access_prevention").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("predefined_acl").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PredefinedACL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rpo").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RPO")).Build(),
	}
}

func (x *TableGcpStorageBucketsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpStorageBucketPoliciesGenerator{}),
	}
}

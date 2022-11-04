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

type TableGcpComputeForwardingRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeForwardingRulesGenerator{}

func (x *TableGcpComputeForwardingRulesGenerator) GetTableName() string {
	return "gcp_compute_forwarding_rules"
}

func (x *TableGcpComputeForwardingRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeForwardingRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeForwardingRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"self_link",
		},
	}
}

func (x *TableGcpComputeForwardingRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListForwardingRulesRequest{
				Project: c.ProjectId,
			}
			it := c.GcpServices.ComputeForwardingRulesClient.AggregatedList(ctx, req)
			for {
				resp, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					maybeError := errors.WithStack(err)
					return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
				}

				resultChannel <- resp.Value.ForwardingRules

			}
			return nil
		},
	}
}

func (x *TableGcpComputeForwardingRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpComputeForwardingRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("all_ports").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata_filters").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ports").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_protocol").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IPProtocol")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnetwork").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("no_automate_dns_zone").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backend_service").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_tier").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("psc_connection_id").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_directory_registrations").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_fingerprint").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port_range").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancing_scheme").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_mirroring_collector").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("psc_connection_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_label").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_global_access").ColumnType(schema.ColumnTypeBool).Build(),
	}
}

func (x *TableGcpComputeForwardingRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}

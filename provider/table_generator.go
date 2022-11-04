package provider

import (
	"github.com/selefra/selefra-provider-sdk/provider/schema"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/bigquery"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/billing"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/compute"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/container"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/dns"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/domains"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/functions"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/iam"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/kms"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/logging"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/monitoring"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/redis"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/resourcemanager"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/run"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/secretmanager"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/serviceusage"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/sql"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator_tables/storage"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&container.TableGcpContainerClustersGenerator{}),
		table_schema_generator.GenTableSchema(&functions.TableGcpFunctionsFunctionsGenerator{}),
		table_schema_generator.GenTableSchema(&monitoring.TableGcpMonitoringAlertPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&resourcemanager.TableGcpResourcemanagerProjectPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&resourcemanager.TableGcpResourcemanagerProjectsGenerator{}),
		table_schema_generator.GenTableSchema(&resourcemanager.TableGcpResourcemanagerFoldersGenerator{}),
		table_schema_generator.GenTableSchema(&sql.TableGcpSqlInstancesGenerator{}),
		table_schema_generator.GenTableSchema(&storage.TableGcpStorageBucketsGenerator{}),
		table_schema_generator.GenTableSchema(&redis.TableGcpRedisInstancesGenerator{}),
		table_schema_generator.GenTableSchema(&run.TableGcpRunServicesGenerator{}),
		table_schema_generator.GenTableSchema(&bigquery.TableGcpBigqueryDatasetsGenerator{}),
		table_schema_generator.GenTableSchema(&logging.TableGcpLoggingSinksGenerator{}),
		table_schema_generator.GenTableSchema(&logging.TableGcpLoggingMetricsGenerator{}),
		table_schema_generator.GenTableSchema(&secretmanager.TableGcpSecretmanagerSecretsGenerator{}),
		table_schema_generator.GenTableSchema(&billing.TableGcpBillingServicesGenerator{}),
		table_schema_generator.GenTableSchema(&billing.TableGcpBillingBillingAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeTargetSslProxiesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeAddressesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeBackendServicesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeDisksGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeSslPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeInterconnectsGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeTargetHttpProxiesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeInstanceGroupsGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeForwardingRulesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeImagesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeDiskTypesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeInstancesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeSslCertificatesGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeVpnGatewaysGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeSubnetworksGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeNetworksGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeUrlMapsGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeAutoscalersGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeProjectsGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableGcpComputeFirewallsGenerator{}),
		table_schema_generator.GenTableSchema(&dns.TableGcpDnsManagedZonesGenerator{}),
		table_schema_generator.GenTableSchema(&dns.TableGcpDnsPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&domains.TableGcpDomainsRegistrationsGenerator{}),
		table_schema_generator.GenTableSchema(&iam.TableGcpIamRolesGenerator{}),
		table_schema_generator.GenTableSchema(&iam.TableGcpIamServiceAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&kms.TableGcpKmsKeyringsGenerator{}),
		table_schema_generator.GenTableSchema(&serviceusage.TableGcpServiceusageServicesGenerator{}),
	}
}

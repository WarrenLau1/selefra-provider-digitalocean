package provider

import (
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/accounts"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/balances"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/billing_history"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/cdns"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/certificates"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/databases"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/domains"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/droplets"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/firewalls"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/floating_ips"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/images"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/keys"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/monitoring"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/projects"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/registries"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/sizes"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/snapshots"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/spaces"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/storage"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator_tables/vpcs"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&registries.TableDigitaloceanRegistriesGenerator{}),
		table_schema_generator.GenTableSchema(&spaces.TableDigitaloceanSpacesGenerator{}),
		table_schema_generator.GenTableSchema(&accounts.TableDigitaloceanAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&balances.TableDigitaloceanBalancesGenerator{}),
		table_schema_generator.GenTableSchema(&billing_history.TableDigitaloceanBillingHistoryGenerator{}),
		table_schema_generator.GenTableSchema(&droplets.TableDigitaloceanDropletsGenerator{}),
		table_schema_generator.GenTableSchema(&floating_ips.TableDigitaloceanFloatingIpsGenerator{}),
		table_schema_generator.GenTableSchema(&sizes.TableDigitaloceanSizesGenerator{}),
		table_schema_generator.GenTableSchema(&snapshots.TableDigitaloceanSnapshotsGenerator{}),
		table_schema_generator.GenTableSchema(&vpcs.TableDigitaloceanVpcsGenerator{}),
		table_schema_generator.GenTableSchema(&domains.TableDigitaloceanDomainsGenerator{}),
		table_schema_generator.GenTableSchema(&images.TableDigitaloceanImagesGenerator{}),
		table_schema_generator.GenTableSchema(&keys.TableDigitaloceanKeysGenerator{}),
		table_schema_generator.GenTableSchema(&monitoring.TableDigitaloceanMonitoringAlertPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&projects.TableDigitaloceanProjectsGenerator{}),
		table_schema_generator.GenTableSchema(&certificates.TableDigitaloceanCertificatesGenerator{}),
		table_schema_generator.GenTableSchema(&firewalls.TableDigitaloceanFirewallsGenerator{}),
		table_schema_generator.GenTableSchema(&cdns.TableDigitaloceanCdnsGenerator{}),
		table_schema_generator.GenTableSchema(&databases.TableDigitaloceanDatabasesGenerator{}),
		table_schema_generator.GenTableSchema(&storage.TableDigitaloceanStorageVolumesGenerator{}),
	}
}

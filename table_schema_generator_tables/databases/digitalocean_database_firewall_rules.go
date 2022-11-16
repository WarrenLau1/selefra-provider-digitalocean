package databases

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanDatabaseFirewallRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanDatabaseFirewallRulesGenerator{}

func (x *TableDigitaloceanDatabaseFirewallRulesGenerator) GetTableName() string {
	return "digitalocean_database_firewall_rules"
}

func (x *TableDigitaloceanDatabaseFirewallRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanDatabaseFirewallRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanDatabaseFirewallRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDigitaloceanDatabaseFirewallRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(godo.Database)
			svc := client.(*digitalocean_client.Client)
			getFunc := func() error {
				response, _, err := svc.Services.Databases.GetFirewallRules(ctx, p.ID)
				if err != nil {
					return err
				}
				resultChannel <- response
				return nil
			}

			err := digitalocean_client.ThrottleWrapper(ctx, svc, getFunc)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			return nil
		},
	}
}

func (x *TableDigitaloceanDatabaseFirewallRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanDatabaseFirewallRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("uuid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_uuid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterUUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("digitalocean_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to digitalocean_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableDigitaloceanDatabaseFirewallRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}

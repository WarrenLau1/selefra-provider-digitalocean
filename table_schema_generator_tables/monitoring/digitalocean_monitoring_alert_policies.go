package monitoring

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanMonitoringAlertPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanMonitoringAlertPoliciesGenerator{}

func (x *TableDigitaloceanMonitoringAlertPoliciesGenerator) GetTableName() string {
	return "digitalocean_monitoring_alert_policies"
}

func (x *TableDigitaloceanMonitoringAlertPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanMonitoringAlertPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanMonitoringAlertPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"uuid",
		},
	}
}

func (x *TableDigitaloceanMonitoringAlertPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)

			opt := &godo.ListOptions{
				PerPage: digitalocean_client.MaxItemsPerPage,
			}

			done := false
			listFunc := func() error {
				data, resp, err := svc.Services.Monitoring.ListAlertPolicies(ctx, opt)
				if err != nil {
					return err
				}

				resultChannel <- data

				if resp.Links == nil || resp.Links.IsLastPage() {
					done = true
					return nil
				}
				page, err := resp.Links.CurrentPage()
				if err != nil {
					return err
				}

				opt.Page = page + 1
				return nil
			}

			for !done {
				err := digitalocean_client.ThrottleWrapper(ctx, svc, listFunc)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
			}
			return nil
		},
	}
}

func (x *TableDigitaloceanMonitoringAlertPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanMonitoringAlertPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compare").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alerts").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uuid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeFloat).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("window").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entities").ColumnType(schema.ColumnTypeStringArray).Build(),
	}
}

func (x *TableDigitaloceanMonitoringAlertPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}

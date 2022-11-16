package balances

import (
	"context"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanBalancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanBalancesGenerator{}

func (x *TableDigitaloceanBalancesGenerator) GetTableName() string {
	return "digitalocean_balances"
}

func (x *TableDigitaloceanBalancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanBalancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanBalancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDigitaloceanBalancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)
			getFunc := func() error {
				response, _, err := svc.Services.Balance.Get(ctx)
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

func (x *TableDigitaloceanBalancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanBalancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("month_to_date_balance").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_balance").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("month_to_date_usage").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generated_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableDigitaloceanBalancesGenerator) GetSubTables() []*schema.Table {
	return nil
}

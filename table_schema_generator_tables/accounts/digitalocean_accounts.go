package accounts

import (
	"context"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanAccountsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanAccountsGenerator{}

func (x *TableDigitaloceanAccountsGenerator) GetTableName() string {
	return "digitalocean_accounts"
}

func (x *TableDigitaloceanAccountsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanAccountsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanAccountsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"uuid",
		},
	}
}

func (x *TableDigitaloceanAccountsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)
			getFunc := func() error {
				response, _, err := svc.Services.Account.Get(ctx)
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

func (x *TableDigitaloceanAccountsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanAccountsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("team").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uuid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reserved_ip_limit").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ReservedIPLimit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volume_limit").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_verified").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("droplet_limit").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("floating_ip_limit").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("FloatingIPLimit")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableDigitaloceanAccountsGenerator) GetSubTables() []*schema.Table {
	return nil
}

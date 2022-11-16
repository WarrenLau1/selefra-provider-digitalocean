package registries

import (
	"context"

	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanRegistriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanRegistriesGenerator{}

func (x *TableDigitaloceanRegistriesGenerator) GetTableName() string {
	return "digitalocean_registries"
}

func (x *TableDigitaloceanRegistriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanRegistriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanRegistriesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"name",
		},
	}
}

func (x *TableDigitaloceanRegistriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)
			getFunc := func() error {
				response, _, err := svc.Services.Registry.Get(ctx)
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

func (x *TableDigitaloceanRegistriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanRegistriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_usage_bytes").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_usage_bytes_updated_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableDigitaloceanRegistriesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableDigitaloceanRegistryRepositoriesGenerator{}),
	}
}

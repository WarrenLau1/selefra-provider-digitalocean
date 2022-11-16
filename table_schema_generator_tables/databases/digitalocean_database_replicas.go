package databases

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanDatabaseReplicasGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanDatabaseReplicasGenerator{}

func (x *TableDigitaloceanDatabaseReplicasGenerator) GetTableName() string {
	return "digitalocean_database_replicas"
}

func (x *TableDigitaloceanDatabaseReplicasGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanDatabaseReplicasGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanDatabaseReplicasGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDigitaloceanDatabaseReplicasGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(godo.Database)
			svc := client.(*digitalocean_client.Client)

			opt := &godo.ListOptions{
				PerPage: digitalocean_client.MaxItemsPerPage,
			}

			done := false
			listFunc := func() error {
				data, resp, err := svc.Services.Databases.ListReplicas(ctx, p.ID, opt)
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

func (x *TableDigitaloceanDatabaseReplicasGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanDatabaseReplicasGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_network_uuid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateNetworkUUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("digitalocean_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to digitalocean_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_connection").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableDigitaloceanDatabaseReplicasGenerator) GetSubTables() []*schema.Table {
	return nil
}

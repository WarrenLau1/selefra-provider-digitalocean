package firewalls

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanFirewallsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanFirewallsGenerator{}

func (x *TableDigitaloceanFirewallsGenerator) GetTableName() string {
	return "digitalocean_firewalls"
}

func (x *TableDigitaloceanFirewallsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanFirewallsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanFirewallsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableDigitaloceanFirewallsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)

			opt := &godo.ListOptions{
				PerPage: digitalocean_client.MaxItemsPerPage,
			}

			done := false
			listFunc := func() error {
				data, resp, err := svc.Services.Firewalls.List(ctx, opt)
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

func (x *TableDigitaloceanFirewallsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanFirewallsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inbound_rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outbound_rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_changes").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("droplet_ids").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("DropletIDs")).Build(),
	}
}

func (x *TableDigitaloceanFirewallsGenerator) GetSubTables() []*schema.Table {
	return nil
}

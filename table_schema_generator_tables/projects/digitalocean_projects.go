package projects

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanProjectsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanProjectsGenerator{}

func (x *TableDigitaloceanProjectsGenerator) GetTableName() string {
	return "digitalocean_projects"
}

func (x *TableDigitaloceanProjectsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanProjectsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanProjectsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableDigitaloceanProjectsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)

			opt := &godo.ListOptions{
				PerPage: digitalocean_client.MaxItemsPerPage,
			}

			done := false
			listFunc := func() error {
				data, resp, err := svc.Services.Projects.List(ctx, opt)
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

func (x *TableDigitaloceanProjectsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanProjectsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("owner_uuid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OwnerUUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purpose").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("OwnerID")).Build(),
	}
}

func (x *TableDigitaloceanProjectsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableDigitaloceanProjectResourcesGenerator{}),
	}
}

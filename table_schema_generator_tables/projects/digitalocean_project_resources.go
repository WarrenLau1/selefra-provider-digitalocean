package projects

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanProjectResourcesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanProjectResourcesGenerator{}

func (x *TableDigitaloceanProjectResourcesGenerator) GetTableName() string {
	return "digitalocean_project_resources"
}

func (x *TableDigitaloceanProjectResourcesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanProjectResourcesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanProjectResourcesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"urn",
		},
	}
}

func (x *TableDigitaloceanProjectResourcesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(godo.Project)
			svc := client.(*digitalocean_client.Client)

			opt := &godo.ListOptions{
				PerPage: digitalocean_client.MaxItemsPerPage,
			}

			done := false
			listFunc := func() error {
				data, resp, err := svc.Services.Projects.ListResources(ctx, p.ID, opt)
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

func (x *TableDigitaloceanProjectResourcesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanProjectResourcesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("links").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("digitalocean_projects_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to digitalocean_projects.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("urn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("URN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assigned_at").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableDigitaloceanProjectResourcesGenerator) GetSubTables() []*schema.Table {
	return nil
}

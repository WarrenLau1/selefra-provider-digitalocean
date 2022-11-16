package cdns

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanCdnsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanCdnsGenerator{}

func (x *TableDigitaloceanCdnsGenerator) GetTableName() string {
	return "digitalocean_cdns"
}

func (x *TableDigitaloceanCdnsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanCdnsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanCdnsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableDigitaloceanCdnsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)

			opt := &godo.ListOptions{
				PerPage: digitalocean_client.MaxItemsPerPage,
			}

			done := false
			listFunc := func() error {
				data, resp, err := svc.Services.Cdn.List(ctx, opt)
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

func (x *TableDigitaloceanCdnsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanCdnsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ttl").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("TTL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_domain").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("origin").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableDigitaloceanCdnsGenerator) GetSubTables() []*schema.Table {
	return nil
}

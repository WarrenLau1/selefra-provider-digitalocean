package storage

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanStorageVolumesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanStorageVolumesGenerator{}

func (x *TableDigitaloceanStorageVolumesGenerator) GetTableName() string {
	return "digitalocean_storage_volumes"
}

func (x *TableDigitaloceanStorageVolumesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanStorageVolumesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanStorageVolumesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableDigitaloceanStorageVolumesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)
			opt := &godo.ListVolumeParams{
				ListOptions: &godo.ListOptions{PerPage: digitalocean_client.MaxItemsPerPage},
			}

			done := false
			listFunc := func() error {
				data, resp, err := svc.Services.Storage.ListVolumes(ctx, opt)
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

				opt.ListOptions.Page = page + 1
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

func (x *TableDigitaloceanStorageVolumesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanStorageVolumesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("droplet_ids").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("DropletIDs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_gigabytes").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("SizeGigaBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filesystem_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filesystem_label").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableDigitaloceanStorageVolumesGenerator) GetSubTables() []*schema.Table {
	return nil
}

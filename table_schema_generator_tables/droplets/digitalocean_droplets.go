package droplets

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanDropletsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanDropletsGenerator{}

func (x *TableDigitaloceanDropletsGenerator) GetTableName() string {
	return "digitalocean_droplets"
}

func (x *TableDigitaloceanDropletsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanDropletsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanDropletsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableDigitaloceanDropletsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)

			opt := &godo.ListOptions{
				PerPage: digitalocean_client.MaxItemsPerPage,
			}

			done := false
			listFunc := func() error {
				data, resp, err := svc.Services.Droplets.List(ctx, opt)
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

func (x *TableDigitaloceanDropletsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanDropletsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vcpus").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_slug").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kernel").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("networks").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_uuid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VPCUUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volume_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("VolumeIDs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_backup_window").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locked").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("features").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_ids").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("BackupIDs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_ids").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("SnapshotIDs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableDigitaloceanDropletsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableDigitaloceanDropletNeighborsGenerator{}),
	}
}

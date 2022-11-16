package droplets

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanDropletNeighborsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanDropletNeighborsGenerator{}

func (x *TableDigitaloceanDropletNeighborsGenerator) GetTableName() string {
	return "digitalocean_droplet_neighbors"
}

func (x *TableDigitaloceanDropletNeighborsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanDropletNeighborsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanDropletNeighborsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"neighbor_id",
		},
	}
}

func (x *TableDigitaloceanDropletNeighborsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*digitalocean_client.Client)
			droplet := task.ParentRawResult.(godo.Droplet)

			neighbors, _, err := svc.Services.Droplets.Neighbors(ctx, droplet.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if neighbors == nil {
				return nil
			}
			nn := make([]NeighborWrapper, len(neighbors))
			for i, n := range neighbors {
				nn[i] = NeighborWrapper{
					DropletId:  droplet.ID,
					NeighborId: n.ID,
				}
			}
			resultChannel <- nn
			return nil
		},
	}
}

type NeighborWrapper struct {
	DropletId  int
	NeighborId int
}

func (x *TableDigitaloceanDropletNeighborsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanDropletNeighborsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("neighbor_id").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("droplet_id").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("digitalocean_droplets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to digitalocean_droplets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableDigitaloceanDropletNeighborsGenerator) GetSubTables() []*schema.Table {
	return nil
}

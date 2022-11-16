package spaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanSpacesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanSpacesGenerator{}

func (x *TableDigitaloceanSpacesGenerator) GetTableName() string {
	return "digitalocean_spaces"
}

func (x *TableDigitaloceanSpacesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanSpacesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanSpacesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableDigitaloceanSpacesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*digitalocean_client.Client)

			buckets, err := c.Services.Spaces.ListBuckets(ctx, &s3.ListBucketsInput{}, func(options *s3.Options) {
				options.Region = c.SpacesRegion
			})
			if err != nil {
				if !c.CredentialStatus.Spaces {

					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			wb := make([]*WrappedBucket, len(buckets.Buckets))
			for i, b := range buckets.Buckets {
				wb[i] = &WrappedBucket{
					Bucket:   b,
					Location: c.SpacesRegion,
				}
			}
			resultChannel <- wb
			return nil
		},
	}
}

func (x *TableDigitaloceanSpacesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return digitalocean_client.ExpandByRegion()
}

func (x *TableDigitaloceanSpacesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("acls").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ACLs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableDigitaloceanSpacesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableDigitaloceanSpaceCorsGenerator{}),
	}
}

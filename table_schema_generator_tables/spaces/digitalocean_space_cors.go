package spaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
	"github.com/selefra/selefra-provider-digitalocean/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableDigitaloceanSpaceCorsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableDigitaloceanSpaceCorsGenerator{}

func (x *TableDigitaloceanSpaceCorsGenerator) GetTableName() string {
	return "digitalocean_space_cors"
}

func (x *TableDigitaloceanSpaceCorsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableDigitaloceanSpaceCorsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableDigitaloceanSpaceCorsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableDigitaloceanSpaceCorsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var ae smithy.APIError
			r := task.ParentRawResult.(*WrappedBucket)
			svc := client.(*digitalocean_client.Client).Services
			corsOutput, err := svc.Spaces.GetBucketCors(ctx, &s3.GetBucketCorsInput{Bucket: r.Name}, func(options *s3.Options) {
				options.Region = r.Location
			})
			if err != nil && !(errors.As(err, &ae) && ae.ErrorCode() == "NoSuchCORSConfiguration") {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if corsOutput != nil {
				resultChannel <- corsOutput.CORSRules
			}
			return nil
		},
	}
}

type WrappedBucket struct {
	types.Bucket
	Location string
	Public   bool
	ACLs     []types.Grant
}

func (x *TableDigitaloceanSpaceCorsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableDigitaloceanSpaceCorsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("max_age_seconds").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("digitalocean_spaces_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to digitalocean_spaces.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_methods").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_origins").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_headers").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expose_headers").ColumnType(schema.ColumnTypeStringArray).Build(),
	}
}

func (x *TableDigitaloceanSpaceCorsGenerator) GetSubTables() []*schema.Table {
	return nil
}

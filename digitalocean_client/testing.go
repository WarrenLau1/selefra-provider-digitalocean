package digitalocean_client

import (
	"context"
	"testing"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/test_helper"
	"github.com/spf13/viper"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, createService func(t *testing.T, ctrl *gomock.Controller) Services, _ TestOptions) {
	testProvider := newTestProvider(t, table, createService)
	config := constants.Testtest
	test_helper.RunProviderPullTables(testProvider, config, constants.Constants_15, constants.Constants_16)
}

func newTestProvider(t *testing.T, table *schema.Table, createService func(t *testing.T, ctrl *gomock.Controller) Services) *provider.Provider {
	return &provider.Provider{
		Name:      constants.Digitalocean,
		Version:   constants.V,
		TableList: []*schema.Table{table},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				services := createService(t, gomock.NewController(t))
				client := &Client{
					SpacesRegion: constants.Nyc,
					Services:     &services,
				}
				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return ``
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var digitaloceanConfig Config
				err := config.Unmarshal(&digitaloceanConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_17,
				constants.NA,
				constants.Notsupported,
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}

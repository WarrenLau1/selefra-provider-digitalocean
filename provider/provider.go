package provider

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"

	"github.com/selefra/selefra-provider-digitalocean/constants"
	"github.com/selefra/selefra-provider-digitalocean/digitalocean_client"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      constants.Digitalocean,
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var digitaloceanConfig digitalocean_client.Config
				err := config.Unmarshal(&digitaloceanConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}

				clients, err := digitalocean_client.NewClients(digitaloceanConfig)

				if err != nil {
					clientMeta.ErrorF(constants.Newclientserrs, err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg(constants.Accountinformationnotfound)
				}
				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `##  Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
#accounts:
#  - token: "<YOUR_API_TOKEN_HERE>" # env DIGITALOCEAN_ACCESS_TOKEN or DIGITALOCEAN_TOKEN
#    spaces_regions: ["region"]
#    spaces_access_key: "<YOUR_SPACES_ACCESS_KEY>" # env SPACES_SECRET_ACCESS_KEY
#    spaces_access_key_id: "<YOUR_SPACES_ACCESS_KEY_ID>" # env SPACES_ACCESS_KEY_ID`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var digitaloceanConfig digitalocean_client.Config
				err := config.Unmarshal(&digitaloceanConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg(constants.Analysisconfigerrs, err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				constants.Constants_18,
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

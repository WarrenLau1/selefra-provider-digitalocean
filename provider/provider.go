package provider

import (
	"context"
	"os"
	"strings"

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

				if digitaloceanConfig.Token == "" {
					digitaloceanConfig.Token = os.Getenv("DIGITALOCEAN_TOKEN")
				}

				if digitaloceanConfig.Token == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing Token in configuration")
				}

				if len(digitaloceanConfig.SpacesRegions) == 0 {
					regionsData := os.Getenv("DIGITALOCEAN_SPACES_REGIONS")

					var regionsList []string

					if regionsData != "" {
						regionsList = strings.Split(regionsData, ",")
					}

					digitaloceanConfig.SpacesRegions = regionsList
				}

				if len(digitaloceanConfig.SpacesRegions) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing SpacesRegions in configuration")
				}

				if digitaloceanConfig.SpacesAccessKey == "" {
					digitaloceanConfig.SpacesAccessKey = os.Getenv("DIGITALOCEAN_SPACES_ACCESS_KEY")
				}

				if digitaloceanConfig.SpacesAccessKey == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing SpacesAccessKey in configuration")
				}

				if digitaloceanConfig.SpacesAccessKeyId == "" {
					digitaloceanConfig.SpacesAccessKeyId = os.Getenv("DIGITALOCEAN_SPACES_ACCESS_KEY_ID")
				}

				if digitaloceanConfig.SpacesAccessKeyId == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing SpacesAccessKeyId in configuration")
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

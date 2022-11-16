package digitalocean_client

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func ExpandByRegion() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		cli := client.(*Client)
		if len(cli.Regions) == 0 {
			return []*schema.ClientTaskContext{&schema.ClientTaskContext{
				Client: cli,
				Task:   task.Clone(),
			}}
		}
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)

		for _, region := range cli.Regions {
			clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
				Client: cli.WithSpacesRegion(region),
				Task:   task.Clone(),
			})
		}
		return clientTaskContextSlice
	}
}

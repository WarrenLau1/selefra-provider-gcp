package gcp_client

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func ExpandByProjects() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		cli := client.(*Client)
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)

		for _, projectId := range cli.projects {
			clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
				Client: cli.withProject(projectId),
				Task:   task.Clone(),
			})
		}
		return clientTaskContextSlice
	}
}

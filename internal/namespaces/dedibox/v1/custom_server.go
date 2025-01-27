package dedibox

import (
	"context"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/dedibox/v1"
)

func serverInstallBuilder(c *core.Command) *core.Command {
	c.AddInterceptors(func(ctx context.Context, argsI interface{}, runner core.CommandRunner) (interface{}, error) {
		args := argsI.(*dedibox.InstallServerRequest)
		if args.Partitions == nil {
			client := core.ExtractClient(ctx)
			api := dedibox.NewAPI(client)

			getDefaultPartition, err := api.GetServerDefaultPartitioning(&dedibox.GetServerDefaultPartitioningRequest{
				Zone:     args.Zone,
				ServerID: args.ServerID,
				OsID:     args.OsID,
			})
			if err != nil {
				return nil, err
			}
			for _, partitions := range getDefaultPartition.Partitions {
				InstallPartition := dedibox.InstallPartition{
					FileSystem: partitions.FileSystem,
					MountPoint: partitions.MountPoint,
					RaidLevel:  partitions.RaidLevel,
					Capacity:   partitions.Capacity,
					Connectors: partitions.Connectors,
				}
				args.Partitions = append(args.Partitions, &InstallPartition)
			}
		} else {
			for _, partitions := range args.Partitions {
				partitions.Capacity = partitions.Capacity / 1000000
			}
		}
		return runner(ctx, args)
	})
	return c
}

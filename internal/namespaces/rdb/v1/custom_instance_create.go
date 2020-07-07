package rdb

import (
	"context"
	"strings"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func instanceCreateBuilder(c *core.Command) *core.Command {
	c.ArgSpecs.GetByName("node-type").Default = core.DefaultValueSetter("DB-DEV-S")
	c.ArgSpecs.GetByName("node-type").EnumValues = nodeTypes

	c.WaitFunc = func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
		api := rdb.NewAPI(core.ExtractClient(ctx))
		return api.WaitForInstance(&rdb.WaitForInstanceRequest{
			InstanceID:    respI.(*rdb.Instance).ID,
			Region:        respI.(*rdb.Instance).Region,
			Timeout:       scw.TimeDurationPtr(instanceActionTimeout),
			RetryInterval: core.DefaultRetryInterval,
		})
	}

	// Waiting for API to accept uppercase node-type
	c.Interceptor = func(ctx context.Context, argsI interface{}, runner core.CommandRunner) (interface{}, error) {
		args := argsI.(*rdb.CreateInstanceRequest)
		args.NodeType = strings.ToLower(args.NodeType)
		return runner(ctx, args)
	}

	return c
}

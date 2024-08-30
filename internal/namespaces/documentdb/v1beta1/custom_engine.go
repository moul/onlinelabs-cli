package documentdb

import (
	"context"
	"time"

	"github.com/scaleway/scaleway-cli/v2/core"
	documentdb "github.com/scaleway/scaleway-sdk-go/api/documentdb/v1beta1"
)

func engineListBuilder(c *core.Command) *core.Command {
	type customEngine struct {
		Name       string     `json:"name"`
		EngineType string     `json:"engine_type"`
		EndOfLife  *time.Time `json:"end_of_life"`
	}

	c.View = &core.View{
		Fields: []*core.ViewField{
			{
				Label:     "Name",
				FieldName: "Name",
			},
			{
				Label:     "Engine Type",
				FieldName: "EngineType",
			},
			{
				Label:     "End of Life",
				FieldName: "EndOfLife",
			},
		},
	}

	c.AddInterceptors(func(ctx context.Context, argsI interface{}, runner core.CommandRunner) (i interface{}, err error) {
		listEngineResp, err := runner(ctx, argsI)
		if err != nil {
			return listEngineResp, err
		}
		engineList := listEngineResp.([]*documentdb.DatabaseEngine)
		var res []customEngine
		for _, engine := range engineList {
			for _, version := range engine.Versions {
				res = append(res, customEngine{
					Name:       version.Name,
					EngineType: engine.Name,
					EndOfLife:  version.EndOfLife,
				})
			}
		}

		return res, nil
	})

	return c
}

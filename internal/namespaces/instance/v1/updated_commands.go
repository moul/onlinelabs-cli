package instance

import (
	"context"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
)

func updateCommands(commands *core.Commands) {
	updateInstancePlacementGroupGet(commands.MustFind("instance", "placement-group", "get"))
}

func updateInstancePlacementGroupGet(c *core.Command) {
	c.Run = func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
		req := argsI.(*instance.GetPlacementGroupRequest)

		client := core.ExtractClient(ctx)
		api := instance.NewAPI(client)
		placementGroupResponse, err := api.GetPlacementGroup(req)
		if err != nil {
			return nil, err
		}

		placementGroupServersResponse, err := api.GetPlacementGroupServers(&instance.GetPlacementGroupServersRequest{
			PlacementGroupID: req.PlacementGroupID,
		})
		if err != nil {
			return nil, err
		}

		return &struct {
			*instance.PlacementGroup
			Servers []*instance.PlacementGroupServer
		}{
			placementGroupResponse.PlacementGroup,
			placementGroupServersResponse.Servers,
		}, nil
	}

	c.View = &core.View{
		Sections: []*core.ViewSection{
			{FieldName: "PlacementGroup", Title: "Placement Group"},
			{FieldName: "servers", Title: "Servers"},
		},
	}
}

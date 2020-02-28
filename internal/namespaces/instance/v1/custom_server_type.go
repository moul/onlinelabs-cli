package instance

import (
	"context"
	"sort"
	"strings"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

//
// Builders
//

// serverTypeListBuilder transforms the server map into a list to display a
// table of server types instead of a flat key/value list.
func serverTypeListBuilder(c *core.Command) *core.Command {
	originalRun := c.Run

	c.Run = func(ctx context.Context, argsI interface{}) (interface{}, error) {
		type customServerType struct {
			Name            string        `json:"name"`
			MonthlyPrice    *scw.Money    `json:"monthly_price"`
			HourlyPrice     *scw.Money    `json:"hourly_price"`
			LocalVolumeSize scw.Size      `json:"local_volume_size"`
			CPU             uint32        `json:"cpu"`
			GPU             *uint64       `json:"gpu"`
			RAM             scw.Size      `json:"ram"`
			Arch            instance.Arch `json:"arch"`
		}

		originalRes, err := originalRun(ctx, argsI)
		if err != nil {
			return nil, err
		}

		listServersTypesResponse := originalRes.(*instance.ListServersTypesResponse)
		serverTypes := []*customServerType(nil)

		for name, serverType := range listServersTypesResponse.Servers {
			serverTypes = append(serverTypes, &customServerType{
				Name:            name,
				MonthlyPrice:    scw.NewMoneyFromFloat(float64(serverType.MonthlyPrice), "EUR", 2),
				HourlyPrice:     scw.NewMoneyFromFloat(float64(serverType.HourlyPrice), "EUR", 3),
				LocalVolumeSize: serverType.VolumesConstraint.MinSize,
				CPU:             serverType.Ncpus,
				GPU:             serverType.Gpu,
				RAM:             scw.Size(serverType.RAM),
				Arch:            serverType.Arch,
			})
		}

		sort.Slice(serverTypes, func(i, j int) bool {
			categoryA := serverTypeCategory(serverTypes[i].Name)
			categoryB := serverTypeCategory(serverTypes[j].Name)
			if categoryA != categoryB {
				return categoryA < categoryB
			}
			return serverTypes[i].MonthlyPrice.ToFloat() < serverTypes[j].MonthlyPrice.ToFloat()
		})

		return serverTypes, nil
	}

	return c
}

func serverTypeCategory(serverTypeName string) (category string) {
	return strings.Split(serverTypeName, "-")[0]
}

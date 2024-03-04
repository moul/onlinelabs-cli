package flexibleip

import (
	"context"
	"time"

	"github.com/fatih/color"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-cli/v2/internal/human"
	flexibleip "github.com/scaleway/scaleway-sdk-go/api/flexibleip/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

var (
	ipStatusMarshalSpecs = human.EnumMarshalSpecs{
		flexibleip.FlexibleIPStatusAttached:  &human.EnumMarshalSpec{Attribute: color.FgGreen},
		flexibleip.FlexibleIPStatusDetaching: &human.EnumMarshalSpec{Attribute: color.FgBlue},
		flexibleip.FlexibleIPStatusError:     &human.EnumMarshalSpec{Attribute: color.FgRed},
		flexibleip.FlexibleIPStatusLocked:    &human.EnumMarshalSpec{Attribute: color.FgRed},
		flexibleip.FlexibleIPStatusReady:     &human.EnumMarshalSpec{Attribute: color.FgGreen},
		flexibleip.FlexibleIPStatusUnknown:   &human.EnumMarshalSpec{Attribute: color.Faint},
		flexibleip.FlexibleIPStatusUpdating:  &human.EnumMarshalSpec{Attribute: color.FgBlue},
	}
)

const (
	flexibleIPTimeout = 60 * time.Second
)

func createIPBuilder(c *core.Command) *core.Command {
	c.WaitFunc = func(ctx context.Context, _, respI interface{}) (interface{}, error) {
		getResp := respI.(*flexibleip.FlexibleIP)
		api := flexibleip.NewAPI(core.ExtractClient(ctx))
		return api.WaitForFlexibleIP(&flexibleip.WaitForFlexibleIPRequest{
			FipID:         getResp.ID,
			Zone:          getResp.Zone,
			Timeout:       scw.TimeDurationPtr(flexibleIPTimeout),
			RetryInterval: core.DefaultRetryInterval,
		})
	}
	return c
}

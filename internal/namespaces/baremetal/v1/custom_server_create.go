package baremetal

import (
	"context"
	"fmt"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	baremetal "github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func serverCreateBuilder(c *core.Command) *core.Command {
	type baremetalCreateServerRequestCustom struct {
		Zone scw.Zone `json:"-"`
		// OrganizationID with which the server will be created
		OrganizationID *string `json:"organization_id"`
		// ProjectID with which the server will be created
		ProjectID *string `json:"project_id"`
		// Name of the server (≠hostname)
		Name string `json:"name"`
		// Description associated to the server, max 255 characters
		Description string `json:"description"`
		// Tags associated with the server
		Tags []string `json:"tags"`
		// Type of the server
		Type string
		// Installation configuration
		Install *baremetal.CreateServerRequestInstall
	}

	c.ArgsType = reflect.TypeOf(baremetalCreateServerRequestCustom{})

	c.ArgSpecs.DeleteByName("offer-id")

	c.ArgSpecs.GetByName("name").Default = core.RandomValueGenerator("bm")
	c.ArgSpecs.GetByName("description").Required = false

	c.ArgSpecs.AddBefore("tags.{index}", &core.ArgSpec{
		Name:  "type",
		Short: "Server commercial type",
	})

	c.Run = func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
		client := core.ExtractClient(ctx)
		api := baremetal.NewAPI(client)

		bareMetalLink := "\u001B[38;2;121;45;212;4m\u001B]8;;https://console.scaleway.com/organization/contracts\u001B\\Bare Metal Specific Conditions\u001B]8;;\u001B\\\u001B[0m (\u001B[38;2;121;45;212mhttps://console.scaleway.com/organization/contracts\u001B[0m)"
		appleLink := "\u001B[4m\u001B[38;2;121;45;212m\u001B]8;;https://www.apple.com/legal/sla/\u001B\\macOS License Agreement\u001B]8;;\u001B\\\u001B[0m (https://www.apple.com/legal/sla/)"

		fmt.Println("Please note: Signing the " + bareMetalLink + " and the " + appleLink + " is mandatory.")

		tmpRequest := argsI.(*baremetalCreateServerRequestCustom)
		request := &baremetal.CreateServerRequest{
			Zone:           tmpRequest.Zone,
			OrganizationID: tmpRequest.OrganizationID,
			ProjectID:      tmpRequest.ProjectID,
			Name:           tmpRequest.Name,
			Description:    tmpRequest.Description,
			Tags:           tmpRequest.Tags,
		}

		if tmpRequest.Install != nil {
			request.Install = tmpRequest.Install
		}

		// We need to find the offer ID.
		// While baremetal does not have list offer name filter we are forced to iterate
		// on the list of offers provided.
		offer, err := api.GetOfferByName(&baremetal.GetOfferByNameRequest{
			OfferName: tmpRequest.Type,
			Zone:      tmpRequest.Zone,
		})
		if err != nil {
			return nil, err
		}
		if offer == nil {
			return nil, fmt.Errorf("could not match an offer with the type: %s", tmpRequest.Type)
		}
		request.OfferID = offer.ID

		return api.CreateServer(request)
	}

	c.SeeAlsos = []*core.SeeAlso{
		{
			Short:   "List os",
			Command: "scw baremetal os list",
		},
		{
			Short:   "Install an OS on your server",
			Command: "scw baremetal server install",
		},
	}

	c.Examples = []*core.Example{
		{
			Short:    "Create instance",
			ArgsJSON: `{}`,
		},
	}

	c.WaitFunc = func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
		api := baremetal.NewAPI(core.ExtractClient(ctx))
		return api.WaitForServer(&baremetal.WaitForServerRequest{
			Zone:          argsI.(*baremetalCreateServerRequestCustom).Zone,
			ServerID:      respI.(*baremetal.Server).ID,
			Timeout:       scw.TimeDurationPtr(ServerActionTimeout),
			RetryInterval: core.DefaultRetryInterval,
		})
	}

	return c
}

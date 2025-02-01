package applesilicon

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/core/human"
	applesilicon "github.com/scaleway/scaleway-sdk-go/api/applesilicon/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	serverActionTimeout = 60 * time.Minute
)

const (
	serverActionCreate = iota
	serverActionDelete
	serverActionReboot
)

var serverStatusMarshalSpecs = human.EnumMarshalSpecs{
	applesilicon.ServerStatusError:     &human.EnumMarshalSpec{Attribute: color.FgRed, Value: "error"},
	applesilicon.ServerStatusReady:     &human.EnumMarshalSpec{Attribute: color.FgGreen, Value: "ready"},
	applesilicon.ServerStatusRebooting: &human.EnumMarshalSpec{Attribute: color.FgBlue, Value: "rebooting"},
	applesilicon.ServerStatusStarting:  &human.EnumMarshalSpec{Attribute: color.FgBlue, Value: "starting"},
	applesilicon.ServerStatusUpdating:  &human.EnumMarshalSpec{Attribute: color.FgBlue, Value: "updating"},
}

func serverCreateBuilder(c *core.Command) *core.Command {
	c.ArgSpecs.GetByName("type").AutoCompleteFunc = autocompleteServerType

	c.AddInterceptors(func(ctx context.Context, argsI interface{}, runner core.CommandRunner) (i interface{}, err error) {
		args := argsI.(*applesilicon.CreateServerRequest)
		bareMetalLink := "\u001B[38;2;121;45;212;4m\u001B]8;;https://console.scaleway.com/organization/contracts\u001B\\Bare Metal Specific Conditions\u001B]8;;\u001B\\\u001B[0m (\u001B[38;2;121;45;212mhttps://console.scaleway.com/organization/contracts\u001B[0m)"
		appleLink := "\u001B[4m\u001B[38;2;121;45;212m\u001B]8;;https://www.apple.com/legal/sla/\u001B\\macOS License Agreement\u001B]8;;\u001B\\\u001B[0m (https://www.apple.com/legal/sla/)"

		fmt.Println("Please note: Signing the " + bareMetalLink + " and the " + appleLink + " is mandatory.")
		return runner(ctx, args)
	},
	)

	c.WaitFunc = waitForServerFunc(serverActionCreate)
	return c
}

func serverDeleteBuilder(c *core.Command) *core.Command {
	c.WaitFunc = waitForServerFunc(serverActionDelete)
	return c
}

func serverRebootBuilder(c *core.Command) *core.Command {
	c.WaitFunc = waitForServerFunc(serverActionReboot)
	return c
}

func waitForServerFunc(action int) core.WaitFunc {
	return func(ctx context.Context, _, respI interface{}) (interface{}, error) {
		server, err := applesilicon.NewAPI(core.ExtractClient(ctx)).WaitForServer(&applesilicon.WaitForServerRequest{
			Zone:          respI.(*applesilicon.Server).Zone,
			ServerID:      respI.(*applesilicon.Server).ID,
			Timeout:       scw.TimeDurationPtr(serverActionTimeout),
			RetryInterval: core.DefaultRetryInterval,
		})

		switch action {
		case serverActionCreate:
			return server, err
		case serverActionReboot:
			return server, err
		case serverActionDelete:
			if err != nil {
				// if we get a 404 here, it means the resource was successfully deleted
				notFoundError := &scw.ResourceNotFoundError{}
				responseError := &scw.ResponseError{}
				if errors.As(err, &responseError) && responseError.StatusCode == http.StatusNotFound || errors.As(err, &notFoundError) {
					return fmt.Sprintf("Server %s successfully deleted.", respI.(*applesilicon.Server).ID), nil
				}
			}
		}
		return nil, err
	}
}

func serverWaitCommand() *core.Command {
	type customServerWaitArgs struct {
		applesilicon.WaitForServerRequest
	}

	return &core.Command{
		Short:     `Wait for a server to reach a stable state`,
		Long:      `Wait for server to reach a stable state. This is similar to using --wait flag on other action commands, but without requiring a new action on the server.`,
		Namespace: "apple-silicon",
		Resource:  "server",
		Verb:      "wait",
		Groups:    []string{"workflow"},
		ArgsType:  reflect.TypeOf(customServerWaitArgs{}),
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, err error) {
			args := argsI.(*customServerWaitArgs)

			api := applesilicon.NewAPI(core.ExtractClient(ctx))
			cluster, err := api.WaitForServer(&applesilicon.WaitForServerRequest{
				Zone:          args.Zone,
				ServerID:      args.ServerID,
				Timeout:       args.Timeout,
				RetryInterval: core.DefaultRetryInterval,
			})
			if err != nil {
				return nil, err
			}

			return cluster, nil
		},
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server.`,
				Required:   true,
				Positional: true,
			},
			core.ZoneArgSpec(),
			core.WaitTimeoutArgSpec(serverActionTimeout),
		},
		Examples: []*core.Example{
			{
				Short:    "Wait for a server to reach a stable state",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

var completeListTypeServerCache *applesilicon.ListServerTypesResponse

func autocompleteServerType(ctx context.Context, prefix string, _ any) core.AutocompleteSuggestions {
	suggestions := core.AutocompleteSuggestions(nil)

	client := core.ExtractClient(ctx)
	api := applesilicon.NewAPI(client)

	if completeListTypeServerCache == nil {
		res, err := api.ListServerTypes(&applesilicon.ListServerTypesRequest{})
		if err != nil {
			return nil
		}
		completeListTypeServerCache = res
	}

	for _, serverType := range completeListTypeServerCache.ServerTypes {
		if strings.HasPrefix(serverType.Name, prefix) {
			suggestions = append(suggestions, serverType.Name)
		}
	}
	return suggestions
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package baremetal

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/baremetal/v3"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		baremetalRoot(),
		baremetalPrivateNetwork(),
		baremetalPrivateNetworkAdd(),
		baremetalPrivateNetworkSet(),
		baremetalPrivateNetworkList(),
		baremetalPrivateNetworkDelete(),
	)
}
func baremetalRoot() *core.Command {
	return &core.Command{
		Short:     `Elastic Metal - Private Networks API`,
		Long:      `Elastic Metal - Private Networks API.`,
		Namespace: "baremetal",
	}
}

func baremetalPrivateNetwork() *core.Command {
	return &core.Command{
		Short: `Private network management command`,
		Long: `A Private Network allows you to interconnect your resources
in an isolated and private
network. Network reachability is limited to the
resources that are on the same Private Network.  A VLAN
interface is available on the server and can be freely
managed (adding IP addresses, shutdown interface etc.).

Note that a resource can be a part of multiple Private Networks.`,
		Namespace: "baremetal",
		Resource:  "private-network",
	}
}

func baremetalPrivateNetworkAdd() *core.Command {
	return &core.Command{
		Short:     `Add a server to a Private Network`,
		Long:      `Add an Elastic Metal server to a Private Network.`,
		Namespace: "baremetal",
		Resource:  "private-network",
		Verb:      "add",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.PrivateNetworkAPIAddServerPrivateNetworkRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `UUID of the Private Network`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ipam-ip-ids.{index}",
				Short:      `IPAM IDs of an IPs to attach to the server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.PrivateNetworkAPIAddServerPrivateNetworkRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewPrivateNetworkAPI(client)
			return api.AddServerPrivateNetwork(request)

		},
	}
}

func baremetalPrivateNetworkSet() *core.Command {
	return &core.Command{
		Short:     `Set multiple Private Networks on a server`,
		Long:      `Configure multiple Private Networks on an Elastic Metal server.`,
		Namespace: "baremetal",
		Resource:  "private-network",
		Verb:      "set",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.PrivateNetworkAPISetServerPrivateNetworksRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "per-private-network-ipam-ip-ids.{key}",
				Short:      `Object where the keys are the UUIDs of Private Networks and the values are arrays of IPAM IDs representing the IPs to assign to this Elastic Metal server on the Private Network. If the array supplied for a Private Network is empty, the next available IP from the Private Network's CIDR block will automatically be used for attachment.`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.PrivateNetworkAPISetServerPrivateNetworksRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewPrivateNetworkAPI(client)
			return api.SetServerPrivateNetworks(request)

		},
	}
}

func baremetalPrivateNetworkList() *core.Command {
	return &core.Command{
		Short:     `List the Private Networks of a server`,
		Long:      `List the Private Networks of an Elastic Metal server.`,
		Namespace: "baremetal",
		Resource:  "private-network",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.PrivateNetworkAPIListServerPrivateNetworksRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order for the returned Private Networks`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc"},
			},
			{
				Name:       "server-id",
				Short:      `Filter Private Networks by server UUID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `Filter Private Networks by Private Network UUID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Filter Private Networks by project UUID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ipam-ip-ids.{index}",
				Short:      `Filter Private Networks by IPAM IP UUIDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Filter Private Networks by organization UUID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.PrivateNetworkAPIListServerPrivateNetworksRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewPrivateNetworkAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListServerPrivateNetworks(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.ServerPrivateNetworks, nil

		},
	}
}

func baremetalPrivateNetworkDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Private Network`,
		Long:      `Delete a Private Network.`,
		Namespace: "baremetal",
		Resource:  "private-network",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.PrivateNetworkAPIDeleteServerPrivateNetworkRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `UUID of the Private Network`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.PrivateNetworkAPIDeleteServerPrivateNetworkRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewPrivateNetworkAPI(client)
			e = api.DeleteServerPrivateNetwork(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "private-network",
				Verb:     "delete",
			}, nil
		},
	}
}

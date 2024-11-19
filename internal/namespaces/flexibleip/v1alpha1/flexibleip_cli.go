// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package flexibleip

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/flexibleip/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		fipRoot(),
		fipIP(),
		fipMac(),
		fipIPCreate(),
		fipIPGet(),
		fipIPList(),
		fipIPUpdate(),
		fipIPDelete(),
		fipIPAttach(),
		fipIPDetach(),
		fipMacCreate(),
		fipMacDuplicate(),
		fipMacMove(),
		fipMacDelete(),
	)
}
func fipRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Elastic Metal servers' flexible public IP addresses`,
		Long:      `This API allows you to manage your Elastic Metal servers' flexible public IP addresses.`,
		Namespace: "fip",
	}
}

func fipIP() *core.Command {
	return &core.Command{
		Short: `Flexible IP management commands`,
		Long: `A flexible IP can be attached to any Elastic Metal server within the same zone.
Multiple flexible IPs can be attached to a server.`,
		Namespace: "fip",
		Resource:  "ip",
	}
}

func fipMac() *core.Command {
	return &core.Command{
		Short:     `MAC address management commands`,
		Long:      `MAC address management commands.`,
		Namespace: "fip",
		Resource:  "mac",
	}
}

func fipIPCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new flexible IP`,
		Long:      `Generate a new flexible IP within a given zone, specifying its configuration including Project ID and description.`,
		Namespace: "fip",
		Resource:  "ip",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.CreateFlexibleIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "description",
				Short:      `Flexible IP description (max. of 255 characters)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags to associate to the flexible IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "server-id",
				Short:      `ID of the server to which the newly created flexible IP will be attached.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "reverse",
				Short:      `Value of the reverse DNS`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "is-ipv6",
				Short:      `Defines whether the flexible IP has an IPv6 address.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.CreateFlexibleIPRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			return api.CreateFlexibleIP(request)

		},
	}
}

func fipIPGet() *core.Command {
	return &core.Command{
		Short:     `Get an existing flexible IP`,
		Long:      `Retrieve information about an existing flexible IP, specified by its ID and zone. Its full details, including Project ID, description and status, are returned in the response object.`,
		Namespace: "fip",
		Resource:  "ip",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.GetFlexibleIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fip-id",
				Short:      `ID of the flexible IP`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.GetFlexibleIPRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			return api.GetFlexibleIP(request)

		},
	}
}

func fipIPList() *core.Command {
	return &core.Command{
		Short:     `List flexible IPs`,
		Long:      `List all flexible IPs within a given zone.`,
		Namespace: "fip",
		Resource:  "ip",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.ListFlexibleIPsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order of the returned flexible IPs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "tags.{index}",
				Short:      `Filter by tag, only flexible IPs with one or more matching tags will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "status.{index}",
				Short:      `Filter by status, only flexible IPs with this status will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "ready", "updating", "attached", "error", "detaching", "locked"},
			},
			{
				Name:       "server-ids.{index}",
				Short:      `Filter by server IDs, only flexible IPs with these server IDs will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Filter by Project ID, only flexible IPs from this Project will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Filter by Organization ID, only flexible IPs from this Organization will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.ListFlexibleIPsRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListFlexibleIPs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.FlexibleIPs, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "IPAddress",
			},
			{
				FieldName: "Status",
			},
			{
				FieldName: "Reverse",
			},
			{
				FieldName: "ServerID",
			},
			{
				FieldName: "Description",
			},
			{
				FieldName: "Tags",
			},
			{
				FieldName: "ProjectID",
			},
			{
				FieldName: "OrganizationID",
			},
			{
				FieldName: "UpdatedAt",
			},
			{
				FieldName: "CreatedAt",
			},
		}},
	}
}

func fipIPUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing flexible IP`,
		Long:      `Update the parameters of an existing flexible IP, specified by its ID and zone. These parameters include tags and description.`,
		Namespace: "fip",
		Resource:  "ip",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.UpdateFlexibleIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fip-id",
				Short:      `ID of the flexible IP to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "description",
				Short:      `Flexible IP description (max. 255 characters)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags associated with the flexible IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "reverse",
				Short:      `Value of the reverse DNS`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.UpdateFlexibleIPRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			return api.UpdateFlexibleIP(request)

		},
	}
}

func fipIPDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an existing flexible IP`,
		Long:      `Delete an existing flexible IP, specified by its ID and zone. Note that deleting a flexible IP is permanent and cannot be undone.`,
		Namespace: "fip",
		Resource:  "ip",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.DeleteFlexibleIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fip-id",
				Short:      `ID of the flexible IP to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.DeleteFlexibleIPRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			e = api.DeleteFlexibleIP(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "ip",
				Verb:     "delete",
			}, nil
		},
	}
}

func fipIPAttach() *core.Command {
	return &core.Command{
		Short:     `Attach an existing flexible IP to a server`,
		Long:      `Attach an existing flexible IP to a specified Elastic Metal server.`,
		Namespace: "fip",
		Resource:  "ip",
		Verb:      "attach",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.AttachFlexibleIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fips-ids.{index}",
				Short:      `List of flexible IP IDs to attach to a server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "server-id",
				Short:      `ID of the server on which to attach the flexible IPs`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.AttachFlexibleIPRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			return api.AttachFlexibleIP(request)

		},
	}
}

func fipIPDetach() *core.Command {
	return &core.Command{
		Short:     `Detach an existing flexible IP from a server`,
		Long:      `Detach an existing flexible IP from a specified Elastic Metal server.`,
		Namespace: "fip",
		Resource:  "ip",
		Verb:      "detach",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.DetachFlexibleIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fips-ids.{index}",
				Short:      `List of flexible IP IDs to detach from a server. Multiple IDs can be provided. Note that flexible IPs must belong to the same MAC group.`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.DetachFlexibleIPRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			return api.DetachFlexibleIP(request)

		},
	}
}

func fipMacCreate() *core.Command {
	return &core.Command{
		Short:     `Generate a virtual MAC address on an existing flexible IP`,
		Long:      `Generate a virtual MAC (Media Access Control) address on an existing flexible IP.`,
		Namespace: "fip",
		Resource:  "mac",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.GenerateMACAddrRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fip-id",
				Short:      `ID of the flexible IP for which to generate a virtual MAC`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "mac-type",
				Short:      `TODO`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "vmware", "xen", "kvm"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.GenerateMACAddrRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			return api.GenerateMACAddr(request)

		},
	}
}

func fipMacDuplicate() *core.Command {
	return &core.Command{
		Short:     `Duplicate a virtual MAC address to another flexible IP`,
		Long:      `Duplicate a virtual MAC address from a given flexible IP to another flexible IP attached to the same server.`,
		Namespace: "fip",
		Resource:  "mac",
		Verb:      "duplicate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.DuplicateMACAddrRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fip-id",
				Short:      `ID of the flexible IP on which to duplicate the virtual MAC`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "duplicate-from-fip-id",
				Short:      `ID of the flexible IP to duplicate the Virtual MAC from`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.DuplicateMACAddrRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			return api.DuplicateMACAddr(request)

		},
	}
}

func fipMacMove() *core.Command {
	return &core.Command{
		Short:     `Relocate an existing virtual MAC address to a different flexible IP`,
		Long:      `Relocate a virtual MAC (Media Access Control) address from an existing flexible IP to a different flexible IP.`,
		Namespace: "fip",
		Resource:  "mac",
		Verb:      "move",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.MoveMACAddrRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fip-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "dst-fip-id",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.MoveMACAddrRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			return api.MoveMACAddr(request)

		},
	}
}

func fipMacDelete() *core.Command {
	return &core.Command{
		Short:     `Detach a given virtual MAC address from an existing flexible IP`,
		Long:      `Detach a given MAC (Media Access Control) address from an existing flexible IP.`,
		Namespace: "fip",
		Resource:  "mac",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(flexibleip.DeleteMACAddrRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "fip-id",
				Short:      `ID of the flexible IP from which to delete the virtual MAC`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*flexibleip.DeleteMACAddrRequest)

			client := core.ExtractClient(ctx)
			api := flexibleip.NewAPI(client)
			e = api.DeleteMACAddr(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "mac",
				Verb:     "delete",
			}, nil
		},
	}
}

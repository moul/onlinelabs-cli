// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package lb

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		lbRoot(),
		lbLB(),
		lbIP(),
		lbBackend(),
		lbFrontend(),
		lbCertificate(),
		lbACL(),
		lbLBTypes(),
		lbPrivateNetwork(),
		lbRoute(),
		lbSubscriber(),
		lbLBList(),
		lbLBCreate(),
		lbLBGet(),
		lbLBUpdate(),
		lbLBDelete(),
		lbLBMigrate(),
		lbIPList(),
		lbIPCreate(),
		lbIPGet(),
		lbIPDelete(),
		lbIPUpdate(),
		lbBackendList(),
		lbBackendCreate(),
		lbBackendGet(),
		lbBackendUpdate(),
		lbBackendDelete(),
		lbBackendAddServers(),
		lbBackendRemoveServers(),
		lbBackendSetServers(),
		lbBackendUpdateHealthcheck(),
		lbFrontendList(),
		lbFrontendCreate(),
		lbFrontendGet(),
		lbFrontendUpdate(),
		lbFrontendDelete(),
		lbRouteList(),
		lbRouteCreate(),
		lbRouteGet(),
		lbRouteUpdate(),
		lbRouteDelete(),
		lbLBGetStats(),
		lbBackendListStatistics(),
		lbACLList(),
		lbACLCreate(),
		lbACLGet(),
		lbACLUpdate(),
		lbACLDelete(),
		lbACLSet(),
		lbCertificateCreate(),
		lbCertificateList(),
		lbCertificateGet(),
		lbCertificateUpdate(),
		lbCertificateDelete(),
		lbLBTypesList(),
		lbSubscriberCreate(),
		lbSubscriberGet(),
		lbSubscriberList(),
		lbSubscriberUpdate(),
		lbSubscriberDelete(),
		lbSubscriberSubscribe(),
		lbSubscriberUnsubscribe(),
		lbPrivateNetworkList(),
		lbPrivateNetworkAttach(),
		lbPrivateNetworkDetach(),
	)
}
func lbRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Scaleway Load Balancer services`,
		Long:      `This API allows you to manage your Scaleway Load Balancer services.`,
		Namespace: "lb",
	}
}

func lbLB() *core.Command {
	return &core.Command{
		Short:     `Load balancer management commands`,
		Long:      `Load balancer management commands.`,
		Namespace: "lb",
		Resource:  "lb",
	}
}

func lbIP() *core.Command {
	return &core.Command{
		Short:     `IP management commands`,
		Long:      `IP management commands.`,
		Namespace: "lb",
		Resource:  "ip",
	}
}

func lbBackend() *core.Command {
	return &core.Command{
		Short:     `Backend management commands`,
		Long:      `Backend management commands.`,
		Namespace: "lb",
		Resource:  "backend",
	}
}

func lbFrontend() *core.Command {
	return &core.Command{
		Short:     `Frontend management commands`,
		Long:      `Frontend management commands.`,
		Namespace: "lb",
		Resource:  "frontend",
	}
}

func lbCertificate() *core.Command {
	return &core.Command{
		Short:     `TLS certificate management commands`,
		Long:      `TLS certificate management commands.`,
		Namespace: "lb",
		Resource:  "certificate",
	}
}

func lbACL() *core.Command {
	return &core.Command{
		Short:     `Access Control List (ACL) management commands`,
		Long:      `Access Control List (ACL) management commands.`,
		Namespace: "lb",
		Resource:  "acl",
	}
}

func lbLBTypes() *core.Command {
	return &core.Command{
		Short:     `Load balancer types management commands`,
		Long:      `Load balancer types management commands.`,
		Namespace: "lb",
		Resource:  "lb-types",
	}
}

func lbPrivateNetwork() *core.Command {
	return &core.Command{
		Short:     `Private networks management commands`,
		Long:      `Private networks management commands.`,
		Namespace: "lb",
		Resource:  "private-network",
	}
}

func lbRoute() *core.Command {
	return &core.Command{
		Short:     `Route rules management commands`,
		Long:      `Route rules management commands.`,
		Namespace: "lb",
		Resource:  "route",
	}
}

func lbSubscriber() *core.Command {
	return &core.Command{
		Short:     `Subscriber management commands`,
		Long:      `Subscriber management commands.`,
		Namespace: "lb",
		Resource:  "subscriber",
	}
}

func lbLBList() *core.Command {
	return &core.Command{
		Short:     `List Load Balancers`,
		Long:      `List all Load Balancers in the specified zone, for a Scaleway Organization or Scaleway Project. By default, the Load Balancers returned in the list are ordered by creation date in ascending order, though this can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "lb",
		Resource:  "lb",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListLBsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Load Balancer name to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Sort order of Load Balancers in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "project-id",
				Short:      `Project ID to filter for, only Load Balancers from this Project will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Filter by tag, only Load Balancers with one or more matching tags will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID to filter for, only Load Balancers from this Organization will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListLBsRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListLBs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.LBs, nil

		},
	}
}

func lbLBCreate() *core.Command {
	return &core.Command{
		Short:     `Create a Load Balancer`,
		Long:      `Create a new Load Balancer. Note that the Load Balancer will be created without frontends or backends; these must be created separately via the dedicated endpoints.`,
		Namespace: "lb",
		Resource:  "lb",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPICreateLBRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "name",
				Short:      `Name for the Load Balancer`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("lb"),
			},
			{
				Name:       "description",
				Short:      `Description for the Load Balancer`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ip-id",
				Short:      `ID of an existing flexible IP address to attach to the Load Balancer`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "assign-flexible-ip",
				Short:      `Defines whether to automatically assign a flexible public IP to the Load Balancer. Default value is ` + "`" + `true` + "`" + ` (assign).`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("true"),
			},
			{
				Name:       "assign-flexible-ipv6",
				Short:      `Defines whether to automatically assign a flexible public IPv6 to the Load Balancer. Default value is ` + "`" + `false` + "`" + ` (do not assign).`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("false"),
			},
			{
				Name:       "ip-ids.{index}",
				Short:      `List of IP IDs to attach to the Load Balancer`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `List of tags for the Load Balancer`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "type",
				Short:      `Load Balancer commercial offer type. Use the Load Balancer types endpoint to retrieve a list of available offer types`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ssl-compatibility-level",
				Short:      `Determines the minimal SSL version which needs to be supported on the client side, in an SSL/TLS offloading context. Intermediate is suitable for general-purpose servers with a variety of clients, recommended for almost all systems. Modern is suitable for services with clients that support TLS 1.3 and do not need backward compatibility. Old is compatible with a small number of very old clients and should be used only as a last resort`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"ssl_compatibility_level_unknown", "ssl_compatibility_level_intermediate", "ssl_compatibility_level_modern", "ssl_compatibility_level_old"},
			},
			core.OrganizationIDArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPICreateLBRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.CreateLB(request)

		},
	}
}

func lbLBGet() *core.Command {
	return &core.Command{
		Short:     `Get a Load Balancer`,
		Long:      `Retrieve information about an existing Load Balancer, specified by its Load Balancer ID. Its full details, including name, status and IP address, are returned in the response object.`,
		Namespace: "lb",
		Resource:  "lb",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetLBRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetLBRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetLB(request)

		},
	}
}

func lbLBUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a Load Balancer`,
		Long:      `Update the parameters of an existing Load Balancer, specified by its Load Balancer ID. Note that the request type is PUT and not PATCH. You must set all parameters.`,
		Namespace: "lb",
		Resource:  "lb",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateLBRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Load Balancer name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Load Balancer description`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `List of tags for the Load Balancer`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ssl-compatibility-level",
				Short:      `Determines the minimal SSL version which needs to be supported on the client side, in an SSL/TLS offloading context. Intermediate is suitable for general-purpose servers with a variety of clients, recommended for almost all systems. Modern is suitable for services with clients that support TLS 1.3 and don't need backward compatibility. Old is compatible with a small number of very old clients and should be used only as a last resort`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"ssl_compatibility_level_unknown", "ssl_compatibility_level_intermediate", "ssl_compatibility_level_modern", "ssl_compatibility_level_old"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateLBRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateLB(request)

		},
	}
}

func lbLBDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Load Balancer`,
		Long:      `Delete an existing Load Balancer, specified by its Load Balancer ID. Deleting a Load Balancer is permanent, and cannot be undone. The Load Balancer's flexible IP address can either be deleted with the Load Balancer, or kept in your account for future use.`,
		Namespace: "lb",
		Resource:  "lb",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIDeleteLBRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `ID of the Load Balancer to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "release-ip",
				Short:      `Defines whether the Load Balancer's flexible IP should be deleted. Set to true to release the flexible IP, or false to keep it available in your account for future Load Balancers`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIDeleteLBRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.DeleteLB(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "lb",
				Verb:     "delete",
			}, nil
		},
	}
}

func lbLBMigrate() *core.Command {
	return &core.Command{
		Short:     `Migrate a Load Balancer`,
		Long:      `Migrate an existing Load Balancer from one commercial type to another. Allows you to scale your Load Balancer up or down in terms of bandwidth or multi-cloud provision.`,
		Namespace: "lb",
		Resource:  "lb",
		Verb:      "migrate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIMigrateLBRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "type",
				Short:      `Load Balancer type to migrate to (use the List all Load Balancer offer types endpoint to get a list of available offer types)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIMigrateLBRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.MigrateLB(request)

		},
	}
}

func lbIPList() *core.Command {
	return &core.Command{
		Short:     `List IP addresses`,
		Long:      `List the Load Balancer flexible IP addresses held in the account (filtered by Organization ID or Project ID). It is also possible to search for a specific IP address.`,
		Namespace: "lb",
		Resource:  "ip",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListIPsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip-address",
				Short:      `IP address to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Project ID to filter for, only Load Balancer IP addresses from this Project will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ip-type",
				Short:      `IP type to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"all", "ipv4", "ipv6"},
			},
			{
				Name:       "tags.{index}",
				Short:      `Tag to filter for, only IPs with one or more matching tags will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID to filter for, only Load Balancer IP addresses from this Organization will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListIPsRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListIPs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.IPs, nil

		},
	}
}

func lbIPCreate() *core.Command {
	return &core.Command{
		Short:     `Create an IP address`,
		Long:      `Create a new Load Balancer flexible IP address, in the specified Scaleway Project. This can be attached to new Load Balancers created in the future.`,
		Namespace: "lb",
		Resource:  "ip",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPICreateIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "reverse",
				Short:      `Reverse DNS (domain name) for the IP address`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "is-ipv6",
				Short:      `If true, creates a Flexible IP with an ipv6 address`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `List of tags for the IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPICreateIPRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.CreateIP(request)

		},
	}
}

func lbIPGet() *core.Command {
	return &core.Command{
		Short:     `Get an IP address`,
		Long:      `Retrieve the full details of a Load Balancer flexible IP address.`,
		Namespace: "lb",
		Resource:  "ip",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip-id",
				Short:      `IP address ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetIPRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetIP(request)

		},
	}
}

func lbIPDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an IP address`,
		Long:      `Delete a Load Balancer flexible IP address. This action is irreversible, and cannot be undone.`,
		Namespace: "lb",
		Resource:  "ip",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIReleaseIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip-id",
				Short:      `IP address ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIReleaseIPRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.ReleaseIP(request)
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

func lbIPUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an IP address`,
		Long:      `Update the reverse DNS of a Load Balancer flexible IP address.`,
		Namespace: "lb",
		Resource:  "ip",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip-id",
				Short:      `IP address ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "reverse",
				Short:      `Reverse DNS (domain name) for the IP address`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "lb-id",
				Short:      `ID of the server on which to attach the flexible IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `List of tags for the IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateIPRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateIP(request)

		},
	}
}

func lbBackendList() *core.Command {
	return &core.Command{
		Short:     `List the backends of a given Load Balancer`,
		Long:      `List all the backends of a Load Balancer, specified by its Load Balancer ID. By default, results are returned in ascending order by the creation date of each backend. The response is an array of backend objects, containing full details of each one including their configuration parameters such as protocol, port and forwarding algorithm.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListBackendsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the backend to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Sort order of backends in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListBackendsRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListBackends(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Backends, nil

		},
	}
}

func lbBackendCreate() *core.Command {
	return &core.Command{
		Short:     `Create a backend for a given Load Balancer`,
		Long:      `Create a new backend for a given Load Balancer, specifying its full configuration including protocol, port and forwarding algorithm.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPICreateBackendRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name for the backend`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("lbb"),
			},
			{
				Name:       "forward-protocol",
				Short:      `Protocol to be used by the backend when forwarding traffic to backend servers`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"tcp", "http"},
			},
			{
				Name:       "forward-port",
				Short:      `Port to be used by the backend when forwarding traffic to backend servers`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "forward-port-algorithm",
				Short:      `Load balancing algorithm to be used when determining which backend server to forward new traffic to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("roundrobin"),
				EnumValues: []string{"roundrobin", "leastconn", "first"},
			},
			{
				Name:       "sticky-sessions",
				Short:      `Defines whether to activate sticky sessions (binding a particular session to a particular backend server) and the method to use if so. None disables sticky sessions. Cookie-based uses an HTTP cookie TO stick a session to a backend server. Table-based uses the source (client) IP address to stick a session to a backend server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("none"),
				EnumValues: []string{"none", "cookie", "table"},
			},
			{
				Name:       "sticky-sessions-cookie-name",
				Short:      `Cookie name for cookie-based sticky sessions`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.port",
				Short:      `Port to use for the backend server health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.check-delay",
				Short:      `Time to wait between two consecutive health checks`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("3s"),
			},
			{
				Name:       "health-check.check-timeout",
				Short:      `Maximum time a backend server has to reply to the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("1s"),
			},
			{
				Name:       "health-check.check-max-retries",
				Short:      `Number of consecutive unsuccessful health checks after which the server will be considered dead`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.mysql-config.user",
				Short:      `MySQL user to use for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.pgsql-config.user",
				Short:      `PostgreSQL user to use for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.http-config.uri",
				Short:      `HTTP URI used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.http-config.method",
				Short:      `HTTP method used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.http-config.code",
				Short:      `HTTP response code expected for a successful health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.http-config.host-header",
				Short:      `HTTP host header used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.https-config.uri",
				Short:      `HTTP URI used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.https-config.method",
				Short:      `HTTP method used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.https-config.code",
				Short:      `HTTP response code expected for a successful health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.https-config.host-header",
				Short:      `HTTP host header used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.https-config.sni",
				Short:      `SNI used for SSL health checks`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.check-send-proxy",
				Short:      `Defines whether proxy protocol should be activated for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.transient-check-delay",
				Short:      `Time to wait between two consecutive health checks when a backend server is in a transient state (going UP or DOWN)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("0.5s"),
			},
			{
				Name:       "server-ip.{index}",
				Short:      `List of backend server IP addresses (IPv4 or IPv6) the backend should forward traffic to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "send-proxy-v2",
				Short:      `Deprecated in favor of proxy_protocol field`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "timeout-server",
				Short:      `Maximum allowed time for a backend server to process a request`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("5m"),
			},
			{
				Name:       "timeout-connect",
				Short:      `Maximum allowed time for establishing a connection to a backend server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("5s"),
			},
			{
				Name:       "timeout-tunnel",
				Short:      `Maximum allowed tunnel inactivity time after Websocket is established (takes precedence over client and server timeout)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("15m"),
			},
			{
				Name:       "on-marked-down-action",
				Short:      `Action to take when a backend server is marked as down`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"on_marked_down_action_none", "shutdown_sessions"},
			},
			{
				Name:       "proxy-protocol",
				Short:      `Protocol to use between the Load Balancer and backend servers. Allows the backend servers to be informed of the client's real IP address. The PROXY protocol must be supported by the backend servers' software`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"proxy_protocol_unknown", "proxy_protocol_none", "proxy_protocol_v1", "proxy_protocol_v2", "proxy_protocol_v2_ssl", "proxy_protocol_v2_ssl_cn"},
			},
			{
				Name:       "failover-host",
				Short:      `Scaleway Object Storage bucket website to be served as failover if all backend servers are down, e.g. failover-website.s3-website.fr-par.scw.cloud`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ssl-bridging",
				Short:      `Defines whether to enable SSL bridging between the Load Balancer and backend servers`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ignore-ssl-server-verify",
				Short:      `Defines whether the server certificate verification should be ignored`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "redispatch-attempt-count",
				Short:      `Whether to use another backend server on each attempt`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-retries",
				Short:      `Number of retries when a backend server connection failed`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-connections",
				Short:      `Maximum number of connections allowed per backend server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "timeout-queue",
				Short:      `Maximum time for a request to be left pending in queue when ` + "`" + `max_connections` + "`" + ` is reached`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPICreateBackendRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.CreateBackend(request)

		},
	}
}

func lbBackendGet() *core.Command {
	return &core.Command{
		Short:     `Get a backend of a given Load Balancer`,
		Long:      `Get the full details of a given backend, specified by its backend ID. The response contains the backend's full configuration parameters including protocol, port and forwarding algorithm.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetBackendRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "backend-id",
				Short:      `Backend ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetBackendRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetBackend(request)

		},
	}
}

func lbBackendUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a backend of a given Load Balancer`,
		Long:      `Update a backend of a given Load Balancer, specified by its backend ID. Note that the request type is PUT and not PATCH. You must set all parameters.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateBackendRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "backend-id",
				Short:      `Backend ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Backend name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "forward-protocol",
				Short:      `Protocol to be used by the backend when forwarding traffic to backend servers`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"tcp", "http"},
			},
			{
				Name:       "forward-port",
				Short:      `Port to be used by the backend when forwarding traffic to backend servers`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "forward-port-algorithm",
				Short:      `Load balancing algorithm to be used when determining which backend server to forward new traffic to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"roundrobin", "leastconn", "first"},
			},
			{
				Name:       "sticky-sessions",
				Short:      `Defines whether to activate sticky sessions (binding a particular session to a particular backend server) and the method to use if so. None disables sticky sessions. Cookie-based uses an HTTP cookie to stick a session to a backend server. Table-based uses the source (client) IP address to stick a session to a backend server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"none", "cookie", "table"},
			},
			{
				Name:       "sticky-sessions-cookie-name",
				Short:      `Cookie name for cookie-based sticky sessions`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "send-proxy-v2",
				Short:      `Deprecated in favor of proxy_protocol field`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "timeout-server",
				Short:      `Maximum allowed time for a backend server to process a request`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("5m"),
			},
			{
				Name:       "timeout-connect",
				Short:      `Maximum allowed time for establishing a connection to a backend server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("5s"),
			},
			{
				Name:       "timeout-tunnel",
				Short:      `Maximum allowed tunnel inactivity time after Websocket is established (takes precedence over client and server timeout)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("15m"),
			},
			{
				Name:       "on-marked-down-action",
				Short:      `Action to take when a backend server is marked as down`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"on_marked_down_action_none", "shutdown_sessions"},
			},
			{
				Name:       "proxy-protocol",
				Short:      `Protocol to use between the Load Balancer and backend servers. Allows the backend servers to be informed of the client's real IP address. The PROXY protocol must be supported by the backend servers' software`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"proxy_protocol_unknown", "proxy_protocol_none", "proxy_protocol_v1", "proxy_protocol_v2", "proxy_protocol_v2_ssl", "proxy_protocol_v2_ssl_cn"},
			},
			{
				Name:       "failover-host",
				Short:      `Scaleway Object Storage bucket website to be served as failover if all backend servers are down, e.g. failover-website.s3-website.fr-par.scw.cloud`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ssl-bridging",
				Short:      `Defines whether to enable SSL bridging between the Load Balancer and backend servers`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ignore-ssl-server-verify",
				Short:      `Defines whether the server certificate verification should be ignored`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "redispatch-attempt-count",
				Short:      `Whether to use another backend server on each attempt`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-retries",
				Short:      `Number of retries when a backend server connection failed`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-connections",
				Short:      `Maximum number of connections allowed per backend server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "timeout-queue",
				Short:      `Maximum time for a request to be left pending in queue when ` + "`" + `max_connections` + "`" + ` is reached`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateBackendRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateBackend(request)

		},
	}
}

func lbBackendDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a backend of a given Load Balancer`,
		Long:      `Delete a backend of a given Load Balancer, specified by its backend ID. This action is irreversible and cannot be undone.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIDeleteBackendRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "backend-id",
				Short:      `ID of the backend to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIDeleteBackendRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.DeleteBackend(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "backend",
				Verb:     "delete",
			}, nil
		},
	}
}

func lbBackendAddServers() *core.Command {
	return &core.Command{
		Short:     `Add a set of backend servers to a given backend`,
		Long:      `For a given backend specified by its backend ID, add a set of backend servers (identified by their IP addresses) it should forward traffic to. These will be appended to any existing set of backend servers for this backend.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "add-servers",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIAddBackendServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "backend-id",
				Short:      `Backend ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "server-ip.{index}",
				Short:      `List of IP addresses to add to backend servers`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIAddBackendServersRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.AddBackendServers(request)

		},
	}
}

func lbBackendRemoveServers() *core.Command {
	return &core.Command{
		Short:     `Remove a set of servers for a given backend`,
		Long:      `For a given backend specified by its backend ID, remove the specified backend servers (identified by their IP addresses) so that it no longer forwards traffic to them.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "remove-servers",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIRemoveBackendServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "backend-id",
				Short:      `Backend ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "server-ip.{index}",
				Short:      `List of IP addresses to remove from backend servers`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIRemoveBackendServersRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.RemoveBackendServers(request)

		},
	}
}

func lbBackendSetServers() *core.Command {
	return &core.Command{
		Short:     `Define all backend servers for a given backend`,
		Long:      `For a given backend specified by its backend ID, define the set of backend servers (identified by their IP addresses) that it should forward traffic to. Any existing backend servers configured for this backend will be removed.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "set-servers",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPISetBackendServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "backend-id",
				Short:      `Backend ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "server-ip.{index}",
				Short:      `List of IP addresses for backend servers. Any other existing backend servers will be removed`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPISetBackendServersRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.SetBackendServers(request)

		},
	}
}

func lbBackendUpdateHealthcheck() *core.Command {
	return &core.Command{
		Short:     `Update a health check for a given backend`,
		Long:      `Update the configuration of the health check performed by a given backend to verify the health of its backend servers, identified by its backend ID. Note that the request type is PUT and not PATCH. You must set all parameters.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "update-healthcheck",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateHealthCheckRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "port",
				Short:      `Port to use for the backend server health check`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "check-delay",
				Short:      `Time to wait between two consecutive health checks`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "check-timeout",
				Short:      `Maximum time a backend server has to reply to the health check`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "check-max-retries",
				Short:      `Number of consecutive unsuccessful health checks after which the server will be considered dead`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "backend-id",
				Short:      `Backend ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "check-send-proxy",
				Short:      `Defines whether proxy protocol should be activated for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "mysql-config.user",
				Short:      `MySQL user to use for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pgsql-config.user",
				Short:      `PostgreSQL user to use for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "http-config.uri",
				Short:      `HTTP URI used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "http-config.method",
				Short:      `HTTP method used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "http-config.code",
				Short:      `HTTP response code expected for a successful health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "http-config.host-header",
				Short:      `HTTP host header used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "https-config.uri",
				Short:      `HTTP URI used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "https-config.method",
				Short:      `HTTP method used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "https-config.code",
				Short:      `HTTP response code expected for a successful health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "https-config.host-header",
				Short:      `HTTP host header used for the health check`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "https-config.sni",
				Short:      `SNI used for SSL health checks`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "transient-check-delay",
				Short:      `Time to wait between two consecutive health checks when a backend server is in a transient state (going UP or DOWN)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("0.5s"),
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateHealthCheckRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateHealthCheck(request)

		},
	}
}

func lbFrontendList() *core.Command {
	return &core.Command{
		Short:     `List frontends of a given Load Balancer`,
		Long:      `List all the frontends of a Load Balancer, specified by its Load Balancer ID. By default, results are returned in ascending order by the creation date of each frontend. The response is an array of frontend objects, containing full details of each one including the port they listen on and the backend they are attached to.`,
		Namespace: "lb",
		Resource:  "frontend",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListFrontendsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the frontend to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Sort order of frontends in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListFrontendsRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListFrontends(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Frontends, nil

		},
	}
}

func lbFrontendCreate() *core.Command {
	return &core.Command{
		Short:     `Create a frontend in a given Load Balancer`,
		Long:      `Create a new frontend for a given Load Balancer, specifying its configuration including the port it should listen on and the backend to attach it to.`,
		Namespace: "lb",
		Resource:  "frontend",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPICreateFrontendRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name for the frontend`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("lbf"),
			},
			{
				Name:       "inbound-port",
				Short:      `Port the frontend should listen on`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID (ID of the Load Balancer to attach the frontend to)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "backend-id",
				Short:      `Backend ID (ID of the backend the frontend should pass traffic to)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "timeout-client",
				Short:      `Maximum allowed inactivity time on the client side`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("5m"),
			},
			{
				Name:       "certificate-id",
				Short:      `Certificate ID, deprecated in favor of certificate_ids array`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "certificate-ids.{index}",
				Short:      `List of SSL/TLS certificate IDs to bind to the frontend`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "enable-http3",
				Short:      `Defines whether to enable HTTP/3 protocol on the frontend`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPICreateFrontendRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.CreateFrontend(request)

		},
	}
}

func lbFrontendGet() *core.Command {
	return &core.Command{
		Short:     `Get a frontend`,
		Long:      `Get the full details of a given frontend, specified by its frontend ID. The response contains the frontend's full configuration parameters including the backend it is attached to, the port it listens on, and any certificates it has.`,
		Namespace: "lb",
		Resource:  "frontend",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetFrontendRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "frontend-id",
				Short:      `Frontend ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetFrontendRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetFrontend(request)

		},
	}
}

func lbFrontendUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a frontend`,
		Long:      `Update a given frontend, specified by its frontend ID. You can update configuration parameters including its name and the port it listens on. Note that the request type is PUT and not PATCH. You must set all parameters.`,
		Namespace: "lb",
		Resource:  "frontend",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateFrontendRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "frontend-id",
				Short:      `Frontend ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Frontend name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "inbound-port",
				Short:      `Port the frontend should listen on`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "backend-id",
				Short:      `Backend ID (ID of the backend the frontend should pass traffic to)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "timeout-client",
				Short:      `Maximum allowed inactivity time on the client side`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("5m"),
			},
			{
				Name:       "certificate-id",
				Short:      `Certificate ID, deprecated in favor of certificate_ids array`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "certificate-ids.{index}",
				Short:      `List of SSL/TLS certificate IDs to bind to the frontend`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "enable-http3",
				Short:      `Defines whether to enable HTTP/3 protocol on the frontend`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateFrontendRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateFrontend(request)

		},
	}
}

func lbFrontendDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a frontend`,
		Long:      `Delete a given frontend, specified by its frontend ID. This action is irreversible and cannot be undone.`,
		Namespace: "lb",
		Resource:  "frontend",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIDeleteFrontendRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "frontend-id",
				Short:      `ID of the frontend to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIDeleteFrontendRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.DeleteFrontend(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "frontend",
				Verb:     "delete",
			}, nil
		},
	}
}

func lbRouteList() *core.Command {
	return &core.Command{
		Short:     `List all routes`,
		Long:      `List all routes for a given frontend. The response is an array of routes, each one  with a specified backend to direct to if a certain condition is matched (based on the value of the SNI field or HTTP Host header).`,
		Namespace: "lb",
		Resource:  "route",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListRoutesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order of routes in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "frontend-id",
				Short:      `Frontend ID to filter for, only Routes from this Frontend will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListRoutesRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListRoutes(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Routes, nil

		},
	}
}

func lbRouteCreate() *core.Command {
	return &core.Command{
		Short:     `Create a route`,
		Long:      `Create a new route on a given frontend. To configure a route, specify the backend to direct to if a certain condition is matched (based on the value of the SNI field or HTTP Host header).`,
		Namespace: "lb",
		Resource:  "route",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPICreateRouteRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "frontend-id",
				Short:      `ID of the source frontend to create the route on`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "backend-id",
				Short:      `ID of the target backend for the route`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.sni",
				Short:      `Server Name Indication (SNI) value to match`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.host-header",
				Short:      `HTTP host header to match`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPICreateRouteRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.CreateRoute(request)

		},
	}
}

func lbRouteGet() *core.Command {
	return &core.Command{
		Short:     `Get a route`,
		Long:      `Retrieve information about an existing route, specified by its route ID. Its full details, origin frontend, target backend and match condition, are returned in the response object.`,
		Namespace: "lb",
		Resource:  "route",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetRouteRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "route-id",
				Short:      `Route ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetRouteRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetRoute(request)

		},
	}
}

func lbRouteUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a route`,
		Long:      `Update the configuration of an existing route, specified by its route ID.`,
		Namespace: "lb",
		Resource:  "route",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateRouteRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "route-id",
				Short:      `Route ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "backend-id",
				Short:      `ID of the target backend for the route`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.sni",
				Short:      `Server Name Indication (SNI) value to match`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.host-header",
				Short:      `HTTP host header to match`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateRouteRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateRoute(request)

		},
	}
}

func lbRouteDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a route`,
		Long:      `Delete an existing route, specified by its route ID. Deleting a route is permanent, and cannot be undone.`,
		Namespace: "lb",
		Resource:  "route",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIDeleteRouteRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "route-id",
				Short:      `Route ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIDeleteRouteRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.DeleteRoute(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "route",
				Verb:     "delete",
			}, nil
		},
	}
}

func lbLBGetStats() *core.Command {
	return &core.Command{
		Short:     `Get usage statistics of a given Load Balancer`,
		Long:      `Get usage statistics of a given Load Balancer.`,
		Namespace: "lb",
		Resource:  "lb",
		Verb:      "get-stats",
		// Deprecated:    true,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetLBStatsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "backend-id",
				Short:      `ID of the backend`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetLBStatsRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetLBStats(request)

		},
	}
}

func lbBackendListStatistics() *core.Command {
	return &core.Command{
		Short:     `List backend server statistics`,
		Long:      `List information about your backend servers, including their state and the result of their last health check.`,
		Namespace: "lb",
		Resource:  "backend",
		Verb:      "list-statistics",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListBackendStatsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "backend-id",
				Short:      `ID of the backend`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListBackendStatsRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListBackendStats(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.BackendServersStats, nil

		},
	}
}

func lbACLList() *core.Command {
	return &core.Command{
		Short:     `List ACLs for a given frontend`,
		Long:      `List the ACLs for a given frontend, specified by its frontend ID. The response is an array of ACL objects, each one representing an ACL that denies or allows traffic based on certain conditions.`,
		Namespace: "lb",
		Resource:  "acl",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListACLsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "frontend-id",
				Short:      `Frontend ID (ACLs attached to this frontend will be returned in the response)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Sort order of ACLs in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "name",
				Short:      `ACL name to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListACLsRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListACLs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.ACLs, nil

		},
	}
}

func lbACLCreate() *core.Command {
	return &core.Command{
		Short:     `Create an ACL for a given frontend`,
		Long:      `Create a new ACL for a given frontend. Each ACL must have a name, an action to perform (allow or deny), and a match rule (the action is carried out when the incoming traffic matches the rule).`,
		Namespace: "lb",
		Resource:  "acl",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPICreateACLRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "frontend-id",
				Short:      `Frontend ID to attach the ACL to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `ACL name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("acl"),
			},
			{
				Name:       "action.type",
				Short:      `Action to take when incoming traffic matches an ACL filter`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"allow", "deny", "redirect"},
			},
			{
				Name:       "action.redirect.type",
				Short:      `Redirect type`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"location", "scheme"},
			},
			{
				Name:       "action.redirect.target",
				Short:      `Redirect target. For a location redirect, you can use a URL e.g. ` + "`" + `https://scaleway.com` + "`" + `. Using a scheme name (e.g. ` + "`" + `https` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `ftp` + "`" + `, ` + "`" + `git` + "`" + `) will replace the request's original scheme. This can be useful to implement HTTP to HTTPS redirects. Valid placeholders that can be used in a ` + "`" + `location` + "`" + ` redirect to preserve parts of the original request in the redirection URL are \{\{host\}\}, \{\{query\}\}, \{\{path\}\} and \{\{scheme\}\}`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "action.redirect.code",
				Short:      `HTTP redirect code to use. Valid values are 301, 302, 303, 307 and 308. Default value is 302`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.ip-subnet.{index}",
				Short:      `List of IPs or CIDR v4/v6 addresses to filter for from the client side`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.http-filter",
				Short:      `Type of HTTP filter to match. Extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part). Defines where to filter for the http_filter_value. Only supported for HTTP backends`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"acl_http_filter_none", "path_begin", "path_end", "regex", "http_header_match"},
			},
			{
				Name:       "match.http-filter-value.{index}",
				Short:      `List of values to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.http-filter-option",
				Short:      `Name of the HTTP header to filter on if ` + "`" + `http_header_match` + "`" + ` was selected in ` + "`" + `http_filter` + "`" + ``,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.invert",
				Short:      `Defines whether to invert the match condition. If set to ` + "`" + `true` + "`" + `, the ACL carries out its action when the condition DOES NOT match`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "index",
				Short:      `Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `ACL description`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPICreateACLRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.CreateACL(request)

		},
	}
}

func lbACLGet() *core.Command {
	return &core.Command{
		Short:     `Get an ACL`,
		Long:      `Get information for a particular ACL, specified by its ACL ID. The response returns full details of the ACL, including its name, action, match rule and frontend.`,
		Namespace: "lb",
		Resource:  "acl",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetACLRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "acl-id",
				Short:      `ACL ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetACLRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetACL(request)

		},
	}
}

func lbACLUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an ACL`,
		Long:      `Update a particular ACL, specified by its ACL ID. You can update details including its name, action and match rule.`,
		Namespace: "lb",
		Resource:  "acl",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateACLRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "acl-id",
				Short:      `ACL ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `ACL name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "action.type",
				Short:      `Action to take when incoming traffic matches an ACL filter`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"allow", "deny", "redirect"},
			},
			{
				Name:       "action.redirect.type",
				Short:      `Redirect type`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"location", "scheme"},
			},
			{
				Name:       "action.redirect.target",
				Short:      `Redirect target. For a location redirect, you can use a URL e.g. ` + "`" + `https://scaleway.com` + "`" + `. Using a scheme name (e.g. ` + "`" + `https` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `ftp` + "`" + `, ` + "`" + `git` + "`" + `) will replace the request's original scheme. This can be useful to implement HTTP to HTTPS redirects. Valid placeholders that can be used in a ` + "`" + `location` + "`" + ` redirect to preserve parts of the original request in the redirection URL are \{\{host\}\}, \{\{query\}\}, \{\{path\}\} and \{\{scheme\}\}`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "action.redirect.code",
				Short:      `HTTP redirect code to use. Valid values are 301, 302, 303, 307 and 308. Default value is 302`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.ip-subnet.{index}",
				Short:      `List of IPs or CIDR v4/v6 addresses to filter for from the client side`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.http-filter",
				Short:      `Type of HTTP filter to match. Extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part). Defines where to filter for the http_filter_value. Only supported for HTTP backends`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"acl_http_filter_none", "path_begin", "path_end", "regex", "http_header_match"},
			},
			{
				Name:       "match.http-filter-value.{index}",
				Short:      `List of values to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.http-filter-option",
				Short:      `Name of the HTTP header to filter on if ` + "`" + `http_header_match` + "`" + ` was selected in ` + "`" + `http_filter` + "`" + ``,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "match.invert",
				Short:      `Defines whether to invert the match condition. If set to ` + "`" + `true` + "`" + `, the ACL carries out its action when the condition DOES NOT match`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "index",
				Short:      `Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `ACL description`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateACLRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateACL(request)

		},
	}
}

func lbACLDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an ACL`,
		Long:      `Delete an ACL, specified by its ACL ID. Deleting an ACL is irreversible and cannot be undone.`,
		Namespace: "lb",
		Resource:  "acl",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIDeleteACLRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "acl-id",
				Short:      `ACL ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIDeleteACLRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.DeleteACL(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "acl",
				Verb:     "delete",
			}, nil
		},
	}
}

func lbACLSet() *core.Command {
	return &core.Command{
		Short:     `Define all ACLs for a given frontend`,
		Long:      `For a given frontend specified by its frontend ID, define and add the complete set of ACLS for that frontend. Any existing ACLs on this frontend will be removed.`,
		Namespace: "lb",
		Resource:  "acl",
		Verb:      "set",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPISetACLsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "acls.{index}.name",
				Short:      `ACL name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acls.{index}.action.type",
				Short:      `Action to take when incoming traffic matches an ACL filter`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"allow", "deny", "redirect"},
			},
			{
				Name:       "acls.{index}.action.redirect.type",
				Short:      `Redirect type`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"location", "scheme"},
			},
			{
				Name:       "acls.{index}.action.redirect.target",
				Short:      `Redirect target. For a location redirect, you can use a URL e.g. ` + "`" + `https://scaleway.com` + "`" + `. Using a scheme name (e.g. ` + "`" + `https` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `ftp` + "`" + `, ` + "`" + `git` + "`" + `) will replace the request's original scheme. This can be useful to implement HTTP to HTTPS redirects. Valid placeholders that can be used in a ` + "`" + `location` + "`" + ` redirect to preserve parts of the original request in the redirection URL are \{\{host\}\}, \{\{query\}\}, \{\{path\}\} and \{\{scheme\}\}`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acls.{index}.action.redirect.code",
				Short:      `HTTP redirect code to use. Valid values are 301, 302, 303, 307 and 308. Default value is 302`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acls.{index}.match.ip-subnet.{index}",
				Short:      `List of IPs or CIDR v4/v6 addresses to filter for from the client side`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acls.{index}.match.http-filter",
				Short:      `Type of HTTP filter to match. Extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part). Defines where to filter for the http_filter_value. Only supported for HTTP backends`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"acl_http_filter_none", "path_begin", "path_end", "regex", "http_header_match"},
			},
			{
				Name:       "acls.{index}.match.http-filter-value.{index}",
				Short:      `List of values to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acls.{index}.match.http-filter-option",
				Short:      `Name of the HTTP header to filter on if ` + "`" + `http_header_match` + "`" + ` was selected in ` + "`" + `http_filter` + "`" + ``,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acls.{index}.match.invert",
				Short:      `Defines whether to invert the match condition. If set to ` + "`" + `true` + "`" + `, the ACL carries out its action when the condition DOES NOT match`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acls.{index}.index",
				Short:      `Priority of this ACL (ACLs are applied in ascending order, 0 is the first ACL executed)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acls.{index}.description",
				Short:      `ACL description`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "frontend-id",
				Short:      `Frontend ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPISetACLsRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.SetACLs(request)

		},
	}
}

func lbCertificateCreate() *core.Command {
	return &core.Command{
		Short:     `Create an SSL/TLS certificate`,
		Long:      `Generate a new SSL/TLS certificate for a given Load Balancer. You can choose to create a Let's Encrypt certificate, or import a custom certificate.`,
		Namespace: "lb",
		Resource:  "certificate",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPICreateCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name for the certificate`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("certificate"),
			},
			{
				Name:       "letsencrypt.common-name",
				Short:      `Main domain name of certificate (this domain must exist and resolve to your Load Balancer IP address)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "letsencrypt.subject-alternative-name.{index}",
				Short:      `Alternative domain names (all domain names must exist and resolve to your Load Balancer IP address)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "custom-certificate.certificate-chain",
				Short:      `Full PEM-formatted certificate, consisting of the entire certificate chain including public key, private key, and (optionally) Certificate Authorities`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPICreateCertificateRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.CreateCertificate(request)

		},
	}
}

func lbCertificateList() *core.Command {
	return &core.Command{
		Short:     `List all SSL/TLS certificates on a given Load Balancer`,
		Long:      `List all the SSL/TLS certificates on a given Load Balancer. The response is an array of certificate objects, which are by default listed in ascending order of creation date.`,
		Namespace: "lb",
		Resource:  "certificate",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListCertificatesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Sort order of certificates in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "name",
				Short:      `Certificate name to filter for, only certificates of this name will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListCertificatesRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListCertificates(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Certificates, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "Type",
			},
			{
				FieldName: "CommonName",
			},
			{
				FieldName: "SubjectAlternativeName",
			},
			{
				FieldName: "Status",
			},
			{
				FieldName: "NotValidBefore",
			},
			{
				FieldName: "NotValidAfter",
			},
			{
				FieldName: "Fingerprint",
			},
			{
				FieldName: "CreatedAt",
			},
			{
				FieldName: "UpdatedAt",
			},
			{
				FieldName: "StatusDetails",
			},
		}},
	}
}

func lbCertificateGet() *core.Command {
	return &core.Command{
		Short:     `Get an SSL/TLS certificate`,
		Long:      `Get information for a particular SSL/TLS certificate, specified by its certificate ID. The response returns full details of the certificate, including its type, main domain name, and alternative domain names.`,
		Namespace: "lb",
		Resource:  "certificate",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "certificate-id",
				Short:      `Certificate ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetCertificateRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetCertificate(request)

		},
	}
}

func lbCertificateUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an SSL/TLS certificate`,
		Long:      `Update the name of a particular SSL/TLS certificate, specified by its certificate ID.`,
		Namespace: "lb",
		Resource:  "certificate",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "certificate-id",
				Short:      `Certificate ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Certificate name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateCertificateRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateCertificate(request)

		},
	}
}

func lbCertificateDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an SSL/TLS certificate`,
		Long:      `Delete an SSL/TLS certificate, specified by its certificate ID. Deleting a certificate is irreversible and cannot be undone.`,
		Namespace: "lb",
		Resource:  "certificate",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIDeleteCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "certificate-id",
				Short:      `Certificate ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIDeleteCertificateRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.DeleteCertificate(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "certificate",
				Verb:     "delete",
			}, nil
		},
	}
}

func lbLBTypesList() *core.Command {
	return &core.Command{
		Short:     `List all Load Balancer offer types`,
		Long:      `List all the different commercial Load Balancer types. The response includes an array of offer types, each with a name, description, and information about its stock availability.`,
		Namespace: "lb",
		Resource:  "lb-types",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListLBTypesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListLBTypesRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListLBTypes(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.LBTypes, nil

		},
	}
}

func lbSubscriberCreate() *core.Command {
	return &core.Command{
		Short:     `Create a subscriber`,
		Long:      `Create a new subscriber, either with an email configuration or a webhook configuration, for a specified Scaleway Project.`,
		Namespace: "lb",
		Resource:  "subscriber",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPICreateSubscriberRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Subscriber name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "email-config.email",
				Short:      `Email address to send alerts to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "webhook-config.uri",
				Short:      `URI to receive POST requests`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectIDArgSpec(),
			core.OrganizationIDArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPICreateSubscriberRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.CreateSubscriber(request)

		},
	}
}

func lbSubscriberGet() *core.Command {
	return &core.Command{
		Short:     `Get a subscriber`,
		Long:      `Retrieve information about an existing subscriber, specified by its subscriber ID. Its full details, including name and email/webhook configuration, are returned in the response object.`,
		Namespace: "lb",
		Resource:  "subscriber",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIGetSubscriberRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "subscriber-id",
				Short:      `Subscriber ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIGetSubscriberRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.GetSubscriber(request)

		},
	}
}

func lbSubscriberList() *core.Command {
	return &core.Command{
		Short:     `List all subscribers`,
		Long:      `List all subscribers to Load Balancer alerts. By default, returns all subscribers to Load Balancer alerts for the Organization associated with the authentication token used for the request.`,
		Namespace: "lb",
		Resource:  "subscriber",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListSubscriberRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order of subscribers in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "name",
				Short:      `Subscriber name to search for`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Filter subscribers by Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Filter subscribers by Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListSubscriberRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListSubscriber(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Subscribers, nil

		},
	}
}

func lbSubscriberUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a subscriber`,
		Long:      `Update the parameters of a given subscriber (e.g. name, webhook configuration, email configuration), specified by its subscriber ID.`,
		Namespace: "lb",
		Resource:  "subscriber",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUpdateSubscriberRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "subscriber-id",
				Short:      `Subscriber ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Subscriber name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "email-config.email",
				Short:      `Email address to send alerts to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "webhook-config.uri",
				Short:      `URI to receive POST requests`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUpdateSubscriberRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UpdateSubscriber(request)

		},
	}
}

func lbSubscriberDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a subscriber`,
		Long:      `Delete an existing subscriber, specified by its subscriber ID. Deleting a subscriber is permanent, and cannot be undone.`,
		Namespace: "lb",
		Resource:  "subscriber",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIDeleteSubscriberRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "subscriber-id",
				Short:      `Subscriber ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIDeleteSubscriberRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.DeleteSubscriber(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "subscriber",
				Verb:     "delete",
			}, nil
		},
	}
}

func lbSubscriberSubscribe() *core.Command {
	return &core.Command{
		Short:     `Subscribe a subscriber to alerts for a given Load Balancer`,
		Long:      `Subscribe an existing subscriber to alerts for a given Load Balancer.`,
		Namespace: "lb",
		Resource:  "subscriber",
		Verb:      "subscribe",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPISubscribeToLBRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "subscriber-id",
				Short:      `Subscriber ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPISubscribeToLBRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.SubscribeToLB(request)

		},
	}
}

func lbSubscriberUnsubscribe() *core.Command {
	return &core.Command{
		Short:     `Unsubscribe a subscriber from alerts for a given Load Balancer`,
		Long:      `Unsubscribe a subscriber from alerts for a given Load Balancer. The subscriber is not deleted, and can be resubscribed in the future if necessary.`,
		Namespace: "lb",
		Resource:  "subscriber",
		Verb:      "unsubscribe",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIUnsubscribeFromLBRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIUnsubscribeFromLBRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.UnsubscribeFromLB(request)

		},
	}
}

func lbPrivateNetworkList() *core.Command {
	return &core.Command{
		Short:     `List Private Networks attached to a Load Balancer`,
		Long:      `List the Private Networks attached to a given Load Balancer, specified by its Load Balancer ID. The response is an array of Private Network objects, giving information including the status, configuration, name and creation date of each Private Network.`,
		Namespace: "lb",
		Resource:  "private-network",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIListLBPrivateNetworksRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order of Private Network objects in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIListLBPrivateNetworksRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListLBPrivateNetworks(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.PrivateNetwork, nil

		},
	}
}

func lbPrivateNetworkAttach() *core.Command {
	return &core.Command{
		Short:     `Attach a Load Balancer to a Private Network`,
		Long:      `Attach a specified Load Balancer to a specified Private Network, defining a static or DHCP configuration for the Load Balancer on the network.`,
		Namespace: "lb",
		Resource:  "private-network",
		Verb:      "attach",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIAttachPrivateNetworkRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load Balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "private-network-id",
				Short:      `Private Network ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ipam-ids.{index}",
				Short:      `IPAM ID of a pre-reserved IP address to assign to the Load Balancer on this Private Network. In the future, it will be possible to specify multiple IPs in this field (IPv4 and IPv6), for now only one ID of an IPv4 address is expected. When null, a new private IP address is created for the Load Balancer on this Private Network.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIAttachPrivateNetworkRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			return api.AttachPrivateNetwork(request)

		},
	}
}

func lbPrivateNetworkDetach() *core.Command {
	return &core.Command{
		Short:     `Detach Load Balancer from Private Network`,
		Long:      `Detach a specified Load Balancer from a specified Private Network.`,
		Namespace: "lb",
		Resource:  "private-network",
		Verb:      "detach",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(lb.ZonedAPIDetachPrivateNetworkRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "lb-id",
				Short:      `Load balancer ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "private-network-id",
				Short:      `Set your instance private network id`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*lb.ZonedAPIDetachPrivateNetworkRequest)

			client := core.ExtractClient(ctx)
			api := lb.NewZonedAPI(client)
			e = api.DetachPrivateNetwork(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "private-network",
				Verb:     "detach",
			}, nil
		},
	}
}

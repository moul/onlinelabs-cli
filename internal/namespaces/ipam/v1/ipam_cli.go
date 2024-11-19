// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package ipam

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/ipam/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		ipamRoot(),
		ipamIP(),
		ipamIPSet(),
		ipamIPCreate(),
		ipamIPDelete(),
		ipamIPSetRelease(),
		ipamIPGet(),
		ipamIPUpdate(),
		ipamIPList(),
	)
}
func ipamRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Scaleway IP addresses with our IP Address Management tool`,
		Long:      `This API allows you to manage your Scaleway IP addresses with our IP Address Management tool.`,
		Namespace: "ipam",
	}
}

func ipamIP() *core.Command {
	return &core.Command{
		Short:     `IP management command`,
		Long:      `*ips_long.`,
		Namespace: "ipam",
		Resource:  "ip",
	}
}

func ipamIPSet() *core.Command {
	return &core.Command{
		Short:     `Management command for sets of IPs`,
		Long:      `*ips_long.`,
		Namespace: "ipam",
		Resource:  "ip-set",
	}
}

func ipamIPCreate() *core.Command {
	return &core.Command{
		Short:     `Reserve a new IP`,
		Long:      `Reserve a new IP from the specified source. Currently IPs can only be reserved from a Private Network.`,
		Namespace: "ipam",
		Resource:  "ip",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(ipam.BookIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "source.zonal",
				Short:      `Zone the IP lives in if the IP is a public zoned IP.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "source.private-network-id",
				Short:      `Private Network the IP lives in if the IP is a private IP.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "source.subnet-id",
				Short:      `Private Network subnet the IP lives in if the IP is a private IP in a Private Network.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "is-ipv6",
				Short:      `Request an IPv6 instead of an IPv4`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "address",
				Short:      `Request this specific IP address in the specified source pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags for the IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "resource.mac-address",
				Short:      `MAC address of the custom resource`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "resource.name",
				Short:      `Name of the custom resource`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*ipam.BookIPRequest)

			client := core.ExtractClient(ctx)
			api := ipam.NewAPI(client)
			return api.BookIP(request)

		},
	}
}

func ipamIPDelete() *core.Command {
	return &core.Command{
		Short:     `Release an IP`,
		Long:      `Release an IP not currently attached to a resource, and returns it to the available IP pool.`,
		Namespace: "ipam",
		Resource:  "ip",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(ipam.ReleaseIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip-id",
				Short:      `IP ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*ipam.ReleaseIPRequest)

			client := core.ExtractClient(ctx)
			api := ipam.NewAPI(client)
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

func ipamIPSetRelease() *core.Command {
	return &core.Command{
		Short:     `Release ipam resources`,
		Long:      `Release ipam resources.`,
		Namespace: "ipam",
		Resource:  "ip-set",
		Verb:      "release",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(ipam.ReleaseIPSetRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip-ids.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*ipam.ReleaseIPSetRequest)

			client := core.ExtractClient(ctx)
			api := ipam.NewAPI(client)
			e = api.ReleaseIPSet(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "ip-set",
				Verb:     "release",
			}, nil
		},
	}
}

func ipamIPGet() *core.Command {
	return &core.Command{
		Short:     `Get an IP`,
		Long:      `Retrieve details of an existing IP, specified by its IP ID.`,
		Namespace: "ipam",
		Resource:  "ip",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(ipam.GetIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip-id",
				Short:      `IP ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*ipam.GetIPRequest)

			client := core.ExtractClient(ctx)
			api := ipam.NewAPI(client)
			return api.GetIP(request)

		},
	}
}

func ipamIPUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an IP`,
		Long:      `Update parameters including tags of the specified IP.`,
		Namespace: "ipam",
		Resource:  "ip",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(ipam.UpdateIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip-id",
				Short:      `IP ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags for the IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "reverses.{index}.hostname",
				Short:      `Reverse domain name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "reverses.{index}.address",
				Short:      `IP corresponding to the hostname`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*ipam.UpdateIPRequest)

			client := core.ExtractClient(ctx)
			api := ipam.NewAPI(client)
			return api.UpdateIP(request)

		},
	}
}

func ipamIPList() *core.Command {
	return &core.Command{
		Short:     `List existing IPs`,
		Long:      `List existing IPs in the specified region using various filters. For example, you can filter for IPs within a specified Private Network, or for public IPs within a specified Project. By default, the IPs returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.`,
		Namespace: "ipam",
		Resource:  "ip",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(ipam.ListIPsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order of the returned IPs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_desc", "created_at_asc", "updated_at_desc", "updated_at_asc", "attached_at_desc", "attached_at_asc"},
			},
			{
				Name:       "project-id",
				Short:      `Project ID to filter for. Only IPs belonging to this Project will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "zonal",
				Short:      `Zone to filter for. Only IPs that are zonal, and in this zone, will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `Private Network to filter for.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "subnet-id",
				Short:      `Subnet ID to filter for.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "vpc-id",
				Short:      `VPC ID to filter for.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "attached",
				Short:      `Defines whether to filter only for IPs which are attached to a resource`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "resource-id",
				Short:      `Resource ID to filter for. Only IPs attached to this resource will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "resource-type",
				Short:      `Resource type to filter for. Only IPs attached to this type of resource will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "custom", "instance_server", "instance_ip", "instance_private_nic", "lb_server", "fip_ip", "vpc_gateway", "vpc_gateway_network", "k8s_node", "k8s_cluster", "rdb_instance", "redis_cluster", "baremetal_server", "baremetal_private_nic", "llm_deployment", "mgdb_instance"},
			},
			{
				Name:       "mac-address",
				Short:      `MAC address to filter for. Only IPs attached to a resource with this MAC address will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags to filter for, only IPs with one or more matching tags will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "is-ipv6",
				Short:      `Defines whether to filter only for IPv4s or IPv6s`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "resource-name",
				Short:      `Attached resource name to filter for, only IPs attached to a resource with this string within their name will be returned.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "resource-types.{index}",
				Short:      `Resource types to filter for. Only IPs attached to these types of resources will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "custom", "instance_server", "instance_ip", "instance_private_nic", "lb_server", "fip_ip", "vpc_gateway", "vpc_gateway_network", "k8s_node", "k8s_cluster", "rdb_instance", "redis_cluster", "baremetal_server", "baremetal_private_nic", "llm_deployment", "mgdb_instance"},
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID to filter for. Only IPs belonging to this Organization will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*ipam.ListIPsRequest)

			client := core.ExtractClient(ctx)
			api := ipam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListIPs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.IPs, nil

		},
	}
}

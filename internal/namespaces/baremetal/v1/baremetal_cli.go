// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package baremetal

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		baremetalRoot(),
		baremetalServer(),
		baremetalOffer(),
		baremetalOs(),
		baremetalBmc(),
		baremetalOptions(),
		baremetalSettings(),
		baremetalPartitioningSchemas(),
		baremetalPrivateNetwork(),
		baremetalServerList(),
		baremetalServerGet(),
		baremetalServerCreate(),
		baremetalServerUpdate(),
		baremetalServerInstall(),
		baremetalServerGetMetrics(),
		baremetalServerDelete(),
		baremetalServerReboot(),
		baremetalServerStart(),
		baremetalServerStop(),
		baremetalServerListEvents(),
		baremetalBmcStart(),
		baremetalBmcGet(),
		baremetalBmcStop(),
		baremetalServerUpdateIP(),
		baremetalOptionsAdd(),
		baremetalOptionsDelete(),
		baremetalOfferList(),
		baremetalOfferGet(),
		baremetalOptionsGet(),
		baremetalOptionsList(),
		baremetalSettingsList(),
		baremetalSettingsUpdate(),
		baremetalOsList(),
		baremetalOsGet(),
		baremetalPrivateNetworkAdd(),
		baremetalPrivateNetworkSet(),
		baremetalPrivateNetworkList(),
		baremetalPrivateNetworkDelete(),
	)
}
func baremetalRoot() *core.Command {
	return &core.Command{
		Short:     `Elastic Metal API`,
		Long:      `Elastic Metal API.`,
		Namespace: "baremetal",
	}
}

func baremetalServer() *core.Command {
	return &core.Command{
		Short:     `Server management commands`,
		Long:      `A server is a denomination of a type of instances provided by Scaleway.`,
		Namespace: "baremetal",
		Resource:  "server",
	}
}

func baremetalOffer() *core.Command {
	return &core.Command{
		Short: `Server offer management commands`,
		Long: `Server offers will answer with all different Elastic Metal server ranges available in a  zone.
Each of them will contain all the features of the server (CPUs, memory, disks) with their associated pricing.`,
		Namespace: "baremetal",
		Resource:  "offer",
	}
}

func baremetalOs() *core.Command {
	return &core.Command{
		Short:     `Operating System (OS) management commands`,
		Long:      `An Operating System (OS) is the underlying software installed on your server.`,
		Namespace: "baremetal",
		Resource:  "os",
	}
}

func baremetalBmc() *core.Command {
	return &core.Command{
		Short: `Baseboard Management Controller (BMC) management commands`,
		Long: `A Baseboard Management Controller (BMC) allows you to remotely access the low-level parameters of your dedicated server.
For instance, your KVM-IP management console could be accessed with it.
You need first to create an Remote Access option. You will find the ID and the price with a call to listOffers (https://developers.scaleway.com/en/products/baremetal/api/#get-78db92). Then you can add the option https://developers.scaleway.com/en/products/baremetal/api/#post-b14abd. Do not forget to delete the Option.
Then you need to create Remote Access https://developers.scaleway.com/en/products/baremetal/api/#post-1af723.
And finally Get Remote Access to get the login/password https://developers.scaleway.com/en/products/baremetal/api/#get-cefc0f.`,
		Namespace: "baremetal",
		Resource:  "bmc",
	}
}

func baremetalOptions() *core.Command {
	return &core.Command{
		Short:     `Server options management commands`,
		Long:      `A Server has additional options that let you personalize it to better fit your needs.`,
		Namespace: "baremetal",
		Resource:  "options",
	}
}

func baremetalSettings() *core.Command {
	return &core.Command{
		Short:     `Settings management commands`,
		Long:      `Allows to configure the general settings for your Elastic Metal server.`,
		Namespace: "baremetal",
		Resource:  "settings",
	}
}

func baremetalPartitioningSchemas() *core.Command {
	return &core.Command{
		Short:     `Partitioning-schemas management commands`,
		Long:      `Allows to customize the partitioning schemas of your servers (available on some offers and OSs).`,
		Namespace: "baremetal",
		Resource:  "partitioning-schemas",
	}
}

func baremetalPrivateNetwork() *core.Command {
	return &core.Command{
		Short: `Private Network management command`,
		Long: `A Private Network allows you to interconnect your resources
(servers, instances, ...) in an isolated and private
network. The network reachability is limited to the
resources that are on the same Private Network .  A VLAN
interface is available on the server and can be freely
managed (adding IP addresses, shutdown interface...).

Note that a resource can be a part of multiple Private Networks.`,
		Namespace: "baremetal",
		Resource:  "private-network",
	}
}

func baremetalServerList() *core.Command {
	return &core.Command{
		Short:     `List Elastic Metal servers for an Organization`,
		Long:      `List Elastic Metal servers for a specific Organization.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.ListServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the servers`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "status.{index}",
				Short:      `Status to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Names to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Project ID to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "option-id",
				Short:      `Option ID to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.ListServersRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListServers(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Servers, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all servers in your default zone",
				ArgsJSON: `null`,
			},
		},
	}
}

func baremetalServerGet() *core.Command {
	return &core.Command{
		Short:     `Get a specific Elastic Metal server`,
		Long:      `Get full details of an existing Elastic Metal server associated with the ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.GetServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.GetServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.GetServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get a specific server",
				ArgsJSON: `{"server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func baremetalServerCreate() *core.Command {
	return &core.Command{
		Short:     `Create an Elastic Metal server`,
		Long:      `Create a new Elastic Metal server. Once the server is created, proceed with the [installation of an OS](#post-3e949e).`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.CreateServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "offer-id",
				Short:      `Offer ID of the new server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "name",
				Short:      `Name of the server (≠hostname)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description associated with the server, max 255 characters`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags to associate to the server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.os-id",
				Short:      `ID of the OS to installation on the server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.hostname",
				Short:      `Hostname of the server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.ssh-key-ids.{index}",
				Short:      `SSH key IDs authorized on the server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.user",
				Short:      `User for the installation`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.password",
				Short:      `Password for the installation`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.service-user",
				Short:      `Regular user that runs the service to be installed on the server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.service-password",
				Short:      `Password used for the service to install`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.disks.{index}.device",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.disks.{index}.partitions.{index}.label",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_partition_label", "uefi", "legacy", "root", "boot", "swap", "data", "home", "raid", "zfs"},
			},
			{
				Name:       "install.partitioning-schema.disks.{index}.partitions.{index}.number",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.disks.{index}.partitions.{index}.size",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.raids.{index}.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.raids.{index}.level",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_raid_level", "raid_level_0", "raid_level_1", "raid_level_5", "raid_level_6", "raid_level_10"},
			},
			{
				Name:       "install.partitioning-schema.raids.{index}.devices.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.filesystems.{index}.device",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.filesystems.{index}.format",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_format", "fat32", "ext4", "swap", "zfs", "xfs"},
			},
			{
				Name:       "install.partitioning-schema.filesystems.{index}.mountpoint",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.zfs.pools.{index}.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.zfs.pools.{index}.type",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "no_raid", "mirror", "raidz1", "raidz2"},
			},
			{
				Name:       "install.partitioning-schema.zfs.pools.{index}.devices.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.zfs.pools.{index}.options.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "install.partitioning-schema.zfs.pools.{index}.filesystem-options.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "option-ids.{index}",
				Short:      `IDs of options to enable on server`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.CreateServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.CreateServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Create a default Elastic Metal server",
				ArgsJSON: `null`,
			},
		},
	}
}

func baremetalServerUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an Elastic Metal server`,
		Long:      `Update the server associated with the ID. You can update parameters such as the server's name, tags and description. Any parameters left null in the request body are not updated.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.UpdateServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the server (≠hostname), not updated if null`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description associated with the server, max 255 characters, not updated if null`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags associated with the server, not updated if null`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.UpdateServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.UpdateServer(request)

		},
	}
}

func baremetalServerInstall() *core.Command {
	return &core.Command{
		Short:     `Install an Elastic Metal server`,
		Long:      `Install an Operating System (OS) on the Elastic Metal server with a specific ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "install",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.InstallServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `Server ID to install`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "os-id",
				Short:      `ID of the OS to installation on the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "hostname",
				Short:      `Hostname of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ssh-key-ids.{index}",
				Short:      `SSH key IDs authorized on the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user",
				Short:      `User used for the installation`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "password",
				Short:      `Password used for the installation`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "service-user",
				Short:      `User used for the service to install`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "service-password",
				Short:      `Password used for the service to install`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.disks.{index}.device",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.disks.{index}.partitions.{index}.label",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_partition_label", "uefi", "legacy", "root", "boot", "swap", "data", "home", "raid", "zfs"},
			},
			{
				Name:       "partitioning-schema.disks.{index}.partitions.{index}.number",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.disks.{index}.partitions.{index}.size",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.raids.{index}.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.raids.{index}.level",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_raid_level", "raid_level_0", "raid_level_1", "raid_level_5", "raid_level_6", "raid_level_10"},
			},
			{
				Name:       "partitioning-schema.raids.{index}.devices.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.filesystems.{index}.device",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.filesystems.{index}.format",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_format", "fat32", "ext4", "swap", "zfs", "xfs"},
			},
			{
				Name:       "partitioning-schema.filesystems.{index}.mountpoint",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.zfs.pools.{index}.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.zfs.pools.{index}.type",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "no_raid", "mirror", "raidz1", "raidz2"},
			},
			{
				Name:       "partitioning-schema.zfs.pools.{index}.devices.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.zfs.pools.{index}.options.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "partitioning-schema.zfs.pools.{index}.filesystem-options.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.InstallServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.InstallServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Install an OS on a  server with a particular SSH key ID",
				ArgsJSON: `{"os_id":"11111111-1111-1111-1111-111111111111","server_id":"11111111-1111-1111-1111-111111111111","ssh_key_ids":["11111111-1111-1111-1111-111111111111"]}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam ssh-key list",
				Short:   "List all SSH keys",
			},
			{
				Command: "scw baremetal os list",
				Short:   "List OS (useful to get all OS IDs)",
			},
			{
				Command: "scw baremetal server create",
				Short:   "Create an Elastic Metal server",
			},
		},
	}
}

func baremetalServerGetMetrics() *core.Command {
	return &core.Command{
		Short:     `Return server metrics`,
		Long:      `Get the ping status of the server associated with the ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "get-metrics",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.GetServerMetricsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `Server ID to get the metrics`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.GetServerMetricsRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.GetServerMetrics(request)

		},
	}
}

func baremetalServerDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an Elastic Metal server`,
		Long:      `Delete the server associated with the ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.DeleteServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.DeleteServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.DeleteServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Delete an Elastic Metal server",
				ArgsJSON: `{"server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func baremetalServerReboot() *core.Command {
	return &core.Command{
		Short:     `Reboot an Elastic Metal server`,
		Long:      `Reboot the Elastic Metal server associated with the ID, use the ` + "`" + `boot_type` + "`" + ` ` + "`" + `rescue` + "`" + ` to reboot the server in rescue mode.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "reboot",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.RebootServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server to reboot`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "boot-type",
				Short:      `The type of boot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_boot_type", "normal", "rescue"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.RebootServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.RebootServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Reboot a server using the same os",
				ArgsJSON: `{"server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Reboot a server in rescue mode",
				ArgsJSON: `{"boot_type":"rescue","server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func baremetalServerStart() *core.Command {
	return &core.Command{
		Short:     `Start an Elastic Metal server`,
		Long:      `Start the server associated with the ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "start",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.StartServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server to start`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "boot-type",
				Short:      `The type of boot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_boot_type", "normal", "rescue"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.StartServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.StartServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Start an Elastic Metalx server",
				ArgsJSON: `{"server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Start a server in rescue mode",
				ArgsJSON: `{"boot_type":"rescue","server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func baremetalServerStop() *core.Command {
	return &core.Command{
		Short:     `Stop an Elastic Metal server`,
		Long:      `Stop the server associated with the ID. The server remains allocated to your account and all data remains on the local storage of the server.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "stop",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.StopServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server to stop`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.StopServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.StopServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Stop an Elastic Metal server",
				ArgsJSON: `{"server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func baremetalServerListEvents() *core.Command {
	return &core.Command{
		Short:     `List server events`,
		Long:      `List event (i.e. start/stop/reboot) associated to the server ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "list-events",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.ListServerEventsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server events searched`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Order of the server events`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.ListServerEventsRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListServerEvents(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Events, nil

		},
	}
}

func baremetalBmcStart() *core.Command {
	return &core.Command{
		Short: `Start BMC access`,
		Long: `Start BMC (Baseboard Management Controller) access associated with the ID.
The BMC (Baseboard Management Controller) access is available one hour after the installation of the server.
You need first to create an option Remote Access. You will find the ID and the price with a call to listOffers (https://developers.scaleway.com/en/products/baremetal/api/#get-78db92). Then add the option https://developers.scaleway.com/en/products/baremetal/api/#post-b14abd.
After adding the BMC option, you need to Get Remote Access to get the login/password https://developers.scaleway.com/en/products/baremetal/api/#get-cefc0f. Do not forget to delete the Option after use.`,
		Namespace: "baremetal",
		Resource:  "bmc",
		Verb:      "start",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.StartBMCAccessRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ip",
				Short:      `The IP authorized to connect to the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.StartBMCAccessRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.StartBMCAccess(request)

		},
	}
}

func baremetalBmcGet() *core.Command {
	return &core.Command{
		Short:     `Get BMC access`,
		Long:      `Get the BMC (Baseboard Management Controller) access associated with the ID, including the URL and login information needed to connect.`,
		Namespace: "baremetal",
		Resource:  "bmc",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.GetBMCAccessRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.GetBMCAccessRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.GetBMCAccess(request)

		},
	}
}

func baremetalBmcStop() *core.Command {
	return &core.Command{
		Short:     `Stop BMC access`,
		Long:      `Stop BMC (Baseboard Management Controller) access associated with the ID.`,
		Namespace: "baremetal",
		Resource:  "bmc",
		Verb:      "stop",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.StopBMCAccessRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.StopBMCAccessRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			e = api.StopBMCAccess(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "bmc",
				Verb:     "stop",
			}, nil
		},
	}
}

func baremetalServerUpdateIP() *core.Command {
	return &core.Command{
		Short:     `Update IP`,
		Long:      `Configure the IP address associated with the server ID and IP ID. You can use this method to set a reverse DNS for an IP address.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "update-ip",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.UpdateIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ip-id",
				Short:      `ID of the IP to update`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "reverse",
				Short:      `New reverse IP to update, not updated if null`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.UpdateIPRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.UpdateIP(request)

		},
	}
}

func baremetalOptionsAdd() *core.Command {
	return &core.Command{
		Short:     `Add server option`,
		Long:      `Add an option, such as Private Networks, to a specific server.`,
		Namespace: "baremetal",
		Resource:  "options",
		Verb:      "add",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.AddOptionServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "option-id",
				Short:      `ID of the option to add`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "expires-at",
				Short:      `Auto expire the option after this date`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.AddOptionServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.AddOptionServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Add an option, such as Private Networks, to a server",
				ArgsJSON: `{"option_id":"11111111-1111-1111-1111-111111111111","server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func baremetalOptionsDelete() *core.Command {
	return &core.Command{
		Short:     `Delete server option`,
		Long:      `Delete an option from a specific server.`,
		Namespace: "baremetal",
		Resource:  "options",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.DeleteOptionServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "option-id",
				Short:      `ID of the option to delete`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.DeleteOptionServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.DeleteOptionServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Delete an option from a server",
				ArgsJSON: `{"option_id":"11111111-1111-1111-1111-111111111111","server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func baremetalOfferList() *core.Command {
	return &core.Command{
		Short:     `List offers`,
		Long:      `List all available Elastic Metal server configurations.`,
		Namespace: "baremetal",
		Resource:  "offer",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.ListOffersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "subscription-period",
				Short:      `Subscription period type to filter offers by`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_subscription_period", "hourly", "monthly"},
			},
			{
				Name:       "name",
				Short:      `Offer name to filter offers by`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.ListOffersRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListOffers(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Offers, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all server offers in the default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "List all server offers in fr-par-1 zone",
				ArgsJSON: `{"zone":"fr-par-1"}`,
			},
		},
	}
}

func baremetalOfferGet() *core.Command {
	return &core.Command{
		Short:     `Get offer`,
		Long:      `Get details of an offer identified by its offer ID.`,
		Namespace: "baremetal",
		Resource:  "offer",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.GetOfferRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "offer-id",
				Short:      `ID of the researched Offer`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.GetOfferRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.GetOffer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get a server offer with the ID",
				ArgsJSON: `{"offer_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
		},
	}
}

func baremetalOptionsGet() *core.Command {
	return &core.Command{
		Short:     `Get option`,
		Long:      `Return specific option for the ID.`,
		Namespace: "baremetal",
		Resource:  "options",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.GetOptionRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "option-id",
				Short:      `ID of the option`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.GetOptionRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.GetOption(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get a server option with the ID",
				ArgsJSON: `{"option_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
		},
	}
}

func baremetalOptionsList() *core.Command {
	return &core.Command{
		Short:     `List options`,
		Long:      `List all options matching with filters.`,
		Namespace: "baremetal",
		Resource:  "options",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.ListOptionsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "offer-id",
				Short:      `Offer ID to filter options for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name to filter options for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.ListOptionsRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListOptions(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Options, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all server options in the default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "List all server options in fr-par-1 zone",
				ArgsJSON: `{"zone":"fr-par-1"}`,
			},
		},
	}
}

func baremetalSettingsList() *core.Command {
	return &core.Command{
		Short:     `List all settings`,
		Long:      `Return all settings for a Project ID.`,
		Namespace: "baremetal",
		Resource:  "settings",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.ListSettingsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order for items in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "project-id",
				Short:      `ID of the Project`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.ListSettingsRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListSettings(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Settings, nil

		},
	}
}

func baremetalSettingsUpdate() *core.Command {
	return &core.Command{
		Short:     `Update setting`,
		Long:      `Update a setting for a Project ID (enable or disable).`,
		Namespace: "baremetal",
		Resource:  "settings",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.UpdateSettingRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "setting-id",
				Short:      `ID of the setting`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "enabled",
				Short:      `Defines whether the setting is enabled`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.UpdateSettingRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.UpdateSetting(request)

		},
	}
}

func baremetalOsList() *core.Command {
	return &core.Command{
		Short:     `List available OSes`,
		Long:      `List all OSes that are available for installation on Elastic Metal servers.`,
		Namespace: "baremetal",
		Resource:  "os",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.ListOSRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "offer-id",
				Short:      `Offer IDs to filter OSes for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.ListOSRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListOS(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Os, nil

		},
	}
}

func baremetalOsGet() *core.Command {
	return &core.Command{
		Short:     `Get OS with an ID`,
		Long:      `Return the specific OS for the ID.`,
		Namespace: "baremetal",
		Resource:  "os",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.GetOSRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "os-id",
				Short:      `ID of the OS`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.GetOSRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.GetOS(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get a specific OS ID",
				ArgsJSON: `{}`,
			},
		},
	}
}

func baremetalPrivateNetworkAdd() *core.Command {
	return &core.Command{
		Short:     `Add a server to a Private Network`,
		Long:      `Add a server to a Private Network.`,
		Namespace: "baremetal",
		Resource:  "private-network",
		Verb:      "add",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.PrivateNetworkAPIAddServerPrivateNetworkRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `The ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `The ID of the Private Network`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
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
		Long:      `Set multiple Private Networks on a server.`,
		Namespace: "baremetal",
		Resource:  "private-network",
		Verb:      "set",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.PrivateNetworkAPISetServerPrivateNetworksRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `The ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-ids.{index}",
				Short:      `The IDs of the Private Networks`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
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
		Long:      `List the Private Networks of a server.`,
		Namespace: "baremetal",
		Resource:  "private-network",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(baremetal.PrivateNetworkAPIListServerPrivateNetworksRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `The sort order for the returned Private Networks`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc"},
			},
			{
				Name:       "server-id",
				Short:      `Filter Private Networks by server ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `Filter Private Networks by Private Network ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Filter Private Networks by Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Filter Private Networks by Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2, scw.Zone(core.AllLocalities)),
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
				Short:      `The ID of the server`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `The ID of the Private Network`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
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

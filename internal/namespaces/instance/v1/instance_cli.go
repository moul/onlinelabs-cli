// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package instance

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		instanceRoot(),
		instanceImage(),
		instanceIP(),
		instancePlacementGroup(),
		instanceSecurityGroup(),
		instanceServer(),
		instanceServerType(),
		instanceVolumeType(),
		instanceSnapshot(),
		instanceUserData(),
		instanceVolume(),
		instancePrivateNic(),
		instanceServerTypeGet(),
		instanceServerTypeList(),
		instanceVolumeTypeList(),
		instanceServerList(),
		instanceServerGet(),
		instanceServerUpdate(),
		instanceServerListActions(),
		instanceUserDataList(),
		instanceUserDataDelete(),
		instanceUserDataSet(),
		instanceUserDataGet(),
		instanceImageList(),
		instanceImageGet(),
		instanceImageCreate(),
		instanceImageUpdate(),
		instanceImageDelete(),
		instanceSnapshotList(),
		instanceSnapshotCreate(),
		instanceSnapshotGet(),
		instanceSnapshotUpdate(),
		instanceSnapshotDelete(),
		instanceSnapshotExport(),
		instanceVolumeList(),
		instanceVolumeCreate(),
		instanceVolumeGet(),
		instanceVolumeUpdate(),
		instanceVolumeDelete(),
		instanceSecurityGroupList(),
		instanceSecurityGroupCreate(),
		instanceSecurityGroupGet(),
		instanceSecurityGroupDelete(),
		instanceSecurityGroupUpdate(),
		instanceSecurityGroupListDefaultRules(),
		instanceSecurityGroupListRules(),
		instanceSecurityGroupCreateRule(),
		instanceSecurityGroupSetRules(),
		instanceSecurityGroupDeleteRule(),
		instanceSecurityGroupGetRule(),
		instanceSecurityGroupUpdateRule(),
		instancePlacementGroupList(),
		instancePlacementGroupCreate(),
		instancePlacementGroupGet(),
		instancePlacementGroupSet(),
		instancePlacementGroupUpdate(),
		instancePlacementGroupDelete(),
		instancePlacementGroupGetServers(),
		instancePlacementGroupSetServers(),
		instancePlacementGroupUpdateServers(),
		instanceIPList(),
		instanceIPCreate(),
		instanceIPGet(),
		instanceIPUpdate(),
		instanceIPDelete(),
		instancePrivateNicList(),
		instancePrivateNicCreate(),
		instancePrivateNicGet(),
		instancePrivateNicUpdate(),
		instancePrivateNicDelete(),
		instanceVolumePlanMigration(),
		instanceVolumeApplyMigration(),
	)
}
func instanceRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Instances`,
		Long:      `This API allows you to manage your Instances.`,
		Namespace: "instance",
	}
}

func instanceImage() *core.Command {
	return &core.Command{
		Short: `Image management commands`,
		Long: `Images are backups of your Instances.
One image will contain all the volumes of your Instance and can be used to restore your Instance and its data. You can also use it to create a series of Instances with a predefined configuration.
To copy not all but only one specified volume of an Instance, you can use the snapshot feature instead.`,
		Namespace: "instance",
		Resource:  "image",
	}
}

func instanceIP() *core.Command {
	return &core.Command{
		Short: `IP management commands`,
		Long: `A flexible IP address is an IP address which you hold independently of any Instance.
You can attach it to any of your Instances and do live migration of the IP address between your Instances.

Note that attaching a flexible IP address to an Instance removes its previous public IP and interrupts any ongoing public connection to the Instance. This does not apply if you have migrated your server to the new Network stack and have at least one flexible IP attached to the Instance.`,
		Namespace: "instance",
		Resource:  "ip",
	}
}

func instancePlacementGroup() *core.Command {
	return &core.Command{
		Short: `Placement group management commands`,
		Long: `Placement groups allow the user to express a preference regarding
the physical position of a group of Instances. The feature lets the user
choose to either group Instances on the same physical hardware for
best network throughput and low latency or to spread Instances across
physically distanced hardware to reduce the risk of physical failure.

The operating mode is selected by a ` + "`" + `policy_type` + "`" + `. Two policy
types are available:
  - ` + "`" + `low_latency` + "`" + ` will group Instances on the same hypervisors
  - ` + "`" + `max_availability` + "`" + ` will spread Instances across physically distanced hypervisors

The ` + "`" + `policy_type` + "`" + ` is set to ` + "`" + `max_availability` + "`" + ` by default.

For each policy types, one of the two ` + "`" + `policy_mode` + "`" + ` may be selected:
  - ` + "`" + `optional` + "`" + ` will start your Instances even if the constraint is not respected
  - ` + "`" + `enforced` + "`" + ` guarantees that if the Instance starts, the constraint is respected

The ` + "`" + `policy_mode` + "`" + ` is set by default to ` + "`" + `optional` + "`" + `.`,
		Namespace: "instance",
		Resource:  "placement-group",
	}
}

func instanceSecurityGroup() *core.Command {
	return &core.Command{
		Short: `Security group management commands`,
		Long: `A security group is a set of firewall rules on a set of Instances.
Security groups enable you to create rules that either drop or allow incoming traffic from certain ports of your Instances.

Security groups are stateful by default which means return traffic is automatically allowed, regardless of any rules.
As a contrary, you have to switch in a stateless mode to define explicitly allowed.`,
		Namespace: "instance",
		Resource:  "security-group",
	}
}

func instanceServer() *core.Command {
	return &core.Command{
		Short: `Instance management commands`,
		Long: `Instances are computing units providing resources to run your applications on.
Scaleway offers various Instance types including **Virtual Instances** and **dedicated GPU Instances**.
**Note: Instances can be referenced as "servers" in API endpoints.**`,
		Namespace: "instance",
		Resource:  "server",
	}
}

func instanceServerType() *core.Command {
	return &core.Command{
		Short: `Instance type management commands`,
		Long: `All Instance types available in a specified zone.
Each type contains all the features of the Instance (CPU, RAM, Storage) as well as their associated pricing.`,
		Namespace: "instance",
		Resource:  "server-type",
	}
}

func instanceVolumeType() *core.Command {
	return &core.Command{
		Short: `Volume type management commands`,
		Long: `All volume types available in a specified zone.
Each of these types will contains all the capabilities and constraints of the volume (min size, max size, snapshot).`,
		Namespace: "instance",
		Resource:  "volume-type",
	}
}

func instanceSnapshot() *core.Command {
	return &core.Command{
		Short: `Snapshot management commands`,
		Long: `Snapshots contain the data of a specified volume at a particular point in time.
The data can include the Instance's operating system,
configuration information and/or files stored on the volume.

A snapshot can be done from a specified volume, e.g. you
have one Instance with a volume containing the OS and another one
containing the application data, and you want to use different
snapshot strategies on both volumes.

A snapshot's volume type is its original volume's type (` + "`" + `l_ssd` + "`" + ` or ` + "`" + `b_ssd` + "`" + `).
Volumes can be created from snapshots of their own type.`,
		Namespace: "instance",
		Resource:  "snapshot",
	}
}

func instanceUserData() *core.Command {
	return &core.Command{
		Short: `User data management commands`,
		Long: `User data is a key/value store you can use to provide your instance with introspective data.

As an example of use, Scaleway images contain the ` + "`" + `scw-generate-ssh-keys` + "`" + ` script, which generates the SSH server’s host keys, then stores their fingerprints as user data under the ` + "`" + `ssh-host-fingerprints` + "`" + ` key.
This way, before connecting to the instance using SSH, one can query the fingerprints to ensure that the remote host is the expected one, thus avoiding any man-in-the-middle attack.

There are two ways of accessing user data:
 - **From within a running Instance**, by requesting the Metadata API at http://169.254.42.42/user_data (or http://[fd00:42::42]/user_data using IPv6).
   The ` + "`" + `scaleway-ecosystem` + "`" + ` package, installed by default on all OS images provided by Scaleway, ships with the ` + "`" + `scw-userdata` + "`" + ` helper command that allows you to easily query the user data from the instance.
   For security reasons, viewing and editing user data is only allowed to queries originating from a port below 1024 (by default, only the super-user can bind to ports below 1024).
   To specify the source port with cURL, use the ` + "`" + `--local-port` + "`" + ` option (e.g. ` + "`" + `curl --local-port 1-1023 http://169.254.42.42/user_data` + "`" + `).
 - **From the Instance API** by using the methods described below.`,
		Namespace: "instance",
		Resource:  "user-data",
	}
}

func instanceVolume() *core.Command {
	return &core.Command{
		Short: `Volume management commands`,
		Long: `A volume is where you store your data inside your Instance. It
appears as a block device on Linux that you can use to create
a filesystem and mount it.

Two different types of volume (` + "`" + `volume_type` + "`" + `) are available:
  - ` + "`" + `l_ssd` + "`" + ` is a local block storage: your data is downloaded on
    the hypervisor and you need to power off your Instance to attach
    or detach a volume.
  - ` + "`" + `b_ssd` + "`" + ` is a remote block storage: your data is stored on a
    centralized cluster. You can plug and unplug a volume while
    your Instance is running.

Minimum and maximum volume sizes for each volume types can be queried
from the zone ` + "`" + `/products/volumes` + "`" + ` API endpoint. _I.e_ for:
  - ` + "`" + `fr-par-1` + "`" + `  use https://api.scaleway.com/instance/v1/zones/fr-par-1/products/volumes
  - ` + "`" + `nl-ams-1` + "`" + `  use https://api.scaleway.com/instance/v1/zones/nl-ams-1/products/volumes

Each type of volume is also subject to a global quota for the sum of all the
volumes. This quota depends of the level of support and may be
changed on demand.

Be wary that when terminating an Instance, if you want to keep
your block storage volume, **you must** detach it before you
issue the ` + "`" + `terminate` + "`" + ` call.

When using multiple block devices, it's advised to mount them by
using their UUID instead of their device name. A device name is
subject to change depending on the volumes order. Block devices
UUIDs can be found in ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
		Namespace: "instance",
		Resource:  "volume",
	}
}

func instancePrivateNic() *core.Command {
	return &core.Command{
		Short: `Private NIC management commands`,
		Long: `A Private NIC is the network interface that connects an Instance to a
Private Network. An Instance can have multiple private NICs at the same
time, but each NIC must belong to a different Private Network.`,
		Namespace: "instance",
		Resource:  "private-nic",
	}
}

func instanceServerTypeGet() *core.Command {
	return &core.Command{
		Short:     `Get availability`,
		Long:      `Get availability for all Instance types.`,
		Namespace: "instance",
		Resource:  "server-type",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetServerTypesAvailabilityRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetServerTypesAvailabilityRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetServerTypesAvailability(request)

		},
	}
}

func instanceServerTypeList() *core.Command {
	return &core.Command{
		Short:     `List Instance types`,
		Long:      `List available Instance types and their technical details.`,
		Namespace: "instance",
		Resource:  "server-type",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListServersTypesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListServersTypesRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.ListServersTypes(request)

		},
		Examples: []*core.Example{
			{
				Short:    "List all server-types in the default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "List all server-types in fr-par-1 zone",
				ArgsJSON: `{"zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceVolumeTypeList() *core.Command {
	return &core.Command{
		Short:     `List volume types`,
		Long:      `List all volume types and their technical details.`,
		Namespace: "instance",
		Resource:  "volume-type",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListVolumesTypesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListVolumesTypesRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.ListVolumesTypes(request)

		},
		Examples: []*core.Example{
			{
				Short:    "List all volume-types in the default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "List all volume-types in fr-par-1 zone",
				ArgsJSON: `{"zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceServerList() *core.Command {
	return &core.Command{
		Short:     `List all Instances`,
		Long:      `List all Instances in a specified Availability Zone, e.g. ` + "`" + `fr-par-1` + "`" + `.`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "project",
				Short:      `List only Instances of this Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Filter Instances by name (eg. "server1" will return "server100" and "server1" but not "foo")`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-ip",
				Short:      `List Instances by private_ip`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "without-ip",
				Short:      `List Instances that are not attached to a public IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "with-ip",
				Short:      `List Instances by IP (both private_ip and public_ip are supported)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "commercial-type",
				Short:      `List Instances of this commercial type`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "state",
				Short:      `List Instances in this state`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"running", "stopped", "stopped in place", "starting", "stopping", "locked"},
			},
			{
				Name:       "tags",
				Short:      `List Instances with these exact tags (to filter with several tags, use commas to separate them)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network",
				Short:      `List Instances in this Private Network`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order",
				Short:      `Define the order of the returned servers`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"creation_date_desc", "creation_date_asc", "modification_date_desc", "modification_date_asc"},
			},
			{
				Name:       "private-networks",
				Short:      `List Instances from the given Private Networks (use commas to separate them)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-nic-mac-address",
				Short:      `List Instances associated with the given private NIC MAC address`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "servers",
				Short:      `List Instances from these server ids (use commas to separate them)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization",
				Short:      `List only Instances of this Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListServersRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
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
				Short:    "List all Instances on your default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "List Instances of this commercial type",
				ArgsJSON: `{"commercial_type":"DEV1-S"}`,
			},
			{
				Short:    "List Instances that are not attached to a public IP",
				ArgsJSON: `{"without_ip":true}`,
			},
			{
				Short:    "List Instances that match the specified name ('server1' will return 'server100' and 'server1' but not 'foo')",
				ArgsJSON: `{"name":"server1"}`,
			},
		},
	}
}

func instanceServerGet() *core.Command {
	return &core.Command{
		Short:     `Get an Instance`,
		Long:      `Get the details of a specified Instance.`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the Instance you want to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetServerRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get the Instance with its specified ID",
				ArgsJSON: `{"server_id":"94ededdf-358d-4019-9886-d754f8a2e78d"}`,
			},
		},
	}
}

func instanceServerUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an Instance`,
		Long:      `Update the Instance information, such as name, boot mode, or tags.`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdateServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the Instance`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "boot-type",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"local", "bootscript", "rescue"},
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of the Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volumes.{key}.id",
				Short:      `UUID of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volumes.{key}.boot",
				Short:      `Force the Instance to boot on this volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("false"),
			},
			{
				Name:       "volumes.{key}.name",
				Short:      `Name of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volumes.{key}.size",
				Short:      `Disk size of the volume, must be a multiple of 512`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volumes.{key}.volume-type",
				Short:      `Type of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"l_ssd", "b_ssd", "unified", "scratch", "sbs_volume", "sbs_snapshot"},
			},
			{
				Name:       "volumes.{key}.base-snapshot",
				Short:      `ID of the snapshot on which this volume will be based`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volumes.{key}.project",
				Short:      `Project ID of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volumes.{key}.organization",
				Short:      `Organization ID of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "dynamic-ip-required",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "routed-ip-enabled",
				Short:      `True to configure the instance so it uses the new routed IP mode (once this is set to True you cannot set it back to False)`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "public-ips.{index}",
				Short:      `A list of reserved IP IDs to attach to the Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "enable-ipv6",
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "protected",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "security-group.id",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "security-group.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "placement-group",
				Short:      `Placement group ID if Instance must be part of a placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-nics.{index}",
				Short:      `Instance private NICs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "commercial-type",
				Short:      `Set the commercial_type for this Instance.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "admin-password-encryption-ssh-key-id",
				Short:      `UUID of the SSH RSA key that will be used to encrypt the initial admin password for OS requiring it. Mandatory for Windows OS.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdateServerRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdateServer(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Update the name of a specified Instance",
				ArgsJSON: `{"name":"foobar","server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Switch a specified Instance to rescue mode (reboot is required to access rescue mode)",
				ArgsJSON: `{"boot_type":"rescue","server_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Overwrite tags of a specified Instance",
				ArgsJSON: `{"server_id":"11111111-1111-1111-1111-111111111111","tags":["foo","bar"]}`,
			},
			{
				Short: "Apply the specified security group to a specified server",
				Raw:   `scw instance server server update 11111111-1111-1111-1111-111111111111 security-group-id=11111111-1111-1111-1111-111111111111`,
			},
			{
				Short: "Put a specified Instance in the specified placement group. Instance must be off",
				Raw:   `scw instance server server update 11111111-1111-1111-1111-111111111111 placement-group-id=11111111-1111-1111-1111-111111111111`,
			},
		},
	}
}

func instanceServerListActions() *core.Command {
	return &core.Command{
		Short:     `List Instance actions`,
		Long:      `List all actions (e.g. power on, power off, reboot) that can currently be performed on an Instance.`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "list-actions",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListServerActionsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListServerActionsRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.ListServerActions(request)

		},
	}
}

func instanceUserDataList() *core.Command {
	return &core.Command{
		Short:     `List user data`,
		Long:      `List all user data keys registered on a specified Instance.`,
		Namespace: "instance",
		Resource:  "user-data",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListServerUserDataRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListServerUserDataRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.ListServerUserData(request)

		},
	}
}

func instanceUserDataDelete() *core.Command {
	return &core.Command{
		Short:     `Delete user data`,
		Long:      `Delete the specified key from an Instance's user data.`,
		Namespace: "instance",
		Resource:  "user-data",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeleteServerUserDataRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "key",
				Short:      `Key of the user data to delete`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeleteServerUserDataRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeleteServerUserData(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "user-data",
				Verb:     "delete",
			}, nil
		},
	}
}

func instanceUserDataSet() *core.Command {
	return &core.Command{
		Short:     `Add/set user data`,
		Long:      `Add or update a user data with the specified key on an Instance.`,
		Namespace: "instance",
		Resource:  "user-data",
		Verb:      "set",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.SetServerUserDataRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "key",
				Short:      `Key of the user data to set`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "content.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "content.content-type",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "content.content",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.SetServerUserDataRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.SetServerUserData(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "user-data",
				Verb:     "set",
			}, nil
		},
	}
}

func instanceUserDataGet() *core.Command {
	return &core.Command{
		Short:     `Get user data`,
		Long:      `Get the content of a user data with the specified key on an Instance.`,
		Namespace: "instance",
		Resource:  "user-data",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetServerUserDataRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "key",
				Short:      `Key of the user data to get`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetServerUserDataRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetServerUserData(request)

		},
	}
}

func instanceImageList() *core.Command {
	return &core.Command{
		Short:     `List Instance images`,
		Long:      `List all existing Instance images.`,
		Namespace: "instance",
		Resource:  "image",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListImagesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "public",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "arch",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListImagesRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListImages(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Images, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all public images in the default zone",
				ArgsJSON: `null`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw marketplace image list",
				Short:   "List marketplace images",
			},
		},
	}
}

func instanceImageGet() *core.Command {
	return &core.Command{
		Short:     `Get an Instance image`,
		Long:      `Get details of an image with the specified ID.`,
		Namespace: "instance",
		Resource:  "image",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetImageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "image-id",
				Short:      `UUID of the image you want to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetImageRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetImage(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get an image in the default zone with the specified ID",
				ArgsJSON: `{"image_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Get an image in fr-par-1 zone with the specified ID",
				ArgsJSON: `{"image_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceImageCreate() *core.Command {
	return &core.Command{
		Short:     `Create an Instance image`,
		Long:      `Create an Instance image from the specified snapshot ID.`,
		Namespace: "instance",
		Resource:  "image",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.CreateImageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the image`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("img"),
			},
			{
				Name:       "root-volume",
				Short:      `UUID of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "arch",
				Short:      `Architecture of the image`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_arch", "x86_64", "arm", "arm64"},
			},
			{
				Name:       "extra-volumes.{key}.id",
				Short:      `UUID of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "extra-volumes.{key}.name",
				Short:      `Name of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "extra-volumes.{key}.size",
				Short:      `Disk size of the volume, must be a multiple of 512`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "extra-volumes.{key}.volume-type",
				Short:      `Type of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"l_ssd", "b_ssd", "unified", "scratch", "sbs_volume", "sbs_snapshot"},
			},
			{
				Name:       "extra-volumes.{key}.project",
				Short:      `Project ID of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "extra-volumes.{key}.organization",
				Short:      `Organization ID of the volume`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			core.ProjectArgSpec(),
			{
				Name:       "tags.{index}",
				Short:      `Tags of the image`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "public",
				Short:      `True to create a public image`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.CreateImageRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.CreateImage(request)

		},
		Examples: []*core.Example{
			{
				Short: "Create an image named 'foobar' for x86_64 Instances from the specified snapshot ID",
				Raw:   `scw instance server image create name=foobar snapshot-id=11111111-1111-1111-1111-111111111111 arch=x86_64`,
			},
		},
	}
}

func instanceImageUpdate() *core.Command {
	return &core.Command{
		Short:     `Update image`,
		Long:      `Update the properties of an image.`,
		Namespace: "instance",
		Resource:  "image",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdateImageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "image-id",
				Short:      `UUID of the image`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the image`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "arch",
				Short:      `Architecture of the image`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_arch", "x86_64", "arm", "arm64"},
			},
			{
				Name:       "extra-volumes.{key}.id",
				Short:      `UUID of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of the image`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "public",
				Short:      `True to set the image as public`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdateImageRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdateImage(request)

		},
	}
}

func instanceImageDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an Instance image`,
		Long:      `Delete the image with the specified ID.`,
		Namespace: "instance",
		Resource:  "image",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeleteImageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "image-id",
				Short:      `UUID of the image you want to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeleteImageRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeleteImage(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "image",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete an image in the default zone with the specified ID",
				ArgsJSON: `{"image_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Delete an image in fr-par-1 zone with the specified ID",
				ArgsJSON: `{"image_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceSnapshotList() *core.Command {
	return &core.Command{
		Short:     `List snapshots`,
		Long:      `List all snapshots of an Organization in a specified Availability Zone.`,
		Namespace: "instance",
		Resource:  "snapshot",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListSnapshotsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "project",
				Short:      `List snapshots only for this Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `List snapshots of the requested name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags",
				Short:      `List snapshots that have the requested tag`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "base-volume-id",
				Short:      `List snapshots originating only from this volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization",
				Short:      `List snapshots only for this Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListSnapshotsRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListSnapshots(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Snapshots, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all snapshots in the default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "List all snapshots in fr-par-1 zone",
				ArgsJSON: `{"zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceSnapshotCreate() *core.Command {
	return &core.Command{
		Short:     `Create a snapshot from a specified volume or from a QCOW2 file`,
		Long:      `Create a snapshot from a specified volume or from a QCOW2 file in a specified Availability Zone.`,
		Namespace: "instance",
		Resource:  "snapshot",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.CreateSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("snp"),
			},
			{
				Name:       "volume-id",
				Short:      `UUID of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectArgSpec(),
			{
				Name:       "volume-type",
				Short:      `Volume type of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_volume_type", "l_ssd", "b_ssd", "unified"},
			},
			{
				Name:       "bucket",
				Short:      `Bucket name for snapshot imports`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "key",
				Short:      `Object key for snapshot imports`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "size",
				Short:      `Imported snapshot size, must be a multiple of 512`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.CreateSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.CreateSnapshot(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Create a snapshot in the default zone from the specified volume ID",
				ArgsJSON: `{"volume_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Create a snapshot in fr-par-1 zone from the specified volume ID",
				ArgsJSON: `{"volume_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
			{
				Short:    "Create a named snapshot from the specified volume ID",
				ArgsJSON: `{"name":"foobar","volume_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Import a QCOW file as an Instance snapshot",
				ArgsJSON: `{"bucket":"my-bucket","key":"my-qcow2-file-name","name":"my-imported-snapshot","volume_type":"b_ssd","zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceSnapshotGet() *core.Command {
	return &core.Command{
		Short:     `Get a snapshot`,
		Long:      `Get details of a snapshot with the specified ID.`,
		Namespace: "instance",
		Resource:  "snapshot",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot you want to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetSnapshot(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get a snapshot in the default zone with the specified ID",
				ArgsJSON: `{"snapshot_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Get a snapshot in fr-par-1 zone with the specified ID",
				ArgsJSON: `{"snapshot_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceSnapshotUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a snapshot`,
		Long:      `Update the properties of a snapshot.`,
		Namespace: "instance",
		Resource:  "snapshot",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdateSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdateSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdateSnapshot(request)

		},
	}
}

func instanceSnapshotDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a snapshot`,
		Long:      `Delete the snapshot with the specified ID.`,
		Namespace: "instance",
		Resource:  "snapshot",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeleteSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot you want to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeleteSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeleteSnapshot(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "snapshot",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete a snapshot in the default zone with the specified ID",
				ArgsJSON: `{"snapshot_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Delete a snapshot in fr-par-1 zone with the specified ID",
				ArgsJSON: `{"snapshot_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceSnapshotExport() *core.Command {
	return &core.Command{
		Short:     `Export a snapshot`,
		Long:      `Export a snapshot to a specified Object Storage bucket in the same region.`,
		Namespace: "instance",
		Resource:  "snapshot",
		Verb:      "export",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ExportSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "bucket",
				Short:      `Object Storage bucket name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "key",
				Short:      `Object key`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "snapshot-id",
				Short:      `Snapshot ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ExportSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.ExportSnapshot(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Export a snapshot to an Object Storage bucket",
				ArgsJSON: `{"bucket":"my-bucket","key":"my-qcow2-file-name","snapshot_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
		},
	}
}

func instanceVolumeList() *core.Command {
	return &core.Command{
		Short:     `List volumes`,
		Long:      `List volumes in the specified Availability Zone. You can filter the output by volume type.`,
		Namespace: "instance",
		Resource:  "volume",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListVolumesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-type",
				Short:      `Filter by volume type`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"l_ssd", "b_ssd", "unified", "scratch", "sbs_volume", "sbs_snapshot"},
			},
			{
				Name:       "project",
				Short:      `Filter volume by Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags",
				Short:      `Filter volumes with these exact tags (to filter with several tags, use commas to separate them)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Filter volume by name (for eg. "vol" will return "myvolume" but not "data")`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization",
				Short:      `Filter volume by Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListVolumesRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListVolumes(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Volumes, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all volumes",
				ArgsJSON: `null`,
			},
			{
				Short:    "List all block storage volumes",
				ArgsJSON: `{"volume_type":"b_ssd"}`,
			},
			{
				Short:    "List all local storage volumes",
				ArgsJSON: `{"volume_type":"l_ssd"}`,
			},
			{
				Short:    "List all volumes that match a name",
				ArgsJSON: `{"name":"foobar"}`,
			},
			{
				Short:    "List all block storage volumes that match a name",
				ArgsJSON: `{"name":"foobar","volume_type":"b_ssd"}`,
			},
		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "State",
			},
			{
				FieldName: "Server.ID",
			},
			{
				FieldName: "Server.Name",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "Project",
			},
			{
				FieldName: "Size",
			},
			{
				FieldName: "VolumeType",
			},
			{
				FieldName: "CreationDate",
			},
			{
				FieldName: "ModificationDate",
			},
			{
				FieldName: "ExportURI",
			},
			{
				FieldName: "Organization",
			},
		}},
	}
}

func instanceVolumeCreate() *core.Command {
	return &core.Command{
		Short:     `Create a volume`,
		Long:      `Create a volume of a specified type in an Availability Zone.`,
		Namespace: "instance",
		Resource:  "volume",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.CreateVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Volume name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("vol"),
			},
			core.ProjectArgSpec(),
			{
				Name:       "tags.{index}",
				Short:      `Volume tags`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volume-type",
				Short:      `Volume type`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"l_ssd", "b_ssd", "unified", "scratch", "sbs_volume", "sbs_snapshot"},
			},
			{
				Name:       "size",
				Short:      `Volume disk size, must be a multiple of 512`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "base-snapshot",
				Short:      `ID of the snapshot on which this volume will be based`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.CreateVolumeRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.CreateVolume(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Create a volume called 'my-volume'",
				ArgsJSON: `{"name":"my-volume"}`,
			},
			{
				Short:    "Create a volume with a size of 50GB",
				ArgsJSON: `{"size":50000000000}`,
			},
		},
	}
}

func instanceVolumeGet() *core.Command {
	return &core.Command{
		Short:     `Get a volume`,
		Long:      `Get details of a volume with the specified ID.`,
		Namespace: "instance",
		Resource:  "volume",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-id",
				Short:      `UUID of the volume you want to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetVolumeRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetVolume(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get a volume with the specified ID",
				ArgsJSON: `{"volume_id":"b70e9a0e-28b1-4542-bb9b-06d2d6debc0f"}`,
			},
		},
	}
}

func instanceVolumeUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a volume`,
		Long:      `Replace the name and/or size properties of a volume specified by its ID, with the specified value(s). Any volume name can be changed, however only ` + "`" + `b_ssd` + "`" + ` volumes can currently be increased in size.`,
		Namespace: "instance",
		Resource:  "volume",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdateVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-id",
				Short:      `UUID of the volume`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Volume name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "size",
				Short:      `Volume disk size, must be a multiple of 512`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdateVolumeRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdateVolume(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Change the volume name",
				ArgsJSON: `{"name":"my-new-name","volume_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Change the volume disk size (bytes)",
				ArgsJSON: `{"size":60000000000,"volume_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Change the volume name and disk size",
				ArgsJSON: `{"name":"a-new-name","size":70000000000,"volume_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func instanceVolumeDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a volume`,
		Long:      `Delete the volume with the specified ID.`,
		Namespace: "instance",
		Resource:  "volume",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeleteVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-id",
				Short:      `UUID of the volume you want to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeleteVolumeRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeleteVolume(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "volume",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete a volume with the specified ID",
				ArgsJSON: `{"volume_id":"af136619-bc59-4b48-a0ed-ed7dceaad9a6"}`,
			},
		},
	}
}

func instanceSecurityGroupList() *core.Command {
	return &core.Command{
		Short:     `List security groups`,
		Long:      `List all existing security groups.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListSecurityGroupsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the security group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project",
				Short:      `Security group Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags",
				Short:      `List security groups with these exact tags (to filter with several tags, use commas to separate them)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-default",
				Short:      `Filter security groups with this value for project_default`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization",
				Short:      `Security group Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListSecurityGroupsRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListSecurityGroups(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.SecurityGroups, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all security groups that match the specified name",
				ArgsJSON: `{"name":"foobar"}`,
			},
		},
	}
}

func instanceSecurityGroupCreate() *core.Command {
	return &core.Command{
		Short:     `Create a security group`,
		Long:      `Create a security group with a specified name and description.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.CreateSecurityGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the security group`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("sg"),
			},
			{
				Name:       "description",
				Short:      `Description of the security group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectArgSpec(),
			{
				Name:       "tags.{index}",
				Short:      `Tags of the security group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-default",
				Short:      `Defines whether this security group becomes the default security group for new Instances`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "project-default",
				Short:      `Whether this security group becomes the default security group for new Instances`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "stateful",
				Short:      `Whether the security group is stateful or not`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("true"),
			},
			{
				Name:       "inbound-default-policy",
				Short:      `Default policy for inbound rules`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("accept"),
				EnumValues: []string{"unknown_policy", "accept", "drop"},
			},
			{
				Name:       "outbound-default-policy",
				Short:      `Default policy for outbound rules`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("accept"),
				EnumValues: []string{"unknown_policy", "accept", "drop"},
			},
			{
				Name:       "enable-default-security",
				Short:      `True to block SMTP on IPv4 and IPv6. This feature is read only, please open a support ticket if you need to make it configurable`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.CreateSecurityGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.CreateSecurityGroup(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Create a security group with a specified name and description",
				ArgsJSON: `{"description":"foobar foobar","name":"foobar"}`,
			},
			{
				Short:    "Create a security group that will be applied as default on all Instances of this Project",
				ArgsJSON: `{"project_default":true}`,
			},
			{
				Short:    "Create a security group that will have a default drop inbound policy (traffic your Instance receives)",
				ArgsJSON: `{"inbound_default_policy":"drop"}`,
			},
			{
				Short:    "Create a security group that will have a default drop outbound policy (traffic your Instance transmits)",
				ArgsJSON: `{"outbound_default_policy":"drop"}`,
			},
			{
				Short:    "Create a stateless security group",
				ArgsJSON: `{"stateful":false}`,
			},
		},
	}
}

func instanceSecurityGroupGet() *core.Command {
	return &core.Command{
		Short:     `Get a security group`,
		Long:      `Get the details of a security group with the specified ID.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetSecurityGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Short:      `UUID of the security group you want to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetSecurityGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetSecurityGroup(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get a security group with the specified ID",
				ArgsJSON: `{"security_group_id":"a3244331-5d32-4e36-9bf9-b60233e201c7"}`,
			},
		},
	}
}

func instanceSecurityGroupDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a security group`,
		Long:      `Delete a security group with the specified ID.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeleteSecurityGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Short:      `UUID of the security group you want to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeleteSecurityGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeleteSecurityGroup(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "security-group",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete the security group with the specified ID",
				ArgsJSON: `{"security_group_id":"69e17c83-9945-47ac-8b29-8c1ad050ee83"}`,
			},
		},
	}
}

func instanceSecurityGroupUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a security group`,
		Long:      `Update the properties of security group.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdateSecurityGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Short:      `UUID of the security group`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the security group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description of the security group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "enable-default-security",
				Short:      `True to block SMTP on IPv4 and IPv6. This feature is read only, please open a support ticket if you need to make it configurable`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "inbound-default-policy",
				Short:      `Default inbound policy`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_policy", "accept", "drop"},
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of the security group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-default",
				Short:      `Please use project_default instead`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "project-default",
				Short:      `True use this security group for future Instances created in this project`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "outbound-default-policy",
				Short:      `Default outbound policy`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_policy", "accept", "drop"},
			},
			{
				Name:       "stateful",
				Short:      `True to set the security group as stateful`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdateSecurityGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdateSecurityGroup(request)

		},
	}
}

func instanceSecurityGroupListDefaultRules() *core.Command {
	return &core.Command{
		Short:     `Get default rules`,
		Long:      `Lists the default rules applied to all the security groups.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "list-default-rules",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListDefaultSecurityGroupRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListDefaultSecurityGroupRulesRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.ListDefaultSecurityGroupRules(request)

		},
	}
}

func instanceSecurityGroupListRules() *core.Command {
	return &core.Command{
		Short:     `List rules`,
		Long:      `List the rules of the a specified security group ID.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "list-rules",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListSecurityGroupRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Short:      `UUID of the security group`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListSecurityGroupRulesRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListSecurityGroupRules(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Rules, nil

		},
	}
}

func instanceSecurityGroupCreateRule() *core.Command {
	return &core.Command{
		Short:     `Create rule`,
		Long:      `Create a rule in the specified security group ID.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "create-rule",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.CreateSecurityGroupRuleRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Short:      `UUID of the security group`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "protocol",
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_protocol", "TCP", "UDP", "ICMP", "ANY"},
			},
			{
				Name:       "direction",
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_direction", "inbound", "outbound"},
			},
			{
				Name:       "action",
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_action", "accept", "drop"},
			},
			{
				Name:       "ip-range",
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("0.0.0.0/0"),
			},
			{
				Name:       "dest-port-from",
				Short:      `Beginning of the range of ports to apply this rule to (inclusive)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "dest-port-to",
				Short:      `End of the range of ports to apply this rule to (inclusive)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "position",
				Short:      `Position of this rule in the security group rules list`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "editable",
				Short:      `Indicates if this rule is editable (will be ignored)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.CreateSecurityGroupRuleRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.CreateSecurityGroupRule(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Allow incoming SSH",
				ArgsJSON: `{"action":"accept","dest_port_from":22,"direction":"inbound","protocol":"TCP","security_group_id":"1248283f-17de-464a-b03b-3f975ada3fa8"}`,
			},
			{
				Short:    "Allow HTTP",
				ArgsJSON: `{"action":"accept","dest_port_from":80,"direction":"inbound","protocol":"TCP","security_group_id":"e8ba77c1-9ccb-4c0c-b08d-555cfd7f57e4"}`,
			},
			{
				Short:    "Allow HTTPS",
				ArgsJSON: `{"action":"accept","dest_port_from":443,"direction":"inbound","protocol":"TCP","security_group_id":"e5906437-8650-4fe2-8ca7-32e1d7320c1b"}`,
			},
			{
				Short:    "Allow a specified IP range",
				ArgsJSON: `{"action":"accept","direction":"inbound","ip_range":"10.0.0.0/16","protocol":"ANY","security_group_id":"b6a58155-a2f8-48bd-9da9-3ff9783fa0d4"}`,
			},
			{
				Short:    "Allow FTP",
				ArgsJSON: `{"action":"accept","dest_port_from":20,"dest_port_to":21,"direction":"inbound","protocol":"TCP","security_group_id":"9c46df03-83c2-46fb-936c-16ecb44860e1"}`,
			},
		},
	}
}

func instanceSecurityGroupSetRules() *core.Command {
	return &core.Command{
		Short:     `Update all the rules of a security group`,
		Long:      `Replaces the existing rules of the security group with the rules provided. This endpoint supports the update of existing rules, creation of new rules and deletion of existing rules when they are not passed in the request.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "set-rules",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.SetSecurityGroupRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Short:      `UUID of the security group to update the rules on`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.id",
				Short:      `UUID of the security rule to update. If no value is provided, a new rule will be created`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.action",
				Short:      `Action to apply when the rule matches a packet`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_action", "accept", "drop"},
			},
			{
				Name:       "rules.{index}.protocol",
				Short:      `Protocol family this rule applies to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_protocol", "TCP", "UDP", "ICMP", "ANY"},
			},
			{
				Name:       "rules.{index}.direction",
				Short:      `Direction the rule applies to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_direction", "inbound", "outbound"},
			},
			{
				Name:       "rules.{index}.ip-range",
				Short:      `Range of IP addresses these rules apply to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.dest-port-from",
				Short:      `Beginning of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.dest-port-to",
				Short:      `End of the range of ports this rule applies to (inclusive). This value will be set to null if protocol is ICMP or ANY, or if it is equal to dest_port_from`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.position",
				Short:      `Position of this rule in the security group rules list. If several rules are passed with the same position, the resulting order is undefined`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.editable",
				Short:      `Indicates if this rule is editable. Rules with the value false will be ignored`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.zone",
				Short:      `Zone of the rule. This field is ignored`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.SetSecurityGroupRulesRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.SetSecurityGroupRules(request)

		},
	}
}

func instanceSecurityGroupDeleteRule() *core.Command {
	return &core.Command{
		Short:     `Delete rule`,
		Long:      `Delete a security group rule with the specified ID.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "delete-rule",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeleteSecurityGroupRuleRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "security-group-rule-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeleteSecurityGroupRuleRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeleteSecurityGroupRule(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "security-group",
				Verb:     "delete-rule",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete a security group rule with the specified ID",
				ArgsJSON: `{"security_group_id":"a01a36e5-5c0c-42c1-ae06-167e587b7ac4","security_group_rule_id":"b8c773ef-a6ea-4b50-a7c1-737864290a3f"}`,
			},
		},
	}
}

func instanceSecurityGroupGetRule() *core.Command {
	return &core.Command{
		Short:     `Get rule`,
		Long:      `Get details of a security group rule with the specified ID.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "get-rule",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetSecurityGroupRuleRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "security-group-rule-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetSecurityGroupRuleRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetSecurityGroupRule(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get details of a security group rule with the specified ID",
				ArgsJSON: `{"security_group_id":"d900fa38-2f0d-4b09-b6d7-f3e46a13f34c","security_group_rule_id":"1f9a16a5-7229-4c03-9327-253e257cf38a"}`,
			},
		},
	}
}

func instanceSecurityGroupUpdateRule() *core.Command {
	return &core.Command{
		Short:     `Update security group rule`,
		Long:      `Update the properties of a rule from a specified security group.`,
		Namespace: "instance",
		Resource:  "security-group",
		Verb:      "update-rule",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdateSecurityGroupRuleRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "security-group-id",
				Short:      `UUID of the security group`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "security-group-rule-id",
				Short:      `UUID of the rule`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "protocol",
				Short:      `Protocol family this rule applies to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_protocol", "TCP", "UDP", "ICMP", "ANY"},
			},
			{
				Name:       "direction",
				Short:      `Direction the rule applies to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_direction", "inbound", "outbound"},
			},
			{
				Name:       "action",
				Short:      `Action to apply when the rule matches a packet`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_action", "accept", "drop"},
			},
			{
				Name:       "ip-range",
				Short:      `Range of IP addresses these rules apply to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "dest-port-from",
				Short:      `Beginning of the range of ports this rule applies to (inclusive). If 0 is provided, unset the parameter.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "dest-port-to",
				Short:      `End of the range of ports this rule applies to (inclusive). If 0 is provided, unset the parameter.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "position",
				Short:      `Position of this rule in the security group rules list`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdateSecurityGroupRuleRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdateSecurityGroupRule(request)

		},
	}
}

func instancePlacementGroupList() *core.Command {
	return &core.Command{
		Short:     `List placement groups`,
		Long:      `List all placement groups in a specified Availability Zone.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListPlacementGroupsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "project",
				Short:      `List only placement groups of this Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags",
				Short:      `List placement groups with these exact tags (to filter with several tags, use commas to separate them)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Filter placement groups by name (for eg. "cluster1" will return "cluster100" and "cluster1" but not "foo")`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization",
				Short:      `List only placement groups of this Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListPlacementGroupsRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListPlacementGroups(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.PlacementGroups, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all placement groups in the default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "List placement groups that match a specified name ('cluster1' will return 'cluster100' and 'cluster1' but not 'foo')",
				ArgsJSON: `{"name":"cluster1"}`,
			},
		},
	}
}

func instancePlacementGroupCreate() *core.Command {
	return &core.Command{
		Short:     `Create a placement group`,
		Long:      `Create a new placement group in a specified Availability Zone.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.CreatePlacementGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("pg"),
			},
			core.ProjectArgSpec(),
			{
				Name:       "tags.{index}",
				Short:      `Tags of the placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "policy-mode",
				Short:      `Operating mode of the placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"optional", "enforced"},
			},
			{
				Name:       "policy-type",
				Short:      `Policy type of the placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"max_availability", "low_latency"},
			},
			core.OrganizationArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.CreatePlacementGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.CreatePlacementGroup(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Create a placement group with default name",
				ArgsJSON: `null`,
			},
			{
				Short:    "Create a placement group with the specified name",
				ArgsJSON: `{"name":"foobar"}`,
			},
			{
				Short:    "Create an enforced placement group",
				ArgsJSON: `{"policy_mode":"enforced"}`,
			},
			{
				Short:    "Create an optional placement group",
				ArgsJSON: `{"policy_mode":"optional"}`,
			},
			{
				Short:    "Create an optional low latency placement group",
				ArgsJSON: `{"policy_mode":"optional","policy_type":"low_latency"}`,
			},
			{
				Short:    "Create an enforced low latency placement group",
				ArgsJSON: `{"policy_mode":"enforced","policy_type":"low_latency"}`,
			},
		},
	}
}

func instancePlacementGroupGet() *core.Command {
	return &core.Command{
		Short:     `Get a placement group`,
		Long:      `Get the specified placement group.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetPlacementGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "placement-group-id",
				Short:      `UUID of the placement group you want to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetPlacementGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetPlacementGroup(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get a placement group with the specified ID",
				ArgsJSON: `{"placement_group_id":"6c15f411-3b6f-402d-8eba-ae24ef9254e9"}`,
			},
		},
	}
}

func instancePlacementGroupSet() *core.Command {
	return &core.Command{
		Short:     `Set placement group`,
		Long:      `Set all parameters of the specified placement group.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "set",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.SetPlacementGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "placement-group-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "policy-mode",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"optional", "enforced"},
			},
			{
				Name:       "policy-type",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"max_availability", "low_latency"},
			},
			core.ProjectArgSpec(),
			{
				Name:       "tags.{index}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.SetPlacementGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.SetPlacementGroup(request)

		},
	}
}

func instancePlacementGroupUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a placement group`,
		Long:      `Update one or more parameter of the specified placement group.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdatePlacementGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "placement-group-id",
				Short:      `UUID of the placement group`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of the placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "policy-mode",
				Short:      `Operating mode of the placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"optional", "enforced"},
			},
			{
				Name:       "policy-type",
				Short:      `Policy type of the placement group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"max_availability", "low_latency"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdatePlacementGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdatePlacementGroup(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Update the name of a placement group",
				ArgsJSON: `{"name":"foobar","placement_group_id":"95053f33-cd3c-4cdc-b2b0-57d2dda97b13"}`,
			},
			{
				Short:    "Update the policy mode of a placement group (all Instances in your placement group MUST be shut down)",
				ArgsJSON: `{"placement_group_id":"1f883434-8c2d-40f0-b686-d0754b3a7bc0","policy_mode":"enforced"}`,
			},
			{
				Short:    "Update the policy type of a placement group (all Instances in your placement group MUST be shutdown)",
				ArgsJSON: `{"placement_group_id":"0954ec26-9917-47b6-8c5c-7bc81d7bb9d2","policy_type":"low_latency"}`,
			},
		},
	}
}

func instancePlacementGroupDelete() *core.Command {
	return &core.Command{
		Short:     `Delete the specified placement group`,
		Long:      `Delete the specified placement group.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeletePlacementGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "placement-group-id",
				Short:      `UUID of the placement group you want to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeletePlacementGroupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeletePlacementGroup(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "placement-group",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete a placement group in the default zone with the specified ID",
				ArgsJSON: `{"placement_group_id":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Delete a placement group in fr-par-1 zone with the specified ID",
				ArgsJSON: `{"placement_group_id":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
		},
	}
}

func instancePlacementGroupGetServers() *core.Command {
	return &core.Command{
		Short:     `Get placement group servers`,
		Long:      `Get all Instances belonging to the specified placement group.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "get-servers",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetPlacementGroupServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "placement-group-id",
				Short:      `UUID of the placement group you want to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetPlacementGroupServersRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetPlacementGroupServers(request)

		},
	}
}

func instancePlacementGroupSetServers() *core.Command {
	return &core.Command{
		Short:     `Set placement group servers`,
		Long:      `Set all Instances belonging to the specified placement group.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "set-servers",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.SetPlacementGroupServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "placement-group-id",
				Short:      `UUID of the placement group you want to set`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "servers.{index}",
				Short:      `An array of the Instances' UUIDs you want to configure`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.SetPlacementGroupServersRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.SetPlacementGroupServers(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Update the complete set of Instances in a specified placement group (all Instances must be shut down)",
				ArgsJSON: `{"placement_group_id":"ced0fd4d-bcf0-4479-85b6-7027e54456e6","servers":["5a250608-24ec-4c31-9631-b3ded8c861cb","e54fd249-0787-4794-ab14-af6ee74df274"]}`,
			},
		},
	}
}

func instancePlacementGroupUpdateServers() *core.Command {
	return &core.Command{
		Short:     `Update placement group servers`,
		Long:      `Update all Instances belonging to the specified placement group.`,
		Namespace: "instance",
		Resource:  "placement-group",
		Verb:      "update-servers",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdatePlacementGroupServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "placement-group-id",
				Short:      `UUID of the placement group you want to update`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "servers.{index}",
				Short:      `An array of the Instances' UUIDs you want to configure`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdatePlacementGroupServersRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdatePlacementGroupServers(request)

		},
	}
}

func instanceIPList() *core.Command {
	return &core.Command{
		Short:     `List all flexible IPs`,
		Long:      `List all flexible IPs in a specified zone.`,
		Namespace: "instance",
		Resource:  "ip",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListIPsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "project",
				Short:      `Project ID in which the IPs are reserved`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags",
				Short:      `Filter IPs with these exact tags (to filter with several tags, use commas to separate them)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Filter on the IP address (Works as a LIKE operation on the IP address)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "type",
				Short:      `Filter on the IP Mobility IP type (whose value should be either 'routed_ipv4', 'routed_ipv6' or 'nat')`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization",
				Short:      `Organization ID in which the IPs are reserved`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListIPsRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
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
		Examples: []*core.Example{
			{
				Short:    "List all IPs in the default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "List all IPs in fr-par-1 zone",
				ArgsJSON: `{"zone":"fr-par-1"}`,
			},
		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Address",
			},
			{
				FieldName: "Prefix",
			},
			{
				FieldName: "Type",
			},
			{
				FieldName: "Reverse",
			},
			{
				FieldName: "Project",
			},
			{
				FieldName: "Server.ID",
			},
			{
				FieldName: "Server.Name",
			},
			{
				FieldName: "Tags",
			},
			{
				FieldName: "Zone",
			},
			{
				FieldName: "Organization",
			},
		}},
	}
}

func instanceIPCreate() *core.Command {
	return &core.Command{
		Short:     `Reserve a flexible IP`,
		Long:      `Reserve a flexible IP and attach it to the specified Instance.`,
		Namespace: "instance",
		Resource:  "ip",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.CreateIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectArgSpec(),
			{
				Name:       "tags.{index}",
				Short:      `Tags of the IP`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "server",
				Short:      `UUID of the Instance you want to attach the IP to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "type",
				Short:      `IP type to reserve (either 'routed_ipv4' or 'routed_ipv6', use of 'nat' is deprecated)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_iptype", "nat", "routed_ipv4", "routed_ipv6"},
			},
			core.OrganizationArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.CreateIPRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.CreateIP(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Create an IP in the default zone",
				ArgsJSON: `null`,
			},
			{
				Short:    "Create an IP in fr-par-1 zone",
				ArgsJSON: `{"zone":"fr-par-1"}`,
			},
			{
				Short:    "Create an IP and attach it to the specified Instance",
				ArgsJSON: `{"server":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func instanceIPGet() *core.Command {
	return &core.Command{
		Short:     `Get a flexible IP`,
		Long:      `Get details of an IP with the specified ID or address.`,
		Namespace: "instance",
		Resource:  "ip",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip",
				Short:      `IP ID or address to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetIPRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetIP(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Get an IP in the default zone with the specified ID",
				ArgsJSON: `{"ip":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Get an IP in fr-par-1 zone with the specified ID",
				ArgsJSON: `{"ip":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
			{
				Short:    "Get an IP, directly using the specified IP address",
				ArgsJSON: `{"ip":"51.15.253.183"}`,
			},
		},
	}
}

func instanceIPUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a flexible IP`,
		Long:      `Update a flexible IP in the specified zone with the specified ID.`,
		Namespace: "instance",
		Resource:  "ip",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdateIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip",
				Short:      `IP ID or IP address`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "reverse",
				Short:      `Reverse domain name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "type",
				Short:      `Convert a 'nat' IP to a 'routed_ipv4'`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_iptype", "nat", "routed_ipv4", "routed_ipv6"},
			},
			{
				Name:       "tags.{index}",
				Short:      `An array of keywords you want to tag this IP with`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdateIPRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdateIP(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Update an IP in the default zone with the specified ID",
				ArgsJSON: `{"ip":"11111111-1111-1111-1111-111111111111","reverse":"example.com"}`,
			},
			{
				Short:    "Update an IP in fr-par-1 zone with the specified ID",
				ArgsJSON: `{"ip":"11111111-1111-1111-1111-111111111111","reverse":"example.com","zone":"fr-par-1"}`,
			},
			{
				Short:    "Update an IP using directly the specified IP address",
				ArgsJSON: `{"ip":"51.15.253.183","reverse":"example.com"}`,
			},
		},
	}
}

func instanceIPDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a flexible IP`,
		Long:      `Delete the IP with the specified ID.`,
		Namespace: "instance",
		Resource:  "ip",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeleteIPRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ip",
				Short:      `ID or address of the IP to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeleteIPRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeleteIP(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "ip",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete an IP in the default zone with the specified ID",
				ArgsJSON: `{"ip":"11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Delete an IP in fr-par-1 zone with the specified ID",
				ArgsJSON: `{"ip":"11111111-1111-1111-1111-111111111111","zone":"fr-par-1"}`,
			},
			{
				Short:    "Delete an IP using directly the specified IP address",
				ArgsJSON: `{"ip":"51.15.253.183"}`,
			},
		},
	}
}

func instancePrivateNicList() *core.Command {
	return &core.Command{
		Short:     `List all private NICs`,
		Long:      `List all private NICs of a specified Instance.`,
		Namespace: "instance",
		Resource:  "private-nic",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ListPrivateNICsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `Instance to which the private NIC is attached`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags",
				Short:      `Private NIC tags`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ListPrivateNICsRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListPrivateNICs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.PrivateNics, nil

		},
		Examples: []*core.Example{
			{
				Short:    "List all private NICs on a specified server",
				ArgsJSON: `null`,
			},
			{
				Short:    "List private NICs of the Instance ID 'my_server_id'",
				ArgsJSON: `{"server_id":"my_server_id"}`,
			},
		},
	}
}

func instancePrivateNicCreate() *core.Command {
	return &core.Command{
		Short:     `Create a private NIC connecting an Instance to a Private Network`,
		Long:      `Create a private NIC connecting an Instance to a Private Network.`,
		Namespace: "instance",
		Resource:  "private-nic",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.CreatePrivateNICRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the Instance the private NIC will be attached to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `UUID of the private network where the private NIC will be attached`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Private NIC tags`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "ip-ids.{index}",
				Short:      `Ip_ids defined from IPAM`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "ipam-ip-ids.{index}",
				Short:      `UUID of IPAM ips, to be attached to the instance in the requested private network`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.CreatePrivateNICRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.CreatePrivateNIC(request)

		},
	}
}

func instancePrivateNicGet() *core.Command {
	return &core.Command{
		Short:     `Get a private NIC`,
		Long:      `Get private NIC properties.`,
		Namespace: "instance",
		Resource:  "private-nic",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.GetPrivateNICRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `Instance to which the private NIC is attached`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-nic-id",
				Short:      `Private NIC unique ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.GetPrivateNICRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.GetPrivateNIC(request)

		},
	}
}

func instancePrivateNicUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a private NIC`,
		Long:      `Update one or more parameter(s) of a specified private NIC.`,
		Namespace: "instance",
		Resource:  "private-nic",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.UpdatePrivateNICRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the Instance the private NIC will be attached to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-nic-id",
				Short:      `Private NIC unique ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags used to select private NIC/s`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.UpdatePrivateNICRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.UpdatePrivateNIC(request)

		},
		Examples: []*core.Example{
			{
				Short:    "Update tags of a private NIC",
				ArgsJSON: `{"private_nic_id":"11111111-1111-1111-1111-111111111111","server_id":"11111111-1111-1111-1111-111111111111","tags":["foo","bar"]}`,
			},
		},
	}
}

func instancePrivateNicDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a private NIC`,
		Long:      `Delete a private NIC.`,
		Namespace: "instance",
		Resource:  "private-nic",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.DeletePrivateNICRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `Instance to which the private NIC is attached`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-nic-id",
				Short:      `Private NIC unique ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.DeletePrivateNICRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.DeletePrivateNIC(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "private-nic",
				Verb:     "delete",
			}, nil
		},
	}
}

func instanceVolumePlanMigration() *core.Command {
	return &core.Command{
		Short:     `Get a volume or snapshot's migration plan`,
		Long:      `Given a volume or snapshot, returns the migration plan for a call to the "Apply a migration plan" endpoint. This plan will include zero or one volume, and zero or more snapshots, which will need to be migrated together. This endpoint does not perform the actual migration itself, the "Apply a migration plan" endpoint must be used. The validation_key value returned by this endpoint must be provided to the call to the "Apply a migration plan" endpoint to confirm that all resources listed in the plan should be migrated.`,
		Namespace: "instance",
		Resource:  "volume",
		Verb:      "plan-migration",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.PlanBlockMigrationRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-id",
				Short:      `The volume for which the migration plan will be generated.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "snapshot-id",
				Short:      `The snapshot for which the migration plan will be generated.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.PlanBlockMigrationRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.PlanBlockMigration(request)

		},
	}
}

func instanceVolumeApplyMigration() *core.Command {
	return &core.Command{
		Short:     `Migrate a volume and/or snapshots to SBS (Scaleway Block Storage)`,
		Long:      `To be used, the call to this endpoint must be preceded by a call to the "Plan a migration" endpoint. To migrate all resources mentioned in the migration plan, the validation_key returned in the plan must be provided.`,
		Namespace: "instance",
		Resource:  "volume",
		Verb:      "apply-migration",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(instance.ApplyBlockMigrationRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-id",
				Short:      `The volume to migrate, along with potentially other resources, according to the migration plan generated with a call to the "Plan a migration" endpoint.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "snapshot-id",
				Short:      `The snapshot to migrate, along with potentially other resources, according to the migration plan generated with a call to the "Plan a migration" endpoint.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "validation-key",
				Short:      `A value to be retrieved from a call to the "Plan a migration" endpoint, to confirm that the volume and/or snapshots specified in said plan should be migrated.`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*instance.ApplyBlockMigrationRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			e = api.ApplyBlockMigration(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "volume",
				Verb:     "apply-migration",
			}, nil
		},
	}
}

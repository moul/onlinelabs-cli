// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package mongodb

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/mongodb/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		mongodbRoot(),
		mongodbNodeType(),
		mongodbVersion(),
		mongodbInstance(),
		mongodbSnapshot(),
		mongodbUser(),
		mongodbEndpoint(),
		mongodbNodeTypeList(),
		mongodbVersionList(),
		mongodbInstanceList(),
		mongodbInstanceGet(),
		mongodbInstanceCreate(),
		mongodbInstanceUpdate(),
		mongodbInstanceDelete(),
		mongodbInstanceUpgrade(),
		mongodbInstanceGetCertificate(),
		mongodbSnapshotCreate(),
		mongodbSnapshotGet(),
		mongodbSnapshotUpdate(),
		mongodbSnapshotRestore(),
		mongodbSnapshotList(),
		mongodbSnapshotDelete(),
		mongodbUserList(),
		mongodbUserUpdate(),
	)
}
func mongodbRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Managed Databases for MongoDB®`,
		Long:      `This API allows you to manage your Managed Databases for MongoDB®.`,
		Namespace: "mongodb",
	}
}

func mongodbNodeType() *core.Command {
	return &core.Command{
		Short: `Node types management commands`,
		Long: `Two node type ranges are available:

* **Cost-Optimized:** a complete and highly reliable node range with shared resources that is made for scaling from development to production needs, at affordable prices.
* **Production-Optimized:** database nodes with dedicated vCPU for the most demanding workloads and mission-critical applications.`,
		Namespace: "mongodb",
		Resource:  "node-type",
	}
}

func mongodbVersion() *core.Command {
	return &core.Command{
		Short:     `MongoDB® version management commands`,
		Long:      `A database engine is the core software that handles the storage, retrieval, and management of data in your Database Instance.`,
		Namespace: "mongodb",
		Resource:  "version",
	}
}

func mongodbInstance() *core.Command {
	return &core.Command{
		Short:     `Instance management commands`,
		Long:      `A Managed Database for MongoDB® Database Instance is composed of one or multiple dedicated compute nodes running a single database engine.`,
		Namespace: "mongodb",
		Resource:  "instance",
	}
}

func mongodbSnapshot() *core.Command {
	return &core.Command{
		Short:     `Snapshot management commands`,
		Long:      `A snapshot is a consistent, instantaneous copy of the Block Storage volume of your Database Instance at a certain point in time.`,
		Namespace: "mongodb",
		Resource:  "snapshot",
	}
}

func mongodbUser() *core.Command {
	return &core.Command{
		Short:     `User management commands`,
		Long:      `Users are profiles to which you can attribute database-level permissions. They allow you to define permissions specific to each type of database usage.`,
		Namespace: "mongodb",
		Resource:  "user",
	}
}

func mongodbEndpoint() *core.Command {
	return &core.Command{
		Short:     `Endpoint management commands`,
		Long:      `Instance endpoints enable connection to your instance.`,
		Namespace: "mongodb",
		Resource:  "endpoint",
	}
}

func mongodbNodeTypeList() *core.Command {
	return &core.Command{
		Short:     `List available node types`,
		Long:      `List available node types.`,
		Namespace: "mongodb",
		Resource:  "node-type",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.ListNodeTypesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "include-disabled-types",
				Short:      `Defines whether or not to include disabled types`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.ListNodeTypesRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListNodeTypes(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.NodeTypes, nil

		},
	}
}

func mongodbVersionList() *core.Command {
	return &core.Command{
		Short:     `List available MongoDB® versions`,
		Long:      `List available MongoDB® versions.`,
		Namespace: "mongodb",
		Resource:  "version",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.ListVersionsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "version",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.ListVersionsRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListVersions(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Versions, nil

		},
	}
}

func mongodbInstanceList() *core.Command {
	return &core.Command{
		Short:     `List MongoDB® Database Instances`,
		Long:      `List all MongoDB® Database Instances in the specified region. By default, the MongoDB® Database Instances returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as ` + "`" + `tags` + "`" + ` and ` + "`" + `name` + "`" + `. For the ` + "`" + `name` + "`" + ` parameter, the value you include will be checked against the whole name string to see if it includes the string you put in the parameter.`,
		Namespace: "mongodb",
		Resource:  "instance",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.ListInstancesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "tags.{index}",
				Short:      `List Database Instances that have a given tag`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Lists Database Instances that match a name pattern`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering Database Instance listings`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc", "status_asc", "status_desc"},
			},
			{
				Name:       "project-id",
				Short:      `Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID of the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.ListInstancesRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListInstances(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Instances, nil

		},
	}
}

func mongodbInstanceGet() *core.Command {
	return &core.Command{
		Short:     `Get a MongoDB® Database Instance`,
		Long:      `Retrieve information about a given MongoDB® Database Instance, specified by the ` + "`" + `region` + "`" + ` and ` + "`" + `instance_id` + "`" + ` parameters. Its full details, including name, status, IP address and port, are returned in the response object.`,
		Namespace: "mongodb",
		Resource:  "instance",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.GetInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.GetInstanceRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.GetInstance(request)

		},
	}
}

func mongodbInstanceCreate() *core.Command {
	return &core.Command{
		Short:     `Create a MongoDB® Database Instance`,
		Long:      `Create a new MongoDB® Database Instance.`,
		Namespace: "mongodb",
		Resource:  "instance",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.CreateInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "name",
				Short:      `Name of the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("mgdb"),
			},
			{
				Name:       "version",
				Short:      `Version of the MongoDB® engine`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags to apply to the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "node-number",
				Short:      `Number of node to use for the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "node-type",
				Short:      `Type of node to use for the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-name",
				Short:      `Username created when the Database Instance is created`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "password",
				Short:      `Password of the initial user`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volume.volume-size",
				Short:      `Volume size`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volume.volume-type",
				Short:      `Type of volume where data is stored`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "sbs_5k", "sbs_15k"},
			},
			{
				Name:       "endpoints.{index}.private-network.private-network-id",
				Short:      `UUID of the Private Network`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.CreateInstanceRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.CreateInstance(request)

		},
	}
}

func mongodbInstanceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a MongoDB® Database Instance`,
		Long:      `Update the parameters of a MongoDB® Database Instance.`,
		Namespace: "mongodb",
		Resource:  "instance",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.UpdateInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of a Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.UpdateInstanceRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.UpdateInstance(request)

		},
	}
}

func mongodbInstanceDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a MongoDB® Database Instance`,
		Long:      `Delete a given MongoDB® Database Instance, specified by the ` + "`" + `region` + "`" + ` and ` + "`" + `instance_id` + "`" + ` parameters. Deleting a MongoDB® Database Instance is permanent, and cannot be undone. Note that upon deletion all your data will be lost.`,
		Namespace: "mongodb",
		Resource:  "instance",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.DeleteInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.DeleteInstanceRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.DeleteInstance(request)

		},
	}
}

func mongodbInstanceUpgrade() *core.Command {
	return &core.Command{
		Short:     `Upgrade a Database Instance`,
		Long:      `Upgrade your current Database Instance specifications like volume size.`,
		Namespace: "mongodb",
		Resource:  "instance",
		Verb:      "upgrade",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.UpgradeInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want to upgrade`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "volume-size",
				Short:      `Increase your Block Storage volume size`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.UpgradeInstanceRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.UpgradeInstance(request)

		},
	}
}

func mongodbInstanceGetCertificate() *core.Command {
	return &core.Command{
		Short:     `Get the certificate of a Database Instance`,
		Long:      `Retrieve the certificate of a given Database Instance, specified by the ` + "`" + `instance_id` + "`" + ` parameter.`,
		Namespace: "mongodb",
		Resource:  "instance",
		Verb:      "get-certificate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.GetInstanceCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.GetInstanceCertificateRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.GetInstanceCertificate(request)

		},
	}
}

func mongodbSnapshotCreate() *core.Command {
	return &core.Command{
		Short:     `Create a Database Instance snapshot`,
		Long:      `Create a new snapshot of a Database Instance. You must define the ` + "`" + `name` + "`" + ` and ` + "`" + `instance_id` + "`" + ` parameters in the request.`,
		Namespace: "mongodb",
		Resource:  "snapshot",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.CreateSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance to snapshot`,
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
				Name:       "expires-at",
				Short:      `Expiration date of the snapshot (must follow the ISO 8601 format)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.CreateSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.CreateSnapshot(request)

		},
	}
}

func mongodbSnapshotGet() *core.Command {
	return &core.Command{
		Short:     `Get a Database Instance snapshot`,
		Long:      `Retrieve information about a given snapshot of a Database Instance. You must specify, in the endpoint, the ` + "`" + `snapshot_id` + "`" + ` parameter of the snapshot you want to retrieve.`,
		Namespace: "mongodb",
		Resource:  "snapshot",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.GetSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.GetSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.GetSnapshot(request)

		},
	}
}

func mongodbSnapshotUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a Database Instance snapshot`,
		Long:      `Update the parameters of a snapshot of a Database Instance. You can update the ` + "`" + `name` + "`" + ` and ` + "`" + `expires_at` + "`" + ` parameters.`,
		Namespace: "mongodb",
		Resource:  "snapshot",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.UpdateSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the Snapshot`,
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
				Name:       "expires-at",
				Short:      `Expiration date of the snapshot (must follow the ISO 8601 format)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.UpdateSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.UpdateSnapshot(request)

		},
	}
}

func mongodbSnapshotRestore() *core.Command {
	return &core.Command{
		Short:     `Restore a Database Instance snapshot`,
		Long:      `Restore a given snapshot of a Database Instance. You must specify, in the endpoint, the ` + "`" + `snapshot_id` + "`" + ` parameter of the snapshot you want to restore, the ` + "`" + `instance_name` + "`" + ` of the new Database Instance, ` + "`" + `node_type` + "`" + ` of the new Database Instance and ` + "`" + `node_number` + "`" + ` of the new Database Instance.`,
		Namespace: "mongodb",
		Resource:  "snapshot",
		Verb:      "restore",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.RestoreSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "instance-name",
				Short:      `Name of the new Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "node-type",
				Short:      `Node type to use for the new Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "node-number",
				Short:      `Number of nodes to use for the new Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volume.volume-type",
				Short:      `Type of volume where data is stored`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "sbs_5k", "sbs_15k"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.RestoreSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.RestoreSnapshot(request)

		},
	}
}

func mongodbSnapshotList() *core.Command {
	return &core.Command{
		Short:     `List snapshots`,
		Long:      `List snapshots. You can include the ` + "`" + `instance_id` + "`" + ` or ` + "`" + `project_id` + "`" + ` in your query to get the list of snapshots for specific Database Instances and/or Projects. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "mongodb",
		Resource:  "snapshot",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.ListSnapshotsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `Instance ID the snapshots belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Lists database snapshots that match a name pattern`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering snapshot listings`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc", "expires_at_asc", "expires_at_desc"},
			},
			{
				Name:       "project-id",
				Short:      `Project ID to list the snapshots of`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID the snapshots belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.ListSnapshotsRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListSnapshots(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Snapshots, nil

		},
	}
}

func mongodbSnapshotDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Database Instance snapshot`,
		Long:      `Delete a given snapshot of a Database Instance. You must specify, in the endpoint, the ` + "`" + `snapshot_id` + "`" + ` parameter of the snapshot you want to delete.`,
		Namespace: "mongodb",
		Resource:  "snapshot",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.DeleteSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.DeleteSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.DeleteSnapshot(request)

		},
	}
}

func mongodbUserList() *core.Command {
	return &core.Command{
		Short:     `List users of a Database Instance`,
		Long:      `List all users of a given Database Instance.`,
		Namespace: "mongodb",
		Resource:  "user",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.ListUsersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the user`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when requesting user listing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc"},
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.ListUsersRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListUsers(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Users, nil

		},
	}
}

func mongodbUserUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a user on a Database Instance`,
		Long:      `Update the parameters of a user on a Database Instance. You can update the ` + "`" + `password` + "`" + ` parameter, but you cannot change the name of the user.`,
		Namespace: "mongodb",
		Resource:  "user",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mongodb.UpdateUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance the user belongs to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the database user`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "password",
				Short:      `Password of the database user`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mongodb.UpdateUserRequest)

			client := core.ExtractClient(ctx)
			api := mongodb.NewAPI(client)
			return api.UpdateUser(request)

		},
	}
}

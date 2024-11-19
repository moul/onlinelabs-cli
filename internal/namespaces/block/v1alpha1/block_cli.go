// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package block

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/block/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		blockRoot(),
		blockVolumeType(),
		blockVolume(),
		blockSnapshot(),
		blockVolumeTypeList(),
		blockVolumeList(),
		blockVolumeCreate(),
		blockVolumeGet(),
		blockVolumeDelete(),
		blockVolumeUpdate(),
		blockSnapshotList(),
		blockSnapshotGet(),
		blockSnapshotCreate(),
		blockSnapshotImportFromObjectStorage(),
		blockSnapshotExportToObjectStorage(),
		blockSnapshotDelete(),
		blockSnapshotUpdate(),
	)
}
func blockRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Block Storage volumes`,
		Long:      `This API allows you to manage your Block Storage volumes.`,
		Namespace: "block",
	}
}

func blockVolumeType() *core.Command {
	return &core.Command{
		Short:     `Block Storage volume types are determined by their storage class and their IOPS. There are two storage classes available: ` + "`" + `bssd` + "`" + ` and ` + "`" + `sbs` + "`" + `. The IOPS can be chosen for volumes of the ` + "`" + `sbs` + "`" + ` storage class`,
		Long:      `Block Storage volume types are determined by their storage class and their IOPS. There are two storage classes available: ` + "`" + `bssd` + "`" + ` and ` + "`" + `sbs` + "`" + `. The IOPS can be chosen for volumes of the ` + "`" + `sbs` + "`" + ` storage class.`,
		Namespace: "block",
		Resource:  "volume-type",
	}
}

func blockVolume() *core.Command {
	return &core.Command{
		Short:     `A Block Storage volume is a logical storage drive on a network-connected storage system. It is exposed to Instances as if it were a physical disk, and can be attached and detached like a hard drive. Several Block volumes can be attached to one Instance at a time`,
		Long:      `Block volumes can be snapshotted, mounted or unmounted.`,
		Namespace: "block",
		Resource:  "volume",
	}
}

func blockSnapshot() *core.Command {
	return &core.Command{
		Short:     `A Block Storage snapshot is a read-only picture of a Block volume, taken at a specific time`,
		Long:      `You can then revert your data to the previous snapshot. You can also create a new read/write Block volume from a previous snapshot.`,
		Namespace: "block",
		Resource:  "snapshot",
	}
}

func blockVolumeTypeList() *core.Command {
	return &core.Command{
		Short:     `List volume types`,
		Long:      `List all available volume types in a specified zone. The volume types listed are ordered by name in ascending order.`,
		Namespace: "block",
		Resource:  "volume-type",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.ListVolumeTypesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.ListVolumeTypesRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Zone == scw.Zone(core.AllLocalities) {
				opts = append(opts, scw.WithZones(api.Zones()...))
				request.Zone = ""
			}
			resp, err := api.ListVolumeTypes(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.VolumeTypes, nil

		},
	}
}

func blockVolumeList() *core.Command {
	return &core.Command{
		Short:     `List volumes`,
		Long:      `List all existing volumes in a specified zone. By default, the volumes listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "block",
		Resource:  "volume",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.ListVolumesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering the list`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "project-id",
				Short:      `Filter by Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Filter the return volumes by their names`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "product-resource-id",
				Short:      `Filter by a product resource ID linked to this volume (such as an Instance ID)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Filter by tags. Only volumes with one or more matching tags will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Filter by Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.ListVolumesRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
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
	}
}

func blockVolumeCreate() *core.Command {
	return &core.Command{
		Short: `Create a volume`,
		Long: `To create a new volume from scratch, you must specify ` + "`" + `from_empty` + "`" + ` and the ` + "`" + `size` + "`" + `.
To create a volume from an existing snapshot, specify ` + "`" + `from_snapshot` + "`" + ` and the ` + "`" + `snapshot_id` + "`" + ` in the request payload instead, size is optional and can be specified if you need to extend the original size. The volume will take on the same volume class and underlying IOPS limitations as the original snapshot.`,
		Namespace: "block",
		Resource:  "volume",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.CreateVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the volume`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("vol"),
			},
			{
				Name:       "perf-iops",
				Short:      `The maximum IO/s expected, according to the different options available in stock (` + "`" + `5000 | 15000` + "`" + `)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "from-empty.size",
				Short:      `Volume size in bytes, with a granularity of 1 GB (10^9 bytes)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "from-snapshot.size",
				Short:      `Volume size in bytes, with a granularity of 1 GB (10^9 bytes)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "from-snapshot.snapshot-id",
				Short:      `Source snapshot from which volume will be created`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `List of tags assigned to the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.CreateVolumeRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			return api.CreateVolume(request)

		},
	}
}

func blockVolumeGet() *core.Command {
	return &core.Command{
		Short:     `Get a volume`,
		Long:      `Retrieve technical information about a specific volume. Details such as size, type, and status are returned in the response.`,
		Namespace: "block",
		Resource:  "volume",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.GetVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-id",
				Short:      `UUID of the volume`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.GetVolumeRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			return api.GetVolume(request)

		},
	}
}

func blockVolumeDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a detached volume`,
		Long:      `You must specify the ` + "`" + `volume_id` + "`" + ` of the volume you want to delete. The volume must not be in the ` + "`" + `in_use` + "`" + ` status.`,
		Namespace: "block",
		Resource:  "volume",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.DeleteVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-id",
				Short:      `UUID of the volume`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.DeleteVolumeRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			e = api.DeleteVolume(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "volume",
				Verb:     "delete",
			}, nil
		},
	}
}

func blockVolumeUpdate() *core.Command {
	return &core.Command{
		Short: `Update a volume`,
		Long: `Update the technical details of a volume, such as its name, tags, or its new size and ` + "`" + `volume_type` + "`" + ` (within the same Block Storage class).
You can only resize a volume to a larger size. It is currently not possible to change your Block Storage Class.`,
		Namespace: "block",
		Resource:  "volume",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.UpdateVolumeRequest{}),
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
				Short:      `When defined, is the new name of the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "size",
				Short:      `Optional field for increasing the size of a volume (size must be equal or larger than the current one)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `List of tags assigned to the volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "perf-iops",
				Short:      `The maximum IO/s expected, according to the different options available in stock (` + "`" + `5000 | 15000` + "`" + `)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.UpdateVolumeRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			return api.UpdateVolume(request)

		},
	}
}

func blockSnapshotList() *core.Command {
	return &core.Command{
		Short:     `List all snapshots`,
		Long:      `List all available snapshots in a specified zone. By default, the snapshots listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "block",
		Resource:  "snapshot",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.ListSnapshotsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering the list`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "project-id",
				Short:      `Filter by Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volume-id",
				Short:      `Filter snapshots by the ID of the original volume`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Filter snapshots by their names`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Filter by Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.Zone(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.ListSnapshotsRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
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
	}
}

func blockSnapshotGet() *core.Command {
	return &core.Command{
		Short:     `Get a snapshot`,
		Long:      `Retrieve technical information about a specific snapshot. Details such as size, volume type, and status are returned in the response.`,
		Namespace: "block",
		Resource:  "snapshot",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.GetSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.GetSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			return api.GetSnapshot(request)

		},
	}
}

func blockSnapshotCreate() *core.Command {
	return &core.Command{
		Short: `Create a snapshot of a volume`,
		Long: `To create a snapshot, the volume must be in the ` + "`" + `in_use` + "`" + ` or the ` + "`" + `available` + "`" + ` status.
If your volume is in a transient state, you need to wait until the end of the current operation.`,
		Namespace: "block",
		Resource:  "snapshot",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.CreateSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "volume-id",
				Short:      `UUID of the volume to snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("snp"),
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "tags.{index}",
				Short:      `List of tags assigned to the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.CreateSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			return api.CreateSnapshot(request)

		},
	}
}

func blockSnapshotImportFromObjectStorage() *core.Command {
	return &core.Command{
		Short: `Import a snapshot from a Scaleway Object Storage bucket`,
		Long: `The bucket must contain a QCOW2 image.
The bucket can be imported into any Availability Zone as long as it is in the same region as the bucket.`,
		Namespace: "block",
		Resource:  "snapshot",
		Verb:      "import-from-object-storage",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.ImportSnapshotFromObjectStorageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "bucket",
				Short:      `Scaleway Object Storage bucket where the object is stored`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "key",
				Short:      `The object key inside the given bucket`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "tags.{index}",
				Short:      `List of tags assigned to the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "size",
				Short:      `Size of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.ImportSnapshotFromObjectStorageRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			return api.ImportSnapshotFromObjectStorage(request)

		},
	}
}

func blockSnapshotExportToObjectStorage() *core.Command {
	return &core.Command{
		Short: `Export a snapshot to a Scaleway Object Storage bucket`,
		Long: `The snapshot is exported in QCOW2 format.
The snapshot must not be in transient state.`,
		Namespace: "block",
		Resource:  "snapshot",
		Verb:      "export-to-object-storage",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.ExportSnapshotToObjectStorageRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "bucket",
				Short:      `Scaleway Object Storage bucket where the object is stored`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "key",
				Short:      `The object key inside the given bucket`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.ExportSnapshotToObjectStorageRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			return api.ExportSnapshotToObjectStorage(request)

		},
	}
}

func blockSnapshotDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a snapshot`,
		Long:      `You must specify the ` + "`" + `snapshot_id` + "`" + ` of the snapshot you want to delete. The snapshot must not be in use.`,
		Namespace: "block",
		Resource:  "snapshot",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.DeleteSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.DeleteSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			e = api.DeleteSnapshot(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "snapshot",
				Verb:     "delete",
			}, nil
		},
	}
}

func blockSnapshotUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a snapshot`,
		Long:      `Update the name or tags of the snapshot.`,
		Namespace: "block",
		Resource:  "snapshot",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(block.UpdateSnapshotRequest{}),
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
				Short:      `When defined, is the name of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `List of tags assigned to the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*block.UpdateSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := block.NewAPI(client)
			return api.UpdateSnapshot(request)

		},
	}
}

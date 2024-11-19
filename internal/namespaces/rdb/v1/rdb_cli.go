// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package rdb

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		rdbRoot(),
		rdbBackup(),
		rdbEngine(),
		rdbInstance(),
		rdbACL(),
		rdbPrivilege(),
		rdbUser(),
		rdbDatabase(),
		rdbNodeType(),
		rdbLog(),
		rdbSnapshot(),
		rdbReadReplica(),
		rdbSetting(),
		rdbEndpoint(),
		rdbEngineList(),
		rdbNodeTypeList(),
		rdbBackupList(),
		rdbBackupCreate(),
		rdbBackupGet(),
		rdbBackupUpdate(),
		rdbBackupDelete(),
		rdbBackupRestore(),
		rdbBackupExport(),
		rdbInstanceUpgrade(),
		rdbInstanceList(),
		rdbInstanceGet(),
		rdbInstanceCreate(),
		rdbInstanceUpdate(),
		rdbInstanceDelete(),
		rdbInstanceClone(),
		rdbInstanceRestart(),
		rdbInstanceGetCertificate(),
		rdbInstanceRenewCertificate(),
		rdbInstanceGetMetrics(),
		rdbReadReplicaCreate(),
		rdbReadReplicaGet(),
		rdbReadReplicaDelete(),
		rdbReadReplicaReset(),
		rdbReadReplicaCreateEndpoint(),
		rdbLogPrepare(),
		rdbLogList(),
		rdbLogGet(),
		rdbLogPurge(),
		rdbLogListDetails(),
		rdbSettingAdd(),
		rdbSettingDelete(),
		rdbSettingSet(),
		rdbACLList(),
		rdbACLAdd(),
		rdbACLSet(),
		rdbACLDelete(),
		rdbUserList(),
		rdbUserCreate(),
		rdbUserUpdate(),
		rdbUserDelete(),
		rdbDatabaseList(),
		rdbDatabaseCreate(),
		rdbDatabaseDelete(),
		rdbPrivilegeList(),
		rdbPrivilegeSet(),
		rdbSnapshotList(),
		rdbSnapshotGet(),
		rdbSnapshotCreate(),
		rdbSnapshotUpdate(),
		rdbSnapshotDelete(),
		rdbSnapshotRestore(),
		rdbEndpointCreate(),
		rdbEndpointDelete(),
		rdbEndpointGet(),
		rdbEndpointMigrate(),
	)
}
func rdbRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Managed Databases for PostgreSQL and MySQL`,
		Long:      `This API allows you to manage your Managed Databases for PostgreSQL and MySQL.`,
		Namespace: "rdb",
	}
}

func rdbBackup() *core.Command {
	return &core.Command{
		Short:     `Backup management commands`,
		Long:      `A database backup is a dated export of a Database Instance stored on an offsite backend located in a different region than your database, by default. Once a backup is created, it can be used to restore the database. Each logical database in a Database Instance is backed up and can be restored separately.`,
		Namespace: "rdb",
		Resource:  "backup",
	}
}

func rdbEngine() *core.Command {
	return &core.Command{
		Short:     `Database engines commands`,
		Long:      `A database engine is the software component that stores and retrieves your data from a database. Currently PostgreSQL 11, 12, 13 and 14 are available. MySQL is available in version 8.`,
		Namespace: "rdb",
		Resource:  "engine",
	}
}

func rdbInstance() *core.Command {
	return &core.Command{
		Short: `Instance management commands`,
		Long: `A Database Instance is made up of one or multiple dedicated compute nodes running a single database engine. Two node settings are available: **High-Availability (HA)**, with a main node and one replica, and **standalone** with a main node. The HA standby node is linked to the main node, using synchronous replication. Synchronous replication offers the ability to confirm that all changes intended by a transaction have been transferred and applied to the synchronous replica node, providing durability to the data.

**Note**: HA standby nodes are not accessible to users unless the main node becomes unavailable and the standby takes over. If you wish to run queries on a read-only node, you can use [Read Replicas](#path-read-replicas-create-a-read-replica)

Read Replicas can be used for certain read-only workflows such as Business Intelligence, or for a read-only scaling of your application. Read Replicas use asynchronous replication to replicate data from the main node.`,
		Namespace: "rdb",
		Resource:  "instance",
	}
}

func rdbACL() *core.Command {
	return &core.Command{
		Short:     `Access Control List (ACL) management commands`,
		Long:      `Network Access Control Lists allow you to control incoming network traffic by setting up ACL rules.`,
		Namespace: "rdb",
		Resource:  "acl",
	}
}

func rdbPrivilege() *core.Command {
	return &core.Command{
		Short: `User privileges management commands`,
		Long: `Privileges are permissions that can be granted to database users. You can manage user permissions either via the console, the Scaleway APIs or SQL. Managed Database for PostgreSQL and MySQL provides a simplified and unified permission model through the API and the console to make things easier to manage and understand.

Each user has associated permissions that give them access to zero or more logical databases. These include:

* **None:** No access to the database
* **Read:** Allow users to read tables and fields in a database
* **Write:** Allow users to write content in databases.
* **Admin:** Read and write access to the data, and extended privileges depending on the database engine.`,
		Namespace: "rdb",
		Resource:  "privilege",
	}
}

func rdbUser() *core.Command {
	return &core.Command{
		Short:     `User management commands`,
		Long:      `Users are profiles to which you can attribute database-level permissions. They allow you to define permissions specific to each type of database usage. For example, users with an ` + "`" + `admin` + "`" + ` role can create new databases and users.`,
		Namespace: "rdb",
		Resource:  "user",
	}
}

func rdbDatabase() *core.Command {
	return &core.Command{
		Short:     `Database management commands`,
		Long:      `Databases can be used to store and manage sets of structured information, or data. The interaction between the user and a database is done using a Database Engine, which provides a structured query language to add, modify or delete information from the database.`,
		Namespace: "rdb",
		Resource:  "database",
	}
}

func rdbNodeType() *core.Command {
	return &core.Command{
		Short: `Node types management commands`,
		Long: `Two node type ranges are available:

* **General Purpose:** production-grade nodes designed for scalable database infrastructures.
* **Development:** sandbox environments and reliable performance for development and testing purposes.`,
		Namespace: "rdb",
		Resource:  "node-type",
	}
}

func rdbLog() *core.Command {
	return &core.Command{
		Short:     `Instance logs management commands`,
		Long:      `Instance logs management commands.`,
		Namespace: "rdb",
		Resource:  "log",
	}
}

func rdbSnapshot() *core.Command {
	return &core.Command{
		Short:     `Block snapshot management`,
		Long:      `A snapshot is a consistent, instantaneous copy of the Block Storage volume of your Database Instance at a certain point in time. They are designed to recover your data in case of failure or accidental alterations of the data by a user. They allow you to quickly create a new Instance from a previous state of your database, regardless of the size of the volume. Their limitation is that, unlike backups, snapshots can only be stored in the same location as the original data.`,
		Namespace: "rdb",
		Resource:  "snapshot",
	}
}

func rdbReadReplica() *core.Command {
	return &core.Command{
		Short: `Read replica management`,
		Long: `A Read Replica is a live copy of a Database Instance that behaves like an Instance, but that only allows read-only connections.
The replica mirrors the data of the primary Database node and any changes made are replicated to the replica asynchronously. Read Replicas allow you to scale your Database Instance for read-heavy database workloads. They can also be used for business intelligence workloads.

A Read Replica can have at most one direct access and one Private Network endpoint. ` + "`" + `Loadbalancer` + "`" + ` endpoints are not available on Read Replicas even if this resource is displayed in the Read Replica response example.

If you want to remove a Read Replica endpoint, you can use [delete a Database Instance endpoint](#path-endpoints-delete-a-database-instance-endpoint) API call.

Instance Access Control Lists (ACL) also apply to Read Replica direct access endpoints.

**Limitations:**
There might be replication lags between the primary node and its Read Replica nodes. You can try to reduce this lag with some good practices:
* All your tables should have a primary key
* Don't run large transactions that modify, delete or insert lots of rows. Try to split it into several small transactions.`,
		Namespace: "rdb",
		Resource:  "read-replica",
	}
}

func rdbSetting() *core.Command {
	return &core.Command{
		Short: `Setting management`,
		Long: `Advanced Database Instance settings allow you to tune the behavior of your database engines to better fit your needs.

Available settings depend on the database engine and its version. Note that some settings can only be defined upon database engine initialization. These are called init settings. You can find a full list of the settings available in the response body of the [list available database engines](#path-databases-list-databases-in-a-database-instance) endpoint.

Each advanced setting entry has a default value that users can override. The deletion of a setting entry will restore the setting to default value. Some of the defaults values can be different from the engine's defaults, as we optimize them to the Scaleway platform.`,
		Namespace: "rdb",
		Resource:  "setting",
	}
}

func rdbEndpoint() *core.Command {
	return &core.Command{
		Short: `Endpoint management`,
		Long: `A point of connection to a Database Instance. The endpoint is associated with an IPv4 address and a port. It contains the information about whether the endpoint is read-write or not. The endpoints always point to the main node of a Database Instance.

All endpoints have TLS enabled. You can use TLS to make your data and your passwords unreadable in transit to anyone but you.

For added security, you can set up ACL rules to restrict access to your endpoint to a set of trusted hosts or networks of your choice.

Load Balancers are used to forward traffic to the right node based on the node state (active/hot standby). The Load Balancers' configuration is set to cut off inactive connections if no TCP traffic is sent within a 6-hour timeframe. We recommend using connection pooling on the application side to renew database connections regularly.`,
		Namespace: "rdb",
		Resource:  "endpoint",
	}
}

func rdbEngineList() *core.Command {
	return &core.Command{
		Short:     `List available database engines`,
		Long:      `List the PostgreSQL and MySQL database engines available at Scaleway.`,
		Namespace: "rdb",
		Resource:  "engine",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListDatabaseEnginesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the database engine`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "version",
				Short:      `Version of the database engine`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListDatabaseEnginesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListDatabaseEngines(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Engines, nil

		},
	}
}

func rdbNodeTypeList() *core.Command {
	return &core.Command{
		Short:     `List available node types`,
		Long:      `List all available node types. By default, the node types returned in the list are ordered by creation date in ascending order, though this can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "rdb",
		Resource:  "node-type",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListNodeTypesRequest{}),
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
			request := args.(*rdb.ListNodeTypesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
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

func rdbBackupList() *core.Command {
	return &core.Command{
		Short:     `List database backups`,
		Long:      `List all backups in a specified region, for a given Scaleway Organization or Scaleway Project. By default, the backups listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListDatabaseBackupsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the database backups`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering database backups listing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc", "status_asc", "status_desc"},
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Project ID of the Project the database backups belong to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID of the Organization the database backups belong to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListDatabaseBackupsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListDatabaseBackups(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.DatabaseBackups, nil

		},
	}
}

func rdbBackupCreate() *core.Command {
	return &core.Command{
		Short:     `Create a database backup`,
		Long:      `Create a new backup. You must set the ` + "`" + `instance_id` + "`" + `, ` + "`" + `database_name` + "`" + `, ` + "`" + `name` + "`" + ` and ` + "`" + `expires_at` + "`" + ` parameters.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "database-name",
				Short:      `Name of the database you want to back up`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the backup`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("bkp"),
			},
			{
				Name:       "expires-at",
				Short:      `Expiration date (must follow the ISO 8601 format)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateDatabaseBackup(request)

		},
	}
}

func rdbBackupGet() *core.Command {
	return &core.Command{
		Short:     `Get a database backup`,
		Long:      `Retrieve information about a given backup, specified by its database backup ID and region. Full details about the backup, like size, URL and expiration date, are returned in the response.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.GetDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-backup-id",
				Short:      `UUID of the database backup`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetDatabaseBackup(request)

		},
	}
}

func rdbBackupUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a database backup`,
		Long:      `Update the parameters of a backup, including name and expiration date.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.UpdateDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-backup-id",
				Short:      `UUID of the database backup to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the Database Backup`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "expires-at",
				Short:      `Expiration date (must follow the ISO 8601 format)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpdateDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpdateDatabaseBackup(request)

		},
	}
}

func rdbBackupDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a database backup`,
		Long:      `Delete a backup, specified by its database backup ID and region. Deleting a backup is permanent, and cannot be undone.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-backup-id",
				Short:      `UUID of the database backup to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteDatabaseBackup(request)

		},
	}
}

func rdbBackupRestore() *core.Command {
	return &core.Command{
		Short:     `Restore a database backup`,
		Long:      `Launch the process of restoring database backup. You must specify the ` + "`" + `instance_id` + "`" + ` of the Database Instance of destination, where the backup will be restored. Note that large database backups can take up to several hours to restore.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "restore",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.RestoreDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-name",
				Short:      `Defines the destination database to restore into a specified database (the default destination is set to the origin database of the backup)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "database-backup-id",
				Short:      `Backup of a logical database`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "instance-id",
				Short:      `Defines the Database Instance where the backup has to be restored`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.RestoreDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.RestoreDatabaseBackup(request)

		},
	}
}

func rdbBackupExport() *core.Command {
	return &core.Command{
		Short:     `Export a database backup`,
		Long:      `Export a backup, specified by the ` + "`" + `database_backup_id` + "`" + ` and the ` + "`" + `region` + "`" + ` parameters. The download URL is returned in the response.`,
		Namespace: "rdb",
		Resource:  "backup",
		Verb:      "export",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ExportDatabaseBackupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "database-backup-id",
				Short:      `UUID of the database backup you want to export`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ExportDatabaseBackupRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.ExportDatabaseBackup(request)

		},
	}
}

func rdbInstanceUpgrade() *core.Command {
	return &core.Command{
		Short:     `Upgrade a Database Instance`,
		Long:      `Upgrade your current Database Instance specifications like node type, high availability, volume, or the database engine version. Note that upon upgrade the ` + "`" + `enable_ha` + "`" + ` parameter can only be set to ` + "`" + `true` + "`" + `.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "upgrade",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.UpgradeInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want to upgrade`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "node-type",
				Short:      `Node type of the Database Instance you want to upgrade to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "enable-ha",
				Short:      `Defines whether or not high availability should be enabled on the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volume-size",
				Short:      `Increase your block storage volume size`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volume-type",
				Short:      `Change your Database Instance storage type`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"lssd", "bssd", "sbs_5k", "sbs_15k"},
			},
			{
				Name:       "upgradable-version-id",
				Short:      `Update your database engine to a newer version`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "major-upgrade-workflow.upgradable-version-id",
				Short:      `Update your database engine to a newer version`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "major-upgrade-workflow.with-endpoints",
				Short:      `Include endpoint during the migration`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "enable-encryption",
				Short:      `Defines whether or not encryption should be enabled on the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpgradeInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpgradeInstance(request)

		},
	}
}

func rdbInstanceList() *core.Command {
	return &core.Command{
		Short:     `List Database Instances`,
		Long:      `List all Database Instances in the specified region, for a given Scaleway Organization or Scaleway Project. By default, the Database Instances returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as ` + "`" + `tags` + "`" + ` and ` + "`" + `name` + "`" + `. For the ` + "`" + `name` + "`" + ` parameter, the value you include will be checked against the whole name string to see if it includes the string you put in the parameter.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListInstancesRequest{}),
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
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc", "region", "status_asc", "status_desc"},
			},
			{
				Name:       "project-id",
				Short:      `Project ID to list the Database Instance of`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Please use project_id instead`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListInstancesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
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
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "NodeType",
			},
			{
				FieldName: "Status",
			},
			{
				FieldName: "Engine",
			},
			{
				FieldName: "Region",
			},
			{
				FieldName: "Tags",
			},
			{
				FieldName: "IsHaCluster",
			},
			{
				FieldName: "BackupSchedule",
			},
			{
				FieldName: "CreatedAt",
			},
		}},
	}
}

func rdbInstanceGet() *core.Command {
	return &core.Command{
		Short:     `Get a Database Instance`,
		Long:      `Retrieve information about a given Database Instance, specified by the ` + "`" + `region` + "`" + ` and ` + "`" + `instance_id` + "`" + ` parameters. Its full details, including name, status, IP address and port, are returned in the response object.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.GetInstanceRequest{}),
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
			request := args.(*rdb.GetInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetInstance(request)

		},
	}
}

func rdbInstanceCreate() *core.Command {
	return &core.Command{
		Short:     `Create a Database Instance`,
		Long:      `Create a new Database Instance. You must set the ` + "`" + `engine` + "`" + `, ` + "`" + `user_name` + "`" + `, ` + "`" + `password` + "`" + ` and ` + "`" + `node_type` + "`" + ` parameters. Optionally, you can specify the volume type and size.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "name",
				Short:      `Name of the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("ins"),
			},
			{
				Name:       "engine",
				Short:      `Database engine of the Database Instance (PostgreSQL, MySQL, ...)`,
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
				Short:      `Password of the user`,
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
				Name:       "is-ha-cluster",
				Short:      `Defines whether or not High-Availability is enabled`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "disable-backup",
				Short:      `Defines whether or not backups are disabled`,
				Required:   false,
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
				Name:       "init-settings.{index}.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "init-settings.{index}.value",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "volume-type",
				Short:      `Type of volume where data is stored (lssd, bssd, ...)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"lssd", "bssd", "sbs_5k", "sbs_15k"},
			},
			{
				Name:       "volume-size",
				Short:      `Volume size when volume_type is not lssd`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "init-endpoints.{index}.private-network.private-network-id",
				Short:      `UUID of the Private Network to be connected to the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "init-endpoints.{index}.private-network.service-ip",
				Short:      `Endpoint IPv4 address with a CIDR notation. Refer to the official Scaleway documentation to learn more about IP and subnet limitations.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "backup-same-region",
				Short:      `Defines whether to or not to store logical backups in the same region as the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "encryption.enabled",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateInstance(request)

		},
	}
}

func rdbInstanceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a Database Instance`,
		Long:      `Update the parameters of a Database Instance, including name, tags and backup schedule details.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.UpdateInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "backup-schedule-frequency",
				Short:      `In hours`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "backup-schedule-retention",
				Short:      `In days`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "is-backup-schedule-disabled",
				Short:      `Defines whether or not the backup schedule is disabled`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags of a Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "logs-policy.max-age-retention",
				Short:      `Max age (in days) of remote logs to keep on the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "logs-policy.total-disk-retention",
				Short:      `Max disk size of remote logs to keep on the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "backup-same-region",
				Short:      `Store logical backups in the same region as the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "backup-schedule-start-hour",
				Short:      `Defines the start time of the autobackup`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpdateInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpdateInstance(request)

		},
	}
}

func rdbInstanceDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Database Instance`,
		Long:      `Delete a given Database Instance, specified by the ` + "`" + `region` + "`" + ` and ` + "`" + `instance_id` + "`" + ` parameters. Deleting a Database Instance is permanent, and cannot be undone. Note that upon deletion all your data will be lost.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteInstanceRequest{}),
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
			request := args.(*rdb.DeleteInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteInstance(request)

		},
	}
}

func rdbInstanceClone() *core.Command {
	return &core.Command{
		Short:     `Clone a Database Instance`,
		Long:      `Clone a given Database Instance, specified by the ` + "`" + `region` + "`" + ` and ` + "`" + `instance_id` + "`" + ` parameters. The clone feature allows you to create a new Database Instance from an existing one. The clone includes all existing databases, users and permissions. You can create a clone on a Database Instance bigger than your current one.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "clone",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CloneInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want to clone`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the Database Instance clone`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "node-type",
				Short:      `Node type of the clone`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CloneInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CloneInstance(request)

		},
	}
}

func rdbInstanceRestart() *core.Command {
	return &core.Command{
		Short:     `Restart Database Instance`,
		Long:      `Restart a given Database Instance, specified by the ` + "`" + `region` + "`" + ` and ` + "`" + `instance_id` + "`" + ` parameters. The status of the Database Instance returned in the response.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "restart",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.RestartInstanceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want to restart`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.RestartInstanceRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.RestartInstance(request)

		},
	}
}

func rdbInstanceGetCertificate() *core.Command {
	return &core.Command{
		Short:     `Get the TLS certificate of a Database Instance`,
		Long:      `Retrieve information about the TLS certificate of a given Database Instance. Details like name and content are returned in the response.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "get-certificate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.GetInstanceCertificateRequest{}),
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
			request := args.(*rdb.GetInstanceCertificateRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetInstanceCertificate(request)

		},
	}
}

func rdbInstanceRenewCertificate() *core.Command {
	return &core.Command{
		Short:     `Renew the TLS certificate of a Database Instance`,
		Long:      `Renew a TLS for a Database Instance. Renewing a certificate means that you will not be able to connect to your Database Instance using the previous certificate. You will also need to download and update the new certificate for all database clients.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "renew-certificate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.RenewInstanceCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want logs of`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.RenewInstanceCertificateRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			e = api.RenewInstanceCertificate(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "instance",
				Verb:     "renew-certificate",
			}, nil
		},
	}
}

func rdbInstanceGetMetrics() *core.Command {
	return &core.Command{
		Short:     `Get Database Instance metrics`,
		Long:      `Retrieve the time series metrics of a given Database Instance. You can define the period from which to retrieve metrics by specifying the ` + "`" + `start_date` + "`" + ` and ` + "`" + `end_date` + "`" + `.`,
		Namespace: "rdb",
		Resource:  "instance",
		Verb:      "get-metrics",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.GetInstanceMetricsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "start-date",
				Short:      `Start date to gather metrics from`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "end-date",
				Short:      `End date to gather metrics from`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "metric-name",
				Short:      `Name of the metric to gather`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetInstanceMetricsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetInstanceMetrics(request)

		},
	}
}

func rdbReadReplicaCreate() *core.Command {
	return &core.Command{
		Short:     `Create a Read Replica`,
		Long:      `Create a new Read Replica of a Database Instance. You must specify the ` + "`" + `region` + "`" + ` and the ` + "`" + `instance_id` + "`" + `. You can only create a maximum of 3 Read Replicas per Database Instance.`,
		Namespace: "rdb",
		Resource:  "read-replica",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateReadReplicaRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want to create a Read Replica from`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "endpoint-spec.{index}.private-network.private-network-id",
				Short:      `UUID of the Private Network to be connected to the Read Replica`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "endpoint-spec.{index}.private-network.service-ip",
				Short:      `Endpoint IPv4 address with a CIDR notation. Refer to the official Scaleway documentation to learn more about IP and subnet limitations.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "same-zone",
				Short:      `Defines whether to create the replica in the same availability zone as the main instance nodes or not.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateReadReplicaRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateReadReplica(request)

		},
	}
}

func rdbReadReplicaGet() *core.Command {
	return &core.Command{
		Short:     `Get a Read Replica`,
		Long:      `Retrieve information about a Database Instance Read Replica. Full details about the Read Replica, like ` + "`" + `endpoints` + "`" + `, ` + "`" + `status` + "`" + `  and ` + "`" + `region` + "`" + ` are returned in the response.`,
		Namespace: "rdb",
		Resource:  "read-replica",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.GetReadReplicaRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "read-replica-id",
				Short:      `UUID of the Read Replica`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetReadReplicaRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetReadReplica(request)

		},
	}
}

func rdbReadReplicaDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Read Replica`,
		Long:      `Delete a Read Replica of a Database Instance. You must specify the ` + "`" + `region` + "`" + ` and ` + "`" + `read_replica_id` + "`" + ` parameters of the Read Replica you want to delete.`,
		Namespace: "rdb",
		Resource:  "read-replica",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteReadReplicaRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "read-replica-id",
				Short:      `UUID of the Read Replica`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteReadReplicaRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteReadReplica(request)

		},
	}
}

func rdbReadReplicaReset() *core.Command {
	return &core.Command{
		Short: `Resync a Read Replica`,
		Long: `When you resync a Read Replica, first it is reset, then its data is resynchronized from the primary node. Your Read Replica remains unavailable during the resync process. The duration of this process is proportional to the size of your Database Instance.
The configured endpoints do not change.`,
		Namespace: "rdb",
		Resource:  "read-replica",
		Verb:      "reset",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ResetReadReplicaRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "read-replica-id",
				Short:      `UUID of the Read Replica`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ResetReadReplicaRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.ResetReadReplica(request)

		},
	}
}

func rdbReadReplicaCreateEndpoint() *core.Command {
	return &core.Command{
		Short:     `Create an endpoint for a Read Replica`,
		Long:      `Create a new endpoint for a Read Replica. Read Replicas can have at most one direct access and one Private Network endpoint.`,
		Namespace: "rdb",
		Resource:  "read-replica",
		Verb:      "create-endpoint",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateReadReplicaEndpointRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "read-replica-id",
				Short:      `UUID of the Read Replica`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "endpoint-spec.{index}.private-network.private-network-id",
				Short:      `UUID of the Private Network to be connected to the Read Replica`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "endpoint-spec.{index}.private-network.service-ip",
				Short:      `Endpoint IPv4 address with a CIDR notation. Refer to the official Scaleway documentation to learn more about IP and subnet limitations.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateReadReplicaEndpointRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateReadReplicaEndpoint(request)

		},
	}
}

func rdbLogPrepare() *core.Command {
	return &core.Command{
		Short:     `Prepare logs of a Database Instance`,
		Long:      `Prepare your Database Instance logs. You can define the ` + "`" + `start_date` + "`" + ` and ` + "`" + `end_date` + "`" + ` parameters for your query. The download URL is returned in the response. Logs are recorded from 00h00 to 23h59 and then aggregated in a ` + "`" + `.log` + "`" + ` file once a day. Therefore, even if you specify a timeframe from which you want to get the logs, you will receive logs from the full 24 hours.`,
		Namespace: "rdb",
		Resource:  "log",
		Verb:      "prepare",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.PrepareInstanceLogsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want logs of`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "start-date",
				Short:      `Start datetime of your log. (RFC 3339 format)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "end-date",
				Short:      `End datetime of your log. (RFC 3339 format)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.PrepareInstanceLogsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.PrepareInstanceLogs(request)

		},
	}
}

func rdbLogList() *core.Command {
	return &core.Command{
		Short:     `List available logs of a Database Instance`,
		Long:      `List the available logs of a Database Instance. By default, the logs returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.`,
		Namespace: "rdb",
		Resource:  "log",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListInstanceLogsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want logs of`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering Database Instance logs listing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListInstanceLogsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.ListInstanceLogs(request)

		},
	}
}

func rdbLogGet() *core.Command {
	return &core.Command{
		Short:     `Get given logs of a Database Instance`,
		Long:      `Retrieve information about the logs of a Database Instance. Specify the ` + "`" + `instance_log_id` + "`" + ` and ` + "`" + `region` + "`" + ` in your request to get information such as ` + "`" + `download_url` + "`" + `, ` + "`" + `status` + "`" + `, ` + "`" + `expires_at` + "`" + ` and ` + "`" + `created_at` + "`" + ` about your logs in the response.`,
		Namespace: "rdb",
		Resource:  "log",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.GetInstanceLogRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-log-id",
				Short:      `UUID of the instance_log you want`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetInstanceLogRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetInstanceLog(request)

		},
	}
}

func rdbLogPurge() *core.Command {
	return &core.Command{
		Short:     `Purge remote Database Instance logs`,
		Long:      `Purge a given remote log from a Database Instance. You can specify the ` + "`" + `log_name` + "`" + ` of the log you wish to clean from your Database Instance.`,
		Namespace: "rdb",
		Resource:  "log",
		Verb:      "purge",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.PurgeInstanceLogsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want logs of`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "log-name",
				Short:      `Given log name to purge`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.PurgeInstanceLogsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			e = api.PurgeInstanceLogs(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "log",
				Verb:     "purge",
			}, nil
		},
	}
}

func rdbLogListDetails() *core.Command {
	return &core.Command{
		Short:     `List remote Database Instance logs details`,
		Long:      `List remote log details. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.`,
		Namespace: "rdb",
		Resource:  "log",
		Verb:      "list-details",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListInstanceLogsDetailsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want logs of`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListInstanceLogsDetailsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.ListInstanceLogsDetails(request)

		},
	}
}

func rdbSettingAdd() *core.Command {
	return &core.Command{
		Short:     `Add Database Instance advanced settings`,
		Long:      `Add an advanced setting to a Database Instance. You must set the ` + "`" + `name` + "`" + ` and the ` + "`" + `value` + "`" + ` of each setting.`,
		Namespace: "rdb",
		Resource:  "setting",
		Verb:      "add",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.AddInstanceSettingsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want to add settings to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "settings.{index}.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "settings.{index}.value",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.AddInstanceSettingsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.AddInstanceSettings(request)

		},
	}
}

func rdbSettingDelete() *core.Command {
	return &core.Command{
		Short:     `Delete Database Instance advanced settings`,
		Long:      `Delete an advanced setting in a Database Instance. You must specify the names of the settings you want to delete in the request.`,
		Namespace: "rdb",
		Resource:  "setting",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteInstanceSettingsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance to delete settings from`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "setting-names.{index}",
				Short:      `Settings names to delete`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteInstanceSettingsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteInstanceSettings(request)

		},
	}
}

func rdbSettingSet() *core.Command {
	return &core.Command{
		Short:     `Set Database Instance advanced settings`,
		Long:      `Update an advanced setting for a Database Instance. Settings added upon database engine initalization can only be defined once, and cannot, therefore, be updated.`,
		Namespace: "rdb",
		Resource:  "setting",
		Verb:      "set",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.SetInstanceSettingsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance where the settings must be set`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "settings.{index}.name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "settings.{index}.value",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.SetInstanceSettingsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.SetInstanceSettings(request)

		},
	}
}

func rdbACLList() *core.Command {
	return &core.Command{
		Short:     `List ACL rules of a Database Instance`,
		Long:      `List the ACL rules for a given Database Instance. The response is an array of ACL objects, each one representing an ACL that denies, allows or redirects traffic based on certain conditions.`,
		Namespace: "rdb",
		Resource:  "acl",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListInstanceACLRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
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
			request := args.(*rdb.ListInstanceACLRulesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListInstanceACLRules(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Rules, nil

		},
	}
}

func rdbACLAdd() *core.Command {
	return &core.Command{
		Short:     `Add an ACL rule to a Database Instance`,
		Long:      `Add an additional ACL rule to a Database Instance.`,
		Namespace: "rdb",
		Resource:  "acl",
		Verb:      "add",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.AddInstanceACLRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want to add ACL rules to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.ip",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.description",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.AddInstanceACLRulesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.AddInstanceACLRules(request)

		},
	}
}

func rdbACLSet() *core.Command {
	return &core.Command{
		Short:     `Set ACL rules for a Database Instance`,
		Long:      `Replace all the ACL rules of a Database Instance.`,
		Namespace: "rdb",
		Resource:  "acl",
		Verb:      "set",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.SetInstanceACLRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance where the ACL rules must be set`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.ip",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.description",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.SetInstanceACLRulesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.SetInstanceACLRules(request)

		},
	}
}

func rdbACLDelete() *core.Command {
	return &core.Command{
		Short:     `Delete ACL rules of a Database Instance`,
		Long:      `Delete one or more ACL rules of a Database Instance.`,
		Namespace: "rdb",
		Resource:  "acl",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteInstanceACLRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you want to delete an ACL rule from`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "acl-rule-ips.{index}",
				Short:      `IP addresses defined in the ACL rules of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteInstanceACLRulesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteInstanceACLRules(request)

		},
	}
}

func rdbUserList() *core.Command {
	return &core.Command{
		Short:     `List users of a Database Instance`,
		Long:      `List all users of a given Database Instance. By default, the users returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.`,
		Namespace: "rdb",
		Resource:  "user",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListUsersRequest{}),
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
				EnumValues: []string{"name_asc", "name_desc", "is_admin_asc", "is_admin_desc"},
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
			request := args.(*rdb.ListUsersRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
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

func rdbUserCreate() *core.Command {
	return &core.Command{
		Short:     `Create a user for a Database Instance`,
		Long:      `Create a new user for a Database Instance. You must define the ` + "`" + `name` + "`" + `, ` + "`" + `password` + "`" + ` and ` + "`" + `is_admin` + "`" + ` parameters.`,
		Namespace: "rdb",
		Resource:  "user",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance in which you want to create a user`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the user you want to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "password",
				Short:      `Password of the user you want to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "is-admin",
				Short:      `Defines whether the user will have administrative privileges`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateUserRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateUser(request)

		},
	}
}

func rdbUserUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a user on a Database Instance`,
		Long:      `Update the parameters of a user on a Database Instance. You can update the ` + "`" + `password` + "`" + ` and ` + "`" + `is_admin` + "`" + ` parameters, but you cannot change the name of the user.`,
		Namespace: "rdb",
		Resource:  "user",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.UpdateUserRequest{}),
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
			{
				Name:       "is-admin",
				Short:      `Defines whether or not this user got administrative privileges`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpdateUserRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpdateUser(request)

		},
	}
}

func rdbUserDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a user on a Database Instance`,
		Long:      `Delete a given user on a Database Instance. You must specify, in the endpoint,  the ` + "`" + `region` + "`" + `, ` + "`" + `instance_id` + "`" + ` and ` + "`" + `name` + "`" + ` parameters of the user you want to delete.`,
		Namespace: "rdb",
		Resource:  "user",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance to delete the user from`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the user`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteUserRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			e = api.DeleteUser(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "user",
				Verb:     "delete",
			}, nil
		},
	}
}

func rdbDatabaseList() *core.Command {
	return &core.Command{
		Short:     `List databases in a Database Instance`,
		Long:      `List all databases of a given Database Instance. By default, the databases returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as ` + "`" + `name` + "`" + `, ` + "`" + `managed` + "`" + ` and ` + "`" + `owner` + "`" + `.`,
		Namespace: "rdb",
		Resource:  "database",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListDatabasesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the database`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "managed",
				Short:      `Defines whether or not the database is managed`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "owner",
				Short:      `User that owns this database`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering database listing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc", "size_asc", "size_desc"},
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance to list the databases of`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListDatabasesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListDatabases(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Databases, nil

		},
	}
}

func rdbDatabaseCreate() *core.Command {
	return &core.Command{
		Short:     `Create a database in a Database Instance`,
		Long:      `Create a new database. You must define the ` + "`" + `name` + "`" + ` parameter in the request.`,
		Namespace: "rdb",
		Resource:  "database",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateDatabaseRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance where to create the database`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the database`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateDatabaseRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateDatabase(request)

		},
	}
}

func rdbDatabaseDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a database in a Database Instance`,
		Long:      `Delete a given database on a Database Instance. You must specify, in the endpoint, the ` + "`" + `region` + "`" + `, ` + "`" + `instance_id` + "`" + ` and ` + "`" + `name` + "`" + ` parameters of the database you want to delete.`,
		Namespace: "rdb",
		Resource:  "database",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteDatabaseRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance where to delete the database`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the database to delete`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteDatabaseRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			e = api.DeleteDatabase(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "database",
				Verb:     "delete",
			}, nil
		},
	}
}

func rdbPrivilegeList() *core.Command {
	return &core.Command{
		Short:     `List user privileges for a database`,
		Long:      `List privileges of a user on a database. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as ` + "`" + `database_name` + "`" + ` and ` + "`" + `user_name` + "`" + `.`,
		Namespace: "rdb",
		Resource:  "privilege",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListPrivilegesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering privileges listing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"user_name_asc", "user_name_desc", "database_name_asc", "database_name_desc"},
			},
			{
				Name:       "database-name",
				Short:      `Name of the database`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-name",
				Short:      `Name of the user`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.ListPrivilegesRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListPrivileges(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Privileges, nil

		},
	}
}

func rdbPrivilegeSet() *core.Command {
	return &core.Command{
		Short:     `Set user privileges for a database`,
		Long:      `Set the privileges of a user on a database. You must define ` + "`" + `database_name` + "`" + `, ` + "`" + `user_name` + "`" + ` and ` + "`" + `permission` + "`" + ` in the request body.`,
		Namespace: "rdb",
		Resource:  "privilege",
		Verb:      "set",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.SetPrivilegeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "database-name",
				Short:      `Name of the database`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-name",
				Short:      `Name of the user`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "permission",
				Short:      `Permission to set (Read, Read/Write, All, Custom)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"readonly", "readwrite", "all", "custom", "none"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.SetPrivilegeRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.SetPrivilege(request)

		},
	}
}

func rdbSnapshotList() *core.Command {
	return &core.Command{
		Short:     `List snapshots`,
		Long:      `List snapshots. You can include the ` + "`" + `instance_id` + "`" + ` or ` + "`" + `project_id` + "`" + ` in your query to get the list of snapshots for specific Database Instances and/or Projects. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "rdb",
		Resource:  "snapshot",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.ListSnapshotsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Criteria to use when ordering snapshot listing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc", "expires_at_asc", "expires_at_desc"},
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Project ID the snapshots belongs to`,
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
			request := args.(*rdb.ListSnapshotsRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
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

func rdbSnapshotGet() *core.Command {
	return &core.Command{
		Short:     `Get a Database Instance snapshot`,
		Long:      `Retrieve information about a given snapshot, specified by its ` + "`" + `snapshot_id` + "`" + ` and ` + "`" + `region` + "`" + `. Full details about the snapshot, like size and expiration date, are returned in the response.`,
		Namespace: "rdb",
		Resource:  "snapshot",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.GetSnapshotRequest{}),
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
			request := args.(*rdb.GetSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetSnapshot(request)

		},
	}
}

func rdbSnapshotCreate() *core.Command {
	return &core.Command{
		Short:     `Create a Database Instance snapshot`,
		Long:      `Create a new snapshot of a Database Instance. You must define the ` + "`" + `name` + "`" + ` parameter in the request.`,
		Namespace: "rdb",
		Resource:  "snapshot",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the snapshot`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("snp"),
			},
			{
				Name:       "expires-at",
				Short:      `Expiration date (must follow the ISO 8601 format)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateSnapshot(request)

		},
	}
}

func rdbSnapshotUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a Database Instance snapshot`,
		Long:      `Update the parameters of a snapshot of a Database Instance. You can update the ` + "`" + `name` + "`" + ` and ` + "`" + `expires_at` + "`" + ` parameters.`,
		Namespace: "rdb",
		Resource:  "snapshot",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.UpdateSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot to update`,
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
				Short:      `Expiration date (must follow the ISO 8601 format)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.UpdateSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.UpdateSnapshot(request)

		},
	}
}

func rdbSnapshotDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Database Instance snapshot`,
		Long:      `Delete a given snapshot of a Database Instance. You must specify, in the endpoint,  the ` + "`" + `region` + "`" + ` and ` + "`" + `snapshot_id` + "`" + ` parameters of the snapshot you want to delete.`,
		Namespace: "rdb",
		Resource:  "snapshot",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `UUID of the snapshot to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.DeleteSnapshot(request)

		},
	}
}

func rdbSnapshotRestore() *core.Command {
	return &core.Command{
		Short:     `Create a new Database Instance from a snapshot`,
		Long:      `Restore a snapshot. When you restore a snapshot, a new Instance is created and billed to your account. Note that is possible to select a larger node type for your new Database Instance. However, the Block volume size will be the same as the size of the restored snapshot. All Instance settings will be restored if you chose a node type with the same or more memory size than the initial Instance. Settings will be reset to the default if your node type has less memory.`,
		Namespace: "rdb",
		Resource:  "snapshot",
		Verb:      "restore",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateInstanceFromSnapshotRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "snapshot-id",
				Short:      `Block snapshot of the Database Instance`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "instance-name",
				Short:      `Name of the Database Instance created with the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "is-ha-cluster",
				Short:      `Defines whether or not High-Availability is enabled on the new Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "node-type",
				Short:      `The node type used to restore the snapshot`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateInstanceFromSnapshotRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateInstanceFromSnapshot(request)

		},
	}
}

func rdbEndpointCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new Database Instance endpoint`,
		Long:      `Create a new endpoint for a Database Instance. You can add ` + "`" + `load_balancer` + "`" + ` and ` + "`" + `private_network` + "`" + ` specifications to the body of the request.`,
		Namespace: "rdb",
		Resource:  "endpoint",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.CreateEndpointRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "instance-id",
				Short:      `UUID of the Database Instance you to which you want to add an endpoint`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "endpoint-spec.private-network.private-network-id",
				Short:      `UUID of the Private Network to be connected to the Database Instance`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "endpoint-spec.private-network.service-ip",
				Short:      `Endpoint IPv4 address with a CIDR notation. Refer to the official Scaleway documentation to learn more about IP and subnet limitations.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.CreateEndpointRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.CreateEndpoint(request)

		},
	}
}

func rdbEndpointDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Database Instance endpoint`,
		Long:      `Delete the endpoint of a Database Instance. You must specify the ` + "`" + `region` + "`" + ` and ` + "`" + `endpoint_id` + "`" + ` parameters of the endpoint you want to delete. Note that might need to update any environment configurations that point to the deleted endpoint.`,
		Namespace: "rdb",
		Resource:  "endpoint",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.DeleteEndpointRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "endpoint-id",
				Short:      `UUID of the endpoint you want to delete`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.DeleteEndpointRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			e = api.DeleteEndpoint(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "endpoint",
				Verb:     "delete",
			}, nil
		},
	}
}

func rdbEndpointGet() *core.Command {
	return &core.Command{
		Short:     `Get a Database Instance endpoint`,
		Long:      `Retrieve information about a Database Instance endpoint. Full details about the endpoint, like ` + "`" + `ip` + "`" + `, ` + "`" + `port` + "`" + `, ` + "`" + `private_network` + "`" + ` and ` + "`" + `load_balancer` + "`" + ` specifications are returned in the response.`,
		Namespace: "rdb",
		Resource:  "endpoint",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.GetEndpointRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "endpoint-id",
				Short:      `UUID of the endpoint you want to get`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.GetEndpointRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.GetEndpoint(request)

		},
	}
}

func rdbEndpointMigrate() *core.Command {
	return &core.Command{
		Short:     `Migrate an existing instance endpoint to another instance`,
		Long:      `Migrate an existing instance endpoint to another instance.`,
		Namespace: "rdb",
		Resource:  "endpoint",
		Verb:      "migrate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(rdb.MigrateEndpointRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "endpoint-id",
				Short:      `UUID of the endpoint you want to migrate`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "instance-id",
				Short:      `UUID of the instance you want to attach the endpoint to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*rdb.MigrateEndpointRequest)

			client := core.ExtractClient(ctx)
			api := rdb.NewAPI(client)
			return api.MigrateEndpoint(request)

		},
	}
}

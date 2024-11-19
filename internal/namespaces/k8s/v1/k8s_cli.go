// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package k8s

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		k8sRoot(),
		k8sCluster(),
		k8sPool(),
		k8sNode(),
		k8sVersion(),
		k8sClusterType(),
		k8sACL(),
		k8sClusterList(),
		k8sClusterCreate(),
		k8sClusterGet(),
		k8sClusterUpdate(),
		k8sClusterDelete(),
		k8sClusterUpgrade(),
		k8sClusterSetType(),
		k8sClusterListAvailableVersions(),
		k8sClusterListAvailableTypes(),
		k8sClusterResetAdminToken(),
		k8sClusterMigrateToSbsCsi(),
		k8sPoolList(),
		k8sPoolCreate(),
		k8sPoolGet(),
		k8sPoolUpgrade(),
		k8sPoolUpdate(),
		k8sPoolDelete(),
		k8sNodeList(),
		k8sNodeGet(),
		k8sNodeReplace(),
		k8sNodeReboot(),
		k8sNodeDelete(),
		k8sVersionList(),
		k8sVersionGet(),
		k8sClusterTypeList(),
	)
}
func k8sRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage Kubernetes Kapsule and Kosmos clusters`,
		Long:      `This API allows you to manage Kubernetes Kapsule and Kosmos clusters.`,
		Namespace: "k8s",
	}
}

func k8sCluster() *core.Command {
	return &core.Command{
		Short: `Kapsule cluster management commands`,
		Long: `A cluster is a fully managed Kubernetes cluster
It is composed of different pools, each pool containing the same kind of nodes.`,
		Namespace: "k8s",
		Resource:  "cluster",
	}
}

func k8sPool() *core.Command {
	return &core.Command{
		Short: `Kapsule pool management commands`,
		Long: `A pool is a set of identical nodes
A pool has a name, a size (its desired number of nodes), node number limits (min, max), and a Scaleway Instance type. Changing those limits increases/decreases the size of a pool. As a result and depending on its load, the pool will grow or shrink within those limits when autoscaling is enabled. A "default pool" is automatically created with every cluster via the console.`,
		Namespace: "k8s",
		Resource:  "pool",
	}
}

func k8sNode() *core.Command {
	return &core.Command{
		Short: `Kapsule node management commands`,
		Long: `A node (short for worker node) is an abstraction for a Scaleway Instance
A node is always part of a pool. Each of them has the Kubernetes software automatically installed and configured by Scaleway.`,
		Namespace: "k8s",
		Resource:  "node",
	}
}

func k8sVersion() *core.Command {
	return &core.Command{
		Short: `Available Kubernetes versions commands`,
		Long: `A version is a vanilla Kubernetes version like ` + "`" + `x.y.z` + "`" + `
It comprises a major version ` + "`" + `x` + "`" + `, a minor version ` + "`" + `y` + "`" + `, and a patch version ` + "`" + `z` + "`" + `. At the minimum, Kapsule (Scaleway's managed Kubernetes), will support the last patch version for the past three minor releases. Also, each version has a different set of CNIs, eventually container runtimes, feature gates, and admission plugins available. See our [Version Support Policy](https://www.scaleway.com/en/docs/containers/kubernetes/reference-content/version-support-policy/).`,
		Namespace: "k8s",
		Resource:  "version",
	}
}

func k8sClusterType() *core.Command {
	return &core.Command{
		Short: `Cluster type management commands`,
		Long: `All cluster types available in a specified region
A cluster type represents the different commercial types of clusters offered by Scaleway.`,
		Namespace: "k8s",
		Resource:  "cluster-type",
	}
}

func k8sACL() *core.Command {
	return &core.Command{
		Short:     `Access Control List (ACL) management commands`,
		Long:      `Network Access Control Lists (ACLs) allow you to manage inbound network traffic by setting up ACL rules.`,
		Namespace: "k8s",
		Resource:  "acl",
	}
}

func k8sClusterList() *core.Command {
	return &core.Command{
		Short:     `List Clusters`,
		Long:      `List all existing Kubernetes clusters in a specific region.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.ListClustersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "project-id",
				Short:      `Project ID on which to filter the returned clusters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Sort order of returned clusters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "name_asc", "name_desc", "status_asc", "status_desc", "version_asc", "version_desc"},
			},
			{
				Name:       "name",
				Short:      `Name to filter on, only clusters containing this substring in their name will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "status",
				Short:      `Status to filter on, only clusters with this status will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "creating", "ready", "deleting", "deleted", "updating", "locked", "pool_required"},
			},
			{
				Name:       "type",
				Short:      `Type to filter on, only clusters with this type will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `Private Network ID to filter on, only clusters within this Private Network will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Organization ID on which to filter the returned clusters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ListClustersRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListClusters(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Clusters, nil

		},
		Examples: []*core.Example{
			{
				Short: "List all clusters on your default region",
				Raw:   `scw k8s cluster list`,
			},
			{
				Short: "List the ready clusters on your default region",
				Raw:   `scw k8s cluster list status=ready`,
			},
			{
				Short: "List the clusters that match the given name on fr-par ('cluster1' will return 'cluster100' and 'cluster1' but not 'foo')",
				Raw:   `scw k8s cluster list region=fr-par name=cluster1`,
			},
		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "Status",
			},
			{
				FieldName: "Version",
			},
			{
				FieldName: "Region",
			},
			{
				FieldName: "ProjectID",
			},
			{
				FieldName: "Tags",
			},
			{
				FieldName: "Cni",
			},
			{
				FieldName: "Description",
			},
			{
				FieldName: "ClusterURL",
			},
			{
				FieldName: "CreatedAt",
			},
			{
				FieldName: "UpdatedAt",
			},
			{
				FieldName: "Type",
			},
		}},
	}
}

func k8sClusterCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new Cluster`,
		Long:      `Create a new Kubernetes cluster in a Scaleway region.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.CreateClusterRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "type",
				Short:      `Type of the cluster (possible values are kapsule, multicloud, kapsule-dedicated-8, kapsule-dedicated-16)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Cluster name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("k8s"),
			},
			{
				Name:       "description",
				Short:      `Cluster description`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags associated with the cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "version",
				Short:      `Kubernetes version of the cluster`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "cni",
				Short:      `Container Network Interface (CNI) plugin running in the cluster`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_cni", "cilium", "calico", "weave", "flannel", "kilo", "none"},
			},
			{
				Name:       "pools.{index}.name",
				Short:      `Name of the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.node-type",
				Short:      `Node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers in a Kosmos Cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.placement-group-id",
				Short:      `Placement group ID in which all the nodes of the pool will be created`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.autoscaling",
				Short:      `Defines whether the autoscaling feature is enabled for the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.size",
				Short:      `Size (number of nodes) of the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.min-size",
				Short:      `Defines the minimum size of the pool. Note that this field is only used when autoscaling is enabled on the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.max-size",
				Short:      `Defines the maximum size of the pool. Note that this field is only used when autoscaling is enabled on the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.container-runtime",
				Short:      `Customization of the container runtime is available for each pool. Note that ` + "`" + `docker` + "`" + ` has been deprecated since version 1.20 and will be removed by version 1.24`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_runtime", "docker", "containerd", "crio"},
			},
			{
				Name:       "pools.{index}.autohealing",
				Short:      `Defines whether the autohealing feature is enabled for the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.tags.{index}",
				Short:      `Tags associated with the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.kubelet-args.{key}",
				Short:      `Kubelet arguments to be used by this pool. Note that this feature is experimental`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.upgrade-policy.max-unavailable",
				Short:      `The maximum number of nodes that can be not ready at the same time`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.upgrade-policy.max-surge",
				Short:      `The maximum number of nodes to be created during the upgrade`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.zone",
				Short:      `Zone in which the pool's nodes will be spawned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.root-volume-type",
				Short:      `Defines the system volume disk type. Two different types of volume (` + "`" + `volume_type` + "`" + `) are provided: ` + "`" + `l_ssd` + "`" + ` is a local block storage which means your system is stored locally on your node's hypervisor. ` + "`" + `b_ssd` + "`" + ` is a remote block storage which means your system is stored on a centralized and resilient cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"default_volume_type", "l_ssd", "b_ssd", "sbs_5k", "sbs_15k"},
			},
			{
				Name:       "pools.{index}.root-volume-size",
				Short:      `System volume disk size`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pools.{index}.public-ip-disabled",
				Short:      `Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.scale-down-disabled",
				Short:      `Disable the cluster autoscaler`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.scale-down-delay-after-add",
				Short:      `How long after scale up that scale down evaluation resumes`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.estimator",
				Short:      `Type of resource estimator to be used in scale up`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_estimator", "binpacking"},
			},
			{
				Name:       "autoscaler-config.expander",
				Short:      `Type of node group expander to be used in scale up`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_expander", "random", "most_pods", "least_waste", "priority", "price"},
			},
			{
				Name:       "autoscaler-config.ignore-daemonsets-utilization",
				Short:      `Ignore DaemonSet pods when calculating resource utilization for scaling down`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.balance-similar-node-groups",
				Short:      `Detect similar node groups and balance the number of nodes between them`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.expendable-pods-priority-cutoff",
				Short:      `Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they won't cause scale up. Pods with null priority (PodPriority disabled) are non expendable`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.scale-down-unneeded-time",
				Short:      `How long a node should be unneeded before it is eligible to be scaled down`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.scale-down-utilization-threshold",
				Short:      `Node utilization level, defined as a sum of requested resources divided by capacity, below which a node can be considered for scale down`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.max-graceful-termination-sec",
				Short:      `Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "auto-upgrade.enable",
				Short:      `Defines whether auto upgrade is enabled for the cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "auto-upgrade.maintenance-window.start-hour",
				Short:      `Start time of the two-hour maintenance window`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "auto-upgrade.maintenance-window.day",
				Short:      `Day of the week for the maintenance window`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"any", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
			{
				Name:       "feature-gates.{index}",
				Short:      `List of feature gates to enable`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "admission-plugins.{index}",
				Short:      `List of admission plugins to enable`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.issuer-url",
				Short:      `URL of the provider which allows the API server to discover public signing keys. Only URLs using the ` + "`" + `https://` + "`" + ` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com"`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.client-id",
				Short:      `A client ID that all tokens must be issued for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.username-claim",
				Short:      `JWT claim to use as the user name. The default is ` + "`" + `sub` + "`" + `, which is expected to be the end user's unique identifier. Admins can choose other claims, such as ` + "`" + `email` + "`" + ` or ` + "`" + `name` + "`" + `, depending on their provider. However, claims other than ` + "`" + `email` + "`" + ` will be prefixed with the issuer URL to prevent name collision`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.username-prefix",
				Short:      `Prefix prepended to username claims to prevent name collision (such as ` + "`" + `system:` + "`" + ` users). For example, the value ` + "`" + `oidc:` + "`" + ` will create usernames like ` + "`" + `oidc:jane.doe` + "`" + `. If this flag is not provided and ` + "`" + `username_claim` + "`" + ` is a value other than ` + "`" + `email` + "`" + `, the prefix defaults to ` + "`" + `( Issuer URL )#` + "`" + ` where ` + "`" + `( Issuer URL )` + "`" + ` is the value of ` + "`" + `issuer_url` + "`" + `. The value ` + "`" + `-` + "`" + ` can be used to disable all prefixing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.groups-claim.{index}",
				Short:      `JWT claim to use as the user's group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.groups-prefix",
				Short:      `Prefix prepended to group claims to prevent name collision (such as ` + "`" + `system:` + "`" + ` groups). For example, the value ` + "`" + `oidc:` + "`" + ` will create group names like ` + "`" + `oidc:engineering` + "`" + ` and ` + "`" + `oidc:infra` + "`" + ``,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.required-claim.{index}",
				Short:      `Multiple key=value pairs describing a required claim in the ID token. If set, the claims are verified to be present in the ID token with a matching value`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "apiserver-cert-sans.{index}",
				Short:      `Additional Subject Alternative Names for the Kubernetes API server certificate`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "private-network-id",
				Short:      `Private network ID for internal cluster communication (cannot be changed later)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.CreateClusterRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.CreateCluster(request)

		},
		Examples: []*core.Example{
			{
				Short: "Create a Kubernetes cluster named foo with cilium as CNI, in version 1.27.0 and with a pool named default composed of 3 DEV1-M",
				Raw:   `scw k8s cluster create name=foo version=1.27.0 pools.0.size=3 pools.0.node-type=DEV1-M pools.0.name=default`,
			},
			{
				Short: "Create a Kubernetes cluster named bar, tagged, calico as CNI, in version 1.27.0 and with a tagged pool named default composed of 2 RENDER-S and autohealing and autoscaling enabled (between 1 and 10 nodes)",
				Raw:   `scw k8s cluster create name=bar version=1.27.0 tags.0=tag1 tags.1=tag2 cni=calico pools.0.size=2 pools.0.node-type=RENDER-S pools.0.min-size=1 pools.0.max-size=10 pools.0.autohealing=true pools.0.autoscaling=true pools.0.tags.0=pooltag1 pools.0.tags.1=pooltag2 pools.0.name=default`,
			},
		},
	}
}

func k8sClusterGet() *core.Command {
	return &core.Command{
		Short:     `Get a Cluster`,
		Long:      `Retrieve information about a specific Kubernetes cluster.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.GetClusterRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `ID of the requested cluster`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.GetClusterRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.GetCluster(request)

		},
		Examples: []*core.Example{
			{
				Short: "Get a cluster information",
				Raw:   `scw k8s cluster get 11111111-1111-1111-111111111111`,
			},
		},
	}
}

func k8sClusterUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a Cluster`,
		Long:      `Update information on a specific Kubernetes cluster. You can update details such as its name, description, tags and configuration. To upgrade a cluster, you will need to use the dedicated endpoint.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.UpdateClusterRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `ID of the cluster to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `New external name for the cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `New description for the cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `New tags associated with the cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.scale-down-disabled",
				Short:      `Disable the cluster autoscaler`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.scale-down-delay-after-add",
				Short:      `How long after scale up that scale down evaluation resumes`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.estimator",
				Short:      `Type of resource estimator to be used in scale up`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_estimator", "binpacking"},
			},
			{
				Name:       "autoscaler-config.expander",
				Short:      `Type of node group expander to be used in scale up`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_expander", "random", "most_pods", "least_waste", "priority", "price"},
			},
			{
				Name:       "autoscaler-config.ignore-daemonsets-utilization",
				Short:      `Ignore DaemonSet pods when calculating resource utilization for scaling down`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.balance-similar-node-groups",
				Short:      `Detect similar node groups and balance the number of nodes between them`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.expendable-pods-priority-cutoff",
				Short:      `Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they won't cause scale up. Pods with null priority (PodPriority disabled) are non expendable`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.scale-down-unneeded-time",
				Short:      `How long a node should be unneeded before it is eligible to be scaled down`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.scale-down-utilization-threshold",
				Short:      `Node utilization level, defined as a sum of requested resources divided by capacity, below which a node can be considered for scale down`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaler-config.max-graceful-termination-sec",
				Short:      `Maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "auto-upgrade.enable",
				Short:      `Defines whether auto upgrade is enabled for the cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "auto-upgrade.maintenance-window.start-hour",
				Short:      `Start time of the two-hour maintenance window`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "auto-upgrade.maintenance-window.day",
				Short:      `Day of the week for the maintenance window`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"any", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"},
			},
			{
				Name:       "feature-gates.{index}",
				Short:      `List of feature gates to enable`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "admission-plugins.{index}",
				Short:      `List of admission plugins to enable`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.issuer-url",
				Short:      `URL of the provider which allows the API server to discover public signing keys. Only URLs using the ` + "`" + `https://` + "`" + ` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com"`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.client-id",
				Short:      `A client ID that all tokens must be issued for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.username-claim",
				Short:      `JWT claim to use as the user name. The default is ` + "`" + `sub` + "`" + `, which is expected to be the end user's unique identifier. Admins can choose other claims, such as ` + "`" + `email` + "`" + ` or ` + "`" + `name` + "`" + `, depending on their provider. However, claims other than ` + "`" + `email` + "`" + ` will be prefixed with the issuer URL to prevent name collision`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.username-prefix",
				Short:      `Prefix prepended to username claims to prevent name collision (such as ` + "`" + `system:` + "`" + ` users). For example, the value ` + "`" + `oidc:` + "`" + ` will create usernames like ` + "`" + `oidc:jane.doe` + "`" + `. If this flag is not provided and ` + "`" + `username_claim` + "`" + ` is a value other than ` + "`" + `email` + "`" + `, the prefix defaults to ` + "`" + `( Issuer URL )#` + "`" + ` where ` + "`" + `( Issuer URL )` + "`" + ` is the value of ` + "`" + `issuer_url` + "`" + `. The value ` + "`" + `-` + "`" + ` can be used to disable all prefixing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.groups-claim.{index}",
				Short:      `JWT claim to use as the user's group`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.groups-prefix",
				Short:      `Prefix prepended to group claims to prevent name collision (such as ` + "`" + `system:` + "`" + ` groups). For example, the value ` + "`" + `oidc:` + "`" + ` will create group names like ` + "`" + `oidc:engineering` + "`" + ` and ` + "`" + `oidc:infra` + "`" + ``,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "open-id-connect-config.required-claim.{index}",
				Short:      `Multiple key=value pairs describing a required claim in the ID token. If set, the claims are verified to be present in the ID token with a matching value`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "apiserver-cert-sans.{index}",
				Short:      `Additional Subject Alternative Names for the Kubernetes API server certificate`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.UpdateClusterRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.UpdateCluster(request)

		},
		Examples: []*core.Example{
			{
				Short: "Add TTLAfterFinished and ServiceNodeExclusion as feature gates on a cluster",
				Raw:   `scw k8s cluster update 11111111-1111-1111-111111111111 feature-gates.0=TTLAfterFinished feature-gates.1=ServiceNodeExclusion`,
			},
		},
	}
}

func k8sClusterDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Cluster`,
		Long:      `Delete a specific Kubernetes cluster and all its associated pools and nodes. Note that this method will not delete any Load Balancer or Block Volume that are associated with the cluster.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.DeleteClusterRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `ID of the cluster to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "with-additional-resources",
				Short:      `Defines whether all volumes (including retain volume type), empty Private Networks and Load Balancers with a name starting with the cluster ID will also be deleted`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.DeleteClusterRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.DeleteCluster(request)

		},
		Examples: []*core.Example{
			{
				Short: "Delete a cluster",
				Raw:   `scw k8s cluster delete 11111111-1111-1111-111111111111`,
			},
			{
				Short: "Delete a cluster with its Block volumes and Load Balancers",
				Raw:   `scw k8s cluster delete 11111111-1111-1111-111111111111 with-additional-resources=true`,
			},
		},
	}
}

func k8sClusterUpgrade() *core.Command {
	return &core.Command{
		Short:     `Upgrade a Cluster`,
		Long:      `Upgrade a specific Kubernetes cluster and possibly its associated pools to a specific and supported Kubernetes version.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "upgrade",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.UpgradeClusterRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `ID of the cluster to upgrade`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "version",
				Short:      `New Kubernetes version of the cluster. Note that the version should either be a higher patch version of the same minor version or the direct minor version after the current one`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "upgrade-pools",
				Short:      `Defines whether pools will also be upgraded once the control plane is upgraded`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.UpgradeClusterRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.UpgradeCluster(request)

		},
		Examples: []*core.Example{
			{
				Short: "Upgrade a cluster version 1.27.0 of Kubernetes (pools *are not* included)",
				Raw:   `scw k8s cluster upgrade 11111111-1111-1111-111111111111 version=1.27.0`,
			},
			{
				Short: "Upgrade a cluster to version 1.27.0 of Kubernetes (pools *are* included)",
				Raw:   `scw k8s cluster upgrade 11111111-1111-1111-111111111111 version=1.27.0 upgrade-pools=true`,
			},
		},
	}
}

func k8sClusterSetType() *core.Command {
	return &core.Command{
		Short:     `Change the Cluster type`,
		Long:      `Change the type of a specific Kubernetes cluster. To see the possible values you can enter for the ` + "`" + `type` + "`" + ` field, [list available cluster types](#path-clusters-list-available-cluster-types-for-a-cluster).`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "set-type",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.SetClusterTypeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `ID of the cluster to migrate from one type to another`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "type",
				Short:      `Type of the cluster. Note that some migrations are not possible (please refer to product documentation)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.SetClusterTypeRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.SetClusterType(request)

		},
		Examples: []*core.Example{
			{
				Short: "Convert a kapsule cluster to a kapsule-dedicated-16 cluster",
				Raw:   `scw k8s cluster set-type 11111111-1111-1111-111111111111 type=kapsule-dedicated-16`,
			},
		},
	}
}

func k8sClusterListAvailableVersions() *core.Command {
	return &core.Command{
		Short:     `List available versions for a Cluster`,
		Long:      `List the versions that a specific Kubernetes cluster is allowed to upgrade to. Results will include every patch version greater than the current patch, as well as one minor version ahead of the current version. Any upgrade skipping a minor version will not work.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "list-available-versions",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.ListClusterAvailableVersionsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `Cluster ID for which the available Kubernetes versions will be listed`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ListClusterAvailableVersionsRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.ListClusterAvailableVersions(request)

		},
		Examples: []*core.Example{
			{
				Short: "List all versions that a cluster can upgrade to",
				Raw:   `scw k8s cluster list-available-versions 11111111-1111-1111-111111111111`,
			},
		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "Name",
			},
			{
				FieldName: "Label",
			},
			{
				FieldName: "AvailableContainerRuntimes",
			},
		}},
	}
}

func k8sClusterListAvailableTypes() *core.Command {
	return &core.Command{
		Short:     `List available cluster types for a cluster`,
		Long:      `List the cluster types that a specific Kubernetes cluster is allowed to switch to.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "list-available-types",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.ListClusterAvailableTypesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `Cluster ID for which the available Kubernetes types will be listed`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ListClusterAvailableTypesRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.ListClusterAvailableTypes(request)

		},
		Examples: []*core.Example{
			{
				Short: "List all cluster types that a cluster can upgrade to",
				Raw:   `scw k8s cluster list-available-types 11111111-1111-1111-111111111111`,
			},
		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "Name",
			},
		}},
	}
}

func k8sClusterResetAdminToken() *core.Command {
	return &core.Command{
		Short:     `Reset the admin token of a Cluster`,
		Long:      `Reset the admin token for a specific Kubernetes cluster. This will revoke the old admin token (which will not be usable afterwards) and create a new one. Note that you will need to download kubeconfig again to keep interacting with the cluster.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "reset-admin-token",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.ResetClusterAdminTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `Cluster ID on which the admin token will be renewed`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ResetClusterAdminTokenRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			e = api.ResetClusterAdminToken(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "cluster",
				Verb:     "reset-admin-token",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short: "Reset the admin token for a cluster",
				Raw:   `scw k8s cluster reset-admin-token 11111111-1111-1111-111111111111`,
			},
		},
	}
}

func k8sClusterMigrateToSbsCsi() *core.Command {
	return &core.Command{
		Short:     `Migrate a cluster to SBS CSI`,
		Long:      `Enable the latest CSI compatible with Scaleway Block Storage (SBS) and migrate all existing PersistentVolumes/VolumeSnapshotContents to SBS.`,
		Namespace: "k8s",
		Resource:  "cluster",
		Verb:      "migrate-to-sbs-csi",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.MigrateClusterToSBSCSIRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `Cluster ID for which the latest CSI compatible with Scaleway Block Storage will be enabled`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.MigrateClusterToSBSCSIRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.MigrateClusterToSBSCSI(request)

		},
		Examples: []*core.Example{
			{
				Short: "Migrate a cluster to SBS CSI",
				Raw:   `scw k8s cluster migrate-to-sbs-csi 11111111-1111-1111-111111111111`,
			},
		},
	}
}

func k8sPoolList() *core.Command {
	return &core.Command{
		Short:     `List Pools in a Cluster`,
		Long:      `List all the existing pools for a specific Kubernetes cluster.`,
		Namespace: "k8s",
		Resource:  "pool",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.ListPoolsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `ID of the cluster whose pools will be listed`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Sort order of returned pools`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "name_asc", "name_desc", "status_asc", "status_desc", "version_asc", "version_desc"},
			},
			{
				Name:       "name",
				Short:      `Name to filter on, only pools containing this substring in their name will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "status",
				Short:      `Status to filter on, only pools with this status will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "ready", "deleting", "deleted", "scaling", "warning", "locked", "upgrading"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ListPoolsRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListPools(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Pools, nil

		},
		Examples: []*core.Example{
			{
				Short: "List all pools for a cluster",
				Raw:   `scw k8s pool list cluster-id=11111111-1111-1111-111111111111`,
			},
			{
				Short: "List all scaling pools for a cluster",
				Raw:   `scw k8s pool list cluster-id=11111111-1111-1111-111111111111 status=scaling`,
			},
			{
				Short: "List all pools for clusters containing 'foo' in their name",
				Raw:   `scw k8s pool list cluster-id=11111111-1111-1111-111111111111 name=foo`,
			},
			{
				Short: "List all pools for a cluster and order them by ascending creation date",
				Raw:   `scw k8s pool list cluster-id=11111111-1111-1111-111111111111 order-by=created_at_asc`,
			},
		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "Status",
			},
			{
				FieldName: "NodeType",
			},
			{
				FieldName: "Size",
			},
			{
				FieldName: "MinSize",
			},
			{
				FieldName: "MaxSize",
			},
			{
				FieldName: "Autoscaling",
			},
			{
				FieldName: "Autohealing",
			},
			{
				FieldName: "Version",
			},
			{
				FieldName: "Tags",
			},
			{
				FieldName: "ContainerRuntime",
			},
			{
				FieldName: "ClusterID",
			},
			{
				FieldName: "Zone",
			},
			{
				FieldName: "Region",
			},
			{
				FieldName: "KubeletArgs",
			},
			{
				FieldName: "PlacementGroupID",
			},
			{
				FieldName: "CreatedAt",
			},
			{
				FieldName: "UpdatedAt",
			},
		}},
	}
}

func k8sPoolCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new Pool in a Cluster`,
		Long:      `Create a new pool in a specific Kubernetes cluster.`,
		Namespace: "k8s",
		Resource:  "pool",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.CreatePoolRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `Cluster ID to which the pool will be attached`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Pool name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("pool"),
			},
			{
				Name:       "node-type",
				Short:      `Node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers in a Kosmos Cluster`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "placement-group-id",
				Short:      `Placement group ID in which all the nodes of the pool will be created`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autoscaling",
				Short:      `Defines whether the autoscaling feature is enabled for the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "size",
				Short:      `Size (number of nodes) of the pool`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "min-size",
				Short:      `Defines the minimum size of the pool. Note that this field is only used when autoscaling is enabled on the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-size",
				Short:      `Defines the maximum size of the pool. Note that this field is only used when autoscaling is enabled on the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "container-runtime",
				Short:      `Customization of the container runtime is available for each pool. Note that ` + "`" + `docker` + "`" + ` has been deprecated since version 1.20 and will be removed by version 1.24`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_runtime", "docker", "containerd", "crio"},
			},
			{
				Name:       "autohealing",
				Short:      `Defines whether the autohealing feature is enabled for the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags associated with the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "kubelet-args.{key}",
				Short:      `Kubelet arguments to be used by this pool. Note that this feature is experimental`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "upgrade-policy.max-unavailable",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "upgrade-policy.max-surge",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "zone",
				Short:      `Zone in which the pool's nodes will be spawned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "root-volume-type",
				Short:      `Defines the system volume disk type. Two different types of volume (` + "`" + `volume_type` + "`" + `) are provided: ` + "`" + `l_ssd` + "`" + ` is a local block storage which means your system is stored locally on your node's hypervisor. ` + "`" + `b_ssd` + "`" + ` is a remote block storage which means your system is stored on a centralized and resilient cluster`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"default_volume_type", "l_ssd", "b_ssd", "sbs_5k", "sbs_15k"},
			},
			{
				Name:       "root-volume-size",
				Short:      `System volume disk size`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "public-ip-disabled",
				Short:      `Defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.CreatePoolRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.CreatePool(request)

		},
		Examples: []*core.Example{
			{
				Short: "Create a pool named 'bar' with 2 DEV1-XL on a cluster",
				Raw:   `scw k8s pool create cluster-id=11111111-1111-1111-111111111111 name=bar node-type=DEV1-XL size=2`,
			},
			{
				Short: "Create a pool named 'fish' with 5 GP1-L, autoscaling within 0 and 10 nodes and autohealing enabled, and containerd as the cluster container runtime",
				Raw:   `scw k8s pool create cluster-id=11111111-1111-1111-111111111111 name=fish node-type=GP1-L size=5 min-size=0 max-size=10 autoscaling=true autohealing=true container-runtime=containerd`,
			},
			{
				Short: "Create a tagged pool named 'turtle' with 1 GP1-S which is using the already created placement group 2222222222222-2222-222222222222 for all the nodes in the pool on a cluster",
				Raw:   `scw k8s pool create cluster-id=11111111-1111-1111-111111111111 name=turtle node-type=GP1-S size=1 placement-group-id=2222222222222-2222-222222222222 tags.0=turtle tags.1=placement-group`,
			},
		},
	}
}

func k8sPoolGet() *core.Command {
	return &core.Command{
		Short:     `Get a Pool in a Cluster`,
		Long:      `Retrieve details about a specific pool in a Kubernetes cluster.`,
		Namespace: "k8s",
		Resource:  "pool",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.GetPoolRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "pool-id",
				Short:      `ID of the requested pool`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.GetPoolRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.GetPool(request)

		},
		Examples: []*core.Example{
			{
				Short: "Get a given pool",
				Raw:   `scw k8s pool get 11111111-1111-1111-111111111111`,
			},
		},
	}
}

func k8sPoolUpgrade() *core.Command {
	return &core.Command{
		Short:     `Upgrade a Pool in a Cluster`,
		Long:      `Upgrade the Kubernetes version of a specific pool. Note that it only works if the targeted version matches the cluster's version.`,
		Namespace: "k8s",
		Resource:  "pool",
		Verb:      "upgrade",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.UpgradePoolRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "pool-id",
				Short:      `ID of the pool to upgrade`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "version",
				Short:      `New Kubernetes version for the pool`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.UpgradePoolRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.UpgradePool(request)

		},
		Examples: []*core.Example{
			{
				Short: "Upgrade a specific pool to the Kubernetes version 1.27.0",
				Raw:   `scw k8s pool upgrade 11111111-1111-1111-111111111111 version=1.27.0`,
			},
		},
	}
}

func k8sPoolUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a Pool in a Cluster`,
		Long:      `Update the attributes of a specific pool, such as its desired size, autoscaling settings, and tags.`,
		Namespace: "k8s",
		Resource:  "pool",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.UpdatePoolRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "pool-id",
				Short:      `ID of the pool to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "autoscaling",
				Short:      `New value for the pool autoscaling enablement`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "size",
				Short:      `New desired pool size`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "min-size",
				Short:      `New minimum size for the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-size",
				Short:      `New maximum size for the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "autohealing",
				Short:      `New value for the pool autohealing enablement`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `New tags associated with the pool`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "kubelet-args.{key}",
				Short:      `New Kubelet arguments to be used by this pool. Note that this feature is experimental`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "upgrade-policy.max-unavailable",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "upgrade-policy.max-surge",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.UpdatePoolRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.UpdatePool(request)

		},
		Examples: []*core.Example{
			{
				Short: "Enable autoscaling on a given pool",
				Raw:   `scw k8s pool update 11111111-1111-1111-111111111111 autoscaling=true`,
			},
			{
				Short: "Reduce the size and maximum size of a given pool to 4",
				Raw:   `scw k8s pool update 11111111-1111-1111-111111111111 size=4 max-size=4`,
			},
			{
				Short: "Modify the tags of a given pool",
				Raw:   `scw k8s pool update 11111111-1111-1111-111111111111 tags.0=my tags.1=new tags.2=pool`,
			},
		},
	}
}

func k8sPoolDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Pool in a Cluster`,
		Long:      `Delete a specific pool from a cluster. Note that all the pool's nodes will also be deleted.`,
		Namespace: "k8s",
		Resource:  "pool",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.DeletePoolRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "pool-id",
				Short:      `ID of the pool to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.DeletePoolRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.DeletePool(request)

		},
		Examples: []*core.Example{
			{
				Short: "Delete a specific pool",
				Raw:   `scw k8s pool delete 11111111-1111-1111-111111111111`,
			},
		},
	}
}

func k8sNodeList() *core.Command {
	return &core.Command{
		Short:     `List Nodes in a Cluster`,
		Long:      `List all the existing nodes for a specific Kubernetes cluster.`,
		Namespace: "k8s",
		Resource:  "node",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.ListNodesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      `Cluster ID from which the nodes will be listed from`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "pool-id",
				Short:      `Pool ID on which to filter the returned nodes`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Sort order of the returned nodes`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "name",
				Short:      `Name to filter on, only nodes containing this substring in their name will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "status",
				Short:      `Status to filter on, only nodes with this status will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "creating", "not_ready", "ready", "deleting", "deleted", "locked", "rebooting", "creation_error", "upgrading", "starting", "registering"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ListNodesRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListNodes(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Nodes, nil

		},
		Examples: []*core.Example{
			{
				Short: "List all the nodes in the cluster",
				Raw:   `scw k8s node list cluster-id=11111111-1111-1111-111111111111`,
			},
			{
				Short: "List all the nodes in the cluster's 2222222222222-2222-222222222222 pool",
				Raw:   `scw k8s node list cluster-id=11111111-1111-1111-111111111111 pool-id=2222222222222-2222-222222222222`,
			},
			{
				Short: "List all cluster nodes that are ready",
				Raw:   `scw k8s node list cluster-id=11111111-1111-1111-111111111111 status=ready`,
			},
		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "Status",
			},
			{
				FieldName: "PoolID",
			},
			{
				FieldName: "ClusterID",
			},
			{
				FieldName: "Region",
			},
			{
				FieldName: "CreatedAt",
			},
			{
				FieldName: "UpdatedAt",
			},
		}},
	}
}

func k8sNodeGet() *core.Command {
	return &core.Command{
		Short:     `Get a Node in a Cluster`,
		Long:      `Retrieve details about a specific Kubernetes Node.`,
		Namespace: "k8s",
		Resource:  "node",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.GetNodeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "node-id",
				Short:      `ID of the requested node`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.GetNodeRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.GetNode(request)

		},
		Examples: []*core.Example{
			{
				Short: "Get a node",
				Raw:   `scw k8s node get 11111111-1111-1111-111111111111`,
			},
		},
	}
}

func k8sNodeReplace() *core.Command {
	return &core.Command{
		Short:     `Replace a Node in a Cluster`,
		Long:      `Replace a specific Node. The node will first be cordoned (scheduling will be disabled on it). The existing pods on the node will then be drained and rescheduled onto another schedulable node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster), disruption of your applications can be expected.`,
		Namespace: "k8s",
		Resource:  "node",
		Verb:      "replace",
		// Deprecated:    true,
		ArgsType: reflect.TypeOf(k8s.ReplaceNodeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "node-id",
				Short:      `ID of the node to replace`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ReplaceNodeRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.ReplaceNode(request)

		},
		Examples: []*core.Example{
			{
				Short: "Replace a node",
				Raw:   `scw k8s node replace 11111111-1111-1111-111111111111`,
			},
		},
	}
}

func k8sNodeReboot() *core.Command {
	return &core.Command{
		Short:     `Reboot a Node in a Cluster`,
		Long:      `Reboot a specific Node. The node will first be cordoned (scheduling will be disabled on it). The existing pods on the node will then be drained and rescheduled onto another schedulable node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster), disruption of your applications can be expected.`,
		Namespace: "k8s",
		Resource:  "node",
		Verb:      "reboot",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.RebootNodeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "node-id",
				Short:      `ID of the node to reboot`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.RebootNodeRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.RebootNode(request)

		},
		Examples: []*core.Example{
			{
				Short: "Reboot a node",
				Raw:   `scw k8s node reboot 11111111-1111-1111-111111111111`,
			},
		},
	}
}

func k8sNodeDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Node in a Cluster`,
		Long:      `Delete a specific Node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster), disruption of your applications can be expected.`,
		Namespace: "k8s",
		Resource:  "node",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.DeleteNodeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "node-id",
				Short:      `ID of the node to replace`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "skip-drain",
				Short:      `Skip draining node from its workload (Note: this parameter is currently inactive)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "replace",
				Short:      `Add a new node after the deletion of this node`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.DeleteNodeRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.DeleteNode(request)

		},
		Examples: []*core.Example{
			{
				Short: "Delete a node",
				Raw:   `scw k8s node delete 11111111-1111-1111-111111111111`,
			},
			{
				Short: "Delete a node without evicting workloads",
				Raw:   `scw k8s node delete 11111111-1111-1111-111111111111 skip-drain=true`,
			},
			{
				Short: "Replace a node by a new one",
				Raw:   `scw k8s node delete 11111111-1111-1111-111111111111 replace=true`,
			},
		},
	}
}

func k8sVersionList() *core.Command {
	return &core.Command{
		Short:     `List all available Versions`,
		Long:      `List all available versions for the creation of a new Kubernetes cluster.`,
		Namespace: "k8s",
		Resource:  "version",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.ListVersionsRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ListVersionsRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.ListVersions(request)

		},
		Examples: []*core.Example{
			{
				Short: "List all available Kubernetes version in Kapsule",
				Raw:   `scw k8s version list`,
			},
		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "Name",
			},
			{
				FieldName: "AvailableCnis",
			},
			{
				FieldName: "AvailableContainerRuntimes",
			},
			{
				FieldName: "AvailableFeatureGates",
			},
			{
				FieldName: "AvailableAdmissionPlugins",
			},
			{
				FieldName: "AvailableKubeletArgs",
			},
		}},
	}
}

func k8sVersionGet() *core.Command {
	return &core.Command{
		Short:     `Get a Version`,
		Long:      `Retrieve a specific Kubernetes version and its details.`,
		Namespace: "k8s",
		Resource:  "version",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.GetVersionRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "version-name",
				Short:      `Requested version name`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.GetVersionRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			return api.GetVersion(request)

		},
		Examples: []*core.Example{
			{
				Short: "Get the Kubernetes version 1.27.0",
				Raw:   `scw k8s version get 1.27.0`,
			},
		},
	}
}

func k8sClusterTypeList() *core.Command {
	return &core.Command{
		Short:     `List cluster types`,
		Long:      `List available cluster types and their technical details.`,
		Namespace: "k8s",
		Resource:  "cluster-type",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(k8s.ListClusterTypesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*k8s.ListClusterTypesRequest)

			client := core.ExtractClient(ctx)
			api := k8s.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListClusterTypes(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.ClusterTypes, nil

		},
	}
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package container

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/container/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		containerRoot(),
		containerNamespace(),
		containerContainer(),
		containerCron(),
		containerDomain(),
		containerToken(),
		containerTrigger(),
		containerNamespaceList(),
		containerNamespaceGet(),
		containerNamespaceCreate(),
		containerNamespaceUpdate(),
		containerNamespaceDelete(),
		containerContainerList(),
		containerContainerGet(),
		containerContainerCreate(),
		containerContainerUpdate(),
		containerContainerDelete(),
		containerContainerDeploy(),
		containerCronList(),
		containerCronGet(),
		containerCronCreate(),
		containerCronUpdate(),
		containerCronDelete(),
		containerDomainList(),
		containerDomainGet(),
		containerDomainCreate(),
		containerDomainDelete(),
		containerTokenCreate(),
		containerTokenGet(),
		containerTokenList(),
		containerTokenDelete(),
		containerTriggerCreate(),
		containerTriggerGet(),
		containerTriggerList(),
		containerTriggerUpdate(),
		containerTriggerDelete(),
	)
}
func containerRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Serverless Containers`,
		Long:      `This API allows you to manage your Serverless Containers.`,
		Namespace: "container",
	}
}

func containerNamespace() *core.Command {
	return &core.Command{
		Short:     `Namespace management commands`,
		Long:      `Namespace management commands.`,
		Namespace: "container",
		Resource:  "namespace",
	}
}

func containerContainer() *core.Command {
	return &core.Command{
		Short:     `Container management commands`,
		Long:      `Container management commands.`,
		Namespace: "container",
		Resource:  "container",
	}
}

func containerCron() *core.Command {
	return &core.Command{
		Short:     `Cron management commands`,
		Long:      `Cron management commands.`,
		Namespace: "container",
		Resource:  "cron",
	}
}

func containerDomain() *core.Command {
	return &core.Command{
		Short:     `Domain management commands`,
		Long:      `Domain management commands.`,
		Namespace: "container",
		Resource:  "domain",
	}
}

func containerToken() *core.Command {
	return &core.Command{
		Short:     `Token management commands`,
		Long:      `Token management commands.`,
		Namespace: "container",
		Resource:  "token",
	}
}

func containerTrigger() *core.Command {
	return &core.Command{
		Short:     `Trigger management commands`,
		Long:      `Trigger management commands.`,
		Namespace: "container",
		Resource:  "trigger",
	}
}

func containerNamespaceList() *core.Command {
	return &core.Command{
		Short:     `List all your namespaces`,
		Long:      `List all namespaces in a specified region.`,
		Namespace: "container",
		Resource:  "namespace",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.ListNamespacesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the namespaces`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "name",
				Short:      `Name of the namespaces`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `UUID of the Project the namespace belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `UUID of the Organization the namespace belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.ListNamespacesRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListNamespaces(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Namespaces, nil

		},
	}
}

func containerNamespaceGet() *core.Command {
	return &core.Command{
		Short:     `Get a namespace`,
		Long:      `Get the namespace associated with the specified ID.`,
		Namespace: "container",
		Resource:  "namespace",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.GetNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.GetNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.GetNamespace(request)

		},
	}
}

func containerNamespaceCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new namespace`,
		Long:      `Create a new namespace in a specified region.`,
		Namespace: "container",
		Resource:  "namespace",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.CreateNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the namespace to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("cns"),
			},
			{
				Name:       "environment-variables.{key}",
				Short:      `Environment variables of the namespace to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "description",
				Short:      `Description of the namespace to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "secret-environment-variables.{index}.key",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "secret-environment-variables.{index}.value",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `[ALPHA] Tags of the Serverless Container Namespace`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.CreateNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.CreateNamespace(request)

		},
	}
}

func containerNamespaceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing namespace`,
		Long:      `Update the space associated with the specified ID.`,
		Namespace: "container",
		Resource:  "namespace",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.UpdateNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "environment-variables.{key}",
				Short:      `Environment variables of the namespace to update`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description of the namespace to update`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "secret-environment-variables.{index}.key",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "secret-environment-variables.{index}.value",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `[ALPHA] Tags of the Serverless Container Namespace`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.UpdateNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.UpdateNamespace(request)

		},
	}
}

func containerNamespaceDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an existing namespace`,
		Long:      `Delete the namespace associated with the specified ID.`,
		Namespace: "container",
		Resource:  "namespace",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.DeleteNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.DeleteNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.DeleteNamespace(request)

		},
	}
}

func containerContainerList() *core.Command {
	return &core.Command{
		Short:     `List all your containers`,
		Long:      `List all containers for a specified region.`,
		Namespace: "container",
		Resource:  "container",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.ListContainersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the containers`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace the container belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `UUID of the Project the container belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `UUID of the Organization the container belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.ListContainersRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListContainers(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Containers, nil

		},
	}
}

func containerContainerGet() *core.Command {
	return &core.Command{
		Short:     `Get a container`,
		Long:      `Get the container associated with the specified ID.`,
		Namespace: "container",
		Resource:  "container",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.GetContainerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "container-id",
				Short:      `UUID of the container to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.GetContainerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.GetContainer(request)

		},
	}
}

func containerContainerCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new container`,
		Long:      `Create a new container in the specified region.`,
		Namespace: "container",
		Resource:  "container",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.CreateContainerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace the container belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "environment-variables.{key}",
				Short:      `Environment variables of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "min-scale",
				Short:      `Minimum number of instances to scale the container to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-scale",
				Short:      `Maximum number of instances to scale the container to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "memory-limit",
				Short:      `Memory limit of the container in MB`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "cpu-limit",
				Short:      `CPU limit of the container in mvCPU`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "timeout",
				Short:      `Processing time limit for the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "privacy",
				Short:      `Privacy setting of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_privacy", "public", "private"},
			},
			{
				Name:       "description",
				Short:      `Description of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "registry-image",
				Short:      `Name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag").`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-concurrency",
				Short:      `Number of maximum concurrent executions of the container`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "protocol",
				Short:      `Protocol the container uses`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_protocol", "http1", "h2c"},
			},
			{
				Name:       "port",
				Short:      `Port the container listens on`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "secret-environment-variables.{index}.key",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "secret-environment-variables.{index}.value",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "http-option",
				Short:      `Configure how HTTP and HTTPS requests are handled`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("enabled"),
				EnumValues: []string{"unknown_http_option", "enabled", "redirected"},
			},
			{
				Name:       "sandbox",
				Short:      `Execution environment of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_sandbox", "v1", "v2"},
			},
			{
				Name:       "local-storage-limit",
				Short:      `Local storage limit of the container (in MB)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scaling-option.concurrent-requests-threshold",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scaling-option.cpu-usage-threshold",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scaling-option.memory-usage-threshold",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.http.path",
				Short:      `Path to use for the HTTP health check.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.failure-threshold",
				Short:      `Number of consecutive health check failures before considering the container unhealthy.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.interval",
				Short:      `Period between health checks.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.CreateContainerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.CreateContainer(request)

		},
	}
}

func containerContainerUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing container`,
		Long:      `Update the container associated with the specified ID.`,
		Namespace: "container",
		Resource:  "container",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.UpdateContainerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "container-id",
				Short:      `UUID of the container to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "environment-variables.{key}",
				Short:      `Environment variables of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "min-scale",
				Short:      `Minimum number of instances to scale the container to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-scale",
				Short:      `Maximum number of instances to scale the container to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "memory-limit",
				Short:      `Memory limit of the container in MB`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "cpu-limit",
				Short:      `CPU limit of the container in mvCPU`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "timeout",
				Short:      `Processing time limit for the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "redeploy",
				Short:      `Defines whether to redeploy failed containers`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "privacy",
				Short:      `Privacy settings of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_privacy", "public", "private"},
			},
			{
				Name:       "description",
				Short:      `Description of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "registry-image",
				Short:      `Name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag").`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-concurrency",
				Short:      `Number of maximum concurrent executions of the container`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "protocol",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_protocol", "http1", "h2c"},
			},
			{
				Name:       "port",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "secret-environment-variables.{index}.key",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "secret-environment-variables.{index}.value",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "http-option",
				Short:      `Configure how HTTP and HTTPS requests are handled`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_http_option", "enabled", "redirected"},
			},
			{
				Name:       "sandbox",
				Short:      `Execution environment of the container`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_sandbox", "v1", "v2"},
			},
			{
				Name:       "local-storage-limit",
				Short:      `Local storage limit of the container (in MB)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scaling-option.concurrent-requests-threshold",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scaling-option.cpu-usage-threshold",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scaling-option.memory-usage-threshold",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.http.path",
				Short:      `Path to use for the HTTP health check.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.failure-threshold",
				Short:      `Number of consecutive health check failures before considering the container unhealthy.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "health-check.interval",
				Short:      `Period between health checks.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.UpdateContainerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.UpdateContainer(request)

		},
	}
}

func containerContainerDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a container`,
		Long:      `Delete the container associated with the specified ID.`,
		Namespace: "container",
		Resource:  "container",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.DeleteContainerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "container-id",
				Short:      `UUID of the container to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.DeleteContainerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.DeleteContainer(request)

		},
	}
}

func containerContainerDeploy() *core.Command {
	return &core.Command{
		Short:     `Deploy a container`,
		Long:      `Deploy a container associated with the specified ID.`,
		Namespace: "container",
		Resource:  "container",
		Verb:      "deploy",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.DeployContainerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "container-id",
				Short:      `UUID of the container to deploy`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.DeployContainerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.DeployContainer(request)

		},
	}
}

func containerCronList() *core.Command {
	return &core.Command{
		Short:     `List all your crons`,
		Long:      `List all your crons.`,
		Namespace: "container",
		Resource:  "cron",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.ListCronsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the crons`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "container-id",
				Short:      `UUID of the container invoked by the cron`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.ListCronsRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListCrons(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Crons, nil

		},
	}
}

func containerCronGet() *core.Command {
	return &core.Command{
		Short:     `Get a cron`,
		Long:      `Get the cron associated with the specified ID.`,
		Namespace: "container",
		Resource:  "cron",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.GetCronRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cron-id",
				Short:      `UUID of the cron to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.GetCronRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.GetCron(request)

		},
	}
}

func containerCronCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new cron`,
		Long:      `Create a new cron.`,
		Namespace: "container",
		Resource:  "cron",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.CreateCronRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "container-id",
				Short:      `UUID of the container to invoke by the cron`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "schedule",
				Short:      `UNIX cron shedule`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "args",
				Short:      `Arguments to pass with the cron`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the cron to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.CreateCronRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.CreateCron(request)

		},
	}
}

func containerCronUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing cron`,
		Long:      `Update the cron associated with the specified ID.`,
		Namespace: "container",
		Resource:  "cron",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.UpdateCronRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cron-id",
				Short:      `UUID of the cron to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "container-id",
				Short:      `UUID of the container invoked by the cron`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "schedule",
				Short:      `UNIX cron schedule`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "args",
				Short:      `Arguments to pass with the cron`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the cron`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.UpdateCronRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.UpdateCron(request)

		},
	}
}

func containerCronDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an existing cron`,
		Long:      `Delete the cron associated with the specified ID.`,
		Namespace: "container",
		Resource:  "cron",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.DeleteCronRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cron-id",
				Short:      `UUID of the cron to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.DeleteCronRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.DeleteCron(request)

		},
	}
}

func containerDomainList() *core.Command {
	return &core.Command{
		Short:     `List all domain name bindings`,
		Long:      `List all domain name bindings in a specified region.`,
		Namespace: "container",
		Resource:  "domain",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.ListDomainsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the domains`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "hostname_asc", "hostname_desc"},
			},
			{
				Name:       "container-id",
				Short:      `UUID of the container the domain belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.ListDomainsRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListDomains(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Domains, nil

		},
	}
}

func containerDomainGet() *core.Command {
	return &core.Command{
		Short:     `Get a domain name binding`,
		Long:      `Get a domain name binding for the container with the specified ID.`,
		Namespace: "container",
		Resource:  "domain",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.GetDomainRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "domain-id",
				Short:      `UUID of the domain to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.GetDomainRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.GetDomain(request)

		},
	}
}

func containerDomainCreate() *core.Command {
	return &core.Command{
		Short:     `Create a domain name binding`,
		Long:      `Create a domain name binding for the container with the specified ID.`,
		Namespace: "container",
		Resource:  "domain",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.CreateDomainRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hostname",
				Short:      `Domain to assign`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "container-id",
				Short:      `UUID of the container to assign the domain to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.CreateDomainRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.CreateDomain(request)

		},
	}
}

func containerDomainDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a domain name binding`,
		Long:      `Delete the domain name binding with the specific ID.`,
		Namespace: "container",
		Resource:  "domain",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.DeleteDomainRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "domain-id",
				Short:      `UUID of the domain to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.DeleteDomainRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.DeleteDomain(request)

		},
	}
}

func containerTokenCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new revocable token`,
		Long:      `Create a new revocable token.`,
		Namespace: "container",
		Resource:  "token",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.CreateTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "container-id",
				Short:      `UUID of the container to create the token for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace to create the token for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description of the token`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "expires-at",
				Short:      `Expiry date of the token`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.CreateTokenRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.CreateToken(request)

		},
	}
}

func containerTokenGet() *core.Command {
	return &core.Command{
		Short:     `Get a token`,
		Long:      `Get a token with a specified ID.`,
		Namespace: "container",
		Resource:  "token",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.GetTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "token-id",
				Short:      `UUID of the token to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.GetTokenRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.GetToken(request)

		},
	}
}

func containerTokenList() *core.Command {
	return &core.Command{
		Short:     `List all tokens`,
		Long:      `List all tokens belonging to a specified Organization or Project.`,
		Namespace: "container",
		Resource:  "token",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.ListTokensRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the tokens`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "container-id",
				Short:      `UUID of the container the token belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace the token belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.ListTokensRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListTokens(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Tokens, nil

		},
	}
}

func containerTokenDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a token`,
		Long:      `Delete a token with a specified ID.`,
		Namespace: "container",
		Resource:  "token",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.DeleteTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "token-id",
				Short:      `UUID of the token to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.DeleteTokenRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.DeleteToken(request)

		},
	}
}

func containerTriggerCreate() *core.Command {
	return &core.Command{
		Short:     `Create a trigger`,
		Long:      `Create a new trigger for a specified container.`,
		Namespace: "container",
		Resource:  "trigger",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.CreateTriggerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the trigger`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "container-id",
				Short:      `ID of the container to trigger`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description of the trigger`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scw-sqs-config.queue",
				Short:      `Name of the SQS queue the trigger should listen to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scw-sqs-config.mnq-project-id",
				Short:      `ID of the Messaging and Queuing project`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scw-sqs-config.mnq-region",
				Short:      `Region in which the Messaging and Queuing project is activated.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scw-nats-config.subject",
				Short:      `Name of the NATS subject the trigger should listen to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scw-nats-config.mnq-nats-account-id",
				Short:      `ID of the Messaging and Queuing NATS account`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scw-nats-config.mnq-project-id",
				Short:      `ID of the Messaging and Queuing project`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scw-nats-config.mnq-region",
				Short:      `Region in which the Messaging and Queuing project is activated.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.CreateTriggerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.CreateTrigger(request)

		},
	}
}

func containerTriggerGet() *core.Command {
	return &core.Command{
		Short:     `Get a trigger`,
		Long:      `Get a trigger with a specified ID.`,
		Namespace: "container",
		Resource:  "trigger",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.GetTriggerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "trigger-id",
				Short:      `ID of the trigger to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.GetTriggerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.GetTrigger(request)

		},
	}
}

func containerTriggerList() *core.Command {
	return &core.Command{
		Short:     `List all triggers`,
		Long:      `List all triggers belonging to a specified Organization or Project.`,
		Namespace: "container",
		Resource:  "trigger",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.ListTriggersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order in which to return results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "container-id",
				Short:      `ID of the container the triggers belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "namespace-id",
				Short:      `ID of the namespace the triggers belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.ListTriggersRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListTriggers(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Triggers, nil

		},
	}
}

func containerTriggerUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a trigger`,
		Long:      `Update a trigger with a specified ID.`,
		Namespace: "container",
		Resource:  "trigger",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.UpdateTriggerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "trigger-id",
				Short:      `ID of the trigger to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the trigger`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description of the trigger`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.UpdateTriggerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.UpdateTrigger(request)

		},
	}
}

func containerTriggerDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a trigger`,
		Long:      `Delete a trigger with a specified ID.`,
		Namespace: "container",
		Resource:  "trigger",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(container.DeleteTriggerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "trigger-id",
				Short:      `ID of the trigger to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*container.DeleteTriggerRequest)

			client := core.ExtractClient(ctx)
			api := container.NewAPI(client)
			return api.DeleteTrigger(request)

		},
	}
}

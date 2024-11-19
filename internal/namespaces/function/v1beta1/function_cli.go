// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package function

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		functionRoot(),
		functionNamespace(),
		functionFunction(),
		functionCron(),
		functionRuntime(),
		functionDomain(),
		functionToken(),
		functionTrigger(),
		functionNamespaceList(),
		functionNamespaceGet(),
		functionNamespaceCreate(),
		functionNamespaceUpdate(),
		functionNamespaceDelete(),
		functionFunctionList(),
		functionFunctionGet(),
		functionFunctionCreate(),
		functionFunctionUpdate(),
		functionFunctionDelete(),
		functionFunctionDeploy(),
		functionRuntimeList(),
		functionFunctionGetUploadURL(),
		functionFunctionGetDownloadURL(),
		functionCronList(),
		functionCronGet(),
		functionCronCreate(),
		functionCronUpdate(),
		functionCronDelete(),
		functionDomainList(),
		functionDomainGet(),
		functionDomainCreate(),
		functionDomainDelete(),
		functionTokenCreate(),
		functionTokenGet(),
		functionTokenList(),
		functionTokenDelete(),
		functionTriggerCreate(),
		functionTriggerGet(),
		functionTriggerList(),
		functionTriggerUpdate(),
		functionTriggerDelete(),
	)
}
func functionRoot() *core.Command {
	return &core.Command{
		Short:     `Function as a Service API`,
		Long:      `Function as a Service API.`,
		Namespace: "function",
	}
}

func functionNamespace() *core.Command {
	return &core.Command{
		Short:     `Function namespace management commands`,
		Long:      `Function namespace management commands.`,
		Namespace: "function",
		Resource:  "namespace",
	}
}

func functionFunction() *core.Command {
	return &core.Command{
		Short:     `Function management commands`,
		Long:      `Function management commands.`,
		Namespace: "function",
		Resource:  "function",
	}
}

func functionCron() *core.Command {
	return &core.Command{
		Short:     `Cron management commands`,
		Long:      `Cron management commands.`,
		Namespace: "function",
		Resource:  "cron",
	}
}

func functionRuntime() *core.Command {
	return &core.Command{
		Short:     `Runtime management commands`,
		Long:      `Runtime management commands.`,
		Namespace: "function",
		Resource:  "runtime",
	}
}

func functionDomain() *core.Command {
	return &core.Command{
		Short:     `Domain management commands`,
		Long:      `Domain management commands.`,
		Namespace: "function",
		Resource:  "domain",
	}
}

func functionToken() *core.Command {
	return &core.Command{
		Short:     `Token management commands`,
		Long:      `Token management commands.`,
		Namespace: "function",
		Resource:  "token",
	}
}

func functionTrigger() *core.Command {
	return &core.Command{
		Short:     `Trigger management commands`,
		Long:      `Trigger management commands.`,
		Namespace: "function",
		Resource:  "trigger",
	}
}

func functionNamespaceList() *core.Command {
	return &core.Command{
		Short:     `List all your namespaces`,
		Long:      `List all existing namespaces in the specified region.`,
		Namespace: "function",
		Resource:  "namespace",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.ListNamespacesRequest{}),
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
				Short:      `Name of the namespace`,
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
			request := args.(*function.ListNamespacesRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
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
				FieldName: "RegistryNamespaceID",
			},
			{
				FieldName: "RegistryEndpoint",
			},
			{
				FieldName: "EnvironmentVariables",
			},
			{
				FieldName: "ErrorMessage",
			},
			{
				FieldName: "Description",
			},
			{
				FieldName: "Region",
			},
			{
				FieldName: "ProjectID",
			},
			{
				FieldName: "OrganizationID",
			},
		}},
	}
}

func functionNamespaceGet() *core.Command {
	return &core.Command{
		Short:     `Get a namespace`,
		Long:      `Get the namespace associated with the specified ID.`,
		Namespace: "function",
		Resource:  "namespace",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.GetNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.GetNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.GetNamespace(request)

		},
	}
}

func functionNamespaceCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new namespace`,
		Long:      `Create a new namespace in a specified Organization or Project.`,
		Namespace: "function",
		Resource:  "namespace",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.CreateNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("ns"),
			},
			{
				Name:       "environment-variables.{key}",
				Short:      `Environment variables of the namespace`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "description",
				Short:      `Description of the namespace`,
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
				Short:      `[ALPHA] Tags of the Serverless Function Namespace`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.CreateNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.CreateNamespace(request)

		},
	}
}

func functionNamespaceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing namespace`,
		Long:      `Update the namespace associated with the specified ID.`,
		Namespace: "function",
		Resource:  "namespace",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.UpdateNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespapce`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "environment-variables.{key}",
				Short:      `Environment variables of the namespace`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description of the namespace`,
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
				Short:      `[ALPHA] Tags of the Serverless Function Namespace`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.UpdateNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.UpdateNamespace(request)

		},
	}
}

func functionNamespaceDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an existing namespace`,
		Long:      `Delete the namespace associated with the specified ID.`,
		Namespace: "function",
		Resource:  "namespace",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.DeleteNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.DeleteNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.DeleteNamespace(request)

		},
	}
}

func functionFunctionList() *core.Command {
	return &core.Command{
		Short:     `List all your functions`,
		Long:      `List all your functions.`,
		Namespace: "function",
		Resource:  "function",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.ListFunctionsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the functions`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace the function belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Name of the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `UUID of the Project the function belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `UUID of the Organziation the function belongs to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.ListFunctionsRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListFunctions(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Functions, nil

		},
	}
}

func functionFunctionGet() *core.Command {
	return &core.Command{
		Short:     `Get a function`,
		Long:      `Get the function associated with the specified ID.`,
		Namespace: "function",
		Resource:  "function",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.GetFunctionRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "function-id",
				Short:      `UUID of the function`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.GetFunctionRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.GetFunction(request)

		},
	}
}

func functionFunctionCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new function`,
		Long:      `Create a new function in the specified region for a specified Organization or Project.`,
		Namespace: "function",
		Resource:  "function",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.CreateFunctionRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the function to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("fn"),
			},
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace the function will be created in`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "environment-variables.{key}",
				Short:      `Environment variables of the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "min-scale",
				Short:      `Minumum number of instances to scale the function to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-scale",
				Short:      `Maximum number of instances to scale the function to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "runtime",
				Short:      `Runtime to use with the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_runtime", "golang", "python", "python3", "node8", "node10", "node14", "node16", "node17", "python37", "python38", "python39", "python310", "go113", "go117", "go118", "node18", "rust165", "go119", "python311", "php82", "node19", "go120", "node20", "go121", "node22", "python312", "php83", "go122", "rust179"},
			},
			{
				Name:       "memory-limit",
				Short:      `Memory limit of the function in MB`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "timeout",
				Short:      `Request processing time limit for the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "handler",
				Short:      `Handler to use with the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "privacy",
				Short:      `Privacy setting of the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_privacy", "public", "private"},
			},
			{
				Name:       "description",
				Short:      `Description of the function`,
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
				Short:      `Execution environment of the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_sandbox", "v1", "v2"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.CreateFunctionRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.CreateFunction(request)

		},
	}
}

func functionFunctionUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing function`,
		Long:      `Update the function associated with the specified ID.`,
		Namespace: "function",
		Resource:  "function",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.UpdateFunctionRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "function-id",
				Short:      `UUID of the function to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "environment-variables.{key}",
				Short:      `Environment variables of the function to update`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "min-scale",
				Short:      `Minumum number of instances to scale the function to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "max-scale",
				Short:      `Maximum number of instances to scale the function to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "runtime",
				Short:      `Runtime to use with the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_runtime", "golang", "python", "python3", "node8", "node10", "node14", "node16", "node17", "python37", "python38", "python39", "python310", "go113", "go117", "go118", "node18", "rust165", "go119", "python311", "php82", "node19", "go120", "node20", "go121", "node22", "python312", "php83", "go122", "rust179"},
			},
			{
				Name:       "memory-limit",
				Short:      `Memory limit of the function in MB`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "timeout",
				Short:      `Processing time limit for the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "redeploy",
				Short:      `Redeploy failed function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "handler",
				Short:      `Handler to use with the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "privacy",
				Short:      `Privacy setting of the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_privacy", "public", "private"},
			},
			{
				Name:       "description",
				Short:      `Description of the function`,
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
				Short:      `Execution environment of the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_sandbox", "v1", "v2"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.UpdateFunctionRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.UpdateFunction(request)

		},
	}
}

func functionFunctionDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a function`,
		Long:      `Delete the function associated with the specified ID.`,
		Namespace: "function",
		Resource:  "function",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.DeleteFunctionRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "function-id",
				Short:      `UUID of the function to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.DeleteFunctionRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.DeleteFunction(request)

		},
	}
}

func functionFunctionDeploy() *core.Command {
	return &core.Command{
		Short:     `Deploy a function`,
		Long:      `Deploy a function associated with the specified ID.`,
		Namespace: "function",
		Resource:  "function",
		Verb:      "deploy",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.DeployFunctionRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "function-id",
				Short:      `UUID of the function to deploy`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.DeployFunctionRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.DeployFunction(request)

		},
	}
}

func functionRuntimeList() *core.Command {
	return &core.Command{
		Short:     `List function runtimes`,
		Long:      `List available function runtimes.`,
		Namespace: "function",
		Resource:  "runtime",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.ListFunctionRuntimesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.ListFunctionRuntimesRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.ListFunctionRuntimes(request)

		},
	}
}

func functionFunctionGetUploadURL() *core.Command {
	return &core.Command{
		Short:     `Get an upload URL of a function`,
		Long:      `Get an upload URL of a function associated with the specified ID.`,
		Namespace: "function",
		Resource:  "function",
		Verb:      "get-upload-url",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.GetFunctionUploadURLRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "function-id",
				Short:      `UUID of the function to get the upload URL for`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "content-length",
				Short:      `Size of the archive to upload in bytes`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.GetFunctionUploadURLRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.GetFunctionUploadURL(request)

		},
	}
}

func functionFunctionGetDownloadURL() *core.Command {
	return &core.Command{
		Short:     `Get a download URL of a function`,
		Long:      `Get a download URL for a function associated with the specified ID.`,
		Namespace: "function",
		Resource:  "function",
		Verb:      "get-download-url",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.GetFunctionDownloadURLRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "function-id",
				Short:      `UUID of the function to get the the download URL for`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.GetFunctionDownloadURLRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.GetFunctionDownloadURL(request)

		},
	}
}

func functionCronList() *core.Command {
	return &core.Command{
		Short:     `List all crons`,
		Long:      `List all the cronjobs in a specified region.`,
		Namespace: "function",
		Resource:  "cron",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.ListCronsRequest{}),
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
				Name:       "function-id",
				Short:      `UUID of the function`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.ListCronsRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
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

func functionCronGet() *core.Command {
	return &core.Command{
		Short:     `Get a cron`,
		Long:      `Get the cron associated with the specified ID.`,
		Namespace: "function",
		Resource:  "cron",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.GetCronRequest{}),
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
			request := args.(*function.GetCronRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.GetCron(request)

		},
	}
}

func functionCronCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new cron`,
		Long:      `Create a new cronjob for a function with the specified ID.`,
		Namespace: "function",
		Resource:  "cron",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.CreateCronRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "function-id",
				Short:      `UUID of the function to use the cron with`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "schedule",
				Short:      `Schedule of the cron in UNIX cron format`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "args",
				Short:      `Arguments to use with the cron`,
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
			request := args.(*function.CreateCronRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.CreateCron(request)

		},
	}
}

func functionCronUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing cron`,
		Long:      `Update the cron associated with the specified ID.`,
		Namespace: "function",
		Resource:  "cron",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.UpdateCronRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cron-id",
				Short:      `UUID of the cron to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "function-id",
				Short:      `UUID of the function to use the cron with`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "schedule",
				Short:      `Schedule of the cron in UNIX cron format`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "args",
				Short:      `Arguments to use with the cron`,
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
			request := args.(*function.UpdateCronRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.UpdateCron(request)

		},
	}
}

func functionCronDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an existing cron`,
		Long:      `Delete the cron associated with the specified ID.`,
		Namespace: "function",
		Resource:  "cron",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.DeleteCronRequest{}),
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
			request := args.(*function.DeleteCronRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.DeleteCron(request)

		},
	}
}

func functionDomainList() *core.Command {
	return &core.Command{
		Short:     `List all domain name bindings`,
		Long:      `List all domain name bindings in a specified region.`,
		Namespace: "function",
		Resource:  "domain",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.ListDomainsRequest{}),
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
				Name:       "function-id",
				Short:      `UUID of the function the domain is assoicated with`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.ListDomainsRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
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

func functionDomainGet() *core.Command {
	return &core.Command{
		Short:     `Get a domain name binding`,
		Long:      `Get a domain name binding for the function with the specified ID.`,
		Namespace: "function",
		Resource:  "domain",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.GetDomainRequest{}),
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
			request := args.(*function.GetDomainRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.GetDomain(request)

		},
	}
}

func functionDomainCreate() *core.Command {
	return &core.Command{
		Short:     `Create a domain name binding`,
		Long:      `Create a domain name binding for the function with the specified ID.`,
		Namespace: "function",
		Resource:  "domain",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.CreateDomainRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hostname",
				Short:      `Hostame to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "function-id",
				Short:      `UUID of the function to associate the domain with`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.CreateDomainRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.CreateDomain(request)

		},
	}
}

func functionDomainDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a domain name binding`,
		Long:      `Delete a domain name binding for the function with the specified ID.`,
		Namespace: "function",
		Resource:  "domain",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.DeleteDomainRequest{}),
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
			request := args.(*function.DeleteDomainRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.DeleteDomain(request)

		},
	}
}

func functionTokenCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new revocable token`,
		Long:      `Create a new revocable token.`,
		Namespace: "function",
		Resource:  "token",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.CreateTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "function-id",
				Short:      `UUID of the function to associate the token with`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace to associate the token with`,
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
				Short:      `Date on which the token expires`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.CreateTokenRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.CreateToken(request)

		},
	}
}

func functionTokenGet() *core.Command {
	return &core.Command{
		Short:     `Get a token`,
		Long:      `Get a token.`,
		Namespace: "function",
		Resource:  "token",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.GetTokenRequest{}),
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
			request := args.(*function.GetTokenRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.GetToken(request)

		},
	}
}

func functionTokenList() *core.Command {
	return &core.Command{
		Short:     `List all tokens`,
		Long:      `List all tokens.`,
		Namespace: "function",
		Resource:  "token",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.ListTokensRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order for the tokens`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "function-id",
				Short:      `UUID of the function the token is assoicated with`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "namespace-id",
				Short:      `UUID of the namespace the token is associated with`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*function.ListTokensRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
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

func functionTokenDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a token`,
		Long:      `Delete a token.`,
		Namespace: "function",
		Resource:  "token",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.DeleteTokenRequest{}),
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
			request := args.(*function.DeleteTokenRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.DeleteToken(request)

		},
	}
}

func functionTriggerCreate() *core.Command {
	return &core.Command{
		Short:     `Create a trigger`,
		Long:      `Create a new trigger for a specified function.`,
		Namespace: "function",
		Resource:  "trigger",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.CreateTriggerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the trigger`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "function-id",
				Short:      `ID of the function to trigger`,
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
			request := args.(*function.CreateTriggerRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.CreateTrigger(request)

		},
	}
}

func functionTriggerGet() *core.Command {
	return &core.Command{
		Short:     `Get a trigger`,
		Long:      `Get a trigger with a specified ID.`,
		Namespace: "function",
		Resource:  "trigger",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.GetTriggerRequest{}),
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
			request := args.(*function.GetTriggerRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.GetTrigger(request)

		},
	}
}

func functionTriggerList() *core.Command {
	return &core.Command{
		Short:     `List all triggers`,
		Long:      `List all triggers belonging to a specified Organization or Project.`,
		Namespace: "function",
		Resource:  "trigger",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.ListTriggersRequest{}),
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
				Name:       "function-id",
				Short:      `ID of the function the triggers belongs to`,
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
			request := args.(*function.ListTriggersRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
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

func functionTriggerUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a trigger`,
		Long:      `Update a trigger with a specified ID.`,
		Namespace: "function",
		Resource:  "trigger",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.UpdateTriggerRequest{}),
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
			request := args.(*function.UpdateTriggerRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.UpdateTrigger(request)

		},
	}
}

func functionTriggerDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a trigger`,
		Long:      `Delete a trigger with a specified ID.`,
		Namespace: "function",
		Resource:  "trigger",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(function.DeleteTriggerRequest{}),
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
			request := args.(*function.DeleteTriggerRequest)

			client := core.ExtractClient(ctx)
			api := function.NewAPI(client)
			return api.DeleteTrigger(request)

		},
	}
}

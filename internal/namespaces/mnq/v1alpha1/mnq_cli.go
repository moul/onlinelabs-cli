// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package mnq

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/mnq/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		mnqRoot(),
		mnqNamespace(),
		mnqCredential(),
		mnqNamespaceList(),
		mnqNamespaceCreate(),
		mnqNamespaceUpdate(),
		mnqNamespaceGet(),
		mnqNamespaceDelete(),
		mnqCredentialCreate(),
		mnqCredentialDelete(),
		mnqCredentialList(),
		mnqCredentialUpdate(),
		mnqCredentialGet(),
	)
}
func mnqRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage Messaging or Queueing brokers`,
		Long:      ``,
		Namespace: "mnq",
	}
}

func mnqNamespace() *core.Command {
	return &core.Command{
		Short:     `MnQ Namespace commands`,
		Long:      `MnQ Namespace commands.`,
		Namespace: "mnq",
		Resource:  "namespace",
	}
}

func mnqCredential() *core.Command {
	return &core.Command{
		Short:     `MnQ Credentials commands`,
		Long:      `MnQ Credentials commands.`,
		Namespace: "mnq",
		Resource:  "credential",
	}
}

func mnqNamespaceList() *core.Command {
	return &core.Command{
		Short:     `List namespaces`,
		Long:      `List namespaces.`,
		Namespace: "mnq",
		Resource:  "namespace",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.ListNamespacesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "project-id",
				Short:      `Will list only the Namespaces contained into the specified project`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Field used for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "id_asc", "id_desc", "name_asc", "name_desc", "project_id_asc", "project_id_desc"},
			},
			{
				Name:       "organization-id",
				Short:      `Will list only the Namespaces owned by the specified organization`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.ListNamespacesRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
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

func mnqNamespaceCreate() *core.Command {
	return &core.Command{
		Short:     `Create a namespace`,
		Long:      `Create a namespace.`,
		Namespace: "mnq",
		Resource:  "namespace",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.CreateNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Namespace name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("mnq"),
			},
			{
				Name:       "protocol",
				Short:      `Namespace protocol`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "nats", "sqs_sns", "amqp"},
			},
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.CreateNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			return api.CreateNamespace(request)

		},
	}
}

func mnqNamespaceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update the name of a namespace`,
		Long:      `Update the name of a namespace.`,
		Namespace: "mnq",
		Resource:  "namespace",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.UpdateNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `ID of the Namespace to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Namespace name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.UpdateNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			return api.UpdateNamespace(request)

		},
	}
}

func mnqNamespaceGet() *core.Command {
	return &core.Command{
		Short:     `Get a namespace`,
		Long:      `Get a namespace.`,
		Namespace: "mnq",
		Resource:  "namespace",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.GetNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `ID of the Namespace to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.GetNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			return api.GetNamespace(request)

		},
	}
}

func mnqNamespaceDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a namespace`,
		Long:      `Delete a namespace.`,
		Namespace: "mnq",
		Resource:  "namespace",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.DeleteNamespaceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `ID of the Namespace to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.DeleteNamespaceRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			e = api.DeleteNamespace(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "namespace",
				Verb:     "delete",
			}, nil
		},
	}
}

func mnqCredentialCreate() *core.Command {
	return &core.Command{
		Short:     `Create a set of credentials`,
		Long:      `Create a set of credentials for a specific namespace.`,
		Namespace: "mnq",
		Resource:  "credential",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.CreateCredentialRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `Namespace containing the Credential`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Credential name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("mnq"),
			},
			{
				Name:       "permissions.can-publish",
				Short:      `Defines if user can publish messages to the service`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "permissions.can-receive",
				Short:      `Defines if user can receive messages from the service`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "permissions.can-manage",
				Short:      `Defines if user can manage the associated resource(s)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.CreateCredentialRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			return api.CreateCredential(request)

		},
	}
}

func mnqCredentialDelete() *core.Command {
	return &core.Command{
		Short:     `Delete credentials`,
		Long:      `Delete credentials.`,
		Namespace: "mnq",
		Resource:  "credential",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.DeleteCredentialRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "credential-id",
				Short:      `ID of the Credential to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.DeleteCredentialRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			e = api.DeleteCredential(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "credential",
				Verb:     "delete",
			}, nil
		},
	}
}

func mnqCredentialList() *core.Command {
	return &core.Command{
		Short:     `List credentials`,
		Long:      `List credentials.`,
		Namespace: "mnq",
		Resource:  "credential",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.ListCredentialsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "namespace-id",
				Short:      `Namespace containing the Credential`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "order-by",
				Short:      `Field used for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"id_asc", "id_desc", "name_asc", "name_desc"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.ListCredentialsRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListCredentials(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Credentials, nil

		},
	}
}

func mnqCredentialUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a set of credentials`,
		Long:      `Update a set of credentials.`,
		Namespace: "mnq",
		Resource:  "credential",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.UpdateCredentialRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "credential-id",
				Short:      `ID of the Credential to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Credential name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "permissions.can-publish",
				Short:      `Defines if user can publish messages to the service`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "permissions.can-receive",
				Short:      `Defines if user can receive messages from the service`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "permissions.can-manage",
				Short:      `Defines if user can manage the associated resource(s)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.UpdateCredentialRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			return api.UpdateCredential(request)

		},
	}
}

func mnqCredentialGet() *core.Command {
	return &core.Command{
		Short:     `Get a set of credentials`,
		Long:      `Get a set of credentials.`,
		Namespace: "mnq",
		Resource:  "credential",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(mnq.GetCredentialRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "credential-id",
				Short:      `ID of the Credential to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*mnq.GetCredentialRequest)

			client := core.ExtractClient(ctx)
			api := mnq.NewAPI(client)
			return api.GetCredential(request)

		},
	}
}

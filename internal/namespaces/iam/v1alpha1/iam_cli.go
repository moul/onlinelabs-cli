// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package iam

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/iam/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		iamRoot(),
		iamSSHKey(),
		iamGroup(),
		iamAPIKey(),
		iamUser(),
		iamApplication(),
		iamPolicy(),
		iamRule(),
		iamPermissionSet(),
		iamJwt(),
		iamLog(),
		iamSSHKeyList(),
		iamSSHKeyCreate(),
		iamSSHKeyGet(),
		iamSSHKeyUpdate(),
		iamSSHKeyDelete(),
		iamUserList(),
		iamUserGet(),
		iamUserUpdate(),
		iamUserDelete(),
		iamUserCreate(),
		iamUserUpdatePassword(),
		iamApplicationList(),
		iamApplicationCreate(),
		iamApplicationGet(),
		iamApplicationUpdate(),
		iamApplicationDelete(),
		iamGroupList(),
		iamGroupCreate(),
		iamGroupGet(),
		iamGroupUpdate(),
		iamGroupSetMembers(),
		iamGroupAddMember(),
		iamGroupAddMembers(),
		iamGroupRemoveMember(),
		iamGroupDelete(),
		iamPolicyList(),
		iamPolicyCreate(),
		iamPolicyGet(),
		iamPolicyUpdate(),
		iamPolicyDelete(),
		iamPolicyClone(),
		iamRuleUpdate(),
		iamRuleList(),
		iamPermissionSetList(),
		iamAPIKeyList(),
		iamAPIKeyCreate(),
		iamAPIKeyGet(),
		iamAPIKeyUpdate(),
		iamAPIKeyDelete(),
		iamJwtList(),
		iamJwtGet(),
		iamJwtDelete(),
		iamLogList(),
		iamLogGet(),
	)
}
func iamRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage Identity and Access Management (IAM) across your Scaleway Organizations, Projects and resources`,
		Long:      `This API allows you to manage Identity and Access Management (IAM) across your Scaleway Organizations, Projects and resources.`,
		Namespace: "iam",
	}
}

func iamSSHKey() *core.Command {
	return &core.Command{
		Short:     `SSH keys management commands`,
		Long:      `SSH keys management commands.`,
		Namespace: "iam",
		Resource:  "ssh-key",
	}
}

func iamGroup() *core.Command {
	return &core.Command{
		Short:     `Groups management commands`,
		Long:      `Groups management commands.`,
		Namespace: "iam",
		Resource:  "group",
	}
}

func iamAPIKey() *core.Command {
	return &core.Command{
		Short:     `API keys management commands`,
		Long:      `API keys management commands.`,
		Namespace: "iam",
		Resource:  "api-key",
	}
}

func iamUser() *core.Command {
	return &core.Command{
		Short:     `Users management commands`,
		Long:      `Users management commands.`,
		Namespace: "iam",
		Resource:  "user",
	}
}

func iamApplication() *core.Command {
	return &core.Command{
		Short:     `Applications management commands`,
		Long:      `Applications management commands.`,
		Namespace: "iam",
		Resource:  "application",
	}
}

func iamPolicy() *core.Command {
	return &core.Command{
		Short:     `Policies management commands`,
		Long:      `Policies management commands.`,
		Namespace: "iam",
		Resource:  "policy",
	}
}

func iamRule() *core.Command {
	return &core.Command{
		Short:     `Rules management commands`,
		Long:      `Rules management commands.`,
		Namespace: "iam",
		Resource:  "rule",
	}
}

func iamPermissionSet() *core.Command {
	return &core.Command{
		Short:     `Permission sets management commands`,
		Long:      `Permission sets management commands.`,
		Namespace: "iam",
		Resource:  "permission-set",
	}
}

func iamJwt() *core.Command {
	return &core.Command{
		Short:     `JWTs management commands`,
		Long:      `JWTs management commands.`,
		Namespace: "iam",
		Resource:  "jwt",
	}
}

func iamLog() *core.Command {
	return &core.Command{
		Short:     `Log management commands`,
		Long:      `Log management commands.`,
		Namespace: "iam",
		Resource:  "log",
	}
}

func iamSSHKeyList() *core.Command {
	return &core.Command{
		Short:     `List SSH keys`,
		Long:      `List SSH keys. By default, the SSH keys listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field. You can define additional parameters for your query such as ` + "`" + `organization_id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `project_id` + "`" + ` and ` + "`" + `disabled` + "`" + `.`,
		Namespace: "iam",
		Resource:  "ssh-key",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListSSHKeysRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order of the SSH keys`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "name",
				Short:      `Name of group to find`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "project-id",
				Short:      `Filter by Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "disabled",
				Short:      `Defines whether to include disabled SSH keys or not`,
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
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListSSHKeysRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListSSHKeys(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.SSHKeys, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "CreatedAt",
			},
			{
				FieldName: "Fingerprint",
			},
			{
				FieldName: "ProjectID",
			},
			{
				FieldName: "Disabled",
			},
		}},
	}
}

func iamSSHKeyCreate() *core.Command {
	return &core.Command{
		Short:     `Create an SSH key`,
		Long:      `Add a new SSH key to a Scaleway Project. You must specify the ` + "`" + `name` + "`" + `, ` + "`" + `public_key` + "`" + ` and ` + "`" + `project_id` + "`" + `.`,
		Namespace: "iam",
		Resource:  "ssh-key",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.CreateSSHKeyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the SSH key. Max length is 1000`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("key"),
			},
			{
				Name:       "public-key",
				Short:      `SSH public key. Currently only the ssh-rsa, ssh-dss (DSA), ssh-ed25519 and ecdsa keys with NIST curves are supported. Max length is 65000`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.CreateSSHKeyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.CreateSSHKey(request)

		},
		Examples: []*core.Example{
			{
				Short: "Add a given ssh key",
				Raw:   `scw iam ssh-key create name=foobar public-key="$(cat <path/to/your/public/key>)"`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam ssh-key list",
				Short:   "List all SSH keys",
			},
			{
				Command: "scw iam ssh-key delete",
				Short:   "Delete an SSH key",
			},
		},
	}
}

func iamSSHKeyGet() *core.Command {
	return &core.Command{
		Short:     `Get an SSH key`,
		Long:      `Retrieve information about a given SSH key, specified by the ` + "`" + `ssh_key_id` + "`" + ` parameter. The SSH key's full details, including ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `public_key` + "`" + `, and ` + "`" + `project_id` + "`" + ` are returned in the response.`,
		Namespace: "iam",
		Resource:  "ssh-key",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.GetSSHKeyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ssh-key-id",
				Short:      `ID of the SSH key`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.GetSSHKeyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.GetSSHKey(request)

		},
	}
}

func iamSSHKeyUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an SSH key`,
		Long:      `Update the parameters of an SSH key, including ` + "`" + `name` + "`" + ` and ` + "`" + `disable` + "`" + `.`,
		Namespace: "iam",
		Resource:  "ssh-key",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.UpdateSSHKeyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ssh-key-id",
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Name of the SSH key. Max length is 1000`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "disabled",
				Short:      `Enable or disable the SSH key`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.UpdateSSHKeyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.UpdateSSHKey(request)

		},
	}
}

func iamSSHKeyDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an SSH key`,
		Long:      `Delete a given SSH key, specified by the ` + "`" + `ssh_key_id` + "`" + `. Deleting an SSH is permanent, and cannot be undone. Note that you might need to update any configurations that used the SSH key.`,
		Namespace: "iam",
		Resource:  "ssh-key",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.DeleteSSHKeyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "ssh-key-id",
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.DeleteSSHKeyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			e = api.DeleteSSHKey(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "ssh-key",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete a given SSH key",
				ArgsJSON: `{"ssh_key_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam ssh-key list",
				Short:   "List all SSH keys",
			},
			{
				Command: "scw iam ssh-key create",
				Short:   "Add a SSH key",
			},
		},
	}
}

func iamUserList() *core.Command {
	return &core.Command{
		Short:     `List users of an Organization`,
		Long:      `List the users of an Organization. By default, the users listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field. You must define the ` + "`" + `organization_id` + "`" + ` in the query path of your request. You can also define additional parameters for your query such as ` + "`" + `user_ids` + "`" + `.`,
		Namespace: "iam",
		Resource:  "user",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListUsersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "email_asc", "email_desc", "last_login_asc", "last_login_desc"},
			},
			{
				Name:       "user-ids.{index}",
				Short:      `Filter by list of IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "mfa",
				Short:      `Filter by MFA status`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tag",
				Short:      `Filter by tags containing a given string`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `ID of the Organization to filter`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListUsersRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListUsers(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Users, nil

		},
	}
}

func iamUserGet() *core.Command {
	return &core.Command{
		Short:     `Get a given user`,
		Long:      `Retrieve information about a user, specified by the ` + "`" + `user_id` + "`" + ` parameter. The user's full details, including ` + "`" + `id` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `organization_id` + "`" + `, ` + "`" + `status` + "`" + ` and ` + "`" + `mfa` + "`" + ` are returned in the response.`,
		Namespace: "iam",
		Resource:  "user",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.GetUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "user-id",
				Short:      `ID of the user to find`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.GetUserRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.GetUser(request)

		},
	}
}

func iamUserUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a user`,
		Long:      `Update the parameters of a user, including ` + "`" + `tags` + "`" + `.`,
		Namespace: "iam",
		Resource:  "user",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.UpdateUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "user-id",
				Short:      `ID of the user to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "tags.{index}",
				Short:      `New tags for the user (maximum of 10 tags)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.UpdateUserRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.UpdateUser(request)

		},
	}
}

func iamUserDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a guest user from an Organization`,
		Long:      `Remove a user from an Organization in which they are a guest. You must define the ` + "`" + `user_id` + "`" + ` in your request. Note that removing a user from an Organization automatically deletes their API keys, and any policies directly attached to them become orphaned.`,
		Namespace: "iam",
		Resource:  "user",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.DeleteUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "user-id",
				Short:      `ID of the user to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.DeleteUserRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
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

func iamUserCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new user`,
		Long:      `Create a new user. You must define the ` + "`" + `organization_id` + "`" + ` and the ` + "`" + `email` + "`" + ` in your request.`,
		Namespace: "iam",
		Resource:  "user",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.CreateUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "email",
				Short:      `Email of the user`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags associated with the user`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "member.email",
				Short:      `Email of the user to create`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "member.send-password-email",
				Short:      `Whether or not to send an email containing the member's password.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "member.send-welcome-email",
				Short:      `Whether or not to send a welcome email that includes onboarding information.`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "member.username",
				Short:      `The member's username`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "member.password",
				Short:      `The member's password`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.CreateUserRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.CreateUser(request)

		},
	}
}

func iamUserUpdatePassword() *core.Command {
	return &core.Command{
		Short:     `Update an user's password`,
		Long:      `Update an user's password.`,
		Namespace: "iam",
		Resource:  "user",
		Verb:      "update-password",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.UpdateUserPasswordRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "user-id",
				Short:      `ID of the user to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "password",
				Short:      `The new password`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "send-email",
				Short:      `Whether or not to send an email alerting the user their password has changed`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.UpdateUserPasswordRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.UpdateUserPassword(request)

		},
	}
}

func iamApplicationList() *core.Command {
	return &core.Command{
		Short:     `List applications of an Organization`,
		Long:      `List the applications of an Organization. By default, the applications listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field. You must define the ` + "`" + `organization_id` + "`" + ` in the query path of your request. You can also define additional parameters for your query such as ` + "`" + `application_ids` + "`" + `.`,
		Namespace: "iam",
		Resource:  "application",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListApplicationsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "name",
				Short:      `Name of the application to filter`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "editable",
				Short:      `Defines whether to filter out editable applications or not`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-ids.{index}",
				Short:      `Filter by list of IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tag",
				Short:      `Filter by tags containing a given string`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListApplicationsRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListApplications(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Applications, nil

		},
	}
}

func iamApplicationCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new application`,
		Long:      `Create a new application. You must define the ` + "`" + `name` + "`" + ` parameter in the request.`,
		Namespace: "iam",
		Resource:  "application",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.CreateApplicationRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the application to create (max length is 64 characters)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("app"),
			},
			{
				Name:       "description",
				Short:      `Description of the application (max length is 200 characters)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags associated with the application (maximum of 10 tags)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.CreateApplicationRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.CreateApplication(request)

		},
	}
}

func iamApplicationGet() *core.Command {
	return &core.Command{
		Short:     `Get a given application`,
		Long:      `Retrieve information about an application, specified by the ` + "`" + `application_id` + "`" + ` parameter. The application's full details, including ` + "`" + `id` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `organization_id` + "`" + `, ` + "`" + `status` + "`" + ` and ` + "`" + `two_factor_enabled` + "`" + ` are returned in the response.`,
		Namespace: "iam",
		Resource:  "application",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.GetApplicationRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "application-id",
				Short:      `ID of the application to find`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.GetApplicationRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.GetApplication(request)

		},
	}
}

func iamApplicationUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an application`,
		Long:      `Update the parameters of an application, including ` + "`" + `name` + "`" + ` and ` + "`" + `description` + "`" + `.`,
		Namespace: "iam",
		Resource:  "application",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.UpdateApplicationRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "application-id",
				Short:      `ID of the application to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `New name for the application (max length is 64 chars)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `New description for the application (max length is 200 chars)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `New tags for the application (maximum of 10 tags)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.UpdateApplicationRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.UpdateApplication(request)

		},
	}
}

func iamApplicationDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an application`,
		Long:      `Delete an application. Note that this action is irreversible and will automatically delete the application's API keys. Policies attached to users and applications via this group will no longer apply.`,
		Namespace: "iam",
		Resource:  "application",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.DeleteApplicationRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "application-id",
				Short:      `ID of the application to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.DeleteApplicationRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			e = api.DeleteApplication(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "application",
				Verb:     "delete",
			}, nil
		},
	}
}

func iamGroupList() *core.Command {
	return &core.Command{
		Short:     `List groups`,
		Long:      `List groups. By default, the groups listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field. You can define additional parameters to filter your query. Use ` + "`" + `user_ids` + "`" + ` or ` + "`" + `application_ids` + "`" + ` to list all groups certain users or applications belong to.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListGroupsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order of groups`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "name_asc", "name_desc"},
			},
			{
				Name:       "name",
				Short:      `Name of group to find`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-ids.{index}",
				Short:      `Filter by a list of application IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-ids.{index}",
				Short:      `Filter by a list of user IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "group-ids.{index}",
				Short:      `Filter by a list of group IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tag",
				Short:      `Filter by tags containing a given string`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListGroupsRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListGroups(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Groups, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "UserIDs",
			},
			{
				FieldName: "ApplicationIDs",
			},
		}},
	}
}

func iamGroupCreate() *core.Command {
	return &core.Command{
		Short:     `Create a group`,
		Long:      `Create a new group. You must define the ` + "`" + `name` + "`" + ` and ` + "`" + `organization_id` + "`" + ` parameters in the request.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.CreateGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the group to create (max length is 64 chars). MUST be unique inside an Organization`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("grp"),
			},
			{
				Name:       "description",
				Short:      `Description of the group to create (max length is 200 chars)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags associated with the group (maximum of 10 tags)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.CreateGroupRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.CreateGroup(request)

		},
		Examples: []*core.Example{
			{
				Short: "Create a group",
				Raw:   `scw iam group create name=foobar`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam group add-member",
				Short:   "Add a group member",
			},
			{
				Command: "scw iam group delete",
				Short:   "Delete a group",
			},
			{
				Command: "scw iam policy create",
				Short:   "Create a policy for a group",
			},
		},
	}
}

func iamGroupGet() *core.Command {
	return &core.Command{
		Short:     `Get a group`,
		Long:      `Retrive information about a given group, specified by the ` + "`" + `group_id` + "`" + ` parameter. The group's full details, including ` + "`" + `user_ids` + "`" + ` and ` + "`" + `application_ids` + "`" + ` are returned in the response.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.GetGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "group-id",
				Short:      `ID of the group`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.GetGroupRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.GetGroup(request)

		},
	}
}

func iamGroupUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a group`,
		Long:      `Update the parameters of group, including ` + "`" + `name` + "`" + ` and ` + "`" + `description` + "`" + `.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.UpdateGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "group-id",
				Short:      `ID of the group to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `New name for the group (max length is 64 chars). MUST be unique inside an Organization`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `New description for the group (max length is 200 chars)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `New tags for the group (maximum of 10 tags)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.UpdateGroupRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.UpdateGroup(request)

		},
	}
}

func iamGroupSetMembers() *core.Command {
	return &core.Command{
		Short:     `Overwrite users and applications of a group`,
		Long:      `Overwrite users and applications configuration in a group. Any information that you add using this command will overwrite the previous configuration.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "set-members",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.SetGroupMembersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "group-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-ids.{index}",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-ids.{index}",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.SetGroupMembersRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.SetGroupMembers(request)

		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam group remove-member",
				Short:   "Remove a group member",
			},
			{
				Command: "scw iam group create",
				Short:   "Create a group",
			},
		},
	}
}

func iamGroupAddMember() *core.Command {
	return &core.Command{
		Short:     `Add a user or an application to a group`,
		Long:      `Add a user or an application to a group. You can specify a ` + "`" + `user_id` + "`" + ` and and ` + "`" + `application_id` + "`" + ` in the body of your request. Note that you can only add one of each per request.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "add-member",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.AddGroupMemberRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "group-id",
				Short:      `ID of the group`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "user-id",
				Short:      `ID of the user to add`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-id",
				Short:      `ID of the application to add`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.AddGroupMemberRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.AddGroupMember(request)

		},
	}
}

func iamGroupAddMembers() *core.Command {
	return &core.Command{
		Short:     `Add multiple users and applications to a group`,
		Long:      `Add multiple users and applications to a group in a single call. You can specify an array of ` + "`" + `user_id` + "`" + `s and ` + "`" + `application_id` + "`" + `s. Note that any existing users and applications in the group will remain. To add new users/applications and delete pre-existing ones, use the [Overwrite users and applications of a group](#path-groups-overwrite-users-and-applications-of-a-group) method.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "add-members",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.AddGroupMembersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "group-id",
				Short:      `ID of the group`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-ids.{index}",
				Short:      `IDs of the users to add`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-ids.{index}",
				Short:      `IDs of the applications to add`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.AddGroupMembersRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.AddGroupMembers(request)

		},
	}
}

func iamGroupRemoveMember() *core.Command {
	return &core.Command{
		Short:     `Remove a user or an application from a group`,
		Long:      `Remove a user or an application from a group. You can specify a ` + "`" + `user_id` + "`" + ` and and ` + "`" + `application_id` + "`" + ` in the body of your request. Note that you can only remove one of each per request. Removing a user from a group means that any permissions given to them via the group (i.e. from an attached policy) will no longer apply. Be sure you want to remove these permissions from the user before proceeding.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "remove-member",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.RemoveGroupMemberRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "group-id",
				Short:      `ID of the group`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "user-id",
				Short:      `ID of the user to remove`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-id",
				Short:      `ID of the application to remove`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.RemoveGroupMemberRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.RemoveGroupMember(request)

		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam group remove-member",
				Short:   "Remove a group member",
			},
			{
				Command: "scw iam group create",
				Short:   "Create a group",
			},
		},
	}
}

func iamGroupDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a group`,
		Long:      `Delete a group. Note that this action is irreversible and could delete permissions for group members. Policies attached to users and applications via this group will no longer apply.`,
		Namespace: "iam",
		Resource:  "group",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.DeleteGroupRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "group-id",
				Short:      `ID of the group to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.DeleteGroupRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			e = api.DeleteGroup(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "group",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete a given group",
				ArgsJSON: `{"group_id":"11111111-1111-1111-1111-111111111111"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam group list",
				Short:   "List all groups",
			},
			{
				Command: "scw iam group delete",
				Short:   "Delete a group",
			},
		},
	}
}

func iamPolicyList() *core.Command {
	return &core.Command{
		Short:     `List policies of an Organization`,
		Long:      `List the policies of an Organization. By default, the policies listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field. You must define the ` + "`" + `organization_id` + "`" + ` in the query path of your request. You can also define additional parameters to filter your query, such as ` + "`" + `user_ids` + "`" + `, ` + "`" + `groups_ids` + "`" + `, ` + "`" + `application_ids` + "`" + `, and ` + "`" + `policy_name` + "`" + `.`,
		Namespace: "iam",
		Resource:  "policy",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListPoliciesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"policy_name_asc", "policy_name_desc", "created_at_asc", "created_at_desc"},
			},
			{
				Name:       "editable",
				Short:      `Defines whether or not filter out editable policies`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-ids.{index}",
				Short:      `Defines whether or not to filter by list of user IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "group-ids.{index}",
				Short:      `Defines whether or not to filter by list of group IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-ids.{index}",
				Short:      `Filter by a list of application IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "no-principal",
				Short:      `Defines whether or not the policy is attributed to a principal`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "policy-name",
				Short:      `Name of the policy to fetch`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tag",
				Short:      `Filter by tags containing a given string`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "policy-ids.{index}",
				Short:      `Filter by a list of IDs`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListPoliciesRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListPolicies(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Policies, nil

		},
	}
}

func iamPolicyCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new policy`,
		Long:      `Create a new application. You must define the ` + "`" + `name` + "`" + ` parameter in the request. You can specify parameters such as ` + "`" + `user_id` + "`" + `, ` + "`" + `groups_id` + "`" + `, ` + "`" + `application_id` + "`" + `, ` + "`" + `no_principal` + "`" + `, ` + "`" + `rules` + "`" + ` and its child attributes.`,
		Namespace: "iam",
		Resource:  "policy",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.CreatePolicyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Name of the policy to create (max length is 64 characters)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("pol"),
			},
			{
				Name:       "description",
				Short:      `Description of the policy to create (max length is 200 characters)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.permission-set-names.{index}",
				Short:      `Names of permission sets bound to the rule`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.condition",
				Short:      `Condition expression to evaluate`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.project-ids.{index}",
				Short:      `List of Project IDs the rule is scoped to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.organization-id",
				Short:      `ID of Organization the rule is scoped to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `Tags associated with the policy (maximum of 10 tags)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-id",
				Short:      `ID of user attributed to the policy`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "group-id",
				Short:      `ID of group attributed to the policy`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-id",
				Short:      `ID of application attributed to the policy`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "no-principal",
				Short:      `Defines whether or not a policy is attributed to a principal`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.CreatePolicyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.CreatePolicy(request)

		},
		Examples: []*core.Example{
			{
				Short: "Add a policy for a group that gives InstanceFullAccess on all projects",
				Raw:   `scw iam policy create group-id=11111111-1111-1111-1111-111111111111 rules.0.organization-id=11111111-1111-1111-1111-111111111111 rules.0.permission-set-names.0=InstancesFullAccess`,
			},
		},
	}
}

func iamPolicyGet() *core.Command {
	return &core.Command{
		Short:     `Get an existing policy`,
		Long:      `Retrieve information about a policy, speficified by the ` + "`" + `policy_id` + "`" + ` parameter. The policy's full details, including ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `organization_id` + "`" + `, ` + "`" + `nb_rules` + "`" + ` and ` + "`" + `nb_scopes` + "`" + `, ` + "`" + `nb_permission_sets` + "`" + ` are returned in the response.`,
		Namespace: "iam",
		Resource:  "policy",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.GetPolicyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "policy-id",
				Short:      `Id of policy to search`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.GetPolicyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.GetPolicy(request)

		},
	}
}

func iamPolicyUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an existing policy`,
		Long:      `Update the parameters of a policy, including ` + "`" + `name` + "`" + `, ` + "`" + `description` + "`" + `, ` + "`" + `user_id` + "`" + `, ` + "`" + `group_id` + "`" + `, ` + "`" + `application_id` + "`" + ` and ` + "`" + `no_principal` + "`" + `.`,
		Namespace: "iam",
		Resource:  "policy",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.UpdatePolicyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "policy-id",
				Short:      `Id of policy to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `New name for the policy (max length is 64 characters)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `New description of policy (max length is 200 characters)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "tags.{index}",
				Short:      `New tags for the policy (maximum of 10 tags)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-id",
				Short:      `New ID of user attributed to the policy`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "group-id",
				Short:      `New ID of group attributed to the policy`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "application-id",
				Short:      `New ID of application attributed to the policy`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "no-principal",
				Short:      `Defines whether or not the policy is attributed to a principal`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.UpdatePolicyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.UpdatePolicy(request)

		},
	}
}

func iamPolicyDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a policy`,
		Long:      `Delete a policy. You must define specify the ` + "`" + `policy_id` + "`" + ` parameter in your request. Note that when deleting a policy, all permissions it gives to its principal (user, group or application) will be revoked.`,
		Namespace: "iam",
		Resource:  "policy",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.DeletePolicyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "policy-id",
				Short:      `Id of policy to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.DeletePolicyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			e = api.DeletePolicy(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "policy",
				Verb:     "delete",
			}, nil
		},
	}
}

func iamPolicyClone() *core.Command {
	return &core.Command{
		Short:     `Clone a policy`,
		Long:      `Clone a policy. You must define specify the ` + "`" + `policy_id` + "`" + ` parameter in your request.`,
		Namespace: "iam",
		Resource:  "policy",
		Verb:      "clone",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ClonePolicyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "policy-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ClonePolicyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.ClonePolicy(request)

		},
	}
}

func iamRuleUpdate() *core.Command {
	return &core.Command{
		Short:     `Set rules of a given policy`,
		Long:      `Overwrite the rules of a given policy. Any information that you add using this command will overwrite the previous configuration. If you include some of the rules you already had in your previous configuration in your new one, but you change their order, the new order of display will apply. While policy rules are ordered, they have no impact on the access logic of IAM because rules are allow-only.`,
		Namespace: "iam",
		Resource:  "rule",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.SetRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "policy-id",
				Short:      `Id of policy to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "rules.{index}.permission-set-names.{index}",
				Short:      `Names of permission sets bound to the rule`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.condition",
				Short:      `Condition expression to evaluate`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.project-ids.{index}",
				Short:      `List of Project IDs the rule is scoped to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rules.{index}.organization-id",
				Short:      `ID of Organization the rule is scoped to`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.SetRulesRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.SetRules(request)

		},
	}
}

func iamRuleList() *core.Command {
	return &core.Command{
		Short:     `List rules of a given policy`,
		Long:      `List the rules of a given policy. By default, the rules listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field. You must define the ` + "`" + `policy_id` + "`" + ` in the query path of your request.`,
		Namespace: "iam",
		Resource:  "rule",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListRulesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "policy-id",
				Short:      `Id of policy to search`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListRulesRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListRules(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Rules, nil

		},
	}
}

func iamPermissionSetList() *core.Command {
	return &core.Command{
		Short:     `List permission sets`,
		Long:      `List permission sets available for given Organization. You must define the ` + "`" + `organization_id` + "`" + ` in the query path of your request.`,
		Namespace: "iam",
		Resource:  "permission-set",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListPermissionSetsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"name_asc", "name_desc", "created_at_asc", "created_at_desc"},
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListPermissionSetsRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListPermissionSets(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.PermissionSets, nil

		},
	}
}

func iamAPIKeyList() *core.Command {
	return &core.Command{
		Short:     `List API keys`,
		Long:      `List API keys. By default, the API keys listed are ordered by creation date in ascending order. This can be modified via the ` + "`" + `order_by` + "`" + ` field. You can define additional parameters for your query such as ` + "`" + `editable` + "`" + `, ` + "`" + `expired` + "`" + `, ` + "`" + `access_key` + "`" + ` and ` + "`" + `bearer_id` + "`" + `.`,
		Namespace: "iam",
		Resource:  "api-key",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListAPIKeysRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "expires_at_asc", "expires_at_desc", "access_key_asc", "access_key_desc"},
			},
			{
				Name:       "application-id",
				Short:      `ID of application that bears the API key`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "user-id",
				Short:      `ID of user that bears the API key`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "editable",
				Short:      `Defines whether to filter out editable API keys or not`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "expired",
				Short:      `Defines whether to filter out expired API keys or not`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "access-key",
				Short:      `Filter by access key (deprecated in favor of ` + "`" + `access_keys` + "`" + `)`,
				Required:   false,
				Deprecated: true,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Filter by description`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "bearer-id",
				Short:      `Filter by bearer ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "bearer-type",
				Short:      `Filter by type of bearer`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_bearer_type", "user", "application"},
			},
			{
				Name:       "access-keys.{index}",
				Short:      `Filter by a list of access keys`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `ID of Organization`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListAPIKeysRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListAPIKeys(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.APIKeys, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "AccessKey",
			},
			{
				FieldName: "SecretKey",
			},
			{
				FieldName: "CreatedAt",
			},
			{
				FieldName: "ExpiresAt",
			},
			{
				FieldName: "DefaultProjectID",
			},
			{
				FieldName: "Description",
			},
		}},
	}
}

func iamAPIKeyCreate() *core.Command {
	return &core.Command{
		Short:     `Create an API key`,
		Long:      `Create an API key. You must specify the ` + "`" + `application_id` + "`" + ` or the ` + "`" + `user_id` + "`" + ` and the description. You can also specify the ` + "`" + `default_project_id` + "`" + ` which is the Project ID of your preferred Project, to use with Object Storage. The ` + "`" + `access_key` + "`" + ` and ` + "`" + `secret_key` + "`" + ` values are returned in the response. Note that he secret key is only showed once. Make sure that you copy and store both keys somewhere safe.`,
		Namespace: "iam",
		Resource:  "api-key",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.CreateAPIKeyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "application-id",
				Short:      `ID of the application`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "user-id",
				Short:      `ID of the user`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "expires-at",
				Short:      `Expiration date of the API key`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "default-project-id",
				Short:      `Default Project ID to use with Object Storage`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Description of the API key (max length is 200 characters)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.CreateAPIKeyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.CreateAPIKey(request)

		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam api-key list",
				Short:   "List all API keys",
			},
			{
				Command: "scw iam api-key delete",
				Short:   "Delete an API key",
			},
		},
	}
}

func iamAPIKeyGet() *core.Command {
	return &core.Command{
		Short:     `Get an API key`,
		Long:      `Retrive information about an API key, specified by the ` + "`" + `access_key` + "`" + ` parameter. The API key's details, including either the ` + "`" + `user_id` + "`" + ` or ` + "`" + `application_id` + "`" + ` of its bearer are returned in the response. Note that the string value for the ` + "`" + `secret_key` + "`" + ` is nullable, and therefore is not displayed in the response. The ` + "`" + `secret_key` + "`" + ` value is only displayed upon API key creation.`,
		Namespace: "iam",
		Resource:  "api-key",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.GetAPIKeyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "access-key",
				Short:      `Access key to search for`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.GetAPIKeyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.GetAPIKey(request)

		},
	}
}

func iamAPIKeyUpdate() *core.Command {
	return &core.Command{
		Short:     `Update an API key`,
		Long:      `Update the parameters of an API key, including ` + "`" + `default_project_id` + "`" + ` and ` + "`" + `description` + "`" + `.`,
		Namespace: "iam",
		Resource:  "api-key",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.UpdateAPIKeyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "access-key",
				Short:      `Access key to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "default-project-id",
				Short:      `New default Project ID to set`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `New description to update`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.UpdateAPIKeyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.UpdateAPIKey(request)

		},
	}
}

func iamAPIKeyDelete() *core.Command {
	return &core.Command{
		Short:     `Delete an API key`,
		Long:      `Delete an API key. Note that this action is irreversible and cannot be undone. Make sure you update any configurations using the API keys you delete.`,
		Namespace: "iam",
		Resource:  "api-key",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.DeleteAPIKeyRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "access-key",
				Short:      `Access key to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.DeleteAPIKeyRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			e = api.DeleteAPIKey(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "api-key",
				Verb:     "delete",
			}, nil
		},
		Examples: []*core.Example{
			{
				Short:    "Delete a given API key",
				ArgsJSON: `{"access_key":"SCW00000000000"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw iam api-key list",
				Short:   "List all API keys",
			},
			{
				Command: "scw iam api-key create",
				Short:   "Create an API key",
			},
		},
	}
}

func iamJwtList() *core.Command {
	return &core.Command{
		Short:     `List JWTs`,
		Long:      `List JWTs.`,
		Namespace: "iam",
		Resource:  "jwt",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListJWTsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc"},
			},
			{
				Name:       "audience-id",
				Short:      `ID of the user to search`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "expired",
				Short:      `Filter out expired JWTs or not`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListJWTsRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListJWTs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Jwts, nil

		},
	}
}

func iamJwtGet() *core.Command {
	return &core.Command{
		Short:     `Get a JWT`,
		Long:      `Get a JWT.`,
		Namespace: "iam",
		Resource:  "jwt",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.GetJWTRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "jti",
				Short:      `JWT ID of the JWT to get`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.GetJWTRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.GetJWT(request)

		},
	}
}

func iamJwtDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a JWT`,
		Long:      `Delete a JWT.`,
		Namespace: "iam",
		Resource:  "jwt",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.DeleteJWTRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "jti",
				Short:      `JWT ID of the JWT to delete`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.DeleteJWTRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			e = api.DeleteJWT(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "jwt",
				Verb:     "delete",
			}, nil
		},
	}
}

func iamLogList() *core.Command {
	return &core.Command{
		Short:     `List logs`,
		Long:      `List logs available for given Organization. You must define the ` + "`" + `organization_id` + "`" + ` in the query path of your request.`,
		Namespace: "iam",
		Resource:  "log",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.ListLogsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Criteria for sorting results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("created_at_asc"),
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:       "created-after",
				Short:      `Defined whether or not to filter out logs created after this timestamp`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "created-before",
				Short:      `Defined whether or not to filter out logs created before this timestamp`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "action",
				Short:      `Defined whether or not to filter out by a specific action`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_action", "created", "updated", "deleted"},
			},
			{
				Name:       "resource-type",
				Short:      `Defined whether or not to filter out by a specific type of resource`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_resource_type", "api_key", "user", "application", "group", "policy"},
			},
			{
				Name:       "search",
				Short:      `Defined whether or not to filter out log by bearer ID or resource ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.OrganizationIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.ListLogsRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListLogs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Logs, nil

		},
	}
}

func iamLogGet() *core.Command {
	return &core.Command{
		Short:     `Get a log`,
		Long:      `Retrieve information about a log, specified by the ` + "`" + `log_id` + "`" + ` parameter. The log's full details, including ` + "`" + `id` + "`" + `, ` + "`" + `ip` + "`" + `, ` + "`" + `user_agent` + "`" + `, ` + "`" + `action` + "`" + `, ` + "`" + `bearer_id` + "`" + `, ` + "`" + `resource_type` + "`" + ` and ` + "`" + `resource_id` + "`" + ` are returned in the response.`,
		Namespace: "iam",
		Resource:  "log",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iam.GetLogRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "log-id",
				Short:      `ID of the log`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iam.GetLogRequest)

			client := core.ExtractClient(ctx)
			api := iam.NewAPI(client)
			return api.GetLog(request)

		},
	}
}

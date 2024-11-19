// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package cockpit

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/cockpit/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		cockpitRoot(),
		cockpitGrafana(),
		cockpitGrafanaUser(),
		cockpitProductDashboards(),
		cockpitPlan(),
		cockpitDataSource(),
		cockpitToken(),
		cockpitAlertManager(),
		cockpitContactPoint(),
		cockpitManagedAlerts(),
		cockpitTestAlert(),
		cockpitUsageOverview(),
		cockpitGrafanaGet(),
		cockpitGrafanaSyncDataSources(),
		cockpitGrafanaUserCreate(),
		cockpitGrafanaUserList(),
		cockpitGrafanaUserDelete(),
		cockpitGrafanaUserResetPassword(),
		cockpitProductDashboardsList(),
		cockpitProductDashboardsGet(),
		cockpitPlanList(),
		cockpitPlanSelect(),
		cockpitPlanGet(),
		cockpitDataSourceCreate(),
		cockpitDataSourceGet(),
		cockpitDataSourceDelete(),
		cockpitDataSourceList(),
		cockpitDataSourceUpdate(),
		cockpitUsageOverviewGet(),
		cockpitTokenCreate(),
		cockpitTokenList(),
		cockpitTokenGet(),
		cockpitTokenDelete(),
		cockpitAlertManagerGet(),
		cockpitAlertManagerEnable(),
		cockpitAlertManagerDisable(),
		cockpitContactPointCreate(),
		cockpitContactPointList(),
		cockpitContactPointDelete(),
		cockpitManagedAlertsEnable(),
		cockpitManagedAlertsDisable(),
		cockpitTestAlertTrigger(),
	)
}
func cockpitRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your Scaleway Cockpit, for storing and visualizing metrics and logs`,
		Long:      `This API allows you to manage your Scaleway Cockpit, for storing and visualizing metrics and logs.`,
		Namespace: "cockpit",
	}
}

func cockpitGrafana() *core.Command {
	return &core.Command{
		Short:     `Grafana user management commands`,
		Long:      `Grafana user management commands.`,
		Namespace: "cockpit",
		Resource:  "grafana",
	}
}

func cockpitGrafanaUser() *core.Command {
	return &core.Command{
		Short:     `Grafana user management commands`,
		Long:      `Grafana user management commands.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
	}
}

func cockpitProductDashboards() *core.Command {
	return &core.Command{
		Short:     `Product dashboards management commands`,
		Long:      `Product dashboards management commands.`,
		Namespace: "cockpit",
		Resource:  "product-dashboards",
	}
}

func cockpitPlan() *core.Command {
	return &core.Command{
		Short:     `Pricing plans management commands`,
		Long:      `Pricing plans management commands.`,
		Namespace: "cockpit",
		Resource:  "plan",
	}
}

func cockpitDataSource() *core.Command {
	return &core.Command{
		Short:     `Datasource management commands`,
		Long:      `Datasource management commands.`,
		Namespace: "cockpit",
		Resource:  "data-source",
	}
}

func cockpitToken() *core.Command {
	return &core.Command{
		Short:     `Token management commands`,
		Long:      `Token management commands.`,
		Namespace: "cockpit",
		Resource:  "token",
	}
}

func cockpitAlertManager() *core.Command {
	return &core.Command{
		Short:     `Alerting management commands`,
		Long:      `Alerting management commands.`,
		Namespace: "cockpit",
		Resource:  "alert-manager",
	}
}

func cockpitContactPoint() *core.Command {
	return &core.Command{
		Short:     `Contact point management commands`,
		Long:      `Contact point management commands.`,
		Namespace: "cockpit",
		Resource:  "contact-point",
	}
}

func cockpitManagedAlerts() *core.Command {
	return &core.Command{
		Short:     `Managed alerts management commands`,
		Long:      `Managed alerts management commands.`,
		Namespace: "cockpit",
		Resource:  "managed-alerts",
	}
}

func cockpitTestAlert() *core.Command {
	return &core.Command{
		Short:     `Test alert management commands`,
		Long:      `Test alert management commands.`,
		Namespace: "cockpit",
		Resource:  "test-alert",
	}
}

func cockpitUsageOverview() *core.Command {
	return &core.Command{
		Short:     `Usage overview management commands`,
		Long:      `Usage overview management commands.`,
		Namespace: "cockpit",
		Resource:  "usage-overview",
	}
}

func cockpitGrafanaGet() *core.Command {
	return &core.Command{
		Short: `Get your Cockpit's Grafana`,
		Long: `Retrieve information on your Cockpit's Grafana, specified by the ID of the Project the Cockpit belongs to.
The output returned displays the URL to access your Cockpit's Grafana.`,
		Namespace: "cockpit",
		Resource:  "grafana",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPIGetGrafanaRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPIGetGrafanaRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			return api.GetGrafana(request)

		},
	}
}

func cockpitGrafanaSyncDataSources() *core.Command {
	return &core.Command{
		Short:     `Synchronize Grafana data sources`,
		Long:      `Trigger the synchronization of all your data sources and the alert manager in the relevant regions. The alert manager will only be synchronized if you have enabled it.`,
		Namespace: "cockpit",
		Resource:  "grafana",
		Verb:      "sync-data-sources",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPISyncGrafanaDataSourcesRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPISyncGrafanaDataSourcesRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			e = api.SyncGrafanaDataSources(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "grafana",
				Verb:     "sync-data-sources",
			}, nil
		},
	}
}

func cockpitGrafanaUserCreate() *core.Command {
	return &core.Command{
		Short: `Create a Grafana user`,
		Long: `Create a Grafana user to connect to your Cockpit's Grafana. Upon creation, your user password displays only once, so make sure that you save it.
Each Grafana user is associated with a role: viewer or editor. A viewer can only view dashboards, whereas an editor can create and edit dashboards. Note that the ` + "`" + `admin` + "`" + ` username is not available for creation.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPICreateGrafanaUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "login",
				Short:      `Username of the Grafana user. Note that the ` + "`" + `admin` + "`" + ` username is not available for creation`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "role",
				Short:      `Role assigned to the Grafana user`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_role", "editor", "viewer"},
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPICreateGrafanaUserRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			return api.CreateGrafanaUser(request)

		},
	}
}

func cockpitGrafanaUserList() *core.Command {
	return &core.Command{
		Short:     `List Grafana users`,
		Long:      `List all Grafana users created in your Cockpit's Grafana. By default, the Grafana users returned in the list are ordered in ascending order.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPIListGrafanaUsersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the Grafana users`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"login_asc", "login_desc"},
			},
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPIListGrafanaUsersRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListGrafanaUsers(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.GrafanaUsers, nil

		},
	}
}

func cockpitGrafanaUserDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Grafana user`,
		Long:      `Delete a Grafana user from your Cockpit's Grafana, specified by the ID of the Project the Cockpit belongs to, and the ID of the Grafana user.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPIDeleteGrafanaUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "grafana-user-id",
				Short:      `ID of the Grafana user`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPIDeleteGrafanaUserRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			e = api.DeleteGrafanaUser(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "grafana-user",
				Verb:     "delete",
			}, nil
		},
	}
}

func cockpitGrafanaUserResetPassword() *core.Command {
	return &core.Command{
		Short: `Reset a Grafana user password`,
		Long: `Reset the password of a Grafana user, specified by the ID of the Project the Cockpit belongs to, and the ID of the Grafana user.
A new password regenerates and only displays once. Make sure that you save it.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
		Verb:      "reset-password",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPIResetGrafanaUserPasswordRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "grafana-user-id",
				Short:      `ID of the Grafana user`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPIResetGrafanaUserPasswordRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			return api.ResetGrafanaUserPassword(request)

		},
	}
}

func cockpitProductDashboardsList() *core.Command {
	return &core.Command{
		Short:     `List Scaleway resources dashboards`,
		Long:      `Retrieve a list of available dashboards in Grafana, for all Scaleway resources which are integrated with Cockpit.`,
		Namespace: "cockpit",
		Resource:  "product-dashboards",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPIListGrafanaProductDashboardsRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "tags.{index}",
				Short:      `Tags to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPIListGrafanaProductDashboardsRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListGrafanaProductDashboards(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Dashboards, nil

		},
	}
}

func cockpitProductDashboardsGet() *core.Command {
	return &core.Command{
		Short:     `Get Scaleway resource dashboard`,
		Long:      `Retrieve information about the dashboard of a Scaleway resource in Grafana, specified by the ID of the Project the Cockpit belongs to, and the name of the dashboard.`,
		Namespace: "cockpit",
		Resource:  "product-dashboards",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPIGetGrafanaProductDashboardRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "dashboard-name",
				Short:      `Name of the dashboard`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPIGetGrafanaProductDashboardRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			return api.GetGrafanaProductDashboard(request)

		},
	}
}

func cockpitPlanList() *core.Command {
	return &core.Command{
		Short: `List plan types`,
		Long: `Retrieve a list of available pricing plan types.
Deprecated, retention is now managed at the data source level.`,
		Namespace: "cockpit",
		Resource:  "plan",
		Verb:      "list",
		// Deprecated:    true,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPIListPlansRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc"},
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPIListPlansRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListPlans(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Plans, nil

		},
	}
}

func cockpitPlanSelect() *core.Command {
	return &core.Command{
		Short: `Apply a pricing plan`,
		Long: `Apply a pricing plan on a given Project. You must specify the ID of the pricing plan type. Note that you will be billed for the plan you apply.
Deprecated, retention is now managed at the data source level.`,
		Namespace: "cockpit",
		Resource:  "plan",
		Verb:      "select",
		// Deprecated:    true,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPISelectPlanRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "plan-name",
				Short:      `Name of the pricing plan`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_name", "free", "premium", "custom"},
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPISelectPlanRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			return api.SelectPlan(request)

		},
	}
}

func cockpitPlanGet() *core.Command {
	return &core.Command{
		Short: `Get current plan`,
		Long: `Retrieve a pricing plan for the given Project, specified by the ID of the Project.
Deprecated, retention is now managed at the data source level.`,
		Namespace: "cockpit",
		Resource:  "plan",
		Verb:      "get",
		// Deprecated:    true,
		ArgsType: reflect.TypeOf(cockpit.GlobalAPIGetCurrentPlanRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GlobalAPIGetCurrentPlanRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewGlobalAPI(client)
			return api.GetCurrentPlan(request)

		},
	}
}

func cockpitDataSourceCreate() *core.Command {
	return &core.Command{
		Short: `Create a data source`,
		Long: `You must specify the data source type upon creation. Available data source types include:
  - metrics
  - logs
  - traces
The name of the data source will then be used as reference to name the associated Grafana data source.`,
		Namespace: "cockpit",
		Resource:  "data-source",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPICreateDataSourceRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "name",
				Short:      `Data source name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "type",
				Short:      `Data source type`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "metrics", "logs", "traces"},
			},
			{
				Name:       "retention-days",
				Short:      `BETA - Duration for which the data will be retained in the data source`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPICreateDataSourceRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.CreateDataSource(request)

		},
	}
}

func cockpitDataSourceGet() *core.Command {
	return &core.Command{
		Short:     `Get a data source`,
		Long:      `Retrieve information about a given data source, specified by the data source ID. The data source's information such as its name, type, URL, origin, and retention period, is returned.`,
		Namespace: "cockpit",
		Resource:  "data-source",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIGetDataSourceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "data-source-id",
				Short:      `ID of the relevant data source`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIGetDataSourceRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.GetDataSource(request)

		},
	}
}

func cockpitDataSourceDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a data source`,
		Long:      `Delete a given data source, specified by the data source ID. Note that deleting a data source is irreversible, and cannot be undone.`,
		Namespace: "cockpit",
		Resource:  "data-source",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIDeleteDataSourceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "data-source-id",
				Short:      `ID of the data source to delete`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIDeleteDataSourceRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			e = api.DeleteDataSource(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "data-source",
				Verb:     "delete",
			}, nil
		},
	}
}

func cockpitDataSourceList() *core.Command {
	return &core.Command{
		Short: `List data sources`,
		Long: `Retrieve the list of data sources available in the specified region. By default, the data sources returned in the list are ordered by creation date, in ascending order.
You can list data sources by Project, type and origin.`,
		Namespace: "cockpit",
		Resource:  "data-source",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIListDataSourcesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order for data sources in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc", "type_asc", "type_desc"},
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "origin",
				Short:      `Origin to filter for, only data sources with matching origin will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_origin", "scaleway", "external"},
			},
			{
				Name:       "types.{index}",
				Short:      `Types to filter for, only data sources with matching types will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_type", "metrics", "logs", "traces"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIListDataSourcesRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListDataSources(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.DataSources, nil

		},
	}
}

func cockpitDataSourceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a data source`,
		Long:      `Update a given data source name, specified by the data source ID.`,
		Namespace: "cockpit",
		Resource:  "data-source",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIUpdateDataSourceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "data-source-id",
				Short:      `ID of the data source to update`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Updated name of the data source`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "retention-days",
				Short:      `BETA - Duration for which the data will be retained in the data source`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIUpdateDataSourceRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.UpdateDataSource(request)

		},
	}
}

func cockpitUsageOverviewGet() *core.Command {
	return &core.Command{
		Short:     `Get data source usage overview`,
		Long:      `Retrieve the data source usage overview per type for the specified Project.`,
		Namespace: "cockpit",
		Resource:  "usage-overview",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIGetUsageOverviewRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "interval",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIGetUsageOverviewRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.GetUsageOverview(request)

		},
	}
}

func cockpitTokenCreate() *core.Command {
	return &core.Command{
		Short: `Create a token`,
		Long: `Give your token the relevant scopes to ensure it has the right permissions to interact with your data sources and the Alert manager. Make sure that you create your token in the same regions as the data sources you want to use it for.
Upon creation, your token's secret key display only once. Make sure that you save it.`,
		Namespace: "cockpit",
		Resource:  "token",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPICreateTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "name",
				Short:      `Name of the token`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "token-scopes.{index}",
				Short:      `Token permission scopes`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_scope", "read_only_metrics", "write_only_metrics", "full_access_metrics_rules", "read_only_logs", "write_only_logs", "full_access_logs_rules", "full_access_alert_manager", "read_only_traces", "write_only_traces"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPICreateTokenRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.CreateToken(request)

		},
	}
}

func cockpitTokenList() *core.Command {
	return &core.Command{
		Short: `List tokens`,
		Long: `Retrieve a list of all tokens in the specified region. By default, tokens returned in the list are ordered by creation date, in ascending order.
You can filter tokens by Project ID and token scopes.`,
		Namespace: "cockpit",
		Resource:  "token",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIListTokensRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order in which to return results`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "token-scopes.{index}",
				Short:      `Token scopes to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_scope", "read_only_metrics", "write_only_metrics", "full_access_metrics_rules", "read_only_logs", "write_only_logs", "full_access_logs_rules", "full_access_alert_manager", "read_only_traces", "write_only_traces"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIListTokensRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
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

func cockpitTokenGet() *core.Command {
	return &core.Command{
		Short:     `Get a token`,
		Long:      `Retrieve information about a given token, specified by the token ID. The token's information such as its scopes, is returned.`,
		Namespace: "cockpit",
		Resource:  "token",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIGetTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "token-id",
				Short:      `Token ID`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIGetTokenRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.GetToken(request)

		},
	}
}

func cockpitTokenDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a token`,
		Long:      `Delete a given token, specified by the token ID. Deleting a token is irreversible and cannot be undone.`,
		Namespace: "cockpit",
		Resource:  "token",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIDeleteTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "token-id",
				Short:      `ID of the token to delete`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIDeleteTokenRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			e = api.DeleteToken(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "token",
				Verb:     "delete",
			}, nil
		},
	}
}

func cockpitAlertManagerGet() *core.Command {
	return &core.Command{
		Short: `Get the Alert manager`,
		Long: `Retrieve information about the Alert manager which is unique per Project and region. By default the Alert manager is disabled.
The output returned displays a URL to access the Alert manager, and whether the Alert manager and managed alerts are enabled.`,
		Namespace: "cockpit",
		Resource:  "alert-manager",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIGetAlertManagerRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIGetAlertManagerRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.GetAlertManager(request)

		},
	}
}

func cockpitAlertManagerEnable() *core.Command {
	return &core.Command{
		Short:     `Enable the Alert manager`,
		Long:      `Enabling the Alert manager allows you to enable managed alerts and create contact points in the specified Project and region, to be notified when your Scaleway resources may require your attention.`,
		Namespace: "cockpit",
		Resource:  "alert-manager",
		Verb:      "enable",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIEnableAlertManagerRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIEnableAlertManagerRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.EnableAlertManager(request)

		},
	}
}

func cockpitAlertManagerDisable() *core.Command {
	return &core.Command{
		Short:     `Disable the Alert manager`,
		Long:      `Disabling the Alert manager deletes the contact points you have created and disables managed alerts in the specified Project and region.`,
		Namespace: "cockpit",
		Resource:  "alert-manager",
		Verb:      "disable",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIDisableAlertManagerRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIDisableAlertManagerRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.DisableAlertManager(request)

		},
	}
}

func cockpitContactPointCreate() *core.Command {
	return &core.Command{
		Short: `Create a contact point`,
		Long: `Contact points are email addresses associated with the default receiver, that the Alert manager sends alerts to.
The source of the alerts are data sources within the same Project and region as the Alert manager.
If you need to receive alerts for other receivers, you can create additional contact points and receivers in Grafana. Make sure that you select the Scaleway Alert manager.`,
		Namespace: "cockpit",
		Resource:  "contact-point",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPICreateContactPointRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "email.to",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPICreateContactPointRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.CreateContactPoint(request)

		},
	}
}

func cockpitContactPointList() *core.Command {
	return &core.Command{
		Short:     `List contact points`,
		Long:      `Retrieve a list of contact points for the specified Project. The response lists all contact points and receivers created in Grafana or via the API.`,
		Namespace: "cockpit",
		Resource:  "contact-point",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIListContactPointsRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIListContactPointsRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListContactPoints(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.ContactPoints, nil

		},
	}
}

func cockpitContactPointDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a contact point`,
		Long:      `Delete a contact point associated with the default receiver.`,
		Namespace: "cockpit",
		Resource:  "contact-point",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIDeleteContactPointRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "email.to",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIDeleteContactPointRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			e = api.DeleteContactPoint(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "contact-point",
				Verb:     "delete",
			}, nil
		},
	}
}

func cockpitManagedAlertsEnable() *core.Command {
	return &core.Command{
		Short:     `Enable managed alerts`,
		Long:      `Enable the sending of managed alerts for the specified Project. Managed alerts are predefined alerts that apply to Scaleway recources integrated with Cockpit by default.`,
		Namespace: "cockpit",
		Resource:  "managed-alerts",
		Verb:      "enable",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIEnableManagedAlertsRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIEnableManagedAlertsRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.EnableManagedAlerts(request)

		},
	}
}

func cockpitManagedAlertsDisable() *core.Command {
	return &core.Command{
		Short:     `Disable managed alerts`,
		Long:      `Disable the sending of managed alerts for the specified Project.`,
		Namespace: "cockpit",
		Resource:  "managed-alerts",
		Verb:      "disable",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPIDisableManagedAlertsRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPIDisableManagedAlertsRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			return api.DisableManagedAlerts(request)

		},
	}
}

func cockpitTestAlertTrigger() *core.Command {
	return &core.Command{
		Short:     `Trigger a test alert`,
		Long:      `Send a test alert to the Alert manager to make sure your contact points get notified.`,
		Namespace: "cockpit",
		Resource:  "test-alert",
		Verb:      "trigger",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.RegionalAPITriggerTestAlertRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.RegionalAPITriggerTestAlertRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewRegionalAPI(client)
			e = api.TriggerTestAlert(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "test-alert",
				Verb:     "trigger",
			}, nil
		},
	}
}

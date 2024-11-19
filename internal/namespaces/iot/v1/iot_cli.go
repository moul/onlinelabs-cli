// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package iot

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-sdk-go/api/iot/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		iotRoot(),
		iotHub(),
		iotDevice(),
		iotRoute(),
		iotNetwork(),
		iotHubList(),
		iotHubCreate(),
		iotHubGet(),
		iotHubUpdate(),
		iotHubEnable(),
		iotHubDisable(),
		iotHubDelete(),
		iotHubSetCa(),
		iotHubGetCa(),
		iotDeviceList(),
		iotDeviceCreate(),
		iotDeviceGet(),
		iotDeviceUpdate(),
		iotDeviceEnable(),
		iotDeviceDisable(),
		iotDeviceRenewCertificate(),
		iotDeviceSetCertificate(),
		iotDeviceGetCertificate(),
		iotDeviceDelete(),
		iotDeviceGetMetrics(),
		iotRouteList(),
		iotRouteCreate(),
		iotRouteUpdate(),
		iotRouteGet(),
		iotRouteDelete(),
		iotNetworkList(),
		iotNetworkCreate(),
		iotNetworkGet(),
		iotNetworkDelete(),
	)
}
func iotRoot() *core.Command {
	return &core.Command{
		Short:     `This API allows you to manage your IoT hubs and devices`,
		Long:      `This API allows you to manage your IoT hubs and devices.`,
		Namespace: "iot",
	}
}

func iotHub() *core.Command {
	return &core.Command{
		Short:     `IoT Hub commands`,
		Long:      `IoT Hub commands.`,
		Namespace: "iot",
		Resource:  "hub",
	}
}

func iotDevice() *core.Command {
	return &core.Command{
		Short:     `IoT Device commands`,
		Long:      `IoT Device commands.`,
		Namespace: "iot",
		Resource:  "device",
	}
}

func iotRoute() *core.Command {
	return &core.Command{
		Short:     `IoT Route commands`,
		Long:      `IoT Route commands.`,
		Namespace: "iot",
		Resource:  "route",
	}
}

func iotNetwork() *core.Command {
	return &core.Command{
		Short:     `IoT Network commands`,
		Long:      `IoT Network commands.`,
		Namespace: "iot",
		Resource:  "network",
	}
}

func iotHubList() *core.Command {
	return &core.Command{
		Short:     `List hubs`,
		Long:      `List all Hubs in the specified zone. By default, returned Hubs are ordered by creation date in ascending order, though this can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.ListHubsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Sort order of Hubs in the response`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc", "status_asc", "status_desc", "product_plan_asc", "product_plan_desc", "created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc"},
			},
			{
				Name:       "project-id",
				Short:      `Only list Hubs of this Project ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Hub name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "organization-id",
				Short:      `Only list Hubs of this Organization ID`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.ListHubsRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListHubs(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Hubs, nil

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
				FieldName: "ProductPlan",
			},
			{
				FieldName: "Enabled",
			},
			{
				FieldName: "DeviceCount",
			},
			{
				FieldName: "ConnectedDeviceCount",
			},
			{
				FieldName: "Endpoint",
			},
			{
				FieldName: "DisableEvents",
			},
			{
				FieldName: "EventsTopicPrefix",
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
			{
				FieldName: "ProjectID",
			},
			{
				FieldName: "OrganizationID",
			},
		}},
	}
}

func iotHubCreate() *core.Command {
	return &core.Command{
		Short:     `Create a hub`,
		Long:      `Create a new Hub in the targeted region, specifying its configuration including name and product plan.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.CreateHubRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Hub name (up to 255 characters)`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("hub"),
			},
			core.ProjectIDArgSpec(),
			{
				Name:       "product-plan",
				Short:      `Hub product plan`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.DefaultValueSetter("plan_shared"),
				EnumValues: []string{"plan_unknown", "plan_shared", "plan_dedicated", "plan_ha"},
			},
			{
				Name:       "disable-events",
				Short:      `Disable Hub events`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "events-topic-prefix",
				Short:      `Topic prefix (default '$SCW/events') of Hub events`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "twins-graphite-config.push-uri",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.CreateHubRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.CreateHub(request)

		},
	}
}

func iotHubGet() *core.Command {
	return &core.Command{
		Short:     `Get a hub`,
		Long:      `Retrieve information about an existing IoT Hub, specified by its Hub ID. Its full details, including name, status and endpoint, are returned in the response object.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.GetHubRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hub-id",
				Short:      `Hub ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.GetHubRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.GetHub(request)

		},
	}
}

func iotHubUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a hub`,
		Long:      `Update the parameters of an existing IoT Hub, specified by its Hub ID.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.UpdateHubRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hub-id",
				Short:      `ID of the Hub you want to update`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "name",
				Short:      `Hub name (up to 255 characters)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "product-plan",
				Short:      `Hub product plan`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"plan_unknown", "plan_shared", "plan_dedicated", "plan_ha"},
			},
			{
				Name:       "disable-events",
				Short:      `Disable Hub events`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "events-topic-prefix",
				Short:      `Topic prefix of Hub events`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "enable-device-auto-provisioning",
				Short:      `Enable device auto provisioning`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "twins-graphite-config.push-uri",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.UpdateHubRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.UpdateHub(request)

		},
	}
}

func iotHubEnable() *core.Command {
	return &core.Command{
		Short:     `Enable a hub`,
		Long:      `Enable an existing IoT Hub, specified by its Hub ID.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "enable",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.EnableHubRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hub-id",
				Short:      `Hub ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.EnableHubRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.EnableHub(request)

		},
	}
}

func iotHubDisable() *core.Command {
	return &core.Command{
		Short:     `Disable a hub`,
		Long:      `Disable an existing IoT Hub, specified by its Hub ID.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "disable",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.DisableHubRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hub-id",
				Short:      `Hub ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.DisableHubRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.DisableHub(request)

		},
	}
}

func iotHubDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a hub`,
		Long:      `Delete an existing IoT Hub, specified by its Hub ID. Deleting a Hub is permanent, and cannot be undone.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.DeleteHubRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hub-id",
				Short:      `Hub ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "delete-devices",
				Short:      `Defines whether to force the deletion of devices added to this Hub or reject the operation`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.DeleteHubRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			e = api.DeleteHub(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "hub",
				Verb:     "delete",
			}, nil
		},
	}
}

func iotHubSetCa() *core.Command {
	return &core.Command{
		Short:     `Set the certificate authority of a hub`,
		Long:      `Set a particular PEM-encoded certificate, specified by the Hub ID.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "set-ca",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.SetHubCARequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hub-id",
				Short:      `Hub ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "ca-cert-pem",
				Short:      `CA's PEM-encoded certificate`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "challenge-cert-pem",
				Short:      `Proof of possession of PEM-encoded certificate`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.SetHubCARequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.SetHubCA(request)

		},
	}
}

func iotHubGetCa() *core.Command {
	return &core.Command{
		Short:     `Get the certificate authority of a hub`,
		Long:      `Get information for a particular PEM-encoded certificate, specified by the Hub ID.`,
		Namespace: "iot",
		Resource:  "hub",
		Verb:      "get-ca",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.GetHubCARequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "hub-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.GetHubCARequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.GetHubCA(request)

		},
	}
}

func iotDeviceList() *core.Command {
	return &core.Command{
		Short:     `List devices`,
		Long:      `List all devices in the specified region. By default, returned devices are ordered by creation date in ascending order, though this can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.ListDevicesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Ordering of requested devices`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc", "status_asc", "status_desc", "hub_id_asc", "hub_id_desc", "created_at_asc", "created_at_desc", "updated_at_asc", "updated_at_desc", "allow_insecure_asc", "allow_insecure_desc"},
			},
			{
				Name:       "name",
				Short:      `Name to filter for, only devices with this name will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "hub-id",
				Short:      `Hub ID to filter for, only devices attached to this Hub will be returned`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "allow-insecure",
				Short:      `Defines wheter to filter the allow_insecure flag`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "status",
				Short:      `Device status (enabled, disabled, etc.)`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "error", "enabled", "disabled"},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.ListDevicesRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListDevices(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Devices, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "Description",
			},
			{
				FieldName: "Status",
			},
			{
				FieldName: "HubID",
			},
			{
				FieldName: "LastActivityAt",
			},
			{
				FieldName: "IsConnected",
			},
			{
				FieldName: "AllowInsecure",
			},
			{
				FieldName: "AllowMultipleConnections",
			},
			{
				FieldName: "MessageFilters",
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

func iotDeviceCreate() *core.Command {
	return &core.Command{
		Short:     `Add a device`,
		Long:      `Attach a device to a given Hub.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.CreateDeviceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Device name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("device"),
			},
			{
				Name:       "hub-id",
				Short:      `Hub ID of the device`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "allow-insecure",
				Short:      `Defines whether to allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "allow-multiple-connections",
				Short:      `Defines whether to allow multiple physical devices to connect with this device's credentials`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "message-filters.publish.policy",
				Short:      `How to use the topic list`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "accept", "reject"},
			},
			{
				Name:       "message-filters.publish.topics.{index}",
				Short:      `List of topics to accept or reject. It must be valid MQTT topics and up to 65535 characters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "message-filters.subscribe.policy",
				Short:      `How to use the topic list`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "accept", "reject"},
			},
			{
				Name:       "message-filters.subscribe.topics.{index}",
				Short:      `List of topics to accept or reject. It must be valid MQTT topics and up to 65535 characters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "description",
				Short:      `Device description`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.CreateDeviceRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.CreateDevice(request)

		},
	}
}

func iotDeviceGet() *core.Command {
	return &core.Command{
		Short:     `Get a device`,
		Long:      `Retrieve information about an existing device, specified by its device ID. Its full details, including name, status and ID, are returned in the response object.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.GetDeviceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.GetDeviceRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.GetDevice(request)

		},
	}
}

func iotDeviceUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a device`,
		Long:      `Update the parameters of an existing device, specified by its device ID.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.UpdateDeviceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "description",
				Short:      `Description for the device`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "allow-insecure",
				Short:      `Defines whether to allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "allow-multiple-connections",
				Short:      `Defines whether to allow multiple physical devices to connect with this device's credentials`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "message-filters.publish.policy",
				Short:      `How to use the topic list`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "accept", "reject"},
			},
			{
				Name:       "message-filters.publish.topics.{index}",
				Short:      `List of topics to accept or reject. It must be valid MQTT topics and up to 65535 characters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "message-filters.subscribe.policy",
				Short:      `How to use the topic list`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "accept", "reject"},
			},
			{
				Name:       "message-filters.subscribe.topics.{index}",
				Short:      `List of topics to accept or reject. It must be valid MQTT topics and up to 65535 characters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "hub-id",
				Short:      `Change Hub for this device, additional fees may apply, see IoT Hub pricing`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.UpdateDeviceRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.UpdateDevice(request)

		},
	}
}

func iotDeviceEnable() *core.Command {
	return &core.Command{
		Short:     `Enable a device`,
		Long:      `Enable a specific device, specified by its device ID.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "enable",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.EnableDeviceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.EnableDeviceRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.EnableDevice(request)

		},
	}
}

func iotDeviceDisable() *core.Command {
	return &core.Command{
		Short:     `Disable a device`,
		Long:      `Disable an existing device, specified by its device ID.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "disable",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.DisableDeviceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.DisableDeviceRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.DisableDevice(request)

		},
	}
}

func iotDeviceRenewCertificate() *core.Command {
	return &core.Command{
		Short:     `Renew a device certificate`,
		Long:      `Renew the certificate of an existing device, specified by its device ID.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "renew-certificate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.RenewDeviceCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.RenewDeviceCertificateRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.RenewDeviceCertificate(request)

		},
	}
}

func iotDeviceSetCertificate() *core.Command {
	return &core.Command{
		Short:     `Set a custom certificate on a device`,
		Long:      `Switch the existing certificate of a given device with an EM-encoded custom certificate.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "set-certificate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.SetDeviceCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "certificate-pem",
				Short:      `PEM-encoded custom certificate`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.SetDeviceCertificateRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.SetDeviceCertificate(request)

		},
	}
}

func iotDeviceGetCertificate() *core.Command {
	return &core.Command{
		Short:     `Get a device's certificate`,
		Long:      `Get information for a particular PEM-encoded certificate, specified by the device ID. The response returns full details of the device, including its type of certificate.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "get-certificate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.GetDeviceCertificateRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.GetDeviceCertificateRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.GetDeviceCertificate(request)

		},
	}
}

func iotDeviceDelete() *core.Command {
	return &core.Command{
		Short:     `Remove a device`,
		Long:      `Remove a specific device from the specific Hub it is attached to.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.DeleteDeviceRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.DeleteDeviceRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			e = api.DeleteDevice(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "device",
				Verb:     "delete",
			}, nil
		},
	}
}

func iotDeviceGetMetrics() *core.Command {
	return &core.Command{
		Short:     `Get a device's metrics`,
		Long:      `Get the metrics of an existing device, specified by its device ID.`,
		Namespace: "iot",
		Resource:  "device",
		Verb:      "get-metrics",
		// Deprecated:    true,
		ArgsType: reflect.TypeOf(iot.GetDeviceMetricsRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "device-id",
				Short:      `Device ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			{
				Name:       "start-date",
				Short:      `Start date used to compute the best scale for the returned metrics`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.GetDeviceMetricsRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.GetDeviceMetrics(request)

		},
	}
}

func iotRouteList() *core.Command {
	return &core.Command{
		Short:     `List routes`,
		Long:      `List all routes in the specified region. By default, returned routes are ordered by creation date in ascending order, though this can be modified via the ` + "`" + `order_by` + "`" + ` field.`,
		Namespace: "iot",
		Resource:  "route",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.ListRoutesRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Ordering of requested routes`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc", "hub_id_asc", "hub_id_desc", "type_asc", "type_desc", "created_at_asc", "created_at_desc"},
			},
			{
				Name:       "hub-id",
				Short:      `Hub ID to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Route name to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.ListRoutesRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListRoutes(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Routes, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "Topic",
			},
			{
				FieldName: "Type",
			},
			{
				FieldName: "HubID",
			},
			{
				FieldName: "CreatedAt",
			},
		}},
	}
}

func iotRouteCreate() *core.Command {
	return &core.Command{
		Short: `Create a route`,
		Long: `Multiple kinds of routes can be created, such as:
- Database Route
  Create a route that will record subscribed MQTT messages into your database.
  <b>You need to manage the database by yourself</b>.
- REST Route.
  Create a route that will call a REST API on received subscribed MQTT messages.
- Amazon S3 Routes.
  Create a route that will put subscribed MQTT messages into an Object Storage bucket.
  You need to create the bucket yourself and grant write access.
  Granting can be done with s3cmd (` + "`" + `s3cmd setacl s3://<my-bucket> --acl-grant=write:555c69c3-87d0-4bf8-80f1-99a2f757d031:555c69c3-87d0-4bf8-80f1-99a2f757d031` + "`" + `).`,
		Namespace: "iot",
		Resource:  "route",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.CreateRouteRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Route name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("route"),
			},
			{
				Name:       "hub-id",
				Short:      `Hub ID of the route`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "topic",
				Short:      `Topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "s3-config.bucket-region",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "s3-config.bucket-name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "s3-config.object-prefix",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "s3-config.strategy",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "per_topic", "per_message"},
			},
			{
				Name:       "db-config.host",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.port",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.dbname",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.username",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.password",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.query",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.engine",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "postgresql", "mysql"},
			},
			{
				Name:       "rest-config.verb",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "get", "post", "put", "patch", "delete"},
			},
			{
				Name:       "rest-config.uri",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rest-config.headers.{key}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.CreateRouteRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.CreateRoute(request)

		},
	}
}

func iotRouteUpdate() *core.Command {
	return &core.Command{
		Short:     `Update a route`,
		Long:      `Update the parameters of an existing route, specified by its route ID.`,
		Namespace: "iot",
		Resource:  "route",
		Verb:      "update",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.UpdateRouteRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "route-id",
				Short:      `Route id`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "name",
				Short:      `Route name`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "topic",
				Short:      `Topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "s3-config.bucket-region",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "s3-config.bucket-name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "s3-config.object-prefix",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "s3-config.strategy",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "per_topic", "per_message"},
			},
			{
				Name:       "db-config.host",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.port",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.dbname",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.username",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.password",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.query",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "db-config.engine",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "postgresql", "mysql"},
			},
			{
				Name:       "rest-config.verb",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "get", "post", "put", "patch", "delete"},
			},
			{
				Name:       "rest-config.uri",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "rest-config.headers.{key}",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.UpdateRouteRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.UpdateRoute(request)

		},
	}
}

func iotRouteGet() *core.Command {
	return &core.Command{
		Short:     `Get a route`,
		Long:      `Get information for a particular route, specified by the route ID. The response returns full details of the route, including its type, the topic it subscribes to and its configuration.`,
		Namespace: "iot",
		Resource:  "route",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.GetRouteRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "route-id",
				Short:      `Route ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.GetRouteRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.GetRoute(request)

		},
	}
}

func iotRouteDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a route`,
		Long:      `Delete an existing route, specified by its route ID. Deleting a route is permanent, and cannot be undone.`,
		Namespace: "iot",
		Resource:  "route",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.DeleteRouteRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "route-id",
				Short:      `Route ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.DeleteRouteRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			e = api.DeleteRoute(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "route",
				Verb:     "delete",
			}, nil
		},
	}
}

func iotNetworkList() *core.Command {
	return &core.Command{
		Short:     `List the networks`,
		Long:      `List the networks.`,
		Namespace: "iot",
		Resource:  "network",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.ListNetworksRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Ordering of requested routes`,
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"name_asc", "name_desc", "type_asc", "type_desc", "created_at_asc", "created_at_desc"},
			},
			{
				Name:       "name",
				Short:      `Network name to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "hub-id",
				Short:      `Hub ID to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "topic-prefix",
				Short:      `Topic prefix to filter for`,
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.Region(core.AllLocalities)),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.ListNetworksRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			if request.Region == scw.Region(core.AllLocalities) {
				opts = append(opts, scw.WithRegions(api.Regions()...))
				request.Region = ""
			}
			resp, err := api.ListNetworks(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Networks, nil

		},
		View: &core.View{Fields: []*core.ViewField{
			{
				FieldName: "ID",
			},
			{
				FieldName: "Name",
			},
			{
				FieldName: "Type",
			},
			{
				FieldName: "Endpoint",
			},
			{
				FieldName: "HubID",
			},
			{
				FieldName: "CreatedAt",
			},
			{
				FieldName: "TopicPrefix",
			},
		}},
	}
}

func iotNetworkCreate() *core.Command {
	return &core.Command{
		Short:     `Create a new network`,
		Long:      `Create a new network for an existing hub.  Beside the default network, you can add networks for different data providers. Possible network types are Sigfox and REST.`,
		Namespace: "iot",
		Resource:  "network",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.CreateNetworkRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Short:      `Network name`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				Default:    core.RandomValueGenerator("network"),
			},
			{
				Name:       "type",
				Short:      `Type of network to connect with`,
				Required:   true,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown", "sigfox", "rest"},
			},
			{
				Name:       "hub-id",
				Short:      `Hub ID to connect the Network to`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "topic-prefix",
				Short:      `Topic prefix for the Network`,
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.CreateNetworkRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.CreateNetwork(request)

		},
	}
}

func iotNetworkGet() *core.Command {
	return &core.Command{
		Short:     `Retrieve a specific network`,
		Long:      `Retrieve an existing network, specified by its network ID.  The response returns full details of the network, including its type, the topic prefix and its endpoint.`,
		Namespace: "iot",
		Resource:  "network",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.GetNetworkRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "network-id",
				Short:      `Network ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.GetNetworkRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			return api.GetNetwork(request)

		},
	}
}

func iotNetworkDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a Network`,
		Long:      `Delete an existing network, specified by its network ID. Deleting a network is permanent, and cannot be undone.`,
		Namespace: "iot",
		Resource:  "network",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(iot.DeleteNetworkRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "network-id",
				Short:      `Network ID`,
				Required:   true,
				Deprecated: false,
				Positional: true,
			},
			core.RegionArgSpec(scw.RegionFrPar),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*iot.DeleteNetworkRequest)

			client := core.ExtractClient(ctx)
			api := iot.NewAPI(client)
			e = api.DeleteNetwork(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "network",
				Verb:     "delete",
			}, nil
		},
	}
}

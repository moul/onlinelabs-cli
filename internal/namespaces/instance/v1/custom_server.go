package instance

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/human"
	"github.com/scaleway/scaleway-cli/v2/internal/interactive"
	block "github.com/scaleway/scaleway-sdk-go/api/block/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/api/vpc/v1"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/validation"
)

const (
	serverActionTimeout = 10 * time.Minute
)

//
// Marshalers
//

// serverStateMarshalSpecs allows to override the displayed instance.ServerState.
var (
	serverStateMarshalSpecs = human.EnumMarshalSpecs{
		instance.ServerStateRunning:        &human.EnumMarshalSpec{Attribute: color.FgGreen},
		instance.ServerStateStopped:        &human.EnumMarshalSpec{Attribute: color.Faint, Value: "archived"},
		instance.ServerStateStoppedInPlace: &human.EnumMarshalSpec{Attribute: color.Faint},
		instance.ServerStateStarting:       &human.EnumMarshalSpec{Attribute: color.FgBlue},
		instance.ServerStateStopping:       &human.EnumMarshalSpec{Attribute: color.FgBlue},
		instance.ServerStateLocked:         &human.EnumMarshalSpec{Attribute: color.FgRed},
	}
)

// serverLocationMarshalerFunc marshals a instance.ServerLocation.
func serverLocationMarshalerFunc(i interface{}, _ *human.MarshalOpt) (string, error) {
	location := i.(instance.ServerLocation)
	zone, err := scw.ParseZone(location.ZoneID)
	if err != nil {
		return "", err
	}
	return string(zone), nil
}

// serversMarshalerFunc marshals a Server.
func serversMarshalerFunc(i interface{}, opt *human.MarshalOpt) (string, error) {
	// humanServerInList is the custom Server type used for list view.
	type humanServerInList struct {
		ID                string
		Name              string
		Type              string
		State             instance.ServerState
		Zone              scw.Zone
		PublicIP          net.IP
		PrivateIP         *string
		Tags              []string
		ImageName         string
		RoutedIPEnabled   *bool
		PlacementGroup    *instance.PlacementGroup
		ModificationDate  *time.Time
		CreationDate      *time.Time
		Volumes           int
		Protected         bool
		SecurityGroupName string
		SecurityGroupID   string
		StateDetail       string
		Arch              instance.Arch
		ImageID           string
	}

	servers := i.([]*instance.Server)
	humanServers := make([]*humanServerInList, 0)
	for _, server := range servers {
		publicIPAddress := net.IP(nil)
		if server.PublicIP != nil {
			publicIPAddress = server.PublicIP.Address
		}
		serverImageID := ""
		serverImageName := ""
		if server.Image != nil {
			serverImageID = server.Image.ID
			serverImageName = server.Image.Name
		}
		humanServers = append(humanServers, &humanServerInList{
			ID:                server.ID,
			Name:              server.Name,
			Type:              server.CommercialType,
			State:             server.State,
			Zone:              server.Zone,
			PublicIP:          publicIPAddress,
			PrivateIP:         server.PrivateIP,
			Tags:              server.Tags,
			ImageName:         serverImageName,
			RoutedIPEnabled:   server.RoutedIPEnabled, //nolint: staticcheck // Field is deprecated but still supported
			PlacementGroup:    server.PlacementGroup,
			ModificationDate:  server.ModificationDate,
			CreationDate:      server.CreationDate,
			Volumes:           len(server.Volumes),
			Protected:         server.Protected,
			SecurityGroupName: server.SecurityGroup.Name,
			SecurityGroupID:   server.SecurityGroup.ID,
			StateDetail:       server.StateDetail,
			Arch:              server.Arch,
			ImageID:           serverImageID,
		})
	}
	return human.Marshal(humanServers, opt)
}

// orderVolumes return an ordered slice based on the volume map key "0", "1", "2",...
func orderVolumes(v map[string]*instance.VolumeServer) []*instance.VolumeServer {
	indexes := []string(nil)
	for index := range v {
		indexes = append(indexes, index)
	}
	sort.Strings(indexes)

	orderedVolumes := make([]*instance.VolumeServer, 0, len(indexes))
	for _, index := range indexes {
		orderedVolumes = append(orderedVolumes, v[index])
	}
	return orderedVolumes
}

// serversMarshalerFunc marshals a BootscriptID.
func bootscriptMarshalerFunc(i interface{}, _ *human.MarshalOpt) (string, error) {
	bootscript := i.(instance.Bootscript)
	return bootscript.Title, nil
}

//
// Builders
//

func serverListBuilder(c *core.Command) *core.Command {
	type customListServersRequest struct {
		*instance.ListServersRequest
		OrganizationID *string
		ProjectID      *string
	}

	renameOrganizationIDArgSpec(c.ArgSpecs)
	renameProjectIDArgSpec(c.ArgSpecs)

	c.ArgsType = reflect.TypeOf(customListServersRequest{})

	c.AddInterceptors(func(ctx context.Context, argsI interface{}, runner core.CommandRunner) (i interface{}, err error) {
		args := argsI.(*customListServersRequest)

		if args.ListServersRequest == nil {
			args.ListServersRequest = &instance.ListServersRequest{}
		}

		request := args.ListServersRequest
		request.Organization = args.OrganizationID
		request.Project = args.ProjectID

		return runner(ctx, request)
	})
	return c
}

func serverUpdateBuilder(c *core.Command) *core.Command {
	type instanceUpdateServerRequestCustom struct {
		*instance.UpdateServerRequest
		IP               *instance.NullableStringValue
		PlacementGroupID *instance.NullableStringValue
		SecurityGroupID  *string
		VolumeIDs        *[]string
		CloudInit        string
	}

	c.ArgsType = reflect.TypeOf(instanceUpdateServerRequestCustom{})

	// Add completion functions
	c.ArgSpecs.GetByName("admin-password-encryption-ssh-key-id").AutoCompleteFunc = completeSSHKeyID

	// Rename modified arg specs.
	c.ArgSpecs.GetByName("placement-group").Name = "placement-group-id"
	c.ArgSpecs.GetByName("security-group.id").Name = "security-group-id"

	// Delete unused arg specs.
	c.ArgSpecs.DeleteByName("security-group.name")
	c.ArgSpecs.DeleteByName("volumes.{key}.name")
	c.ArgSpecs.DeleteByName("volumes.{key}.size")
	c.ArgSpecs.DeleteByName("volumes.{key}.id")
	c.ArgSpecs.DeleteByName("volumes.{key}.volume-type")

	// Add new arg specs.
	c.ArgSpecs.AddBefore("placement-group-id", &core.ArgSpec{
		Name:  "volume-ids.{index}",
		Short: "Will update ALL volume IDs at once, including the root volume of the server (use volume-ids=none to detach all volumes)",
	})
	c.ArgSpecs.AddBefore("boot-type", &core.ArgSpec{
		Name:  "ip",
		Short: `IP that should be attached to the server (use ip=none to detach)`,
	})
	c.ArgSpecs.AddBefore("boot-type", &core.ArgSpec{
		Name:        "cloud-init",
		Short:       "The cloud-init script to use",
		CanLoadFile: true,
	})

	c.Run = func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
		customRequest := argsI.(*instanceUpdateServerRequestCustom)

		updateServerRequest := customRequest.UpdateServerRequest
		updateServerRequest.PlacementGroup = customRequest.PlacementGroupID
		if customRequest.SecurityGroupID != nil {
			updateServerRequest.SecurityGroup = &instance.SecurityGroupTemplate{
				ID: *customRequest.SecurityGroupID,
			}
		}

		attachIPRequest := (*instance.UpdateIPRequest)(nil)

		detachIP := false

		client := core.ExtractClient(ctx)
		api := instance.NewAPI(client)

		getServerResponse, err := api.GetServer(&instance.GetServerRequest{
			Zone:     updateServerRequest.Zone,
			ServerID: customRequest.ServerID,
		})
		if err != nil {
			return nil, err
		}

		switch {
		case customRequest.IP == nil:
			// ip is not set
			// do nothing

		case customRequest.IP.Null:
			// ip=none
			// detach IP from the server, only if it was attached.
			if getServerResponse.Server.PublicIP != nil {
				detachIP = true
			}

		default:
			// ip=<anything>
			// update ip
			if getServerResponse.Server.PublicIP != nil {
				detachIP = true
			}
			attachIPRequest = &instance.UpdateIPRequest{
				IP: customRequest.IP.Value,
				Server: &instance.NullableStringValue{
					Value: customRequest.ServerID,
				},
			}
		}

		// Instance API does not support detaching the existing IP and then attaching a new one to the same server
		// in 1 call only.
		// We need to do it manually in 2 calls.

		if detachIP {
			_, err = api.UpdateIP(&instance.UpdateIPRequest{
				IP: getServerResponse.Server.PublicIP.ID,
				Server: &instance.NullableStringValue{
					Null: true,
				},
			})
			if err != nil {
				return nil, err
			}
		}

		if attachIPRequest != nil {
			_, err = api.UpdateIP(attachIPRequest)
			if err != nil {
				return nil, err
			}
		}

		// Update all volume IDs at once.
		if customRequest.VolumeIDs != nil {
			volumes := make(map[string]*instance.VolumeServerTemplate)
			for i, volumeID := range *customRequest.VolumeIDs {
				index := strconv.Itoa(i)
				volumes[index] = &instance.VolumeServerTemplate{
					ID:   scw.StringPtr(volumeID),
					Name: scw.StringPtr(getServerResponse.Server.Name + "-" + index),
				}
			}
			customRequest.Volumes = &volumes
		}

		// Set cloud-init
		if customRequest.CloudInit != "" {
			err := api.SetServerUserData(&instance.SetServerUserDataRequest{
				Zone:     updateServerRequest.Zone,
				ServerID: customRequest.ServerID,
				Key:      "cloud-init",
				Content:  bytes.NewBufferString(customRequest.CloudInit),
			})
			if err != nil {
				return nil, err
			}
		}

		updateServerResponse, err := api.UpdateServer(updateServerRequest)
		if err != nil {
			return nil, err
		}

		return updateServerResponse, nil
	}

	return c
}

func serverGetBuilder(c *core.Command) *core.Command {
	// This method is here as a proof of concept before we find the correct way to implement it at larger scale
	c.ArgSpecs.GetPositionalArg().AutoCompleteFunc = func(ctx context.Context, prefix string, request any) core.AutocompleteSuggestions {
		req := request.(*instance.GetServerRequest)
		api := instance.NewAPI(core.ExtractClient(ctx))
		resp, err := api.ListServers(&instance.ListServersRequest{
			Zone: req.Zone,
		}, scw.WithAllPages())
		if err != nil {
			return nil
		}

		suggestion := core.AutocompleteSuggestions{}
		for _, s := range resp.Servers {
			if strings.HasPrefix(s.ID, prefix) {
				suggestion = append(suggestion, s.ID)
			}
		}
		return suggestion
	}

	c.Interceptor = func(ctx context.Context, argsI interface{}, runner core.CommandRunner) (interface{}, error) {
		rawResp, err := runner(ctx, argsI)
		if err != nil {
			return rawResp, err
		}
		getServerResp := rawResp.(*instance.GetServerResponse)

		client := core.ExtractClient(ctx)
		vpcAPI := vpc.NewAPI(client)

		type customNICs struct {
			ID                 string `json:"id"`
			MacAddress         string `json:"mac_address"`
			PrivateNetworkName string `json:"private_network_name"`
			PrivateNetworkID   string `json:"private_network_id"`
		}

		nics := []customNICs{}

		for _, nic := range getServerResp.Server.PrivateNics {
			pn, err := vpcAPI.GetPrivateNetwork(&vpc.GetPrivateNetworkRequest{
				PrivateNetworkID: nic.PrivateNetworkID,
				Zone:             getServerResp.Server.Zone,
			})
			if err != nil {
				return nil, err
			}
			nics = append(nics, customNICs{
				ID:                 nic.ID,
				PrivateNetworkID:   pn.ID,
				PrivateNetworkName: pn.Name,
				MacAddress:         nic.MacAddress,
			})
		}

		return &struct {
			*instance.Server
			Volumes     []*instance.VolumeServer
			PrivateNics []customNICs `json:"private_nics"`
		}{
			getServerResp.Server,
			orderVolumes(getServerResp.Server.Volumes),
			nics,
		}, nil
	}

	c.View = &core.View{
		Sections: []*core.ViewSection{
			{
				FieldName: "Image",
				Title:     "Server Image",
			},
			{
				FieldName: "AllowedActions",
				Title:     "Allowed Actions",
			},
			{
				FieldName: "Volumes",
				Title:     "Volumes",
			},
			{
				Title:     "Public IPs",
				FieldName: "PublicIPs",
			},
			{
				Title:     "IPv6",
				FieldName: "IPv6",
			},
			{
				FieldName: "PrivateNics",
				Title:     "Private NICs",
			},
		},
	}

	return c
}

//
// Commands
//

func serverAttachVolumeCommand() *core.Command {
	return &core.Command{
		Short:     `Attach a volume to a server`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "attach-volume",
		ArgsType:  reflect.TypeOf(instance.AttachVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "server-id",
				Short:    `ID of the server`,
				Required: true,
			},
			{
				Name:     "volume-id",
				Short:    `ID of the volume to attach`,
				Required: true,
			},
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, err error) {
			request := argsI.(*instance.AttachVolumeRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.AttachVolume(request)
		},
		Examples: []*core.Example{
			{
				Short:    "Attach a volume to a server",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111","volume_id": "22222222-1111-5555-2222-666666111111"}`,
			},
		},
	}
}

func serverDetachVolumeCommand() *core.Command {
	return &core.Command{
		Short:     `Detach a volume from its server`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "detach-volume",
		ArgsType:  reflect.TypeOf(instance.DetachVolumeRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "volume-id",
				Short:    `ID of the volume to detach`,
				Required: true,
			},
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, err error) {
			request := argsI.(*instance.DetachVolumeRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			return api.DetachVolume(request)
		},
		Examples: []*core.Example{
			{
				Short:    "Detach a volume from its server",
				ArgsJSON: `{"volume_id": "22222222-1111-5555-2222-666666111111"}`,
			},
		},
	}
}

func serverAttachIPCommand() *core.Command {
	type customIPAttachRequest struct {
		OrganizationID *string
		ProjectID      *string
		// Server: UUID of the server you want to attach the IP to
		ServerID string   `json:"server,omitempty"`
		IP       string   `json:"-"`
		Zone     scw.Zone `json:"zone"`
	}

	return &core.Command{
		Short:     `Attach an IP to a server`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "attach-ip",
		ArgsType:  reflect.TypeOf(customIPAttachRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server`,
				Required:   true,
				Positional: true,
			},
			{
				Name:     "ip",
				Short:    `UUID of the IP to attach or its UUID`,
				Required: true,
			},
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, err error) {
			api := instance.NewAPI(core.ExtractClient(ctx))
			args := argsI.(*customIPAttachRequest)

			var ipID string
			switch {
			case validation.IsUUID(args.IP):
				ipID = args.IP
			case net.ParseIP(args.IP) != nil:
				// Find the corresponding flexible IP UUID.
				logger.Debugf("finding public IP UUID from address: %s", args.IP)
				res, err := api.GetIP(&instance.GetIPRequest{
					Zone: args.Zone,
					IP:   args.IP,
				})
				if err != nil { // FIXME: isNotFoundError
					return nil, fmt.Errorf("%s does not belong to you", args.IP)
				}
				ipID = res.IP.ID
			default:
				return nil, fmt.Errorf(`invalid IP "%s", should be either an IP address ID or a reserved flexible IP address`, args.IP)
			}

			_, err = api.UpdateIP(&instance.UpdateIPRequest{
				IP: ipID,
				Server: &instance.NullableStringValue{
					Value: args.ServerID,
				},
				Zone: args.Zone,
			})
			if err != nil {
				return nil, err
			}
			return api.GetServer(&instance.GetServerRequest{ServerID: args.ServerID})
		},
		Examples: []*core.Example{
			{
				Short:    "Attach an IP to a server",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111","ip": "11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Attach an IP to a server",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111","ip": "1.2.3.4"}`,
			},
		},
	}
}

func serverDetachIPCommand() *core.Command {
	type customIPDetachRequest struct {
		OrganizationID *string
		ProjectID      *string
		Zone           scw.Zone `json:"zone"`
		ServerID       string
	}

	return &core.Command{
		Short:     `Detach an IP from a server`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "detach-ip",
		ArgsType:  reflect.TypeOf(customIPDetachRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `UUID of the server.`,
				Required:   true,
				Positional: true,
			},
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, err error) {
			args := argsI.(*customIPDetachRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			serverResponse, err := api.GetServer(&instance.GetServerRequest{ServerID: args.ServerID})
			if err != nil {
				return nil, err
			}

			if server := serverResponse.Server; server != nil {
				if ip := server.PublicIP; ip != nil {
					_, err := api.UpdateIP(&instance.UpdateIPRequest{
						Zone: args.Zone,
						// We detach an ip by specifying no serverResponse
						Server: &instance.NullableStringValue{
							Null: true,
						},
						IP: ip.ID,
					})
					if err != nil {
						return nil, err
					}
					return api.GetServer(&instance.GetServerRequest{ServerID: args.ServerID})
				}
				return nil, errors.New("no public ip found")
			}
			return nil, errors.New("no server found")
		},
		Examples: []*core.Example{
			{
				Short:    "Detach IP from a given server",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

type serverWaitRequest struct {
	Zone     scw.Zone
	ServerID string
	Timeout  time.Duration
}

func serverWaitCommand() *core.Command {
	return &core.Command{
		Short:     `Wait for server to reach a stable state`,
		Long:      `Wait for server to reach a stable state. This is similar to using --wait flag on other action commands, but without requiring a new action on the server.`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "wait",
		Groups:    []string{"workflow"},
		ArgsType:  reflect.TypeOf(serverWaitRequest{}),
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, err error) {
			args := argsI.(*serverWaitRequest)

			return instance.NewAPI(core.ExtractClient(ctx)).WaitForServer(&instance.WaitForServerRequest{
				Zone:          args.Zone,
				ServerID:      args.ServerID,
				Timeout:       scw.TimeDurationPtr(args.Timeout),
				RetryInterval: core.DefaultRetryInterval,
			})
		},
		ArgSpecs: core.ArgSpecs{
			core.WaitTimeoutArgSpec(serverActionTimeout),
			{
				Name:       "server-id",
				Short:      `ID of the server affected by the action.`,
				Required:   true,
				Positional: true,
			},
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
		},
		Examples: []*core.Example{
			{
				Short:    "Wait for a server to reach a stable state",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

type customDeleteServerRequest struct {
	Zone          scw.Zone
	ServerID      string
	WithVolumes   withVolumes
	WithIP        bool
	ForceShutdown bool
}

type withVolumes string

const (
	withVolumesNone  = withVolumes("none")
	withVolumesLocal = withVolumes("local")
	withVolumesBlock = withVolumes("block")
	withVolumesRoot  = withVolumes("root")
	withVolumesAll   = withVolumes("all")
)

func serverDeleteCommand() *core.Command {
	return &core.Command{
		Short:     `Delete server`,
		Long:      `Delete a server with the given ID.`,
		Namespace: "instance",
		Verb:      "delete",
		Resource:  "server",
		ArgsType:  reflect.TypeOf(customDeleteServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Required:   true,
				Positional: true,
			},
			{
				Name:    "with-volumes",
				Short:   "Delete the volumes attached to the server",
				Default: core.DefaultValueSetter("all"),
				EnumValues: []string{
					string(withVolumesNone),
					string(withVolumesLocal),
					string(withVolumesBlock),
					string(withVolumesRoot),
					string(withVolumesAll),
				},
			},
			{
				Name:  "with-ip",
				Short: "Delete the IP attached to the server",
			},
			{
				Name:  "force-shutdown",
				Short: "Force shutdown of the instance server before deleting it",
			},
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
		},
		Examples: []*core.Example{
			{
				Short:    "Delete a server in the default zone with a given id",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Delete a server in fr-par-1 zone with a given id",
				ArgsJSON: `{"zone":"fr-par-1", "server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw instance server terminate",
				Short:   "Terminate a running server",
			},
			{
				Command: "scw instance server stop",
				Short:   "Stop a running server",
			},
		},
		WaitUsage: "wait until the server and its resources are deleted",
		WaitFunc: func(ctx context.Context, _, respI interface{}) (interface{}, error) {
			server := respI.(*core.SuccessResult).TargetResource.(*instance.Server)
			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)

			notFoundErr := &scw.ResourceNotFoundError{}

			_, err := api.WaitForServer(&instance.WaitForServerRequest{
				Zone:          server.Zone,
				ServerID:      server.ID,
				Timeout:       scw.TimeDurationPtr(serverActionTimeout),
				RetryInterval: core.DefaultRetryInterval,
			})
			if err != nil {
				err = errors.Unwrap(err)
				if !errors.As(err, &notFoundErr) {
					return nil, err
				}
			}
			return respI, nil
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			deleteServerArgs := argsI.(*customDeleteServerRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)

			server, err := api.GetServer(&instance.GetServerRequest{
				Zone:     deleteServerArgs.Zone,
				ServerID: deleteServerArgs.ServerID,
			})
			if err != nil {
				return nil, err
			}

			if deleteServerArgs.ForceShutdown {
				finalStateServer, err := api.WaitForServer(&instance.WaitForServerRequest{
					Zone:          deleteServerArgs.Zone,
					ServerID:      deleteServerArgs.ServerID,
					Timeout:       scw.TimeDurationPtr(serverActionTimeout),
					RetryInterval: core.DefaultRetryInterval,
				})
				if err != nil {
					return nil, err
				}

				if finalStateServer.State != instance.ServerStateStopped {
					err = api.ServerActionAndWait(&instance.ServerActionAndWaitRequest{
						Zone:          deleteServerArgs.Zone,
						ServerID:      deleteServerArgs.ServerID,
						Action:        instance.ServerActionPoweroff,
						Timeout:       scw.TimeDurationPtr(serverActionTimeout),
						RetryInterval: core.DefaultRetryInterval,
					})
					if err != nil {
						return nil, err
					}
				}
			}

			err = api.DeleteServer(&instance.DeleteServerRequest{
				Zone:     deleteServerArgs.Zone,
				ServerID: deleteServerArgs.ServerID,
			})
			if err != nil {
				return nil, err
			}

			if deleteServerArgs.WithIP && server.Server.PublicIPs != nil {
				for _, ip := range server.Server.PublicIPs {
					if ip.Dynamic {
						continue
					}
					err = api.DeleteIP(&instance.DeleteIPRequest{
						Zone: deleteServerArgs.Zone,
						IP:   ip.ID,
					})
					if err != nil {
						return nil, err
					}
					_, _ = interactive.Printf("successfully deleted ip %s\n", ip.Address.String())
				}
			}

			deletedVolumeMessages := [][2]string(nil)
		volumeDelete:
			for index, volume := range server.Server.Volumes {
				switch {
				case deleteServerArgs.WithVolumes == withVolumesNone:
					break volumeDelete
				case deleteServerArgs.WithVolumes == withVolumesRoot && index != "0":
					continue
				case deleteServerArgs.WithVolumes == withVolumesLocal && volume.VolumeType != instance.VolumeServerVolumeTypeLSSD:
					continue
				case deleteServerArgs.WithVolumes == withVolumesBlock && volume.VolumeType != instance.VolumeServerVolumeTypeBSSD && volume.VolumeType != instance.VolumeServerVolumeTypeSbsVolume:
					continue
				case volume.VolumeType == instance.VolumeServerVolumeTypeScratch:
					continue
				}
				if volume.VolumeType == instance.VolumeServerVolumeTypeSbsVolume {
					err = block.NewAPI(client).DeleteVolume(&block.DeleteVolumeRequest{
						Zone:     deleteServerArgs.Zone,
						VolumeID: volume.ID,
					})
				} else {
					err = api.DeleteVolume(&instance.DeleteVolumeRequest{
						Zone:     deleteServerArgs.Zone,
						VolumeID: volume.ID,
					})
				}
				if err != nil {
					return nil, &core.CliError{
						Err:  err,
						Hint: "Make sure this resource have been deleted or try to delete it manually.",
					}
				}
				humanSize, err := human.Marshal(volume.Size, nil)
				if err != nil {
					logger.Debugf("cannot marshal human size %v", volume.Size)
				}
				deletedVolumeMessages = append(deletedVolumeMessages, [2]string{
					index,
					fmt.Sprintf("successfully deleted volume %s (%s %s)", volume.Name, humanSize, volume.VolumeType),
				})
			}

			// Sort and print deleted volume messages
			sort.Slice(deletedVolumeMessages, func(i, j int) bool {
				return deletedVolumeMessages[i][0] < deletedVolumeMessages[j][0]
			})
			for _, message := range deletedVolumeMessages {
				_, _ = interactive.Println(message[1])
			}

			return &core.SuccessResult{
				TargetResource: server.Server,
			}, nil
		},
	}
}

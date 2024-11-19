package instance

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/interactive"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func serverStartCommand() *core.Command {
	return &core.Command{
		Short:     `Power on server`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "start",
		ArgsType:  reflect.TypeOf(instanceUniqueActionRequest{}),
		Run:       getRunServerAction(instance.ServerActionPoweron),
		WaitFunc:  waitForServerFunc(),
		ArgSpecs:  serverActionArgSpecs,
		Examples: []*core.Example{
			{
				Short:    "Start a server in the default zone with a given id",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Start a server in fr-par-1 zone with a given id",
				ArgsJSON: `{"zone":"fr-par-1", "server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func serverStopCommand() *core.Command {
	return &core.Command{
		Short:     `Power off server`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "stop",
		ArgsType:  reflect.TypeOf(instanceUniqueActionRequest{}),
		Run:       getRunServerAction(instance.ServerActionPoweroff),
		WaitFunc:  waitForServerFunc(),
		ArgSpecs:  serverActionArgSpecs,
		Examples: []*core.Example{
			{
				Short:    "Stop a server in the default zone with a given id",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Stop a server in fr-par-1 zone with a given id",
				ArgsJSON: `{"zone":"fr-par-1", "server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func serverStandbyCommand() *core.Command {
	return &core.Command{
		Short:     `Put server in standby mode`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "standby",
		ArgsType:  reflect.TypeOf(instanceUniqueActionRequest{}),
		Run:       getRunServerAction(instance.ServerActionStopInPlace),
		WaitFunc:  waitForServerFunc(),
		ArgSpecs:  serverActionArgSpecs,
		Examples: []*core.Example{
			{
				Short:    "Put in standby a server in the default zone with a given id",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Put in standby a server in fr-par-1 zone with a given id",
				ArgsJSON: `{"zone":"fr-par-1", "server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func serverRebootCommand() *core.Command {
	return &core.Command{
		Short:     `Reboot server`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "reboot",
		ArgsType:  reflect.TypeOf(instanceUniqueActionRequest{}),
		Run:       getRunServerAction(instance.ServerActionReboot),
		WaitFunc:  waitForServerFunc(),
		ArgSpecs:  serverActionArgSpecs,
		Examples: []*core.Example{
			{
				Short:    "Reboot a server in the default zone with a given id",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Reboot a server in fr-par-1 zone with a given id",
				ArgsJSON: `{"zone":"fr-par-1", "server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func serverEnableRoutedIPCommand() *core.Command {
	return &core.Command{
		Short: `Migrate server to IP mobility`,
		Long: `Enable routed IP for this server and migrate the nat public IP to routed
Server will reboot !
https://www.scaleway.com/en/docs/compute/instances/api-cli/using-ip-mobility/
`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "enable-routed-ip",
		ArgsType:  reflect.TypeOf(instanceUniqueActionRequest{}),
		Run:       getRunServerAction("enable_routed_ip"),
		WaitFunc:  waitForServerFunc(),
		ArgSpecs:  serverActionArgSpecs,
		Examples: []*core.Example{
			{
				Short:    "Migrate a server with legacy network to IP mobility",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

func serverBackupCommand() *core.Command {
	type instanceBackupRequest struct {
		Zone     scw.Zone
		ServerID string
		Name     string
		Unified  bool
	}

	return &core.Command{
		Short: `Backup server`,
		Long: `Create a new image based on the server.

This command:
  - creates a snapshot of all attached volumes.
  - creates an image based on all these snapshots.

Once your image is ready you will be able to create a new server based on this image.
`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "backup",
		ArgsType:  reflect.TypeOf(instanceBackupRequest{}),
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, err error) {
			args := argsI.(*instanceBackupRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)
			server, err := api.GetServer(&instance.GetServerRequest{
				ServerID: args.ServerID,
				Zone:     args.Zone,
			})
			if err != nil {
				return nil, err
			}

			req := &instance.ServerActionRequest{
				Zone:     args.Zone,
				ServerID: args.ServerID,
				Action:   instance.ServerActionBackup,
				Name:     &args.Name,
				Volumes:  map[string]*instance.ServerActionRequestVolumeBackupTemplate{},
			}
			for _, v := range server.Server.Volumes {
				var template *instance.ServerActionRequestVolumeBackupTemplate
				if args.Unified {
					template = &instance.ServerActionRequestVolumeBackupTemplate{
						VolumeType: instance.SnapshotVolumeTypeUnified,
					}
				} else {
					if v.VolumeType == instance.VolumeServerVolumeTypeSbsVolume {
						template = &instance.ServerActionRequestVolumeBackupTemplate{
							VolumeType: instance.SnapshotVolumeType("sbs_snapshot"),
						}
					} else {
						template = &instance.ServerActionRequestVolumeBackupTemplate{
							VolumeType: instance.SnapshotVolumeType(v.VolumeType),
						}
					}
				}
				req.Volumes[v.ID] = template
			}
			res, err := api.ServerAction(req)
			if err != nil {
				return nil, err
			}

			tmp := strings.Split(res.Task.HrefResult, "/")
			if len(tmp) != 3 {
				return nil, errors.New("cannot extract image id from task")
			}
			return api.GetImage(&instance.GetImageRequest{Zone: args.Zone, ImageID: tmp[2]})
		},
		WaitFunc: func(ctx context.Context, _, respI interface{}) (i interface{}, err error) {
			resp := respI.(*instance.GetImageResponse)
			api := instance.NewAPI(core.ExtractClient(ctx))
			return api.WaitForImage(&instance.WaitForImageRequest{
				ImageID:       resp.Image.ID,
				Zone:          resp.Image.Zone,
				Timeout:       scw.TimeDurationPtr(serverActionTimeout),
				RetryInterval: core.DefaultRetryInterval,
			})
		},
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Short:      `ID of the server to backup.`,
				Required:   true,
				Positional: true,
			},
			{
				Name:    "name",
				Short:   `Name of your backup.`,
				Default: core.RandomValueGenerator("backup"),
			},
			{
				Name:  "unified",
				Short: "Whether or not the type of the snapshot is unified.",
			},
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
		},
		Examples: []*core.Example{
			{
				Short:    "Create a new image based on a server",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
	}
}

type customTerminateServerRequest struct {
	Zone      scw.Zone
	ServerID  string
	WithIP    bool
	WithBlock withBlock
}

type withBlock string

const (
	withBlockPrompt = withBlock("prompt")
	withBlockTrue   = withBlock("true")
	withBlockFalse  = withBlock("false")
)

func serverTerminateCommand() *core.Command {
	return &core.Command{
		Short:     `Terminate server`,
		Long:      `Terminates a server with the given ID and all of its volumes.`,
		Namespace: "instance",
		Verb:      "terminate",
		Resource:  "server",
		ArgsType:  reflect.TypeOf(customTerminateServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "server-id",
				Required:   true,
				Positional: true,
			},
			{
				Name:  "with-ip",
				Short: "Delete the IP attached to the server",
			},
			{
				Name:    "with-block",
				Short:   "Delete the Block Storage volumes attached to the server",
				Default: core.DefaultValueSetter("prompt"),
				EnumValues: []string{
					string(withBlockPrompt),
					string(withBlockTrue),
					string(withBlockFalse),
				},
			},
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
		},
		Examples: []*core.Example{
			{
				Short:    "Terminate a server in the default zone with a given id",
				ArgsJSON: `{"server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Terminate a server in fr-par-1 zone with a given id",
				ArgsJSON: `{"zone":"fr-par-1", "server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
			{
				Short:    "Terminate a server and also delete its flexible IPs",
				ArgsJSON: `{"with_ip":true, "server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw instance server delete",
				Short:   "delete a running server",
			},
			{
				Command: "scw instance server stop",
				Short:   "Stop a running server",
			},
		},
		WaitUsage: "wait until the server and its resources are deleted",
		WaitFunc: func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
			terminateServerArgs := argsI.(*customTerminateServerRequest)
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

			if terminateServerArgs.WithBlock == withBlockTrue {
				for _, volume := range server.Volumes {
					if volume.VolumeType != instance.VolumeServerVolumeTypeBSSD {
						continue
					}
					_, err := api.WaitForVolume(&instance.WaitForVolumeRequest{
						VolumeID: volume.ID,
						Zone:     volume.Zone,
					})
					if err != nil {
						if errors.As(err, &notFoundErr) {
							continue
						}
						return nil, err
					}
				}
			}

			return respI, nil
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			terminateServerArgs := argsI.(*customTerminateServerRequest)

			client := core.ExtractClient(ctx)
			api := instance.NewAPI(client)

			server, err := api.GetServer(&instance.GetServerRequest{
				Zone:     terminateServerArgs.Zone,
				ServerID: terminateServerArgs.ServerID,
			})
			if err != nil {
				return nil, err
			}

			deleteBlockVolumes, err := shouldDeleteBlockVolumes(ctx, server, terminateServerArgs.WithBlock)
			if err != nil {
				return nil, err
			}

			if !deleteBlockVolumes {
				// detach block storage volumes before terminating the instance to preserve them
				for _, volume := range server.Server.Volumes {
					if volume.VolumeType != instance.VolumeServerVolumeTypeBSSD {
						continue
					}

					if _, err := api.DetachVolume(&instance.DetachVolumeRequest{
						Zone:     terminateServerArgs.Zone,
						VolumeID: volume.ID,
					}); err != nil {
						return nil, err
					}

					_, _ = interactive.Printf("successfully detached volume %s\n", volume.Name)
				}
			}

			if _, err := api.ServerAction(&instance.ServerActionRequest{
				Zone:     terminateServerArgs.Zone,
				ServerID: terminateServerArgs.ServerID,
				Action:   instance.ServerActionTerminate,
			}); err != nil {
				return nil, err
			}

			if terminateServerArgs.WithIP && server.Server.PublicIP != nil && !server.Server.PublicIP.Dynamic {
				err = api.DeleteIP(&instance.DeleteIPRequest{
					Zone: terminateServerArgs.Zone,
					IP:   server.Server.PublicIP.ID,
				})
				if err != nil {
					return nil, err
				}
				_, _ = interactive.Printf("successfully deleted ip %s\n", server.Server.PublicIP.Address.String())
			}

			return &core.SuccessResult{
				TargetResource: server.Server,
			}, err
		},
	}
}

func shouldDeleteBlockVolumes(ctx context.Context, server *instance.GetServerResponse, terminateWithBlock withBlock) (bool, error) {
	switch terminateWithBlock {
	case withBlockTrue:
		return true, nil
	case withBlockFalse:
		return false, nil
	case withBlockPrompt:
		// Only prompt user if at least one block volume is attached to the instance
		for _, volume := range server.Server.Volumes {
			if volume.VolumeType != instance.VolumeServerVolumeTypeBSSD {
				continue
			}

			return interactive.PromptBoolWithConfig(&interactive.PromptBoolConfig{
				Prompt:       "Do you also want to delete block volumes attached to this instance ?",
				DefaultValue: false,
				Ctx:          ctx,
			})
		}
		return false, nil
	default:
		return false, fmt.Errorf("unsupported with-block value %v", terminateWithBlock)
	}
}

type instanceUniqueActionRequest struct {
	Zone     scw.Zone
	ServerID string
}

var serverActionArgSpecs = core.ArgSpecs{
	{
		Name:       "server-id",
		Short:      `ID of the server affected by the action.`,
		Required:   true,
		Positional: true,
	},
	core.ZoneArgSpec((*instance.API)(nil).Zones()...),
}

func getRunServerAction(action instance.ServerAction) core.CommandRunner {
	return func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
		args := argsI.(*instanceUniqueActionRequest)

		client := core.ExtractClient(ctx)
		api := instance.NewAPI(client)

		_, err := api.ServerAction(&instance.ServerActionRequest{
			Zone:     args.Zone,
			ServerID: args.ServerID,
			Action:   action,
		})
		return &core.SuccessResult{Message: fmt.Sprintf("%s successfully started for the server", action)}, err
	}
}

func waitForServerFunc() core.WaitFunc {
	return func(ctx context.Context, argsI, _ interface{}) (interface{}, error) {
		return instance.NewAPI(core.ExtractClient(ctx)).WaitForServer(&instance.WaitForServerRequest{
			Zone:          argsI.(*instanceUniqueActionRequest).Zone,
			ServerID:      argsI.(*instanceUniqueActionRequest).ServerID,
			Timeout:       scw.TimeDurationPtr(serverActionTimeout),
			RetryInterval: core.DefaultRetryInterval,
		})
	}
}

type instanceActionRequest struct {
	Zone     scw.Zone
	ServerID string
	Action   string
}

func serverActionCommand() *core.Command {
	argSpecs := serverActionArgSpecs
	argSpecs.AddBefore("server-id", &core.ArgSpec{
		Name:     "action",
		Short:    "The raw API action to perform, as listed with 'scw instance server list-actions'",
		Required: true,
	})

	return &core.Command{
		Short:     `Perform a raw API action on a server`,
		Namespace: "instance",
		Resource:  "server",
		Verb:      "action",
		ArgsType:  reflect.TypeOf(instanceActionRequest{}),
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*instanceActionRequest)

			return getRunServerAction(instance.ServerAction(args.Action))(ctx, &instanceUniqueActionRequest{
				Zone:     args.Zone,
				ServerID: args.ServerID,
			})
		},
		WaitFunc: func(ctx context.Context, argsI, _ interface{}) (interface{}, error) {
			return instance.NewAPI(core.ExtractClient(ctx)).WaitForServer(&instance.WaitForServerRequest{
				Zone:          argsI.(*instanceActionRequest).Zone,
				ServerID:      argsI.(*instanceActionRequest).ServerID,
				Timeout:       scw.TimeDurationPtr(serverActionTimeout),
				RetryInterval: core.DefaultRetryInterval,
			})
		},
		ArgSpecs: argSpecs,
		Examples: []*core.Example{
			{
				Short:    "Start a server in the default zone with a given id",
				ArgsJSON: `{"action":"poweron", "server_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw instance server list-actions",
				Short:   "List available actions for a server",
			},
			{
				Command: "scw instance server reboot",
				Short:   "Perform reboot action",
			},
			{
				Command: "scw instance server start",
				Short:   "Perform start action",
			},
			{
				Command: "scw instance server stop",
				Short:   "Perform stop action",
			},
		},
	}
}

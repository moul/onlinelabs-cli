package baremetal_test

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/baremetal/v1"
	baremetalSDK "github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Test_StartServerErrors(t *testing.T) {
	t.Run("Error: cannot be started while not delivered", core.Test(&core.TestConfig{
		BeforeFunc: createServer("Server"),
		Commands:   baremetal.GetCommands(),
		Cmd:        "scw baremetal server start zone=nl-ams-1 {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(1),
		),
		AfterFunc: core.AfterFuncCombine(
			func(ctx *core.AfterFuncCtx) error {
				api := baremetalSDK.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*baremetalSDK.Server)
				_, err := api.WaitForServer(&baremetalSDK.WaitForServerRequest{
					ServerID:      server.ID,
					Zone:          server.Zone,
					Timeout:       scw.TimeDurationPtr(baremetal.ServerActionTimeout),
					RetryInterval: core.DefaultRetryInterval,
				})
				return err
			},
			deleteServer("Server"),
		),
	}))
}

func Test_StopServerErrors(t *testing.T) {
	t.Run("Error: cannot be stopped while not delivered", core.Test(&core.TestConfig{
		BeforeFunc: createServer("Server"),
		Commands:   baremetal.GetCommands(),
		Cmd:        "scw baremetal server stop zone=nl-ams-1 {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(1),
		),
		AfterFunc: core.AfterFuncCombine(
			func(ctx *core.AfterFuncCtx) error {
				api := baremetalSDK.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*baremetalSDK.Server)
				_, err := api.WaitForServer(&baremetalSDK.WaitForServerRequest{
					ServerID:      server.ID,
					Zone:          server.Zone,
					Timeout:       scw.TimeDurationPtr(baremetal.ServerActionTimeout),
					RetryInterval: core.DefaultRetryInterval,
				})
				return err
			},
			deleteServer("Server"),
		),
	}))
}

func Test_RebootServerErrors(t *testing.T) {
	t.Run("Error: cannot be rebooted while not delivered", core.Test(&core.TestConfig{
		BeforeFunc: createServer("Server"),
		Commands:   baremetal.GetCommands(),
		Cmd:        "scw baremetal server reboot zone-nl-ams-1 {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(1),
		),
		AfterFunc: core.AfterFuncCombine(
			func(ctx *core.AfterFuncCtx) error {
				api := baremetalSDK.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*baremetalSDK.Server)
				_, err := api.WaitForServer(&baremetalSDK.WaitForServerRequest{
					ServerID:      server.ID,
					Zone:          server.Zone,
					Timeout:       scw.TimeDurationPtr(baremetal.ServerActionTimeout),
					RetryInterval: core.DefaultRetryInterval,
				})
				return err
			},
			deleteServer("Server"),
		),
	}))
}

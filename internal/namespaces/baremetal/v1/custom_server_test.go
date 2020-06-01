package baremetal

import (
	"testing"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Test_StartServerErrors(t *testing.T) {
	t.Run("Error: cannot be started while not delivered", core.Test(&core.TestConfig{
		BeforeFunc: createServer("Server"),
		Commands:   GetCommands(),
		Cmd:        "scw baremetal server start {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(1),
		),
		AfterFunc: core.AfterFuncCombine(
			func(ctx *core.AfterFuncCtx) error {
				api := baremetal.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*baremetal.Server)
				server, err := api.WaitForServer(&baremetal.WaitForServerRequest{
					ServerID:      server.ID,
					Zone:          server.Zone,
					Timeout:       serverActionTimeout,
					RetryInterval: defaultRetryInterval,
				})
				return err
			},
			deleteServer("Server"),
		),
		DefaultZone: scw.ZoneFrPar2,
	}))
}

func Test_StopServerErrors(t *testing.T) {
	t.Run("Error: cannot be stopped while not delivered", core.Test(&core.TestConfig{
		BeforeFunc: createServer("Server"),
		Commands:   GetCommands(),
		Cmd:        "scw baremetal server stop {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(1),
		),
		AfterFunc: core.AfterFuncCombine(
			func(ctx *core.AfterFuncCtx) error {
				api := baremetal.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*baremetal.Server)
				server, err := api.WaitForServer(&baremetal.WaitForServerRequest{
					ServerID:      server.ID,
					Zone:          server.Zone,
					Timeout:       serverActionTimeout,
					RetryInterval: defaultRetryInterval,
				})
				return err
			},
			deleteServer("Server"),
		),
		DefaultZone: scw.ZoneFrPar2,
	}))
}

func Test_RebootServerErrors(t *testing.T) {
	t.Run("Error: cannot be rebooted while not delivered", core.Test(&core.TestConfig{
		BeforeFunc: createServer("Server"),
		Commands:   GetCommands(),
		Cmd:        "scw baremetal server reboot {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(1),
		),
		AfterFunc: core.AfterFuncCombine(
			func(ctx *core.AfterFuncCtx) error {
				api := baremetal.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*baremetal.Server)
				server, err := api.WaitForServer(&baremetal.WaitForServerRequest{
					ServerID:      server.ID,
					Zone:          server.Zone,
					Timeout:       serverActionTimeout,
					RetryInterval: defaultRetryInterval,
				})
				return err
			},
			deleteServer("Server"),
		),
		DefaultZone: scw.ZoneFrPar2,
	}))
}

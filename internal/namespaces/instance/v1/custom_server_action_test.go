package instance_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/interactive"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/instance/v1"
	instanceSDK "github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/stretchr/testify/require"
)

// These tests needs to be run in sequence
// since they are using the interactive print
func Test_ServerTerminate(t *testing.T) {
	interactive.IsInteractive = true

	t.Run("without IP", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: core.ExecStoreBeforeCmd("Server", testServerCommand("image=ubuntu-jammy -w")),
		Cmd:        `scw instance server terminate {{ .Server.ID }}`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				api := instanceSDK.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*instanceSDK.Server)
				_, err := api.GetIP(&instanceSDK.GetIPRequest{
					IP: server.PublicIP.ID,
				})
				assert.NoError(t, err)
			},
		),
		AfterFunc:       core.ExecAfterCmd(`scw instance ip delete {{ index .Server.PublicIP.ID }}`),
		DisableParallel: true,
	}))

	t.Run("with IP", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: core.ExecStoreBeforeCmd("Server", testServerCommand("image=ubuntu-jammy -w")),
		Cmd:        `scw instance server terminate {{ .Server.ID }} with-ip=true`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				api := instanceSDK.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*instanceSDK.Server)
				_, err := api.GetIP(&instanceSDK.GetIPRequest{
					IP: server.PublicIP.ID,
				})
				require.IsType(t, &scw.ResponseError{}, err)
				assert.Equal(t, 403, err.(*scw.ResponseError).StatusCode)
			},
		),
		DisableParallel: true,
	}))

	t.Run("without block", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: core.ExecStoreBeforeCmd("Server", testServerCommand("image=ubuntu-jammy additional-volumes.0=block:10G -w")),
		Cmd:        `scw instance server terminate {{ .Server.ID }} with-ip=true with-block=false`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			core.ExecAfterCmd(`scw instance volume wait {{ (index .Server.Volumes "1").ID }}`),
			core.ExecAfterCmd(`scw instance volume delete {{ (index .Server.Volumes "1").ID }}`),
		),
		DisableParallel: true,
	}))

	t.Run("with block", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: core.ExecStoreBeforeCmd("Server", testServerCommand("image=ubuntu-jammy additional-volumes.0=block:10G -w")),
		Cmd:        `scw instance server terminate {{ .Server.ID }} with-ip=true with-block=true -w`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				api := instanceSDK.NewAPI(ctx.Client)
				server := ctx.Meta["Server"].(*instanceSDK.Server)
				_, err := api.GetVolume(&instanceSDK.GetVolumeRequest{
					VolumeID: server.Volumes["0"].ID,
					Zone:     server.Zone,
				})
				require.IsType(t, &scw.ResourceNotFoundError{}, err)
			},
		),
		DisableParallel: true,
	}))

	interactive.IsInteractive = false
}

// These tests needs to be run in sequence
// since they are using the interactive print
func Test_ServerBackup(t *testing.T) {
	t.Run("simple", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: core.ExecStoreBeforeCmd("Server", testServerCommand("stopped=true image=ubuntu-jammy")),
		Cmd:        `scw instance server backup {{ .Server.ID }} name=backup`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			core.ExecAfterCmd("scw instance image delete {{ .CmdResult.Image.ID }} with-snapshots=true"),
			core.ExecAfterCmd("scw instance server delete {{ .Server.ID }} with-ip=true with-volumes=local"),
		),
	}))

	t.Run("With SBS volumes", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: core.ExecStoreBeforeCmd("Server", testServerCommand("root-volume=sbs:20G stopped=true image=ubuntu-jammy")),
		Cmd:        `scw instance server backup {{ .Server.ID }} name=backup`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			core.ExecAfterCmd("scw instance image delete {{ .CmdResult.Image.ID }} with-snapshots=true"),
			core.ExecAfterCmd("scw instance server delete {{ .Server.ID }} with-ip=true with-volumes=local"),
		),
	}))
}

func Test_ServerAction(t *testing.T) {
	t.Run("manual poweron", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: core.ExecStoreBeforeCmd("Server", testServerCommand("stopped=true image=ubuntu_jammy")),
		Cmd:        `scw instance server action {{ .Server.ID }} action=poweron --wait`,
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(0),
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				storedServer := ctx.Meta["Server"].(*instanceSDK.Server)
				api := instanceSDK.NewAPI(ctx.Client)
				resp, err := api.GetServer(&instanceSDK.GetServerRequest{
					Zone:     storedServer.Zone,
					ServerID: storedServer.ID,
				})
				assert.Nil(t, err)
				assert.Equal(t, instanceSDK.ServerStateRunning, resp.Server.State)
			},
			core.TestCheckGolden(),
		),
		AfterFunc: core.AfterFuncCombine(
			core.ExecAfterCmd("scw instance server delete {{ .Server.ID }} with-ip=true with-volumes=local force-shutdown=true"),
		),
	}))
}

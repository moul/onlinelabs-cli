package instance_test

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/instance/v1"
	instanceSDK "github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Server
func Test_ListServer(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		Cmd:      "scw instance server list",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
	}))
}

func Test_GetServer(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: createServerBionic("Server"),
		Cmd:        "scw instance server get {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteServer("Server"),
	}))
}

// Volume
func Test_CreateVolume(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		Cmd:      "scw instance volume create name=test volume-type=b_ssd size=20G",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(0),
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				require.NotNil(t, ctx.Result)
				assert.Equal(t, "test", ctx.Result.(*instanceSDK.CreateVolumeResponse).Volume.Name)
			},
		),
		AfterFunc: core.ExecAfterCmd("scw instance volume delete {{ .CmdResult.Volume.ID }}"),
	}))

	t.Run("Bad size unit", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		Cmd:      "scw instance volume create name=test volume-type=b_ssd size=20",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(1),
		),
	}))
}

func Test_ServerUpdate(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: createServerBionic("Server"),
		Cmd:        "scw instance server update {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(0),
			core.TestCheckGolden(),
		),
		AfterFunc: deleteServer("Server"),
	}))

	t.Run("No initial placement group & placement-group-id=none", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: createServerBionic("Server"),
		Cmd:        "scw instance server update {{ .Server.ID }} placement-group-id=none",
		Check: core.TestCheckCombine(
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				require.NoError(t, ctx.Err)
				assert.Nil(t, ctx.Result.(*instanceSDK.UpdateServerResponse).Server.PlacementGroup)
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteServer("Server"),
	}))

	t.Run(`No initial placement group & placement-group-id=<existing pg id>`, core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createPlacementGroup("PlacementGroup"),
			createServerBionic("Server"),
		),
		Cmd: `scw instance server update {{ .Server.ID }} placement-group-id={{ .PlacementGroup.ID }}`,
		Check: core.TestCheckCombine(
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				require.NoError(t, ctx.Err)
				assert.Equal(t,
					ctx.Meta["PlacementGroup"].(*instanceSDK.PlacementGroup).ID,
					ctx.Result.(*instanceSDK.UpdateServerResponse).Server.PlacementGroup.ID,
				)
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deletePlacementGroup("PlacementGroup"),
		),
	}))

	t.Run(`No initial placement group & placement-group-id=<valid, but non existing pg id>`, core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: createServerBionic("Server"),
		Cmd:        `scw instance server update {{ .Server.ID }} placement-group-id=11111111-1111-1111-1111-111111111111`,
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(1),
			core.TestCheckGolden(),
		),
		AfterFunc: deleteServer("Server"),
	}))

	t.Run(`No initial placement group & placement-group-id=<invalid pg id>`, core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: createServerBionic("Server"),
		Cmd:        `scw instance server update {{ .Server.ID }} placement-group-id=1111111`,
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(1),
			core.TestCheckGolden(),
		),
		AfterFunc: deleteServer("Server"),
	}))

	t.Run(`Initial placement group & placement-group-id=none`, core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createPlacementGroup("PlacementGroup"),
			core.ExecStoreBeforeCmd("Server", testServerCommand("image=ubuntu-bionic placement-group-id={{ .PlacementGroup.ID }} stopped=true")),
		),
		Cmd: `scw instance server update {{ .Server.ID }} placement-group-id=none`,
		Check: core.TestCheckCombine(
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				require.NoError(t, ctx.Err)
				assert.Nil(t, ctx.Result.(*instanceSDK.UpdateServerResponse).Server.PlacementGroup)
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deletePlacementGroup("PlacementGroup"),
		),
	}))

	t.Run(`Initial placement group & placement-group-id=<current pg id>`, core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createPlacementGroup("PlacementGroup"),
			core.ExecStoreBeforeCmd("Server", testServerCommand("image=ubuntu-bionic placement-group-id={{ .PlacementGroup.ID }} stopped=true")),
		),
		Cmd: `scw instance server update {{ .Server.ID }} placement-group-id={{ .PlacementGroup.ID }}`,
		Check: core.TestCheckCombine(
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				require.NoError(t, ctx.Err)
				assert.Equal(t,
					ctx.Meta["PlacementGroup"].(*instanceSDK.PlacementGroup).ID,
					ctx.Result.(*instanceSDK.UpdateServerResponse).Server.PlacementGroup.ID,
				)
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deletePlacementGroup("PlacementGroup"),
		),
	}))

	t.Run(`Initial placement group & placement-group-id=<new pg id>`, core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			createPlacementGroup("PlacementGroup1"),
			createPlacementGroup("PlacementGroup2"),
			core.ExecStoreBeforeCmd("Server", testServerCommand("image=ubuntu-bionic placement-group-id={{ .PlacementGroup1.ID }} stopped=true")),
		),
		Cmd: `scw instance server update {{ .Server.ID }} placement-group-id={{ .PlacementGroup2.ID }}`,
		Check: core.TestCheckCombine(
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				require.NoError(t, ctx.Err)
				assert.Equal(t,
					ctx.Meta["PlacementGroup2"].(*instanceSDK.PlacementGroup).ID,
					ctx.Result.(*instanceSDK.UpdateServerResponse).Server.PlacementGroup.ID,
				)
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deletePlacementGroup("PlacementGroup1"),
			deletePlacementGroup("PlacementGroup2"),
		),
	}))
}

// Snapshot
func Test_SnapshotCreate(t *testing.T) {
	t.Run("simple", core.Test(&core.TestConfig{
		Commands:   instance.GetCommands(),
		BeforeFunc: createServerBionic("Server"),
		Cmd:        `scw instance snapshot create volume-id={{ (index .Server.Volumes "0").ID }}`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			core.ExecAfterCmd("scw instance snapshot delete {{ .CmdResult.Snapshot.ID }}"),
			deleteServer("Server"),
		),
	}))
}

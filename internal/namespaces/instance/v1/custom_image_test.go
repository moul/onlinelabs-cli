package instance_test

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/instance/v1"
	instanceSDK "github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Test_ImageCreate(t *testing.T) {
	t.Run("Create simple image", core.Test(&core.TestConfig{
		BeforeFunc: core.BeforeFuncCombine(
			createServer("Server"),
			core.ExecStoreBeforeCmd("Snapshot", `scw instance snapshot create volume-id={{ (index .Server.Volumes "0").ID }}`),
		),
		Commands: instance.GetCommands(),
		Cmd:      "scw instance image create snapshot-id={{ .Snapshot.Snapshot.ID }} arch=x86_64",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			core.ExecAfterCmd("scw instance image delete {{ .CmdResult.Image.ID }}"),
			deleteSnapshot("Snapshot"),
		),
	}))

	t.Run("Use additional snapshots", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			core.ExecStoreBeforeCmd("Server", "scw instance server create type=DEV1-S image=ubuntu_focal root-volume=local:10GB additional-volumes.0=local:10GB -w"),
			core.ExecStoreBeforeCmd("SnapshotA", `scw instance snapshot create -w name=cli-test-image-create-snapshotA volume-id={{ (index .Server.Volumes "0").ID }}`),
			core.ExecStoreBeforeCmd("SnapshotB", `scw instance snapshot create -w name=cli-test-image-create-snapshotB volume-id={{ (index .Server.Volumes "1").ID }}`),
		),
		Cmd: "scw instance image create snapshot-id={{ .SnapshotA.ID }} extra-volumes.0.id={{ .SnapshotB.ID }} arch=x86_64",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			core.ExecAfterCmd("scw instance image delete {{ .CmdResult.Image.ID }}"),
			core.ExecAfterCmd("scw instance snapshot delete {{ .SnapshotA.ID }}"),
			core.ExecAfterCmd("scw instance snapshot delete {{ .SnapshotB.ID }}"),
		),
	}))
}

func Test_ImageDelete(t *testing.T) {
	t.Run("simple", core.Test(&core.TestConfig{
		BeforeFunc: createImage("Image"),
		Commands:   instance.GetCommands(),
		Cmd:        "scw instance image delete {{ .Image.Image.ID }} with-snapshots=true",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				// Assert snapshot are deleted with the image
				api := instanceSDK.NewAPI(ctx.Client)
				_, err := api.GetSnapshot(&instanceSDK.GetSnapshotRequest{
					SnapshotID: ctx.Meta["Snapshot"].(*instanceSDK.CreateSnapshotResponse).Snapshot.ID,
				})
				assert.IsType(t, &scw.ResourceNotFoundError{}, err)
			},
		),
		AfterFunc: deleteServer("Server"),
	}))
}

func createImage(metaKey string) core.BeforeFunc {
	return core.BeforeFuncCombine(
		createServer("Server"),
		core.ExecStoreBeforeCmd("Snapshot", `scw instance snapshot create volume-id={{ (index .Server.Volumes "0").ID }}`),
		core.ExecStoreBeforeCmd(metaKey, `scw instance image create snapshot-id={{ .Snapshot.Snapshot.ID }} arch=x86_64`),
	)
}

func deleteImage(metaKey string) core.AfterFunc {
	return core.ExecAfterCmd(`scw instance image delete {{ .` + metaKey + `.Image.ID }} with-snapshots=true`)
}

func Test_ImageList(t *testing.T) {
	t.Run("simple", core.Test(&core.TestConfig{
		BeforeFunc: createImage("Image"),
		Commands:   instance.GetCommands(),
		Cmd:        "scw instance image list",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteImage("Image"),
	}))
}

func Test_ImageUpdate(t *testing.T) {
	t.Run("Change name", core.Test(&core.TestConfig{
		BeforeFunc: createImage("ImageName"),
		Commands:   instance.GetCommands(),
		Cmd:        "scw instance image update {{ .ImageName.Image.ID }} name=foo",
		Check: core.TestCheckCombine(
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				assert.NotNil(t, ctx.Result)
				assert.Equal(t, "foo", ctx.Result.(*instanceSDK.UpdateImageResponse).Image.Name)
			},
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deleteImage("ImageName"),
		),
	}))

	t.Run("Change public from default false to true", core.Test(&core.TestConfig{
		BeforeFunc: createImage("ImagePub"),
		Commands:   instance.GetCommands(),
		Cmd:        "scw instance image update {{ .ImagePub.Image.ID }} public=true",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				assert.NotNil(t, ctx.Result)
				assert.Equal(t, true, ctx.Result.(*instanceSDK.UpdateImageResponse).Image.Public)
			},
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deleteImage("ImagePub"),
		),
	}))

	t.Run("Add extra volume", core.Test(&core.TestConfig{
		BeforeFunc: core.BeforeFuncCombine(
			createVolume("Volume", 20, instanceSDK.VolumeVolumeTypeBSSD),
			core.ExecStoreBeforeCmd("SnapshotVol", `scw instance snapshot create -w name=snapVol volume-id={{ .Volume.ID }}`),
			createImage("ImageExtraVol"),
		),
		Commands: instance.GetCommands(),
		Cmd:      "scw instance image update {{ .ImageExtraVol.Image.ID }} extra-volumes.1.id={{ .SnapshotVol.ID }}",
		Check: core.TestCheckCombine(
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				assert.NotNil(t, ctx.Result)
				assert.Equal(t, "snapVol", ctx.Result.(*instanceSDK.UpdateImageResponse).Image.ExtraVolumes["1"].Name)
			},
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deleteImage("ImageExtraVol"),
			deleteVolume("Volume"),
		),
	}))
}

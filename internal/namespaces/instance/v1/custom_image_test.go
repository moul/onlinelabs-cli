package instance

import (
	"testing"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
)

func Test_ImageCreate(t *testing.T) {
	t.Run("simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		BeforeFunc: func(ctx *core.BeforeFuncCtx) error {
			ctx.Meta["Server"] = ctx.ExecuteCmd("scw instance server create image=ubuntu_bionic stopped")
			ctx.Meta["SnapshotResponse"] = ctx.ExecuteCmd(`scw instance snapshot create volume-id={{ (index .Server.Volumes "0").ID }}`)
			return nil
		},
		Cmd: "scw instance image create snapshot-id={{ .SnapshotResponse.Snapshot.ID }} arch=x86_64",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: func(ctx *core.AfterFuncCtx) error {
			ctx.ExecuteCmd("scw instance image delete image-id=" + ctx.CmdResult.(*instance.CreateImageResponse).Image.ID)
			ctx.ExecuteCmd("scw instance snapshot delete snapshot-id={{ .SnapshotResponse.Snapshot.ID }}")
			ctx.ExecuteCmd("scw instance server delete server-id={{ .Server.ID }} delete-ip=true delete-volumes=true force-shutdown=true")
			return nil
		},
	}))
}

package instance_test

import (
	"os"
	"testing"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/instance/v1"
)

func Test_UserDataGet(t *testing.T) {
	t.Run("Get an existing key", core.Test(&core.TestConfig{
		BeforeFunc: core.BeforeFuncCombine(
			createServerBionic("Server"),
			core.ExecBeforeCmd("scw instance user-data set server-id={{.Server.ID}} key=happy content=true"),
		),
		Commands:  instance.GetCommands(),
		Cmd:       "scw instance user-data get server-id={{.Server.ID}} key=happy",
		AfterFunc: deleteServer("Server"),
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
	}))

	t.Run("Get an nonexistent key", core.Test(&core.TestConfig{
		BeforeFunc: createServerBionic("Server"),
		Commands:   instance.GetCommands(),
		Cmd:        "scw instance user-data get server-id={{.Server.ID}} key=happy",
		AfterFunc:  deleteServer("Server"),
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(1),
		),
	}))
}

func Test_UserDataList(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		BeforeFunc: core.BeforeFuncCombine(
			createServerBionic("Server"),
			core.ExecBeforeCmd("scw instance user-data set server-id={{ .Server.ID }} key=foo content=bar"),
			core.ExecBeforeCmd("scw instance user-data set server-id={{ .Server.ID }} key=bar content=foo"),
		),
		Commands:  instance.GetCommands(),
		Cmd:       "scw instance user-data list server-id={{ .Server.ID }}",
		AfterFunc: deleteServer("Server"),
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
	}))
}

func Test_UserDataFileUpload(t *testing.T) {
	content := "cloud-init file content"

	t.Run("on-cloud-init", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			core.ExecStoreBeforeCmd("Server", testServerCommand("stopped=true image=ubuntu-bionic")),
			func(ctx *core.BeforeFuncCtx) error {
				file, _ := os.CreateTemp("", "test")
				_, _ = file.WriteString(content)
				ctx.Meta["filePath"] = file.Name()
				return nil
			},
		),
		Cmd: `scw instance user-data set key=cloud-init server-id={{ .Server.ID }} content=@{{ .filePath }}`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
		),
		AfterFunc: core.AfterFuncCombine(
			func(ctx *core.AfterFuncCtx) error {
				_ = os.RemoveAll(ctx.Meta["filePath"].(string))
				return nil
			},
		),
	}))

	t.Run("on-random-key", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			core.ExecStoreBeforeCmd("Server", testServerCommand("stopped=true image=ubuntu-bionic")),
			func(ctx *core.BeforeFuncCtx) error {
				file, _ := os.CreateTemp("", "test")
				_, _ = file.WriteString(content)
				ctx.Meta["filePath"] = file.Name()
				return nil
			},
		),
		Cmd: `scw instance user-data set key=foobar server-id={{ .Server.ID }} content=@{{ .filePath }}`,
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
		),
		AfterFunc: core.AfterFuncCombine(
			func(ctx *core.AfterFuncCtx) error {
				_ = os.RemoveAll(ctx.Meta["filePath"].(string))
				return nil
			},
		),
	}))
}

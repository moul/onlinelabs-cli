package instance_test

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/instance/v1"
)

func Test_IPAttach(t *testing.T) {
	t.Run("With UUID", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			core.ExecStoreBeforeCmd("Server", "scw instance server create type=DEV1-S stopped=true ip=none"),
			createIP("Ip"),
		),
		Cmd: "scw instance ip attach {{ .Ip.Address }} server-id={{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
		),
		DisableParallel: true,
	}))

	t.Run("With IP", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			core.ExecStoreBeforeCmd("Server", "scw instance server create type=DEV1-S stopped=true ip=none"),
			createIP("Ip"),
		),
		Cmd: "scw instance ip attach {{ .Ip.Address }} server-id={{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
		),
		DisableParallel: true,
	}))
}

func Test_IPDetach(t *testing.T) {
	t.Run("With UUID", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			core.ExecStoreBeforeCmd("Server", "scw instance server create type=DEV1-S stopped=true ip=none"),
			createIP("Ip"),
			core.ExecBeforeCmd("scw instance ip attach {{ .Ip.Address }} server-id={{ .Server.ID }}"),
		),
		Cmd: "scw instance ip detach {{ .Ip.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deleteIP("Ip"),
		),
		DisableParallel: true,
	}))

	t.Run("With IP", core.Test(&core.TestConfig{
		Commands: instance.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			core.ExecStoreBeforeCmd("Server", "scw instance server create type=DEV1-S stopped=true ip=none"),
			createIP("Ip"),
			core.ExecBeforeCmd("scw instance ip attach {{ .Ip.Address }} server-id={{ .Server.ID }}"),
		),
		Cmd: "scw instance ip detach {{ .Ip.Address }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			core.TestCheckExitCode(0),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServer("Server"),
			deleteIP("Ip"),
		),
		DisableParallel: true,
	}))
}

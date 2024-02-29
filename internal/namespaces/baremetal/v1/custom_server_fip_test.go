package baremetal

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-cli/v2/internal/interactive"
	flexibleip "github.com/scaleway/scaleway-cli/v2/internal/namespaces/flexibleip/v1alpha1"
)

func Test_CreateFlexibleIPInteractive(t *testing.T) {
	promptResponse := []string{
		`" "`,
	}
	interactive.IsInteractive = true
	cmds := GetCommands()
	cmds.Merge(flexibleip.GetCommands())
	t.Run("Simple Interactive", core.Test(&core.TestConfig{
		Commands: cmds,
		BeforeFunc: core.BeforeFuncCombine(
			createServerAndWaitDefault("Server"),
		),
		Cmd: "scw baremetal server add-flexible-ip {{ .Server.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServerDefault("Server"),
			core.ExecAfterCmd("scw fip ip delete {{ .CmdResult.ID }}"),
		),
		PromptResponseMocks: promptResponse,
	}))
}

func Test_CreateFlexibleIP(t *testing.T) {
	interactive.IsInteractive = false
	cmds := GetCommands()
	cmds.Merge(flexibleip.GetCommands())
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: cmds,
		BeforeFunc: core.BeforeFuncCombine(
			createServerAndWaitDefault("Server"),
		),
		Cmd: "scw baremetal server add-flexible-ip {{ .Server.ID }} ip-type=IPv4",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
		),
		AfterFunc: core.AfterFuncCombine(
			deleteServerDefault("Server"),
			core.ExecAfterCmd("scw fip ip delete {{ .CmdResult.ID }}"),
		),
	}))
}

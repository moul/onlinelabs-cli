package instance_test

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/instance/v1"
)

func Test_GetPlacementGroup(t *testing.T) {
	t.Run("simple", core.Test(&core.TestConfig{
		BeforeFunc: core.BeforeFuncCombine(
			core.ExecStoreBeforeCmd("PlacementGroup", "scw instance placement-group create"),
			core.ExecStoreBeforeCmd("ServerA", "scw instance server create type=DEV1-S image=ubuntu_jammy stopped=true placement-group-id={{ .PlacementGroup.PlacementGroup.ID }}"),
		),
		Commands: instance.GetCommands(),
		Cmd:      "scw instance placement-group get {{ .PlacementGroup.PlacementGroup.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
		),
		AfterFunc: core.AfterFuncCombine(
			core.ExecAfterCmd("scw instance placement-group delete {{ .PlacementGroup.PlacementGroup.ID }}"),
			core.ExecAfterCmd("scw instance server delete {{ .ServerA.ID }}"),
		),
	}))
}

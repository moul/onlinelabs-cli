package vpcgw

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/vpc/v2"
	"github.com/scaleway/scaleway-cli/v2/internal/testhelpers"
)

func Test_vpcGwGatewayGet(t *testing.T) {
	cmds := GetCommands()
	cmds.Merge(vpc.GetCommands())

	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: cmds,
		BeforeFunc: core.BeforeFuncCombine(
			testhelpers.CreatePN(),
			testhelpers.CreateGateway("GW"),
			testhelpers.CreateGatewayNetwork("GW"),
		),
		Cmd:   "scw vpc-gw gateway get {{ .GW.ID }}",
		Check: core.TestCheckGolden(),
		AfterFunc: core.AfterFuncCombine(
			testhelpers.DeleteGatewayNetwork(),
			testhelpers.DeletePN(),
			testhelpers.DeleteGateway("GW"),
			testhelpers.DeleteIPVpcGw("GW"),
		),
	}))
}

package llm_inference_test

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
	llm_inference "github.com/scaleway/scaleway-cli/v2/internal/namespaces/llm_inference/v1beta1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/vpc/v2"
)

func Test_createEndpoint(t *testing.T) {
	t.Skip("No stock to run test")
	cmds := llm_inference.GetCommands()
	cmds.Merge(vpc.GetCommands())

	t.Run("Create Private Endpoint", core.Test(&core.TestConfig{
		Commands: cmds,
		BeforeFunc: core.BeforeFuncCombine(
			CreatePN(),
			CreateDeploymentPublicEndpoint(),
		),
		Cmd: "scw llm-inference endpoint create deployment-id={{ .DEPLOYMENT.ID }} endpoint.private-network.private-network-id={{ .PN.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
		),
		AfterFunc: core.AfterFuncCombine(
			core.ExecAfterCmd("scw llm-inference endpoint delete {{ .CmdResult.ID }}"),
			DeletePrivateNetwork(),
			DeleteDeployment(),
		),
	}))

	t.Run("Create Public Endpoint", core.Test(&core.TestConfig{
		Commands: cmds,
		BeforeFunc: core.BeforeFuncCombine(
			CreatePN(),
			CreateDeploymentPrivateEndpoint(),
		),
		Cmd: "scw llm-inference endpoint create deployment-id={{ .DEPLOYMENT.ID }} endpoint.is-public=true",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
		),
		AfterFunc: core.AfterFuncCombine(
			core.ExecAfterCmd("scw llm-inference endpoint delete {{ .CmdResult.ID }}"),
			DeletePrivateNetwork(),
			DeleteDeployment(),
		),
	}))
}

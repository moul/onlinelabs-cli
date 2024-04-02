package llm_inference_test

import "github.com/scaleway/scaleway-cli/v2/internal/core"

func CreateDeploymentPublicEndpoint() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"DEPLOYMENT",
		"scw llm-inference deployment create node-type=H100 accept-eula=true model-name=meta/llama-2-70b-chat:fp8 -w",
	)
}

func CreateDeploymentPrivateEndpoint() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"DEPLOYMENT",
		"scw llm-inference deployment create node-type=H100 accept-eula=true model-name=meta/llama-2-70b-chat:fp8 endpoints.0.private-network.private-network-id={{ .PN.ID }} -w",
	)
}

func CreatePN() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"PN",
		"scw vpc private-network create",
	)
}

func DeletePrivateNetwork() core.AfterFunc {
	return core.ExecAfterCmd("scw vpc private-network delete {{ .PN.ID }}")
}

func DeleteDeployment() core.AfterFunc {
	return core.ExecAfterCmd("scw llm-inference deployment delete {{ .DEPLOYMENT.ID }}")
}

package k8s_test

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/k8s/v1"
)

func Test_GetKubeconfig(t *testing.T) {
	// simple, auth-mode= not provided
	t.Run("simple", core.Test(&core.TestConfig{
		Commands:   k8s.GetCommands(),
		BeforeFunc: createClusterAndWaitAndKubeconfig("get-kubeconfig", "Cluster", "Kubeconfig", kapsuleVersion),
		Cmd:        "scw k8s kubeconfig get {{ .Cluster.ID }}",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			func(t *testing.T, _ *core.CheckFuncCtx) {
				t.Helper()
				// config, err := yaml.Marshal(ctx.Meta["Kubeconfig"].(api.Config))
				// assert.Equal(t, err, nil)
				// assert.Equal(t, ctx.Result.(string), string(config))
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteCluster("Cluster"),
	}))

	t.Run("legacy", core.Test(&core.TestConfig{
		Commands:   k8s.GetCommands(),
		BeforeFunc: createClusterAndWaitAndKubeconfig("get-kubeconfig", "Cluster", "Kubeconfig", kapsuleVersion),
		Cmd:        "scw k8s kubeconfig get {{ .Cluster.ID }} auth-method=legacy",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			func(t *testing.T, _ *core.CheckFuncCtx) {
				t.Helper()
				// config, err := yaml.Marshal(ctx.Meta["Kubeconfig"].(api.Config))
				// assert.Equal(t, err, nil)
				// assert.Equal(t, ctx.Result.(string), string(config))
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteCluster("Cluster"),
	}))

	t.Run("cli", core.Test(&core.TestConfig{
		Commands:   k8s.GetCommands(),
		BeforeFunc: createClusterAndWaitAndKubeconfig("get-kubeconfig", "Cluster", "Kubeconfig", kapsuleVersion),
		Cmd:        "scw k8s kubeconfig get {{ .Cluster.ID }} auth-method=cli",
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			func(t *testing.T, _ *core.CheckFuncCtx) {
				t.Helper()
				// config, err := yaml.Marshal(ctx.Meta["Kubeconfig"].(api.Config))
				// assert.Equal(t, err, nil)
				// assert.Equal(t, ctx.Result.(string), string(config))
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteCluster("Cluster"),
	}))

	// t.Run("copy-token", core.Test(&core.TestConfig{
	//	Commands:   k8s.GetCommands(),
	//	BeforeFunc: createClusterAndWaitAndKubeconfig("get-kubeconfig", "Cluster", "Kubeconfig", kapsuleVersion),
	//	Cmd:        "scw k8s kubeconfig get {{ .Cluster.ID }} auth-method=copy-token",
	//	Check: core.TestCheckCombine(
	//		core.TestCheckGoldenAndReplacePatterns(
	//			core.GoldenReplacement{
	//				Pattern:       regexp.MustCompile("token: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"),
	//				Replacement:   "token: 11111111-1111-1111-1111-111111111111",
	//				OptionalMatch: false,
	//			},
	//		),
	//		func(t *testing.T, _ *core.CheckFuncCtx) {
	//			// config, err := yaml.Marshal(ctx.Meta["Kubeconfig"].(api.Config))
	//			// assert.Equal(t, err, nil)
	//			// assert.Equal(t, ctx.Result.(string), string(config))
	//		},
	//		core.TestCheckExitCode(0),
	//	),
	//	AfterFunc: deleteCluster("Cluster"),
	// }))
}

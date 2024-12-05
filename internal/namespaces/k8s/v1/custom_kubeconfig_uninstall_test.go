package k8s_test

import (
	"os"
	"path"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/ghodss/yaml"
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/k8s/v1"
	api "github.com/scaleway/scaleway-cli/v2/internal/namespaces/k8s/v1/types"
	k8sSDK "github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
)

// testIfKubeconfigNotInFile checks if the given kubeconfig is not in the given file
// it tests if the user, cluster and context of the kubeconfig file are not in the given file
func testIfKubeconfigNotInFile(t *testing.T, filePath string, suffix string, kubeconfig api.Config) {
	t.Helper()
	kubeconfigBytes, err := os.ReadFile(filePath)
	assert.Nil(t, err)
	var existingKubeconfig k8sSDK.Kubeconfig
	err = yaml.Unmarshal(kubeconfigBytes, &existingKubeconfig)
	assert.Nil(t, err)

	found := false
	for _, cluster := range existingKubeconfig.Clusters {
		if cluster.Name == kubeconfig.Clusters[0].Name+suffix {
			found = true
			break
		}
	}
	assert.False(t, found, "cluster found in kubeconfig for cluster with suffix %s", suffix)

	found = false
	for _, context := range existingKubeconfig.Contexts {
		if context.Name == kubeconfig.Clusters[0].Name+suffix {
			found = true
			break
		}
	}
	assert.False(t, found, "context found in kubeconfig for cluster with suffix %s", suffix)

	found = false
	for _, user := range existingKubeconfig.Users {
		if user.Name == kubeconfig.Clusters[0].Name+suffix {
			found = true
			break
		}
	}
	assert.False(t, found, "user found in kubeconfig with suffix %s", suffix)
}

func Test_UninstallKubeconfig(t *testing.T) {
	////
	// Simple use cases
	////
	t.Run("uninstall", core.Test(&core.TestConfig{
		Commands:   k8s.GetCommands(),
		BeforeFunc: createClusterAndWaitAndInstallKubeconfig("uninstall-kubeconfig", "Cluster", "Kubeconfig", kapsuleVersion),
		Cmd:        "scw k8s kubeconfig uninstall {{ .Cluster.ID }}",
		Check: core.TestCheckCombine(
			// no golden tests since it's os specific
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				testIfKubeconfigNotInFile(t, path.Join(os.TempDir(), "cli-uninstall-test"), "-"+ctx.Meta["Cluster"].(*k8sSDK.Cluster).ID, ctx.Meta["Kubeconfig"].(api.Config))
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteCluster("Cluster"),
		OverrideEnv: map[string]string{
			"KUBECONFIG": path.Join(os.TempDir(), "cli-uninstall-test"),
		},
	}))
	t.Run("empty file", core.Test(&core.TestConfig{
		Commands:   k8s.GetCommands(),
		BeforeFunc: createClusterAndWaitAndKubeconfig("uninstall-kubeconfig-empty", "EmptyCluster", "EmptyKubeconfig", kapsuleVersion),
		Cmd:        "scw k8s kubeconfig uninstall {{ .EmptyCluster.ID }}",
		Check: core.TestCheckCombine(
			// no golden tests since it's os specific
			func(t *testing.T, _ *core.CheckFuncCtx) {
				t.Helper()
				_, err := os.Stat(path.Join(os.TempDir(), "emptyfile"))
				assert.True(t, os.IsNotExist(err))
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteCluster("EmptyCluster"),
		OverrideEnv: map[string]string{
			"KUBECONFIG": path.Join(os.TempDir(), "emptyfile"),
		},
	}))
	t.Run("uninstall-merge", core.Test(&core.TestConfig{
		Commands:   k8s.GetCommands(),
		BeforeFunc: createClusterAndWaitAndKubeconfigAndPopulateFileAndInstall("uninstall-kubeconfig-merge", "Cluster", "Kubeconfig", kapsuleVersion, path.Join(os.TempDir(), "cli-uninstall-merge-test"), []byte(existingKubeconfig)),
		Cmd:        "scw k8s kubeconfig uninstall {{ .Cluster.ID }}",
		Check: core.TestCheckCombine(
			// no golden tests since it's os specific
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				testIfKubeconfigNotInFile(t, path.Join(os.TempDir(), "cli-uninstall-merge-test"), "-"+ctx.Meta["Cluster"].(*k8sSDK.Cluster).ID, ctx.Meta["Kubeconfig"].(api.Config))
			},
			core.TestCheckExitCode(0),
		),
		AfterFunc: deleteCluster("Cluster"),
		OverrideEnv: map[string]string{
			"KUBECONFIG": path.Join(os.TempDir(), "cli-uninstall-merge-test"),
		},
	}))
}

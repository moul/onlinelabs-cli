package k8s

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/scaleway/scaleway-cli/v2/core"
)

type k8sKubeconfigUninstallRequest struct {
	ClusterID string
}

func k8sKubeconfigUninstallCommand() *core.Command {
	return &core.Command{
		Short: `Uninstall a kubeconfig`,
		Long: `Remove specified cluster from kubeconfig file specified by the KUBECONFIG env, if empty it will default to $HOME/.kube/config.
If the current context points to this cluster, it will be set to an empty context.`,
		Namespace: "k8s",
		Verb:      "uninstall",
		Resource:  "kubeconfig",
		ArgsType:  reflect.TypeOf(k8sKubeconfigUninstallRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      "Cluster ID from which to uninstall the kubeconfig",
				Required:   true,
				Positional: true,
			},
		},
		Run: k8sKubeconfigUninstallRun,
		Examples: []*core.Example{
			{
				Short:    "Uninstall the kubeconfig for a given cluster",
				ArgsJSON: `{"cluster_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw k8s kubeconfig install",
				Short:   "Install a kubeconfig",
			},
		},
	}
}

// k8sKubeconfigUninstallRun use the specified cluster ID to remove it from the wanted kubeconfig file
// it removes all the users, contexts and clusters that contains this ID from the file
func k8sKubeconfigUninstallRun(ctx context.Context, argsI interface{}) (i interface{}, e error) {
	request := argsI.(*k8sKubeconfigUninstallRequest)

	kubeconfigPath, err := getKubeconfigPath(ctx)
	if err != nil {
		return nil, err
	}

	// if the file does not exist, the cluster is not there
	if _, err := os.Stat(kubeconfigPath); os.IsNotExist(err) {
		return fmt.Sprintf("File %s does not exist.", kubeconfigPath), nil
	}

	kmc, err := LoadKubeMapConfig(ctx, kubeconfigPath)
	if err != nil {
		return nil, err
	}
	kubeconfig := kmc.Kubeconfig()

	for _, kubeconfigContext := range kubeconfig.Contexts {
		if strings.HasSuffix(kubeconfigContext.Name, request.ClusterID) {
			err = kmc.RemoveContext(kubeconfigContext.Name)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, cluster := range kubeconfig.Clusters {
		if strings.HasSuffix(cluster.Name, request.ClusterID) {
			err = kmc.RemoveCluster(cluster.Name)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, user := range kubeconfig.AuthInfos {
		if strings.HasSuffix(user.Name, request.ClusterID) {
			err = kmc.RemoveUser(user.Name)
			if err != nil {
				return nil, err
			}
		}
	}

	// reset the current context
	if strings.HasSuffix(kubeconfig.CurrentContext, request.ClusterID) {
		kmc.CurrentContext = ""
	}

	err = kmc.Save(kubeconfigPath)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("Cluster %s successfully deleted from %s", request.ClusterID, kubeconfigPath), nil
}

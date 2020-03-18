package k8s

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/scaleway/scaleway-cli/internal/core"
	k8s "github.com/scaleway/scaleway-sdk-go/api/k8s/v1beta4"
)

type k8sKubeconfigUninstallRequest struct {
	ClusterID string
}

func k8sKubeconfigUninstallCommand() *core.Command {
	return &core.Command{
		Short:     `Uninstall a kubeconfig`,
		Long:      `Remove specified cluster from kubeconfig file specified by the KUBECONFIG env, if empty it will default to $HOME/.kube/config.`,
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
		return fmt.Sprintf("File %s does not exists.", kubeconfigPath), nil
	}

	existingKubeconfig, err := openAndUnmarshalKubeconfig(kubeconfigPath)
	if err != nil {
		return nil, err
	}

	// delete the wanted cluster from the file
	newClusters := []*k8s.KubeconfigClusterWithName{}
	for _, cluster := range existingKubeconfig.Clusters {
		if !strings.HasSuffix(cluster.Name, request.ClusterID) {
			newClusters = append(newClusters, cluster)
		}
	}

	// delete the wanted context from the file
	newContexts := []*k8s.KubeconfigContextWithName{}
	for _, kubeconfigContext := range existingKubeconfig.Contexts {
		if !strings.HasSuffix(kubeconfigContext.Name, request.ClusterID) {
			newContexts = append(newContexts, kubeconfigContext)
		}
	}

	// delete the wanted user from the file
	newUsers := []*k8s.KubeconfigUserWithName{}
	for _, user := range existingKubeconfig.Users {
		if !strings.HasSuffix(user.Name, request.ClusterID) {
			newUsers = append(newUsers, user)
		}
	}

	// reset the current context
	existingKubeconfig.CurrentContext = ""

	// write the modification
	existingKubeconfig.Clusters = newClusters
	existingKubeconfig.Contexts = newContexts
	existingKubeconfig.Users = newUsers

	err = marshalAndWriteKubeconfig(existingKubeconfig, kubeconfigPath)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("Cluster %s successfully deleted from %s", request.ClusterID, kubeconfigPath), nil
}

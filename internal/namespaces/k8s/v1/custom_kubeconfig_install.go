package k8s

import (
	"context"
	"errors"
	"fmt"
	"hash/crc32"
	"os"
	"reflect"

	"github.com/ghodss/yaml"
	"github.com/scaleway/scaleway-cli/v2/core"
	api "github.com/scaleway/scaleway-cli/v2/internal/namespaces/k8s/v1/types"
	k8s "github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type k8sKubeconfigInstallRequest struct {
	ClusterID          string
	Region             scw.Region
	KeepCurrentContext bool
	AuthMethod         authMethods
}

func k8sKubeconfigInstallCommand() *core.Command {
	return &core.Command{
		Short: `Install a kubeconfig`,
		Long: `Retrieve the kubeconfig for a specified cluster and write it on disk. 
It will merge the new kubeconfig in the file pointed by the KUBECONFIG variable. If empty it will default to $HOME/.kube/config.`,
		Namespace: "k8s",
		Verb:      "install",
		Resource:  "kubeconfig",
		ArgsType:  reflect.TypeOf(k8sKubeconfigInstallRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      "Cluster ID from which to retrieve the kubeconfig",
				Required:   true,
				Positional: true,
			},
			{
				Name:  "keep-current-context",
				Short: "Whether or not to keep the current kubeconfig context unmodified",
			},
			{
				Name:    "auth-method",
				Short:   `Which method to use to authenticate using kubelet`,
				Default: core.DefaultValueSetter(string(authMethodLegacy)),
				EnumValues: []string{
					string(authMethodLegacy),
					string(authMethodCLI),
					string(authMethodCopyToken),
				},
			},
			core.RegionArgSpec(),
		},
		Run: k8sKubeconfigInstallRun,
		Examples: []*core.Example{
			{
				Short:    "Install the kubeconfig for a given cluster and using the new context",
				ArgsJSON: `{"cluster_id": "11111111-1111-1111-1111-111111111111"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Command: "scw k8s kubeconfig uninstall",
				Short:   "Uninstall a kubeconfig",
			},
		},
	}
}

func k8sKubeconfigInstallRun(ctx context.Context, argsI interface{}) (i interface{}, e error) {
	request := argsI.(*k8sKubeconfigInstallRequest)

	apiKubeconfigResp, err := k8s.NewAPI(core.ExtractClient(ctx)).GetClusterKubeConfig(&k8s.GetClusterKubeConfigRequest{
		Region:    request.Region,
		ClusterID: request.ClusterID,
	})
	if err != nil {
		return nil, err
	}

	var clusterKubeconfig api.Config
	err = yaml.Unmarshal(apiKubeconfigResp.GetRaw(), &clusterKubeconfig)
	if err != nil {
		return nil, err
	}

	kubeconfigPath, err := getKubeconfigPath(ctx)
	if err != nil {
		return nil, err
	}

	kmc := NewKubeMapConfig()
	if _, err := os.Stat(kubeconfigPath); err == nil {
		kmc, err = LoadKubeMapConfig(ctx, kubeconfigPath)
		if err != nil {
			return nil, err
		}
	}

	// insert
	newNamedUser, err := buildAuthUser(ctx, clusterKubeconfig, request.ClusterID, request.AuthMethod)
	if err != nil {
		return nil, err
	}

	err = kmc.SetUser(newNamedUser.Name, newNamedUser.AuthInfo, true)
	if err != nil {
		return nil, err
	}

	clusterNameWithID := fmt.Sprintf("%s-%s", clusterKubeconfig.Clusters[0].Name, request.ClusterID)
	err = kmc.SetCluster(clusterNameWithID, clusterKubeconfig.Clusters[0].Cluster, true)
	if err != nil {
		return nil, err
	}

	err = kmc.SetContext(clusterNameWithID, api.Context{Cluster: clusterNameWithID, AuthInfo: newNamedUser.Name}, true)
	if err != nil {
		return nil, err
	}

	if !request.KeepCurrentContext {
		kmc.CurrentContext = clusterNameWithID
	}

	err = kmc.Save(kubeconfigPath)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("Kubeconfig for cluster %s successfully written at %s", request.ClusterID, kubeconfigPath), nil
}

func buildAuthUser(ctx context.Context, config api.Config, clusterID string, met authMethods) (*api.NamedAuthInfo, error) {
	switch met {
	case authMethodLegacy:
		if config.AuthInfos[0].AuthInfo.Token == RedactedAuthInfoToken {
			return nil, errors.New("this cluster does not support legacy authentication")
		}

		return &api.NamedAuthInfo{
			Name: fmt.Sprintf("%s-%s", config.Clusters[0].Name, clusterID),
			AuthInfo: api.AuthInfo{
				Token: config.AuthInfos[0].AuthInfo.Token,
			},
		}, nil
	case authMethodCLI:
		args := []string{}
		profileName := core.ExtractProfileName(ctx)
		if profileName != scw.DefaultProfileName {
			args = append(args, "--profile", profileName)
		}

		var configPath string
		switch {
		case core.ExtractConfigPathFlag(ctx) != "":
			configPath = core.ExtractConfigPathFlag(ctx)
			args = append(args, "--config", configPath)
		case core.ExtractEnv(ctx, scw.ScwConfigPathEnv) != "":
			configPath = core.ExtractEnv(ctx, scw.ScwConfigPathEnv)
			args = append(args, "--config", configPath)
		}

		configPathSum := crc32.ChecksumIEEE([]byte(configPath))
		return &api.NamedAuthInfo{
			Name: fmt.Sprintf("cli-%s-%08x", profileName, configPathSum),
			AuthInfo: api.AuthInfo{
				Exec: &api.ExecConfig{
					APIVersion: "client.authentication.k8s.io/v1",
					Command:    core.ExtractBinaryName(ctx),
					Args: append(args,
						"k8s",
						"exec-credential",
					),
					InstallHint: installHint,
				},
			},
		}, nil
	case authMethodCopyToken:
		token, err := SecretKey(ctx)
		if err != nil {
			return nil, err
		}

		tokenSum := crc32.ChecksumIEEE([]byte(token))
		return &api.NamedAuthInfo{
			Name: fmt.Sprintf("token-cli-%08x", tokenSum),
			AuthInfo: api.AuthInfo{
				Token: token,
			},
		}, nil
	}
	return nil, errors.New("unknown auth method")
}

package k8s

import (
	"context"
	"errors"
	"fmt"
	"hash/crc32"
	"reflect"

	"github.com/ghodss/yaml"
	"github.com/scaleway/scaleway-cli/v2/core"
	api "github.com/scaleway/scaleway-cli/v2/internal/namespaces/k8s/v1/types"
	k8s "github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type k8sKubeconfigGetRequest struct {
	ClusterID  string
	Region     scw.Region
	AuthMethod authMethods
}

func k8sKubeconfigGetCommand() *core.Command {
	return &core.Command{
		Short:     `Retrieve a kubeconfig`,
		Long:      `Retrieve the kubeconfig for a specified cluster.`,
		Namespace: "k8s",
		Verb:      "get",
		Resource:  "kubeconfig",
		ArgsType:  reflect.TypeOf(k8sKubeconfigGetRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "cluster-id",
				Short:      "Cluster ID from which to retrieve the kubeconfig",
				Required:   true,
				Positional: true,
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
		Run: k8sKubeconfigGetRun,
		Examples: []*core.Example{
			{
				Short:    "Get the kubeconfig for a given cluster",
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

func k8sKubeconfigGetRun(ctx context.Context, argsI interface{}) (i interface{}, e error) {
	request := argsI.(*k8sKubeconfigGetRequest)

	apiKubeconfig, err := k8s.NewAPI(core.ExtractClient(ctx)).GetClusterKubeConfig(&k8s.GetClusterKubeConfigRequest{
		Region:    request.Region,
		ClusterID: request.ClusterID,
	})
	if err != nil {
		return nil, err
	}

	var kubeconfig api.Config
	err = yaml.Unmarshal(apiKubeconfig.GetRaw(), &kubeconfig)
	if err != nil {
		return nil, err
	}

	namedClusterInfo := api.NamedCluster{
		Name:    fmt.Sprintf("%s-%s", kubeconfig.Clusters[0].Name, request.ClusterID),
		Cluster: kubeconfig.Clusters[0].Cluster,
	}

	var namedAuthInfo api.NamedAuthInfo
	switch request.AuthMethod {
	case authMethodLegacy:
		if kubeconfig.AuthInfos[0].AuthInfo.Token == RedactedAuthInfoToken {
			return nil, errors.New("this cluster does not support legacy authentication")
		}

		namedAuthInfo.Name = fmt.Sprintf("%s-%s", kubeconfig.Clusters[0].Name, request.ClusterID)
		namedAuthInfo.AuthInfo.Token = kubeconfig.AuthInfos[0].AuthInfo.Token
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
		namedAuthInfo.Name = fmt.Sprintf("cli-%s-%08x", profileName, configPathSum)
		namedAuthInfo.AuthInfo = api.AuthInfo{
			Exec: &api.ExecConfig{
				APIVersion: "client.authentication.k8s.io/v1",
				Command:    core.ExtractBinaryName(ctx),
				Args: append(args,
					"k8s",
					"exec-credential",
				),
				InstallHint: installHint,
			},
		}
	case authMethodCopyToken:
		token, err := SecretKey(ctx)
		if err != nil {
			return nil, err
		}

		tokenSum := crc32.ChecksumIEEE([]byte(token))
		namedAuthInfo.Name = fmt.Sprintf("token-cli-%08x", tokenSum)
		namedAuthInfo.AuthInfo = api.AuthInfo{
			Token: token,
		}
	default:
		return nil, errors.New("unknown auth method")
	}

	namedContext := api.NamedContext{
		Name: fmt.Sprintf("%s-%s", kubeconfig.Clusters[0].Name, request.ClusterID),
		Context: api.Context{
			Cluster:  namedClusterInfo.Name,
			AuthInfo: namedAuthInfo.Name,
		},
	}

	resultingKubeconfig := api.Config{
		APIVersion:     KubeconfigAPIVersion,
		Kind:           KubeconfigKind,
		Clusters:       []api.NamedCluster{namedClusterInfo},
		AuthInfos:      []api.NamedAuthInfo{namedAuthInfo},
		Contexts:       []api.NamedContext{namedContext},
		CurrentContext: namedContext.Name,
	}

	config, err := yaml.Marshal(resultingKubeconfig)
	if err != nil {
		return nil, err
	}

	return string(config), nil
}

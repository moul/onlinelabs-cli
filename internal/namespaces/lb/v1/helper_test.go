package lb_test

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/scaleway/scaleway-cli/v2/core"
	go_api "github.com/scaleway/scaleway-cli/v2/internal/namespaces/k8s/v1/types"
	"github.com/scaleway/scaleway-sdk-go/api/k8s/v1"
	"github.com/scaleway/scaleway-sdk-go/api/lb/v1"
)

func createLB() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"LB",
		"scw lb lb create name=cli-test description=cli-test --wait",
	)
}

func deleteLB() core.AfterFunc {
	return core.ExecAfterCmd("scw lb lb delete {{ .LB.ID }}")
}

func createInstance() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"Instance",
		"scw instance server create type=DEV1-S stopped=true image=ubuntu_focal",
	)
}

func createRunningInstance() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"Instance",
		"scw instance server create type=DEV1-S image=ubuntu_bionic -w",
	)
}

func createRunningInstanceWithTag() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"Instance",
		"scw instance server create type=DEV1-S image=ubuntu_bionic tags.0=foo -w",
	)
}

func deleteInstance() core.AfterFunc {
	return core.ExecAfterCmd("scw instance server delete {{ .Instance.ID }}")
}

func deleteRunningInstance() core.AfterFunc {
	return core.ExecAfterCmd("scw instance server delete {{ .Instance.ID }} force-shutdown=true")
}

func createBackend(forwardPort int32) core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"Backend",
		fmt.Sprintf("scw lb backend create lb-id={{ .LB.ID }} name=cli-test forward-protocol=tcp forward-port=%d forward-port-algorithm=roundrobin sticky-sessions=none health-check.port=8888 health-check.check-max-retries=5", forwardPort),
	)
}

func addIP2Backend(ip string) core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"AddIP2Backend",
		"scw lb backend add-servers {{ .Backend.ID }} server-ip.0="+ip,
	)
}

func createFrontend(inboundPort int32) core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"Frontend",
		fmt.Sprintf("scw lb frontend create lb-id={{ .LB.ID }} backend-id={{ .Backend.ID }} name=cli-test inbound-port=%d", inboundPort),
	)
}

func createClusterAndWaitAndInstallKubeconfig(metaKey string, kubeconfigMetaKey string, version string) core.BeforeFunc {
	return func(ctx *core.BeforeFuncCtx) error {
		cmd := fmt.Sprintf("scw k8s cluster create name=cli-test version=%s cni=cilium pools.0.node-type=DEV1-M pools.0.size=1 pools.0.name=default --wait", version)
		res := ctx.ExecuteCmd(strings.Split(cmd, " "))
		cluster := res.(*k8s.Cluster)
		ctx.Meta[metaKey] = cluster
		api := k8s.NewAPI(ctx.Client)
		apiKubeconfig, err := api.GetClusterKubeConfig(&k8s.GetClusterKubeConfigRequest{
			Region:    cluster.Region,
			ClusterID: cluster.ID,
		})
		if err != nil {
			return err
		}

		var kubeconfig go_api.Config

		err = yaml.Unmarshal(apiKubeconfig.GetRaw(), &kubeconfig)
		if err != nil {
			return err
		}

		ctx.Meta[kubeconfigMetaKey] = kubeconfig
		cmd = "scw k8s kubeconfig install " + cluster.ID
		_ = ctx.ExecuteCmd(strings.Split(cmd, " "))
		return nil
	}
}

func deleteCluster(metaKey string) core.AfterFunc {
	return core.ExecAfterCmd("scw k8s cluster delete {{ ." + metaKey + ".ID }} --wait")
}

func retrieveLBID(metaKey string) core.BeforeFunc {
	return func(ctx *core.BeforeFuncCtx) error {
		_, err := exec.Command("bash", "-c", "kubectl create -f testfixture/lb.yaml").Output()
		if err != nil {
			return err
		}
		// We let enough time for kubeconfig to install
		time.Sleep(5 * time.Second)

		cmd, err := exec.Command("bash", "-c", "kubectl -n kube-system get service/traefik-ingress -o jsonpath='{.metadata.annotations.service\\.beta\\.kubernetes\\.io/scw-loadbalancer-id}'").Output()
		if err != nil {
			return err
		}
		parseID := strings.Split(string(cmd), "/")
		if len(parseID) != 2 {
			return fmt.Errorf("can't parse ID: %s", parseID)
		}
		lbID := parseID[1]

		api := lb.NewZonedAPI(ctx.Client)
		getLB, err := api.GetLB(&lb.ZonedAPIGetLBRequest{
			LBID: lbID,
		})
		if err != nil {
			return err
		}
		ctx.Meta[metaKey] = getLB.ID

		return nil
	}
}

func createPN() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"PN",
		"scw vpc private-network create",
	)
}

func deletePN() core.AfterFunc {
	return core.ExecAfterCmd("scw vpc private-network delete {{ .PN.ID }}")
}

func attachPN() core.BeforeFunc {
	return core.ExecBeforeCmd(
		"scw lb private-network attach {{ .LB.ID }} private-network-id={{ .PN.ID }}",
	)
}

func detachPN() core.AfterFunc {
	return core.ExecAfterCmd(
		"scw lb private-network detach {{ .LB.ID }} private-network-id={{ .PN.ID }}",
	)
}

func createIP() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"IP",
		"scw lb ip create is-ipv6=true",
	)
}

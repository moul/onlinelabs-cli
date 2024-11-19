package vpc_test

import (
	"github.com/scaleway/scaleway-cli/v2/core"
)

func createInstance() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"Instance",
		"scw instance server create type=DEV1-S stopped=true image=ubuntu_focal",
	)
}

func deleteInstance() core.AfterFunc {
	return core.ExecAfterCmd("scw instance server delete {{ .Instance.ID }} --wait")
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

func createNIC() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"NIC",
		"scw instance private-nic create server-id={{ .Instance.ID }} private-network-id={{ .PN.ID }}",
	)
}

func createLB() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"LB",
		"scw lb lb create name=cli-test description=cli-test --wait",
	)
}

func attachLB() core.BeforeFunc {
	return core.ExecBeforeCmd(
		"scw lb private-network attach {{ .LB.ID }} private-network-id={{ .PN.ID }}",
	)
}

func detachLB() core.AfterFunc {
	return core.ExecAfterCmd(
		"scw lb private-network detach {{ .LB.ID }} private-network-id={{ .PN.ID }}",
	)
}

func deleteLB() core.AfterFunc {
	return core.ExecAfterCmd("scw lb lb delete {{ .LB.ID }}")
}

func createRdbInstance() core.BeforeFunc {
	return core.ExecStoreBeforeCmd(
		"RDB",
		"scw rdb instance create node-type=DB-DEV-S is-ha-cluster=false name=cli-test engine=PostgreSQL-12 user-name=foobar password={4xdl*#QOoP+&3XRkGA)] init-endpoints.0.private-network.private-network-id={{ .PN.ID }} init-endpoints.0.private-network.service-ip=192.168.0.1/24 --wait",
	)
}

func detachRdbInstance() core.AfterFunc {
	return core.ExecAfterCmd(
		"scw rdb endpoint delete {{ (index .RDB.Endpoints 0).ID  }} instance-id={{ .RDB.ID }}",
	)
}

func waitRdbInstance() core.AfterFunc {
	return core.ExecAfterCmd(
		"scw rdb instance wait {{ .RDB.ID }}",
	)
}

func deleteRdbInstance() core.AfterFunc {
	return core.ExecAfterCmd("scw rdb instance delete {{ .RDB.ID }}")
}

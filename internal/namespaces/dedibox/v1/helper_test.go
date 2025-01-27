package dedibox_test

import (
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/dedibox/v1"
)

func createServer(metaKey string) core.BeforeFunc {
	return core.ExecStoreBeforeCmd(metaKey, "scw dedibox server create offer-id=28483 zone=fr-par-2")
}

func stopInstall(metaKey string) core.AfterFunc {
	return core.ExecAfterCmd("scw dedibox server cancel-install server-id={{ ." + metaKey + ".ID }}")
}

func deleteServer(metaKey string) core.AfterFunc {
	return func(ctx *core.AfterFuncCtx) error {
		server := ctx.Meta[metaKey].(*dedibox.Server)
		if server.Status == dedibox.ServerStatusReady {
			err := core.ExecAfterCmd("scw dedibox server stop server-id={{ ." + metaKey + ".ID }}")(ctx)
			if err != nil {
				return err
			}
		}
		return core.ExecAfterCmd("scw dedibox server delete server-id={{ ." + metaKey + ".ID }}")(ctx)
	}
}

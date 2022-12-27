package baremetal

import (
	"github.com/scaleway/scaleway-cli/v2/internal/core"
)

// createServerAndWait creates a baremetal instance
// register it in the context Meta at metaKey.
func createServerAndWait(metaKey string) core.BeforeFunc {
	return core.ExecStoreBeforeCmd(metaKey, "scw baremetal server create zone=fr-par-1 type=EM-A210R-HDD -w")
}

func createServer(metaKey string) core.BeforeFunc {
	return core.ExecStoreBeforeCmd(metaKey, "scw baremetal server create zone=fr-par-1 type=EM-A210R-HDD")
}

// deleteServer deletes a server
// previously registered in the context Meta at metaKey.
//
//nolint:unparam
func deleteServer(metaKey string) core.AfterFunc {
	return core.ExecAfterCmd("scw baremetal server delete zone=fr-par-1 {{ ." + metaKey + ".ID }}")
}

// add an ssh key with a given meta key
func addSSH(metaKey string, key string) core.BeforeFunc {
	return func(ctx *core.BeforeFuncCtx) error {
		ctx.Meta[metaKey] = ctx.ExecuteCmd([]string{
			"scw", "iam", "ssh-key", "create", "public-key=" + key,
		})
		return nil
	}
}

// delete an ssh key with a given meta key
func deleteSSH(metaKey string) core.AfterFunc {
	return core.ExecAfterCmd("scw iam ssh-key delete {{ ." + metaKey + ".ID }}")
}

package baremetal

import (
	"testing"

	"github.com/scaleway/scaleway-cli/internal/core"
	account "github.com/scaleway/scaleway-cli/internal/namespaces/account/v2alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func Test_InstallServer(t *testing.T) {
	// All test below should succeed to create an instance.
	t.Run("Simple", func(t *testing.T) {
		// baremetal api requires that the key must be at least 1024 bits long. Regardless of the algorithm
		sshKey := `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCbJuYSOQc01zjHsMyn4OUsW61cqRvttKt3StJgbvt2WBuGpwi1/5RtSoMQpudYlZpdeivFb21S8QRas8zcOc+6WqgWa2nj/8yA+cauRlV6CMWY+hOTkkg39xaekstuQ+WR2/AP7O/9hjVx5735+9ZNIxxHsFjVYdBEuk9gEX+1Rw== foobar@foobar`
		osID := `d859aa89-8b4a-4551-af42-ff7c0c27260a` // Ubuntu 18.04
		cmds := GetCommands()
		cmds.Merge(account.GetCommands())

		t.Run("With ID", core.Test(&core.TestConfig{
			BeforeFunc: core.BeforeFuncCombine(
				addSSH("key", sshKey),
				createServerAndWait("Server"),
			),
			Commands: cmds,
			Cmd:      "scw baremetal server install {{ .Server.ID }} hostname=test-install-server ssh-key-ids.0={{ .key.ID }} os-id=" + osID + " -w",
			Check: core.TestCheckCombine(
				core.TestCheckGolden(),
				core.TestCheckExitCode(0),
			),
			AfterFunc: core.AfterFuncCombine(
				deleteSSH("key"),
				deleteServer("Server"),
			),
			DefaultZone: scw.ZoneFrPar2,
		}))

		t.Run("All SSH keys", core.Test(&core.TestConfig{
			Commands: cmds,
			BeforeFunc: core.BeforeFuncCombine(
				addSSH("key", sshKey),
				createServerAndWait("Server"),
			),
			Cmd: "scw baremetal server install {{ .Server.ID }} hostname=test-install-server all-ssh-keys=true os-id=" + osID + " -w",
			Check: core.TestCheckCombine(
				core.TestCheckGolden(),
				core.TestCheckExitCode(0),
			),
			AfterFunc: core.AfterFuncCombine(
				deleteSSH("key"),
				deleteServer("Server"),
			),
			DefaultZone: scw.ZoneFrPar2,
		}))
	})
}

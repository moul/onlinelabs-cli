package baremetal

import (
	"time"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-cli/internal/human"
	baremetal "github.com/scaleway/scaleway-sdk-go/api/baremetal/v1"
)

var (
	defaultRetryInterval = 15 * time.Second
)

func GetCommands() *core.Commands {
	cmds := GetGeneratedCommands()

	cmds.Merge(core.NewCommands(
		serverWaitCommand(),
	))

	human.RegisterMarshalerFunc(baremetal.ServerPingStatus(0), human.EnumMarshalFunc(serverPingStatusMarshalSpecs))

	cmds.MustFind("baremetal", "server", "create").Override(serverCreateBuilder)
	cmds.MustFind("baremetal", "server", "install").Override(serverInstallBuilder)
	cmds.MustFind("baremetal", "server", "list").Override(serverListBuilder)

	// Action commands
	cmds.MustFind("baremetal", "server", "start").Override(serverStartBuilder)
	cmds.MustFind("baremetal", "server", "stop").Override(serverStopBuilder)
	cmds.MustFind("baremetal", "server", "reboot").Override(serverRebootBuilder)

	return cmds
}

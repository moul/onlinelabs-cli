package dedibox

import (
	"context"

	"github.com/scaleway/scaleway-cli/v2/core"
)

func GetCommands() *core.Commands {
	cmds := GetGeneratedCommands()

	cmds.MustFind("dedibox", "server", "install").Override(serverInstallBuilder)

	for _, commandPath := range [][]string{
		{"dedibox", "server", "list"},
		{"dedibox", "service", "list"},
		{"dedibox", "offer", "list"},
		{"dedibox", "offer", "get"},
		{"dedibox", "os", "list"},
		{"dedibox", "fip", "list"},
		{"dedibox", "fip", "get-quota"},
		{"dedibox", "billing", "list-invoice"},
		{"dedibox", "billing", "list-refund"},
		{"dedibox", "ipv6-block", "get-quota"},
		{"dedibox", "ipv6-block", "create"},
		{"dedibox", "ipv6-block", "get"},
		{"dedibox", "rpn-info", "list"},
		{"dedibox", "rpn-info", "get"},
		{"dedibox", "san", "list"},
		{"dedibox", "rpn-v1", "list"},
		{"dedibox", "rpn-v1", "list-members"},
		{"dedibox", "rpn-v1", "list-capable-server"},
		{"dedibox", "rpn-v1", "list-capable-san-server"},
		{"dedibox", "rpn-v2", "list"},
		{"dedibox", "rpn-v2", "list-capable-resources"},
	} {
		cmds.MustFind(commandPath...).Override(setProjectDefaultValue)
	}

	return cmds
}

func setProjectDefaultValue(c *core.Command) *core.Command {
	c.ArgSpecs.GetByName("project-id").Default = func(ctx context.Context) (value string, doc string) {
		projectID := core.GetProjectIDFromContext(ctx)
		return projectID, "<retrieved from config>"
	}
	return c
}

package dedibox_test

import (
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/dedibox/v1"
	"testing"
)

func Test_InstallServer(t *testing.T) {
	t.Run("Install server without partitions setup", core.Test(&core.TestConfig{
		Commands:            dedibox.GetCommands(),
		BeforeFunc:          createServer(),
		Cmd:                 "",
		Args:                nil,
		Check:               nil,
		AfterFunc:           nil,
		DisableParallel:     false,
		BuildInfo:           nil,
		TmpHomeDir:          false,
		OverrideEnv:         nil,
		OverrideExec:        nil,
		Client:              nil,
		Ctx:                 nil,
		PromptResponseMocks: nil,
		Stdin:               nil,
		EnableAliases:       false,
	}))
}

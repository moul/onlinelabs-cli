package core_test

import (
	"testing"

	"github.com/scaleway/scaleway-cli/v2/internal/core"

	"github.com/alecthomas/assert"
	"github.com/scaleway/scaleway-cli/v2/internal/alias"
)

func TestCommandMatchAlias(t *testing.T) {
	commandWithArg := &core.Command{
		Namespace: "first",
		Resource:  "command",
		ArgSpecs: core.ArgSpecs{
			{
				Name: "arg",
			},
		},
	}
	commandWithoutArg := &core.Command{
		Namespace: "second",
		Resource:  "command",
		ArgSpecs: core.ArgSpecs{
			{
				Name: "other-arg",
			},
		},
	}

	testAlias := alias.Alias{
		Name:    "alias",
		Command: []string{"command"},
	}

	assert.True(t, commandWithArg.MatchAlias(testAlias))
	assert.True(t, commandWithoutArg.MatchAlias(testAlias))

	testAliasWithArg := alias.Alias{
		Name:    "alias",
		Command: []string{"command", "arg=value"},
	}

	assert.True(t, commandWithArg.MatchAlias(testAliasWithArg))
	assert.False(t, commandWithoutArg.MatchAlias(testAliasWithArg))
}

func TestAliasChildCommand(t *testing.T) {
	namespace := &core.Command{
		Namespace: "namespace",
	}
	resource := &core.Command{
		Namespace: "namespace",
		Resource:  "first",
	}

	commands := core.NewCommands(
		namespace,
		resource,
	)

	validAlias := alias.Alias{
		Name:    "alias",
		Command: []string{"namespace", "first"},
	}

	assert.True(t, commands.AliasIsValidCommandChild(namespace, validAlias))

	invalidAlias := alias.Alias{
		Name:    "alias",
		Command: []string{"namespace", "random"},
	}

	assert.False(t, commands.AliasIsValidCommandChild(namespace, invalidAlias))
}

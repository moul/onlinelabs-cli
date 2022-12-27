package namespaces

import (
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	accountv2 "github.com/scaleway/scaleway-cli/v2/internal/namespaces/account/v2"
	account "github.com/scaleway/scaleway-cli/v2/internal/namespaces/account/v2alpha1"
	applesilicon "github.com/scaleway/scaleway-cli/v2/internal/namespaces/applesilicon/v1alpha1"
	autocompleteNamespace "github.com/scaleway/scaleway-cli/v2/internal/namespaces/autocomplete"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/baremetal/v1"
	configNamespace "github.com/scaleway/scaleway-cli/v2/internal/namespaces/config"
	container "github.com/scaleway/scaleway-cli/v2/internal/namespaces/container/v1beta1"
	domain "github.com/scaleway/scaleway-cli/v2/internal/namespaces/domain/v2beta1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/feedback"
	flexibleip "github.com/scaleway/scaleway-cli/v2/internal/namespaces/flexibleip/v1alpha1"
	function "github.com/scaleway/scaleway-cli/v2/internal/namespaces/function/v1beta1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/help"
	iam "github.com/scaleway/scaleway-cli/v2/internal/namespaces/iam/v1alpha1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/info"
	initNamespace "github.com/scaleway/scaleway-cli/v2/internal/namespaces/init"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/instance/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/iot/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/k8s/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/lb/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/marketplace/v1"
	mnq "github.com/scaleway/scaleway-cli/v2/internal/namespaces/mnq/v1alpha1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/object/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/rdb/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/redis/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/registry/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/shell"
	tem "github.com/scaleway/scaleway-cli/v2/internal/namespaces/tem/v1alpha1"
	versionNamespace "github.com/scaleway/scaleway-cli/v2/internal/namespaces/version"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/vpc/v1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/vpcgw/v1"
)

// GetCommands returns a list of all commands in the CLI.
// It is used by both scw and scw-qa.
// We can not put it in `core` package as it would result in a import cycle `core` -> `namespaces/autocomplete` -> `core`.
func GetCommands(beta ...bool) *core.Commands {
	// Import all commands available in CLI from various packages.
	// NB: Merge order impacts scw usage sort.
	commands := core.NewCommandsMerge(
		iam.GetCommands(),
		instance.GetCommands(),
		baremetal.GetCommands(),
		k8s.GetCommands(),
		marketplace.GetCommands(),
		initNamespace.GetCommands(),
		configNamespace.GetCommands(),
		account.GetCommands(),
		accountv2.GetCommands(),
		autocompleteNamespace.GetCommands(),
		object.GetCommands(),
		versionNamespace.GetCommands(),
		registry.GetCommands(),
		feedback.GetCommands(),
		info.GetCommands(),
		rdb.GetCommands(),
		lb.GetCommands(),
		iot.GetCommands(),
		help.GetCommands(),
		vpc.GetCommands(),
		domain.GetCommands(),
		applesilicon.GetCommands(),
		flexibleip.GetCommands(),
		container.GetCommands(),
		function.GetCommands(),
		vpcgw.GetCommands(),
		redis.GetCommands(),
		shell.GetCommands(),
		tem.GetCommands(),
		mnq.GetCommands(),
	)

	return commands
}

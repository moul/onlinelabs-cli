package mongodb

import (
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/human"
	mongodb "github.com/scaleway/scaleway-sdk-go/api/mongodb/v1alpha1"
)

func GetCommands() *core.Commands {
	cmds := GetGeneratedCommands()

	human.RegisterMarshalerFunc(mongodb.SnapshotStatus(""), human.EnumMarshalFunc(snapshotStatusMarshalSpecs))
	human.RegisterMarshalerFunc(mongodb.InstanceStatus(""), human.EnumMarshalFunc(instanceStatusMarshalSpecs))
	human.RegisterMarshalerFunc(mongodb.NodeTypeStock(""), human.EnumMarshalFunc(nodeTypeStockMarshalSpecs))

	return cmds
}

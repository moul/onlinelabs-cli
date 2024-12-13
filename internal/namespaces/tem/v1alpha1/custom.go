package tem

import (
	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/human"
	tem "github.com/scaleway/scaleway-sdk-go/api/tem/v1alpha1"
)

func GetCommands() *core.Commands {
	cmds := GetGeneratedCommands()

	human.RegisterMarshalerFunc(tem.DomainStatus(""), human.EnumMarshalFunc(domainStatusMarshalSpecs))
	human.RegisterMarshalerFunc(tem.EmailStatus(""), human.EnumMarshalFunc(emailStatusMarshalSpecs))

	cmds.MustFind("tem", "domain", "get").Override(domainGetBuilder)

	return cmds
}

package instance

import (
	"fmt"
	"reflect"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-cli/internal/human"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
)

// GetCommands returns instance commands.
//
// This function:
// - Gets the generated commands
// - Register handwritten marshalers
// - Apply handwritten overrides (of Command.Run and Command.View)
// - Merge handwritten commands
func GetCommands() *core.Commands {
	cmds := GetGeneratedCommands()

	//
	// Server
	//
	human.RegisterMarshalerFunc(instance.CreateServerResponse{}, marshallNestedField("Server"))
	human.RegisterMarshalerFunc(instance.ServerState(0), serverStateMarshalerFunc)
	human.RegisterMarshalerFunc(instance.ServerLocation{}, serverLocationMarshalerFunc)
	human.RegisterMarshalerFunc([]*instance.Server{}, serversMarshalerFunc)
	human.RegisterMarshalerFunc(instance.GetServerResponse{}, getServerResponseMarshalerFunc)
	human.RegisterMarshalerFunc(instance.Bootscript{}, bootscriptMarshalerFunc)

	cmds.MustFind("instance", "server", "update").Override(serverUpdateBuilder)

	cmds.Merge(core.NewCommands(
		serverCreateCommand(),
		serverStartCommand(),
		serverStopCommand(),
		serverStandbyCommand(),
		serverRebootCommand(),
		serverDeleteCommand(),
	))

	//
	// Server-Type
	//
	cmds.MustFind("instance", "server-type", "list").Override(serverTypeListBuilder)

	//
	// IP
	//
	human.RegisterMarshalerFunc(instance.CreateIPResponse{}, marshallNestedField("IP"))

	cmds.MustFind("instance", "image", "list").Override(imageListBuilder)

	//
	// Image
	//
	human.RegisterMarshalerFunc(instance.CreateImageResponse{}, marshallNestedField("Image"))

	//
	// Snapshot
	//
	human.RegisterMarshalerFunc(instance.CreateSnapshotResponse{}, marshallNestedField("Snapshot"))

	//
	// Volume
	//
	human.RegisterMarshalerFunc(instance.CreateVolumeResponse{}, marshallNestedField("Volume"))
	human.RegisterMarshalerFunc(instance.VolumeState(0), human.BindAttributesMarshalFunc(volumeStateAttributes))
	human.RegisterMarshalerFunc(instance.VolumeSummary{}, volumeSummaryMarshalerFunc)
	human.RegisterMarshalerFunc(map[string]*instance.Volume{}, volumeMapMarshalerFunc)

	//
	// Security Group
	//
	human.RegisterMarshalerFunc(instance.CreateSecurityGroupResponse{}, marshallNestedField("SecurityGroup"))
	human.RegisterMarshalerFunc(instance.SecurityGroupPolicy(0), human.BindAttributesMarshalFunc(securityGroupPolicyAttribute))

	cmds.MustFind("instance", "security-group", "get").Override(securityGroupGetBuilder)
	cmds.MustFind("instance", "security-group", "delete").Override(securityGroupDeleteBuilder)

	cmds.Merge(core.NewCommands(
		securityGroupClearCommand(),
		securityGroupUpdateCommand(),
	))

	//
	// Security Group Rule
	//
	human.RegisterMarshalerFunc(instance.CreateSecurityGroupRuleResponse{}, marshallNestedField("Rule"))
	human.RegisterMarshalerFunc(instance.SecurityGroupRuleAction(0), human.BindAttributesMarshalFunc(securityGroupRuleActionAttribute))

	//
	// Placement Group
	//
	human.RegisterMarshalerFunc(instance.CreatePlacementGroupResponse{}, marshallNestedField("PlacementGroup"))

	cmds.MustFind("instance", "placement-group", "get").Override(placementGroupGetBuilder)

	//
	// User Data
	//
	cmds.Merge(core.NewCommands(
		userDataCommand(),
		userDataListCommand(),
		userDataSetCommand(),
		userDataDeleteCommand(),
		userDataGetCommand(),
	))

	return cmds
}

// marshallNestedField will marshal only the given field of a struct.
func marshallNestedField(nestedKey string) human.MarshalerFunc {
	return func(i interface{}, opt *human.MarshalOpt) (s string, err error) {
		if reflect.TypeOf(i).Kind() != reflect.Struct {
			return "", fmt.Errorf("%T must be a struct", i)
		}
		nestedValue := reflect.ValueOf(i).FieldByName(nestedKey)
		return human.Marshal(nestedValue.Interface(), opt)
	}
}

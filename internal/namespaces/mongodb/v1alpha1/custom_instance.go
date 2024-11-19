package mongodb

import (
	"github.com/fatih/color"
	"github.com/scaleway/scaleway-cli/v2/internal/human"
	mongodb "github.com/scaleway/scaleway-sdk-go/api/mongodb/v1alpha1"
)

var instanceStatusMarshalSpecs = human.EnumMarshalSpecs{
	mongodb.InstanceStatusConfiguring:  &human.EnumMarshalSpec{Attribute: color.FgBlue, Value: "configuring"},
	mongodb.InstanceStatusDeleting:     &human.EnumMarshalSpec{Attribute: color.FgBlue, Value: "deleting"},
	mongodb.InstanceStatusError:        &human.EnumMarshalSpec{Attribute: color.FgRed, Value: "error"},
	mongodb.InstanceStatusInitializing: &human.EnumMarshalSpec{Attribute: color.FgBlue, Value: "initializing"},
	mongodb.InstanceStatusLocked:       &human.EnumMarshalSpec{Attribute: color.FgRed, Value: "locked"},
	mongodb.InstanceStatusProvisioning: &human.EnumMarshalSpec{Attribute: color.FgBlue, Value: "provisioning"},
	mongodb.InstanceStatusReady:        &human.EnumMarshalSpec{Attribute: color.FgGreen, Value: "ready"},
	mongodb.InstanceStatusSnapshotting: &human.EnumMarshalSpec{Attribute: color.FgBlue, Value: "snapshotting"},
}

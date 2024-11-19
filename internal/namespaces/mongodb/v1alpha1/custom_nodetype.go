package mongodb

import (
	"github.com/fatih/color"
	"github.com/scaleway/scaleway-cli/v2/internal/human"
	mongodb "github.com/scaleway/scaleway-sdk-go/api/mongodb/v1alpha1"
)

var nodeTypeStockMarshalSpecs = human.EnumMarshalSpecs{
	mongodb.NodeTypeStockAvailable:  &human.EnumMarshalSpec{Attribute: color.FgGreen, Value: "available"},
	mongodb.NodeTypeStockLowStock:   &human.EnumMarshalSpec{Attribute: color.FgYellow, Value: "low stock"},
	mongodb.NodeTypeStockOutOfStock: &human.EnumMarshalSpec{Attribute: color.FgRed, Value: "out of stock"},
}

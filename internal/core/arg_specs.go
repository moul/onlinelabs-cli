package core

import (
	"context"
	"fmt"
	"strings"

	"github.com/scaleway/scaleway-sdk-go/scw"
)

type ArgSpecs []*ArgSpec

func (s ArgSpecs) GetPositionalArg() *ArgSpec {
	var positionalArg *ArgSpec
	for _, argSpec := range s {
		if argSpec.Positional {
			if positionalArg != nil {
				panic(fmt.Errorf("more than one positional parameter detected: %s and %s are flagged as positional arg", positionalArg.Name, argSpec.Name))
			}
			positionalArg = argSpec
		}
	}
	return positionalArg
}

func (s ArgSpecs) GetByName(name string) *ArgSpec {
	for _, spec := range s {
		if spec.Name == name {
			return spec
		}
	}
	return nil
}

func (s *ArgSpecs) DeleteByName(name string) {
	for i, spec := range *s {
		if spec.Name == name {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return
		}
	}
	panic(fmt.Errorf("in DeleteByName: %s not found", name))
}

func (s *ArgSpecs) AddBefore(name string, argSpec *ArgSpec) {
	for i, spec := range *s {
		if spec.Name == name {
			newSpecs := ArgSpecs(nil)
			newSpecs = append(newSpecs, (*s)[:i]...)
			newSpecs = append(newSpecs, argSpec)
			newSpecs = append(newSpecs, (*s)[i:]...)
			*s = newSpecs
			return
		}
	}
	panic(fmt.Errorf("in AddBefore: %s not found", name))
}

type ArgSpec struct {
	// Name of the argument.
	Name string

	// Short description.
	Short string

	// Required defines whether the argument is required.
	Required bool

	// Default is the argument default value.
	Default DefaultFunc

	// EnumValues contains all possible values of an enum.
	EnumValues []string

	// AutoCompleteFunc is used to autocomplete possible values for a given argument.
	AutoCompleteFunc AutoCompleteArgFunc

	// ValidateFunc validates an argument.
	ValidateFunc ArgSpecValidateFunc

	// Positional defines whether the argument is a positional argument. NB: a positional argument is required.
	Positional bool
}

func (a *ArgSpec) Prefix() string {
	return a.Name + "="
}

func (a *ArgSpec) IsPartOfMapOrSlice() bool {
	return strings.Contains(a.Name, sliceSchema) || strings.Contains(a.Name, mapSchema)
}

type DefaultFunc func(ctx context.Context) (value string, doc string)

func ZoneArgSpec(zones ...scw.Zone) *ArgSpec {
	enumValues := []string(nil)
	for _, zone := range zones {
		enumValues = append(enumValues, zone.String())
	}
	return &ArgSpec{
		Name:       "zone",
		Short:      "Zone to target. If none is passed will use default zone from the config",
		EnumValues: enumValues,
		Default: func(ctx context.Context) (value string, doc string) {
			client := ExtractClient(ctx)
			zone, _ := client.GetDefaultZone()
			return zone.String(), zone.String()
		},
	}
}

func RegionArgSpec(regions ...scw.Region) *ArgSpec {
	enumValues := []string(nil)
	for _, region := range regions {
		enumValues = append(enumValues, region.String())
	}
	return &ArgSpec{
		Name:       "region",
		Short:      "Region to target. If none is passed will use default region from the config",
		EnumValues: enumValues,
		Default: func(ctx context.Context) (value string, doc string) {
			client := ExtractClient(ctx)
			region, _ := client.GetDefaultRegion()
			return region.String(), region.String()
		},
	}
}

func OrganizationIDArgSpec() *ArgSpec {
	return &ArgSpec{
		Name:         "organization-id",
		Short:        "Organization ID to use. If none is passed will use default organization ID from the config",
		ValidateFunc: ValidateOrganizationID(),
	}
}

func OrganizationArgSpec() *ArgSpec {
	return &ArgSpec{
		Name:         "organization",
		Short:        "Organization ID to use. If none is passed will use default organization ID from the config",
		ValidateFunc: ValidateOrganizationID(),
	}
}

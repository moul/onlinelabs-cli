package config

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"reflect"

	"github.com/scaleway/scaleway-sdk-go/validation"

	"github.com/fatih/color"
	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-cli/internal/interactive"
	"github.com/scaleway/scaleway-cli/internal/tabwriter"
	"github.com/scaleway/scaleway-cli/internal/terminal"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/strcase"
)

func GetCommands() *core.Commands {
	return core.NewCommands(
		configRoot(),
		configGetCommand(),
		configSetCommand(),
		configUnsetCommand(),
		configDumpCommand(),
		configProfileCommand(),
		configDeleteProfileCommand(),
		configResetCommand(),
	)
}

func configRoot() *core.Command {
	configPath := scw.GetConfigPath()
	envVarTable := bytes.Buffer{}
	w := tabwriter.NewWriter(&envVarTable, 5, 1, 2, ' ', tabwriter.ANSIGraphicsRendition)
	for _, envVar := range [][2]string{
		{"SCW_ACCESS_KEY", "The access key of a token (create a token at https://console.scaleway.com/account/credentials)"},
		{"SCW_SECRET_KEY", "The secret key of a token (create a token at https://console.scaleway.com/account/credentials)"},
		{"SCW_DEFAULT_ORGANIZATION_ID", "The default organization ID (get your organization ID at https://console.scaleway.com/account/credentials)"},
		{"SCW_DEFAULT_REGION", "The default region"},
		{"SCW_DEFAULT_ZONE", "The default availability zone"},
		{"SCW_API_URL", "URL of the API"},
		{"SCW_INSECURE", "Set this to true to enable the insecure mode"},
		{"SCW_PROFILE", "Set the config profile to use"},
	} {
		fmt.Fprintf(w, "  %s\t%s\n", terminal.Style(envVar[0], color.Bold, color.FgBlue), envVar[1])
	}
	w.Flush()
	return &core.Command{
		Short: `Config file management`,
		Long: interactive.RemoveIndent(`
			Config management engine is common across all Scaleway developer tools (CLI, terraform, SDK, ... ). It allows to handle Scaleway config through two ways: environment variables and/or config file.

			Scaleway config file is self-documented. We recommend you to have a look at it at least once before using Scaleway developer tools: ` + terminal.Style(configPath, color.Bold, color.FgBlue) + `

			In this CLI, ` + terminal.Style(`environment variables have priority over the configuration file`, color.Bold) + `.

			The following environment variables are supported:
			` + envVarTable.String() + `
			Read more about the config management engine at https://github.com/scaleway/scaleway-sdk-go/tree/master/scw#scaleway-config
		`),
		Namespace: "config",
		SeeAlsos: []*core.SeeAlso{
			{
				Short:   "Init your Scaleway config",
				Command: "scw config init",
			},
			{
				Short:   "Set a config attribute",
				Command: "scw config set --help",
			},
			{
				Short:   "Set a config attribute",
				Command: "scw config get --help",
			},
			{
				Short:   "Dump the config",
				Command: "scw config dump",
			},
			{
				Short:   "Display the actual config file",
				Command: "cat " + configPath,
			},
		},
	}
}

// configGetCommand gets one or many values for the scaleway config
func configGetCommand() *core.Command {

	type configGetArgs struct {
		Key string
	}

	return &core.Command{
		Short:                `Get a value from the config file`,
		Namespace:            "config",
		Resource:             "get",
		AllowAnonymousClient: true,
		ArgsType:             reflect.TypeOf(configGetArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "key",
				Short:      "the key to get from the configt",
				Required:   true,
				EnumValues: getProfileKeys(),
				Positional: true,
			},
		},
		Examples: []*core.Example{
			{
				Short: "Get the default organization ID",
				Raw:   "scw config get default_organization_id",
			},
			{
				Short: "Get the default region of the profile 'prod'",
				Raw:   "scw -p prod config get default_region",
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Short:   "Config management help",
				Command: "scw config --help",
			},
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {

			config, err := scw.LoadConfigFromPath(extractConfigPath(ctx))
			if err != nil {
				return nil, err
			}
			key := argsI.(*configGetArgs).Key

			profileName := core.ExtractProfileName(ctx)
			profile, err := getProfile(config, profileName)
			if err != nil {
				return nil, err
			}

			return getProfileValue(profile, key)
		},
	}
}

// configSetCommand sets a value for the scaleway config
func configSetCommand() *core.Command {
	allRegions := []string(nil)
	for _, region := range scw.AllRegions {
		allRegions = append(allRegions, region.String())
	}
	allZones := []string(nil)
	for _, zone := range scw.AllZones {
		allZones = append(allZones, zone.String())
	}

	return &core.Command{
		Short: `Set a line from the config file`,
		Long: `This commands overwrites the configuration file parameters with user input.
The only allowed attributes are access_key, secret_key, default_organization_id, default_region, default_zone, api_url, insecure`,
		Namespace:            "config",
		Resource:             "set",
		AllowAnonymousClient: true,
		ArgsType:             reflect.TypeOf(scw.Profile{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:  "access-key",
				Short: "A Scaleway access key",
				ValidateFunc: func(argSpec *core.ArgSpec, value interface{}) error {
					if !reflect.ValueOf(value).IsNil() && !validation.IsAccessKey(*value.(*string)) {
						return core.InvalidAccessKeyError(*value.(*string))
					}
					return nil
				},
			},
			{
				Name:  "secret-key",
				Short: "A Scaleway secret key",
				ValidateFunc: func(argSpec *core.ArgSpec, value interface{}) error {
					if !reflect.ValueOf(value).IsNil() && !validation.IsSecretKey(*value.(*string)) {
						return core.InvalidSecretKeyError(*value.(*string))
					}
					return nil
				},
			},
			{
				Name:  "api-url",
				Short: "Scaleway API URL",
				ValidateFunc: func(argSpec *core.ArgSpec, value interface{}) error {
					if !reflect.ValueOf(value).IsNil() && !validation.IsURL(*value.(*string)) {
						return fmt.Errorf("%s is not a valid URL", *value.(*string))
					}
					return nil
				},
			},
			{
				Name:  "insecure",
				Short: "Set to true to allow insecure HTTPS connections",
			},
			{
				Name:  "default-organization-id",
				Short: "A default Scaleway organization id",
				ValidateFunc: func(argSpec *core.ArgSpec, value interface{}) error {
					if !reflect.ValueOf(value).IsNil() && !validation.IsOrganizationID(*value.(*string)) {
						return core.InvalidOrganizationIDError(*value.(*string))
					}
					return nil
				},
			},
			{
				Name:       "default-region",
				Short:      "A default Scaleway region",
				EnumValues: allZones,
			},
			{
				Name:       "default-zone",
				Short:      "A default Scaleway zone",
				EnumValues: allZones,
			},
			{
				Name:  "send-telemetry",
				Short: "Set to false to disable telemetry",
			},
		},
		Examples: []*core.Example{
			{
				Short: "Update the default organization ID",
				Raw:   "scw config set default_organization_id=12903058-d0e8-4366-89c3-6e666abe1f6f",
			},
			{
				Short: "Update the default region of the profile 'prod'",
				Raw:   "scw -p prod config set default_region=nl-ams",
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Short:   "Config management help",
				Command: "scw config --help",
			},
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, err error) {

			// Validate arguments
			args := argsI.(*scw.Profile)

			// Execute
			configPath := extractConfigPath(ctx)
			config, err := scw.LoadConfigFromPath(configPath)
			if err != nil {
				return nil, err
			}

			// send_telemetry is the only key that is not in a profile but in the config object directly
			profileName := core.ExtractProfileName(ctx)
			profile := &config.Profile
			if profileName != "" {
				var exist bool
				profile, exist = config.Profiles[profileName]
				if !exist {
					if config.Profiles == nil {
						config.Profiles = map[string]*scw.Profile{}
					}
					config.Profiles[profileName] = &scw.Profile{}
					profile = config.Profiles[profileName]
				}
			}

			argValue := reflect.ValueOf(args).Elem()
			profileValue := reflect.ValueOf(profile).Elem()
			for i := 0; i < argValue.NumField(); i++ {
				field := argValue.Field(i)
				if !field.IsNil() {
					profileValue.Field(i).Set(field)
				}
			}

			// Save
			err = config.SaveTo(configPath)
			if err != nil {
				return nil, err
			}

			return &core.SuccessResult{
				Message: fmt.Sprintf("successfully update config"),
			}, nil
		},
	}
}

// configDumpCommand unsets a value for the scaleway config
func configUnsetCommand() *core.Command {

	type configUnsetArgs struct {
		Key string
	}

	return &core.Command{
		Short:                `Unset a line from the config file`,
		Namespace:            "config",
		Resource:             "unset",
		AllowAnonymousClient: true,
		ArgsType:             reflect.TypeOf(configUnsetArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "key",
				Short:      "the config config key name to unset",
				Required:   true,
				EnumValues: getProfileKeys(),
				Positional: true,
			},
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			configPath := extractConfigPath(ctx)
			config, err := scw.LoadConfigFromPath(configPath)
			if err != nil {
				return nil, err
			}
			key := argsI.(*configUnsetArgs).Key

			profileName := core.ExtractProfileName(ctx)
			profile, err := getProfile(config, profileName)
			if err != nil {
				return nil, err
			}
			err = unsetProfileValue(profile, key)
			if err != nil {
				return nil, err
			}

			err = config.SaveTo(configPath)
			if err != nil {
				return nil, err
			}

			return &core.SuccessResult{
				Message: fmt.Sprintf("successfully unset %s", key),
			}, nil
		},
	}
}

// configDumpCommand dumps the scaleway config
func configDumpCommand() *core.Command {

	type configDumpArgs struct{}

	return &core.Command{
		Short:                `Dump the config file`,
		Namespace:            "config",
		Resource:             "dump",
		AllowAnonymousClient: true,
		ArgsType:             reflect.TypeOf(configDumpArgs{}),
		SeeAlsos: []*core.SeeAlso{
			{
				Short:   "Config management help",
				Command: "scw config --help",
			},
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			configPath := extractConfigPath(ctx)
			config, err := scw.LoadConfigFromPath(configPath)
			if err != nil {
				return nil, err
			}
			return config, nil
		},
	}
}

func configProfileCommand() *core.Command {
	return &core.Command{
		Short:                `Allows the deletion of a profile from the config file`,
		Namespace:            "config",
		Resource:             "profile",
		AllowAnonymousClient: true,
	}
}

// configDeleteProfileCommand deletes a profile from the config
func configDeleteProfileCommand() *core.Command {

	type configDeleteProfileArgs struct {
		Name string
	}

	return &core.Command{
		Short:                `Delete a profile from the config file`,
		Namespace:            "config",
		Resource:             "profile",
		Verb:                 "delete",
		AllowAnonymousClient: true,
		ArgsType:             reflect.TypeOf(configDeleteProfileArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Required:   true,
				Positional: true,
			},
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			profileName := argsI.(*configDeleteProfileArgs).Name
			configPath := extractConfigPath(ctx)
			config, err := scw.LoadConfigFromPath(configPath)
			if err != nil {
				return nil, err
			}
			if _, exists := config.Profiles[profileName]; exists {
				delete(config.Profiles, profileName)
			} else {
				return nil, unknownProfileError(profileName)
			}
			err = config.SaveTo(configPath)
			if err != nil {
				return nil, err
			}

			return &core.SuccessResult{
				Message: fmt.Sprintf("successfully delete profile %s", profileName),
			}, nil
		},
	}
}

// configResetCommand resets the config
func configResetCommand() *core.Command {

	type configResetArgs struct{}

	return &core.Command{
		Short:                `Reset the config`,
		Namespace:            "config",
		Resource:             "reset",
		AllowAnonymousClient: true,
		ArgsType:             reflect.TypeOf(configResetArgs{}),
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			_, err := scw.LoadConfig()
			if err != nil {
				return nil, err
			}
			config := &scw.Config{}
			err = config.Save()
			if err != nil {
				return nil, err
			}
			return &core.SuccessResult{
				Message: "successfully reset config",
			}, nil
		},
	}
}

//
// Helper functions
//
func getProfileValue(profile *scw.Profile, fieldName string) (interface{}, error) {
	field, err := getProfileField(profile, fieldName)
	if err != nil {
		return nil, err
	}
	return field.Interface(), nil
}

func unsetProfileValue(profile *scw.Profile, key string) error {
	field, err := getProfileField(profile, key)
	if err != nil {
		return err
	}
	field.Set(reflect.Zero(field.Type()))
	return nil
}

func getProfileField(profile *scw.Profile, key string) (reflect.Value, error) {
	field := reflect.ValueOf(profile).Elem().FieldByName(strcase.ToPublicGoName(key))
	if !field.IsValid() {
		return reflect.ValueOf(nil), invalidProfileKeyError(key)
	}
	return field, nil
}

func getProfileKeys() []string {
	t := reflect.TypeOf(scw.Profile{})
	keys := []string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		switch field.Name {
		case "APIURL":
			keys = append(keys, "api-url")
		default:
			keys = append(keys, strcase.ToBashArg(t.Field(i).Name))
		}
	}
	return keys
}

// This func should be removes when core implement it
func extractConfigPath(ctx context.Context) string {
	homeDir := core.ExtractUserHomeDir(ctx)
	return path.Join(homeDir, ".config", "scw", "config.yaml")
}

func getProfile(config *scw.Config, profileName string) (*scw.Profile, error) {
	if profileName == "" {
		return &config.Profile, nil
	}
	profile, exist := config.Profiles[profileName]
	if !exist {
		return nil, unknownProfileError(profileName)
	}
	return profile, nil
}

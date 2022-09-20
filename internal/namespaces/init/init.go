package init

import (
	"context"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/scaleway/scaleway-cli/v2/internal/account"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-cli/v2/internal/interactive"
	accountcommands "github.com/scaleway/scaleway-cli/v2/internal/namespaces/account/v2alpha1"
	"github.com/scaleway/scaleway-cli/v2/internal/namespaces/autocomplete"
	"github.com/scaleway/scaleway-cli/v2/internal/terminal"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/validation"
)

/*
See below the schema `scw init` follows to ask for default config:

                 yes   +----------+
               +-------+Config ok?|
               |       +----------+
+---+  no +----v----+       |no
|out+<----+Override?|       v
+---+     +----+----+  +----+-----+
               |       |Read      +-----------+
               +------>+    token |  token    |
                 yes   +----------+           |
                                              |
                                              v
                                       +------+---+
                                       |Get access|
                                       |   key    |
                                       +------+---+
                                              |
                                              |
                                              |
                                              |
                                              |
                                              |
                                              |
                                              |
                                              |
                                              |
                                              |
                                              |
                    +-------+----------+      |
                    |ask default config+<-----+
                    +------------------+
*/

func GetCommands() *core.Commands {
	return core.NewCommands(initCommand())
}

type initArgs struct {
	AccessKey           string
	SecretKey           string
	Region              scw.Region
	Zone                scw.Zone
	SendTelemetry       *bool
	WithSSHKey          *bool
	InstallAutocomplete *bool
	RemoveV1Config      *bool
}

func initCommand() *core.Command {
	return &core.Command{
		Short: `Initialize the config`,
		Long: `Initialize the active profile of the config.
Default path for configuration file is based on the following priority order:

- $SCW_CONFIG_PATH
- $XDG_CONFIG_HOME/scw/config.yaml
- $HOME/.config/scw/config.yaml
- $USERPROFILE/.config/scw/config.yaml`,
		Namespace:            "init",
		AllowAnonymousClient: true,
		ArgsType:             reflect.TypeOf(initArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:         "secret-key",
				Short:        "Scaleway secret-key",
				ValidateFunc: core.ValidateSecretKey(),
			},
			{
				Name:  "send-telemetry",
				Short: "Send usage statistics and diagnostics",
			},
			{
				Name:    "with-ssh-key",
				Short:   "Whether the SSH key for managing instances should be uploaded automatically",
				Default: core.DefaultValueSetter("true"),
			},
			{
				Name:  "install-autocomplete",
				Short: "Whether the autocomplete script should be installed during initialisation",
			},
			{
				Name:  "remove-v1-config",
				Short: "Whether to remove the v1 configuration file if it exists",
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
			core.ZoneArgSpec(scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1),
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Short:   "Config management help",
				Command: "scw config --help",
			},
		},
		PreValidateFunc: func(ctx context.Context, argsI interface{}) error {
			args := argsI.(*initArgs)

			// Show logo banner, or simple welcome message
			if terminal.GetWidth() >= 80 {
				interactive.Printf("%s\n%s\n\n", interactive.Center(logo), interactive.Line("-"))
			} else {
				interactive.Printf("Welcome to the Scaleway Cli\n\n")
			}

			config, err := scw.LoadConfigFromPath(core.ExtractConfigPath(ctx))

			// If it is not a new config, ask if we want to override the existing config
			if err == nil && !config.IsEmpty() {
				_, _ = interactive.PrintlnWithoutIndent(`
					Current config is located at ` + core.ExtractConfigPath(ctx) + `
					` + terminal.Style(fmt.Sprint(config), color.Faint) + `
				`)
				overrideConfig, err := interactive.PromptBoolWithConfig(&interactive.PromptBoolConfig{
					Prompt:       "Do you want to override the current config?",
					DefaultValue: true,
					Ctx:          ctx,
				})
				if err != nil {
					return err
				}
				if !overrideConfig {
					return fmt.Errorf("initialization canceled")
				}
			}

			// Manually prompt for missing args:

			// Credentials
			if args.SecretKey == "" {
				_, _ = interactive.Println()
				args.SecretKey, err = promptSecret(ctx)
				if err != nil {
					return err
				}
			}

			// Zone
			if args.Zone == "" {
				_, _ = interactive.Println()
				zone, err := interactive.PromptStringWithConfig(&interactive.PromptStringConfig{
					Ctx:             ctx,
					Prompt:          "Select a zone",
					DefaultValueDoc: "fr-par-1",
					DefaultValue:    "fr-par-1",
					ValidateFunc: func(s string) error {
						logger.Debugf("s: %v", s)
						if !validation.IsZone(s) {
							return fmt.Errorf("invalid zone")
						}
						return nil
					},
				})
				if err != nil {
					return err
				}
				args.Zone, err = scw.ParseZone(zone)
				if err != nil {
					return err
				}
			}

			// Deduce Region from Zone
			if args.Region == "" {
				args.Region, err = args.Zone.Region()
				if err != nil {
					return err
				}
			}

			// Ask for send usage permission
			if args.SendTelemetry == nil {
				_, _ = interactive.Println()
				_, _ = interactive.PrintlnWithoutIndent(`
					To improve this tool we rely on diagnostic and usage data.
					Sending such data is optional and can be disabled at any time by running "scw config set send-telemetry=false".
				`)

				sendTelemetry, err := interactive.PromptBoolWithConfig(&interactive.PromptBoolConfig{
					Prompt:       "Do you want to send usage statistics and diagnostics?",
					DefaultValue: true,
					Ctx:          ctx,
				})
				if err != nil {
					return err
				}

				args.SendTelemetry = scw.BoolPtr(sendTelemetry)
			}

			// Ask whether we should install autocomplete
			if args.InstallAutocomplete == nil {
				_, _ = interactive.Println()
				_, _ = interactive.PrintlnWithoutIndent(`
					To fully enjoy Scaleway CLI we recommend you install autocomplete support in your shell.
				`)

				installAutocomplete, err := interactive.PromptBoolWithConfig(&interactive.PromptBoolConfig{
					Ctx:          ctx,
					Prompt:       "Do you want to install autocomplete?",
					DefaultValue: true,
				})
				if err != nil {
					return err
				}

				args.InstallAutocomplete = scw.BoolPtr(installAutocomplete)
			}

			// Ask whether to remove v1 configuration file if it exists
			if args.RemoveV1Config == nil {
				homeDir := core.ExtractUserHomeDir(ctx)
				if err == nil {
					configPath := path.Join(homeDir, ".scwrc")
					if _, err := os.Stat(configPath); err == nil {
						removeV1ConfigFile, err := interactive.PromptBoolWithConfig(&interactive.PromptBoolConfig{
							Ctx:          ctx,
							Prompt:       "Do you want to permanently remove old configuration file (" + configPath + ")?",
							DefaultValue: false,
						})
						if err != nil {
							return err
						}

						args.RemoveV1Config = &removeV1ConfigFile
					}
				}
			}

			return nil
		},
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			args := argsI.(*initArgs)
			// Check if a config exists
			// Creates a new one if it does not
			configPath := core.ExtractConfigPath(ctx)
			config, err := scw.LoadConfigFromPath(configPath)
			if err != nil {
				config = &scw.Config{}
				interactive.Printf("Creating new config at %s\n", configPath)
			}

			if args.SendTelemetry != nil {
				config.SendTelemetry = args.SendTelemetry
			}

			// Get access key
			apiKey, err := account.GetAPIKey(ctx, args.SecretKey)
			if err != nil {
				return "", &core.CliError{
					Err:     err,
					Details: "Failed to retrieve Access Key from the given Secret Key.",
				}
			}

			profile := &scw.Profile{
				AccessKey:             &apiKey.AccessKey,
				SecretKey:             &args.SecretKey,
				DefaultZone:           scw.StringPtr(args.Zone.String()),
				DefaultRegion:         scw.StringPtr(args.Region.String()),
				DefaultOrganizationID: &apiKey.OrganizationID,
				DefaultProjectID:      &apiKey.ProjectID, // An API key is always bound to a project.
			}

			// Save the profile as default or as a named profile
			profileName := core.ExtractProfileName(ctx)
			if profileName == scw.DefaultProfileName {
				// Default configuration
				config.Profile = *profile
			} else {
				if config.Profiles == nil {
					config.Profiles = make(map[string]*scw.Profile)
				}
				config.Profiles[profileName] = profile
			}

			// Persist configuration on disk
			interactive.Printf("Config saved at %s:\n%s\n", configPath, terminal.Style(fmt.Sprint(config), color.Faint))
			err = config.SaveTo(configPath)
			if err != nil {
				return nil, err
			}

			// Now that the config has been save we reload the client with the new config
			err = core.ReloadClient(ctx)
			if err != nil {
				return nil, err
			}
			successDetails := []string(nil)

			// Install autocomplete
			if *args.InstallAutocomplete {
				_, _ = interactive.Println()
				_, err := autocomplete.InstallCommandRun(ctx, &autocomplete.InstallArgs{})
				if err != nil {
					successDetails = append(successDetails, "Except for autocomplete: "+err.Error())
				}
			}

			// Init SSH Key
			if *args.WithSSHKey {
				_, _ = interactive.Println()
				_, err := accountcommands.InitRun(ctx, nil)
				if err != nil {
					successDetails = append(successDetails, "Except for SSH key: "+err.Error())
				}
			}

			// Remove old configuration file
			if args.RemoveV1Config != nil && *args.RemoveV1Config {
				homeDir := core.ExtractUserHomeDir(ctx)
				err = os.Remove(path.Join(homeDir, ".scwrc"))
				if err != nil {
					successDetails = append(successDetails, "Except for removing old configuration: "+err.Error())
				}
			}

			_, _ = interactive.Println()

			return &core.SuccessResult{
				Message: "Initialization completed with success",
				Details: strings.Join(successDetails, "\n"),
			}, nil
		},
	}
}

func promptSecret(ctx context.Context) (string, error) {
	secret, err := interactive.Readline(&interactive.ReadlineConfig{
		Ctx: ctx,
		PromptFunc: func(value string) string {
			secretKey := "secret-key"
			switch {
			case validation.IsUUID(value):
				secretKey = terminal.Style(secretKey, color.FgBlue)
			}
			return terminal.Style(fmt.Sprintf("Enter a valid %s: ", secretKey), color.Bold)
		},
		ValidateFunc: func(s string) error {
			if validation.IsSecretKey(s) {
				return nil
			}
			return fmt.Errorf("invalid secret-key")
		},
	})
	if err != nil {
		return "", err
	}

	switch {
	case validation.IsUUID(secret):
		return secret, nil

	default:
		return "", fmt.Errorf("invalid secret-key: '%v'", secret)
	}
}

const logo = `
  @@@@@@@@@@@@@@@.
@@@@@@@@@@@@@@@@@@@@        __          __  _
@@@               @@@@      \ \        / / | |
@@@    @@@@@@@     .@@@      \ \  /\  / /__| | ___ ___  _ __ ___   ___
@@@   @@@@@@@@      @@@       \ \/  \/ / _ \ |/ __/ _ \| '_ ` + "`" + ` _ \ / _ \
@@@   @@@           @@@        \  /\  /  __/ | (_| (_) | | | | | |  __/
@@@   @@@     @@@   @@@         \/  \/ \___|_|\___\___/|_| |_| |_|\___|
@@@   @@@     @@@   @@@                                         _  _
@@@           @@@   @@@                                        | |(_)
@@@      .@@@@@@@   @@@             ___   ___ __      __   ___ | | _
@@@      @@@@@@@    @@@            / __| / __|\ \ /\ / /  / __|| || |
 @@@.               @@@            \__ \| (__  \ V  V /  | (__ | || |
  @@@@@@.         .@@@@            |___/ \___|  \_/\_/    \___||_||_|
     @@@@@@@@@@@@@@@@.
`

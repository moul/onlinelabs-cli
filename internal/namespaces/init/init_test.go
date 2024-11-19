package init_test

import (
	"fmt"
	"path"
	"regexp"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/scaleway/scaleway-cli/v2/core"
	initCLI "github.com/scaleway/scaleway-cli/v2/internal/namespaces/init" // alias required to not collide with go init func
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/stretchr/testify/require"
)

func checkConfig(check func(t *testing.T, ctx *core.CheckFuncCtx, config *scw.Config)) core.TestCheck {
	return func(t *testing.T, ctx *core.CheckFuncCtx) {
		t.Helper()
		homeDir := ctx.OverrideEnv["HOME"]
		config, err := scw.LoadConfigFromPath(path.Join(homeDir, ".config", "scw", "config.yaml"))
		require.NoError(t, err)
		check(t, ctx, config)
	}
}

func appendArgs(prefix string, args map[string]string) string {
	cmd := prefix
	for k, v := range args {
		cmd += fmt.Sprintf(" %s=%s", k, v)
	}
	return cmd
}

func beforeFuncSaveConfig(config *scw.Config) core.BeforeFunc {
	return func(ctx *core.BeforeFuncCtx) error {
		// Persist the dummy Config in the temp directory
		return config.SaveTo(path.Join(ctx.OverrideEnv["HOME"], ".config", "scw", "config.yaml"))
	}
}

func TestInit(t *testing.T) {
	defaultArgs := map[string]string{
		"access-key":           "{{ .AccessKey }}",
		"secret-key":           "{{ .SecretKey }}",
		"send-telemetry":       "true",
		"install-autocomplete": "false",
		"with-ssh-key":         "false",
		"organization-id":      "{{ .OrganizationID }}",
		"project-id":           "{{ .ProjectID }}",
	}

	t.Run("Simple", core.Test(&core.TestConfig{
		Commands:   initCLI.GetCommands(),
		BeforeFunc: baseBeforeFunc(),
		TmpHomeDir: true,
		Cmd:        appendArgs("scw init", defaultArgs),
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			checkConfig(func(t *testing.T, ctx *core.CheckFuncCtx, config *scw.Config) {
				t.Helper()
				secretKey, _ := ctx.Client.GetSecretKey()
				assert.Equal(t, secretKey, *config.SecretKey)
				assert.NotEmpty(t, *config.DefaultProjectID)
				assert.Equal(t, *config.DefaultProjectID, *config.DefaultProjectID)
			}),
		),
	}))

	t.Run("Configuration Path", core.Test(&core.TestConfig{
		Commands: initCLI.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			baseBeforeFunc(),
			func(ctx *core.BeforeFuncCtx) error {
				ctx.Meta["CONFIG_PATH"] = path.Join(ctx.Meta["HOME"].(string), "new_config_path.yml")
				return nil
			},
		),
		TmpHomeDir: true,
		Cmd:        appendArgs("scw -c {{ .CONFIG_PATH }} init", defaultArgs),
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			func(t *testing.T, ctx *core.CheckFuncCtx) {
				t.Helper()
				config, err := scw.LoadConfigFromPath(ctx.Meta["CONFIG_PATH"].(string))
				require.NoError(t, err)
				secretKey, _ := ctx.Client.GetSecretKey()
				assert.Equal(t, secretKey, *config.SecretKey)
			},
		),
	}))

	t.Run("Profile", core.Test(&core.TestConfig{
		Commands:   initCLI.GetCommands(),
		BeforeFunc: baseBeforeFunc(),
		Cmd:        appendArgs("scw -p foobar init", defaultArgs),
		Check: core.TestCheckCombine(
			core.TestCheckGolden(),
			checkConfig(func(t *testing.T, ctx *core.CheckFuncCtx, config *scw.Config) {
				t.Helper()
				secretKey, _ := ctx.Client.GetSecretKey()
				assert.Equal(t, secretKey, *config.Profiles["foobar"].SecretKey)
			}),
		),
		TmpHomeDir: true,
	}))

	t.Run("CLIv2Config", func(t *testing.T) {
		dummySecretKey := "22222222-2222-2222-2222-222222222222"
		dummyAccessKey := "SCW22222222222222222"
		dummyConfig := &scw.Config{
			Profile: scw.Profile{
				AccessKey: &dummyAccessKey,
				SecretKey: &dummySecretKey,
			},
			Profiles: map[string]*scw.Profile{
				"test": {
					AccessKey:   &dummyAccessKey,
					SecretKey:   &dummySecretKey,
					DefaultZone: scw.StringPtr("fr-test"), // Used to check profile override
				},
			},
		}

		t.Run("NoOverwrite", core.Test(&core.TestConfig{
			Commands: initCLI.GetCommands(),
			BeforeFunc: core.BeforeFuncCombine(
				baseBeforeFunc(),
				beforeFuncSaveConfig(dummyConfig),
			),
			Cmd: appendArgs("scw init", defaultArgs),
			Check: core.TestCheckCombine(
				core.TestCheckGolden(),
				checkConfig(func(t *testing.T, _ *core.CheckFuncCtx, config *scw.Config) {
					t.Helper()
					assert.Equal(t, dummyConfig.String(), config.String())
				}),
			),
			TmpHomeDir: true,
			PromptResponseMocks: []string{
				// Do you want to override the current config?
				"no",
			},
		}))

		t.Run("Overwrite", core.Test(&core.TestConfig{
			Commands: initCLI.GetCommands(),
			BeforeFunc: core.BeforeFuncCombine(
				baseBeforeFunc(),
				beforeFuncSaveConfig(dummyConfig),
			),
			Cmd: appendArgs("scw init", defaultArgs),
			Check: core.TestCheckCombine(
				core.TestCheckGolden(),
				checkConfig(func(t *testing.T, ctx *core.CheckFuncCtx, config *scw.Config) {
					t.Helper()
					secretKey, _ := ctx.Client.GetSecretKey()
					assert.Equal(t, secretKey, *config.SecretKey)
				}),
			),
			TmpHomeDir: true,
			PromptResponseMocks: []string{
				// Do you want to override the current config?
				"yes",
			},
		}))

		t.Run("No Prompt Overwrite for new profile", core.Test(&core.TestConfig{
			Commands: initCLI.GetCommands(),
			BeforeFunc: core.BeforeFuncCombine(
				baseBeforeFunc(),
				beforeFuncSaveConfig(dummyConfig),
			),
			Cmd: appendArgs("scw -p test2 init", defaultArgs),
			Check: core.TestCheckCombine(
				core.TestCheckGolden(),
				checkConfig(func(t *testing.T, _ *core.CheckFuncCtx, config *scw.Config) {
					t.Helper()
					assert.NotNil(t, config.Profiles["test2"], "new profile should have been created")
				}),
			),
			TmpHomeDir: true,
			PromptResponseMocks: []string{
				// Do you want to override the current config? (Should not be prompted as profile is a new one)
				"no",
			},
		}))

		t.Run("Prompt Overwrite for existing profile", core.Test(&core.TestConfig{
			Commands: initCLI.GetCommands(),
			BeforeFunc: core.BeforeFuncCombine(
				baseBeforeFunc(),
				beforeFuncSaveConfig(dummyConfig),
			),
			Cmd: appendArgs("scw -p test init", defaultArgs),
			Check: core.TestCheckCombine(
				core.TestCheckGolden(),
				checkConfig(func(t *testing.T, _ *core.CheckFuncCtx, config *scw.Config) {
					t.Helper()
					assert.NotNil(t, config.Profiles["test"].DefaultZone)
					assert.Equal(t, *config.Profiles["test"].DefaultZone, "fr-test")
				}),
			),
			TmpHomeDir: true,
			PromptResponseMocks: []string{
				// Do you want to override the current config? (Should not be prompted as profile is a new one)
				"no",
			},
		}))

		t.Run("Default profile activated", core.Test(&core.TestConfig{
			Commands:   initCLI.GetCommands(),
			BeforeFunc: baseBeforeFunc(),
			TmpHomeDir: true,
			Cmd:        appendArgs("scw -p newprofile init", defaultArgs),
			Check: core.TestCheckCombine(
				core.TestCheckGolden(),
				checkConfig(func(t *testing.T, _ *core.CheckFuncCtx, config *scw.Config) {
					t.Helper()
					assert.NotNil(t, config.ActiveProfile)
					assert.Equal(t, "newprofile", *config.ActiveProfile)
				}),
			),
		}))
	})
}

func TestInit_Prompt(t *testing.T) {
	promptResponse := []string{
		"secret-key",      // Secret key prompt, should be replaced in BeforeFunc.
		"access-key",      // Access key prompt, should be replaced in BeforeFunc.
		"organization-id", // Organization prompt, should be replaced in BeforeFunc.
		" ",               // default-project-id list prompt, space is validation, it will pick default organization project.
		"",                // Telemetry prompt, use default value.
		"y",               // Autocomplete prompt, enable it but the tests should override a SHELL variable to avoid breaking because of local configuration.
	}

	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: initCLI.GetCommands(),
		BeforeFunc: core.BeforeFuncCombine(
			baseBeforeFunc(),
			func(ctx *core.BeforeFuncCtx) error {
				promptResponse[0] = ctx.Meta["SecretKey"].(string)
				promptResponse[1] = ctx.Meta["AccessKey"].(string)
				promptResponse[2] = ctx.Meta["OrganizationID"].(string)

				return nil
			}),
		TmpHomeDir: true,
		Cmd:        "scw init",
		Check: core.TestCheckCombine(
			core.TestCheckGoldenAndReplacePatterns(
				core.GoldenReplacement{
					Pattern:       regexp.MustCompile("\\s\\sExcept for autocomplete: unsupported OS 'windows'\n"),
					Replacement:   "",
					OptionalMatch: true,
				},
				core.GoldenReplacement{
					Pattern:       regexp.MustCompile(`Except for autocomplete: unsupported OS 'windows'\\n`),
					Replacement:   "",
					OptionalMatch: true,
				},
			),
			checkConfig(func(t *testing.T, ctx *core.CheckFuncCtx, config *scw.Config) {
				t.Helper()
				secretKey, _ := ctx.Client.GetSecretKey()
				assert.Equal(t, secretKey, *config.SecretKey)
				assert.NotEmpty(t, *config.DefaultProjectID)
				assert.Equal(t, *config.DefaultProjectID, *config.DefaultProjectID)
			}),
		),
		OverrideEnv: map[string]string{
			"SHELL": "/bin/bash",
		},
		PromptResponseMocks: promptResponse,
	}))
}

package core_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/scaleway/scaleway-cli/v2/internal/core"

	"github.com/scaleway/scaleway-cli/v2/internal/args"
)

type testType struct {
	NameID string
	Tag    string
}

type testDate struct {
	Date *time.Time
}

func testGetCommands() *core.Commands {
	return core.NewCommands(
		&core.Command{
			Namespace: "test",
			ArgSpecs: core.ArgSpecs{
				{
					Name: "name-id",
				},
			},
			AllowAnonymousClient: true,
			ArgsType:             reflect.TypeOf(testType{}),
			Run: func(_ context.Context, _ interface{}) (i interface{}, e error) {
				return "", nil
			},
		},
		&core.Command{
			Namespace: "test",
			Resource:  "positional",
			ArgSpecs: core.ArgSpecs{
				{
					Name:       "name-id",
					Positional: true,
				},
				{
					Name: "tag",
				},
			},
			AllowAnonymousClient: true,
			ArgsType:             reflect.TypeOf(testType{}),
			Run: func(_ context.Context, argsI interface{}) (i interface{}, e error) {
				return argsI, nil
			},
		},
		&core.Command{
			Namespace:            "test",
			Resource:             "raw-args",
			ArgsType:             reflect.TypeOf(args.RawArgs{}),
			AllowAnonymousClient: true,
			Run: func(_ context.Context, argsI interface{}) (i interface{}, e error) {
				res := ""
				rawArgs := *argsI.(*args.RawArgs)
				for i, arg := range rawArgs {
					res += arg
					if i != len(rawArgs)-1 {
						res += " "
					}
				}
				return res, nil
			},
		},
		&core.Command{
			Namespace:            "test",
			Resource:             "date",
			ArgsType:             reflect.TypeOf(testDate{}),
			AllowAnonymousClient: true,
			Run: func(_ context.Context, argsI interface{}) (i interface{}, e error) {
				a := argsI.(*testDate)
				return a.Date, nil
			},
		},
	)
}

func Test_handleUnmarshalErrors(t *testing.T) {
	t.Run("underscore", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test name_id",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(1),
			core.TestCheckError(&core.CliError{
				Err:  fmt.Errorf("invalid argument 'name_id': arg name must only contain lowercase letters, numbers or dashes"),
				Hint: "Valid arguments are: name-id",
			}),
		),
	}))

	t.Run("value only", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test ubuntu_focal",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(1),
			core.TestCheckError(&core.CliError{
				Err:  fmt.Errorf("invalid argument 'ubuntu_focal': arg name must only contain lowercase letters, numbers or dashes"),
				Hint: "Valid arguments are: name-id",
			}),
		),
	}))

	t.Run("relative date", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test date date=+3R",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(1),
			core.TestCheckError(&core.CliError{
				Message: "could not parse +3R as either an absolute time (RFC3339) nor a relative time (+/-)RFC3339",
				Details: `Absolute time error: parsing time "+3R" as "2006-01-02T15:04:05Z07:00": cannot parse "+3R" as "2006"
Relative time error: unknown unit in duration: "R"
`,
				Err:  fmt.Errorf("date parsing error: +3R"),
				Hint: "Run `scw help date` to learn more about date parsing",
			}),
		),
	}))
}

func Test_RawArgs(t *testing.T) {
	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test raw-args -- blabla",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(0),
			core.TestCheckStdout("blabla\n"),
		),
	}))
	t.Run("Multiple", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test raw-args -- blabla foo bar",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(0),
			core.TestCheckStdout("blabla foo bar\n"),
		),
	}))
}

func Test_PositionalArg(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		t.Run("Missing1", core.Test(&core.TestConfig{
			Commands: testGetCommands(),
			Cmd:      "scw test positional",
			Check: core.TestCheckCombine(
				core.TestCheckExitCode(1),
				core.TestCheckError(&core.CliError{
					Err:  fmt.Errorf("a positional argument is required for this command"),
					Hint: "Try running: scw test positional <name-id>",
				}),
			),
		}))

		t.Run("Missing2", core.Test(&core.TestConfig{
			Commands: testGetCommands(),
			Cmd:      "scw test positional tag=world",
			Check: core.TestCheckCombine(
				core.TestCheckExitCode(1),
				core.TestCheckError(&core.CliError{
					Err:  fmt.Errorf("a positional argument is required for this command"),
					Hint: "Try running: scw test positional <name-id> tag=world",
				}),
			),
		}))

		t.Run("Invalid1", core.Test(&core.TestConfig{
			Commands: testGetCommands(),
			Cmd:      "scw test positional name-id=plop tag=world",
			Check: core.TestCheckCombine(
				core.TestCheckExitCode(1),
				core.TestCheckError(&core.CliError{
					Err:  fmt.Errorf("a positional argument is required for this command"),
					Hint: "Try running: scw test positional plop tag=world",
				}),
			),
		}))

		t.Run("Invalid2", core.Test(&core.TestConfig{
			Commands: testGetCommands(),
			Cmd:      "scw test positional tag=world name-id=plop",
			Check: core.TestCheckCombine(
				core.TestCheckExitCode(1),
				core.TestCheckError(&core.CliError{
					Err:  fmt.Errorf("a positional argument is required for this command"),
					Hint: "Try running: scw test positional plop tag=world",
				}),
			),
		}))

		t.Run("Invalid3", core.Test(&core.TestConfig{
			Commands: testGetCommands(),
			Cmd:      "scw test positional plop name-id=plop",
			Check: core.TestCheckCombine(
				core.TestCheckExitCode(1),
				core.TestCheckError(&core.CliError{
					Err:  fmt.Errorf("a positional argument is required for this command"),
					Hint: "Try running: scw test positional plop",
				}),
			),
		}))
	})

	t.Run("simple", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test positional plop",
		Check:    core.TestCheckExitCode(0),
	}))

	t.Run("full command", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test positional plop tag=world",
		Check:    core.TestCheckExitCode(0),
	}))

	t.Run("full command", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test positional -h",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(0),
			core.TestCheckGolden(),
		),
	}))

	t.Run("multi positional", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test positional tag=tag01 test1 test2",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(0),
			core.TestCheckGolden(),
		),
	}))

	t.Run("multi positional json", core.Test(&core.TestConfig{
		Commands: testGetCommands(),
		Cmd:      "scw test positional -o json tag=tag01 test1 test2",
		Check: core.TestCheckCombine(
			core.TestCheckExitCode(0),
			core.TestCheckGolden(),
		),
	}))
}

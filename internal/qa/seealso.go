package qa

import (
	"fmt"
	"strings"

	"github.com/scaleway/scaleway-cli/v2/core"
)

type CommandInvalidSeeAlsoError struct {
	Command        *core.Command
	SeeAlsoCommand string
}

func (err CommandInvalidSeeAlsoError) Error() string {
	return fmt.Sprintf("command has invalid see_also commands, '%s' has '%s'",
		err.Command.GetCommandLine("scw"),
		err.SeeAlsoCommand,
	)
}

// testArgSpecInvalidError tests that all commands' see_also commands exist.
func testCommandInvalidSeeAlsoError(commands *core.Commands) []error {
	errors := []error(nil)

	for _, command := range commands.GetAll() {
		for _, seeAlso := range command.SeeAlsos {
			seeAlsoCommand := strings.Fields(seeAlso.Command)

			// Only check scw commands
			if len(seeAlsoCommand) <= 1 || seeAlsoCommand[0] != "scw" {
				continue
			}
			seeAlsoCommand = seeAlsoCommand[1:]

			if commands.Find(seeAlsoCommand...) == nil {
				errors = append(errors, &CommandInvalidSeeAlsoError{
					Command:        command,
					SeeAlsoCommand: seeAlso.Command,
				})
			}
		}
	}

	return errors
}

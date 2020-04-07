package core

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/scaleway/scaleway-cli/internal/human"
)

// CliError is an all-in-one error structure that can be used in commands to return useful errors to the user.
// CliError implements JSON and human marshaler for a smooth experience.
type CliError struct {

	// The original error that triggers this CLI error.
	// The Err.String() will be print in red to the user.
	Err error

	// Message allow to override the red message shown to the use.
	// By default we will use Err.String() but in same case you may want to keep Err
	// to avoid loosing detail in json output.
	Message string

	Details string
	Hint    string
}

func (s *CliError) Error() string {
	return s.Err.Error()
}

func (s *CliError) MarshalHuman() (string, error) {
	sections := []string(nil)
	if s.Err != nil {
		humanError := s.Err
		if s.Message != "" {
			humanError = fmt.Errorf(s.Message)
		}
		str, err := human.Marshal(humanError, nil)
		if err != nil {
			return "", err
		}
		sections = append(sections, str)
	}

	if s.Details != "" {
		str, err := human.Marshal(human.Capitalize(s.Details), &human.MarshalOpt{Title: "Details"})
		if err != nil {
			return "", err
		}
		sections = append(sections, str)
	}

	if s.Hint != "" {
		str, err := human.Marshal(human.Capitalize(s.Hint), &human.MarshalOpt{Title: "Hint"})
		if err != nil {
			return "", err
		}
		sections = append(sections, str)
	}

	return strings.Join(sections, "\n\n"), nil
}

func (s *CliError) MarshalJSON() ([]byte, error) {
	type tmpRes struct {
		Message string `json:"message"`
		Error   error  `json:"error"`
		Details string `json:"details"`
		Hint    string `json:"hint"`
	}
	return json.Marshal(&tmpRes{
		Message: s.Message,
		Error:   s.Err,
		Details: s.Details,
		Hint:    s.Hint,
	})
}

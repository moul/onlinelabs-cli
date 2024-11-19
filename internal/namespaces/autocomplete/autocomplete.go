package autocomplete

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/scaleway/scaleway-cli/v2/core"
	"github.com/scaleway/scaleway-cli/v2/internal/args"
	"github.com/scaleway/scaleway-cli/v2/internal/interactive"
	"github.com/scaleway/scaleway-sdk-go/logger"
)

func GetCommands() *core.Commands {
	cmds := core.NewCommands(
		autocompleteRootCommand(),
		autocompleteInstallCommand(),
		autocompleteCompleteBashCommand(),
		autocompleteCompleteFishCommand(),
		autocompleteCompleteZshCommand(),
		autocompleteScriptCommand(),
	)

	for _, cmd := range cmds.GetAll() {
		cmd.DisableAfterChecks = true
	}

	return cmds
}

func autocompleteRootCommand() *core.Command {
	return &core.Command{
		Short:     `Autocomplete related commands`,
		Long:      ``,
		Namespace: "autocomplete",
	}
}

type autocompleteScript struct {
	CompleteScript         string
	CompleteFunc           string
	ShellConfigurationFile map[string]string
}

// autocompleteScripts regroups the autocomplete scripts for the different shells
// The key is the path of the shell.
func autocompleteScripts(ctx context.Context, basename string) map[string]autocompleteScript {
	homePath := core.ExtractUserHomeDir(ctx)
	return map[string]autocompleteScript{
		"bash": {
			// If `scw` is the first word on the command line,
			// after hitting [tab] arguments are sent to `scw autocomplete complete bash`:
			//  - COMP_LINE: the complete command line
			//  - cword:     the index of the word being completed (source COMP_CWORD)
			//  - words:     the words composing the command line (source COMP_WORDS)
			//
			// Note that `=` signs are excluding from $COMP_WORDBREAKS. As a result, they are NOT be
			// considered as breaking words and arguments like `image=` will not be split.
			//
			// Then `scw autocomplete complete bash` process the line, and tries to returns suggestions.
			// These scw suggestions are put into `COMPREPLY` which is used by Bash to provides the shell suggestions.
			CompleteFunc: fmt.Sprintf(`
			_%[1]s() {
				_get_comp_words_by_ref -n = cword words

				output=$(%[1]s autocomplete complete bash -- "$COMP_LINE" "$cword" "${words[@]}")
				COMPREPLY=($output)
				# apply compopt option and ignore failure for older bash versions
				[[ $COMPREPLY == *= ]] && compopt -o nospace 2> /dev/null || true
				return
			}
			complete -F _%[1]s %[1]s
		`, basename),
			CompleteScript: fmt.Sprintf(`eval "$(%s autocomplete script shell=bash)"`, basename),
			ShellConfigurationFile: map[string]string{
				"darwin": path.Join(homePath, ".bash_profile"),
				"linux":  path.Join(homePath, ".bashrc"),
			},
		},
		"fish": {
			// (commandline)                             complete command line
			// (commandline --cursor)                    position of the cursor, as number of chars in the command line
			// (commandline --current-token)             word to complete
			// (commandline --tokenize --cut-at-cursor)  tokenized selection up until the current cursor position
			//                                           formatted as one string-type token per line
			//
			// If files are shown although --no-files is set,
			// it might be because you are using an alias for scw, such as :
			// 		alias scw='go run "$HOME"/scaleway-cli/cmd/scw/main.go'
			// You might want to run 'complete --erase --command go' during development.
			//
			// TODO: send rightWords
			CompleteFunc: fmt.Sprintf(`
			complete --erase --command %[1]s;
			complete --command %[1]s --no-files;
			complete --command %[1]s --arguments '(%[1]s autocomplete complete fish -- (commandline) (commandline --cursor) (commandline --current-token) (commandline --current-process --tokenize --cut-at-cursor))';
		`, basename),
			CompleteScript: fmt.Sprintf(`eval (%s autocomplete script shell=fish)`, basename),
			ShellConfigurationFile: map[string]string{
				"darwin": path.Join(homePath, ".config/fish/config.fish"),
				"linux":  path.Join(homePath, ".config/fish/config.fish"),
			},
		},
		"zsh": {
			// If you are using an alias for scw, such as :
			// 		alias scw='go run "$HOME"/scaleway-cli/cmd/scw/main.go'
			// you might want to run 'compdef _scw go' during development.
			CompleteFunc: fmt.Sprintf(`
			autoload -U compinit && compinit
			_%[1]s () {
				output=($(%[1]s autocomplete complete zsh -- ${CURRENT} ${words}))
				opts=('-S' ' ')
				if [[ $output == *= ]]; then
					opts=('-S' '')
				fi
				compadd "${opts[@]}" -- "${output[@]}"
			}
			compdef _%[1]s %[1]s
		`, basename),
			CompleteScript: fmt.Sprintf(`eval "$(%s autocomplete script shell=zsh)"`, basename),
			ShellConfigurationFile: map[string]string{
				"darwin": path.Join(homePath, ".zshrc"),
				"linux":  path.Join(homePath, ".zshrc"),
			},
		},
	}
}

type InstallArgs struct {
	Shell    string
	Basename string
}

func autocompleteInstallCommand() *core.Command {
	return &core.Command{
		Short:                `Install autocomplete script`,
		Long:                 `Install autocomplete script for a given shell and OS.`,
		Namespace:            "autocomplete",
		Resource:             "install",
		AllowAnonymousClient: true,
		ArgSpecs: core.ArgSpecs{
			{
				Name: "shell",
			},
			{
				Name: "basename",
				Default: func(ctx context.Context) (value string, doc string) {
					resp := core.ExtractBinaryName(ctx)
					return resp, resp
				},
			},
		},
		ArgsType: reflect.TypeOf(InstallArgs{}),
		Run:      InstallCommandRun,
	}
}

func InstallCommandRun(ctx context.Context, argsI interface{}) (i interface{}, e error) {
	// Warning
	_, _ = interactive.Println("To enable autocomplete, scw needs to update your shell configuration.")

	// If `shell=` is empty, ask for a value for `shell=`.
	shellArg := argsI.(*InstallArgs).Shell
	logger.Debugf("shellArg: %v", shellArg)
	if shellArg == "" {
		defaultShellName := "bash"

		if core.ExtractEnv(ctx, "SHELL") != "" {
			defaultShellName = filepath.Base(core.ExtractEnv(ctx, "SHELL"))
		}

		promptedShell, err := interactive.PromptStringWithConfig(&interactive.PromptStringConfig{
			Ctx:             ctx,
			Prompt:          "What type of shell are you using",
			DefaultValue:    defaultShellName,
			DefaultValueDoc: defaultShellName,
		})
		if err != nil {
			return nil, err
		}
		shellArg = promptedShell
	}

	shellName := filepath.Base(shellArg)
	basename := argsI.(*InstallArgs).Basename
	script, exists := autocompleteScripts(ctx, basename)[shellName]
	if !exists {
		return nil, unsupportedShellError(shellName)
	}

	// Find destination file depending on the OS.
	shellConfigurationFilePath, exists := script.ShellConfigurationFile[runtime.GOOS]
	if !exists {
		return nil, unsupportedOsError(runtime.GOOS)
	}

	// If the file doesn't exist, create it
	f, err := os.OpenFile(shellConfigurationFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return nil, installationNotFound(shellName, shellConfigurationFilePath, script.CompleteScript)
	}

	// Early exit if eval line is already present in the shell configuration.
	shellConfigurationFileContent, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if strings.Contains(string(shellConfigurationFileContent), script.CompleteScript) {
		_, _ = interactive.Println()
		_, _ = interactive.Println("Autocomplete looks already installed. If it does not work properly, try to open a new shell.")
		return "", nil
	}

	// Autocomplete script content
	autoCompleteScript := "\n# Scaleway CLI autocomplete initialization.\n" + script.CompleteScript

	// Warning
	_, _ = interactive.Println()
	_, _ = interactive.PrintlnWithoutIndent("To enable autocomplete we need to append to " + shellConfigurationFilePath + " the following lines:")
	_, _ = interactive.Println(strings.ReplaceAll(autoCompleteScript, "\n", "\n\t"))

	// Early exit if user disagrees
	_, _ = interactive.Println()
	continueInstallation, err := interactive.PromptBoolWithConfig(&interactive.PromptBoolConfig{
		Ctx:          ctx,
		Prompt:       "Do you want to proceed with these changes?",
		DefaultValue: true,
	})
	if err != nil {
		return nil, err
	}
	if !continueInstallation {
		return nil, installationCancelledError(shellName, script.CompleteScript)
	}

	// Append to file
	_, err = f.WriteString(autoCompleteScript + "\n")
	if err != nil {
		return nil, err
	}

	// Ack
	return &core.SuccessResult{
		Message: fmt.Sprintf("Autocomplete has been successfully installed for your %v shell.\nUpdated %v.", shellName, shellConfigurationFilePath),
	}, nil
}

func autocompleteCompleteBashCommand() *core.Command {
	return &core.Command{
		Short:     `Autocomplete for Bash`,
		Long:      `Autocomplete for Bash.`,
		Namespace: "autocomplete",
		Resource:  "complete",
		Verb:      "bash",
		// TODO: Switch AllowAnonymousClient to true when cache will be implemented.
		AllowAnonymousClient: false,
		Hidden:               true,
		DisableTelemetry:     true,
		ArgsType:             reflect.TypeOf(args.RawArgs{}),
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			rawArgs := *argsI.(*args.RawArgs)
			if len(rawArgs) < 3 {
				return nil, errors.New("not enough arguments")
			}
			wordIndex, err := strconv.Atoi(rawArgs[1])
			if err != nil {
				return nil, err
			}
			words := rawArgs[2:]
			if len(words) <= wordIndex {
				return nil, errors.New("index to complete is invalid")
			}

			aliases := core.ExtractAliases(ctx)

			leftWords := aliases.ResolveAliases(words[:wordIndex])
			wordToComplete := words[wordIndex]
			rightWords := aliases.ResolveAliases(words[wordIndex+1:])

			// If the wordToComplete is an argument label (cf. `arg=`), remove
			// this prefix for all suggestions.
			res := core.AutoComplete(ctx, leftWords, wordToComplete, rightWords)
			if strings.Contains(wordToComplete, "=") {
				prefix := strings.SplitAfterN(wordToComplete, "=", 2)[0]
				for k, p := range res.Suggestions {
					res.Suggestions[k] = strings.TrimPrefix(p, prefix)
				}
			}

			return strings.Join(res.Suggestions, " "), nil
		},
	}
}

func autocompleteCompleteFishCommand() *core.Command {
	return &core.Command{
		Short:     `Autocomplete for Fish`,
		Long:      `Autocomplete for Fish.`,
		Namespace: "autocomplete",
		Resource:  "complete",
		Verb:      "fish",
		// TODO: Switch AllowAnonymousClient to true when cache will be implemented.
		AllowAnonymousClient: false,
		Hidden:               true,
		DisableTelemetry:     true,
		ArgsType:             reflect.TypeOf(args.RawArgs{}),
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			rawArgs := *argsI.(*args.RawArgs)
			if len(rawArgs) < 4 {
				return nil, errors.New("not enough arguments")
			}

			aliases := core.ExtractAliases(ctx)

			leftWords := aliases.ResolveAliases(rawArgs[3:])
			wordToComplete := rawArgs[2]

			// TODO: compute rightWords once used by core.AutoComplete()
			// line := rawArgs[0]
			// charIndex, _ := strconv.Atoi(rawArgs[1])
			rightWords := []string(nil)

			res := core.AutoComplete(ctx, leftWords, wordToComplete, rightWords)

			// TODO: decide if we want to add descriptions
			// see https://stackoverflow.com/a/20879411
			// "followed optionally by a tab and a short description."
			return strings.Join(res.Suggestions, "\n"), nil
		},
	}
}

func autocompleteCompleteZshCommand() *core.Command {
	return &core.Command{
		Short:     `Autocomplete for Zsh`,
		Long:      `Autocomplete for Zsh.`,
		Namespace: "autocomplete",
		Resource:  "complete",
		Verb:      "zsh",
		// TODO: Switch AllowAnonymousClient to true when cache will be implemented.
		AllowAnonymousClient: false,
		Hidden:               true,
		DisableTelemetry:     true,
		ArgsType:             reflect.TypeOf(args.RawArgs{}),
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			rawArgs := *argsI.(*args.RawArgs)
			if len(rawArgs) < 2 {
				return nil, errors.New("not enough arguments")
			}

			// First arg is the word index.
			wordIndex, err := strconv.Atoi(rawArgs[0])
			if err != nil {
				return nil, err
			}
			wordIndex-- // In zsh word index starts at 1.

			if wordIndex <= 0 {
				return nil, errors.New("index cannot be 1 (0) or lower")
			}

			// Other args are all the words.
			words := rawArgs[1:]
			if len(words) <= wordIndex {
				words = append(words, "") // Handle case when last word is empty.
			}

			aliases := core.ExtractAliases(ctx)

			leftWords := aliases.ResolveAliases(words[:wordIndex])
			wordToComplete := words[wordIndex]
			rightWords := aliases.ResolveAliases(words[wordIndex+1:])

			res := core.AutoComplete(ctx, leftWords, wordToComplete, rightWords)
			return strings.Join(res.Suggestions, " "), nil
		},
	}
}

type autocompleteShowArgs struct {
	Shell    string
	Basename string
}

func autocompleteScriptCommand() *core.Command {
	return &core.Command{
		Short:                `Show autocomplete script for current shell`,
		Long:                 `Show autocomplete script for current shell.`,
		Namespace:            "autocomplete",
		Resource:             "script",
		AllowAnonymousClient: true,
		DisableTelemetry:     true,
		ArgSpecs: core.ArgSpecs{
			{
				Name:    "shell",
				Default: core.DefaultValueSetter(os.Getenv("SHELL")),
			},
			{
				Name: "basename",
				Default: func(ctx context.Context) (value string, doc string) {
					resp := core.ExtractBinaryName(ctx)
					return resp, resp
				},
			},
		},
		ArgsType: reflect.TypeOf(autocompleteShowArgs{}),
		Run: func(ctx context.Context, argsI interface{}) (i interface{}, e error) {
			shell := filepath.Base(argsI.(*autocompleteShowArgs).Shell)
			basename := argsI.(*autocompleteShowArgs).Basename
			script, exists := autocompleteScripts(ctx, basename)[shell]
			if !exists {
				return nil, unsupportedShellError(shell)
			}
			return TrimText(script.CompleteFunc), nil
		},
	}
}

func TrimText(str string) string {
	foundFirstNonEmptyLine := false
	strToRemove := ""
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if !foundFirstNonEmptyLine {
			if len(line) > 0 {
				for _, c := range line {
					if c == ' ' || c == '\t' {
						strToRemove += string(c)
						continue
					}
					break
				}
				foundFirstNonEmptyLine = true
			}
		}
		for _, c := range strToRemove {
			lines[i] = strings.Replace(lines[i], string(c), "", 1)
		}
	}
	lines = removeStartingAndEndingEmptyLines(lines)
	return strings.Join(lines, "\n")
}

func removeStartingAndEndingEmptyLines(lines []string) []string {
	lines = removeStartingEmptyLines(lines)
	lines = reverseLines(lines)
	lines = removeStartingEmptyLines(lines)
	lines = reverseLines(lines)
	return lines
}

func removeStartingEmptyLines(lines []string) []string {
	doAdd := false
	lines2 := []string(nil)
	for _, line := range lines {
		if len(line) > 0 {
			doAdd = true
		}
		if doAdd {
			lines2 = append(lines2, line)
		}
	}
	return lines2
}

func reverseLines(lines []string) []string {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
	return lines
}

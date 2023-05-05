package core

import (
	"context"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/stretchr/testify/assert"
)

func testAutocompleteGetCommands() *Commands {
	return NewCommands(
		&Command{
			Namespace: "test",
			Resource:  "flower",
			Verb:      "create",
			ArgsType:  reflect.TypeOf(struct{}{}),
			ArgSpecs: ArgSpecs{
				{
					Name: "name",
				},
				{
					Name:       "species",
					EnumValues: []string{"rose", "violet", "petunia", "virginia bluebell"},
				},
				{
					Name: "size",
					AutoCompleteFunc: func(ctx context.Context, prefix string) AutocompleteSuggestions {
						return []string{regexp.MustCompile("[a-z]").ReplaceAllString(prefix, "")}
					},
					EnumValues: []string{"S", "M", "L", "XL", "XXL"},
				},
				{
					Name:       "colours.{index}",
					EnumValues: []string{"blue", "red", "pink"},
				},
				{
					Name:       "leaves.{key}.size",
					EnumValues: []string{"S", "M", "L", "XL", "XXL"},
				},
			},
			WaitFunc: func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
				return nil, nil
			},
		},
		&Command{
			Namespace: "test",
			Resource:  "flower",
			Verb:      "delete",
			ArgsType: reflect.TypeOf(struct {
				WithLeaves bool
			}{}),
			ArgSpecs: ArgSpecs{
				{
					Name:       "name",
					EnumValues: []string{"hibiscus", "anemone"},
					Positional: true,
				},
				{
					Name: "with-leaves",
				},
			},
		},
	)
}

type autoCompleteTestCase struct {
	Suggestions         AutocompleteSuggestions
	WordToCompleteIndex int
	Words               []string
}

func runAutocompleteTest(ctx context.Context, tc *autoCompleteTestCase) func(*testing.T) {
	return func(t *testing.T) {
		words := tc.Words
		if len(words) == 0 {
			name := strings.Replace(t.Name(), "TestAutocomplete/", "", -1)
			name = strings.Replace(name, "_", " ", -1)
			// Test can contain a sharp if duplicated
			// MyTest/scw_-flag_#01
			sharpIndex := strings.Index(name, "#")
			if sharpIndex != -1 {
				name = name[:sharpIndex]
			}
			words = strings.Split(name, " ")
		}

		wordToCompleteIndex := len(words) - 1
		if tc.WordToCompleteIndex != 0 {
			wordToCompleteIndex = tc.WordToCompleteIndex
		}
		leftWords := words[:wordToCompleteIndex]
		wordToComplete := words[wordToCompleteIndex]
		rightWord := words[wordToCompleteIndex+1:]

		result := AutoComplete(ctx, leftWords, wordToComplete, rightWord)
		assert.Equal(t, tc.Suggestions, result.Suggestions)
	}
}

func TestAutocomplete(t *testing.T) {
	ctx := injectMeta(context.Background(), &meta{
		Commands: testAutocompleteGetCommands(),
	})

	type testCase = autoCompleteTestCase

	run := func(tc *testCase) func(*testing.T) {
		return runAutocompleteTest(ctx, tc)
	}

	t.Run("scw ", run(&testCase{Suggestions: AutocompleteSuggestions{"test"}}))
	t.Run("scw te", run(&testCase{Suggestions: AutocompleteSuggestions{"test"}}))
	t.Run("scw test", run(&testCase{Suggestions: AutocompleteSuggestions{"test"}}))
	t.Run("scw  flower create name=plop", run(&testCase{WordToCompleteIndex: 1, Suggestions: AutocompleteSuggestions{"test"}}))
	t.Run("scw te flower create name=plop", run(&testCase{WordToCompleteIndex: 1, Suggestions: AutocompleteSuggestions{"test"}}))
	t.Run("scw test ", run(&testCase{Suggestions: AutocompleteSuggestions{"flower"}}))
	t.Run("scw test fl", run(&testCase{Suggestions: AutocompleteSuggestions{"flower"}}))
	t.Run("scw test flower ", run(&testCase{Suggestions: AutocompleteSuggestions{"create", "delete"}}))
	t.Run("scw test flower cr", run(&testCase{Suggestions: AutocompleteSuggestions{"create"}}))
	t.Run("scw test flower d", run(&testCase{Suggestions: AutocompleteSuggestions{"delete"}}))
	t.Run("scw test flower create ", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=", "leaves.", "name=", "size=", "species="}}))
	t.Run("scw test flower create n", run(&testCase{Suggestions: AutocompleteSuggestions{"name="}}))
	t.Run("scw test flower create name", run(&testCase{Suggestions: AutocompleteSuggestions{"name="}}))
	t.Run("scw test flower create name=", run(&testCase{Suggestions: nil}))
	t.Run("scw test flower create name n", run(&testCase{Suggestions: AutocompleteSuggestions{"name="}}))
	t.Run("scw test flower create name=p", run(&testCase{Suggestions: nil}))
	t.Run("scw test flower create name=p ", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=", "leaves.", "size=", "species="}}))
	t.Run("scw test flower create name=plop n", run(&testCase{Suggestions: nil}))
	t.Run("scw test flower create n name=plop", run(&testCase{WordToCompleteIndex: 4, Suggestions: nil}))
	t.Run("scw test flower create s", run(&testCase{Suggestions: AutocompleteSuggestions{"size=", "species="}}))
	t.Run("scw test flower create species=", run(&testCase{Suggestions: AutocompleteSuggestions{"species=petunia", "species=rose", "species=violet", "species=virginia bluebell"}}))
	t.Run("scw test flower create species=v", run(&testCase{Suggestions: AutocompleteSuggestions{"species=violet", "species=virginia bluebell"}}))
	t.Run("scw test flower create size=a1b2c", run(&testCase{Suggestions: AutocompleteSuggestions{"size=12"}}))
	t.Run("scw test flower create colo", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0="}}))
	t.Run("scw test flower create colours.0", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0="}}))
	t.Run("scw test flower create colours.0=", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=blue", "colours.0=pink", "colours.0=red"}}))
	t.Run("scw test flower create colours.0=r", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=red"}}))
	t.Run("scw test flower create colo colours.1=red", run(&testCase{WordToCompleteIndex: 4, Suggestions: AutocompleteSuggestions{"colours.0="}}))
	t.Run("scw test flower create colo colours.0=blue colours.1=red", run(&testCase{WordToCompleteIndex: 4, Suggestions: AutocompleteSuggestions{"colours.2="}}))
	t.Run("scw test flower create colours.0=blue colours.1=r", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.1=red"}}))
	t.Run("scw test flower create leaves.", run(&testCase{Suggestions: AutocompleteSuggestions{"leaves."}}))
	t.Run("scw test flower create leaves.0", run(&testCase{Suggestions: AutocompleteSuggestions{"leaves.0.size="}}))
	t.Run("scw test flower create leaves.0.", run(&testCase{Suggestions: AutocompleteSuggestions{"leaves.0.size="}}))
	t.Run("scw test flower create leaves.0.size=M", run(&testCase{Suggestions: AutocompleteSuggestions{"leaves.0.size=M"}}))
	t.Run("scw test flower create leaves.0.size=M leaves", run(&testCase{Suggestions: AutocompleteSuggestions{"leaves."}}))
	t.Run("scw test flower create leaves.0.size=M leaves leaves.1.size=M", run(&testCase{WordToCompleteIndex: 5, Suggestions: AutocompleteSuggestions{"leaves."}}))
	t.Run("scw test flower delete ", run(&testCase{Suggestions: AutocompleteSuggestions{"anemone", "hibiscus", "with-leaves="}}))
	t.Run("scw test flower delete w", run(&testCase{Suggestions: AutocompleteSuggestions{"with-leaves="}}))
	t.Run("scw test flower delete h", run(&testCase{Suggestions: AutocompleteSuggestions{"hibiscus"}}))
	t.Run("scw test flower delete with-leaves=true ", run(&testCase{Suggestions: AutocompleteSuggestions{"anemone", "hibiscus"}})) // invalid notation
	t.Run("scw test flower delete hibiscus n", run(&testCase{Suggestions: nil}))
	t.Run("scw test flower delete hibiscus w", run(&testCase{Suggestions: AutocompleteSuggestions{"with-leaves="}}))
	t.Run("scw test flower delete hibiscus with-leaves=true", run(&testCase{Suggestions: AutocompleteSuggestions{"with-leaves=true"}}))
	t.Run("scw test flower delete hibiscus with-leaves=true ", run(&testCase{Suggestions: AutocompleteSuggestions{"anemone"}}))
	t.Run("scw test flower delete hibiscus with-leaves=", run(&testCase{Suggestions: AutocompleteSuggestions{"with-leaves=false", "with-leaves=true"}}))
	t.Run("scw test flower delete hibiscus with-leaves=tr", run(&testCase{Suggestions: AutocompleteSuggestions{"with-leaves=true"}}))
	t.Run("scw test flower delete hibiscus with-leaves=yes", run(&testCase{Suggestions: nil}))
	t.Run("scw test flower create leaves.0.size=", run(&testCase{Suggestions: AutocompleteSuggestions{"leaves.0.size=L", "leaves.0.size=M", "leaves.0.size=S", "leaves.0.size=XL", "leaves.0.size=XXL"}}))
	t.Run("scw -", run(&testCase{Suggestions: AutocompleteSuggestions{"--config", "--debug", "--help", "--output", "--profile", "-D", "-c", "-h", "-o", "-p"}}))
	t.Run("scw test -o j", run(&testCase{Suggestions: AutocompleteSuggestions{"json"}}))
	t.Run("scw test flower -o ", run(&testCase{Suggestions: AutocompleteSuggestions{PrinterTypeHuman.String(), PrinterTypeJSON.String(), PrinterTypeTemplate.String(), PrinterTypeYAML.String()}}))
	t.Run("scw test flower -o json create -", run(&testCase{Suggestions: AutocompleteSuggestions{"--config", "--debug", "--help", "--output", "--profile", "--wait", "-D", "-c", "-h", "-p", "-w"}}))
	t.Run("scw test flower create name=p -o j", run(&testCase{Suggestions: AutocompleteSuggestions{"json"}}))
	t.Run("scw test flower create name=p -o json ", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=", "leaves.", "size=", "species="}}))
	t.Run("scw test flower create name=p -o=json ", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=", "leaves.", "size=", "species="}}))
	t.Run("scw test flower create name=p -o=jso", run(&testCase{Suggestions: nil})) // TODO: make this work
	t.Run("scw test flower create name=p -o", run(&testCase{Suggestions: AutocompleteSuggestions{"-o"}}))
	t.Run("scw test -o json flower create ", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=", "leaves.", "name=", "size=", "species="}}))
	t.Run("scw test flower create name=p --profile xxxx ", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=", "leaves.", "size=", "species="}}))
	t.Run("scw test --profile xxxx flower create name=p ", run(&testCase{Suggestions: AutocompleteSuggestions{"colours.0=", "leaves.", "size=", "species="}}))
	t.Run("scw test flower create name=p --profile xxxx", run(&testCase{Suggestions: nil}))

	t.Run("scw test flower -o json delete -", run(&testCase{Suggestions: AutocompleteSuggestions{"--config", "--debug", "--help", "--output", "--profile", "-D", "-c", "-h", "-p"}}))
	t.Run("scw test flower delete -o ", run(&testCase{Suggestions: AutocompleteSuggestions{PrinterTypeHuman.String(), PrinterTypeJSON.String(), PrinterTypeTemplate.String(), PrinterTypeYAML.String()}}))
	t.Run("scw test flower delete -o j", run(&testCase{Suggestions: AutocompleteSuggestions{"json"}}))
	t.Run("scw test flower delete -o json ", run(&testCase{Suggestions: AutocompleteSuggestions{"anemone", "hibiscus", "with-leaves="}}))
	t.Run("scw test flower delete -o=json ", run(&testCase{Suggestions: AutocompleteSuggestions{"anemone", "hibiscus", "with-leaves="}}))
	t.Run("scw test flower delete -o json hibiscus w", run(&testCase{Suggestions: AutocompleteSuggestions{"with-leaves="}}))
	t.Run("scw test flower delete -o=json hibiscus w", run(&testCase{Suggestions: AutocompleteSuggestions{"with-leaves="}}))
}

func TestAutocompleteArgs(t *testing.T) {
	commands := testAutocompleteGetCommands()
	commands.Add(&Command{
		Namespace: "test",
		Resource:  "flower",
		Verb:      "get",
		ArgsType: reflect.TypeOf(struct {
			Name         string
			MaterialName string
		}{}),
		ArgSpecs: ArgSpecs{
			{
				Name:       "name",
				Positional: true,
			},
			{
				Name: "material-name",
			},
		},
	})
	commands.Add(&Command{
		Namespace: "test",
		Resource:  "flower",
		Verb:      "list",
		ArgsType: reflect.TypeOf(struct {
		}{}),
		ArgSpecs: ArgSpecs{},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			return []*struct {
				Name string
			}{
				{
					Name: "flower1",
				},
				{
					Name: "flower2",
				},
			}, nil
		},
	})
	commands.Add(&Command{
		Namespace: "test",
		Resource:  "material",
		Verb:      "list",
		ArgsType: reflect.TypeOf(struct {
		}{}),
		ArgSpecs: ArgSpecs{},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			return []*struct {
				Name string
			}{
				{
					Name: "material1",
				},
				{
					Name: "material2",
				},
			}, nil
		},
	})
	ctx := injectMeta(context.Background(), &meta{
		Commands: commands,
		betaMode: true,
	})

	type testCase = autoCompleteTestCase

	run := func(tc *testCase) func(*testing.T) {
		return runAutocompleteTest(ctx, tc)
	}

	t.Run("scw test flower get ", run(&testCase{Suggestions: AutocompleteSuggestions{"flower1", "flower2", "material-name="}}))
	t.Run("scw test flower get material-name=", run(&testCase{Suggestions: AutocompleteSuggestions{"material-name=material1", "material-name=material2"}}))
	t.Run("scw test flower get material-name=mat ", run(&testCase{Suggestions: AutocompleteSuggestions{"flower1", "flower2"}}))
	t.Run("scw test flower create name=", run(&testCase{Suggestions: AutocompleteSuggestions(nil)}))
}

func TestAutocompleteProfiles(t *testing.T) {
	commands := testAutocompleteGetCommands()
	ctx := injectMeta(context.Background(), &meta{
		Commands: commands,
		betaMode: true,
	})

	type testCase = autoCompleteTestCase

	run := func(tc *testCase) func(*testing.T) {
		return runAutocompleteTest(ctx, tc)
	}
	t.Run("scw -p ", run(&testCase{Suggestions: nil}))
	t.Run("scw test -p ", run(&testCase{Suggestions: nil}))
	t.Run("scw test flower --profile ", run(&testCase{Suggestions: nil}))

	injectConfig(ctx, &scw.Config{
		Profiles: map[string]*scw.Profile{
			"p1": nil,
			"p2": nil,
		},
	})

	t.Run("scw -p ", run(&testCase{Suggestions: AutocompleteSuggestions{"p1", "p2"}}))
	t.Run("scw test -p ", run(&testCase{Suggestions: AutocompleteSuggestions{"p1", "p2"}}))
	t.Run("scw test flower --profile ", run(&testCase{Suggestions: AutocompleteSuggestions{"p1", "p2"}}))
}

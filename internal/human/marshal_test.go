package human

import (
	"strings"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/dustin/go-humanize"
)

type Struct struct {
	String      string
	Int         int
	Bool        bool
	Strings     []string
	Time        time.Time
	Struct      *Struct
	Nil         *Struct
	Structs     []*Struct
	Map         map[string]string
	Stringer    Stringer
	StringerPtr *Stringer
}

type Address struct {
	Street string
	City   string
}

type Acquaintance struct {
	Name string
	Link string
}

type Human struct {
	Name          string
	Age           int
	Address       *Address
	Acquaintances []*Acquaintance
}

type NestedAnonymous struct {
	Name string
}

type Anonymous struct {
	NestedAnonymous
	Name string
}

type Stringer struct{}

func (s Stringer) String() string {
	return "a stringer"
}

func TestMarshal(t *testing.T) {
	type testCase struct {
		data   interface{}
		opt    *MarshalOpt
		result string
		err    error
	}

	run := func(tc *testCase) func(*testing.T) {
		return func(t *testing.T) {
			result, err := Marshal(tc.data, tc.opt)

			// Format expected to allow indentation when writing test
			expected := tc.result
			expected = strings.Replace(expected, "\t", "", -1)
			expected = strings.Trim(expected, "\n")

			if tc.result != "" {
				assert.NoError(t, err)
				assert.Equal(t, expected, result)
			} else {
				assert.Equal(t, err, err)
			}
		}
	}

	date := time.Date(1990, 11, 17, 20, 20, 0, 0, time.UTC)
	humanDate := humanize.Time(date)
	t.Run("struct", run(&testCase{
		data: &Struct{
			String:  "This is a string",
			Int:     42,
			Bool:    true,
			Strings: []string{"s1", "s2"},
			Time:    date,
			Struct:  &Struct{},
			Nil:     nil,
			Structs: []*Struct{
				{
					String: "Nested string",
				},
			},
			Map: map[string]string{
				"key1": "v1",
				"key2": "v2",
			},
			Stringer:    Stringer{},
			StringerPtr: &Stringer{},
		},
		result: `
			String              This is a string
			Int                 42
			Bool                true
			Strings.0           s1
			Strings.1           s2
			Time                ` + humanDate + `
			Struct.String       -
			Struct.Int          0
			Struct.Bool         false
			Struct.Time         a long while ago
			Struct.Stringer     a stringer
			Structs.0.String    Nested string
			Structs.0.Int       0
			Structs.0.Bool      false
			Structs.0.Time      a long while ago
			Structs.0.Stringer  a stringer
			Map.key1            v1
			Map.key2            v2
			Stringer            a stringer
			StringerPtr         a stringer
		`,
	}))

	t.Run("struct2", run(&testCase{
		data: &Human{
			Name:    "Sherlock Holmes",
			Age:     42,
			Address: &Address{Street: "221b Baker St", City: "London"},
			Acquaintances: []*Acquaintance{
				{Name: "Dr watson", Link: "Assistant"},
				{Name: "Mrs. Hudson", Link: "Landlady"},
			},
		},
		opt: &MarshalOpt{
			Title: "Personal Information",
			Sections: []*MarshalSection{
				{FieldName: "Address"},
				{Title: "Relationship", FieldName: "Acquaintances"},
			},
		},
		result: `
			Personal Information:
			Name  Sherlock Holmes
			Age   42
			
			Address:
			Street  221b Baker St
			City    London
			
			Relationship:
			NAME         LINK
			Dr watson    Assistant
			Mrs. Hudson  Landlady
		`,
	}))

	t.Run("empty string", run(&testCase{
		data:   "",
		result: `-`,
	}))

	t.Run("nil", run(&testCase{
		data:   nil,
		result: `-`,
	}))

	t.Run("anonymous", run(&testCase{
		data: &Anonymous{
			NestedAnonymous: NestedAnonymous{
				Name: "John",
			},
			Name: "Paul",
		},
		result: `Name  Paul`,
	}))
}

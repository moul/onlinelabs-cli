package core

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/scaleway/scaleway-cli/internal/human"
)

// Type defines an formatter format.
type PrinterType string

func (p PrinterType) String() string {
	return string(p)
}

var (
	// PrinterTypeJSON defines a JSON formatter.
	PrinterTypeJSON = PrinterType("json")

	// PrinterTypeHuman defines a human readable formatted formatter.
	PrinterTypeHuman = PrinterType("human")
)

type PrinterConfig struct {
	OutputFlag string
	Stdout     io.Writer
	Stderr     io.Writer
}

// NewPrinter returns an initialized formatter corresponding to a given FormatterType.
func NewPrinter(config *PrinterConfig) (*Printer, error) {
	printer := &Printer{
		stdout: config.Stdout,
		stderr: config.Stderr,
	}

	// First we parse OutputFlag to extract printerName and printerOpt (e.g json=pretty)
	tmp := strings.SplitN(config.OutputFlag, "=", 2)
	printerName := tmp[0]
	printerOpt := ""
	if len(tmp) > 1 {
		printerOpt = tmp[1]
	}

	// We call the correct setup method depending on the printer type
	switch printerName {
	case PrinterTypeHuman.String():
		err := setupHumanPrinter(printer, printerOpt)
		if err != nil {
			return nil, err
		}

	case PrinterTypeJSON.String():
		err := setupJSONPrinter(printer, printerOpt)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("invalid output format: %s", printerName)
	}

	return printer, nil
}

func setupJSONPrinter(printer *Printer, opts string) error {
	printer.printerType = PrinterTypeJSON
	switch opts {
	case "pretty":
		printer.jsonPretty = true
	case "":
	default:
		return fmt.Errorf("invalid option %s for json outout. Valid options are: pretty", opts)
	}
	return nil
}

func setupHumanPrinter(printer *Printer, _ string) error {
	printer.printerType = PrinterTypeHuman
	return nil
}

type Printer struct {
	printerType PrinterType
	stdout      io.Writer
	stderr      io.Writer
	jsonPretty  bool
}

func (p *Printer) Print(data interface{}, opt *human.MarshalOpt) error {
	// No matter the printer type if data is a RawResult we should print it as is.
	if rawResult, isRawResult := data.(RawResult); isRawResult {
		_, err := p.stdout.Write(rawResult)
		return err
	}

	switch p.printerType {
	case PrinterTypeHuman:
		return p.printHuman(data, opt)
	case PrinterTypeJSON:
		return p.printJSON(data)
	default:
		return fmt.Errorf("unkonwn format: %s", p.printerType)
	}
}

func (p *Printer) printHuman(data interface{}, opt *human.MarshalOpt) error {
	str, err := human.Marshal(data, opt)
	if err != nil {
		return err
	}

	// If human marshal return an empty string we avoid printing empty line
	if str == "" {
		return nil
	}

	if _, isError := data.(error); isError {
		_, err = fmt.Fprintln(p.stderr, str)
	} else {
		_, err = fmt.Fprintln(p.stdout, str)
	}
	return err
}

func (p *Printer) printJSON(data interface{}) error {
	_, implementMarshaler := data.(json.Marshaler)
	err, isError := data.(error)

	if isError && !implementMarshaler {
		data = map[string]string{
			"error": err.Error(),
		}
	}

	writer := p.stdout
	if isError {
		writer = p.stderr
	}
	encoder := json.NewEncoder(writer)
	if p.jsonPretty {
		encoder.SetIndent("", "  ")
	}

	// We handle special case to make sure that a nil slice is marshal as `[]`
	if reflect.TypeOf(data).Kind() == reflect.Slice && reflect.ValueOf(data).IsNil() {
		_, err := p.stdout.Write([]byte("[]\n"))
		return err
	}

	return encoder.Encode(data)
}

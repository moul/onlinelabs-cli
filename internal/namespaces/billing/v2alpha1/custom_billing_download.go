package billing

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-cli/v2/internal/interactive"
	"github.com/scaleway/scaleway-cli/v2/internal/terminal"
	billing "github.com/scaleway/scaleway-sdk-go/api/billing/v2alpha1"
)

const (
	invoiceDefaultPrefix = "scaleway-invoice"
)

type billingDownloadRequest struct {
	billing.DownloadInvoiceRequest
	// extra arguments
	FilePath     string
	ForceReplace bool
}

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func buildDownloadCommand(command *core.Command) *core.Command {
	command.ArgsType = reflect.TypeOf(billingDownloadRequest{})
	command.ArgSpecs = core.ArgSpecs{
		{
			Name:       "invoice-id",
			Short:      `Invoice ID`,
			Required:   true,
			Deprecated: false,
			Positional: false,
		},
		{
			Name:       "file-path",
			Short:      `Wanted file path`,
			Required:   false,
			Deprecated: false,
			Positional: false,
		},
		{
			Name:       "file-type",
			Short:      `Wanted file extension`,
			Required:   false,
			Deprecated: false,
			Positional: false,
			Default:    core.DefaultValueSetter(billing.DownloadInvoiceRequestFileTypePdf.String()),
		},
		{
			Name:       "force-replace",
			Short:      `Force file replacement`,
			Required:   false,
			Deprecated: false,
			Positional: false,
			Default:    core.DefaultValueSetter("false"),
		},
	}
	command.Run = billingDownloadRun
	command.PreValidateFunc = func(ctx context.Context, argsI interface{}) error {
		args := argsI.(*billingDownloadRequest)
		askPrompt := false
		request := &billing.DownloadInvoiceRequest{
			InvoiceID: args.InvoiceID,
			FileType:  args.FileType,
		}
		client := core.ExtractClient(ctx)
		billingAPI := billing.NewAPI(client)
		resp, err := billingAPI.DownloadInvoice(request)
		if err != nil {
			return err
		}

		date, err := trimDateFromFileName(resp.Name)
		if err != nil {
			return fmt.Errorf("parse date on file name")
		}

		dir, file := filepath.Split(args.FilePath)
		if len(file) > 0 {
			fileExtension := filepath.Ext(file)
			if extensionOnFile := checkInvoiceExt(fileExtension); !extensionOnFile {
				return fmt.Errorf("file has not supported extension")
			}
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			return err
		}

		if args.ForceReplace {
			return nil
		}

		// check default name
		defaultFileName := fmt.Sprintf("%s-%s-%s.%s", invoiceDefaultPrefix, date, args.InvoiceID, args.FileType)
		// read only in the parent path
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			// case default name on directory
			if len(file) == 0 && e.Name() == defaultFileName {
				file = defaultFileName
				askPrompt = true
			}

			if file == e.Name() {
				askPrompt = true
			}
		}

		if askPrompt {
			_, _ = interactive.PrintlnWithoutIndent(`
					Current file exist is located at ` + terminal.Style(args.FilePath, color.Faint))
			overrideFile, err := interactive.PromptBoolWithConfig(&interactive.PromptBoolConfig{
				Prompt:       fmt.Sprintf("Do you want to override the current file: %s ?", file),
				DefaultValue: false,
				Ctx:          ctx,
			})
			if err != nil {
				return err
			}
			if !overrideFile {
				return fmt.Errorf("download file canceled")
			}
		}

		return nil
	}

	return command
}

func addExt(fileName, contentType string) string {
	switch contentType {
	case "application/pdf":
		fileName = fmt.Sprintf("%s.pdf", fileName)
	}

	return fileName
}

func checkInvoiceExt(ext string) bool {
	switch ext {
	case ".pdf":
		return true
	}

	return false
}

func billingDownloadRun(ctx context.Context, argsI interface{}) (interface{}, error) {
	argsDownload := argsI.(*billingDownloadRequest)

	request := &billing.DownloadInvoiceRequest{
		InvoiceID: argsDownload.InvoiceID,
		FileType:  argsDownload.FileType,
	}

	client := core.ExtractClient(ctx)
	billingAPI := billing.NewAPI(client)
	resp, err := billingAPI.DownloadInvoice(request)
	if err != nil {
		return nil, err
	}

	fileName := fileNameWithoutExtTrimSuffix(argsDownload.FilePath)
	// check file path argument
	fInfo, err := os.Stat(argsDownload.FilePath)
	if err == nil {
		if fInfo.IsDir() {
			// case when filepath is a directory: join default name with custom path
			date, err := trimDateFromFileName(resp.Name)
			if err != nil {
				return nil, fmt.Errorf("parse date on file name")
			}

			defaultFileName := fmt.Sprintf("%s-%s-%s", invoiceDefaultPrefix, date, argsDownload.InvoiceID)
			pathAbs, err := filepath.Abs(argsDownload.FilePath)
			if err != nil {
				return nil, err
			}
			fileName = filepath.Join(pathAbs, defaultFileName)
		}
	}
	// add supported extension
	fileName = addExt(fileName, resp.ContentType)

	fileOutput, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer fileOutput.Close()

	_, err = io.Copy(fileOutput, resp.Content)
	if err != nil {
		return nil, err
	}

	return &core.SuccessResult{Empty: true}, nil
}

func trimDateFromFileName(filename string) (string, error) {
	formatLayout := "2006-01"
	m := strings.Split(filename, "-")
	d, err := time.Parse(formatLayout, fmt.Sprintf("%s-%s", m[1], m[2]))
	if err != nil {
		return "", err
	}
	return d.Format(formatLayout), nil
}

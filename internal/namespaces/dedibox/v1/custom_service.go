package dedibox

import (
	"context"
	"errors"
	"fmt"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/dedibox/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"net/http"
	"time"
)

const (
	serviceActionTimeout = time.Minute * 60
)

const (
	serviceActionCreate = iota
	serviceActionDelete
)

func serviceCreateBuilder(c *core.Command) *core.Command {
	c.WaitFunc = waitForServiceFunc(serviceActionCreate)
	return c
}

func waitForServiceFunc(action int) core.WaitFunc {
	return func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
		service, err := dedibox.NewAPI(core.ExtractClient(ctx)).WaitForService(&dedibox.WaitForServiceRequest{
			ServiceID:     respI.(*dedibox.Service).ID,
			Zone:          argsI.(*dedibox.CreateServerRequest).Zone,
			Timeout:       scw.TimeDurationPtr(serviceActionTimeout),
			RetryInterval: core.DefaultRetryInterval,
		})
		switch action {
		case serviceActionCreate:
			return service, err
		case serviceActionDelete:
			if err != nil {
				notFoundError := &scw.ResourceNotFoundError{}
				responseError := &scw.ResponseError{}
				if errors.As(err, &responseError) && responseError.StatusCode == http.StatusNotFound || errors.As(err, &notFoundError) {
					return fmt.Sprintf("Server %s successfully deleted.", respI.(*dedibox.Service).ID), nil
				}
			}
		}
		return nil, err
	}
}

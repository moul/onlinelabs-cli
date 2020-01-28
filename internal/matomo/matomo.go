package matomo

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/scaleway/scaleway-cli/internal/terminal"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	// full reference API https://developer.matomo.org/api-reference/tracking-api
	endpoint     = "kindermoumoute.matomo.cloud/matomo.php"
	project      = "4"
	enableRecord = "1"
	apiVersion   = "1"
)

var ForceTracking = os.Getenv("SCW_FORCE_TRACKING") == "true"

type TrackCommandRequest struct {
	RunCommand    string
	Version       string
	ExecutionTime time.Duration
}

func TrackCommand(request *TrackCommandRequest) error {
	// disable tracking when no command was run, or on completion command
	if request.RunCommand == "" || strings.HasPrefix(request.RunCommand, "autocomplete complete") {
		return nil
	}

	// compute or retrieve tracking parameters
	terminalResolution := fmt.Sprintf("%dx%d", terminal.GetWidth(), terminal.GetHeight())
	commandDurationInMs := fmt.Sprintf("%d", request.ExecutionTime/time.Millisecond)
	randNumber := generateRandNumber()
	action, actionURL := commandToAction(request.RunCommand)
	userAgent := fakeUserAgent(request.Version)
	organizationID := ""
	config, err := scw.LoadConfig()
	if err == nil {
		profile, err := config.GetActiveProfile()
		if err == nil && profile.DefaultOrganizationID != nil {
			organizationID = *profile.DefaultOrganizationID
		}
	}

	// build the query parameters in the URL
	query := url.Values{}

	// required
	query.Add("idsite", project)
	query.Add("rec", enableRecord)

	// recommended
	query.Add("action_name", action)
	query.Add("url", actionURL)
	query.Add("rand", randNumber)
	query.Add("apiv", apiVersion)

	// optional
	query.Add("res", terminalResolution)
	query.Add("ua", userAgent)
	query.Add("uid", organizationID)
	query.Add("gt_ms", commandDurationInMs)

	matomoURL := url.URL{
		Path:     endpoint,
		Scheme:   "https",
		RawQuery: query.Encode(),
	}

	// send the report
	resp, err := http.Get(matomoURL.String())
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("non-success status code %d: %s", resp.StatusCode, matomoURL.String())
	}

	logger.Debugf("Telemetry successfully sent %s", matomoURL.String())

	return nil
}

func fakeUserAgent(version string) string {
	return fmt.Sprintf("scaleway-cli/%s (%s; %s; %s)", version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func commandToAction(command string) (action string, url string) {
	command = "scw " + strings.Replace(command, ".", " ", -1)
	return command, "https://" + strings.Replace(command, " ", "/", -1)
}

func generateRandNumber() string {
	bigRand, err := rand.Int(rand.Reader, big.NewInt(int64(1<<uint64(32))-1))
	if err != nil {
		return ""
	}
	return bigRand.String()
}

// IsTelemetryEnabled returns true when the Opt-In send_usage attribute in the config is set.
func IsTelemetryEnabled() bool {
	config, err := scw.LoadConfig()
	if err != nil {
		return false
	}
	return config.SendUsage
}

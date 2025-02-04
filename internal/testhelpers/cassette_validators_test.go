package testhelpers_test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/stretchr/testify/require"
)

func exceptionsCassettesCases() map[string]struct{} {
	return map[string]struct{}{
		"../namespaces/baremetal/v1/testdata/test-reboot-server-errors-error-cannot-be-rebooted-while-not-delivered.cassette.yaml":         {},
		"../namespaces/baremetal/v1/testdata/test-start-server-errors-error-cannot-be-started-while-not-delivered.cassette.yaml":           {},
		"../namespaces/baremetal/v1/testdata/test-stop-server-errors-error-cannot-be-stopped-while-not-delivered.cassette.yaml":            {},
		"../namespaces/init/testdata/test-init-cl-iv2-config-no-prompt-overwrite-for-new-profile.cassette.yaml":                            {},
		"../namespaces/init/testdata/test-init-cl-iv2-config-prompt-overwrite-for-existing-profile.cassette.yaml":                          {},
		"../namespaces/init/testdata/test-init-ssh-key-unregistered.cassette.yaml":                                                         {},
		"../namespaces/init/testdata/test-init-ssh-with-local-ed25519-key.cassette.yaml":                                                   {},
		"../namespaces/instance/v1/testdata/test-server-update-no-initial-placement-group&-placement-group-id=invalid-pg-id.cassette.yaml": {},
		"../namespaces/mnq/v1beta1/testdata/test-create-context-with-wrond-id-simple.cassette.yaml":                                        {},
		"../namespaces/mnq/v1beta1/testdata/test-create-context-with-wrong-id-wrong-account-id.cassette.yaml":                              {},
		"../namespaces/redis/v1/testdata/test-endpoints-edge-cases-private-endpoint-with-both-attributes-set.cassette.yaml":                {},
		"../namespaces/redis/v1/testdata/test-endpoints-edge-cases-private-endpoint-with-none-set.cassette.yaml":                           {},
		"../namespaces/registry/v1/testdata/test-registry-install-docker-helper-command-simple.cassette.yaml":                              {},
		"../namespaces/registry/v1/testdata/test-registry-install-docker-helper-command-with-profile.cassette.yaml":                        {},
	}
}

func fileNameWithoutExtSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// getTestFiles returns a map of cassettes files
func getTestFiles() (map[string]struct{}, error) {
	filesMap := make(map[string]struct{})
	exceptions := exceptionsCassettesCases()
	err := filepath.WalkDir("../namespaces", func(path string, _ fs.DirEntry, _ error) error {
		isCassette := strings.Contains(path, "cassette")
		_, isException := exceptions[path]
		if isCassette && !isException {
			filesMap[fileNameWithoutExtSuffix(path)] = struct{}{}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return filesMap, nil
}

func checkErrCodeExcept(i *cassette.Interaction, c *cassette.Cassette, codes ...int) bool {
	exceptions := exceptionsCassettesCases()
	_, isException := exceptions[c.File]
	if isException {
		return isException
	}
	if i.Response.Code >= 400 {
		for _, httpCode := range codes {
			if i.Response.Code == httpCode {
				return true
			}
		}
		return false
	}
	return true
}

// isTransientStateError checks if the interaction response is a transient state error
// Transient state error are expected when creating resource linked to each other
// example:
// creating a gateway_network will set its public gateway to a transient state
// when creating 2 gateway_network, one will fail with a transient state error
// but the transient state error will be caught, it will wait again for the resource to be ready
func isTransientStateError(i *cassette.Interaction) bool {
	if i.Response.Code != 409 {
		return false
	}

	scwError := struct {
		Type string `json:"type"`
	}{}

	err := json.Unmarshal([]byte(i.Response.Body), &scwError)
	if err != nil {
		return false
	}

	return scwError.Type == "transient_state"
}

func checkErrorCode(c *cassette.Cassette) error {
	for _, i := range c.Interactions {
		if !checkErrCodeExcept(i, c, http.StatusNotFound, http.StatusTooManyRequests, http.StatusForbidden, http.StatusGone) &&
			!isTransientStateError(i) {
			return fmt.Errorf("status: %v found on %s. method: %s, url %s\nrequest body = %v\nresponse body = %v", i.Response.Code, c.Name, i.Request.Method, i.Request.URL, i.Request.Body, i.Response.Body)
		}
	}

	return nil
}

func TestAccCassettes_Validator(t *testing.T) {
	paths, err := getTestFiles()
	require.NoError(t, err)

	for path := range paths {
		c, err := cassette.Load(path)
		require.NoError(t, err)
		require.NoError(t, checkErrorCode(c))
	}
}

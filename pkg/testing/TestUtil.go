package commonTesting

import (
	"os"
	"strconv"
	"testing"
)

const DisableIntegrationTestEnvKey = "INTEGRATION_TEST_DISABLE"

func SkipIntegrationTestIfNeed(T *testing.T) {
	if IsSkipIntegrationTest() {
		T.Skip("Integration test is disabled")
	}
}

func IsSkipIntegrationTest() bool {
	integrationDisabledStr := os.Getenv(DisableIntegrationTestEnvKey)
	if integrationDisabledStr == "" {
		return false
	}

	integrationDisabled, err := strconv.ParseBool(integrationDisabledStr)
	if err != nil {
		panic(err)
	}

	return integrationDisabled
}

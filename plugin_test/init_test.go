package plugin_test

import (
	"os"
	"testing"
)

const (
	defTimeoutSecond     = 10
	defTimeoutFileSecond = 30
)

var (
	envDebug    = false
	envUserName = ""
	envPassword = ""

	envPluginWebhook = ""
)

func init() {
	envDebug = os.Getenv("ENV_DEBUG") == "true"

	envUserName = os.Getenv("ENV_PLUGIN_USERNAME")
	envPassword = os.Getenv("ENV_PLUGIN_PASSWORD")

	envPluginWebhook = os.Getenv("PLUGIN_WEBHOOK")
}

func envCheck(t *testing.T) bool {
	mustSetEnvList := []string{
		"PLUGIN_WEBHOOK",
	}
	for _, item := range mustSetEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %v, than run test", mustSetEnvList)
			return true
		}
	}

	return false
}

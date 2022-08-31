package config_test

import (
	"testing"

	cfg "github.com/mrzacarias/bootstrapper/config"
)

func TestMain(t *testing.T) {
	config := cfg.GetConfig()
	checkConfig(t, "AppName", config.AppName, "NOT SET")
	checkConfig(t, "AppPrefix", config.AppPrefix, "NOT SET")
	checkConfig(t, "AppPath", config.AppPath, "NOT SET")
}

// DRY helpers for checking a config attribute
func checkConfig(t *testing.T, attr string, got string, want string) {
	if want != got {
		t.Fatalf("Attribute '%s' should be %v, but it was %v", attr, want, got)
	}
}

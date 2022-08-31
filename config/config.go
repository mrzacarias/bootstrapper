package config

import (
	"flag"

	"github.com/spf13/viper"
)

// AppConfig will hold our environment variables
type AppConfig struct {
	AppName   string
	AppPrefix string
	AppPath   string
}

func init() {
	viper.SetEnvPrefix("BTSPR")
	viper.AutomaticEnv()
	viper.SetDefault("app_name", "NOT SET")
	viper.SetDefault("app_prefix", "NOT SET")
	viper.SetDefault("app_path", "NOT SET")
}

// GetConfig will generate the standard AppConfig
func GetConfig() AppConfig {
	AppName := flag.String("app_name", viper.GetString("app_name"), "Application name (REQUIRED)")
	AppPrefix := flag.String("app_prefix", viper.GetString("app_prefix"), "Application Prefix (REQUIRED)")
	AppPath := flag.String("app_path", viper.GetString("app_path"), "Application Path (default: github.com/mrzacarias)")
	flag.Parse()

	return AppConfig{
		AppName:   *AppName,
		AppPrefix: *AppPrefix,
		AppPath:   *AppPath,
	}
}

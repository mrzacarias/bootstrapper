package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mrzacarias/bootstrapper/config"
	"github.com/mrzacarias/bootstrapper/internal/service"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
}

func main() {
	var err error
	var appName, appPrefix, appPath string
	cfg := config.GetConfig()

	// Handling App Name and Prefix
	scanner := bufio.NewScanner(os.Stdin)
	if notSet(cfg.AppName) {
		fmt.Print("What is your app name? ==> ")
		if scanner.Scan() {
			appName = scanner.Text()
		}
	} else {
		appName = cfg.AppName
	}

	if notSet(cfg.AppPrefix) {
		fmt.Print("What is your app prefix? (3 letters are better) ==> ")
		if scanner.Scan() {
			appPrefix = scanner.Text()
		}
	} else {
		appPrefix = cfg.AppPrefix
	}

	if notSet(cfg.AppPath) {
		fmt.Print("What is your app path? (default: github.com/mrzacarias) ==> ")
		if scanner.Scan() {
			appPath = scanner.Text()
		}
		if appPath == "" {
			appPath = "github.com/mrzacarias"
		}
	} else {
		appPath = cfg.AppPath
	}

	// Ensuring readable naming
	appName = strings.ReplaceAll(strings.ToLower(appName), " ", "-")
	appPrefix = strings.ToUpper(appPrefix)
	appPath = strings.ToLower(appPath)

	err = service.BuildService(appName, appPrefix, appPath)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error on `service.BuildService`")
	}

	fmt.Printf("==> Done! Check the directory `%s`!\n", fmt.Sprintf("%s/src/%s/%s", os.Getenv("GOPATH"), appPath, appName))
}

func notSet(p string) bool {
	return p == "NOT SET"
}

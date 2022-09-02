package service

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrzacarias/bootstrapper/templates"
)

const executable = true
const notExecutable = false

// StatelessHTTPEndpoint and other App Template options
const (
	StatelessHTTPEndpoint = "stateless_http_endpoint"
	StatefulHTTPEndpoint  = "stateful_http_endpoint"
)

// BuildService will create the boilerplate service under appName directory
func BuildService(appName, appPrefix, appPath string) (err error) {
	basePath := fmt.Sprintf("%s/src/%s/%s", os.Getenv("GOPATH"), appPath, appName)
	fmt.Printf("===> Creating new service (%s) on `%s` directory!\n", appName, basePath)

	// Generate replacer
	replacer := strings.NewReplacer(
		"{app_name}", appName,
		"{app_prefix}", appPrefix,
		"{app_path}", appPath,
	)

	// Scaffolding general directories
	_ = os.MkdirAll(basePath, os.ModePerm)

	// Generate `root` files
	rootFiles := []string{".gitignore", "docker-compose.yml", "Dockerfile", "go.mod", "go.sum", "README.md"}
	err = loadAndGenerate(replacer, rootFiles, "", basePath, notExecutable)
	if err != nil {
		return err
	}

	// Generate `.github/workflows` files
	cmdGHWorkflowsDir := basePath + "/.github/workflows"
	_ = os.MkdirAll(cmdGHWorkflowsDir, os.ModePerm)
	cmdGHWorkflowsFiles := []string{"build-and-publish.yml"}
	err = loadAndGenerate(replacer, cmdGHWorkflowsFiles, ".github/workflows/", cmdGHWorkflowsDir, notExecutable)
	if err != nil {
		return err
	}

	// Generate `cmd/app` files
	cmdAppDir := basePath + "/cmd/app"
	_ = os.MkdirAll(cmdAppDir, os.ModePerm)
	cmdAppFiles := []string{"app.go", "app_test.go"}
	err = loadAndGenerate(replacer, cmdAppFiles, "cmd/app/", cmdAppDir, notExecutable)
	if err != nil {
		return err
	}

	// Generate `config` files
	configDir := basePath + "/config"
	_ = os.MkdirAll(configDir, os.ModePerm)
	configFiles := []string{"config.go", "config_test.go"}
	err = loadAndGenerate(replacer, configFiles, "config/", configDir, notExecutable)
	if err != nil {
		return err
	}

	// Generate `k8s/prod` files
	k8sProdDir := basePath + "/k8s/prod"
	_ = os.MkdirAll(k8sProdDir, os.ModePerm)
	k8sProdFiles := []string{"deployment.yml", "horizontal_autoscaler.yml", "ingress.yml", "network_policy.yml", "pod_disruption_budget.yml", "service_monitor.yml", "service.yml"}
	err = loadAndGenerate(replacer, k8sProdFiles, "k8s/prod/", k8sProdDir, notExecutable)
	if err != nil {
		return err
	}

	// Generate `script` files
	scriptDir := basePath + "/script"
	_ = os.MkdirAll(scriptDir, os.ModePerm)
	scriptFiles := []string{"nuke.sh", "run.sh", "setup.sh", "test.sh", "update.sh"}
	err = loadAndGenerate(replacer, scriptFiles, "script/", scriptDir, executable)
	if err != nil {
		return err
	}

	// Generate `internal/metrics` files
	metricsDir := basePath + "/internal/metrics"
	_ = os.MkdirAll(metricsDir, os.ModePerm)
	metricsFiles := []string{"metrics.go"}
	err = loadAndGenerate(replacer, metricsFiles, "internal/metrics/", metricsDir, notExecutable)
	if err != nil {
		return err
	}

	// Generate `internal/emoji` files
	emojiDir := basePath + "/internal/emoji"
	_ = os.MkdirAll(emojiDir, os.ModePerm)
	emojiFiles := []string{"emoji.go", "emoji_test.go"}
	err = loadAndGenerate(replacer, emojiFiles, "internal/emoji/", emojiDir, notExecutable)
	if err != nil {
		return err
	}

	// Generate `internal/mock` files
	mockDir := basePath + "/internal/mock"
	_ = os.MkdirAll(mockDir, os.ModePerm)
	mockFiles := []string{"emoji.go"}
	err = loadAndGenerate(replacer, mockFiles, "internal/mock/", mockDir, notExecutable)
	if err != nil {
		return err
	}

	// Generate `tools` files
	toolsDir := basePath + "/tools"
	_ = os.MkdirAll(toolsDir, os.ModePerm)
	toolsFiles := []string{"simulator.sh"}
	err = loadAndGenerate(replacer, toolsFiles, "tools/", toolsDir, executable)
	if err != nil {
		return err
	}

	fmt.Println("===> New application created!")
	return nil
}

func loadAndGenerate(replacer *strings.Replacer, list []string, templDir string, genDir string, exec bool) (err error) {
	for _, file := range list {
		// Load template
		template, err := templates.Asset(fmt.Sprintf("templates/service/stateless_http_endpoint/%s%s.btsppr", templDir, file))
		if err != nil {
			return err
		}

		generated := fmt.Sprintf("%s/%s", genDir, file)
		fmt.Printf("====> Generating '%s'...\n", generated)

		// Replace with new values
		result := replacer.Replace(string(template))

		// Create file (checking for executables first)
		if exec {
			err = os.WriteFile(generated, []byte(result), 0755)
		} else {
			err = os.WriteFile(generated, []byte(result), 0644)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

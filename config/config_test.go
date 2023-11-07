package config

import (
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/keshu12345/truecaller/toolkit"
)

func GetConfiguration(configDirPath, overridePath string) (*Configuration, error) {
	var conf Configuration
	if len(configDirPath) == 0 {
		configDirPath = os.Getenv("CONFIG_PATH")
	}
	err := toolkit.NewConfig(&conf, path.Join(configDirPath, serverYML), overridePath)
	return &conf, err
}

func TestGetConfiguration(t *testing.T) {
	// Create a temporary directory for testing and set an override path
	tempDir := t.TempDir()
	overridePath := path.Join(tempDir, "server.yml")

	// Create a sample server.yml file for testing
	serverYMLContent := `
environmentName: "test"
server:
  restServicePort: 8080
  readTimeout: 30
  writeTimeout: 30
  idleTimeout: 60
swagger:
  host: "localhost:8080"
`
	serverYMLPath := path.Join(tempDir, serverYML)

	if err := os.WriteFile(serverYMLPath, []byte(serverYMLContent), 0644); err != nil {
		t.Fatalf("Failed to create server.yml: %v", err)
	}

	// Set the CONFIG_PATH environment variable to the test directory
	oldConfigPath := os.Getenv("CONFIG_PATH")
	os.Setenv("CONFIG_PATH", tempDir)
	defer os.Setenv("CONFIG_PATH", oldConfigPath)

	// Get the Configuration struct
	config, err := GetConfiguration(tempDir, overridePath)

	// Check if there was no error
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check if the individual fields of the Configuration struct match the expected values
	expectedConf := Configuration{
		EnvironmentName: "test",
		Server: Server{
			RestServicePort: 8080,
			ReadTimeout:     30,
			WriteTimeout:    30,
			IdleTimeout:     60,
		},
		Swagger: Swagger{
			Host: "localhost:8080",
		},
	}

	if !reflect.DeepEqual(config, &expectedConf) {
		t.Errorf("Configuration does not match the expected values. Got %+v, expected %+v", config, &expectedConf)
	}
}

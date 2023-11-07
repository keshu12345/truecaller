package toolkit

import (
	"testing"
)

type AppConfig struct {
	// Define fields and tags for your configuration struct
}

// Define a test struct to hold test cases
type NewConfigTestCase struct {
	Name         string
	Config       interface{}
	ConfigPath   string
	OverridePath string
	EnvVars      []map[string]string
	ExpectError  bool
}

func TestNewConfig(t *testing.T) {
	testCases := []NewConfigTestCase{
		{
			Name:         "Test case 1: Successful configuration loading",
			Config:       &AppConfig{},
			ConfigPath:   "config/config.yaml",
			OverridePath: "config/override.yaml",
			EnvVars: []map[string]string{
				{
					"some_key": "SOME_ENV_VAR",
				},
			},
			ExpectError: false,
		},
		{
			Name:         "Test case 2: Error in config loading",
			Config:       &AppConfig{},
			ConfigPath:   "",
			OverridePath: "",
			EnvVars:      nil,
			ExpectError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := NewConfig(tc.Config, tc.ConfigPath, tc.OverridePath, tc.EnvVars...)
			if tc.ExpectError {
				if err == nil {
					t.Error("Expected an error but got nil.")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

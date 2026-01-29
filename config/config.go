/*
Config Module - Application configuration management

Functions:
  - LoadConfig: Loads the JSON configuration file from disk
  - ResolveProfile: Resolves a profile name to its full configuration
  - EnsureDataDirectory: Creates the data directory if it doesn't exist
*/
package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Profile represents a user profile configuration
type Profile struct {
	Name     string `json:"name"`
	DataPath string `json:"data_path"`
}

// Config represents the application configuration
type Config struct {
	Profiles map[string]Profile `json:"profiles"`
}

// ProfileInfo contains the resolved profile information
type ProfileInfo struct {
	AppTitle string
	DataPath string
}

// CreateConfigIfMissing creates a default configuration file if it doesn't exist
func CreateConfigIfMissing(configPath string) {
	if _, err := os.Stat(configPath); err == nil {
		return // File exists, nothing to do
	}

	configTemplate := `{
  "profiles": {
    "default": {
      "name": "Personal",
      "data_path": "data/personal"
    },
    "business": {
      "name": "Business Account",
      "data_path": "data/business"
    },
    "gaming": {
      "name": "Gaming Community",
      "data_path": "data/gaming"
    }
  }
}`

	err := os.WriteFile(configPath, []byte(configTemplate), 0644)
	if err != nil {
		// Log error or handle it, but for now we'll just print since we can't panic in production easily
		println("Error creating default config:", err.Error())
	}
}

func LoadConfig(configPath string) (*Config, error) {
	var cfg Config
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return &cfg, err
	}

	err = json.Unmarshal(configFile, &cfg)
	return &cfg, err
}

func ResolveProfile(cfg *Config, profileName, exeDir string) ProfileInfo {
	var info ProfileInfo

	if profile, ok := cfg.Profiles[profileName]; ok {
		info.AppTitle = "WhatsX - " + profile.Name
		if filepath.IsAbs(profile.DataPath) {
			info.DataPath = profile.DataPath
		} else {
			info.DataPath = filepath.Join(exeDir, profile.DataPath)
		}
	} else {
		// Fallback defaults
		if profileName == "default" {
			info.AppTitle = "WhatsX"
			info.DataPath = filepath.Join(exeDir, "data", "default")
		} else {
			info.AppTitle = "WhatsX - " + profileName
			info.DataPath = filepath.Join(exeDir, "data", profileName)
		}
	}

	return info
}

func EnsureDataDirectory(dataPath string) error {
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return os.MkdirAll(dataPath, 0755)
	}
	return nil
}

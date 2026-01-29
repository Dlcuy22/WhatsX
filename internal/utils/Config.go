package utils

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

// LoadConfig loads the configuration file from the given path
func LoadConfig(configPath string) (*Config, error) {
	var config Config
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return &config, err
	}

	err = json.Unmarshal(configFile, &config)
	return &config, err
}

// ResolveProfile resolves a profile name to its configuration
func ResolveProfile(config *Config, profileName, exeDir string) ProfileInfo {
	var info ProfileInfo

	if profile, ok := config.Profiles[profileName]; ok {
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

// EnsureDataDirectory creates the data directory if it doesn't exist
func EnsureDataDirectory(dataPath string) error {
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return os.MkdirAll(dataPath, 0755)
	}
	return nil
}

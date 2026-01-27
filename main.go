package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Config structures
type Profile struct {
	Name     string `json:"name"`
	DataPath string `json:"data_path"`
}

type Config struct {
	Profiles map[string]Profile `json:"profiles"`
}

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/windows/icon.ico
var iconData []byte

func setupSystray(ctx context.Context, title string) {
	systray.Run(func() {
		systray.SetIcon(iconData)
		systray.SetTitle(title)
		systray.SetTooltip(title)

		mShow := systray.AddMenuItem("Open WhatsX", "Show the main window")
		mQuit := systray.AddMenuItem("Quit", "Quit the application")

		go func() {
			for {
				select {
				case <-mShow.ClickedCh:
					runtime.WindowShow(ctx)
				case <-mQuit.ClickedCh:
					systray.Quit()
					runtime.Quit(ctx)
				}
			}
		}()
	}, func() {
		// onExit cleanup
	})
}

func main() {
	// Parse CLI flags
	profileFlag := flag.String("profile", "default", "The profile name to use")
	flag.Parse()

	// Create an instance of the app structure
	app := NewApp()

	// Get executable directory
	exePath, err := os.Executable()
	if err != nil {
		println("Error getting executable path:", err.Error())
		return
	}
	exeDir := filepath.Dir(exePath)

	// Load Configuration
	configPath := filepath.Join(exeDir, "WhatsX.config.json")
	var config Config
	configFile, err := os.ReadFile(configPath)
	if err == nil {
		_ = json.Unmarshal(configFile, &config)
	}

	// Resolve Profile
	profileName := *profileFlag
	var appTitle string
	var dataPath string

	if profile, ok := config.Profiles[profileName]; ok {
		appTitle = "WhatsX - " + profile.Name
		if filepath.IsAbs(profile.DataPath) {
			dataPath = profile.DataPath
		} else {
			dataPath = filepath.Join(exeDir, profile.DataPath)
		}
	} else {
		// Fallback defaults
		if profileName == "default" {
			appTitle = "WhatsX"
			dataPath = filepath.Join(exeDir, "data", "default")
		} else {
			appTitle = "WhatsX - " + profileName
			dataPath = filepath.Join(exeDir, "data", profileName)
		}
	}

	// Create data directory if it doesn't exist
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		os.MkdirAll(dataPath, 0755)
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:             appTitle,
		Width:             1000,
		Height:            600,
		HideWindowOnClose: true, // Hide instead of close
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			// Initialize Systray with dynamic title
			go setupSystray(ctx, appTitle)
		},
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewUserDataPath: dataPath,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

package main

import (
	"WhatsX/internal/utils"
	"context"
	"embed"
	"flag"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	_ "github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

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
	config, _ := utils.LoadConfig(configPath)

	// Resolve Profile
	profileName := *profileFlag
	profileInfo := utils.ResolveProfile(config, profileName, exeDir)

	// Create data directory if it doesn't exist
	_ = utils.EnsureDataDirectory(profileInfo.DataPath)

	// Create application with options
	err = wails.Run(&options.App{
		Title:             profileInfo.AppTitle,
		Width:             1000,
		Height:            600,
		HideWindowOnClose: hideWindowOnClose, // Platform-specific behavior
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			// Initialize Systray (platform-specific implementation)
			// We run it in a goroutine so it doesn't block startup on platforms where it runs in-process.
			go setupSystray(ctx, profileInfo.AppTitle)
		},
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewUserDataPath: profileInfo.DataPath,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

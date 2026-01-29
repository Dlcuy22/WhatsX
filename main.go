package main

import (
	"WhatsX/config"
	"context"
	"embed"
	"flag"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

// appContext stores the application context for menu callbacks
var appContext context.Context

func createAppMenu() *menu.Menu {
	appMenu := menu.NewMenu()

	// File Menu
	fileMenu := appMenu.AddSubmenu("File")
	fileMenu.AddText("Reload", keys.CmdOrCtrl("r"), func(_ *menu.CallbackData) {
		if appContext != nil {
			runtime.WindowReload(appContext)
		}
	})
	fileMenu.AddSeparator()
	fileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		if appContext != nil {
			runtime.Quit(appContext)
		}
	})

	// View Menu
	viewMenu := appMenu.AddSubmenu("View")
	viewMenu.AddText("Toggle Fullscreen", keys.Key("F11"), func(_ *menu.CallbackData) {
		if appContext != nil {
			runtime.WindowToggleMaximise(appContext)
		}
	})
	viewMenu.AddText("Minimize", keys.CmdOrCtrl("m"), func(_ *menu.CallbackData) {
		if appContext != nil {
			runtime.WindowMinimise(appContext)
		}
	})

	// Help Menu
	helpMenu := appMenu.AddSubmenu("Help")
	helpMenu.AddText("About WhatsX", nil, func(_ *menu.CallbackData) {
		if appContext != nil {
			runtime.MessageDialog(appContext, runtime.MessageDialogOptions{
				Type:    runtime.InfoDialog,
				Title:   "About WhatsX",
				Message: "WhatsX - WhatsApp Desktop Wrapper\nVersion 1.0.0",
			})
		}
	})

	return appMenu
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
	config.CreateConfigIfMissing(configPath)
	cfg, _ := config.LoadConfig(configPath)

	// Resolve Profile
	profileName := *profileFlag
	profileInfo := config.ResolveProfile(cfg, profileName, exeDir)

	// Create data directory if it doesn't exist
	_ = config.EnsureDataDirectory(profileInfo.DataPath)

	// Create native menu
	appMenu := createAppMenu()

	// Create application with options
	err = wails.Run(&options.App{
		Title:             profileInfo.AppTitle,
		Width:             1000,
		Height:            600,
		HideWindowOnClose: hideWindowOnClose,
		Menu:              appMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			appContext = ctx
			app.startup(ctx)
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

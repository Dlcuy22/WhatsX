//go:build windows

package main

import (
	"context"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// On Windows, we want to hide the window on close and use systray
const hideWindowOnClose = true

// setupSystray runs the systray on Windows. It uses iconData (embedded in main.go).
func setupSystray(ctx context.Context, title string) {
	// Run systray in a goroutine so it doesn't block OnStartup (ok on Windows)
	go systray.Run(func() {
		if len(iconData) > 0 {
			systray.SetIcon(iconData)
		}
		systray.SetTitle(title)
		systray.SetTooltip(title)

		mShow := systray.AddMenuItem("Open WhatsX", "Show the main window")
		mQuit := systray.AddMenuItem("Quit", "Quit the application")

		for {
			select {
			case <-mShow.ClickedCh:
				runtime.WindowShow(ctx)
			case <-mQuit.ClickedCh:
				systray.Quit()
				runtime.Quit(ctx)
				return
			}
		}
	}, func() {
		// onExit (optional)
	})
}

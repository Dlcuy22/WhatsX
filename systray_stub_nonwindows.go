//go:build !windows

package main

import (
	"context"
)

// On Linux, we don't want to hide the window on close - we want to exit
const hideWindowOnClose = false

// no-op implementation for non-windows platforms
func setupSystray(ctx context.Context, title string) {
	// Intentionally empty: systray is disabled on non-windows build
}

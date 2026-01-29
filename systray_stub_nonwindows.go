//go:build !windows

package main

import (
	"context"
)

const hideWindowOnClose = false

// no-op implementation for non-windows platforms
func setupSystray(ctx context.Context, title string) {
	// Intentionally empty: systray is disabled on non-windows build
}

package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) OpenURL(url string) {
	runtime.BrowserOpenURL(a.ctx, url)
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

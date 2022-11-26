package app

import (
	"context"

	"golang.design/x/mainthread"
)

func (a *App) handleAppStartup(ctx context.Context) {
	a.SetContext(ctx)
	mainthread.Init(a.RegisterHotkey)
}

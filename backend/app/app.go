package app

import (
	"context"
	"log"

	hk "golang.design/x/hotkey"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) GetContext() context.Context {
	return a.ctx
}

func (a *App) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) RegisterHotkey() {
	registerHotkey(a)
}

func registerHotkey(a *App) {
	hkey := hk.New([]hk.Modifier{hk.ModCtrl, hk.ModShift}, hk.KeyS)
	err := hkey.Register()
	if err != nil {
		log.Panicln("failed to register hotkey", err)
		return
	}

	<-hkey.Keyup()
	log.Println("hotkey triggered")

	hkey.Unregister()

	// reattach listener
	registerHotkey(a)
}

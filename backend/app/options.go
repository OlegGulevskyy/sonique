package app

import (
	"embed"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func (a *App) GetOptions(assets embed.FS) *options.App {
	return &options.App{
		Title:  "Sonique",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: a.handleAppStartup,
		Bind: []interface{}{
			a,
		},
	}
}

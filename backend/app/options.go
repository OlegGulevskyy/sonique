package app

import (
	"embed"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func (a *App) GetOptions(assets embed.FS) *options.App {
	macOpts := &mac.Options{
		TitleBar: &mac.TitleBar{
			HideTitle:                  true,
			TitlebarAppearsTransparent: true,
			HideToolbarSeparator:       true,
		},
	}

	linuxOpts := &linux.Options{
		WindowIsTranslucent: true,
	}

	windowsOpts := &windows.Options{
		WebviewIsTransparent: false,
		WindowIsTranslucent:  false,
		DisableWindowIcon:    true,
	}

	return &options.App{
		Title:  "Sonique",
		Width:  1024,
		Height: 768,
		DisableResize: true,
		Frameless: true,
		AlwaysOnTop: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.handleAppStartup,
		Bind: []interface{}{
			a,
		},
		Mac:     macOpts,
		Linux:   linuxOpts,
		Windows: windowsOpts,
	}
}

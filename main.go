package main

import (
	"embed"

	"github.com/OlegGulevskyy/sonique/backend/app"

	"github.com/wailsapp/wails/v2"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	a := app.NewApp()

	// Create application with options
	err := wails.Run(a.GetOptions(assets))

	if err != nil {
		println("Error:", err.Error())
	}
}

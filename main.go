package main

import (
	"embed"
	"log"

	"github.com/amrllkmn/sorted/internal/database"
	"github.com/amrllkmn/sorted/internal/tasks"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {

	// Init DB
	db := database.InitDB()
	dbErr := db.AutoMigrate(&tasks.Task{})
	if dbErr != nil {
		log.Fatal("Migration failed")
	}

	taskRepo := tasks.NewSQLiteTaskRepository(db)

	taskService := tasks.NewTaskService(taskRepo)
	// Create an instance of the app structure
	app := NewApp(taskService)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Sorted.",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

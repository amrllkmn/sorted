package main

import (
	"context"
	"log"

	"github.com/amrllkmn/sorted/internal/tasks"
)

// App struct
type App struct {
	ctx         context.Context
	taskService tasks.TasksService
}

// NewApp creates a new App application struct
func NewApp(taskService tasks.TasksService) *App {
	return &App{
		taskService: taskService,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAllTasks() []tasks.Task {
	tasks, err := a.taskService.GetTasks()

	if err != nil {
		log.Fatal(err)
	}

	return tasks
}

func (a *App) CreateTask(description string) (*tasks.Task, error) {
	task, err := a.taskService.CreateTask(description)

	if err != nil {
		log.Fatal(err)

	}

	return task, nil
}

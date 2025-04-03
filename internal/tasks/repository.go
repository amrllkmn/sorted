package tasks

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint   `json:"id"`
	Description string `json:"description"`
	Urgent      *bool  `json:"urgent"`
	Important   *bool  `json:"important"`
}

type TaskRepository interface {
	GetAllTasks(tasks *[]Task) error
	CreateTask(taskDesc string) (*Task, error)
	// DeleteTaskById(id int) error
}

type SQLiteTaskRepository struct {
	DB *gorm.DB
}

func (tr *SQLiteTaskRepository) GetAllTasks(tasks *[]Task) error {
	result := tr.DB.Find(tasks)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (tr *SQLiteTaskRepository) CreateTask(taskDesc string) (*Task, error) {
	task := Task{Description: taskDesc}
	result := tr.DB.Create(&task)

	if result.Error != nil {
		return &Task{}, result.Error
	}
	return &task, nil
}

// func (tr *SQLiteTaskRepository) DeleteTaskById(id int) error

func NewSQLiteTaskRepository(db *gorm.DB) TaskRepository {
	return &SQLiteTaskRepository{
		DB: db,
	}
}

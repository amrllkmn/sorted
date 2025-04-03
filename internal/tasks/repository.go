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
	// CreateTask(task *Task) error
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

// func (tr *SQLiteTaskRepository) CreateTask(task *Task) error
// func (tr *SQLiteTaskRepository) DeleteTaskById(id int) error

func NewSQLiteTaskRepository(db *gorm.DB) TaskRepository {
	return &SQLiteTaskRepository{
		DB: db,
	}
}

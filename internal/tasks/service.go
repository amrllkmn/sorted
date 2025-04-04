package tasks

type TasksService interface {
	GetTasks() ([]Task, error)
	// DeleteTask()
	CreateTask(description string) (*Task, error)
}

type Service struct {
	repo TaskRepository
}

func (svc *Service) GetTasks() ([]Task, error) {
	var tasks []Task

	err := svc.repo.GetAllTasks(&tasks)

	if err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

// func (svc *Service) DeleteTask()
func (svc *Service) CreateTask(description string) (*Task, error) {
	task, err := svc.repo.CreateTask(description)

	if err != nil {
		return &Task{}, err
	}

	return task, nil
}

func NewTaskService(repo TaskRepository) TasksService {
	return &Service{
		repo: repo,
	}
}

package tasks

type TasksService interface {
	GetTasks() ([]Task, error)
	// DeleteTask()
	// CreateTask()
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
// func (svc *Service) CreateTask()

func NewTaskService(repo TaskRepository) TasksService {
	return &Service{
		repo: repo,
	}
}

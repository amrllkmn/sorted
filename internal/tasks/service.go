package tasks

type TasksService interface {
	GetTasks() ([]Task, error)
	DeleteTask(id int) error
	CreateTask(description string) (*Task, error)
	UpdateTaskStatus(id int, urgent bool, important bool) error
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

func (svc *Service) DeleteTask(id int) error {
	err := svc.repo.DeleteTaskById(id)

	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) CreateTask(description string) (*Task, error) {
	task, err := svc.repo.CreateTask(description)

	if err != nil {
		return &Task{}, err
	}

	return task, nil
}

func (svc *Service) UpdateTaskStatus(id int, urgent bool, important bool) error {
	err := svc.repo.UpdateTaskStatus(id, urgent, important)

	if err != nil {
		return err
	}

	return nil
}

func NewTaskService(repo TaskRepository) TasksService {
	return &Service{
		repo: repo,
	}
}

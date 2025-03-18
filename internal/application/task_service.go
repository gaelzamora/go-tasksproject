package application

import (
	"github.com/gaelzamora/go-rest-crud/internal/domain"
	"github.com/gaelzamora/go-rest-crud/internal/ports"
)

type TaskService struct {
	repo ports.TaskRepository
}

func NewTaskService(repo ports.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAllTasks() ([]domain.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetTaskByID(id uint) (domain.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) CreateTask(task domain.Task) error {
	return s.repo.Create(task)
}

func (s *TaskService) UpdateTask(task domain.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}

package application

import (
	"fmt"

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
	
	fmt.Println("------------------")
	fmt.Println(id)
	return s.repo.GetByID(id)
}

func (s *TaskService) GetAllTasksByUser(id uint) ([]domain.Task, error) {
	return s.repo.GetAllTasksById(id)
}

func (s *TaskService) CreateTask(userID uint, task *domain.Task) (domain.Task, error) {
    task.UserID = userID
	err := s.repo.Create(*task)
    if err != nil {
        return domain.Task{}, err
    }
    return *task, nil
}

func (s *TaskService) UpdateTask(task domain.Task) error {
	return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}

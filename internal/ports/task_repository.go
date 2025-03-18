package ports

import "github.com/gaelzamora/go-rest-crud/internal/domain"

type TaskRepository interface {
	GetAll() ([]domain.Task, error)
	GetByID(id uint) (domain.Task, error)
	Create(task domain.Task) error
	Update(task domain.Task) error
	Delete(id uint) error
}
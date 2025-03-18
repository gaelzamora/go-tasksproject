package database

import (
	"github.com/gaelzamora/go-rest-crud/internal/ports"
	"github.com/gaelzamora/go-rest-crud/internal/domain"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ports.TaskRepository {
	return &TaskRepositoryImpl{DB: db}
}

func (r *TaskRepositoryImpl) GetAll() ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepositoryImpl) GetByID(id uint) (domain.Task, error) {
	var task domain.Task
	err := r.DB.First(&task, id).Error
	return task, err
}

func (r *TaskRepositoryImpl) Create(task domain.Task) error {
	return r.DB.Create(&task).Error
}

func (r *TaskRepositoryImpl) Update(task domain.Task) error {
	return r.DB.Save(&task).Error
}

func (r *TaskRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&domain.Task{}, id).Error
}
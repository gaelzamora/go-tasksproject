package database

import (
	"fmt"

	"github.com/gaelzamora/go-rest-crud/internal/domain"
	"github.com/gaelzamora/go-rest-crud/internal/ports"
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
	fmt.Println("--------------------")
	fmt.Println("Task: ", task)
	return task, err
}

func (r *TaskRepositoryImpl) GetAllTasksById(user_id uint) ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.DB.Where("user_id = ?", user_id).Find(&tasks).Error

	return tasks, err
}

func (r *TaskRepositoryImpl) Create(task domain.Task) error {
	err := r.DB.Create(&task).Error
	if err != nil {
		return err
	}
	fmt.Printf("Task after DB.Create: %+v\n", task) // Depuración
	return nil
}

func (r *TaskRepositoryImpl) Update(task domain.Task) error {
	return r.DB.Save(&task).Error
}

func (r *TaskRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&domain.Task{}, id).Error
}

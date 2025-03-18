package database

import (
	"github.com/gaelzamora/go-rest-crud/internal/domain"
	"github.com/gaelzamora/go-rest-crud/internal/ports"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := r.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}
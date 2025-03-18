package ports

import "github.com/gaelzamora/go-rest-crud/internal/domain"

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
	Create(user *domain.User) error
}
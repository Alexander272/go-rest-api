package store

import "github.com/golang-api-server/iternal/app/models"

type UserRepository interface {
	Create(user *models.User) error
	Find(int) (*models.User, error)
	FindByEmail(string) (*models.User, error)
}

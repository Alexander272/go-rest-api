package teststore

import (
	"github.com/golang-api-server/iternal/app/models"
	"github.com/golang-api-server/iternal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*models.User
}

func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.ID] = u
	u.ID = len(r.users)
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrorRecordNotFound
}

func (r *UserRepository) Find(id int) (*models.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrorRecordNotFound
	}

	return u, nil
}
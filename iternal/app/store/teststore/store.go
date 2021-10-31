package teststore

import (
	"github.com/golang-api-server/iternal/app/models"
	"github.com/golang-api-server/iternal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}


func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*models.User),
	}
	return s.userRepository
}

package teststore_test

import (
	"github.com/golang-api-server/iternal/app/models"
	"github.com/golang-api-server/iternal/app/store"
	"github.com/golang-api-server/iternal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	user := models.TestUser(t)
	assert.NoError(t, s.User().Create(user))
	assert.NotNil(t, user)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	email := "user@example.com"
	s := teststore.New()
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrorRecordNotFound.Error())


	u := models.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	email := "user@example.com"
	s := teststore.New()

	u1 := models.TestUser(t)
	u1.Email = email
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
}
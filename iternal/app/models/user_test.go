package models_test

import (
	"github.com/golang-api-server/iternal/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct{
		name string
		user func() *models.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *models.User {
				return models.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "with encrypted password",
			user: func() *models.User {
				u := models.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "EncryptedPassword"
				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			user: func() *models.User {
				u := models.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *models.User {
				u := models.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *models.User {
				u := models.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			user: func() *models.User {
				u := models.TestUser(t)
				u.Password = "123"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.user().Validate())
			} else {
				assert.Error(t, tc.user().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := models.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotNil(t, u.EncryptedPassword)
}

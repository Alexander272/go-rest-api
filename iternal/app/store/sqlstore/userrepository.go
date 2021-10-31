package sqlstore

import (
	"database/sql"
	"github.com/golang-api-server/iternal/app/models"
	"github.com/golang-api-server/iternal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1", email,
	).Scan(&user.ID, &user.Email, &user.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Find(id int) (*models.User, error) {
	user := &models.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1", email,
	).Scan(&user.ID, &user.Email, &user.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrorRecordNotFound
		}
		return nil, err
	}

	return user, nil
}
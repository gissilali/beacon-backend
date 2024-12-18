package data

import (
	"beacon.silali.com/internal/api/dtos"
	"database/sql"
	"errors"
)

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Create(user *dtos.User) (*dtos.User, error) {
	err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email`
	args := []interface{}{user.Name, user.Email, user.Password}

	err = u.DB.QueryRow(query, args...).Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserModel) Update() (UserModel, error) {
	return UserModel{}, nil
}

func (u *UserModel) Delete() (UserModel, error) {
	return UserModel{}, nil
}

func (u *UserModel) GetByEmail(email string) (*dtos.User, error) {
	query := `SELECT id, name, email, created_at,password FROM users WHERE email = $1 LIMIT 1`
	var user dtos.User
	err := u.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.Password)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, err
}

func (u *UserModel) GetById(id int) (*dtos.User, error) {
	query := `SELECT id, name, email, created_at FROM users WHERE email = $1 LIMIT 1`
	var user dtos.User
	err := u.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, err
}

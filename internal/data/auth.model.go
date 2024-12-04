package data

import (
	"database/sql"
	"time"
)

type AuthTokenModel struct {
	DB *sql.DB
}

func (model *AuthTokenModel) Create(refreshToken string, userId int64) error {

	if err := model.DeleteByUserID(userId); err != nil {
		return err
	}

	query := `INSERT INTO auth_tokens (user_id, refresh_token, expires_at) VALUES ($1, $2, $3) RETURNING user_id, refresh_token, expires_at`
	expiresAt := time.Now().Add(14 * 24 * time.Hour)

	var returnedUserId int64
	var returnedRefreshToken string
	var returnedExpiresAt time.Time

	err := model.DB.QueryRow(query, userId, refreshToken, expiresAt).Scan(&returnedUserId, &returnedRefreshToken, &returnedExpiresAt)

	if err != nil {
		return err
	}

	return nil
}

func (model *AuthTokenModel) DeleteByUserID(userId int64) error {
	query := `DELETE FROM auth_tokens WHERE user_id = $1`

	_, err := model.DB.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}

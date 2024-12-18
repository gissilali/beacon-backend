package data

import (
	"beacon.silali.com/internal/api/dtos"
	"database/sql"
	"fmt"
)

type AccessKeyModel struct {
	DB *sql.DB
}

func (model *AccessKeyModel) CreateAccessKey(name string, key string, userId int64) (*dtos.AccessKey, error) {
	query := `INSERT INTO api_keys (name, key, user_id) VALUES ($1, $2, $3) RETURNING id, user_id, name, key, last_used_at, revoked`

	accessKey := &dtos.AccessKey{}

	err := model.DB.QueryRow(query, name, key, userId).Scan(&accessKey.ID, &accessKey.UserId, &accessKey.Name, &accessKey.Key, &accessKey.LastUsed, &accessKey.Revoked)

	if err != nil {
		return nil, err
	}

	return accessKey, nil
}

func (model *AccessKeyModel) GetUserAccessKeys(userId int64) ([]dtos.AccessKey, error) {
	query := `SELECT id,user_id,name,revoked,key,last_used_at FROM api_keys WHERE user_id = $1`
	rows, err := model.DB.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query api_keys: %w", err)
	}
	defer rows.Close()
	accessKeys := make([]dtos.AccessKey, 0)

	for rows.Next() {
		var accessKey dtos.AccessKey
		if err := rows.Scan(&accessKey.ID, &accessKey.UserId, &accessKey.Name, &accessKey.Revoked, &accessKey.Key, &accessKey.LastUsed); err != nil {
			return nil, fmt.Errorf("failed to scan api_keys: %w", err)
		}
		accessKeys = append(accessKeys, accessKey)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return accessKeys, nil
}

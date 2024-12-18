package dtos

import "time"

type CreateAccessKeyRequest struct {
	Name string `json:"name"`
}

type AccessKey struct {
	ID       string     `json:"id"`
	UserId   string     `json:"userId"`
	Name     string     `json:"name"`
	Key      string     `json:"key"`
	Revoked  bool       `json:"revoked"`
	LastUsed *time.Time `json:"lastUsed"`
}

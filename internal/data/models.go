package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	User      UserModel
	AuthToken AuthTokenModel
	Workspace WorkspaceModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		User:      UserModel{DB: db},
		AuthToken: AuthTokenModel{DB: db},
		Workspace: WorkspaceModel{DB: db},
	}
}

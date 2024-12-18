package data

import (
	"beacon.silali.com/internal/api/dtos"
	"database/sql"
	"fmt"
)

type WorkspaceModel struct {
	DB *sql.DB
}

func (model *WorkspaceModel) GetUserWorkspaces(userId int64) ([]dtos.Workspace, error) {
	query := `SELECT id,name,owner_id FROM workspaces WHERE id IN (SELECT workspace_id FROM workspace_members WHERE user_id = $1)`
	rows, err := model.DB.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to query workspaces: %w", err)
	}

	defer rows.Close()

	workspaces := make([]dtos.Workspace, 0)

	for rows.Next() {
		var workspace dtos.Workspace
		if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.Owner); err != nil {
			return nil, fmt.Errorf("failed to scan workspace: %w", err)
		}
		workspaces = append(workspaces, workspace)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return workspaces, nil
}

func (model *WorkspaceModel) CreateWorkspace(userId int64, name string) (*dtos.Workspace, error) {
	query := `INSERT INTO workspaces (owner_id, name) VALUES ($1, $2) RETURNING id, owner_id, name`

	workspace := &dtos.Workspace{}
	err := model.DB.QueryRow(query, userId, name).Scan(&workspace.ID, &workspace.Owner, &workspace.Name)

	if err != nil {
		return nil, err
	}

	attachErr := model.AttachWorkspaceToUser(userId, workspace.ID)
	if attachErr != nil {
		return nil, attachErr
	}

	return workspace, nil
}

func (model *WorkspaceModel) AttachWorkspaceToUser(userId int64, workspaceId int64) error {
	query := `INSERT INTO workspace_members (user_id, workspace_id) VALUES ($1, $2)`
	_, err := model.DB.Exec(query, userId, workspaceId)
	if err != nil {
		return err
	}

	return nil
}

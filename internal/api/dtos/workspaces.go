package dtos

type Workspace struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type CreateWorkspaceRequest struct {
	Name string `json:"name"`
}

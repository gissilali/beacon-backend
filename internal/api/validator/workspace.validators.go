package validator

import (
	"beacon.silali.com/internal/api/dtos"
)

func (v *Validator) ValidateCreateWorkspaceRequest(request dtos.CreateWorkspaceRequest) {
	v.Check(len(request.Name) > 0, "name", "Name should be provided")
}

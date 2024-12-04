package validator

import "beacon.silali.com/internal/data"

func (v *Validator) ValidateCreateWorkspaceRequest(request data.CreateWorkspaceRequest) {
	v.Check(len(request.Name) > 0, "name", "Name should be provided")
}

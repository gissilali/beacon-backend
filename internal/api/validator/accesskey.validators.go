package validator

import (
	"beacon.silali.com/internal/api/dtos"
)

func (v *Validator) ValidateCreateAccessKeyRequest(request *dtos.CreateAccessKeyRequest) {
	v.Check(len(request.Name) > 0, "name", "Name should be provided")
}

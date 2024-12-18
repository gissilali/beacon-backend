package validator

import (
	"beacon.silali.com/internal/api/dtos"
)

func (v *Validator) ValidateRegisterUserRequest(request *dtos.RegisterUserRequest) {
	v.Check(len(request.Name) > 0, "name", "name is required")
	v.Check(len(request.Email) > 0, "email", "email is required")
	v.Check(len(request.Password) > 0, "password", "password is required")
	v.CheckWith(v.IsEmailRule, request.Email, "email", "email {{.value}} is invalid")
	v.CheckWith(v.EmailExistsRule, request.Email, "email", "email {{.value}} already exists")
}

func (v *Validator) ValidateLoginUserRequest(request *dtos.LoginUserRequest) {
	v.Check(len(request.Email) > 0, "email", "email is required")
	v.Check(len(request.Password) > 0, "password", "password is required")
	v.CheckWith(v.IsEmailRule, request.Email, "email", "email {{.value}} is invalid")
	v.CheckWith(v.EmailDoesNotExistRule, request.Email, "email", "email {{.value}} does not exist in our records")
}

func (v *Validator) IsEmailRule(email string) bool {
	return emailRegex.MatchString(email)
}

func (v *Validator) EmailExistsRule(email string) bool {
	value, err := v.Models.User.GetByEmail(email)

	if err != nil || value == nil {
		return true
	}

	return value.Email != email
}

func (v *Validator) EmailDoesNotExistRule(email string) bool {
	value, err := v.Models.User.GetByEmail(email)

	if err != nil || value == nil {
		return false
	}

	return value.Email == email
}

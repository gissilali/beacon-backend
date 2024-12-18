package validator

import (
	"beacon.silali.com/internal/api/data"
	"bytes"
	"regexp"
	"text/template"
)

type Validation struct {
	Validator Validator
}

type Validator struct {
	Errors map[string]string
	Models data.Models
}

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type ValidationRule func(value string) bool

func New(models data.Models) *Validator {
	return &Validator{
		Errors: make(map[string]string),
		Models: models,
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key string, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) ClearErrors() {
	v.Errors = make(map[string]string)
}

func (v *Validator) Check(isOk bool, attribute string, message string) {
	if isOk == false {
		v.AddError(attribute, message)
	}
}

func (v *Validator) CheckWith(validationFunc ValidationRule, value string, attribute string, message string) {
	templateMap := map[string]string{
		"value":     value,
		"attribute": attribute,
	}

	tmpl, err := template.New("validationMessage").Parse(message)
	if err != nil {
		panic(err)
	}

	var result bytes.Buffer

	if err := tmpl.Execute(&result, templateMap); err != nil {
		panic(err)
	}

	v.Check(validationFunc(value), attribute, result.String())
}

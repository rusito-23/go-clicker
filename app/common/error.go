package common

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// ValidatorError -
// Contains customized validation error information
type ValidatorError struct {
	Errors map[string]interface{} `json:"errors"`
}

// NewValidatorError -
// Creates the validation error based on the Bind errors
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) ValidatorError {
	res := ValidatorError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}
	}
	return res
}

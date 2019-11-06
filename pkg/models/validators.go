package models

import "regexp"

var emailRegExp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type ValidationErrors struct {
	Errors []*ValidationError `json:"errors"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidationErrors() *ValidationErrors {
	return &ValidationErrors{}
}

func (e *ValidationErrors) Add(field, msg string) {
	e.Errors = append(e.Errors, &ValidationError{
		Field:   field,
		Message: msg,
	})
}

func (e ValidationErrors) Present() bool {
	return len(e.Errors) != 0
}

func (u User) Validate(validateName bool) *ValidationErrors {
	errs := u.ValidateEmail()

	if validateName && u.Name == "" {
		errs.Add("name", "field is required")
	}
	if validateName && u.Name != "" && len(u.Name) > 50 {
		errs.Add("name", "max length is 50 characters")
	}
	if u.Password == "" {
		errs.Add("password", "field is required")
	}
	if u.Password != "" && len(u.Password) < 5 || len(u.Password) > 30 {
		errs.Add("password", "min length is 5 characters, max length is 30 characters")
	}
	return errs
}

func (u User) ValidateEmail() *ValidationErrors {
	errs := NewValidationErrors()

	if u.Email == "" {
		errs.Add("email", "field is required")
	}
	if u.Email != "" && len(u.Email) > 120 {
		errs.Add("email", "max length is 120 characters")
	}
	if u.Email != "" && !emailRegExp.MatchString(u.Email) {
		errs.Add("email", "invalid email address")
	}
	return errs
}

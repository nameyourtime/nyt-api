package models

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

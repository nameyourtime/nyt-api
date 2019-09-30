package main

import (
	"fmt"
	"nameyourtime.com/api/pkg/models"
	"net/http"
	"runtime/debug"
)

type ApiError struct {
	Code             string                    `json:"code"`
	Message          string                    `json:"message"`
	ValidationErrors []*models.ValidationError `json:"validation_errors,omitempty"`
}

func (e ApiError) str() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	_ = app.errorLog.Output(2, trace)
	w.WriteHeader(http.StatusInternalServerError)
	writeJsonResponse(w, &ApiError{
		Code:    "001",
		Message: "server error",
	})
}

func (app *application) badRequest(w http.ResponseWriter) {
	handleError(app, w, http.StatusBadRequest, &ApiError{
		Code:    "002",
		Message: "can't read request body",
	})
}

func (app *application) validationError(w http.ResponseWriter, errs *models.ValidationErrors) {
	handleError(app, w, http.StatusBadRequest, &ApiError{
		Code:             "003",
		Message:          "request validation failed",
		ValidationErrors: errs.Errors,
	})
}

func handleError(app *application, w http.ResponseWriter, status int, e *ApiError) {
	w.WriteHeader(status)
	app.errorLog.Println(e.str())
	writeJsonResponse(w, e)
}

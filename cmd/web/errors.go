package main

import (
	"fmt"
	"log"
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

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	_ = app.errorLog.Output(2, trace)
	w.WriteHeader(http.StatusInternalServerError)
	writeJsonResponse(w, &ApiError{
		Code:    "001",
		Message: "server error",
	})
}

func (app *Application) badRequest(w http.ResponseWriter) {
	handleError(app, w, http.StatusBadRequest, &ApiError{
		Code:    "002",
		Message: "can't read request body",
	})
}

func (app *Application) validationError(w http.ResponseWriter, errs *models.ValidationErrors) {
	handleError(app, w, http.StatusBadRequest, &ApiError{
		Code:             "003",
		Message:          "request validation failed",
		ValidationErrors: errs.Errors,
	})
}

func (app *Application) duplicatedEmail(w http.ResponseWriter) {
	handleError(app, w, http.StatusBadRequest, &ApiError{
		Code:    "007",
		Message: "user with this email already registered",
	})
}

func handleError(app *Application, w http.ResponseWriter, status int, e *ApiError) {
	w.WriteHeader(status)
	app.errorLog.Println(e.str())
	writeJsonResponse(w, e)
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	log.Printf("%s\n%s", err.Error(), debug.Stack())
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}

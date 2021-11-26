package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *Application) routes() http.Handler {
	router := mux.NewRouter()
	// Middleware
	router.Use(app.logRequest)
	// Handlers
	router.HandleFunc("/v0/hc", app.hc).Methods("GET")

	// users
	router.HandleFunc("/v0/users/register", app.registerUser).Methods("POST")

	router.Handle("/", NoSurf(app.Home)).Methods("GET")

	return router
}

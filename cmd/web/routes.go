package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()
	// Middleware
	router.Use(app.logRequest)
	// Handlers
	router.HandleFunc("/v0/hc", app.hc).Methods("GET")
	return router
}

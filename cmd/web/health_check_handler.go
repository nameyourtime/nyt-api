package main

import (
	"nameyourtime.com/api/pkg/models"
	"net/http"
	"time"
)

func (app *application) hc(w http.ResponseWriter, r *http.Request) {
	h := &models.Healthcheck{
		Status:     "OK",
		ServerTime: time.Now(),
	}
	reply(w, http.StatusOK, h)
}

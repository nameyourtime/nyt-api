package models

import (
	"time"
)

type Healthcheck struct {
	Status     string    `json:"status"`
	ServerTime time.Time `json:"server_time"`
}

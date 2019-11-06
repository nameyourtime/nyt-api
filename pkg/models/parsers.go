package models

import (
	"encoding/json"
	"net/http"
)

func ParseUser(r *http.Request) (*User, bool) {
	if r.Body == nil {
		return nil, false
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, false
	}
	return &user, true
}

package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestRegister_Positive(t *testing.T) {
	app := newTestApp(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	req := `{
            "name": "test", 
            "email": "test@test.com", 
            "password": "123456"
            }`
	status, _, resp := ts.post(t, "/v0/users/register", req)
	u, valid := parseUser(string(resp))
	if !valid {
		t.Error()
	}
	if status != http.StatusCreated {
		t.Errorf("status: want 201; got %d", status)
	}
	if u.Email != "test_user_email@example.com" {
		t.Errorf("email: want test_user_email@example.com; got %s", u.Email)
	}
	if u.Name != "test_user_name" {
		t.Errorf("name: want test_user_name; got %s", u.Name)
	}
	if u.Token.RefreshToken == "" {
		t.Errorf("refreshToken: empty")
	}
	if u.Token.AccessToken == "" {
		t.Errorf("accessToken: empty")
	}
	if u.Password != "" {
		t.Errorf("password: password should be empty; got %s", u.Password)
	}
}

func TestRegister_EmptyRequest(t *testing.T) {
	app := newTestApp(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	status, _, resp := ts.post(t, "/v0/users/register", "")

	if status != http.StatusBadRequest {
		t.Errorf("status: want 400; got %d", status)
	}
	wr := `{"code":"002","message":"can't read request body"}`
	if !bytes.Contains([]byte(wr), resp) {
		t.Errorf("response: want %s; got %s", wr, string(resp))
	}
}

func TestRegister_DuplicateEmail(t *testing.T) {
	app := newTestApp(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	req := `{
            "name": "test", 
            "email": "test_user_email@example.com", 
            "password": "123456"
            }`
	status, _, resp := ts.post(t, "/v0/users/register", req)

	if status != http.StatusBadRequest {
		t.Errorf("status: want 400; got %d", status)
	}
	wr := `{"code":"007","message":"user with this email already registered"}`
	if !bytes.Contains([]byte(wr), resp) {
		t.Errorf("response: want %s; got %s", wr, string(resp))
	}
}

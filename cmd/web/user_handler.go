package main

import (
	"github.com/dchest/uniuri"
	"nameyourtime.com/api/pkg/models"
	"net/http"
	"time"
)

func (app *Application) registerUser(w http.ResponseWriter, r *http.Request) {
	user, valid := models.ParseUser(r)

	if !valid {
		app.badRequest(w)
		return
	}

	if errs := user.Validate(true); errs.Present() {
		app.validationError(w, errs)
		return
	}

	existingUser, err := app.users.GetByEmail(user.Email)
	if err != nil && err != models.ErrNoRecord {
		app.serverError(w, err)
		return
	}

	if existingUser != nil {
		app.duplicatedEmail(w)
		return
	}
	pwdHash, err := hashAndSalt(user.Password)
	if err != nil {
		app.serverError(w, err)
	}
	user.Password = pwdHash
	user.Token.RefreshToken, user.Token.RefreshTokenExp = generateRefreshToken()
	user.Status = "PENDING"
	userId, err := app.users.Create(user)
	if err != nil {
		app.serverError(w, err)
	}
	user, err = app.users.Get(userId)
	if err != nil {
		app.serverError(w, err)
	}
	accessToken, err := generateAccessToken(userId)
	if err != nil {
		app.serverError(w, err)
	}
	user.Token.AccessToken = accessToken
	user.Password = ""

	go app.sendConfirmation(user.ID, user.Name, user.Email)
	reply(w, http.StatusCreated, user)
}

func (app *Application) sendConfirmation(userID, userName, email string) {
	code := generateVerificationCode(userID)
	c, err := app.verification.Create(code)
	if err != nil {
		app.errorLog.Println(err.Error())
		return
	}
	result, err := app.mailSender.SendConfirmation(email, userName, c)
	if err != nil {
		app.errorLog.Println(err)
		return
	}
	app.infoLog.Println(result)
}

func generateVerificationCode(userID string) models.VerificationCode {
	return models.VerificationCode{
		UserID:  userID,
		Code:    uniuri.NewLen(40),
		CodeExp: time.Now().Add(24 * time.Hour * 5),
	}
}

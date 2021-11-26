package main

import (
	"database/sql"
	"flag"
	"log"
	"nameyourtime.com/api/pkg/models"
	"nameyourtime.com/api/pkg/models/pg"
	"net/http"
	"os"
	"time"
)

type Application struct {
	infoLog  *log.Logger
	errorLog *log.Logger

	users interface {
		GetByEmail(string) (*models.User, error)
		Create(user *models.User) (string, error)
		Get(string) (*models.User, error)
	}

	verification interface {
		Create(code models.VerificationCode) (string, error)
	}

	mailSender interface {
		SendConfirmation(email, username, code string) (*SendMessageResponse, error)
	}
}

func main() {
	port := flag.String("port", ":5000", "Application port")
	key := flag.String("appkey", "test-key", "Application key")
	mailUrl := flag.String("mail_api_url", "", "URL to the email service")
	mailKey := flag.String("mail_api_key", "", "API key to the email service")
	//dsn := flag.String("dsn", "postgres://nameyourtime:nameyourtime@nameyourtime-test-db/nameyourtime?sslmode=disable", "Postgres data source")
	dsn := flag.String("dsn", "postgres://nameyourtime:nameyourtime@localhost/nameyourtime?sslmode=disable", "Postgres data source")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := connectDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &Application{
		users:        &pg.UserModel{DB: db},
		verification: &pg.VerificationModel{DB: db},
		mailSender: &MailSender{
			baseUrl: *mailUrl,
			apiKey:  *mailKey,
		},
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:         *port,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	initJwt(*key)
	infoLog.Printf("Starting server on %s port", *port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func connectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

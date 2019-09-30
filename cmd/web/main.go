package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	port := flag.String("port", ":5000", "Application port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
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
	infoLog.Printf("Starting server on %s port", *port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

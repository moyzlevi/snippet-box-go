package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	cfg := initConf()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	db, err := openDb(cfg.dsn)
	if err != nil {
		app.errorLog.Fatal(err)
	}
	defer db.Close()

	srv := &http.Server {
		Addr: cfg.addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting Server on %s...", cfg.addr)
	errorLog.Fatal(srv.ListenAndServe())
}

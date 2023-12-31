package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"snippetbox.moypietsch.com/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	snnipets  *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	cfg := initConf()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDb(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()

	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snnipets: &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server {
		Addr: cfg.addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting Server on %s...", cfg.addr)
	errorLog.Fatal(srv.ListenAndServe())
}

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

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.Handle("/custom-handler/", &cstmhandler{})
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server {
		Addr: cfg.addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Starting Server on %s...", cfg.addr)
	errorLog.Fatal(srv.ListenAndServe())
}

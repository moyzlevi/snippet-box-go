package main

import (
	"log"
	"net/http"
)



func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting Server on 4000...")
	log.Fatal(http.ListenAndServe(":4000", mux))
}

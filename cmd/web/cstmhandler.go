package main

import "net/http"

type cstmhandler struct {}

func (h *cstmhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my handler"))
}
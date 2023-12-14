package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"../../ui/html/base.tmpl.html",
		"../../ui/html/partials/nav.tmpl.html",
		"../../ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	// w.Write([]byte("Display a snippet view"))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "oh snail"
	content :="my snail my snail you are very trail"
	expire := 7

	id, err := app.snnipets.Insert(title, content, expire)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

	w.Write([]byte("Create a snippet "))
}
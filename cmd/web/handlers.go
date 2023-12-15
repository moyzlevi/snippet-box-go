package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"html/template"
	"snippetbox.moypietsch.com/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snnipets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// files := []string{
	// 	"../../ui/html/base.tmpl.html",
	// 	"../../ui/html/partials/nav.tmpl.html",
	// 	"../../ui/html/pages/home.tmpl.html",
	// }

	// ts, err := template.ParseFiles(files...)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// err = ts.ExecuteTemplate(w, "base", nil)

	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	app.infoLog.Printf("[HANDLERS] - Viewing snippet for id: %d\n", id)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snnipets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

		files := []string{
		"../../ui/html/base.tmpl.html",
		"../../ui/html/partials/nav.tmpl.html",
		"../../ui/html/pages/view.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", snippet)

	if err != nil {
		app.serverError(w, err)
	}
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
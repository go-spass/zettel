package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling home page request", "method", r.Method, "url", r.URL.Path)
	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message, use
	// the http.Error() function to send an Internal Server Error response to the
	// user, and then return from the handler so no subsequent code is executed.
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		logger.Error("could not parse home.tmpl.html", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Then we use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		logger.Error("could not execute template home.tmpl.html", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func zettelView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		logger.Error("invalid zettel view request", "method", r.Method, "url", r.URL.Path, "id", r.PathValue("id"))
		http.NotFound(w, r)
		return
	}
	logger.Info("handling zettel view request", "method", r.Method, "url", r.URL.Path, "id", id)
	_, _ = fmt.Fprintf(w, "Display zettel with id=%d...", id)
}

func zettelCreate(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling zettel create request", "method", r.Method, "url", r.URL.Path)
	_, _ = w.Write([]byte("Display a form for creating a new zettel..."))
}

func zettelCreatePost(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling zettel create post request", "method", r.Method, "url", r.URL.Path)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Save a new zettel to the DB..."))
}

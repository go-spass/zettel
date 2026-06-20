package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling home page request", "method", r.Method, "url", r.URL.Path)
	ts, err := template.ParseFiles("./ui/html/base.tmpl.html", "./ui/html/partials/nav.tmpl.html", "./ui/html/pages/home.tmpl.html")
	if err != nil {
		logger.Error("could not parse home template", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("could not execute home template", "error", err.Error())
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
	ts, err := template.ParseFiles("./ui/html/base.tmpl.html", "./ui/html/partials/nav.tmpl.html", "./ui/html/pages/zettel-view.tmpl.html")
	if err != nil {
		logger.Error("could not parse zettel view template", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("could not execute zettel view template", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func zettelCreate(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling zettel create request", "method", r.Method, "url", r.URL.Path)
	ts, err := template.ParseFiles("./ui/html/base.tmpl.html", "./ui/html/partials/nav.tmpl.html", "./ui/html/pages/zettel-create.tmpl.html")
	if err != nil {
		logger.Error("could not parse zettel create template", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		logger.Error("could not execute zettel create template", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func zettelCreatePost(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling zettel create post request", "method", r.Method, "url", r.URL.Path)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("Save a new zettel to the DB..."))
}

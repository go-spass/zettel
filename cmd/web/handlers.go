package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling home page request", "method", r.Method, "url", r.URL.Path)
	w.Header().Add("Server", "Go")
	_, _ = w.Write([]byte("Hello from Zettel!"))
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

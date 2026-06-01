package main

import (
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger

func main() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/zettel/view", zettelView)
	mux.HandleFunc("/zettel/create", zettelCreate)

	logger.Info("starting zettel server", "addr", ":4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		logger.Error("zettel server error", "err", err)
		os.Exit(1)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling home page request", "method", r.Method, "url", r.URL.Path)
	_, _ = w.Write([]byte("Hello from Zettel!"))
}

func zettelView(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling zettel view request", "method", r.Method, "url", r.URL.Path)
	_, _ = w.Write([]byte("Display a specific zettel..."))
}

func zettelCreate(w http.ResponseWriter, r *http.Request) {
	logger.Info("handling zettel create request", "method", r.Method, "url", r.URL.Path)
	_, _ = w.Write([]byte("Display a form for creating a new zettel..."))
}

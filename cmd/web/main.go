package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

var logger *slog.Logger

func main() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()

	// To prevent subtree path patterns from acting like they have a wildcard at
	// the end, you can append the special character sequence {$} to the end of
	// the pattern — like "/{$}" or "/static/{$}".
	//
	// So if you have the route pattern "/{$}", it effectively means match a single
	// slash, followed by nothing else. It will only match requests where the
	// URL path is exactly /.
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/zettel/view/{id}", zettelView)
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
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
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

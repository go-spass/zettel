package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger

func validateTemplates() error {
	pages := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/pages/zettel-view.tmpl.html",
		"./ui/html/pages/zettel-create.tmpl.html",
	}
	for _, page := range pages {
		if _, err := template.ParseFiles(
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := validateTemplates(); err != nil {
		logger.Error("template validation failed", "error", err)
		os.Exit(1)
	}

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
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /zettel/view/{id}", zettelView)
	mux.HandleFunc("GET /zettel/create", zettelCreate)
	mux.HandleFunc("POST /zettel/create", zettelCreatePost)

	logger.Info("starting zettel server", "addr", ":4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		logger.Error("zettel server error", "err", err)
		os.Exit(1)
	}
}

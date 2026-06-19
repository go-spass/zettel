package main

import (
	"bytes"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// go test runs from the package directory; navigate to project root so
	// template paths like "./ui/html/..." resolve correctly.
	if err := os.Chdir("../.."); err != nil {
		panic("could not chdir to project root: " + err.Error())
	}
	logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	os.Exit(m.Run())
}

func makeGetRequest(t *testing.T, handler http.Handler, path string) *http.Response {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr.Result()
}

func assertBodyContains(t *testing.T, body, want string) {
	t.Helper()
	if !strings.Contains(body, want) {
		t.Errorf("expected body to contain %q", want)
	}
}

// T002: TestHomeChrome — home page must include shared chrome elements.
func TestHomeChrome(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)

	resp := makeGetRequest(t, mux, "/")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	raw, _ := io.ReadAll(resp.Body)
	body := string(raw)

	assertBodyContains(t, body, "<header>")
	assertBodyContains(t, body, "<nav>")
	assertBodyContains(t, body, "<title>Home - Zettel</title>")
}

// T003: TestZettelViewChrome — zettel view page must include shared chrome elements.
func TestZettelViewChrome(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /zettel/view/{id}", zettelView)

	resp := makeGetRequest(t, mux, "/zettel/view/1")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	raw, _ := io.ReadAll(resp.Body)
	body := string(raw)

	assertBodyContains(t, body, "<header>")
	assertBodyContains(t, body, "<nav>")
	assertBodyContains(t, body, "<title>View Zettel - Zettel</title>")
}

// T004: TestZettelCreateChrome — zettel create page must include shared chrome elements.
func TestZettelCreateChrome(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /zettel/create", zettelCreate)

	resp := makeGetRequest(t, mux, "/zettel/create")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	raw, _ := io.ReadAll(resp.Body)
	body := string(raw)

	assertBodyContains(t, body, "<header>")
	assertBodyContains(t, body, "<nav>")
	assertBodyContains(t, body, "<title>Create Zettel - Zettel</title>")
}

// T014: TestNewPageInheritsChrome — a minimal template (title+main only) must render
// with full shared chrome from the base template, proving zero-duplication architecture.
func TestNewPageInheritsChrome(t *testing.T) {
	ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
	if err != nil {
		t.Fatalf("could not parse base template: %v", err)
	}
	ts, err = ts.Parse(`{{define "title"}}New Page{{end}}{{define "main"}}<p>New content</p>{{end}}`)
	if err != nil {
		t.Fatalf("could not parse page blocks: %v", err)
	}

	var buf bytes.Buffer
	if err := ts.ExecuteTemplate(&buf, "base", nil); err != nil {
		t.Fatalf("could not execute base template: %v", err)
	}

	body := buf.String()
	assertBodyContains(t, body, "<header>")
	assertBodyContains(t, body, "<nav>")
	assertBodyContains(t, body, "<title>New Page - Zettel</title>")
}

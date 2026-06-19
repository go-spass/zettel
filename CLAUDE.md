# CLAUDE.md — zettel

## Project Overview
`zettel` is a note-taking web application built in Go, modeled on the Snippet Box app from the [Let's Go](https://www.goodreads.com/book/show/43429043-let-s-go) book (v2.26.0). The book uses MySQL; **this project uses PostgreSQL instead.**

- **Module**: `github.com/go-spass/zettel`
- **Go version**: 1.26.3
- **Entry point**: `cmd/web/main.go`
- **Server port**: `:4000`

## Book Progress
Currently implemented through **Chapter 2.8** (HTML template composition). See `README.md` for detailed notes on each section.

## Build & Run
```bash
make build   # fmt → lint → vet → go build → ./build/zettel
make run     # build + run the binary
make test    # go test -v -cover ./...
```

The built binary is `./build/zettel`.

## Project Structure
```
cmd/web/
  main.go           # server setup, routing, logger declaration
  handlers.go       # HTTP handler functions
  handlers_test.go  # handler integration tests
ui/html/
  base.tmpl.html    # shared base template (chrome: head, header, nav, footer)
  pages/            # page-specific templates (*.tmpl.html)
build/              # compiled binary output (gitignored)
scripts/            # SQL scripts (e.g. seed data)
```

## Routes (current)
| Pattern          | Handler         |
|------------------|-----------------|
| `/{$}`           | `home`          |
| `/zettel/view`   | `zettelView`    |
| `/zettel/create` | `zettelCreate`  |

## Naming Convention
The book uses "snippet" throughout; this project uses **"zettel"** as the equivalent domain term. Apply this consistently when adding new handlers, routes, templates, and DB models.

## Database
- **PostgreSQL** (not MySQL as in the book)
- Adapt all SQL syntax and Go driver usage accordingly (`lib/pq` or `pgx`)
- DB name: `zettel` (equivalent to book's `snippetbox`)
- App DB user: `web`

## Logging
Uses `log/slog` with `slog.NewJSONHandler(os.Stdout, nil)` — structured JSON logs to stdout.

## Deviations from the Book
- PostgreSQL instead of MySQL — use `$1`, `$2` placeholders (not `?`)
- Domain entity is "zettel" not "snippet"

<!-- SPECKIT START -->
For additional context about technologies to be used, project structure,
shell commands, and other important information, read the current plan:
`specs/001-base-template-composition/plan.md`
<!-- SPECKIT END -->

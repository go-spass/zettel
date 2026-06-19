<!--
Sync Impact Report
==================
Version change: 1.0.0 → 1.1.0 (MINOR — new principle added)
Modified principles: None renamed
Added sections: VI. Test-Driven Development
Removed sections: None
Templates requiring updates:
  ✅ plan-template.md — Constitution Check section is generic; no updates required
  ✅ spec-template.md — No constitution-specific references; no updates required
  ✅ tasks-template.md — Updated: tests marked MANDATORY (were OPTIONAL), per Principle VI
Follow-up TODOs: None
-->

# Zettel Constitution

## Core Principles

### I. Book-Faithful Implementation

This project MUST follow the "Let's Go" book (v2.26.0) structure and patterns as its primary
reference. Where the book uses technology or naming that conflicts with this project's
deliberate deviations (see Principle II and III), the deviation takes precedence.
All other design decisions — middleware ordering, handler structure, template layout,
error handling patterns — MUST mirror the book unless a specific reason to deviate is
documented in CLAUDE.md.

### II. PostgreSQL-First

All database interactions MUST use PostgreSQL semantics:

- Positional placeholders `$1`, `$2`, … (not `?`)
- Driver: `lib/pq` or `pgx` (not `go-sql-driver/mysql`)
- SQL syntax compatible with PostgreSQL (e.g., `RETURNING`, sequences, not MySQL-specific constructs)
- DB name: `zettel`; application DB user: `web`

Any SQL lifted directly from the book MUST be adapted to PostgreSQL before use.

### III. Consistent Domain Naming

The domain entity MUST be **zettel / zettels** throughout the entire codebase:
routes, handler names, template filenames, HTML copy, database table/column names,
Go struct names, variable names, and test helpers.

The word "snippet" MUST NOT appear in any new or modified code, templates, or SQL.
Legacy occurrences inherited from the book MUST be renamed at the point of first
implementation.

### IV. Idiomatic Go

All code MUST conform to standard Go conventions:

- Formatted with `gofmt` (enforced via `make build`)
- Error handling MUST follow Go idioms: check-and-return, no `panic` in request paths
- No over-engineering: abstractions are added only when required by two or more concrete
  use cases already present in the code
- Entry point: `cmd/web/main.go`; HTTP handlers: `cmd/web/handlers.go`

### V. Structured Logging

All server-side event logging MUST use `log/slog` with a
`slog.NewJSONHandler(os.Stdout, nil)` handler. Structured JSON fields MUST be used
for all variable data (e.g., request IDs, error values, entity IDs). Plain `fmt.Println`
or the standard `log` package MUST NOT be used for application logging.

### VI. Test-Driven Development (NON-NEGOTIABLE)

All new features MUST follow the Red-Green-Refactor TDD cycle:

1. **Red**: Write a failing test that specifies the desired behaviour before writing
   any implementation code. The test MUST be verified to fail before proceeding.
2. **Green**: Write the minimum implementation code required to make the test pass.
3. **Refactor**: Clean up the implementation without breaking the passing test.

Tests are NOT optional. Every new handler, DB method, and model MUST have a
corresponding test written before its implementation is committed. `make test` MUST
pass in full before any feature branch is considered complete. Skipping the Red phase
(writing tests after implementation) is a constitution violation.

## Technology Stack

This section documents non-negotiable technology constraints. Deviating from these
requires amending this constitution.

- **Language**: Go 1.26 (`go version go1.26.x`)
- **HTTP**: Standard library `net/http` with Go 1.22+ enhanced ServeMux
- **Database**: PostgreSQL; driver `lib/pq` or `pgx`
- **Logging**: `log/slog` (structured JSON, stdout)
- **Templating**: `html/template` (standard library)
- **Build & Run**: `make build` / `make run` / `make test` (see `Makefile`)
- **Module path**: `github.com/go-spass/zettel`
- **Server port**: `:4000`

No third-party web frameworks (e.g., Gin, Echo) may be introduced. New dependencies
MUST be justified in the implementing PR.

## Development Workflow

- **Build gate**: `make build` runs `fmt → lint → vet → go build` in sequence.
  All steps MUST pass before a change is considered complete.
- **Tests**: `make test` runs `go test -v -cover ./...`. New handlers and DB methods
  MUST have corresponding tests before the feature branch is merged.
- **Commits**: Conventional commit prefixes (`feat:`, `fix:`, `refactor:`, `test:`,
  `docs:`) are REQUIRED. No commits with `--no-verify` unless explicitly approved.
- **Branch progress**: The project advances chapter by chapter through the book.
  Each chapter's work SHOULD be a distinct commit or PR. `CLAUDE.md` and `README.md`
  MUST be updated to reflect the current chapter milestone after each increment.

## Governance

This constitution supersedes all other runtime preferences and per-session instructions.
Amendments MUST:

1. Increment `CONSTITUTION_VERSION` according to semver (MAJOR for principle
   removals/redefinitions, MINOR for additions, PATCH for clarifications).
2. Update `LAST_AMENDED_DATE` to the date of the change.
3. Update `CLAUDE.md` ("Deviations from the Book" and any affected sections) to stay
   consistent with the new principle text.
4. Include a Sync Impact Report (HTML comment at top of this file) listing all changed
   principles and affected templates.

All feature plans generated by `/speckit-plan` MUST include a Constitution Check gate
that verifies compliance with Principles I–VI before Phase 0 research begins.

**Version**: 1.1.0 | **Ratified**: 2026-06-19 | **Last Amended**: 2026-06-19

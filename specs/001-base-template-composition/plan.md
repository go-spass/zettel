# Implementation Plan: Page Template Composition

**Branch**: `001-base-template-composition` | **Date**: 2026-06-19 | **Spec**: [spec.md](spec.md)

**Input**: Feature specification from `specs/001-base-template-composition/spec.md`

## Summary

Introduce a shared base HTML template (`ui/html/base.tmpl.html`) containing the common page chrome — `<head>` metadata, `<header>`, `<nav>`, and `<footer>`. All existing page handlers are updated to parse and compose the base template with their page-specific content block, eliminating duplicated markup and ensuring a single source of truth for site-wide layout.

## Technical Context

**Language/Version**: Go 1.26

**Primary Dependencies**: `html/template` (standard library) — no new external dependencies

**Storage**: N/A — no database interactions in this feature

**Testing**: `go test ./...` via `make test`; `net/http/httptest` for handler-level integration tests

**Target Platform**: Local HTTP server on `:4000`

**Project Type**: Web application (server-rendered HTML)

**Performance Goals**: No change from baseline; template parsing is per-request at this chapter stage (pre-caching)

**Constraints**: Standard library only; no third-party template engines (constitution Principle I + IV)

**Scale/Scope**: 3 existing pages migrated; ~5 files created or modified

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

| Principle | Status | Notes |
|-----------|--------|-------|
| I. Book-Faithful Implementation | ✅ PASS | Directly implements ch2.7–2.8 pattern using named `{{define}}`/`{{template}}` blocks |
| II. PostgreSQL-First | ✅ PASS | N/A — no database interactions |
| III. Consistent Domain Naming | ✅ PASS | Home page "Latest Snippets" → "Latest Zettels" renamed as part of this feature |
| IV. Idiomatic Go | ✅ PASS | Standard library only; `gofmt`; check-and-return error handling; no premature abstraction |
| V. Structured Logging | ✅ PASS | Existing `slog` handler logging preserved unchanged |
| VI. TDD (NON-NEGOTIABLE) | ✅ PASS | Tests written first (Red), verified failing, then implementation (Green); `httptest` assertions on chrome presence |

**Gate result**: All principles pass. Proceed.

## Project Structure

### Documentation (this feature)

```text
specs/001-base-template-composition/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output
├── quickstart.md        # Phase 1 output
├── contracts/
│   └── template-contract.md   # Phase 1 output
└── tasks.md             # Phase 2 output (/speckit-tasks — not yet created)
```

### Source Code Changes

```text
ui/html/
├── base.tmpl.html                    # NEW — shared base template (chrome)
└── pages/
    ├── home.tmpl.html                # MODIFIED — stripped to title+main blocks only
    ├── zettel-view.tmpl.html         # NEW — page template for zettel view
    └── zettel-create.tmpl.html       # NEW — page template for zettel create

cmd/web/
├── handlers.go                       # MODIFIED — parse base + page template in each handler
└── handlers_test.go                  # NEW — TDD tests for chrome presence on all pages
```

**Structure Decision**: Go web app layout (`cmd/web/`, `ui/html/`). No structural changes to the project layout — this feature only adds and modifies files within the existing layout.

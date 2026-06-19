---
description: "Task list for page template composition feature"
---

# Tasks: Page Template Composition

**Input**: Design documents from `specs/001-base-template-composition/`

**Prerequisites**: plan.md ✅ | spec.md ✅ | research.md ✅ | data-model.md ✅ | contracts/template-contract.md ✅

**Tests**: MANDATORY per Constitution Principle VI (TDD). Tests MUST be written before implementation and verified to FAIL first (Red phase). Each user story phase begins with test tasks.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story?] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1, US2)
- Include exact file paths in all descriptions

## Path Conventions

Source files follow the existing project layout:

```text
cmd/web/           # Handler logic and tests
ui/html/           # Base template (root level)
ui/html/pages/     # Page-specific templates
specs/001-base-template-composition/  # Feature docs (read-only reference)
```

---

## Phase 1: Setup

**Purpose**: Create the test file that US1 and US2 tests will be written into.

- [ ] T001 Create `cmd/web/handlers_test.go` with `package main` declaration and a `makeGetRequest(t *testing.T, mux http.Handler, path string) *http.Response` helper that uses `httptest.NewRecorder` and returns the recorded response

**Checkpoint**: `cmd/web/handlers_test.go` exists and compiles with `make build`

---

## Phase 2: Foundational

**Purpose**: No shared blocking infrastructure is needed beyond Phase 1. Proceed to user stories.

**Checkpoint**: Phase 1 complete — begin user story phases.

---

## Phase 3: User Story 1 — Consistent Page Chrome on Every Page (Priority: P1) 🎯 MVP

**Goal**: All three existing pages (home, zettel view, zettel create) render an identical `<header>`, `<nav>`, and `<head>` metadata section, defined in a single base template.

**Independent Test**: Load home, zettel view, and zettel create in sequence. Every page must contain `<header>`, `<nav>`, and a `<title>` tag matching the pattern `[page] - Zettel`. Changing the footer text in `base.tmpl.html` alone must update all three pages on next load.

### Tests for User Story 1 (MANDATORY — write first, verify they FAIL) ⚠️

> **Per Constitution Principle VI (TDD): write these tests FIRST and confirm they FAIL before creating any template files.**

- [ ] T002 [US1] Add `TestHomeChrome` to `cmd/web/handlers_test.go`: call `makeGetRequest` on `GET /`, assert status 200, body contains `<header>`, `<nav>`, and `<title>Home - Zettel</title>`
- [ ] T003 [US1] Add `TestZettelViewChrome` to `cmd/web/handlers_test.go`: call `makeGetRequest` on `GET /zettel/view/1`, assert status 200, body contains `<header>`, `<nav>`, and `<title>View Zettel - Zettel</title>`
- [ ] T004 [US1] Add `TestZettelCreateChrome` to `cmd/web/handlers_test.go`: call `makeGetRequest` on `GET /zettel/create`, assert status 200, body contains `<header>`, `<nav>`, and `<title>Create Zettel - Zettel</title>`
- [ ] T005 [US1] Run `make test` — verify T002, T003, T004 all FAIL (expected: template parse errors or assertion failures). Do NOT proceed to implementation until failures are confirmed.

### Implementation for User Story 1

- [ ] T006 [US1] Create `ui/html/base.tmpl.html`: define `{{define "base"}}` block containing the full HTML skeleton — `<!doctype html>`, `<head>` with charset and `<title>{{template "title" .}} - Zettel</title>`, `<header>` with site name link, `<nav>` with links to `/`, `/zettel/create`, `<main>{{template "main" .}}</main>`, `<footer>` with Go attribution
- [ ] T007 [US1] Rewrite `ui/html/pages/home.tmpl.html`: replace the full HTML document with `{{template "base" .}}`, `{{define "title"}}Home{{end}}`, and `{{define "main"}}` block containing the latest zettels section (rename "Latest Snippets" → "Latest Zettels" per Constitution Principle III)
- [ ] T008 [P] [US1] Create `ui/html/pages/zettel-view.tmpl.html`: `{{template "base" .}}` + `{{define "title"}}View Zettel{{end}}` + `{{define "main"}}` block with a placeholder view content area
- [ ] T009 [P] [US1] Create `ui/html/pages/zettel-create.tmpl.html`: `{{template "base" .}}` + `{{define "title"}}Create Zettel{{end}}` + `{{define "main"}}` block with a placeholder create form content area
- [ ] T010 [US1] Update `home` handler in `cmd/web/handlers.go`: change `template.ParseFiles` call to `template.ParseFiles("./ui/html/base.tmpl.html", "./ui/html/pages/home.tmpl.html")`
- [ ] T011 [US1] Update `zettelView` handler in `cmd/web/handlers.go`: change response to use `template.ParseFiles("./ui/html/base.tmpl.html", "./ui/html/pages/zettel-view.tmpl.html")` and `ts.Execute` instead of `fmt.Fprintf`
- [ ] T012 [US1] Update `zettelCreate` handler in `cmd/web/handlers.go`: change response to use `template.ParseFiles("./ui/html/base.tmpl.html", "./ui/html/pages/zettel-create.tmpl.html")` and `ts.Execute` instead of `w.Write`
- [ ] T013 [US1] Run `make test` — verify T002, T003, T004 now PASS (Green phase). All other existing tests must also pass.

**Checkpoint**: User Story 1 fully functional. All pages share a single source of chrome. Run `make run` and manually verify all three pages look consistent.

---

## Phase 4: User Story 2 — New Page Inherits Shared Layout Automatically (Priority: P2)

**Goal**: Demonstrate and verify that adding a new page requires authoring only `title` and `main` blocks — zero shared chrome markup.

**Independent Test**: A test creates a page template containing only `{{define "title"}}` and `{{define "main"}}` blocks (no `<html>`, no `<header>`, no `<nav>`), parses it together with the base template, renders it, and asserts the output contains `<header>` and `<nav>`. This proves the zero-duplication property of the architecture.

### Tests for User Story 2 (MANDATORY — write first, verify they FAIL) ⚠️

> **Per Constitution Principle VI (TDD): write T014 BEFORE Phase 3 implementation is complete if possible; otherwise write it after T013 and verify it fails before T015.**

- [ ] T014 [US2] Add `TestNewPageInheritsChrome` to `cmd/web/handlers_test.go`: parse `"./ui/html/base.tmpl.html"` into a template set, then call `.Parse("{{define \"title\"}}New Page{{end}}{{define \"main\"}}<p>New content</p>{{end}}")` on it, execute the `"base"` template to a `bytes.Buffer`, assert output contains `<header>` and `<nav>` and `<title>New Page - Zettel</title>`

### Implementation for User Story 2

- [ ] T015 [US2] Run `make test` — verify T014 passes. No new source files are required; the base template architecture already satisfies US2. If T014 fails, diagnose the template execution order and fix `base.tmpl.html` accordingly.

**Checkpoint**: User Stories 1 and 2 both independently verified. The architecture provably supports zero-chrome-duplication for any future page.

---

## Phase 5: Polish & Cross-Cutting Concerns

**Purpose**: Build gate, smoke test, and regression verification.

- [ ] T016 Run `make build` — confirm `fmt → lint → vet → go build` all pass with zero errors or warnings
- [ ] T017 Manual smoke test per `specs/001-base-template-composition/quickstart.md` — run `make run`, load all three pages, perform the change-propagation check (edit footer in `base.tmpl.html`, restart, verify all three pages reflect the change)

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies — start immediately
- **User Story 1 (Phase 3)**: Depends on Phase 1 (handlers_test.go must exist)
- **User Story 2 (Phase 4)**: Depends on T006 (base.tmpl.html must exist for T014 to be meaningful)
- **Polish (Phase 5)**: Depends on Phase 3 and Phase 4 completion

### Within User Story 1

```
T002 → T003 → T004 → T005 (Red gate)
       ↓
      T006 → T007
              T008 [P] ─┐
              T009 [P] ─┘→ T010 → T011 → T012 → T013 (Green gate)
```

### Within User Story 2

```
T014 → T015
```

T008 and T009 can run in parallel (different new files).
T002–T004 write to the same file and must run sequentially.
T010–T012 write to the same file and must run sequentially.

---

## Parallel Opportunities

### User Story 1 — template file creation

```bash
# After T007, these two can run simultaneously:
Task: "Create ui/html/pages/zettel-view.tmpl.html" (T008)
Task: "Create ui/html/pages/zettel-create.tmpl.html" (T009)
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (T001)
2. Complete Phase 3 Tests: Write T002–T004, confirm Red (T005)
3. Complete Phase 3 Implementation: T006–T012
4. **STOP and VALIDATE**: Run `make test` (T013) — all chrome tests pass
5. Run `make run` and visually confirm all pages share the base chrome

### Incremental Delivery

1. Phase 1 → Phase 3 Tests (Red) → Phase 3 Implementation (Green) → US1 validated ✅
2. Phase 4 Tests + verify → US2 validated ✅
3. Phase 5 → Full build + smoke test → Feature complete ✅

---

## Notes

- `[P]` tasks operate on different files and have no incomplete dependencies
- TDD gates (T005, T013, T015) are NOT optional — skipping them violates Constitution Principle VI
- "Latest Snippets" → "Latest Zettels" is required in T007 per Constitution Principle III
- The `zettelCreatePost` handler does NOT need a template (it responds with status + text); leave it unchanged
- Template files use `.tmpl.html` extension matching the existing project convention

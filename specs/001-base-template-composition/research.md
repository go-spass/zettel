# Research: Page Template Composition

## Decision 1: Template Composition Mechanism

**Decision**: Use Go's `html/template` named template system with `{{define}}` and `{{template}}` actions.

**Rationale**: The standard library `html/template` package supports named template blocks natively. The book (ch2.7–2.8) uses exactly this pattern. A base template defines the page skeleton and declares named slots (`{{template "title" .}}`, `{{template "main" .}}`); each page template fills those slots with `{{define "title"}}` and `{{define "main"}}` blocks.

**Alternatives considered**:
- Third-party template engines (templ, Jet): rejected — constitution Principle I requires book-faithful implementation and Principle IV prohibits unnecessary dependencies.
- Server-side includes: not applicable in Go's template system.

---

## Decision 2: Template File Layout

**Decision**: Single `base.tmpl.html` at `ui/html/` root; page-specific templates remain in `ui/html/pages/`.

**Rationale**: Mirrors the "Let's Go" book directory convention. Keeps the base template visually distinct from per-page templates.

**Alternatives considered**:
- Putting base template in `ui/html/partials/`: overly nested for a single file at this stage; the book uses the root.

---

## Decision 3: Template Parsing in Handlers

**Decision**: Each handler calls `template.ParseFiles()` with both the base template and the relevant page template. Files are parsed in the order: base first, page second.

**Rationale**: Keeps parsing explicit and co-located with the handler that uses it. The book follows this pattern at this chapter. Extraction into a shared render helper is a future refactor (ch3.x), deferred per constitution Principle IV (no premature abstraction).

**Alternatives considered**:
- Pre-parse all templates at startup: correct long-term approach (book introduces this later); deferred to the appropriate chapter.
- `template.ParseGlob()`: would parse all templates in a directory, mixing unrelated page templates; not appropriate yet.

---

## Decision 4: "Latest Snippets" Rename

**Decision**: The home page copy "Latest Snippets" MUST be renamed to "Latest Zettels" as part of this feature.

**Rationale**: Constitution Principle III: "snippet" MUST NOT appear in modified code or templates. The home template is being modified during this feature, making this the point of first modification.

---

## Decision 5: Test Approach

**Decision**: Use `net/http/httptest` with a test HTTP server to assert that rendered pages contain the expected shared chrome elements.

**Rationale**: Constitution Principle VI mandates TDD. Handler-level integration tests using `httptest.NewRecorder()` and `http.Request` are the idiomatic Go approach for testing HTTP handlers without mocking the template engine.

Tests must be written and verified to FAIL before the base template is created.

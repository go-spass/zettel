# Data Model: Page Template Composition

This feature introduces no new data entities or database schema changes. The "model" here is the **template composition contract** — the structural relationship between the base template and page templates.

## Template Entities

### Base Template (`ui/html/base.tmpl.html`)

The single source of truth for all shared page chrome.

| Named Block | Required? | Purpose |
|-------------|-----------|---------|
| `base`      | N/A (defines the root) | The top-level named template that wraps the full HTML document |
| `title`     | MUST be defined by each page template | Page-specific title text inserted into `<title>[title] - Zettel</title>` |
| `main`      | MUST be defined by each page template | Page-specific body content rendered inside `<main>` |

**Invariants**:
- `<head>` with `charset`, `viewport`, and `<title>` is always present
- `<header>` with site name link (`/`) is always present
- `<nav>` with links to all current routes is always present
- `<footer>` is always present
- Content is rendered in `<main>` and is the only area that varies per page

### Page Template (one per page, `ui/html/pages/*.tmpl.html`)

Each page template composes with the base by invoking `{{template "base" .}}` and defining the required named blocks.

| Page | File | `title` block | `main` block |
|------|------|---------------|--------------|
| Home | `home.tmpl.html` | `"Home"` | Latest zettels listing |
| Zettel View | `zettel-view.tmpl.html` | `"View Zettel"` | Single zettel display |
| Zettel Create | `zettel-create.tmpl.html` | `"Create Zettel"` | Zettel creation form |

## State Transitions

No state transitions — template composition is stateless at render time. Dynamic data (zettel content, IDs) is passed to templates via the handler's data argument (currently `nil`; typed data structs are a future chapter concern).

## Validation Rules

- A page template that does not define `{{define "main"}}` will cause a template execution error at runtime. Tests must cover this contract.
- A page template that does not define `{{define "title"}}` will render a blank title. Tests must assert the title is non-empty.

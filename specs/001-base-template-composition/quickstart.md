# Quickstart Validation Guide: Page Template Composition

## Prerequisites

- Go 1.26 installed
- Repository cloned and `make build` passing on `main`
- Current branch: `001-base-template-composition`

## Validation Steps

### 1. Red Phase — Confirm Tests Fail First

Before implementing the base template, write the handler tests (see tasks.md T001–T003) and confirm they fail:

```bash
make test
```

**Expected**: Tests fail because `base.tmpl.html` does not yet exist and the handlers do not parse it. You should see template-not-found or assertion errors.

### 2. Green Phase — Implement Base Template

After confirming the Red phase, implement the base template and update the handlers per the template contract. Then re-run:

```bash
make test
```

**Expected**: All new tests pass. No existing tests regress.

### 3. Build Gate

```bash
make build
```

**Expected**: `fmt → lint → vet → go build` all pass with zero errors.

### 4. Manual Smoke Test

```bash
make run
```

Open each page in a browser and verify:

| Page | URL | Expected Chrome | Expected Content |
|------|-----|----------------|-----------------|
| Home | `http://localhost:4000/` | Header with "Zettel" link, nav with all routes, footer | "Latest Zettels" heading |
| Zettel View | `http://localhost:4000/zettel/view/1` | Same header + nav + footer | View zettel content area |
| Zettel Create | `http://localhost:4000/zettel/create` | Same header + nav + footer | Create zettel form area |

### 5. Change-Propagation Check

1. Edit `ui/html/base.tmpl.html` — change the footer text.
2. Restart the server (`make run`).
3. Load all three pages.

**Expected**: All three pages show the updated footer text. No per-page file was changed.

### 6. Regression Check

```bash
make test
```

**Expected**: All tests pass. `make build` still clean.

## Success Criteria Mapping

| Success Criterion | Validation Step |
|-------------------|----------------|
| SC-001: Identical chrome on every page | Step 4 — manual side-by-side |
| SC-002: Zero shared markup in page templates | Code review — each page template defines only `title` and `main` blocks |
| SC-003: Single edit propagates to all pages | Step 5 — change-propagation check |
| SC-004: No regression in existing pages | Step 6 — full test run |

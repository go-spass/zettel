# Feature Specification: HTML Template Nav Partial

**Feature Branch**: `002-support-embedded-partials`

**Created**: 2026-06-20

**Status**: Draft

**Input**: User description: "Add support for embedded partials in the HTML templates. Specifically, introduce a new 'nav' partial that will contain all of the site navigation and embed it into the site pages."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Consistent Navigation Across All Pages (Priority: P1)

A site visitor navigates between the home page, zettel view, and zettel create pages and sees the same navigation on every page. The navigation is always present and always up to date, regardless of which page is visited.

**Why this priority**: This is the core visible outcome of the feature — navigation consistency is the end-user benefit that the partial exists to provide.

**Independent Test**: Can be fully tested by visiting each route (`/`, `/zettel/view`, `/zettel/create`) and confirming navigation links are rendered identically on all of them.

**Acceptance Scenarios**:

1. **Given** the home page is loaded, **When** the page is rendered, **Then** the navigation partial is present and contains all site navigation links
2. **Given** a zettel view page is loaded, **When** the page is rendered, **Then** the navigation partial is present and identical to the home page navigation
3. **Given** the zettel create page is loaded, **When** the page is rendered, **Then** the navigation partial is present and identical to the other pages

---

### User Story 2 - Navigation Managed in a Single Location (Priority: P2)

A developer updating the site navigation edits a single nav partial file and the change is reflected on all pages without touching individual page templates or the base template.

**Why this priority**: This is the maintainability benefit that motivates the partial pattern; it depends on Story 1 being complete first.

**Independent Test**: Can be tested by modifying the nav partial's content and confirming the change appears on every rendered page without any other file being edited.

**Acceptance Scenarios**:

1. **Given** the nav partial file exists, **When** a navigation link is added or changed in that file, **Then** the updated navigation appears on all site pages
2. **Given** the nav partial file exists, **When** the base template and page templates are inspected, **Then** no navigation markup appears outside the nav partial

---

### Edge Cases

- What happens when the nav partial template file is missing or has a parse error at server start?
- How does the system handle a page template that does not embed the nav partial?

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: A dedicated nav partial template file MUST exist at a path separate from the base template and page templates
- **FR-002**: The nav partial MUST define the template block that renders all site navigation links, with content identical to what currently exists in the base template — no links added, removed, or renamed
- **FR-003**: The base template MUST embed the nav partial so it is included in every page rendered through it
- **FR-004**: All existing page templates (home, zettel view, zettel create) MUST render with the nav partial present
- **FR-005**: No navigation markup MUST exist outside the nav partial file
- **FR-006**: The application MUST fail at startup (not at request time) if the nav partial file cannot be parsed

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: All three existing page routes (`/`, `/zettel/view`, `/zettel/create`) render with identical navigation, confirmed by inspection or automated test
- **SC-002**: Navigation content exists in exactly one file; a grep for navigation link text finds matches only in the nav partial
- **SC-003**: A single edit to the nav partial is sufficient to change navigation on all pages — no other template file requires modification
- **SC-004**: All existing automated tests pass after being updated to include the nav partial file in their template parsing sets; no test that verifies navigation chrome may be removed or skipped

## Clarifications

### Session 2026-06-20

- Q: SC-004 claimed existing tests pass "without modification" — T014 parses only the base template and would break when nav is extracted. How should SC-004 be revised? → A: SC-004 revised to require tests pass after being updated to include the nav partial in their template parsing sets; no nav-chrome tests may be removed or skipped.
- Q: Should the nav partial extract the current navigation exactly, or is this an opportunity to revise navigation links? → A: Extract exactly as-is — same two links (Home, Create zettel), same structure; no content changes in scope.

## Assumptions

- Navigation content currently exists in the base template (`ui/html/base.tmpl.html`) and will be extracted into the new partial
- The nav partial will follow the same Go `html/template` define/block mechanism already used for page templates in this project
- The partial file will reside under `ui/html/partials/` consistent with the directory structure convention used for `ui/html/pages/`
- Scope is limited to the three routes currently implemented (`/`, `/zettel/view`, `/zettel/create`)
- No new routes or navigation destinations are introduced as part of this feature

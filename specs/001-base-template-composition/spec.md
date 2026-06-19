# Feature Specification: Page Template Composition

**Feature Branch**: `001-base-template-composition`

**Created**: 2026-06-19

**Status**: Draft

**Input**: User description: "As we add more pages to our web application, there will be some shared, boilerplate, HTML markup that we want to include on every page — like the header, navigation and metadata inside the `<head>` HTML element. To prevent duplication and save typing, it's a good idea to create a base (or master) template which contains this shared content, which we can then compose with the page-specific markup for the individual pages."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Consistent Page Chrome on Every Page (Priority: P1)

A visitor navigating between any two pages of the application sees an identical header, navigation, and page metadata — there are no missing nav links, stray titles, or mismatched branding caused by copy-paste errors or omissions.

**Why this priority**: This is the core value of template composition. Inconsistent chrome erodes trust and creates maintenance debt that compounds with every new page added.

**Independent Test**: Open the home page and the zettel view page side by side. Both must display the same header and navigation structure. The `<head>` metadata (charset, viewport, title pattern) must follow the same format on both pages.

**Acceptance Scenarios**:

1. **Given** a visitor loads the home page, **When** they view the page, **Then** the page includes a consistent header, navigation, and HTML head metadata identical in structure to all other pages.
2. **Given** a visitor loads the zettel view page, **When** they view the page, **Then** the header and navigation match the home page exactly, with only the page-specific content area differing.
3. **Given** a navigation link is updated in the shared layout, **When** any page is loaded, **Then** the updated link appears on all pages without any per-page edits.

---

### User Story 2 - New Page Inherits Shared Layout Automatically (Priority: P2)

A developer creating a new page only writes the content unique to that page. The shared chrome (header, nav, metadata) is automatically included by the composition system — no copying or pasting of boilerplate.

**Why this priority**: Without this, every new page introduced in future chapters requires manual duplication of the shared markup, increasing the risk of drift and inconsistency.

**Independent Test**: Add a new page template containing only its unique content block. Serve the page and confirm that the header, navigation, and head metadata appear without having been explicitly written in the new page template.

**Acceptance Scenarios**:

1. **Given** a new page template is created with only page-specific content, **When** the page is served, **Then** the full rendered output includes the shared header, nav, and head metadata.
2. **Given** the shared base template is the single source of truth for chrome, **When** the base template is modified, **Then** all pages immediately reflect the change on next load.

---

### Edge Cases

- What happens if a page template omits the content block? The page should still render with the shared chrome but with an empty content area (no error).
- What happens when the application serves a 404 or error page? Error pages should also use the shared layout for consistency.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Every page MUST include a consistent HTML head section (charset declaration, viewport settings, and a page title following a common pattern).
- **FR-002**: Every page MUST include a shared header and navigation section rendered identically across all pages.
- **FR-003**: Each page MUST render its own unique content within a designated content area, separate from the shared chrome.
- **FR-004**: The shared layout MUST be defined in a single location so that a change to it propagates to all pages without per-page edits.
- **FR-005**: Adding a new page to the application MUST NOT require duplicating any shared markup — only page-specific content needs to be authored.

### Key Entities

- **Base Template**: The single shared layout containing the HTML skeleton, head metadata, header, and navigation. All pages are composed using this template.
- **Page Template**: A page-specific content block that slots into the base template's designated content area. Each page (home, zettel view, zettel create, etc.) has its own page template.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Every page in the application displays an identical header and navigation structure with zero visual or structural inconsistencies across pages.
- **SC-002**: Adding a new page requires authoring zero lines of shared chrome markup — only page-specific content is written.
- **SC-003**: A single edit to the shared layout is reflected on all pages; no secondary edits to individual page files are needed.
- **SC-004**: All existing pages (home, zettel view, zettel create) continue to render correctly after the shared layout is introduced, with no regression in displayed content.

## Assumptions

- All pages in the current application (home, zettel view, zettel create) will be migrated to use the shared layout as part of this feature.
- The navigation links included in the shared layout reflect the current route structure (`/`, `/zettel/view`, `/zettel/create`).
- Error and 404 pages are out of scope for this feature but are expected to adopt the shared layout in a future iteration.
- The shared layout does not include user authentication state or dynamic per-user content (the application does not have auth yet).

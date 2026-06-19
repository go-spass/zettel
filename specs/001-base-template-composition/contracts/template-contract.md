# Template Composition Contract

## Overview

This contract defines the interface between the base template and page templates. Any page template that fails to satisfy this contract will produce a runtime error or visually broken page.

## Base Template Interface

The base template (`ui/html/base.tmpl.html`) exposes two named block slots that page templates MUST fill:

```
{{define "base"}}
  ...
  <title>{{template "title" .}} - Zettel</title>
  ...
  <main>
    {{template "main" .}}
  </main>
  ...
{{end}}
```

## Page Template Requirements

Every page template MUST:

1. Invoke the base template as the first line:
   ```
   {{template "base" .}}
   ```

2. Define a `title` block returning a short, non-empty string:
   ```
   {{define "title"}}Home{{end}}
   ```

3. Define a `main` block containing the page-specific HTML content:
   ```
   {{define "main"}}
   <h2>Page Heading</h2>
   <p>Page content here.</p>
   {{end}}
   ```

## Parsing Contract

Handlers MUST parse the base template and the page template together in a single `template.ParseFiles()` call, with the base template listed first:

```
template.ParseFiles(
    "./ui/html/base.tmpl.html",
    "./ui/html/pages/<page>.tmpl.html",
)
```

The resulting template set is then executed by invoking the `"base"` named template (this happens implicitly when the first file defines `{{define "base"}}`).

## Violation Behaviour

| Violation | Result |
|-----------|--------|
| Page template missing `{{define "main"}}` | Runtime template execution error → 500 response |
| Page template missing `{{define "title"}}` | `<title> - Zettel</title>` (empty title) — not a crash but fails SC-001 |
| Base template not included in `ParseFiles` | Template not found error → 500 response |
| Base template listed after page template in `ParseFiles` | May cause incorrect template execution order |

## Test Assertions (per constitution Principle VI)

Tests must assert, for each page:
- Response status is 200
- Response body contains `<header>`
- Response body contains `<nav>`
- Response body contains a non-empty `<title>` tag matching `<title>[page title] - Zettel</title>`
- Response body contains page-specific content unique to that page

# zettel
`zettel` is modeled on the Snippet Box application from the [Let's Go](https://www.goodreads.com/book/show/43429043-let-s-go)
Version Version 2.26.0 of the book, last updated 2026-03-07.
The code is based on the latest version of the book which uses Go version 1.26.

## Running zettel
To run the `zettel` application, you need to have Go installed on your
machine. You can download Go from the [official website](https://golang.org/).

### Build the Project
To build the project, navigate to the root of the project and run the
following:
```bash
make web
```

### Run the Project
To run the project, navigate to the root of the project and run the
following:
```bash
./build/web
```

# Project Notes
## Chapter 2
### Section 2.3
This section covered routing requests to different handlers based on the
path of the request.

|   Route Pattern |         Handler | Action                                |
|----------------:|----------------:|:--------------------------------------|
|               / |          home() | Display the home page                 |
|   /snippet/view |   snippetView() | Display a specific snippet            |
| /snippet/create | snippetCreate() | Display a form for creating a snippet |

#### Trailing slashes in route patterns
It’s important to know that Go’s `servemux` has different matching rules
depending on whether a route pattern ends with a trailing slash or not.

**exact path**:
- If the path does not have a trailing slash (e.g. `/snippet/create`), the
  pattern matches that path exactly.

**sub-tree path**:
- When a route ends with a trailing slash (e.g. "/" or "/snippet/"), the
  pattern matches that path and any path that has that prefix.
- To prevent subtree path patterns from acting like they have a wildcard
  at the end, you can append the special character sequence `{$}` to the
  end of the pattern  — like `/{$}` or `/static/{$}`.

### Section 2.4 Wildcard Segments in Route Patterns
This section covers how to use wildcard patterns to match dynamic parts
of a URL path. This essentially allows you to capture parts of the URL
that define query parameters or other dynamic parts of the URL.

Wildcard segments in a route pattern are denoted by an wildcard identifier
inside `{}` brackets. Like this:

```go
mux.HandleFunc("/products/{category}/item/{itemID}", exampleHandler)
```

Inside your handler, you can retrieve the corresponding value for a
wildcard segment using its identifier and the `r.PathValue()` method.
For example:
```go
func exampleHandler(w http.ResponseWriter, r *http.Request) {
    category := r.PathValue("category")
    itemID := r.PathValue("itemID")

    ...
}
```

For this section the routes will be updated to include wildcard segments
as follows:

|      Route Pattern |         Handler | Action                                |
|-------------------:|----------------:|:--------------------------------------|
|               /{$} |          home() | Display the home page                 |
| /snippet/view/{id} |   snippetView() | Display a specific snippet            |
|    /snippet/create | snippetCreate() | Display a form for creating a snippet |

### Section 2.5 Specify HTTP Methods on Routes
This section covers how to specify the HTTP methods that a route should
match. This is useful when you want to restrict a route to only respond
to a specific HTTP method (e.g. POST or GET).

To restrict a route to a specific HTTP method, you can prefix the route
pattern with the necessary HTTP method when declaring it, like so:

```golang
mux.HandleFunc("GET /{$}", home)
```

In this section the routes will updated as follows:


|          Route Pattern |              Handler | Action                                |
|-----------------------:|---------------------:|:--------------------------------------|
|               GET /{$} |               home() | Display the home page                 |
| GET /snippet/view/{id} |        snippetView() | Display a specific snippet            |
|    GET /snippet/create |      snippetCreate() | Display a form for creating a snippet |
|   POST /snippet/create | snippetCreatePost()  | Save a new snippet to the database    |

### Section 2.6 Customizing HTTP Responses
In Go, by default, every response that your handlers send has the HTTP
status code 200 OK (which indicates to the user that their request was
received and processed successfully), plus three automatic system-generated headers:
- a Date header,
- the Content-Length header, and
- the Content-Type of the response body

In this section the code will be updated to send appropriate HTTP status codes.

#### Update createSnippetPost to return a 201 Created status code
Use the `writeHeader()` method on the `http.ResponseWriter` to set the
status code to 201 Created.

```golang
// Use the w.WriteHeader() method to send a 201 status code.
w.WriteHeader(201)
```

#### Updating Headers Using Header().Add()
You can use the `Header().Add()` method to add custom headers to the
response. For example, to add a `Location` header to the response, you
can do the following:

```golang
w.Header().Add("Location", "/snippet/create")
```

### Section 2.9 Service Static Files
In section 2.9 we will add support for serving static files such as CSS,
so that we can style our application. To do this, we will use the
`http.FileServer` and `http.StripPrefix` functions from the `net/http`.

Go’s net/http package ships with a built-in http.FileServer handler which
you can use to serve files over HTTP from a specific directory. Let’s add
a new route to our application so that all GET requests which begin with
"/static/" are handled using this.


|          Route Pattern |              Handler | Action                                |
|-----------------------:|---------------------:|:--------------------------------------|
|               GET /{$} |               home() | Display the home page                 |
| GET /snippet/view/{id} |        snippetView() | Display a specific snippet            |
|    GET /snippet/create |      snippetCreate() | Display a form for creating a snippet |
|   POST /snippet/create |  snippetCreatePost() | Save a new snippet to the database    |
|           GET /static/ |    http.FileServer() | Serve a specific static file          |

> Remember: The pattern "GET /static/" is a subtree path pattern, so it acts a bit like
> there is a wildcard at the end.

## Section 4 Databases
### Getting Started
**Connecting to the Database via the Terminal**
```sh
# Connect as Root
mysql -u root -p
# Connect as web application user
mysql -D snippetbox -u web -p
```

**Loading Sample Database Records**
```sh
 mysql -D snippetbox -u web -p < ./scripts/insert_snippets.sql
```

### 4.8 Multiple Record SQL Queries
In this section the book covered the patern for executing SQL statements
that return multiple rows.  I’ll demonstrate this by updating the
`SnippetModel.Latest()` method to return the ten most-recently created
snippets (so long as they haven’t expired) using the following SQL query:

```sql
SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10
```

**Using the Query in the Latest() method**
```golang
func (m *SnippetModel) Latest() ([]Snippet, error) {
    // Write the SQL statement we want to execute.
    stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

    // Use the Query() method on the connection pool to execute our
    // SQL statement. This returns a sql.Rows resultset containing the result of
    // our query.
    rows, err := m.DB.Query(stmt)
    if err != nil {
        return nil, err
    }

    // We defer rows.Close() to ensure the sql.Rows resultset is
    // always properly closed before the Latest() method returns. This defer
    // statement should come *after* you check for an error from the Query()
    // method. Otherwise, if Query() returns an error, you'll get a panic
    // trying to close a nil resultset.
    defer rows.Close()

    // Initialize an empty slice to hold the Snippet structs.
    var snippets []Snippet

    // Use rows.Next to iterate through the rows in the resultset. This
    // prepares the first (and then each subsequent) row to be acted on by the
    // rows.Scan() method. If iteration over all the rows completes then the
    // resultset automatically closes itself and frees up the underlying
    // database connection.
    for rows.Next() {
        // Create a new zero value Snippet struct.
        var s Snippet
        // Use rows.Scan() to copy the values from each field in the row[…]
        // to the
        // new Snippet struct that we created. Again, the arguments to row.Scan()
        // must be pointers to the place you want to copy the data into, and the
        // number of arguments must be exactly the same as the number of
        // columns returned by your statement.
        err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
        if err != nil {
            return nil, err
        }
        // Append it to the slice of snippets.
        snippets = append(snippets, s)
    }

    // When the rows.Next() loop has finished we call rows.Err() to retrieve any
    // error that was encountered during the iteration. It's important to
    // call this - don't assume the iteration completed successfully over the
    // entire result set.
    if err = rows.Err(); err != nil {
        return nil, err
    }

    // If everything went OK then return the Snippets slice.
    return snippets, nil
}

```

> Important: Closing a resultset with defer rows.Close() is critical in
> the code above. As long as a resultset is open it will keep the underlying
> database connection open… so if something goes wrong in this method and
> the resultset isn’t closed, it can rapidly lead to all the connections in
> your pool being used up.

## Section 5 Dynamic HTML
### 5.1 Displaying Dynamic Data
Currently our snippetView handler function fetches a models.Snippet value
from the database and then dumps the contents out in a plain-text HTTP
response. In this chapter we’ll improve this so that the data is displayed
in a proper HTML web page.

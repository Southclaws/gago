# gago

Quick Google analytics middleware for Golang.

## Usage

```go
ga := Client{
    ID:                 "UA-12345678",
    ClientIDContextKey: "MY_APP_user_id",
    Errors:             func(err error) { myCustomLogger(err) },
}

router.Use(
    extractUserID,
    ga.Middleware,
)

// where extractUserID is something along the lines of:

func extractUserID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := getUserID(...)

        next.ServeHTTP(w, r.WithContext(context.WithValue(
            r.Context(),
            "MY_APP_user_id",
            userID,
        )))
    })
}

// and do the same for Hit Type if you need.
```

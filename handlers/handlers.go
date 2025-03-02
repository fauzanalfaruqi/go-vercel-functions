package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var n = 0

func RootPathHandler(w http.ResponseWriter, r *http.Request) {
	output := `
# Available Endpoints

This API provides the following endpoints:

## Standard Endpoints

- /hello        : Returns a simple JSON "hello world" message.
- /global-state : Increments and returns a global counter (demonstrates state).

## Endpoints with CORS Middleware

- /hello-cors   : Same as /hello, but with CORS enabled.
`

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, output)
}

// Returns a simple JSON "hello world" message.
func HelloJsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"code":    http.StatusOK,
		"message": "hello world!",
	})
}

// Increments and returns a global counter (demonstrates state).
func GlobalStateHandler(w http.ResponseWriter, r *http.Request) {
	n++
	fmt.Fprintln(w, n)
}

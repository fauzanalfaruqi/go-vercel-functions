package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var n = 0

func RootPathHandler(w http.ResponseWriter, r *http.Request) {
	output := "Available endpoints:\n- /hello\n- /global-state"
	fmt.Fprintln(w, output)
}

func HelloJsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"code":    http.StatusOK,
		"message": "hello world!",
	})
}

func GlobalStateHandler(w http.ResponseWriter, r *http.Request) {
	n++
	fmt.Fprintln(w, n)
}

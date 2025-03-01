package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Routes struct{}

var n = 0

func (rt *Routes) NStateHandler(w http.ResponseWriter, r *http.Request) {
	n++
	fmt.Fprintln(w, n)
}

func (rt *Routes) HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Code":    http.StatusOK,
		"Message": "hello world!",
	})
}

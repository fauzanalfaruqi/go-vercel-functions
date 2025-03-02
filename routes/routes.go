package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var n = 0

func NStateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("log from NStateHandler()")
	n++
	fmt.Fprintln(w, n)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("log from HelloHandler()")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Code":    http.StatusOK,
		"Message": "hello world!",
	})
}

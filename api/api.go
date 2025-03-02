package api

import (
	"api_state/handlers"
	"net/http"
)

var mux *http.ServeMux

func Handler(w http.ResponseWriter, r *http.Request) {
	mux = http.NewServeMux()

	mux.HandleFunc("/", handlers.RootPathHandler)
	mux.HandleFunc("/hello", handlers.HelloJsonHandler)
	mux.HandleFunc("/global-state", handlers.GlobalStateHandler)

	mux.ServeHTTP(w, r)
}

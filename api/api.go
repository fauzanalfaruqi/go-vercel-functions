package api

import (
	"api_state/handlers"
	"api_state/middlewares"
	"net/http"
)

var mux *http.ServeMux

func Handler(w http.ResponseWriter, r *http.Request) {
	mux = http.NewServeMux()

	mux.HandleFunc("/", handlers.RootPathHandler)
	mux.HandleFunc("/hello", handlers.HelloJsonHandler)
	mux.HandleFunc("/global-state", handlers.GlobalStateHandler)
	mux.HandleFunc("/sse", handlers.SseHandler)

	// Endpoints with CORS middleware
	mux.Handle("/hello-cors", middlewares.CorsMiddleware(http.HandlerFunc(handlers.HelloJsonHandler)))
	mux.Handle("/sse-cors", middlewares.CorsMiddleware(http.HandlerFunc(handlers.SseHandler)))

	mux.ServeHTTP(w, r)
}

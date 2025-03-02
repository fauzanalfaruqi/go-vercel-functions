package api

import (
	"api_state/routes"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	rts := routes.Routes{}
	path := r.URL.Path

	switch path {
	case "/n-state":
		rts.NStateHandler(w, r)
	case "/hello":
		rts.HelloHandler(w, r)
	case "/":
		fmt.Fprintln(w, "Welcome")
	default:
		http.NotFound(w, r) // Handle 404 for unknown paths
	}
}

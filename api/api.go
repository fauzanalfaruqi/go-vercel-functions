package api

import (
	"api_state/routes"
	"fmt"
	"net/http"
)

func FunctionsHandler(w http.ResponseWriter, r *http.Request) {
	rts := routes.Routes{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome")
	})
	http.HandleFunc("/n-state", rts.NStateHandler)
	http.HandleFunc("/hello", rts.HelloHandler)

	http.ListenAndServe(":8001", nil)
}

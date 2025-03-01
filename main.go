package main

import (
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", api.HandleNState)
	// api.FunctionsHandler()
	log.Println("Listening at port :8080")
	http.ListenAndServe(":8080", nil)
}

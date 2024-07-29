package main

import (
	"api_state/api"
	"log"
	"net/http"
)

var n = 0

func main() {
	http.HandleFunc("/", api.HandleNState)
	log.Println("Listening at port :8080")
	http.ListenAndServe(":8080", nil)
}

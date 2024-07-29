package main

import (
	"fmt"
	"log"
	"net/http"
)

var n = 0

func main() {
	http.HandleFunc("/", handleNState)
	log.Println("Listening at port :8080")
	http.ListenAndServe(":8080", nil)
}

func handleNState(w http.ResponseWriter, r *http.Request) {
	n++
	fmt.Fprintln(w, n)
}

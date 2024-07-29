package api

import (
	"fmt"
	"net/http"
)

var n = 0

func HandleNState(w http.ResponseWriter, r *http.Request) {
	n++
	fmt.Fprintln(w, n)
}

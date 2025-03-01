package api

import (
	"encoding/json"
	"net/http"
)

func HellowHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"Code":    200,
		"Message": "Hellow",
	})
}

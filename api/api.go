package api

import (
	"api_state/routes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var mux *http.ServeMux

func registerRoutes(r *http.ServeMux) {
	log.Println("log from registerRoute()")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("log from \"/\"")
		fmt.Fprintln(w, "Welcome")
	})
	r.HandleFunc("/n-state", routes.NStateHandler)
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("log from HelloHandler()")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"Code":    http.StatusOK,
			"Message": "hello world!",
		})
	})
}

func init() {
	log.Println("log from init()")
	mux = http.NewServeMux()
	registerRoutes(mux)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("mux: %#v\n", mux)
	if mux == nil {
		log.Println("log from 'mux == nil'")
		mux = http.NewServeMux()
		registerRoutes(mux)
	}
	mux.ServeHTTP(w, r)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var n = 0

func RootPathHandler(w http.ResponseWriter, r *http.Request) {
	output := `
# Available Endpoints

This API provides the following endpoints:

## Standard Endpoints

- /hello        : Returns a simple JSON "hello world" message.
- /global-state : Increments and returns a global counter (demonstrates state).
- /sse          : Server-Sent Events endpoint (sends time updates).

## Endpoints with CORS Middleware

- /hello-cors   : Same as /hello, but with CORS enabled.
- /sse-cors     : Same as /sse, but with CORS enabled.
`

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, output)
}

// Returns a simple JSON "hello world" message.
func HelloJsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"code":    http.StatusOK,
		"message": "hello world!",
	})
}

// Increments and returns a global counter (demonstrates state).
func GlobalStateHandler(w http.ResponseWriter, r *http.Request) {
	n++
	fmt.Fprintln(w, n)
}

// Server-Sent Events endpoint (sends time updates).
func SseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SseHandler called")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	log.Println("SSE headers set")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	for i := 0; i < 7; i++ {
		currentTime := time.Now().Format(time.RFC1123)
		fmt.Fprintf(w, "data: The server time is = %s\n\n", currentTime)
		flusher.Flush()
		log.Println("Minimal SSE data sent and flushed")
		time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(w, "data: Connection closed by the server!\n\n")
	flusher.Flush()
	log.Println("SseHandler finished. Connection closed by the server.")
}

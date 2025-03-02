package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contains := func(slice []string, s string) bool {
			for _, item := range slice {
				if item == s {
					return true
				}
			}
			return false
		}

		// Get the Origin header from the request
		origin := r.Header.Get("Origin")
		if origin == "" {
			// If no Origin header, proceed without CORS headers
			next.ServeHTTP(w, r)
			return
		}

		// CORS Configs
		allowedOrigins := []string{"*"}
		allowedMethods := []string{"POST", "DELETE", "GET", "OPTIONS", "PUT"}
		allowedHeaders := []string{"Origin", "Content-Type", "Authorization"}
		allowCredentials := true
		maxAge := 12 * time.Hour

		// Origin Check
		allowOrigin := ""
		if contains(allowedOrigins, "*") || contains(allowedOrigins, origin) {
			allowOrigin = origin
			if contains(allowedOrigins, "*") {
				allowOrigin = "*"
			} else {
				allowOrigin = origin
			}

		}

		// If origin is not allowed, you might choose to reject the request or proceed without CORS headers.
		// For this example, we proceed without CORS if origin is not allowed (no CORS headers set).
		if allowOrigin != "" {
			// 2. Set CORS headers on the response
			w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(allowedHeaders, ", "))
			w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
			if allowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			w.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", int(maxAge.Hours())))
		}

		// Handle Preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain for actual requests
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"log"
	"net/http"
)

// LogMiddleware returns handler function decorated with log statement.
func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer log.Printf("[HTTP] %s %s", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	}
}

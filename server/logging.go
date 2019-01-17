package server

import (
	"log"
	"net/http"
)

// LogMiddleware returns handler decorated with log statement.
func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer log.Printf("[HTTP] %s %s", r.Method, r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

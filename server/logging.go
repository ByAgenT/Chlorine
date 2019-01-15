package server

import (
	"log"
	"net/http"
)

func logHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[REQUEST] %s %s", r.RequestURI, r.Method)
		h.ServeHTTP(w, r)
	}
}

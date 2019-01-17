package middleware

import "net/http"

// Middleware type describes function that accepts HandlerFunc and returns altered HandlerFunc.
type Middleware func(http.Handler) http.Handler

// ApplyMiddlewares apply chain of middleware to a handler.
func ApplyMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

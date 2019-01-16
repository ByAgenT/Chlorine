package middleware

import "net/http"

// Middleware type describes function that accepts HandlerFunc and returns altered HandlerFunc.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ApplyMiddlewares apply chain of middleware to a handler function.
func ApplyMiddlewares(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

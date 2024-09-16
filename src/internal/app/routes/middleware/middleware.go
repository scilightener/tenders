package middleware

import (
	"net/http"
)

// Middleware is a function that wraps a http.Handler.
type Middleware func(http.Handler) http.Handler

// wrappedResponseWriter is a custom implementation of http.ResponseWriter WriteHeader method.
// It stores information about response status code.
type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader writes header to the request and stores its status code.
func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Chain creates a new middleware that chains the provided middlewares.
func Chain(ms ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			next = ms[i](next)
		}

		return next
	}
}

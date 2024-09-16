package middleware

import (
	"log/slog"
	"net/http"
	"tenders-management/internal/lib/api"
	"time"
)

// NewLoggingMiddleware creates a new logging middleware.
// It logs the request method, path, remote address, employee agent, and request ID,
// response status code and its duration.
func NewLoggingMiddleware(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := logger.With(
				"method", r.Method,
				"path", r.URL.Path,
				"remote_addr", r.RemoteAddr,
				"user_agent", r.UserAgent(),
				api.RequestIDKey, api.RequestID(r),
			)

			log.Info("request started")

			wrw := &wrappedResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			t1 := time.Now()

			next.ServeHTTP(wrw, r)

			log.Info("request completed",
				"status", wrw.statusCode,
				"duration", time.Since(t1).String(),
			)
		})
	}
}

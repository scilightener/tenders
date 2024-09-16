package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"tenders-management/internal/lib/api"
	"tenders-management/internal/lib/api/jsn"
	"tenders-management/internal/lib/api/msg"
	"tenders-management/internal/lib/logger/sl"
)

// NewRecovererMiddleware is a middleware that recovers from panics and returns
// a 500 Internal Server Error in such cases.
// It sets an error message in the response body.
func NewRecovererMiddleware(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("panic occurred. recovered.", sl.Err(fmt.Errorf("%v", err)))
					w.Header().Set("Content-Type", "application/json")
					jsn.EncodeResponse(w, http.StatusInternalServerError, api.ErrResponse(msg.APIInternalErr), logger)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

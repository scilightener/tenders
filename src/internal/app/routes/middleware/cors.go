package middleware

import "net/http"

// CORSEnableMiddleware sets the CORS header to allow-all for every request.
func CORSEnableMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

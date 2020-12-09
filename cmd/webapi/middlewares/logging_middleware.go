package middlewares

import (
	"log"
	"net/http"
)

// Logging is a middleware that logs every request received.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Method + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

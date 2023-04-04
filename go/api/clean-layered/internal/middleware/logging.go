package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware function that logs the
// request details.
func LoggingMiddleware(next http.Handler) http.Handler {
	// Your logging middleware logic goes here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request details
		fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"net/http"
)

// CachingMiddleware is a middleware function that checks the
// request against a caching policy. If the response is cached,
// the middleware returns the cached response. Otherwise, the
// middleware calls the next handler in the chain and caches
// the response.
func CachingMiddleware(next http.Handler) http.Handler {
	// Your caching middleware logic goes here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request against the caching policy
		// If the response is cached, return the cached response
		// Otherwise, call the next handler in the chain and cache the response
		next.ServeHTTP(w, r)
	})
}

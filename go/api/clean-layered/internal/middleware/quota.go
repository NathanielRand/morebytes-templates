package middleware

import (
    "net/http"
    "sync"
    "time"

    "golang.org/x/time/rate"
)

type TokenBucket struct {
    tokens         int
    lastRefillTime time.Time
}

var userTokenBuckets = make(map[string]*TokenBucket)
var wg sync.WaitGroup

func QuotaMiddleware(next http.Handler) http.Handler {
    // Allow a maximum of 100 requests per second
    rateLimiter := rate.NewLimiter(rate.Limit(100), 100)

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := r.Header.Get("X-User-ID")
        if userID == "" {
            http.Error(w, "X-User-ID header is missing", http.StatusBadRequest)
            return
        }

        // Wait for the rate limiter to allow the request to proceed
        if !rateLimiter.Allow() {
            http.Error(w, "Too many requests", http.StatusTooManyRequests)
            return
        }

		// Get the user's token bucket, or create a new one if it doesn't exist
        bucket, ok := userTokenBuckets[userID]
        if !ok {
            bucket = &TokenBucket{tokens: 10, lastRefillTime: time.Now()}
            userTokenBuckets[userID] = bucket
        }

        // Refill the bucket with new tokens
        now := time.Now()
        elapsed := now.Sub(bucket.lastRefillTime)
        tokensToAdd := int(elapsed.Seconds()) // add one token per second
        if tokensToAdd > 0 {
            bucket.tokens = min(bucket.tokens+tokensToAdd, 10)
            bucket.lastRefillTime = now
        }

        // Check if the bucket has enough tokens for the request
        if bucket.tokens <= 0 {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }

		// Decrement the number of tokens in the bucket
        bucket.tokens--

        // Increment the WaitGroup counter to indicate that a request is being processed
        wg.Add(1)
        defer wg.Done()

        // Call the next middleware/handler in the chain
        next.ServeHTTP(w, r)
    })
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

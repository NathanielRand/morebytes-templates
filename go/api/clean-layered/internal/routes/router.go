package routes

import (
	"github.com/NathanielRand/morebytes-templates/go/api/clean-layered/internal/handlers"
	"github.com/NathanielRand/morebytes-templates/go/api/clean-layered/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func SetupRouter() *mux.Router {
	// Initialize a new router from the Gorilla Mux library
	router := mux.NewRouter()

	// Create a new middleware chain using the Alice library
	chain := alice.New()

	// Add middleware to the chain for authentication, rate limiting, caching, and quotas
	chain = chain.Append(middleware.AuthenticationMiddleware)
	chain = chain.Append(middleware.RateLimitingMiddleware)
	chain = chain.Append(middleware.CachingMiddleware)
	chain = chain.Append(middleware.QuotaMiddleware)
	chain = chain.Append(middleware.SecurityMiddleware)
	chain = chain.Append(middleware.LoggingMiddleware)

	// Add your API endpoints to the router
	router.Handle("/api/v1/hello", chain.ThenFunc(handlers.HelloHandler)).Methods("GET")

	return router
}

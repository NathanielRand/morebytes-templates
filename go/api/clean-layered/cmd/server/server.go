package server

import (
	"net/http"

	"github.com/NathanielRand/morebytes-templates/go/api/clean-layered/internal/routes"
)

func StartServer() error {
	// Get the router from the routes package
	router := routes.SetupRouter()

	// Start the HTTP server
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return err
	}

	return nil
}

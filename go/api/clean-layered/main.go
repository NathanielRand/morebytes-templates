package main

import (
	"log"

	"github.com/NathanielRand/morebytes-templates/go/api/clean-layered/internal/server"
)

func main() {
	err := server.StartServer()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

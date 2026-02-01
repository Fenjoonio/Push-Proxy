package main

import (
	"log"
	"net/http"
	"os"

	"github.com/freakingeek/phoxy/internals/routes"
)

func main() {
	routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := "0.0.0.0:" + port
	log.Printf("Proxy server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

package main

import (
	"log"
	"net/http"

	"github.com/freakingeek/phoxy/internals/routes"
)



func main() {
	routes.SetupRoutes()

	log.Println("Proxy server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

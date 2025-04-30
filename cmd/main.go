package main

import (
	"log"
	"net/http"

	"github.com/freakingeek/phoxy/internals/routes"
)



func main() {
	routes.SetupRoutes()

	log.Println("Proxy server running on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

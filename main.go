package main

import (
	"log"
	"net/http"

	server "ascii-art-web/server"
)

func main() {
	server.InitTemplates()
	server.SetupRoutes()
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

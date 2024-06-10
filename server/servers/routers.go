package server

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
}

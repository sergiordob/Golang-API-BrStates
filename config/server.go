package config

import (
	"log"
	"net/http"
)

var Mux = http.NewServeMux()

func InitServer() {
	server := http.Server{
		Addr:    ":8080",
		Handler: Mux,
	}
	log.Println("Starting server at :8080")
	server.ListenAndServe()
}

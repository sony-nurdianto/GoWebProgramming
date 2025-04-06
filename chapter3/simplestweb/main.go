package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}

	log.Println("simplestweb Running On 0.0.0.0:8080")

	server.ListenAndServe()
}

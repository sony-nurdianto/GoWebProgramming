package main

import "net/http"

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	server.ListenAndServeTLS("cmd/cert.pem", "cmd/key.pem")
}

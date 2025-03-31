package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Server Running On 0.0.0.0:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

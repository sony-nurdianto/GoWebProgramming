package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello %s!\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

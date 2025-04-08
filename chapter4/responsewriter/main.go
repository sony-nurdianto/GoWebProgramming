package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `
    <!DOCTYPE html>
    <html lang="en">
    
    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <title>Another Example</title>
      <link href="css/style.css" rel="stylesheet">
    </head>
    
    <body>
      <h1>Hello Go</h1>
    </body>
    
    </html>
  `

	w.Write([]byte(str))
}

func writeHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "No Such Service, try next door")
}

func headerRedirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

type Users struct {
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Another string `json:"duknow,omitempty"`
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	u := &Users{
		Name:    "Impostor Engineer",
		Age:     0,
		Another: "Another",
	}

	json, err := json.Marshal(u)
	if err != nil {
		log.Println("Failed To Parsing Json")
		return
	}

	fmt.Fprintln(w, string(json))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/write", writeExample)
	mux.HandleFunc("/writeheader", writeHeader)
	mux.HandleFunc("/redirect", headerRedirect)
	mux.HandleFunc("/json", jsonExample)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

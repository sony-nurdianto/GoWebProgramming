package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, "Hello World")
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":8080", nil)
}

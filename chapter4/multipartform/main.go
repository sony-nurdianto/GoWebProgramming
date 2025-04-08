package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func processfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(1084); err != nil {
		log.Println("Error Parsing Form")
		http.Error(w, fmt.Sprintf("Error Parsing Form: %s", err), http.StatusInternalServerError)
		return
	}

	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("Error Open File")
		http.Error(w, "Error Open File", http.StatusInternalServerError)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error Read File")
		http.Error(w, fmt.Sprintf("Error Read File: %s", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(data))
}

func singleProcessFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err != nil {
		http.Error(w, "Error Retriving File", http.StatusInternalServerError)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error Read File ", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(data))
}

func previewImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("uploaded")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	switch {
	case strings.HasSuffix(header.Filename, ".jpg"), strings.HasSuffix(header.Filename, ".jpeg"):
		w.Header().Set("Content-Type", "image/jpeg")
	case strings.HasSuffix(header.Filename, ".png"):
		w.Header().Set("Content-Type", "image/png")
	case strings.HasSuffix(header.Filename, ".gif"):
		w.Header().Set("Content-Type", "image/gif")
	default:
		http.Error(w, "unsuported Image File", http.StatusBadRequest)

	}

	io.Copy(w, file)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("POST /uploaded", previewImage)
	server.ListenAndServe()
}

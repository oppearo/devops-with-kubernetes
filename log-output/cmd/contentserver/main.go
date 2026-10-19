package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var filePath string



func statusHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Tiedostoa ei ole olemassa", http.StatusServiceUnavailable)
			return
		}
		http.Error(w, "Tiedoston lukeminen epäonnistui", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(data)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love this! See /status")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	filePath = os.Getenv("FILE_PATH")
	if filePath == "" {
		filePath = "/shared/status.txt"
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/status", statusHandler)

	fmt.Printf("Server listening on port %s, reading from %s\n", port, filePath)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
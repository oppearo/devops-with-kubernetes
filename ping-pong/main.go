package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var counter int

func pingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "pong %d", counter)
    fmt.Printf("pong %d\n", counter) // myös konsoliin
    counter++
}

func main() {
    port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

    http.HandleFunc("/pingpong", pingHandler)
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


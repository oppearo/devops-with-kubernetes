package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// tulostetaan URLista / jälkeinen osa
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// gon dokumentaation mukaan
	http.HandleFunc("/", handler)
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

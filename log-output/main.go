package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	currentString string
	mu sync.RWMutex
)


func handler(w http.ResponseWriter, r *http.Request) {
	// tulostetaan URLista / jälkeinen osa
	fmt.Fprintf(w, "Hi there, I love this! See /status")
}

// statusHandler is called every time someone makes an HTTP request to "/status"
func statusHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()         // Lock for reading
	val := currentString
	mu.RUnlock()       // Release the lock

	// Write the string back to the browser / curl
	fmt.Fprint(w, val)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	uuid, _ := uuid.NewRandom()

	go func() {
		for {
			now := time.Now().Format(time.RFC3339)
			output := fmt.Sprintf("%s: %s\n", now, uuid)
			fmt.Print(output) // stdout for debugging

			// Lock for writing
			mu.Lock()
			currentString = output
			mu.Unlock()

			time.Sleep(5 * time.Second)
		}
	}()

	// gon dokumentaation mukaan
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/", handler)
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func main() {
	filePath := os.Getenv("FILE_PATH")
	if filePath == "" {
		filePath = "/shared/status.txt"
	}

	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		log.Fatalf("Kansion luominen epäonnistui: %v", err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("UUID:n generointi epäonnistui: %v", err)
	}

	for {
		now := time.Now().Format(time.RFC3339)
		output := fmt.Sprintf("%s: %s\n", now, id)

		// tulosta konsoliin
		fmt.Print(output)

		// kirjoita tiedostoon
		if err := os.WriteFile(filePath, []byte(output), 0644); err != nil {
			log.Printf("Virhe tiedostoon kirjoittamisessa: %v", err)
		}

		time.Sleep(5 * time.Second)
	}
}

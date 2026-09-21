package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	uuid, _ := uuid.NewRandom()
	for {
		now := time.Now().Format(time.RFC3339)
		fmt.Printf("%s: %s\n", now, uuid)
		time.Sleep(5 * time.Second)
	}
}

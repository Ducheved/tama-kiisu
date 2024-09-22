package main

import (
	"log"

	"github.com/Ducheved/tama-kiisu/cmd/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

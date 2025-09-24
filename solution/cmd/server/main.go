package main

import (
	"log"

	"github.com/haohanl/oolio-challenge/solution/pkg/server"
)

func main() {
	// Create server instance
	srv := server.New(":8080")

	// Start server
	if err := srv.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

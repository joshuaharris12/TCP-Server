package main

import (
	"fmt"
	"os"

	"github.com/joshuharris12/tcp-server/pkg/server"
)

const (
	PORT string = "4567"
)

func main() {
	server, err := server.NewServer(PORT)
	if err != nil {
		fmt.Println("LOG: Failed to create server: %w", err)
		os.Exit(1)
	}

	if err := server.Run(); err != nil {
		fmt.Println("LOG: Failed to run server: %w", err)
		os.Exit(1)
	}
}

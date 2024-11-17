package main

import (
	"log"
	"main/internal/config"
	"main/internal/server"
)

func main() {
	config.Config()

	log.Printf("Server starting... %s", config.Server)
	err := server.Run(config.Server)
	if err != nil {
		panic(err)
	}
}

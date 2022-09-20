package main

import (
	"log"
	"moderation_service/interfaces/middleware/server"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	server.InitServer().RunServer()
}

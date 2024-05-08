package main

import "github.com/somos831/somos-backend/handlers"

func main() {
	server := handlers.Server{}

	server.InitServer()
	server.Run(":8080")
}

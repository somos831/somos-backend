package server

import "github.com/somos831/somos-backend/handlers"

func Run() {
	server := handlers.Server{}

	server.InitServer()
	server.Run(":8080")
}

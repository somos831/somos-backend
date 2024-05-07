package server

import "github.com/somos831/somos-backend/handlers"

var server = handlers.Server{}

func Run() {
	server.Initialize()
	server.Run(":8080")
}

package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	conn "github.com/somos831/somos-backend/db"
	"github.com/somos831/somos-backend/validators"
)

type Server struct {
	db        *sql.DB
	Router    *mux.Router
	Validator validators.Validator
}

func (server *Server) InitServer() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db := conn.Connect()
	server.db = db

	// Initialize new router:
	server.Router = mux.NewRouter()
	server.Router.Use(ContextMiddleware)

	// Initialize routes:
	server.InitRoutes()

	// Initialize validator:
	server.Validator = validators.NewValidator(db)
}

func (server *Server) Run(addr string) {

	go func() {
		// Handle termination signals
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

		// Block until a signal is received
		<-sigint

		// Shutdown gracefully
		log.Println("Shutting down server...")

		// Perform cleanup tasks before exiting
		conn.Disconnect(server.db)

		// Exit
		os.Exit(0)
	}()

	log.Printf("Listening on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

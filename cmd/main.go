package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	conn "github.com/somos831/somos-backend/db"
	"github.com/somos831/somos-backend/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := "somos_events"
	db := conn.Connect(dbUser, dbPassword, dbName)

	defer conn.Disconnect(db)
	conn.ExecuteSchemaFromFile(db, "db/schema.sql")
	handler := handlers.New(db)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /events",
		handlers.HttpHandler(handler.ListAllEvents))
	mux.HandleFunc("GET /events/{id}",
		handlers.HttpHandler(handler.GetEvent))
	mux.HandleFunc("POST /events",
		handlers.HttpHandler(handler.CreateEvent))
	mux.HandleFunc("PATCH /events/{id}",
		handlers.HttpHandler(handler.UpdateEvent))
	mux.HandleFunc("DELETE /events/{id}",
		handlers.HttpHandler(handler.DeleteEvent))

	mux.HandleFunc("GET /categories",
		handlers.HttpHandler(handler.ListAllCategories))
	mux.HandleFunc("GET /categories/{id}",
		handlers.HttpHandler(handler.GetCategory))
	mux.HandleFunc("POST /categories",
		handlers.HttpHandler(handler.CreateCategory))
	mux.HandleFunc("PATCH /categories/{id}",
		handlers.HttpHandler(handler.UpdateCategory))
	mux.HandleFunc("DELETE /categories/{id}",
		handlers.HttpHandler(handler.DeleteCategory))

	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

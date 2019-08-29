package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/trucntt1101/testusername/database"
	"github.com/trucntt1101/testusername/handler"
)

func main() {

	database.ConnectDB
	r := chi.NewRouter()

	r.Get("/users", handler.Get)
	r.Get("/users/{id}", handler.Gett)
	r.Post("/users", handler.Post)
	r.Put("/users", handler.Put)
	r.Patch("/users", handler.Patch)
	r.Delete("/users", handler.Delete)

	http.ListenAndServe(":3333", r)
}

package main

import (
	"net/http"

	"github.com/trucntt1101/testusername/handler"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handler.Print())
	http.ListenAndServe(":3333", r)
}

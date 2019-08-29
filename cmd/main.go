package main

import (
	"handler"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handler.Print())
	http.ListenAndServe(":3333", r)
}

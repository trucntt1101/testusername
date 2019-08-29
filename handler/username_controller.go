package handler

import (
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
}

func Gett(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
}
func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
}

func Patch(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PATCH")
}
func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
}

package main

import (
	"log"
	"net/http"
	a "./api"
	"gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	a.CreateApi(*r)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

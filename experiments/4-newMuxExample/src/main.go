package main

import (
	h "./handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getall", h.GetAll)

	log.Println("Listening...")
	http.ListenAndServe(":3000", r)
}
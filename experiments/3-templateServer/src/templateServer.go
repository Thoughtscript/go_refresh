package main

import (
	"log"
	"net/http"
	h "./handlers"
)

func main() {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/", http.StripPrefix("/public/", fileServer))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	http.HandleFunc("/home", h.ServeHome)

	log.Println("Listening...")
	http.ListenAndServe(":7777", nil)
}
package api

import (
	"gorilla/mux"
	"net/http"
	"encoding/json"
	s "../data"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	exampleSchema := s.ExampleSchema{
		ID:        1,
		FirstName: "Adam",
		LastName:  "Gerard",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exampleSchema)
}

func CreateApi(r mux.Router) {
	r.HandleFunc("/example", GetAll).Methods("GET")
}
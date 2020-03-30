package handlers

import (
	"net/http"
	"fmt"
)

func PrintOut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

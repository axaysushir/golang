package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	r.Methods("GET")
}

func HomeHandler() {
	fmt.Println("Welcome Home")
}

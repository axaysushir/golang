package main

import (
	"fmt"
	"net/http"
)

type Course struct {
	Name   string
	Id     int
	Author string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	r.Methods("GET")
}

func HomeHandler() {
	fmt.Println("Welcome Home")
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course : range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
}

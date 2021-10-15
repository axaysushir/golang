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

func creteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	var course Course
	_= json.NewDecoder(r.body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data")
		return 
	}

	// generate ID and append course
	rand.Seed(time.Now().UnixNano())
	Course.CourseId = rand.Intn(12345)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}


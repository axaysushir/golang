package router

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.getMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.createMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.markAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.deleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.deleteAllMovies).Methods("DELETE")

	return router
}

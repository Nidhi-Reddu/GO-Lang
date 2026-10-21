package routers

import (
	"temp/handlers"

	"github.com/gorilla/mux"
)

func SetUpRoutes() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	mux.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	mux.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
	mux.HandleFunc("/posts/{id}", handlers.DeletePosts).Methods("DELETE")
	mux.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")

	return mux
}

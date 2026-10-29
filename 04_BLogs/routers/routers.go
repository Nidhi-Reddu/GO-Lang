package routers

import (
	"temp/handlers"

	"github.com/gorilla/mux"
)

func SetUpRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/about", handlers.AboutHandler)
	r.HandleFunc("/post", handlers.GetPost).Methods("GET")

	r.HandleFunc("/posts", handlers.RequireLogin(handlers.CreatePost)).Methods("GET", "POST")
	r.HandleFunc("/posts/{id}", handlers.RequireLogin(handlers.UpdatePost)).Methods("GET", "POST", "PUT")
	r.HandleFunc("/posts/{id}", handlers.RequireLogin(handlers.DeletePosts)).Methods("POST", "DELETE")

	// REST API routes
	r.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.RequireLogin(handlers.UpdatePost)).Methods("POST", "PUT")
	r.HandleFunc("/posts/{id}", handlers.RequireLogin(handlers.DeletePosts)).Methods("DELETE")

	r.HandleFunc("/my_posts", handlers.RequireLogin(handlers.MyPosts)).Methods("GET")

	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/signup", handlers.SignUpHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")

	return r
}

package main

import (
	"github.com/gorilla/mux"
)

func initRoutes(r *mux.Router) {
	r.HandleFunc("/posts", GetAllPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", GetPostByID).Methods("GET")
	r.HandleFunc("/posts", CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", DeletePost).Methods("DELETE")
}

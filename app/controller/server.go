package controller

import (
	"fmt"
	"log"
	"net/http"

	controller "example.com/go_rest_postgre/controller/blog"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/post/{id}", controller.GetPost).Methods("GET")
	router.HandleFunc("/api/post/new", controller.CreatePost).Methods("POST")
	router.HandleFunc("/api/post/update", controller.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/post/delete/{id}", controller.DeletePost).Methods("DELETE")
}

func Init() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("Router initialized and listening :9090 \n")
	log.Fatal(http.ListenAndServe(":9090", router))
}

package gorillamux

import (
	"github.com/gorilla/mux"
	"go-routers/middleware"
	"log"
	"net/http"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.Auth, middleware.Log)
	router.HandleFunc("/", handle)
	return router
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("handling user request")
	w.Write([]byte("Hello world"))
}

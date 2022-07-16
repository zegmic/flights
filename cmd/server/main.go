package main

import (
	"flights/internal/server"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	s := server.RouteServer{}
	r.HandleFunc("/flight", s.Find).Methods("POST")
	http.ListenAndServe(":8080", r)
}

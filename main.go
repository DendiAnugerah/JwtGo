package main

import (
	"net/http"

	"github.com/DendiAnugerah/JwtGo/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/register", controller.Register).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("Get")

	http.ListenAndServe(":8000", r)
}

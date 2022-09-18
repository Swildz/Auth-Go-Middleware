package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swildz/go-jwt-siddiq/controllers/authcontroller"
	"github.com/swildz/go-jwt-siddiq/controllers/productcontroller"
	"github.com/swildz/go-jwt-siddiq/middleware"
	"github.com/swildz/go-jwt-siddiq/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.Use(middleware.JWTMiddleware)
	log.Fatal(http.ListenAndServe(":8082", r))

}

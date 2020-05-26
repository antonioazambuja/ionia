package main

import (
	"net/http"

	handler_v1 "github.com/antonioazambuja/ionia/handlers/api/v1"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/summoner/byname/{name}", handler_v1.GetByName).Methods("GET")
	router.HandleFunc("/health", handler_v1.HealthCheck).Methods("GET")

	http.ListenAndServe(":5000", router)
}

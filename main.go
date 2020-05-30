package main

import (
	"log"
	"net/http"
	"os"

	handler_v1 "github.com/antonioazambuja/ionia/handlers/api/v1"
	"github.com/gorilla/mux"
)

func main() {
	var logOperation = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	router := mux.NewRouter()
	router.HandleFunc("/summoner/byname/{name}", handler_v1.GetByName).Methods("GET")
	router.HandleFunc("/summoner/byname/{name}/matches", handler_v1.GetMatchesByName).Methods("GET")
	router.HandleFunc("/summoner/byname/{name}/league", handler_v1.GetLeagueByName).Methods("GET")
	router.HandleFunc("/summoner/byname/{name}/info", handler_v1.GetInfoByName).Methods("GET")
	router.HandleFunc("/health", handler_v1.HealthCheck).Methods("GET")
	logOperation.Print("Custom Golang API Riot!")
	logOperation.Print("Serving at port 5000!")
	log.Fatal(http.ListenAndServe(":5000", router))
}

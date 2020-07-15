package main

import (
	"log"
	"net/http"
	"os"

	handler_v1 "github.com/antonioazambuja/ionia/app/handlers/api/v1"
	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	utils "github.com/antonioazambuja/ionia/utils"
	"github.com/gorilla/mux"
)

func init() {
	rsc_v1.RedisClientConnected = rsc_v1.NewRedisClient(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PORT"))
	rsc_v1.RedisClientConnected.Connect(os.Getenv("REDIS_PWD"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/v1/summoner/byname/{name}/matches/current/year", handler_v1.GetMatchesCurrentYear).Methods("GET")
	router.HandleFunc("/v1/summoner/byname/{name}/matches/current/month", handler_v1.GetMatchesCurrentMonth).Methods("GET")
	router.HandleFunc("/v1/summoner/byname/{name}/matches/current/day", handler_v1.GetMatchesCurrentDay).Methods("GET")
	router.HandleFunc("/v1/summoner/byname/{name}", handler_v1.GetByName).Methods("GET")
	router.HandleFunc("/v1/summoner/byname/{name}/matches", handler_v1.GetMatchesByName).Methods("GET")
	router.HandleFunc("/v1/summoner/byname/{name}/league", handler_v1.GetLeagueByName).Methods("GET")
	router.HandleFunc("/v1/summoner/byname/{name}/info", handler_v1.GetInfoByName).Methods("GET")
	router.HandleFunc("/v1/health", handler_v1.HealthCheck).Methods("GET")
	utils.LogOperation.Print("Custom Golang API Riot!")
	utils.LogOperation.Print("Serving at port 5000!")
	log.Fatal(http.ListenAndServe(":5000", router))
}

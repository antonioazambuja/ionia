package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	svc_v1 "github.com/antonioazambuja/ionia/services/api/v1"
	"github.com/gorilla/mux"
)

var logOperation = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)

// GetByName - get summoner by name
func GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetByName(params["name"])
	if err != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed get summoner in service GetByName")
	}
	json.NewEncoder(w).Encode(summoner)
	logOperation.Print("Perform GetByName")
}

// GetMachesByName - get matches summoner by name
func GetMachesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetByName(params["name"])
	if err != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed get summoner in service GetMachesByName")
	}
	json.NewEncoder(w).Encode(summoner)
	logOperation.Print("Perform GetMachesByName")
}

// GetInfoByName - get info summoner by name
func GetInfoByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetInfoByName(params["name"])
	if err != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed get summoner in service GetInfoByName")
	}
	json.NewEncoder(w).Encode(summoner)
	logOperation.Print("Perform GetInfoByName")
}

// GetMatchesByName - get info summoner by name
func GetMatchesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetMatchesByName(params["name"])
	if err != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed get summoner in service GetMatchesByName")
	}
	json.NewEncoder(w).Encode(summoner)
	logOperation.Print("Perform GetMatchesByName")
}

// GetLeagueByName - get info summoner by name
func GetLeagueByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetLeagueByName(params["name"])
	if err != nil {
		logOperation.Print("Failed get summoner in service GetLeagueByName")
	}
	logOperation.Print("Perform GetLeagueByName")
	json.NewEncoder(w).Encode(summoner)
}

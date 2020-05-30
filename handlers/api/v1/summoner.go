package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	svc_v1 "github.com/antonioazambuja/ionia/services/api/v1"
	"github.com/gorilla/mux"
)

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
}

// GetMachesByName - get matches summoner by name
func GetMachesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetByName(params["name"])
	if err != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed get summoner in service GetByName")
	}
	json.NewEncoder(w).Encode(summoner)
}

// GetInfoByName - get info summoner by name
func GetInfoByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetInfoByName(params["name"])
	if err != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed get summoner in service GetByName")
	}
	json.NewEncoder(w).Encode(summoner)
}

// GetMatchesByName - get info summoner by name
func GetMatchesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetMatchesByName(params["name"])
	if err != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed get summoner in service GetByName")
	}
	json.NewEncoder(w).Encode(summoner)
}

// GetLeagueByName - get info summoner by name
func GetLeagueByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetLeagueByName(params["name"])
	if err != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed get summoner in service GetByName")
	}
	json.NewEncoder(w).Encode(summoner)
}

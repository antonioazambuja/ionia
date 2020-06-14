package v1

import (
	"encoding/json"
	"net/http"

	svc_v1 "github.com/antonioazambuja/ionia/app/services/api/v1"
	utils "github.com/antonioazambuja/ionia/utils"
	"github.com/gorilla/mux"
)

// GetByName - get summoner by name
func GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetByName(params["name"])
	if err != nil {
		utils.LogOperation.Print("Failed get summoner in service GetByName")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(summoner)
	utils.LogOperation.Print("Perform GetByName")
}

// GetInfoByName - get info summoner by name
func GetInfoByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetInfoByName(params["name"])
	if err != nil {
		utils.LogOperation.Print("Failed get summoner in service GetInfoByName")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(summoner)
	utils.LogOperation.Print("Perform GetInfoByName")
}

// GetMatchesByName - get info summoner by name
func GetMatchesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetMatchesByName(params["name"])
	if err != nil {
		utils.LogOperation.Print("Failed get summoner in service GetMatchesByName")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(summoner)
	utils.LogOperation.Print("Perform GetMatchesByName")
}

// GetLeagueByName - get info summoner by name
func GetLeagueByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetLeagueByName(params["name"])
	if err != nil {
		utils.LogOperation.Print("Failed get summoner in service GetLeagueByName")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(summoner)
	utils.LogOperation.Print("Perform GetLeagueByName")
}
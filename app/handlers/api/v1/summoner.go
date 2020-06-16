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
	var statusCode int
	params := mux.Vars(r)
	summoner, err := svc_v1.GetByName(params["name"])
	if err != nil {
		utils.LogOperation.Print(err)
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusOK
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetByName")
}

// GetInfoByName - get info summoner by name
func GetInfoByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	params := mux.Vars(r)
	summoner, err := svc_v1.GetInfoByName(params["name"])
	if err != nil {
		utils.LogOperation.Print(err)
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusOK
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetInfoByName")
}

// GetMatchesByName - get info summoner by name
func GetMatchesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	params := mux.Vars(r)
	summoner, err := svc_v1.GetMatchesByName(params["name"])
	if err != nil {
		utils.LogOperation.Print(err)
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusOK
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetMatchesByName")
}

// GetLeagueByName - get info summoner by name
func GetLeagueByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	params := mux.Vars(r)
	summoner, err := svc_v1.GetLeagueByName(params["name"])
	if err != nil {
		utils.LogOperation.Print(err)
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusOK
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetLeagueByName")
}

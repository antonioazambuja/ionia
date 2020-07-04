package v1

import (
	"encoding/json"
	"net/http"
	"os"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	svc_v1 "github.com/antonioazambuja/ionia/app/services/api/v1"
	utils "github.com/antonioazambuja/ionia/utils"
	"github.com/gorilla/mux"
)

// GetByName - get summoner by name
func GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var statusCode int
	params := mux.Vars(r)
	riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
	summoner, errGetByName := svc_v1.GetByName(riotAPIClient, params["name"])
	if errGetByName != nil {
		utils.LogOperation.Print(errGetByName)
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
	riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
	summoner, errGetInfoByName := svc_v1.GetInfoByName(riotAPIClient, params["name"])
	if errGetInfoByName != nil {
		utils.LogOperation.Print(errGetInfoByName)
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
	riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
	summoner, errGetMatchesByName := svc_v1.GetMatchesByName(riotAPIClient, params["name"])
	if errGetMatchesByName != nil {
		utils.LogOperation.Print(errGetMatchesByName)
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
	riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
	summoner, errGetLeagueByName := svc_v1.GetLeagueByName(riotAPIClient, params["name"])
	if errGetLeagueByName != nil {
		utils.LogOperation.Print(errGetLeagueByName)
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusOK
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetLeagueByName")
}

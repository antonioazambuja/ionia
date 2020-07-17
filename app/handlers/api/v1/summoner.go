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

// GetByName - Get summoner by name
func GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := 200
	summoner := new(rsc_v1.Summoner)
	params := mux.Vars(r)
	summoner, erroSummonerCached := svc_v1.GetRedisSummoner(rsc_v1.RedisClientConnected, params["name"], "full")
	if erroSummonerCached != nil {
		utils.LogOperation.Println(erroSummonerCached)
		riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
		summonerRiotAPI, errGetByName := svc_v1.GetByName(riotAPIClient, params["name"])
		if errGetByName != nil {
			utils.LogOperation.Print(errGetByName)
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusOK
		}
		summoner = summonerRiotAPI
		if errSaveCacheSummoner := svc_v1.SaveRedisSummoner(rsc_v1.RedisClientConnected, summoner, "full"); errSaveCacheSummoner != nil {
			utils.LogOperation.Print(errSaveCacheSummoner)
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetByName")
}

// GetInfoByName - Get info summoner by name
func GetInfoByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := 200
	params := mux.Vars(r)
	summoner, erroSummonerCached := svc_v1.GetRedisSummoner(rsc_v1.RedisClientConnected, params["name"], "info")
	if erroSummonerCached != nil {
		utils.LogOperation.Println(erroSummonerCached)
		riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
		summonerRiotAPI, errGetByName := svc_v1.GetInfoByName(riotAPIClient, params["name"])
		if errGetByName != nil {
			utils.LogOperation.Print(errGetByName)
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusOK
		}
		summoner = summonerRiotAPI
		if errSaveCacheSummoner := svc_v1.SaveRedisSummoner(rsc_v1.RedisClientConnected, summoner, "info"); errSaveCacheSummoner != nil {
			utils.LogOperation.Print(errSaveCacheSummoner)
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetInfoByName")
}

// GetMatchesByName - Get info summoner by name
func GetMatchesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := 200
	params := mux.Vars(r)
	summoner, erroSummonerCached := svc_v1.GetRedisSummoner(rsc_v1.RedisClientConnected, params["name"], "matches")
	if erroSummonerCached != nil {
		utils.LogOperation.Println(erroSummonerCached)
		riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
		summonerRiotAPI, errGetByName := svc_v1.GetMatchesByName(riotAPIClient, params["name"])
		if errGetByName != nil {
			utils.LogOperation.Print(errGetByName)
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusOK
		}
		summoner = summonerRiotAPI
		if errSaveCacheSummoner := svc_v1.SaveRedisSummoner(rsc_v1.RedisClientConnected, summoner, "matches"); errSaveCacheSummoner != nil {
			utils.LogOperation.Print(errSaveCacheSummoner)
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetMatchesByName")
}

// GetLeagueByName - Get info summoner by name
func GetLeagueByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := 200
	params := mux.Vars(r)
	summoner, erroSummonerCached := svc_v1.GetRedisSummoner(rsc_v1.RedisClientConnected, params["name"], "leagues")
	if erroSummonerCached != nil {
		utils.LogOperation.Println(erroSummonerCached)
		riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
		summonerRiotAPI, errGetByName := svc_v1.GetLeagueByName(riotAPIClient, params["name"])
		if errGetByName != nil {
			utils.LogOperation.Print(errGetByName)
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusOK
		}
		summoner = summonerRiotAPI
		if errSaveCacheSummoner := svc_v1.SaveRedisSummoner(rsc_v1.RedisClientConnected, summoner, "leagues"); errSaveCacheSummoner != nil {
			utils.LogOperation.Print(errSaveCacheSummoner)
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetLeagueByName")
}


// GetMatchesCurrentYear - Get summoner by name
func GetMatchesCurrentYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := 200
	summoner := new(rsc_v1.Summoner)
	params := mux.Vars(r)
	summoner, erroSummonerCached := svc_v1.GetRedisSummoner(rsc_v1.RedisClientConnected, params["name"], "matches_current_year")
	if erroSummonerCached != nil {
		utils.LogOperation.Println(erroSummonerCached)
		riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
		summonerRiotAPI, errGetByName := svc_v1.GetMatchesCurrentYear(riotAPIClient, params["name"])
		if errGetByName != nil {
			utils.LogOperation.Print(errGetByName)
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusOK
			summoner = summonerRiotAPI
			if errSaveCacheSummoner := svc_v1.SaveRedisSummoner(rsc_v1.RedisClientConnected, summoner, "matches_current_year"); errSaveCacheSummoner != nil {
				utils.LogOperation.Print(errSaveCacheSummoner)
			}
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetMatchesCurrentYear")
}

// GetMatchesCurrentMonth - Get summoner by name
func GetMatchesCurrentMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := 200
	summoner := new(rsc_v1.Summoner)
	params := mux.Vars(r)
	summoner, erroSummonerCached := svc_v1.GetRedisSummoner(rsc_v1.RedisClientConnected, params["name"], "matches_current_month")
	if erroSummonerCached != nil {
		utils.LogOperation.Println(erroSummonerCached)
		riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
		summonerRiotAPI, errGetByName := svc_v1.GetMatchesCurrentMonth(riotAPIClient, params["name"])
		if errGetByName != nil {
			utils.LogOperation.Print(errGetByName)
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusOK
			summoner = summonerRiotAPI
			if errSaveCacheSummoner := svc_v1.SaveRedisSummoner(rsc_v1.RedisClientConnected, summoner, "matches_current_month"); errSaveCacheSummoner != nil {
				utils.LogOperation.Print(errSaveCacheSummoner)
			}
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetMatchesCurrentMonth")
}

// GetMatchesCurrentDay - Get summoner by name
func GetMatchesCurrentDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := 200
	summoner := new(rsc_v1.Summoner)
	params := mux.Vars(r)
	summoner, erroSummonerCached := svc_v1.GetRedisSummoner(rsc_v1.RedisClientConnected, params["name"], "matches_current_day")
	if erroSummonerCached != nil {
		utils.LogOperation.Println(erroSummonerCached)
		riotAPIClient := rsc_v1.NewRiotAPIClient(os.Getenv("ENDPOINT_REGION"), os.Getenv("API_KEY"), os.Getenv("HEADER_API_KEY"))
		summonerRiotAPI, errGetByName := svc_v1.GetMatchesCurrentDay(riotAPIClient, params["name"])
		if errGetByName != nil {
			utils.LogOperation.Print(errGetByName)
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = http.StatusOK
		}
		summoner = summonerRiotAPI
		if errSaveCacheSummoner := svc_v1.SaveRedisSummoner(rsc_v1.RedisClientConnected, summoner, "matches_current_day"); errSaveCacheSummoner != nil {
			utils.LogOperation.Print(errSaveCacheSummoner)
		}
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(summoner)
	utils.ServiceLog(statusCode, r, "GetMatchesCurrentDay")
}

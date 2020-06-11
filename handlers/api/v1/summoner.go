package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"

	v1 "github.com/antonioazambuja/ionia/resources/api/v1"
	svc_v1 "github.com/antonioazambuja/ionia/services/api/v1"
	"github.com/gorilla/mux"
)

var logOperation = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)

// GetByName - get summoner by name
func GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if checkSummonerName(params["name"]) {
		summoner, err := svc_v1.GetByName(params["name"])
		if err != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Failed get summoner in service GetByName")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(summoner)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v1.Summoner{})
	}
	logOperation.Print("Perform GetByName")
}

// GetInfoByName - get info summoner by name
func GetInfoByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if checkSummonerName(params["name"]) {
		summoner, err := svc_v1.GetInfoByName(params["name"])
		if err != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Failed get summoner in service GetInfoByName")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(summoner)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v1.Summoner{})
	}
	logOperation.Print("Perform GetInfoByName")
}

// GetMatchesByName - get info summoner by name
func GetMatchesByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if checkSummonerName(params["name"]) {
		summoner, err := svc_v1.GetMatchesByName(params["name"])
		if err != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Failed get summoner in service GetMatchesByName")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(summoner)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v1.Summoner{})
	}
	logOperation.Print("Perform GetMatchesByName")
}

// GetLeagueByName - get info summoner by name
func GetLeagueByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if checkSummonerName(params["name"]) {
		summoner, err := svc_v1.GetLeagueByName(params["name"])
		if err != nil {
			logOperation.Print("Failed get summoner in service GetLeagueByName")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(summoner)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(v1.Summoner{})
	}
	logOperation.Print("Perform GetLeagueByName")
}

func checkSummonerName(summonerName string) bool {
	checkSummonerName, errEspecialCharacters := regexp.MatchString("[$&+,:;=?@#|'<>.^*()%!-]", summonerName)
	if errEspecialCharacters != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Error validate summoner name - " + summonerName)
		return false
	}
	return !checkSummonerName
}

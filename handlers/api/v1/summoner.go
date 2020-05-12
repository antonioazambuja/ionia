package v1

import (
	"encoding/json"
	"log"
	"net/http"

	svc_v1 "github.com/antonioazambuja/ionia/services/api/v1"
	"github.com/gorilla/mux"
)

// GetByName - get summoner by name
func GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetByName(params["name"])
	if err != nil {
		log.Fatal("Error handler...")
		panic(err)
		// json.NewEncoder(w).Encode(rsc_v1.Summoner{})
	}
	json.NewEncoder(w).Encode(summoner)
}

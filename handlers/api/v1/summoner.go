package v1

import (
	"encoding/json"
	"net/http"

	svc_v1 "github.com/antonioazambuja/ionia/services/api/v1"
	"github.com/gorilla/mux"
)

// responsavel por criar rotas e escutar requisições

// GetByName - get summoner by name
func GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	summoner, err := svc_v1.GetByName(params["name"], params["region"])
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(summoner)
}

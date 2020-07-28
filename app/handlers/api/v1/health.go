package v1

import (
	"net/http"
	"os"
	"time"

	utils "github.com/antonioazambuja/ionia/utils"
)

// HealthCheck - get summoner by name
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clientHealthCheck := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	_, errRequestHealthCheckRiotAPI := clientHealthCheck.Get(os.Getenv("ENDPOINT_REGION"))
	if errRequestHealthCheckRiotAPI != nil {
		utils.LogOperation.Println(errRequestHealthCheckRiotAPI.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

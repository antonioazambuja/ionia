package v1

import (
	"encoding/json"
	"net/http"
	"time"

	resources_v1 "github.com/antonioazambuja/ionia/resources/api/v1"
)

// GetByName - get summoner by name
func GetByName(summonerName string, regionName string) (resources_v1.SummonerResponse, error) {
	regions := GetAvailableRegions()
	request, errNewRequest := http.NewRequest("GET", regions[regionName]+"/lol/summoner/v4/summoners/by-name/"+summonerName, nil)
	if errNewRequest != nil {
		panic(errNewRequest)
	}
	request.Header.Set(HeaderAPIKey, APIKey)
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	response, errResponse := client.Do(request)
	if errResponse != nil && response.StatusCode != 200 {
		panic(errResponse)
	}
	defer response.Body.Close()
	var summoner resources_v1.SummonerResponse
	json.NewDecoder(response.Body).Decode(&summoner)
	return summoner, nil
}

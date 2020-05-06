package v1

import (
	"encoding/json"
	"net/http"
	"time"

	resources_v1 "github.com/antonioazambuja/ionia/resources/api/v1"
)

// GetByName - get summoner by name
func GetByName(summonerName string, regionName string) (resources_v1.Summoner, error) {
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

	var summonerResponse resources_v1.SummonerResponse
	json.NewDecoder(response.Body).Decode(&summonerResponse)

	leagueEntriesResponse, err := _GetLeagueEntries(summonerResponse.ID, regionName)
	if err != nil {
		panic(err)
	}

	return resources_v1.Summoner{
		SummonerName:  summonerResponse.Name,
		SummonerLevel: summonerResponse.SummonerLevel,
		CurrentRank:   leagueEntriesResponse.Rank,
		CurrentTier:   leagueEntriesResponse.Tier,
		LeaguePoints:  leagueEntriesResponse.LeaguePoints,
		Wins:          leagueEntriesResponse.Wins,
		Losses:        leagueEntriesResponse.Losses,
		QueueType:     leagueEntriesResponse.QueueType,
	}, nil
}

// _GetLeagueEntries - get data league of summoner
func _GetLeagueEntries(summonerID string, regionName string) (resources_v1.LeagueEntriesResponse, error) {
	regions := GetAvailableRegions()
	request, errNewRequest := http.NewRequest("GET", regions[regionName]+"/lol/league/v4/entries/by-summoner/"+summonerID, nil)
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
	var leagueEntriesResponse []resources_v1.LeagueEntriesResponse
	json.NewDecoder(response.Body).Decode(&leagueEntriesResponse)
	return leagueEntriesResponse[0], nil
}

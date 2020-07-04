package v1

import (
	"encoding/json"
	"net/http"
	"time"

	// svc_v1 "github.com/antonioazambuja/ionia/app/services/api/v1"
	"github.com/antonioazambuja/ionia/utils"
)

const summonerV4 string = "/lol/summoner/v4/summoners/by-name/"

const leagueV4 string = "/lol/league/v4/entries/by-summoner/"

const matchesV4 string = "/lol/match/v4/matchlists/by-account/"

// RiotAPIClient -
type RiotAPIClient struct {
	ServerURL         string
	TokenAPI          string
	HeaderAPI         string
	RiotAPIClientFunc interface{}
}

func NewRiotAPIClient(serverURL, tokenAPI, headerAPI string) *RiotAPIClient {
	return &RiotAPIClient{
		ServerURL: serverURL,
		TokenAPI:  tokenAPI,
		HeaderAPI: headerAPI,
	}
}

func (riotAPIClient *RiotAPIClient) GetSummonerByName(summonerName string) *SummonerDTO {
	newRequest, errNewRequest := http.NewRequest("GET", riotAPIClient.ServerURL+summonerV4+summonerName, nil)
	if errNewRequest != nil {
		utils.LogOperation.Println(errNewRequest.Error())
		// return nil
	}
	newRequest.Header.Set(riotAPIClient.HeaderAPI, riotAPIClient.TokenAPI)
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	response, errResponse := client.Do(newRequest)
	if errResponse != nil {
		utils.LogOperation.Println(errResponse)
		// utils.LogOperation.Print("Failed perform request: " + endpoint + " - " + errResponse.Error())
		// return nil
	} else if response.StatusCode != 200 {
		// utils.LogOperation.Print("Failed get request with following info: '" + req.pathParam + "' in: '" + req.endpoint + "' - Invalid status code: " + response.Status)
		// return nil
	}
	summonerDTO := new(SummonerDTO)
	if errDecodeSummonerResponse := json.NewDecoder(response.Body).Decode(&summonerDTO); errDecodeSummonerResponse != nil {
		utils.LogOperation.Println(errDecodeSummonerResponse)
		// return nil
	}
	defer response.Body.Close()
	// summoner := new(Summoner)
	// summoner.SummonerName = summonerDTO.Name
	// summoner.SummonerLevel = summonerDTO.SummonerLevel
	// summoner.SummonerID = summonerDTO.ID
	// summoner.AccountID = summonerDTO.AccountID
	// summoner.Puuid = summonerDTO.Puuid
	// summoner.ProfileIconID = summonerDTO.ProfileIconID
	// summoner.RevisionDate = summonerDTO.RevisionDate
	return summonerDTO
}

func (riotAPIClient *RiotAPIClient) GetSummonerLeaguesByID(summonerID string) []LeagueEntryDTO {
	newRequest, errNewRequest := http.NewRequest("GET", riotAPIClient.ServerURL+leagueV4+summonerID, nil)
	if errNewRequest != nil {
		utils.LogOperation.Println(errNewRequest.Error())
	}
	newRequest.Header.Set(riotAPIClient.HeaderAPI, riotAPIClient.TokenAPI)
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	response, errResponse := client.Do(newRequest)
	if errResponse != nil {
		utils.LogOperation.Println(errResponse)
		// utils.LogOperation.Print("Failed perform request: " + endpoint + " - " + errResponse.Error())
	} else if response.StatusCode != 200 {
		// utils.LogOperation.Print("Failed get request with following info: '" + req.pathParam + "' in: '" + req.endpoint + "' - Invalid status code: " + response.Status)
	}
	var leagueEntryDTO []LeagueEntryDTO
	if errDecodeSummonerResponse := json.NewDecoder(response.Body).Decode(&leagueEntryDTO); errDecodeSummonerResponse != nil {
		utils.LogOperation.Println(errDecodeSummonerResponse)
	}
	defer response.Body.Close()
	return leagueEntryDTO
}

func (riotAPIClient *RiotAPIClient) GetSummonerMatchesByAccountID(accountID string) *MatchlistDto {
	newRequest, errNewRequest := http.NewRequest("GET", riotAPIClient.ServerURL+matchesV4+accountID, nil)
	if errNewRequest != nil {
		utils.LogOperation.Println(errNewRequest.Error())
	}
	newRequest.Header.Set(riotAPIClient.HeaderAPI, riotAPIClient.TokenAPI)
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	response, errResponse := client.Do(newRequest)
	if errResponse != nil {
		utils.LogOperation.Println(errResponse)
		// utils.LogOperation.Print("Failed perform request: " + endpoint + " - " + errResponse.Error())
	} else if response.StatusCode != 200 {
		// utils.LogOperation.Print("Failed get request with following info: '" + req.pathParam + "' in: '" + req.endpoint + "' - Invalid status code: " + response.Status)
	}
	matchlistDto := new(MatchlistDto)
	if errDecodeSummonerResponse := json.NewDecoder(response.Body).Decode(&matchlistDto); errDecodeSummonerResponse != nil {
		utils.LogOperation.Println(errDecodeSummonerResponse)
	}
	defer response.Body.Close()
	return matchlistDto
}

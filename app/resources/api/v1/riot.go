package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/antonioazambuja/ionia/utils"
)

const summonerV4 string = "/lol/summoner/v4/summoners/by-name/"
const leagueV4 string = "/lol/league/v4/entries/by-summoner/"
const matchesV4 string = "/lol/match/v4/matchlists/by-account/"

// RiotAPIClient - client of Riot API
type RiotAPIClient struct {
	ServerURL string
	TokenAPI  string
	HeaderAPI string
}

// NewRiotAPIClient - create new client Riot API
func NewRiotAPIClient(serverURL, tokenAPI, headerAPI string) *RiotAPIClient {
	return &RiotAPIClient{
		ServerURL: serverURL,
		TokenAPI:  tokenAPI,
		HeaderAPI: headerAPI,
	}
}

// GetSummonerByName - Get summoner by name: "/lol/summoner/v4/summoners/by-name/"
func (riotAPIClient *RiotAPIClient) GetSummonerByName(summonerName string) (*SummonerDTO, error) {
	newRequest, errNewRequest := http.NewRequest("GET", riotAPIClient.ServerURL+summonerV4+summonerName, nil)
	if errNewRequest != nil {
		utils.LogOperation.Println("Error found - GetSummonerByName: failed create new request")
		utils.LogOperation.Println(errNewRequest.Error())
		return nil, errNewRequest
	}
	newRequest.Header.Set(riotAPIClient.HeaderAPI, riotAPIClient.TokenAPI)
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	response, errResponse := client.Do(newRequest)
	if errResponse != nil {
		utils.LogOperation.Print("Error found - GetSummonerByName: " + summonerV4 + " - " + errResponse.Error())
		utils.LogOperation.Println(errResponse)
		return nil, errResponse
	} else if response.StatusCode != 200 {
		utils.LogOperation.Print("Error found: Response of '/lol/summoner/v4/summoners/by-name/' with invalid status code: '" + response.Status + "'")
		return nil, errors.New("Error found - Invalid status code: '" + response.Status + "'")
	}
	summonerDTO := new(SummonerDTO)
	if errDecodeSummonerResponse := json.NewDecoder(response.Body).Decode(&summonerDTO); errDecodeSummonerResponse != nil {
		utils.LogOperation.Println(errDecodeSummonerResponse)
		return nil, errDecodeSummonerResponse
	}
	defer response.Body.Close()
	return summonerDTO, nil
}

// GetSummonerLeaguesByID - Get leagues of summoner by ID: "/lol/league/v4/entries/by-summoner/"
func (riotAPIClient *RiotAPIClient) GetSummonerLeaguesByID(summonerID string) ([]LeagueEntryDTO, error) {
	newRequest, errNewRequest := http.NewRequest("GET", riotAPIClient.ServerURL+leagueV4+summonerID, nil)
	if errNewRequest != nil {
		utils.LogOperation.Println("Error found - GetSummonerLeaguesByID: failed create new request")
		utils.LogOperation.Println(errNewRequest.Error())
		return nil, errNewRequest
	}
	newRequest.Header.Set(riotAPIClient.HeaderAPI, riotAPIClient.TokenAPI)
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	response, errResponse := client.Do(newRequest)
	if errResponse != nil {
		utils.LogOperation.Print("Error found - GetSummonerLeaguesByID: " + summonerV4 + " - " + errResponse.Error())
		utils.LogOperation.Println(errResponse)
		return nil, errResponse
	} else if response.StatusCode != 200 {
		utils.LogOperation.Print("Error found: Response of '/lol/league/v4/entries/by-summoner/' with invalid status code: '" + response.Status + "'")
		return nil, errors.New("Error found - Invalid status code: '" + response.Status + "'")
	}
	var leagueEntryDTO []LeagueEntryDTO
	if errDecodeSummonerResponse := json.NewDecoder(response.Body).Decode(&leagueEntryDTO); errDecodeSummonerResponse != nil {
		utils.LogOperation.Println(errDecodeSummonerResponse)
	}
	defer response.Body.Close()
	return leagueEntryDTO, nil
}

// GetSummonerMatchesByAccountID - Get matches of summoner by account ID: "/lol/match/v4/matchlists/by-account/"
func (riotAPIClient *RiotAPIClient) GetSummonerMatchesByAccountID(accountID string) (*MatchlistDto, error) {
	newRequest, errNewRequest := http.NewRequest("GET", riotAPIClient.ServerURL+matchesV4+accountID, nil)
	if errNewRequest != nil {
		utils.LogOperation.Println("Error found - GetSummonerMatchesByAccountID: failed create new request")
		utils.LogOperation.Println(errNewRequest.Error())
		return nil, errNewRequest
	}
	newRequest.Header.Set(riotAPIClient.HeaderAPI, riotAPIClient.TokenAPI)
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	response, errResponse := client.Do(newRequest)
	if errResponse != nil {
		utils.LogOperation.Print("Error found - GetSummonerMatchesByAccountID: " + summonerV4 + " - " + errResponse.Error())
		utils.LogOperation.Println(errResponse)
		return nil, errResponse
	} else if response.StatusCode != 200 {
		utils.LogOperation.Print("Error found: Response of '/lol/match/v4/matchlists/by-account/' with invalid status code: '" + response.Status + "'")
		return nil, errors.New("Error found - Invalid status code: '" + response.Status + "'")
	}
	matchlistDto := new(MatchlistDto)
	if errDecodeSummonerResponse := json.NewDecoder(response.Body).Decode(&matchlistDto); errDecodeSummonerResponse != nil {
		utils.LogOperation.Println(errDecodeSummonerResponse)
	}
	defer response.Body.Close()
	return matchlistDto, nil
}

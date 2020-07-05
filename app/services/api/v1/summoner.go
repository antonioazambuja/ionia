package v1

import (
	"regexp"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	utils "github.com/antonioazambuja/ionia/utils"
)

// RiotAPIClientFunc - interface with functions of RiotAPIClient
type RiotAPIClientFunc interface {
	GetSummonerByName(summonerName string) (*rsc_v1.SummonerDTO, error)
	GetSummonerLeaguesByID(summonerID string) ([]rsc_v1.LeagueEntryDTO, error)
	GetSummonerMatchesByAccountID(accountID string) (*rsc_v1.MatchlistDto, error)
}

// SummonerV4 -
const SummonerV4 string = "/lol/summoner/v4/summoners/by-name/"

// LeagueV4 -
const LeagueV4 string = "/lol/league/v4/entries/by-summoner/"

// MatchesV4 -
const MatchesV4 string = "/lol/match/v4/matchlists/by-account/"

// GetByName - Service complex info summoner by name
func GetByName(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName)
		return nil, errCheckSummonerName
	}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	if errSummonerDTO != nil {
		return nil, errSummonerDTO
	}
	leagueEntryDTO, errLeagueEntryDTO := riotAPIClient.GetSummonerLeaguesByID(summonerDTO.ID)
	if errLeagueEntryDTO != nil {
		return nil, errLeagueEntryDTO
	}
	matchlistDTO, errMatchlistDTO := riotAPIClient.GetSummonerMatchesByAccountID(summonerDTO.AccountID)
	if errMatchlistDTO != nil {
		return nil, errMatchlistDTO
	}
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithLeagueInfo(leagueEntryDTO)
	summoner.WithMatchesInfo(matchlistDTO)
	return summoner, nil
}

// GetInfoByName - Service info summoner by name
func GetInfoByName(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName)
		return nil, errCheckSummonerName
	}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	if errSummonerDTO != nil {
		return nil, errSummonerDTO
	}
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	return summoner, nil
}

// GetLeagueByName - Service league info summoner by name
func GetLeagueByName(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName)
		return nil, errCheckSummonerName
	}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	if errSummonerDTO != nil {
		return nil, errSummonerDTO
	}
	leagueEntryDTO, errLeagueEntryDTO := riotAPIClient.GetSummonerLeaguesByID(summonerDTO.ID)
	if errLeagueEntryDTO != nil {
		return nil, errLeagueEntryDTO
	}
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithLeagueInfo(leagueEntryDTO)
	return summoner, nil
}

// GetMatchesByName - Service matches info summoner by name
func GetMatchesByName(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName)
		return nil, errCheckSummonerName
	}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	if errSummonerDTO != nil {
		return nil, errSummonerDTO
	}
	matchlistDTO, errMatchlistDTO := riotAPIClient.GetSummonerMatchesByAccountID(summonerDTO.AccountID)
	if errMatchlistDTO != nil {
		return nil, errMatchlistDTO
	}
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithMatchesInfo(matchlistDTO)
	return summoner, nil
}

// CheckSummonerName - Validate summoner name before perform get info in service
func CheckSummonerName(summonerName string) error {
	_, errEspecialCharacters := regexp.MatchString("[$&+,:;=?@#|'<>.^*()%!-]", summonerName)
	if errEspecialCharacters != nil {
		utils.LogOperation.Print("Error validate summoner name - " + summonerName)
		utils.LogOperation.Print(errEspecialCharacters.Error())
		return errEspecialCharacters
	}
	return nil
}

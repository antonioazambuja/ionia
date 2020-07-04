package v1

import (
	"regexp"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	utils "github.com/antonioazambuja/ionia/utils"
)

// RiotAPIClientFunc -
type RiotAPIClientFunc interface {
	GetSummonerByName(summonerName string) *rsc_v1.SummonerDTO
	GetSummonerLeaguesByID(summonerID string) []rsc_v1.LeagueEntryDTO
	GetSummonerMatchesByAccountID(accountID string) *rsc_v1.MatchlistDto
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
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	// summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "full")
	// if errSummonerCacheRedis == nil {
	// 	return &summonerCacheRedis, nil
	// }
	summonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	leagueEntryDTO := riotAPIClient.GetSummonerLeaguesByID(summonerDTO.ID)
	matchlistDto := riotAPIClient.GetSummonerMatchesByAccountID(summonerDTO.AccountID)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithLeagueInfo(leagueEntryDTO)
	summoner.WithMatchesInfo(matchlistDto)
	// if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "full"); errNewCacheSummoner != nil {
	// 	utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
	// 	utils.LogOperation.Print(errNewCacheSummoner.Error())
	// }
	return summoner, nil
}

// GetInfoByName - Service info summoner by name
func GetInfoByName(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName.Error())
		return nil, errCheckSummonerName
	}
	// summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "info")
	// if errSummonerCacheRedis == nil {
	// 	return &summonerCacheRedis, nil
	// }
	summonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	// if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "info"); errNewCacheSummoner != nil {
	// 	utils.LogOperation.Print("Error found! Failed service GetInfoByName: errNewCacheSummoner")
	// 	utils.LogOperation.Print(errNewCacheSummoner.Error())
	// }
	return summoner, nil
}

// GetLeagueByName - Service league info summoner by name
func GetLeagueByName(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	// summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "league")
	// if errSummonerCacheRedis == nil {
	// 	return &summonerCacheRedis, nil
	// }
	summonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	leagueEntryDTO := riotAPIClient.GetSummonerLeaguesByID(summonerDTO.ID)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithLeagueInfo(leagueEntryDTO)
	// if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "league"); errNewCacheSummoner != nil {
	// 	utils.LogOperation.Print("Error found! Failed service GetLeagueByName: errNewCacheSummoner")
	// 	utils.LogOperation.Print(errNewCacheSummoner.Error())
	// }
	return summoner, nil
}

// GetMatchesByName - Service matches info summoner by name
func GetMatchesByName(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	// summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "matches")
	// if errSummonerCacheRedis == nil {
	// 	return &summonerCacheRedis, nil
	// }
	summonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	matchlistDto := riotAPIClient.GetSummonerMatchesByAccountID(summonerDTO.AccountID)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithMatchesInfo(matchlistDto)
	// if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "full"); errNewCacheSummoner != nil {
	// 	utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
	// 	utils.LogOperation.Print(errNewCacheSummoner.Error())
	// }
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

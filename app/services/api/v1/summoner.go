package v1

import (
	"regexp"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	"github.com/antonioazambuja/ionia/utils"
)

const summonerV4 string = "/lol/summoner/v4/summoners/by-name/"

const leagueV4 string = "/lol/league/v4/entries/by-summoner/"

const matchesV4 string = "/lol/match/v4/matchlists/by-account/"

// GetInfoByName - Service info summoner by name
func GetInfoByName(summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName.Error())
		return nil, errCheckSummonerName
	}
	summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "info")
	if errSummonerCacheRedis == nil {
		return &summonerCacheRedis, nil
	}
	request := rsc_v1.NewRequestBuilder(summonerV4)
	request.WithPathParam(summonerName)
	summonerHTTPResponse, errSummonerHTTPResponse := request.Run()
	if errSummonerHTTPResponse != nil {
		utils.LogOperation.Println(errSummonerHTTPResponse.Error())
		return nil, errSummonerHTTPResponse
	}
	defer summonerHTTPResponse.Body.Close()
	summoner := rsc_v1.NewSummoner(summonerHTTPResponse)
	if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "info"); errNewCacheSummoner != nil {
		utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
		utils.LogOperation.Print(errNewCacheSummoner.Error())
	}
	return summoner, nil
}

// GetByName - Service complex info summoner by name
func GetByName(summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "full")
	if errSummonerCacheRedis == nil {
		return &summonerCacheRedis, nil
	}
	summonerHTTPResponse, errSummonerHTTPResponse := rsc_v1.NewRequestBuilder(summonerV4).WithPathParam(summonerName).Run()
	if errSummonerHTTPResponse != nil {
		utils.LogOperation.Println(errSummonerHTTPResponse.Error())
		return nil, errSummonerHTTPResponse
	}
	summoner := rsc_v1.NewSummoner(summonerHTTPResponse)
	summonerLeagueHTTPResponse, errorHTTPResponseLeague := rsc_v1.NewRequestBuilder(leagueV4).WithPathParam(summoner.SummonerID).Run()
	if errorHTTPResponseLeague != nil {
		utils.LogOperation.Println(errorHTTPResponseLeague.Error())
		return nil, errorHTTPResponseLeague
	}
	defer summonerHTTPResponse.Body.Close()
	summonerHTTPMatchesResponse, errorSummonerHTTPMatchesResponse := rsc_v1.NewRequestBuilder(matchesV4).WithPathParam(summoner.AccountID).WithQueries([]string{"beginIndex", "endIndex"}, []string{"0", "15"}).Run()
	if errorSummonerHTTPMatchesResponse != nil {
		utils.LogOperation.Print("Failed build summoner, get matches info")
		return nil, errorSummonerHTTPMatchesResponse
	}
	defer summonerHTTPMatchesResponse.Body.Close()
	summoner.WithLeagueInfo(summonerLeagueHTTPResponse)
	summoner.WithMatchesInfo(summonerHTTPMatchesResponse)
	if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "full"); errNewCacheSummoner != nil {
		utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
		utils.LogOperation.Print(errNewCacheSummoner.Error())
	}
	return summoner, nil
}

// GetLeagueByName - Service league info summoner by name
func GetLeagueByName(summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "league")
	if errSummonerCacheRedis == nil {
		return &summonerCacheRedis, nil
	}
	summonerHTTPResponse, errSummonerHTTPResponse := rsc_v1.NewRequestBuilder(summonerV4).WithPathParam(summonerName).Run()
	if errSummonerHTTPResponse != nil {
		utils.LogOperation.Println(errSummonerHTTPResponse.Error())
		return nil, errSummonerHTTPResponse
	}
	defer summonerHTTPResponse.Body.Close()
	summoner := rsc_v1.NewSummoner(summonerHTTPResponse)
	summonerLeagueHTTPResponse, errorHTTPResponseLeague := rsc_v1.NewRequestBuilder(leagueV4).WithPathParam(summoner.SummonerID).Run()
	if errorHTTPResponseLeague != nil {
		utils.LogOperation.Println(errorHTTPResponseLeague.Error())
		return nil, errorHTTPResponseLeague
	}
	defer summonerLeagueHTTPResponse.Body.Close()
	summoner.WithLeagueInfo(summonerLeagueHTTPResponse)
	if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "league"); errNewCacheSummoner != nil {
		utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
		utils.LogOperation.Print(errNewCacheSummoner.Error())
	}
	return summoner, nil
}

// GetMatchesByName - Service matches info summoner by name
func GetMatchesByName(summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "matches")
	if errSummonerCacheRedis == nil {
		return &summonerCacheRedis, nil
	}
	request := rsc_v1.NewRequestBuilder(summonerV4)
	request.WithPathParam(summonerName)
	summonerHTTPResponse, errSummonerHTTPResponse := request.Run()
	if errSummonerHTTPResponse != nil {
		utils.LogOperation.Println(errSummonerHTTPResponse.Error())
		return nil, errSummonerHTTPResponse
	}
	defer summonerHTTPResponse.Body.Close()
	summoner := rsc_v1.NewSummoner(summonerHTTPResponse)
	summonerMatchesHTTPResponse, errorMatchesHTTPResponse := rsc_v1.NewRequestBuilder(matchesV4).WithPathParam(summoner.AccountID).WithQueries([]string{"beginIndex", "endIndex"}, []string{"0", "15"}).Run()
	if errorMatchesHTTPResponse != nil {
		utils.LogOperation.Print(errorMatchesHTTPResponse.Error())
		return nil, errorMatchesHTTPResponse
	}
	defer summonerMatchesHTTPResponse.Body.Close()
	summoner.WithMatchesInfo(summonerMatchesHTTPResponse)
	if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "matches"); errNewCacheSummoner != nil {
		utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
		utils.LogOperation.Print(errNewCacheSummoner.Error())
	}
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

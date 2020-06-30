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
	if errCheckSummonerName := checkSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "info")
	if errSummonerCacheRedis == nil {
		return &summonerCacheRedis, nil
	}
	request := rsc_v1.NewRequestBuilder(summonerV4)
	request.WithPathParam(summonerName)
	summonerHTTPResponse, err := request.Run()
	if err != nil {
		panic(err)
	}
	summoner := rsc_v1.NewSummoner(summonerHTTPResponse)
	newSummoner := summoner
	if errNewCacheSummoner := rsc_v1.NewCacheSummoner(newSummoner, "info"); errNewCacheSummoner != nil {
		utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
		utils.LogOperation.Print(errNewCacheSummoner.Error())
	}
	return newSummoner, nil
}

// GetByName - Service complex info summoner by name
func GetByName(summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := checkSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "full")
	if errSummonerCacheRedis == nil {
		return &summonerCacheRedis, nil
	}
	summonerHTTPResponse, err := rsc_v1.NewRequestBuilder(summonerV4).WithPathParam(summonerName).Run()
	if err != nil {
		panic(err)
	}
	summoner := rsc_v1.NewSummoner(summonerHTTPResponse)
	summonerLeagueHTTPResponse, errorHTTPResponseLeague := rsc_v1.NewRequestBuilder(leagueV4).WithPathParam(summoner.SummonerID).Run()
	if errorHTTPResponseLeague != nil {
		panic(errorHTTPResponseLeague)
	}
	summonerHTTPMatchesResponse, errorSummonerHTTPMatchesResponse := rsc_v1.NewRequestBuilder(matchesV4).WithPathParam(summoner.AccountID).WithQueries([]string{"beginIndex", "endIndex"}, []string{"0", "15"}).Run()
	if errorSummonerHTTPMatchesResponse != nil {
		utils.LogOperation.Print("Failed build summoner, get matches info")
		return nil, errorSummonerHTTPMatchesResponse
	}
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
	if errCheckSummonerName := checkSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "league")
	if errSummonerCacheRedis == nil {
		return &summonerCacheRedis, nil
	}
	summonerHTTPResponse, err := rsc_v1.NewRequestBuilder(summonerV4).WithPathParam(summonerName).Run()
	if err != nil {
		panic(err)
	}
	summoner := rsc_v1.NewSummoner(summonerHTTPResponse)
	summonerLeagueHTTPResponse, errorHTTPResponseLeague := rsc_v1.NewRequestBuilder(leagueV4).WithPathParam(summoner.SummonerID).Run()
	if errorHTTPResponseLeague != nil {
		panic(errorHTTPResponseLeague)
	}
	summoner.WithLeagueInfo(summonerLeagueHTTPResponse)
	utils.LogOperation.Println(&summoner)
	if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "league"); errNewCacheSummoner != nil {
		utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
		utils.LogOperation.Print(errNewCacheSummoner.Error())
	}
	return summoner, nil
}

// GetMatchesByName - Service matches info summoner by name
func GetMatchesByName(summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := checkSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println("Not validate summoner name")
		return nil, errCheckSummonerName
	}
	summonerCacheRedis, errSummonerCacheRedis := rsc_v1.GetCacheSummoner(summonerName, "matches")
	if errSummonerCacheRedis == nil {
		return &summonerCacheRedis, nil
	}
	request := rsc_v1.NewRequestBuilder(summonerV4)
	request.WithPathParam(summonerName)
	summonerHTTPResponse, err := request.Run()
	if err != nil {
		panic(err)
	}
	summoner := rsc_v1.NewSummoner(summonerHTTPResponse)
	summonerMatchesHTTPResponse, errorMatchesHTTPResponse := rsc_v1.NewRequestBuilder(matchesV4).WithPathParam(summoner.AccountID).WithQueries([]string{"beginIndex", "endIndex"}, []string{"0", "15"}).Run()
	if errorMatchesHTTPResponse != nil {
		utils.LogOperation.Print("Failed build summoner, get matches info")
		utils.LogOperation.Print(errorMatchesHTTPResponse.Error())
		return nil, errorMatchesHTTPResponse
	}
	summoner.WithMatchesInfo(summonerMatchesHTTPResponse)
	if errNewCacheSummoner := rsc_v1.NewCacheSummoner(summoner, "matches"); errNewCacheSummoner != nil {
		utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
		utils.LogOperation.Print(errNewCacheSummoner.Error())
	}
	return summoner, nil
}

func checkSummonerName(summonerName string) error {
	_, errEspecialCharacters := regexp.MatchString("[$&+,:;=?@#|'<>.^*()%!-]", summonerName)
	if errEspecialCharacters != nil {
		utils.LogOperation.Print("Error validate summoner name - " + summonerName)
		utils.LogOperation.Print(errEspecialCharacters.Error())
		return errEspecialCharacters
	}
	return nil
}

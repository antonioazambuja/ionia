package v1

import (
	"log"
	"os"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
)

// GetByName - Service summoner by name
func GetByName(summonerName string) (rsc_v1.Summoner, error) {
	if checkSummonerName(summonerName) {
		summoner, errSummoner := rsc_v1.NewSummonerBuilder(summonerName).WithSummonerInfo().WithLeagueInfo().WithMatchesInfo().Build()
		if errSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Failed service GetByName")
			return rsc_v1.Summoner{}, errSummoner
		} else {
			return summoner, nil
		}
	} else {
		return rsc_v1.Summoner{}, errors.New("Not validate summoner name")
	}
}

// GetInfoByName - Service main info summoner by name
func GetInfoByName(summonerName string) (rsc_v1.Summoner, error) {
	summoner, errSummoner := rsc_v1.NewSummonerBuilder(summonerName).WithSummonerInfo().Build()
	if errSummoner != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed service GetInfoByName")
		return rsc_v1.Summoner{}, errSummoner
	}
	return summoner, nil
}

// GetLeagueByName - Service league info summoner by name
func GetLeagueByName(summonerName string) (rsc_v1.Summoner, error) {
	summoner, errSummoner := rsc_v1.NewSummonerBuilder(summonerName).WithLeagueInfo().Build()
	if errSummoner != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed service GetLeagueByName")
		return rsc_v1.Summoner{}, errSummoner
	}
	return summoner, nil
}

// GetMatchesByName - Service matches info summoner by name
func GetMatchesByName(summonerName string) (rsc_v1.Summoner, error) {
	summoner, errSummoner := rsc_v1.NewSummonerBuilder(summonerName).WithMatchesInfo().Build()
	if errSummoner != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed service GetMatchesByName")
		return rsc_v1.Summoner{}, errSummoner
	}
	return summoner, nil
}

func checkSummonerName(summonerName string) bool {
	checkSummonerName, errEspecialCharacters := regexp.MatchString("[$&+,:;=?@#|'<>.^*()%!-]", summonerName)
	if errEspecialCharacters != nil {
		utils.LogOperation.Print("Error validate summoner name - " + summonerName)
		return false
	}
	return !checkSummonerName
}

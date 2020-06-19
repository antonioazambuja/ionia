package v1

import (
	"errors"
	"log"
	"os"
	"regexp"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	"github.com/antonioazambuja/ionia/utils"
)

// GetByName - Service summoner by name
func GetByName(summonerName string) (rsc_v1.Summoner, error) {
	if checkSummonerName(summonerName) {
		summoner, errCreateNewSummoner := rsc_v1.NewSummonerBuilder(summonerName)
		if errCreateNewSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Error found! Failed service GetByName: errCreateNewSummoner")
			return rsc_v1.Summoner{}, errCreateNewSummoner
		}
		newSummoner, errBuildNewSummoner := summoner.WithSummonerInfo().WithLeagueInfo().WithMatchesInfo().Build()
		if errBuildNewSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Error found! Failed service GetByName: errBuildNewSummoner")
			return rsc_v1.Summoner{}, errBuildNewSummoner
		}
		return newSummoner, nil
	}
	return rsc_v1.Summoner{}, errors.New("Not validate summoner name")
}

// GetInfoByName - Service main info summoner by name
func GetInfoByName(summonerName string) (rsc_v1.Summoner, error) {
	if checkSummonerName(summonerName) {
		summoner, errCreateNewSummoner := rsc_v1.NewSummonerBuilder(summonerName)
		if errCreateNewSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Error found! Failed service GetInfoByName: errCreateNewSummoner")
			return rsc_v1.Summoner{}, errCreateNewSummoner
		}
		newSummoner, errBuildNewSummoner := summoner.WithSummonerInfo().Build()
		if errBuildNewSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Error found! Failed service GetInfoByName: errBuildNewSummoner")
			return rsc_v1.Summoner{}, errBuildNewSummoner
		}
		return newSummoner, nil
	} 
	return rsc_v1.Summoner{}, errors.New("Not validate summoner name")
}

// GetLeagueByName - Service league info summoner by name
func GetLeagueByName(summonerName string) (rsc_v1.Summoner, error) {
	if checkSummonerName(summonerName) {
		summoner, errCreateNewSummoner := rsc_v1.NewSummonerBuilder(summonerName)
		if errCreateNewSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Error found! Failed service GetLeagueByName: errCreateNewSummoner")
			return rsc_v1.Summoner{}, errCreateNewSummoner
		}
		newSummoner, errBuildNewSummoner := summoner.WithLeagueInfo().Build()
		if errBuildNewSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Error found! Failed service GetLeagueByName: errBuildNewSummoner")
			return rsc_v1.Summoner{}, errBuildNewSummoner
		}
		return newSummoner, nil
	}
	return rsc_v1.Summoner{}, errors.New("Not validate summoner name")
}

// GetMatchesByName - Service matches info summoner by name
func GetMatchesByName(summonerName string) (rsc_v1.Summoner, error) {
	if checkSummonerName(summonerName) {
		summoner, errCreateNewSummoner := rsc_v1.NewSummonerBuilder(summonerName)
		if errCreateNewSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Error found! Failed service GetMatchesByName: errCreateNewSummoner")
			return rsc_v1.Summoner{}, errCreateNewSummoner
		}
		newSummoner, errBuildNewSummoner := summoner.WithMatchesInfo().Build()
		if errBuildNewSummoner != nil {
			logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
			logOperation.Print("Error found! Failed service GetMatchesByName: errBuildNewSummoner")
			return rsc_v1.Summoner{}, errBuildNewSummoner
		}
		return newSummoner, nil
	}
	return rsc_v1.Summoner{}, errors.New("Not validate summoner name")
}

func checkSummonerName(summonerName string) bool {
	checkSummonerName, errEspecialCharacters := regexp.MatchString("[$&+,:;=?@#|'<>.^*()%!-]", summonerName)
	if errEspecialCharacters != nil {
		utils.LogOperation.Print("Error validate summoner name - " + summonerName)
		return false
	}
	return !checkSummonerName
}

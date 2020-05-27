package v1

import (
	"log"
	"os"

	rsc_v1 "github.com/antonioazambuja/ionia/resources/api/v1"
)

// GetByName - get summoner by name
func GetByName(summonerName string) (rsc_v1.Summoner, error) {
	summoner, errSummoner := rsc_v1.NewSummonerBuilder(summonerName).WithSummonerInfo().WithLeagueInfo().WithMatchesInfo().Build()
	if errSummoner != nil {
		logOperation := log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
		logOperation.Print("Failed builder summoner in service GetByName")
		return rsc_v1.Summoner{}, errSummoner
	}
	return summoner, nil
}

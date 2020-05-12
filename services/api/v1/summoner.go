package v1

import (
	"log"

	rsc_v1 "github.com/antonioazambuja/ionia/resources/api/v1"
)

// GetByName - get summoner by name
func GetByName(summonerName string) (rsc_v1.Summoner, error) {
	summoner, errSummoner := rsc_v1.NewSummonerBuilder(summonerName).WithSummonerInfo().WithLeagueInfo().Build()
	if errSummoner != nil {
		log.Fatal("Error service...")
		panic(errSummoner)
		// return rsc_v1.Summoner{}, errSummoner
	}
	return summoner, nil
}

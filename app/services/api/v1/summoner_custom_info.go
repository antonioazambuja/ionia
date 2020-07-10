package v1

import (
	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	utils "github.com/antonioazambuja/ionia/utils"
)

// GetMatchesCurrentYear - Service complex info summoner by name
func GetMatchesCurrentYear(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName)
		return nil, errCheckSummonerName
	}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	if errSummonerDTO != nil {
		return nil, errSummonerDTO
	}
	matchlistDTO, errMatchlistDTO := riotAPIClient.GetMatchesCurrentYear(summonerDTO.AccountID)
	if errMatchlistDTO != nil {
		return nil, errMatchlistDTO
	}
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithMatchesInfo(matchlistDTO)
	return summoner, nil
}

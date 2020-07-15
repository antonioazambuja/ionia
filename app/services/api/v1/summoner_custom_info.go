package v1

import (
	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	utils "github.com/antonioazambuja/ionia/utils"
)

// GetMatchesCurrentYear - Service complex info summoner by name
func GetMatchesCurrentYear(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	summoner := rsc_v1.NewSummoner()
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName)
		return &rsc_v1.Summoner{}, errCheckSummonerName
	}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	if errSummonerDTO != nil {
		return &rsc_v1.Summoner{}, errSummonerDTO
	}
	summoner.WithSummonerInfo(summonerDTO)
	matchlistDTO, errMatchlistDTO := riotAPIClient.GetMatchesCurrentYear(summonerDTO.AccountID)
	if errMatchlistDTO != nil {
		return summoner, errMatchlistDTO
	}
	summoner.WithMatchesInfo(matchlistDTO)
	return summoner, nil
}

// GetMatchesCurrentMonth - Service complex info summoner by name
func GetMatchesCurrentMonth(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	summoner := rsc_v1.NewSummoner()
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName)
		return &rsc_v1.Summoner{}, errCheckSummonerName
	}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	if errSummonerDTO != nil {
		return &rsc_v1.Summoner{}, errSummonerDTO
	}
	summoner.WithSummonerInfo(summonerDTO)
	matchlistDTO, errMatchlistDTO := riotAPIClient.GetMatchesCurrentMonth(summonerDTO.AccountID)
	if errMatchlistDTO != nil {
		return summoner, errMatchlistDTO
	}
	summoner.WithMatchesInfo(matchlistDTO)
	return summoner, nil
}

// GetMatchesCurrentDay - Service complex info summoner by name
func GetMatchesCurrentDay(riotAPIClient RiotAPIClientFunc, summonerName string) (*rsc_v1.Summoner, error) {
	summoner := rsc_v1.NewSummoner()
	if errCheckSummonerName := CheckSummonerName(summonerName); errCheckSummonerName != nil {
		utils.LogOperation.Println(errCheckSummonerName)
		return &rsc_v1.Summoner{}, errCheckSummonerName
	}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	if errSummonerDTO != nil {
		return &rsc_v1.Summoner{}, errSummonerDTO
	}
	summoner.WithSummonerInfo(summonerDTO)
	matchlistDTO, errMatchlistDTO := riotAPIClient.GetMatchesCurrentDay(summonerDTO.AccountID)
	if errMatchlistDTO != nil {
		return summoner, errMatchlistDTO
	}
	summoner.WithMatchesInfo(matchlistDTO)
	return summoner, nil
}

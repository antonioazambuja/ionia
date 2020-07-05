package v1

import (
	"testing"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestRiotAPIGetSummonerByName(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	summonerName := "IsBlackPanther"
	riotAPIClient := mockRiotClient{}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	assert.Nil(errSummonerDTO)
	leagueEntryDTO, errLeagueEntryDTO := riotAPIClient.GetSummonerLeaguesByID(summonerDTO.ID)
	assert.Nil(errLeagueEntryDTO)
	matchlistDTO, errMatchlistDTO := riotAPIClient.GetSummonerMatchesByAccountID(summonerDTO.AccountID)
	assert.Nil(errMatchlistDTO)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithLeagueInfo(leagueEntryDTO)
	summoner.WithMatchesInfo(matchlistDTO)
	assert.Equal("IsBlackPanther", summoner.SummonerName)
}

func TestRiotAPIGetInfoSummonerByName(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	summonerName := "IsBlackPanther"
	riotAPIClient := mockRiotClient{}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	assert.Nil(errSummonerDTO)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	assert.Equal("IsBlackPanther", summoner.SummonerName)
}

func TestRiotAPIGetLeagueSummonerByName(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	summonerName := "IsBlackPanther"
	riotAPIClient := mockRiotClient{}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	assert.Nil(errSummonerDTO)
	leagueEntryDTO, errLeagueEntryDTO := riotAPIClient.GetSummonerLeaguesByID(summonerDTO.ID)
	assert.Nil(errLeagueEntryDTO)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithLeagueInfo(leagueEntryDTO)
	assert.Equal("IsBlackPanther", summoner.SummonerName)
}

func TestRiotAPIGetMatchesSummonerByName(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	summonerName := "IsBlackPanther"
	riotAPIClient := mockRiotClient{}
	summonerDTO, errSummonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	assert.Nil(errSummonerDTO)
	matchlistDTO, errMatchlistDTO := riotAPIClient.GetSummonerMatchesByAccountID(summonerDTO.AccountID)
	assert.Nil(errMatchlistDTO)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithMatchesInfo(matchlistDTO)
	assert.Equal("IsBlackPanther", summoner.SummonerName)
}

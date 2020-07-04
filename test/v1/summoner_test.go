package v1

import (
	"testing"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	"github.com/stretchr/testify/assert"
)

type mockRiotClient struct{}

func (riotAPIClient *mockRiotClient) GetSummonerByName(summonerName string) *rsc_v1.SummonerDTO {
	return &rsc_v1.SummonerDTO{
		AccountID: "",
		Name:      "IsBlackPanther",
	}
}

func (riotAPIClient *mockRiotClient) GetSummonerLeaguesByID(summonerID string) []rsc_v1.LeagueEntryDTO {
	var leagueEntryDTO []rsc_v1.LeagueEntryDTO
	return leagueEntryDTO
}

func (riotAPIClient *mockRiotClient) GetSummonerMatchesByAccountID(accountID string) *rsc_v1.MatchlistDto {
	return &rsc_v1.MatchlistDto{
		TotalGames: 190,
	}
}

func TestGetSummonerByName(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	summonerName := "IsBlackPanther"
	riotAPIClient := mockRiotClient{}
	summonerDTO := riotAPIClient.GetSummonerByName(summonerName)
	leagueEntryDTO := riotAPIClient.GetSummonerLeaguesByID(summonerDTO.ID)
	matchlistDto := riotAPIClient.GetSummonerMatchesByAccountID(summonerDTO.AccountID)
	summoner := new(rsc_v1.Summoner)
	summoner.WithSummonerInfo(summonerDTO)
	summoner.WithLeagueInfo(leagueEntryDTO)
	summoner.WithMatchesInfo(matchlistDto)
	assert.Equal("IsBlackPanther", summoner.SummonerName)
}

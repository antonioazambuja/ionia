package v1

import (
	"testing"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	"github.com/stretchr/testify/assert"
)

type mockRedisClient struct{}

func (mockRedisClient *mockRedisClient) SaveSummoner(summoner *rsc_v1.Summoner, informationID string) error {
	return nil
}

func (mockRedisClient *mockRedisClient) SearchSummoner(summonerID string, informationID string) (*rsc_v1.Summoner, error) {
	return &rsc_v1.Summoner{
		SummonerName: "IsBlackPanther",
	}, nil
}

type mockRiotClient struct{}

func (riotAPIClient *mockRiotClient) GetSummonerByName(summonerName string) (*rsc_v1.SummonerDTO, error) {
	return &rsc_v1.SummonerDTO{
		AccountID: "",
		Name:      "IsBlackPanther",
	}, nil
}

func (riotAPIClient *mockRiotClient) GetSummonerLeaguesByID(summonerID string) ([]rsc_v1.LeagueEntryDTO, error) {
	var leagueEntryDTO []rsc_v1.LeagueEntryDTO
	return leagueEntryDTO, nil
}

func (riotAPIClient *mockRiotClient) GetSummonerMatchesByAccountID(accountID string) (*rsc_v1.MatchlistDto, error) {
	return &rsc_v1.MatchlistDto{
		TotalGames: 190,
	}, nil
}

func TestGetSummonerByName(test *testing.T) {
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
	redisClient := mockRedisClient{}
	summonerCache, errSummonerCache := redisClient.SearchSummoner("IsBlackPanther", "info")
	assert.Nil(errSummonerCache)
	assert.Equal("IsBlackPanther", summonerCache.SummonerName)
}

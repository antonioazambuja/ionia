package v1

import (
	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
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

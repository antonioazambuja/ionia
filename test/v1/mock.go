package v1

import (
	"errors"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
)

type mockRedisClient struct{}

func (mockRedisClient *mockRedisClient) SaveSummoner(summoner *rsc_v1.Summoner, informationID string) error {
	return nil
}

func (mockRedisClient *mockRedisClient) SearchSummoner(summonerID string, informationID string) (*rsc_v1.Summoner, error) {
	switch informationID {
	case "full":
		return &rsc_v1.Summoner{
			SummonerName:  "IsBlackPanther",
			AccountID:     "rsvKSJAKKjeis90",
			ProfileIconID: 1,
			Puuid:         "1",
			RevisionDate:  12122345,
			SummonerID:    "90opdlsjdnJAKdso",
			SummonerLevel: 1,
			MatchesInfo: []rsc_v1.MatchReferenceDto{
				rsc_v1.MatchReferenceDto{
					GameID:     1995205271,
					Role:       "SOLO",
					Season:     13,
					PlatformID: "BR1",
					Champion:   39,
					Queue:      440,
					Lane:       "MID",
					Timestamp:  1594174353091,
				},
				rsc_v1.MatchReferenceDto{
					GameID:     1995180379,
					Role:       "SOLO",
					Season:     13,
					PlatformID: "BR1",
					Champion:   39,
					Queue:      440,
					Lane:       "TOP",
					Timestamp:  1594171998098,
				},
			},
			TotalGames: 2,
			LeagueInfo: []rsc_v1.LeagueInfo{
				rsc_v1.LeagueInfo{
					LeagueID:      "b80abef5-2b55-4fc6-ba48-88331795255c",
					QueueType:     "RANKED_SOLO_5x5",
					Tier:          "PLATINUM",
					Rank:          "II",
					LeaguePoints:  12,
					Wins:          17,
					Losses:        24,
					MiniSeriesDTO: rsc_v1.MiniSeriesDTO{},
				},
				rsc_v1.LeagueInfo{
					LeagueID:      "ca46fd88-db98-4719-a06b-239e44c1b9c2",
					QueueType:     "RANKED_FLEX_SR",
					Tier:          "GOLD",
					Rank:          "II",
					LeaguePoints:  13,
					Wins:          23,
					Losses:        11,
					FreshBlood:    true,
					MiniSeriesDTO: rsc_v1.MiniSeriesDTO{},
				},
			},
		}, nil
	case "info":
		return &rsc_v1.Summoner{
			SummonerName:  "IsBlackPanther",
			AccountID:     "rsvKSJAKKjeis90",
			ProfileIconID: 1,
			Puuid:         "1",
			RevisionDate:  12122345,
			SummonerID:    "90opdlsjdnJAKdso",
			SummonerLevel: 1,
		}, nil
	case "league":
		return &rsc_v1.Summoner{
			SummonerName:  "IsBlackPanther",
			AccountID:     "rsvKSJAKKjeis90",
			ProfileIconID: 1,
			Puuid:         "1",
			RevisionDate:  12122345,
			SummonerID:    "90opdlsjdnJAKdso",
			SummonerLevel: 1,
			LeagueInfo: []rsc_v1.LeagueInfo{
				rsc_v1.LeagueInfo{
					LeagueID:      "b80abef5-2b55-4fc6-ba48-88331795255c",
					QueueType:     "RANKED_SOLO_5x5",
					Tier:          "PLATINUM",
					Rank:          "II",
					LeaguePoints:  12,
					Wins:          17,
					Losses:        24,
					MiniSeriesDTO: rsc_v1.MiniSeriesDTO{},
				},
				rsc_v1.LeagueInfo{
					LeagueID:      "ca46fd88-db98-4719-a06b-239e44c1b9c2",
					QueueType:     "RANKED_FLEX_SR",
					Tier:          "GOLD",
					Rank:          "II",
					LeaguePoints:  13,
					Wins:          23,
					Losses:        11,
					FreshBlood:    true,
					MiniSeriesDTO: rsc_v1.MiniSeriesDTO{},
				},
			},
		}, nil
	case "matches":
		return &rsc_v1.Summoner{
			SummonerName:  "IsBlackPanther",
			AccountID:     "rsvKSJAKKjeis90",
			ProfileIconID: 1,
			Puuid:         "1",
			RevisionDate:  12122345,
			SummonerID:    "90opdlsjdnJAKdso",
			SummonerLevel: 1,
			MatchesInfo: []rsc_v1.MatchReferenceDto{
				rsc_v1.MatchReferenceDto{
					GameID:     1995205271,
					Role:       "SOLO",
					Season:     13,
					PlatformID: "BR1",
					Champion:   39,
					Queue:      440,
					Lane:       "MID",
					Timestamp:  1594174353091,
				},
				rsc_v1.MatchReferenceDto{
					GameID:     1995180379,
					Role:       "SOLO",
					Season:     13,
					PlatformID: "BR1",
					Champion:   39,
					Queue:      440,
					Lane:       "TOP",
					Timestamp:  1594171998098,
				},
			},
			TotalGames: 2,
		}, nil
	default:
		return nil, errors.New("Not found summoner in search")
	}
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

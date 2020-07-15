package v1

import (
	"testing"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestRedisSearchFullSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summonerCache, errSummonerCache := redisClient.SearchSummoner("IsBlackPanther", "full")
	assert.Nil(errSummonerCache)
	assert.Equal("IsBlackPanther", summonerCache.SummonerName)
	assert.Equal("rsvKSJAKKjeis90", summonerCache.AccountID)
	assert.Equal(1, summonerCache.ProfileIconID)
	assert.Equal("1", summonerCache.Puuid)
	assert.Equal(12122345, summonerCache.RevisionDate)
	assert.Equal("90opdlsjdnJAKdso", summonerCache.SummonerID)
	assert.Equal(1, summonerCache.SummonerLevel)
	assert.Len(summonerCache.MatchesInfo, summonerCache.TotalGames)
	assert.Equal(1995205271, summonerCache.MatchesInfo[0].GameID)
	assert.Equal("SOLO", summonerCache.MatchesInfo[0].Role)
	assert.Equal(13, summonerCache.MatchesInfo[0].Season)
	assert.Equal("BR1", summonerCache.MatchesInfo[0].PlatformID)
	assert.Equal(39, summonerCache.MatchesInfo[0].Champion)
	assert.Equal(440, summonerCache.MatchesInfo[0].Queue)
	assert.Equal("MID", summonerCache.MatchesInfo[0].Lane)
	assert.Equal(1594174353091, summonerCache.MatchesInfo[0].Timestamp)
	assert.Equal(2, summonerCache.TotalGames)
	assert.Len(summonerCache.MatchesInfo, 2)
	assert.Equal("b80abef5-2b55-4fc6-ba48-88331795255c", summonerCache.LeagueInfo[0].LeagueID)
	assert.Equal("RANKED_SOLO_5x5", summonerCache.LeagueInfo[0].QueueType)
	assert.Equal("PLATINUM", summonerCache.LeagueInfo[0].Tier)
	assert.Equal("II", summonerCache.LeagueInfo[0].Rank)
	assert.Equal(12, summonerCache.LeagueInfo[0].LeaguePoints)
	assert.Equal(17, summonerCache.LeagueInfo[0].Wins)
	assert.Equal(24, summonerCache.LeagueInfo[0].Losses)
	assert.Equal(rsc_v1.MiniSeriesDTO{}, summonerCache.LeagueInfo[0].MiniSeriesDTO)
}

func TestRedisSearchInfoSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summonerCache, errSummonerCache := redisClient.SearchSummoner("IsBlackPanther", "info")
	assert.Nil(errSummonerCache)
	assert.Equal("IsBlackPanther", summonerCache.SummonerName)
	assert.Equal("rsvKSJAKKjeis90", summonerCache.AccountID)
	assert.Equal(1, summonerCache.ProfileIconID)
	assert.Equal("1", summonerCache.Puuid)
	assert.Equal(12122345, summonerCache.RevisionDate)
	assert.Equal("90opdlsjdnJAKdso", summonerCache.SummonerID)
}

func TestRedisSearchLeagueSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summonerCache, errSummonerCache := redisClient.SearchSummoner("IsBlackPanther", "league")
	assert.Nil(errSummonerCache)
	assert.Equal("IsBlackPanther", summonerCache.SummonerName)
	assert.Equal("rsvKSJAKKjeis90", summonerCache.AccountID)
	assert.Equal(1, summonerCache.ProfileIconID)
	assert.Equal("1", summonerCache.Puuid)
	assert.Equal(12122345, summonerCache.RevisionDate)
	assert.Equal("90opdlsjdnJAKdso", summonerCache.SummonerID)
	assert.Equal(1, summonerCache.SummonerLevel)
	assert.Equal("b80abef5-2b55-4fc6-ba48-88331795255c", summonerCache.LeagueInfo[0].LeagueID)
	assert.Equal("RANKED_SOLO_5x5", summonerCache.LeagueInfo[0].QueueType)
	assert.Equal("PLATINUM", summonerCache.LeagueInfo[0].Tier)
	assert.Equal("II", summonerCache.LeagueInfo[0].Rank)
	assert.Equal(12, summonerCache.LeagueInfo[0].LeaguePoints)
	assert.Equal(17, summonerCache.LeagueInfo[0].Wins)
	assert.Equal(24, summonerCache.LeagueInfo[0].Losses)
	assert.Equal(rsc_v1.MiniSeriesDTO{}, summonerCache.LeagueInfo[0].MiniSeriesDTO)
}

func TestRedisSearchMatchesSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summonerCache, errSummonerCache := redisClient.SearchSummoner("IsBlackPanther", "matches")
	assert.Nil(errSummonerCache)
	assert.Equal("IsBlackPanther", summonerCache.SummonerName)
	assert.Equal("rsvKSJAKKjeis90", summonerCache.AccountID)
	assert.Equal(1, summonerCache.ProfileIconID)
	assert.Equal("1", summonerCache.Puuid)
	assert.Equal(12122345, summonerCache.RevisionDate)
	assert.Equal("90opdlsjdnJAKdso", summonerCache.SummonerID)
	assert.Equal(1, summonerCache.SummonerLevel)
	assert.Len(summonerCache.MatchesInfo, summonerCache.TotalGames)
	assert.Equal(1995205271, summonerCache.MatchesInfo[0].GameID)
	assert.Equal("SOLO", summonerCache.MatchesInfo[0].Role)
	assert.Equal(13, summonerCache.MatchesInfo[0].Season)
	assert.Equal("BR1", summonerCache.MatchesInfo[0].PlatformID)
	assert.Equal(39, summonerCache.MatchesInfo[0].Champion)
	assert.Equal(440, summonerCache.MatchesInfo[0].Queue)
	assert.Equal("MID", summonerCache.MatchesInfo[0].Lane)
	assert.Equal(1594174353091, summonerCache.MatchesInfo[0].Timestamp)
	assert.Equal(2, summonerCache.TotalGames)
	assert.Len(summonerCache.MatchesInfo, 2)
}

func TestRedisSaveInfoSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summoner := &rsc_v1.Summoner{}
	errSaveSummonerCache := redisClient.SaveSummoner(summoner, "info")
	assert.Nil(errSaveSummonerCache)
}

func TestRedisSaveErrorSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	_, errSaveSummonerCache := redisClient.SearchSummoner("IsBlackPanther", "")
	assert.NotNil(errSaveSummonerCache)
}

func TestRedisSaveFullSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summoner := &rsc_v1.Summoner{}
	errSaveSummonerCache := redisClient.SaveSummoner(summoner, "full")
	assert.Nil(errSaveSummonerCache)
}

func TestRedisSaveLeagueSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summoner := &rsc_v1.Summoner{}
	errSaveSummonerCache := redisClient.SaveSummoner(summoner, "league")
	assert.Nil(errSaveSummonerCache)
}

func TestRedisSaveMatchesSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summoner := &rsc_v1.Summoner{}
	errSaveSummonerCache := redisClient.SaveSummoner(summoner, "matches")
	assert.Nil(errSaveSummonerCache)
}

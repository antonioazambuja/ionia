package v1

import (
	"testing"

	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestRedisSearchSummonerByName(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summonerCache, errSummonerCache := redisClient.SearchSummoner("IsBlackPanther", "info")
	assert.Nil(errSummonerCache)
	assert.Equal("IsBlackPanther", summonerCache.SummonerName)
	errSaveSummonerCache := redisClient.SaveSummoner(summonerCache, "info")
	assert.Nil(errSaveSummonerCache)
}

func TestRedisSaveSummoner(test *testing.T) {
	test.Parallel()
	assert := assert.New(test)
	redisClient := mockRedisClient{}
	summoner := &rsc_v1.Summoner{}
	errSaveSummonerCache := redisClient.SaveSummoner(summoner, "info")
	assert.Nil(errSaveSummonerCache)
}

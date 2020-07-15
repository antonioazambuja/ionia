package v1

import (
	rsc_v1 "github.com/antonioazambuja/ionia/app/resources/api/v1"
	utils "github.com/antonioazambuja/ionia/utils"
)

// RedisClientFunc - interface with functions of RedisClient
type RedisClientFunc interface {
	SaveSummoner(summoner *rsc_v1.Summoner, informationID string) error
	SearchSummoner(summonerID string, informationID string) (*rsc_v1.Summoner, error)
}

// GetRedisSummoner - Service complex info summoner by name
func GetRedisSummoner(redisClient RedisClientFunc, summonerName, informationID string) (*rsc_v1.Summoner, error) {
	summonerCacheRedis, errSummonerCacheRedis := redisClient.SearchSummoner(summonerName, informationID)
	if errSummonerCacheRedis != nil {
		return &rsc_v1.Summoner{}, errSummonerCacheRedis
	}
	return summonerCacheRedis, nil
}

// SaveRedisSummoner - Service info summoner by name
func SaveRedisSummoner(redisClient RedisClientFunc, summoner *rsc_v1.Summoner, informationID string) error {
	if errNewCacheSummoner := redisClient.SaveSummoner(summoner, informationID); errNewCacheSummoner != nil {
		utils.LogOperation.Print("Error found! Failed service GetByName: errNewCacheSummoner")
		utils.LogOperation.Print(errNewCacheSummoner.Error())
		return errNewCacheSummoner
	}
	return nil
}

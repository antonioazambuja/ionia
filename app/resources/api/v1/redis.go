package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/antonioazambuja/ionia/utils"

	"github.com/go-redis/redis"
)

// var redisClient *redis.Client

// RedisClient - Client for Redis
type RedisClient struct {
	client *redis.Client
	URI    string
}

// RedisClientConnected -
var RedisClientConnected *RedisClient

// NewRedisClient - New client for Redis
func NewRedisClient(host, port string) *RedisClient {
	return &RedisClient{
		URI:    host + ":" + port,
		client: nil,
	}
}

// Connect - Connect client Redis
func (redisClient *RedisClient) Connect(pwd string) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisClient.URI,
		Password: pwd,
		DB:       0,
	})
	pingResult, errPingResult := client.Ping(context.Background()).Result()
	if errPingResult != nil {
		utils.LogOperation.Println("Error found pingResultError: Failed connect in Redis host: " + os.Getenv("REDIS_MASTER_URL") + ".")
		// panic(errPingResult)
	}
	utils.LogOperation.Println("Connected in Redis: " + pingResult)
	redisClient.client = client
}

// GetConn - Get Redis connection with client Redis
func (redisClient *RedisClient) GetConn() *redis.Client {
	_, errPingResult := redisClient.client.Ping(context.Background()).Result()
	if errPingResult != nil {
		utils.LogOperation.Println("Error found pingResultError: Failed connect in Redis host: " + os.Getenv("REDIS_MASTER_URL") + ".")
		// redisClient.Connect()
	}
	return redisClient.client
}

// CloseRedisConn - Close client Redis
func (redisClient *RedisClient) CloseRedisConn() {
	if errCloseConn := redisClient.client.Close(); errCloseConn != nil {
		utils.LogOperation.Println("Error found: " + errCloseConn.Error())
		// panic(errCloseConn)
	}
	utils.LogOperation.Println("Closed connection with Redis")
}

// SaveSummoner - Save summoner informations in Redis
func (redisClient *RedisClient) SaveSummoner(summoner *Summoner, informationID string) error {
	summonerRedisKey := fmt.Sprint(summoner.SummonerName + "_" + informationID)
	summonerJSON, errParseStructToJSON := json.Marshal(summoner)
	if errParseStructToJSON != nil {
		utils.LogOperation.Println("Failed cached summoner of id service: " + informationID)
		utils.LogOperation.Println(errParseStructToJSON.Error())
		return errParseStructToJSON
	}
	setSummonerRedisResult, errSetSummonerRedisResult := redisClient.client.Do(context.TODO(), "SET", summonerRedisKey, summonerJSON).Result()
	if errSetSummonerRedisResult != nil {
		utils.LogOperation.Println("Failed cached summoner of id service: " + informationID)
		utils.LogOperation.Println(errSetSummonerRedisResult.Error())
		return errSetSummonerRedisResult
	}
	utils.LogOperation.Printf("Succesfull cached summoner of id service: %s. Result Redis: %s\n", summonerRedisKey, setSummonerRedisResult)
	setExpireSummonerRedisResult, errSetExpireSummonerRedisResult := redisClient.client.Expire(context.TODO(), summonerRedisKey, time.Duration(86400*time.Second)).Result()
	if errSetExpireSummonerRedisResult != nil && setExpireSummonerRedisResult {
		utils.LogOperation.Println("Failed cached summoner of id service: " + informationID)
		utils.LogOperation.Println(errSetSummonerRedisResult.Error())
		return errSetSummonerRedisResult
	}
	utils.LogOperation.Printf("Succesfull set TTL to cached summoner of id service: %s. Result Redis: %s\n", summonerRedisKey, setSummonerRedisResult)
	return nil
}

// SearchSummoner - Search summoner informations in Redis
func (redisClient *RedisClient) SearchSummoner(summonerName string, informationID string) (*Summoner, error) {
	var summoner *Summoner
	summonerCacheRedis, errGetSummonerCacheRedis := redisClient.client.Get(context.TODO(), summonerName+"_"+informationID).Result()
	if errGetSummonerCacheRedis != nil {
		utils.LogOperation.Print("Not found cache data in Redis - errGetSummonerCacheRedis. Result Redis: " + errGetSummonerCacheRedis.Error())
		return &Summoner{}, errGetSummonerCacheRedis
	}
	errParseJSONToStruct := json.Unmarshal([]byte(summonerCacheRedis), &summoner)
	if errParseJSONToStruct != nil {
		utils.LogOperation.Println("Error found cache data in Redis - errParseJSONToStruct. Result Redis: " + errParseJSONToStruct.Error())
		return &Summoner{}, errParseJSONToStruct
	}
	utils.LogOperation.Print("Found cache data in Redis")
	return summoner, nil
}

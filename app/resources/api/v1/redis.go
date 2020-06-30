package v1

import (
	"context"
	"os"

	"github.com/antonioazambuja/ionia/utils"

	"github.com/go-redis/redis"
)

var redisClientMaster *redis.Client

// CreateRedisConnection - create new connection with Redis
func CreateRedisConnection() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_MASTER_URL") + ":6379",
		Password: os.Getenv("REDIS_MASTER_PWD"),
		DB: 0,
	})
	pingResult, errPingResult := client.Ping(context.Background()).Result()
	if errPingResult != nil {
		utils.LogOperation.Println("Error found pingResultError: Failed connect in Redis host: " + os.Getenv("REDIS_MASTER_URL") + ".")
		panic(errPingResult)
	} else {
		utils.LogOperation.Println("Connected in Redis: " + pingResult)
	}
	redisClientMaster = client
}

// GetConn - test connection with Redis, if failed create new connection
func GetConn() *redis.Client {
	if pingResult := redisClientMaster.Ping(context.TODO()); pingResult.Err() != nil {
		CreateRedisConnection()
	}
	return redisClientMaster
}

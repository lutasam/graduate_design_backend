package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/lutasam/doctors/biz/utils"
)

var redisDB *redis.Client

func init() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     utils.GetConfigString("redis.address"),
		Password: utils.GetConfigString("redis.password"),
	})
	_, err := redisDB.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return redisDB
}

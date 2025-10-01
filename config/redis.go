package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Redis *redis.Client
var Ctx = context.Background()

func InitRedis() {
	addr := fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT"))

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: viper.GetString("REDIS_USER"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB"),
	})

	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}
	log.Println("✅ Redis connected")
}

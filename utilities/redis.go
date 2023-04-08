package utilities

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *Redis

type Redis struct {
	Client *redis.Client
}

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	RDB = &Redis{Client: client}
}

func (r *Redis) Set(key, value string, exp time.Duration) error {
	err := r.Client.Set(context.Background(), key, value, exp).Err()
	return err
}

func (r *Redis) Get(key string) (string, error) {
	val, err := r.Client.Get(context.Background(), key).Result()
	return val, err
}

package cache

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type redisClient struct {
	rdb *redis.Client
}

func InitRedis() Redis {

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err.Error())
	}

	return &redisClient{
		rdb: rdb,
	}

}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}) error {

	valueJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.rdb.Set(ctx, key, valueJSON, 10*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil

}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {

	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil

}

func (r *redisClient) Delete(ctx context.Context, key string) error {

	err := r.rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

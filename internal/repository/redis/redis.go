package redis

import (
	"api-orderfood/internal/model"
	"context"
	"github.com/go-redis/redis"
)

type RedisRepo struct {
	Redis *redis.Client
	CTX   context.Context
}

func NewRedisRepo(r *redis.Client, ctx context.Context) model.Cache {
	return &RedisRepo{Redis: r, CTX: ctx}
}

func (r *RedisRepo) Set(key string, value string) error {
	err := r.Redis.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisRepo) Get(key string) (string, error) {
	val, err := r.Redis.Get(key).Result()

	if err != nil {
		return "", err
	}
	return val, nil
}

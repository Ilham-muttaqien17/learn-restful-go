package config

import (
	"fmt"
	"runtime"

	"github.com/gofiber/storage/redis/v3"
)

var RedisStore *redis.Storage 


func RegisterRedis() error {

	if RedisStore != nil {
		return fmt.Errorf("redis connection is already initialized")
	}

	store := redis.New(redis.Config{
		Host: Env.RedisHost,
		Port: Env.RedisPort,
		Username: Env.RedisUsername,
		Password: Env.RedisPassword,
		Database: 0,
		TLSConfig: nil,
		PoolSize: 10 * runtime.GOMAXPROCS(0),
	})

	if store == nil {
		return fmt.Errorf("failed to initialize redis storage")
	}

	RedisStore = store

	return nil
}

func CloseRedis() error {

	if RedisStore == nil {
		return fmt.Errorf("redis connection is already closed")
	}

	if RedisStore != nil {
		if err := RedisStore.Close(); err != nil {
			return err
		}
	}

	return nil
}
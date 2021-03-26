package Common

import (
	"github.com/go-redis/redis"
	"fmt"
	"log"
	"time"
)

var RedisClient *redis.Client

// Redis cache implement
type Redis struct {
	Client *redis.Client
}

// Setup connection
func RedisConnect(cfg *Config) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.PassWord,
		DB:       cfg.Redis.DB,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not connected to redis : %s", err.Error()))
	}
	return err
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.Client.Get(key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val string, expire int) error {
	return r.Client.Set(key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.Client.Del(key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.Client.HGet(hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.Client.HDel(hk, key).Err()
}

// Increase
func (r *Redis) Increase(key string) error {
	return r.Client.Incr(key).Err()
}

// Set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.Client.Expire(key, dur).Err()
}

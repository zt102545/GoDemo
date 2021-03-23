package Common

import (
	"github.com/go-redis/redis"
	"fmt"
	"log"
	"time"
)

// Redis cache implement
type Redis struct {
	client *redis.Client
}

// Setup connection
func (r *Redis) Connect(cfg *Config) {
	r.client = redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Addr,
		Password:     cfg.Redis.PassWord,
		DB:           cfg.Redis.DB,
	})
	_, err := r.client.Ping().Result()
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not connected to redis : %s", err.Error()))
	}
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val string, expire int) error {
	return r.client.Set(key, val, time.Duration(expire)).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.client.Del(key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.client.HGet(hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.client.HDel(hk, key).Err()
}

// Increase
func (r *Redis) Increase(key string) error {
	return r.client.Incr(key).Err()
}

// Set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.client.Expire(key, dur).Err()
}

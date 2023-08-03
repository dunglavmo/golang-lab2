package repositories

import (
	"context"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(redisAddr string) (*RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &RedisRepository{
		client: client,
	}, nil
}

func (r *RedisRepository) IsRateLimited(ip string, limit int) bool {
	key := "rate_limit:" + ip

	// Increment the counter for the IP
	currentCount, err := r.client.IncrBy(ctx, key, 1).Result()
	if err != nil {
		log.Printf("Error incrementing Redis key: %v", err)
		return true
	}

	// Set the TTL to 30 seconds if the key is newly created
	if currentCount == 1 {
		_, err := r.client.Expire(ctx, key, 30*time.Second).Result()
		if err != nil {
			log.Printf("Error setting Redis key expiration: %v", err)
			return true
		}
	}
	log.Println(currentCount)

	// Check if the request count exceeds the limit
	if currentCount > int64(limit) {
		log.Printf("Rate limit exceeded for IP: %s", ip)
		return true
	}

	return false
}

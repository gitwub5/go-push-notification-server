package storage

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	Client *redis.Client
}

// NewRedisStore는 새로운 Redis 클라이언트를 생성합니다.
func NewRedisStore(addr, password string, db int) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &RedisStore{Client: rdb}
}

// AddNotification은 Redis에 알림을 저장합니다.
func (r *RedisStore) AddNotification(ctx context.Context, notification string) error {
	err := r.Client.LPush(ctx, "notifications", notification).Err()
	if err != nil {
		log.Printf("Failed to add notification to Redis: %v", err)
		return err
	}
	return nil
}

// GetAllNotifications은 Redis에서 모든 알림을 가져옵니다.
func (r *RedisStore) GetAllNotifications(ctx context.Context) ([]string, error) {
	notifications, err := r.Client.LRange(ctx, "notifications", 0, -1).Result()
	if err != nil {
		log.Printf("Failed to retrieve notifications from Redis: %v", err)
		return nil, err
	}
	return notifications, nil
}

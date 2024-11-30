package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// TODO: Redis를 사용하여 알림을 저장하고 가져오는 메소드를 구현합니다.
// TODO: Redis에 구독자 저장 및 해당 구독자에게 알림을 전송하는 메소드를 구현합니다.

type Notification struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	Token    string `json:"token"`
	Priority string `json:"priority"`
	Platform int    `json:"platform"`
	Status   string `json:"status"`
}

type RedisStore struct {
	Client *redis.Client
}

// NewRedisStore는 새로운 Redis 클라이언트를 생성합니다.
func NewRedisStore(addr, password string, db int) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // 설정 파일이나 환경 변수에서 받아온 주소 사용
		Password: password, // 환경 변수나 설정에서 비밀번호 사용
		DB:       db,       // 기본 DB 설정
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

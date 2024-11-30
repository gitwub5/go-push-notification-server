package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLStore struct {
	DB *gorm.DB
}

// Subscriber는 구독자 스키마를 정의하는 구조체입니다.
type Subscriber struct {
	gorm.Model
	Token    string `json:"token" gorm:"type:varchar(255);uniqueIndex"` // 디바이스 토큰 (고유 인덱스)
	Platform int    `json:"platform" gorm:"type:int"`                   // 플랫폼 정보 (예: 1 = iOS, 2 = Android)
	Topic    string `json:"topic" gorm:"type:varchar(255)"`             // 구독할 주제
}

// NewMySQLStore는 MySQL 연결을 설정하고, 필요한 경우 데이터베이스를 생성합니다.
func NewMySQLStore(user, password, host string, port int, dbName string) (*MySQLStore, error) {
	// 데이터베이스 지정 없이 DSN 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port)

	// 데이터베이스 연결 시도
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to MySQL: %v", err)
		return nil, err
	}

	// 데이터베이스 생성 쿼리 실행
	createDBQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	if err := db.Exec(createDBQuery).Error; err != nil {
		log.Printf("Failed to create database: %v", err)
		return nil, err
	}

	// 생성된 데이터베이스를 사용하도록 다시 연결
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to MySQL database: %v", err)
		return nil, err
	}

	// 테이블 생성 (자동 마이그레이션)
	if err := db.AutoMigrate(&Subscriber{}); err != nil {
		log.Printf("Failed to migrate database: %v", err)
		return nil, err
	}

	return &MySQLStore{DB: db}, nil
}

// / AddSubscriber는 새로운 구독자를 데이터베이스에 추가합니다.
func (m *MySQLStore) AddSubscriber(subscriber Subscriber) error {
	result := m.DB.Create(&subscriber)
	if result.Error != nil {
		log.Printf("Failed to add subscriber: %v", result.Error)
		return result.Error
	}
	return nil
}

// DeleteSubscriber는 주어진 토큰, 토픽 및 플랫폼을 기준으로 구독자를 삭제합니다.
func (m *MySQLStore) DeleteSubscriber(token string, topic string, platform int) error {
	var subscriber Subscriber
	result := m.DB.Where("token = ? AND topic = ? AND platform = ?", token, topic, platform).First(&subscriber)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("Subscriber not found for token: %s, topic: %s, platform: %d", token, topic, platform)
			return result.Error // 구독자가 존재하지 않으면 에러 반환
		}
		log.Printf("Failed to find subscriber: %v", result.Error)
		return result.Error
	}

	// 구독자 삭제
	if delErr := m.DB.Delete(&subscriber).Error; delErr != nil {
		log.Printf("Failed to delete subscriber: %v", delErr)
		return delErr
	}
	log.Printf("Successfully deleted subscriber: token=%s, topic=%s, platform=%d", token, topic, platform)
	return nil
}

// GetAllSubscribers는 모든 구독자를 반환합니다.
func (m *MySQLStore) GetAllSubscribers() ([]Subscriber, error) {
	var subscribers []Subscriber
	result := m.DB.Find(&subscribers)
	if result.Error != nil {
		return nil, result.Error
	}
	return subscribers, nil
}

// GetSubscriberByToken은 특정 디바이스 토큰으로 구독자를 조회합니다.
func (m *MySQLStore) GetSubscriberByToken(token string) (*Subscriber, error) {
	var subscriber Subscriber
	result := m.DB.First(&subscriber, "token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subscriber, nil
}

// GetSubscribersByTopic은 특정 토픽을 구독한 구독자들을 조회합니다.
func (m *MySQLStore) GetSubscribersByTopic(topic string) ([]Subscriber, error) {
	var subscribers []Subscriber
	result := m.DB.Where("topic = ?", topic).Find(&subscribers)
	if result.Error != nil {
		return nil, result.Error
	}
	return subscribers, nil
}

package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MySQLStore struct {
	DB *gorm.DB
}

// Subscriber는 구독자 데이터를 정의하는 구조체입니다.
type Subscriber struct {
	gorm.Model
	Token string `json:"token" gorm:"type:varchar(255);uniqueIndex"` // 디바이스 토큰 (고유 인덱스)
	Topic string `json:"topic" gorm:"type:varchar(255)"`             // 구독할 주제
}

// NewMySQLStore는 SQLite 연결을 설정합니다.
func NewMySQLStore(dsn string) (*MySQLStore, error) {
	// SQLite 사용 (임시로 인메모리 데이터베이스 사용)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to initialize SQLite database: %v", err)
		return nil, err
	}

	// Notification 및 Subscriber 테이블 생성 (자동 마이그레이션)
	db.AutoMigrate(&Notification{}, &Subscriber{})

	return &MySQLStore{DB: db}, nil
}

// AddSubscriber는 새로운 구독자를 데이터베이스에 추가합니다.
func (m *MySQLStore) AddSubscriber(subscriber Subscriber) error {
	result := m.DB.Create(&subscriber)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteSubscriber는 주어진 토큰과 토픽을 기준으로 구독자를 삭제합니다.
func (m *MySQLStore) DeleteSubscriber(token string, topic string) error {
	var subscriber Subscriber
	result := m.DB.Where("token = ? AND topic = ?", token, topic).First(&subscriber)
	if result.Error != nil {
		return result.Error // 구독자가 존재하지 않으면 에러 반환
	}

	// 구독자 삭제
	if delErr := m.DB.Delete(&subscriber).Error; delErr != nil {
		log.Printf("Failed to delete subscriber: %v", delErr)
		return delErr
	}
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

// TODO: sqlite 대신 MySQL을 사용하도록 변경
// import (
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

//dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

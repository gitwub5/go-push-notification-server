package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MySQLStore struct {
	DB *gorm.DB
}

// Notification은 푸시 알림의 데이터 구조를 정의합니다.
type Notification struct {
	gorm.Model        // 기본적으로 ID, CreatedAt, UpdatedAt, DeletedAt 필드를 포함합니다.
	Title      string `json:"title" gorm:"type:varchar(255)"`   // 제목
	Message    string `json:"message" gorm:"type:text"`         // 메시지 내용
	Token      string `json:"token" gorm:"type:varchar(255)"`   // 디바이스 토큰
	Priority   string `json:"priority" gorm:"type:varchar(10)"` // 우선순위 (high, normal 등)
	Platform   int    `json:"platform"`                         // 플랫폼 (1 = iOS, 2 = Android)
	Status     string `json:"status" gorm:"type:varchar(50)"`   // 상태 (delivered, failed 등)
}

// NewMySQLStore는 SQLite 연결을 임시로 설정합니다.
func NewMySQLStore(dsn string) (*MySQLStore, error) {
	// SQLite 사용 (임시로 인메모리 데이터베이스 사용)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to initialize SQLite database: %v", err)
		return nil, err
	}

	// Notification 테이블 생성 (자동 마이그레이션)
	db.AutoMigrate(&Notification{})

	return &MySQLStore{DB: db}, nil
}

// AddNotification은 SQLite에 알림을 저장하는 함수입니다.
func (m *MySQLStore) AddNotification(notification Notification) error {
	result := m.DB.Create(&notification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllNotifications은 SQLite에서 모든 알림을 가져옵니다.
func (m *MySQLStore) GetAllNotifications() ([]Notification, error) {
	var notifications []Notification
	result := m.DB.Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifications, nil
}

// GetNotificationByID는 알림 ID로 특정 알림을 조회하는 함수입니다.
func (m *MySQLStore) GetNotificationByID(id uint) (*Notification, error) {
	var notification Notification
	result := m.DB.First(&notification, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &notification, nil
}

//dsn := "username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

// package storage

// import (
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type MySQLStore struct {
// 	DB *gorm.DB
// }

// // Notification은 푸시 알림의 데이터 구조를 정의합니다.
// type Notification struct {
// 	gorm.Model        // 기본적으로 ID, CreatedAt, UpdatedAt, DeletedAt 필드를 포함합니다.
// 	Title      string `json:"title" gorm:"type:varchar(255)"`   // 제목
// 	Message    string `json:"message" gorm:"type:text"`         // 메시지 내용
// 	Token      string `json:"token" gorm:"type:varchar(255)"`   // 디바이스 토큰
// 	Priority   string `json:"priority" gorm:"type:varchar(10)"` // 우선순위 (high, normal 등)
// 	Platform   int    `json:"platform"`                         // 플랫폼 (1 = iOS, 2 = Android)
// 	Status     string `json:"status" gorm:"type:varchar(50)"`   // 상태 (delivered, failed 등)
// }

// // NewMySQLStore는 MySQL 연결을 설정합니다.
// func NewMySQLStore(dsn string) (*MySQLStore, error) {
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Printf("Failed to connect to MySQL: %v", err)
// 		return nil, err
// 	}

// 	// Notification 테이블 생성 (자동 마이그레이션)
// 	db.AutoMigrate(&Notification{})

// 	return &MySQLStore{DB: db}, nil
// }

// // AddNotification은 MySQL에 알림을 저장하는 함수입니다.
// func (m *MySQLStore) AddNotification(notification Notification) error {
// 	result := m.DB.Create(&notification)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// // GetAllNotifications은 MySQL에서 모든 알림을 가져옵니다.
// func (m *MySQLStore) GetAllNotifications() ([]Notification, error) {
// 	var notifications []Notification
// 	result := m.DB.Find(&notifications)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return notifications, nil
// }

// // GetNotificationByID는 알림 ID로 특정 알림을 조회하는 함수입니다.
// func (m *MySQLStore) GetNotificationByID(id uint) (*Notification, error) {
// 	var notification Notification
// 	result := m.DB.First(&notification, id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &notification, nil
// }

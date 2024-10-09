package storage

import "sync"

// MemoryStore는 메모리 내에서 알림을 저장합니다.
type MemoryStore struct {
	Notifications []string
	mu            sync.Mutex
}

// AddNotification은 새로운 알림을 저장소에 추가합니다.
func (m *MemoryStore) AddNotification(notification string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Notifications = append(m.Notifications, notification)
}

// GetAllNotifications은 저장된 모든 알림을 반환합니다.
func (m *MemoryStore) GetAllNotifications() []string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Notifications
}

package gosession

import (
	"errors"
	"fmt"
	"sync"
)

type memoryProvider struct {
	sessions map[string]*memorySession
	mutex    sync.Mutex
}

func NewMemoryProvider() *memoryProvider {
	return &memoryProvider{
		sessions: make(map[string]*memorySession),
	}
}

func (m *memoryProvider) Start() (Session, error) {
	memorySession := newMemorySession()
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.sessions[memorySession.id] = memorySession
	return memorySession, nil
}

func (m *memoryProvider) Stop(sessionID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, exists := m.sessions[sessionID]
	if !exists {
		errorMessage := fmt.Sprintf("%s does not exist", sessionID)
		return errors.New(errorMessage)
	}
	delete(m.sessions, sessionID)
	return nil
}

func (m *memoryProvider) Get(sessionID string) (Session, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	session, exists := m.sessions[sessionID]
	if !exists {
		errorMessage := fmt.Sprintf("%s does not exist", sessionID)
		return nil, errors.New(errorMessage)
	}
	return session, nil
}

func (m *memoryProvider) F() []string {
	keys := []string{}
	for key := range m.sessions {
		keys = append(keys, key)
	}
	return keys
}

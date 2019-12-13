package gosession

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/sk409/gotype"
)

type memorySession struct {
	id    string
	data  map[string]interface{}
	mutex sync.Mutex
}

func newMemorySession() *memorySession {
	return &memorySession{
		id:   makeSessionID(),
		data: make(map[string]interface{}),
	}
}

func (m *memorySession) ID() string {
	return m.id
}

func (m *memorySession) Store(key string, value interface{}) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.data[key] = value
	return nil
}

func (m *memorySession) String(key string) (string, error) {
	i, err := m.get(key, "string")
	if err != nil {
		return "", err
	}
	return i.(string), nil
}

func (m *memorySession) Int(key string) (int, error) {
	i, err := m.get(key, "int")
	if err != nil {
		return 0, err
	}
	return i.(int), nil
}

func (m *memorySession) Uint(key string) (uint, error) {
	i, err := m.get(key, "uint")
	if err != nil {
		return 0, err
	}
	return i.(uint), nil
}

func (m *memorySession) Object(key string, ptr interface{}) error {
	if !gotype.IsPointer(ptr) {
		return errors.New("The second argument must be a pointer to interface")
	}
	drt := reflect.TypeOf(ptr).Elem()
	drv := reflect.ValueOf(ptr).Elem()
	if !gotype.IsStruct(drv) {
		return errors.New("The second argument must be a pointer to struct")
	}
	i, err := m.get(key, "struct")
	if err != nil {
		return err
	}
	srt := reflect.TypeOf(i)
	srv := reflect.ValueOf(i)
	if srt.Name() != drt.Name() {
		errorMessage := fmt.Sprintf("Tried to get %s but specified %s for second argumant", srt.Name(), drt.Name())
		return errors.New(errorMessage)
	}
	for index := 0; index < srt.NumField(); index++ {
		drv.Field(index).Set(srv.Field(index))
	}
	return nil
}

func (m *memorySession) get(key, _type string) (interface{}, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	i, exists := m.data[key]
	if !exists {
		return nil, errorSpecifiedKeyDoesNotExsistInThisSession(key)
	}
	rt := reflect.TypeOf(i)
	kind := fmt.Sprint(rt.Kind())
	if kind != _type {
		return nil, errorValueCorrespondingToKeyIsNotSpecifiedType(key, _type)
	}
	return i, nil
}

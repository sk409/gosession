package gosession

import (
	"crypto/rand"
	"encoding/base64"
)

type Session interface {
	ID() string
	Store(string, interface{}) error
	String(string) (string, error)
	Int(string) (int, error)
	Uint(string) (uint, error)
	Object(string, interface{}) error
}

func makeSessionID() string {
	bytes := make([]byte, 128)
	rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}

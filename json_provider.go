package gosession

// import (
// 	"os"
// 	"path/filepath"
// 	"sync"
// )

// type jsonProvider struct {
// 	directory string
// 	sessions  map[string]*memorySession
// 	mutex     sync.Mutex
// }

// func NewJSONProvider(directory string) *jsonProvider {
// 	return &jsonProvider{
// 		directory: directory,
// 		sessions:  make(map[string]*memorySession),
// 	}
// }

// func (j *jsonProvider) Start() (Session, error) {
// 	sessionID := makeSessionID()
// 	path := filepath.Join(j.directory, sessionID)
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	jsonSession :=
// }

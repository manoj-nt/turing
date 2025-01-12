package fileversionmanager

import (
	"fmt"
	"sync"
	"time"
)

// Version represents a single version of a file.
type Version struct {
	ID   int
	Data []byte // Store content of the file version, this can be anything suitable
	Time int64  // Time when the version was created, in unix seconds
}

// FileVersionManager manages versions for a file.
type FileVersionManager struct {
	versions map[string][]Version
	lock     sync.Mutex
}

// Singleton instance of FileVersionManager.
var (
	instance *FileVersionManager
	once     sync.Once
)

// GetInstance returns a single global instance of FileVersionManager.
func GetInstance() *FileVersionManager {
	once.Do(func() {
		instance = &FileVersionManager{
			versions: make(map[string][]Version),
		}
	})
	return instance
}

// AddVersion adds a new version for a given file.
func (v *FileVersionManager) AddVersion(fileName string, data []byte) error {
	v.lock.Lock()
	defer v.lock.Unlock()

	if v.versions[fileName] == nil {
		v.versions[fileName] = make([]Version, 0)
	}

	newVersion := Version{
		ID:   len(v.versions[fileName]) + 1,
		Data: data,
		Time: time.Now().Unix(),
	}

	v.versions[fileName] = append(v.versions[fileName], newVersion)
	return nil
}

// GetVersion retrieves a specific version for a given file by ID.
func (v *FileVersionManager) GetVersion(fileName string, versionID int) (*Version, error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	for _, v := range v.versions[fileName] {
		if v.ID == versionID {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("version not found for file: %s, id: %d", fileName, versionID)
}

// ListVersions lists all versions for a given file.
func (v *FileVersionManager) ListVersions(fileName string) ([]Version, error) {
	v.lock.Lock()
	defer v.lock.Unlock()

	if v.versions[fileName] == nil {
		return nil, fmt.Errorf("no versions found for file: %s", fileName)
	}
	return v.versions[fileName], nil
}

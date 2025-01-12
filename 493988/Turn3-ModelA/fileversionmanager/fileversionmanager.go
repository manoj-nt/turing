package fileversionmanager

import (
	"fmt"
	"sync"
	"time"
)

// Version represents a single version of a file.
type Version interface {
	ID() int
	Data() []byte
	Timestamp() int64
	UserID() string
	Description() string
}

// TextVersion implements the Version interface for text files.
type TextVersion struct {
	ID          int
	Data        []byte
	Timestamp   int64
	UserID      string
	Description string
}

func (tv *TextVersion) ID() int             { return tv.ID }
func (tv *TextVersion) Data() []byte        { return tv.Data }
func (tv *TextVersion) Timestamp() int64    { return tv.Timestamp }
func (tv *TextVersion) UserID() string      { return tv.UserID }
func (tv *TextVersion) Description() string { return tv.Description }

// VersionFactory is an interface that defines the factory method.
type VersionFactory interface {
	CreateVersion(id int, data []byte, timestamp int64, userID string, description string) Version
}

// TextVersionFactory implements VersionFactory for text files.
type TextVersionFactory struct{}

func (tvf *TextVersionFactory) CreateVersion(id int, data []byte, timestamp int64, userID string, description string) Version {
	return &TextVersion{
		ID:          id,
		Data:        data,
		Timestamp:   timestamp,
		UserID:      userID,
		Description: description,
	}
}

// FileVersionManager manages versions for a file.
type FileVersionManager struct {
	versions       map[string][]Version
	versionFactory VersionFactory
	rwlock         sync.RWMutex
}

// Singleton instance of FileVersionManager.
var (
	instance *FileVersionManager
	once     sync.Once
)

// GetInstance returns a single global instance of FileVersionManager.
func GetInstance(versionFactory VersionFactory) *FileVersionManager {
	once.Do(func() {
		instance = &FileVersionManager{
			versions:       make(map[string][]Version),
			versionFactory: versionFactory,
		}
	})
	return instance
}

// AddVersion adds a new version for a given file.
func (v *FileVersionManager) AddVersion(fileName string, data []byte, userID string, description string) error {
	v.rwlock.Lock()
	defer v.rwlock.Unlock()

	if v.versions[fileName] == nil {
		v.versions[fileName] = make([]Version, 0)
	}

	timestamp := time.Now().Unix()
	newVersionID := len(v.versions[fileName]) + 1

	newVersion := v.versionFactory.CreateVersion(
		newVersionID,
		data,
		timestamp,
		userID,
		description,
	)

	v.versions[fileName] = append(v.versions[fileName], newVersion)
	return nil
}

// GetVersion retrieves a specific version for a given file by ID.
func (v *FileVersionManager) GetVersion(fileName string, versionID int) (*Version, error) {
	v.rwlock.RLock()
	defer v.rwlock.RUnlock()

	for _, v := range v.versions[fileName] {
		if v.ID() == versionID {
			return v, nil
		}
	}
	return nil, fmt.Errorf("version not found for file: %s, id: %d", fileName, versionID)
}

// ListVersions lists all versions for a given file.
func (v *FileVersionManager) ListVersions(fileName string) ([]Version, error) {
	v.rwlock.RLock()
	defer v.rwlock.RUnlock()

	if v.versions[fileName] == nil {
		return nil, fmt.Errorf("no versions found for file: %s", fileName)
	}

	// Return a copy of the slice to prevent data race
	return append([]Version{}, v.versions[fileName]...), nil
}

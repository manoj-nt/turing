// user_repository.go
package persistence

import (
	"encoding/xml"
	"os"

	"app/domain"
)

type FileUserRepository struct {
	filePath string
	users    []domain.User
}

func NewFileUserRepository(filePath string) *FileUserRepository {
	return &FileUserRepository{filePath: filePath}
}

func (r *FileUserRepository) Save(user domain.User) error {
	// Implement saving logic to a file or database
	r.users = append(r.users, user)
	// Here you could serialize back to XML or another format as needed
	return nil
}

func (r *FileUserRepository) FindAll() ([]domain.User, error) {
	err := r.loadUsers()
	if err != nil {
		return nil, err
	}
	return r.users, nil
}

func (r *FileUserRepository) loadUsers() error {
	file, err := os.Open(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	for {
		var user domain.User
		if err := decoder.Decode(&user); err != nil {
			break
		}
		r.users = append(r.users, user)
	}

	return nil
}

// models.go
package domain

type User struct {
	ID    string
	Name  string
	Email string
}

type UserRepository interface {
	Save(user User) error
	FindAll() ([]User, error)
}

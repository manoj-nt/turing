package queryparamorm

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// User is an ORM model representing a user.
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// CreateUserFromQueryParams creates a new User record in the database from the given query parameters.
func CreateUserFromQueryParams(db *sqlx.DB, r *mux.Router) error {
	user := User{}
	if err := populateStructFromQueryParams(&user, r); err != nil {
		return err
	}

	query := "INSERT INTO users (name, age) VALUES (?, ?)"
	result, err := db.Exec(query, user.Name, user.Age)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(lastInsertID)

	fmt.Println("User created successfully:", user)
	return nil
}

// UpdateUserFromQueryParams updates an existing User record in the database from the given query parameters.
func UpdateUserFromQueryParams(db *sqlx.DB, r *mux.Router) error {
	user := User{}
	if err := populateStructFromQueryParams(&user, r); err != nil {
		return err
	}

	query := "UPDATE users SET name = ?, age = ? WHERE id = ?"
	result, err := db.Exec(query, user.Name, user.Age, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("User updated successfully. Rows affected:", rowsAffected)
	return nil
}

func populateStructFromQueryParams(s interface{}, r *mux.Router) error {
	v := reflect.ValueOf(s).Elem()
	t := reflect.TypeOf(s).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbTag := field.Tag.Get("db")
		if dbTag == "" {
			continue
		}

		paramValue := r.QueryParam(dbTag)
		if paramValue == "" {
			continue
		}

		switch field.Type.Kind() {
		case reflect.Int:
			value, err := strconv.Atoi(paramValue)
			if err != nil {
				return err
			}
			v.Field(i).SetInt(int64(value))
		case reflect.String:
			v.Field(i).SetString(paramValue)
		default:
			return fmt.Errorf("unsupported field type: %s", field.Type.Kind())
		}
	}
	return nil
}

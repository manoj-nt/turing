package paramserializer

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func SerializeQueryParams(rawQuery string) (*User, error) {
	params, err := url.ParseQuery(rawQuery)
	if err != nil {
		return nil, err
	}
	user := &User{}
	err = decode(user, params)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func decode(val interface{}, params url.Values) error {
	v := reflect.ValueOf(val).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("query")
		if tag == "-" {
			continue
		}

		if tag != "" {
			err := handleNestedStruct(value, params, tag)
			if err != nil {
				return err
			}
		} else {
			err := handleSimpleField(value, params, field.Name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func handleNestedStruct(v reflect.Value, params url.Values, tag string) error {
	if v.Kind() == reflect.Struct {
		err := decode(v.Addr().Interface(), params)
		if err != nil {
			return fmt.Errorf("failed to decode nested struct %q: %v", tag, err)
		}
	} else if v.Kind() == reflect.Ptr {
		elem := v.Elem()
		if elem.Kind() != reflect.Struct {
			return fmt.Errorf("nested pointer must point to a struct: %v", tag)
		}
		// Set the pointer to a new struct value
		v.Set(reflect.New(elem.Type()))
		err := decode(v.Interface(), params)
		if err != nil {
			return fmt.Errorf("failed to decode nested struct %q: %v", tag, err)
		}
	} else {
		return fmt.Errorf("invalid nested type for %q: %v", tag, v.Kind())
	}
	return nil
}

func handleSimpleField(v reflect.Value, params url.Values, fieldName string) error {
	key := fieldName
	if name := params.Get(key); name != "" {
		switch v.Kind() {
		case reflect.String:
			v.SetString(name)
		case reflect.Int:
			id, err := strconv.Atoi(name)
			if err != nil {
				return fmt.Errorf("invalid integer value for %q: %v", key, err)
			}
			v.SetInt(int64(id))
		default:
			return fmt.Errorf("unsupported simple field type for %q: %v", key, v.Kind())
		}
	}
	return nil
}

// User is the ORM model with embedded Address struct.
type User struct {
	ID      int      `query:"-"` // "-" tag will skip this field from parsing
	Name    string   `query:"name"`
	Email   string   `query:"email"`
	Address *Address `query:"address"` // "address" tag for nested struct
}

type Address struct {
	City  string `query:"city"`
	State string `query:"state"`
}

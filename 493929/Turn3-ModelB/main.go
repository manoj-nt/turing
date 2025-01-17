package main

import (
	"errors"
	"fmt"
	"sync"
)

// MultiError is a struct to hold multiple errors
type MultiError struct {
	errors []error
}

func (me *MultiError) Error() string {
	return fmt.Sprintf("%d errors occurred: %v", len(me.errors), me.errors)
}

func (me *MultiError) Add(err error) {
	me.errors = append(me.errors, err)
}

// Callback types representing different kinds of operations on UserData
type ValidationCallback func(data UserData) error
type TransformationCallback func(data UserData) UserData
type LoggingCallback func(data UserData)

// UserData represents a user's data
type UserData struct {
	ID    int
	Name  string
	Email string
}

// CallbackManager manages different types of callbacks
type CallbackManager struct {
	validationCallbacks     []ValidationCallback
	transformationCallbacks []TransformationCallback
	loggingCallbacks        []LoggingCallback
}

// AddValidationCallback registers a new validation callback
func (cm *CallbackManager) AddValidationCallback(cb ValidationCallback) {
	cm.validationCallbacks = append(cm.validationCallbacks, cb)
}

// AddTransformationCallback registers a new transformation callback
func (cm *CallbackManager) AddTransformationCallback(cb TransformationCallback) {
	cm.transformationCallbacks = append(cm.transformationCallbacks, cb)
}

// AddLoggingCallback registers a new logging callback
func (cm *CallbackManager) AddLoggingCallback(cb LoggingCallback) {
	cm.loggingCallbacks = append(cm.loggingCallbacks, cb)
}

// ExecuteCallbacks runs through the registered callbacks concurrently
func (cm *CallbackManager) ExecuteCallbacks(data UserData) (UserData, error) {
	var wg sync.WaitGroup
	var me MultiError
	errorCh := make(chan error, 1)

	// Run validation callbacks concurrently
	wg.Add(len(cm.validationCallbacks))
	for _, vcb := range cm.validationCallbacks {
		go func(cb ValidationCallback) {
			defer wg.Done()
			if err := cb(data); err != nil {
				errorCh <- err
			}
		}(vcb)
	}

	// Run transformation callbacks concurrently
	wg.Add(len(cm.transformationCallbacks))
	for _, tcb := range cm.transformationCallbacks {
		go func(cb TransformationCallback) {
			defer wg.Done()
			data = cb(data)
		}(tcb)
	}

	// Run logging callbacks concurrently
	wg.Add(len(cm.loggingCallbacks))
	for _, lcb := range cm.loggingCallbacks {
		go func(cb LoggingCallback) {
			defer wg.Done()
			cb(data)
		}(lcb)
	}

	// Wait for all callbacks to complete
	go func() {
		wg.Wait()
		close(errorCh)
	}()

	// Collect errors from the error channel
	for err := range errorCh {
		me.Add(err)
	}

	if len(me.errors) > 0 {
		return data, &me
	}
	return data, nil
}

// Example usage:
func main() {
	cm := &CallbackManager{}

	// Register validation callbacks
	cm.AddValidationCallback(func(data UserData) error {
		if data.Email == "" {
			return errors.New("email cannot be empty")
		}
		return nil
	})

	cm.AddValidationCallback(func(data UserData) error {
		if data.Name == "" {
			return errors.New("name cannot be empty")
		}
		return nil
	})

	// Register transformation callbacks
	cm.AddTransformationCallback(func(data UserData) UserData {
		data.Name = string(data.Name[0]-32) + data.Name[1:] // Capitalize name
		return data
	})

	// Register logging callback
	cm.AddLoggingCallback(func(data UserData) {
		fmt.Printf("Processed User: %+v\n", data)
	})

	// Execute with sample data
	userData := UserData{ID: 1, Name: "john doe", Email: "john.doe@example.com"}
	processedData, err := cm.ExecuteCallbacks(userData)
	if err != nil {
		fmt.Printf("Errors occurred: %v\n", err)
	} else {
		fmt.Printf("Final Transformed Data: %+v\n", processedData)
	}
}

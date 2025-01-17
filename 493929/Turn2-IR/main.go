package main

// Callback interface
type Callback interface {
	ProcessData(data string) string
}

// Function that accepts a callback interface
func ProcessData(callback Callback) string {
	return callback.ProcessData("data")
}

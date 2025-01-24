package main

import (
	"fmt"
	"os"
)

// FileHandler interface now includes methods for additional file operations
type FileHandler interface {
	Open(name string) (*os.File, error)
	Close(file *os.File) error
	Write(file *os.File, data []byte) (int, error)
	Exists(name string) bool
}

// OSFileHandler implementation of the FileHandler interface
type OSFileHandler struct{}

// Open implements the method to open a file
func (f *OSFileHandler) Open(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
}

// Close implements the method to close a file
func (f *OSFileHandler) Close(file *os.File) error {
	return file.Close()
}

// Write implements the method to write data to a file
func (f *OSFileHandler) Write(file *os.File, data []byte) (int, error) {
	return file.Write(data)
}

// Exists checks if the file exists
func (f *OSFileHandler) Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// Application core logic which manipulates file resources
type FileProcessor struct {
	fileHandler FileHandler
}

// NewFileProcessor is a constructor that uses dependency injection
func NewFileProcessor(fh FileHandler) *FileProcessor {
	return &FileProcessor{fileHandler: fh}
}

// ProcessFile demonstrates file processing
func (fp *FileProcessor) ProcessFile(filename string) error {
	if !fp.fileHandler.Exists(filename) {
		fmt.Println("File does not exist. Creating and writing to file.")

		// Open the file for writing
		file, err := fp.fileHandler.Open(filename)
		if err != nil {
			return fmt.Errorf("failed to open file for writing: %v", err)
		}
		defer fp.fileHandler.Close(file)

		// Write to the file
		_, err = fp.fileHandler.Write(file, []byte("Hello, World!"))
		if err != nil {
			return fmt.Errorf("failed to write to file: %v", err)
		}
	}

	file, err := fp.fileHandler.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer fp.fileHandler.Close(file)

	// Additional processing logic could be implemented here
	fmt.Println("Processing file:", file.Name())
	return nil
}

func main() {
	fileHandler := &OSFileHandler{}
	fileProcessor := NewFileProcessor(fileHandler)

	err := fileProcessor.ProcessFile("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

package main

import (
	"fmt"
	"os"
)

// FileHandler interface will define methods for file operations
type FileHandler interface {
	Open(name string) (*os.File, error)
	Close(file *os.File) error
}

// OSFileHandler is a concrete implementation of the FileHandler interface
type OSFileHandler struct{}

// Open implements the method to open a file
func (f *OSFileHandler) Open(name string) (*os.File, error) {
	return os.Open(name)
}

// Close implements the method to close a file
func (f *OSFileHandler) Close(file *os.File) error {
	return file.Close()
}

// Application core logic which manipulates file resources
type FileProcessor struct {
	fileHandler FileHandler
}

// NewFileProcessor is a constructor that uses dependency injection
func NewFileProcessor(fh FileHandler) *FileProcessor {
	return &FileProcessor{fileHandler: fh}
}

// ProcessFile is a method to demonstrate file processing
func (fp *FileProcessor) ProcessFile(filename string) error {
	file, err := fp.fileHandler.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer fp.fileHandler.Close(file)

	// Here you would implement the file processing logic
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

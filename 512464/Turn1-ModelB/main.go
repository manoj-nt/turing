package main

import (
	"fmt"
	"os"
)

// FileManager interface defines the operations that can be performed on files.
type FileManager interface {
	Open(name string) (*os.File, error)
	Close(file *os.File) error
}

// fileManager implements the FileManager interface.
type fileManager struct{}

// Open opens a file using the os package.
func (fm fileManager) Open(name string) (*os.File, error) {
	return os.Open(name)
}

// Close closes a file using the os package.
func (fm fileManager) Close(file *os.File) error {
	return file.Close()
}

// ProcessingService performs some processing using the file manager.
type ProcessingService struct {
	fm FileManager
}

// NewProcessingService creates a new ProcessingService with a file manager.
func NewProcessingService(fm FileManager) *ProcessingService {
	return &ProcessingService{fm: fm}
}

// ProcessFile processes the specified file.
func (ps *ProcessingService) ProcessFile(fileName string) error {
	file, err := ps.fm.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer ps.fm.Close(file)

	// Perform processing on the file
	// For simplicity, we'll just read the file's content.
	content, err := os.ReadFile(file.Name())
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	fmt.Println("File content:", string(content))
	return nil
}

func main() {
	// Create an instance of the fileManager
	fm := fileManager{}

	// Create a ProcessingService with the fileManager
	ps := NewProcessingService(fm)

	// Process a file using the service
	err := ps.ProcessFile("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

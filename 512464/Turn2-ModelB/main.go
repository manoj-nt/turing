package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileHandler interface with additional methods
type FileHandler interface {
	Open(name string) (*os.File, error)
	Close(file *os.File) error
	WriteFile(filename string, data []byte) error
	FileExists(filename string) bool
}

// OSFileHandler implementation with additional file operations
type OSFileHandler struct{}

func (f *OSFileHandler) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (f *OSFileHandler) Close(file *os.File) error {
	return file.Close()
}

func (f *OSFileHandler) WriteFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}

func (f *OSFileHandler) FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}

// FileProcessor struct remains the same
type FileProcessor struct {
	fileHandler FileHandler
}

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

// Function to demonstrate writing to a file
func (fp *FileProcessor) WriteDataToFile(filename string, data []byte) error {
	return fp.fileHandler.WriteFile(filename, data)
}

// Function to check if a file exists
func (fp *FileProcessor) IsFileExist(filename string) bool {
	return fp.fileHandler.FileExists(filename)
}

func main() {
	fileHandler := &OSFileHandler{}
	fileProcessor := NewFileProcessor(fileHandler)

	// Example usage of additional operations
	err := fileProcessor.WriteDataToFile("example_write.txt", []byte("Hello, World!"))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	if fileProcessor.IsFileExist("example_write.txt") {
		fmt.Println("File exists.")
	} else {
		fmt.Println("File does not exist.")
	}

	// ... (Other operations)
}

package main

import (
	"os"
	"testing"
)

// Function to generate a large file using 'defer'
func generateFileDefer(fileName string, size int) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	// Generate data and write to the file
	data := make([]byte, size)
	_, err = f.Write(data)
	return err
}

// Function to generate a large file manually
func generateFileManual(fileName string, size int) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	// Generate data and write to the file
	data := make([]byte, size)
	_, err = f.Write(data)
	if err != nil {
		// Clean up if there's an error during writing
		f.Close()
		return err
	}
	// Close the file when done
	err = f.Close()
	return err
}

func TestGenerateFile(t *testing.T) {
	const fileName = "test_file.txt"
	const fileSize = 1000000 // 1MB

	// Test case 1: Generate file using 'defer'
	t.Run("GenerateFileDefer", func(t *testing.T) {
		err := generateFileDefer(fileName, fileSize)
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(fileName)
	})

	// Test case 2: Generate file manually
	t.Run("GenerateFileManual", func(t *testing.T) {
		err := generateFileManual(fileName, fileSize)
		if err != nil {
			t.Fatal(err)
		}
		err = os.Remove(fileName)
		if err != nil {
			t.Fatal(err)
		}
	})
}

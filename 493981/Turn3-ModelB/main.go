package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Function to simulate processing each line of a file
func processLine(line string) error {
	// Create a temporary file for processing the line
	tempFile, err := ioutil.TempFile("", "temp-line-")
	if err != nil {
		return err
	}
	// Simulate some work on the line
	fmt.Fprintln(tempFile, line)
	// Close the temporary file
	tempFile.Close()
	return nil
}

// Function to manage the resource lifetime of the temporary file
func withTemporaryFile(line string, f func(*os.File) error) error {
	tempFile, err := ioutil.TempFile("", "temp-line-")
	if err != nil {
		return err
	}
	defer tempFile.Close()
	// Write the line to the temporary file
	fmt.Fprintln(tempFile, line)
	// Call the provided function with the temporary file as an argument
	return f(tempFile)
}

func processLargeFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Original approach (risk of using a closed temporary file)
		// defer func() {
		//     tempFile, _ := ioutil.TempFile("", "temp-line-")
		//     defer tempFile.Close()
		//     processLine(line, tempFile)
		// }()

		// Mitigated approach using withTemporaryFile helper
		if err := withTemporaryFile(line, func(tempFile *os.File) error {
			return processLine(line)
		}); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func main() {
	// Example file content for testing
	exampleFileContent := `
line1
line2
line3
`
	// Create a temporary file to process
	tmpFile, err := ioutil.TempFile("", "example-file-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	if _, err := tmpFile.WriteString(exampleFileContent); err != nil {
		panic(err)
	}
	tmpFile.Close()

	err = processLargeFile(tmpFile.Name())
	if err != nil {
		fmt.Println("Error processing file:", err)
	} else {
		fmt.Println("File processing completed successfully.")
	}
}

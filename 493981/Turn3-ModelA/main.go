package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func processFile(filePath string) error {
	// Open the file and defer its closure
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %s: %v", filePath, err)
	}
	defer file.Close() // Defer ensures that each file is closed as soon as processFile returns

	// Do something with the file, for example, read it
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Process the line
		fmt.Println(scanner.Text()) // For demonstration purposes
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", filePath, err)
	}

	return nil
}

func processFiles(filePaths []string) {
	for _, filePath := range filePaths {
		// Handle each file separately
		if err := processFile(filePath); err != nil {
			log.Printf("Failed to process %s: %v\n", filePath, err)
		}
		// Files are closed immediately after processing in `processFile`
		// No deferred calls stacking
	}
}

func main() {
	filePaths := []string{"file1.txt", "file2.txt", "file3.txt"}

	start := time.Now()
	processFiles(filePaths)
	elapsed := time.Since(start)
	fmt.Printf("Processed files in %s\n", elapsed)
}

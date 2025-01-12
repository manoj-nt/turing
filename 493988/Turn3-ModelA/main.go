package main

import (
	"fmt"
	"modela/fileversionmanager"
)

func main() {
	// Initialize the version factory and get the singleton instance of FileVersionManager
	versionFactory := &fileversionmanager.TextVersionFactory{}
	fvManager := fileversionmanager.GetInstance(versionFactory)

	// Add some versions for a file
	err := fvManager.AddVersion("testfile.txt", []byte("Version 1 data"), "user123", "Initial version")
	if err != nil {
		fmt.Println("Error adding version 1:", err)
		return
	}

	err = fvManager.AddVersion("testfile.txt", []byte("Version 2 data"), "user456", "Second version")
	if err != nil {
		fmt.Println("Error adding version 2:", err)
		return
	}

	// Retrieve a specific version by ID
	version1, err := fvManager.GetVersion("testfile.txt", 1)
	if err != nil {
		fmt.Println("Error retrieving version 1:", err)
		return
	}

	fmt.Printf("Retrieved Version 1: %d, Description: %s, Timestamp: %d\n", version1.ID(), version1.Description(), version1.Timestamp())

	version2, err := fvManager.GetVersion("testfile.txt", 2)
	if err != nil {
		fmt.Println("Error retrieving version 2:", err)
		return
	}

	fmt.Printf("Retrieved Version 2: %d, Description: %s, Timestamp: %d\n", version2.ID(), version2.Description(), version2.Timestamp())

	// List all versions for a file
	versions, err := fvManager.ListVersions("testfile.txt")
	if err != nil {
		fmt.Println("Error listing versions:", err)
		return
	}

	fmt.Println("\nAll Versions for 'testfile.txt':")
	for _, v := range versions {
		fmt.Printf("Version ID: %d, Description: %s, Timestamp: %d\n", v.ID(), v.Description(), v.Timestamp())
	}

	// Simulate adding more versions and listing them
	err = fvManager.AddVersion("testfile.txt", []byte("Version 3 data"), "user789", "Third version")
	if err != nil {
		fmt.Println("Error adding version 3:", err)
		return
	}

	versions, err = fvManager.ListVersions("testfile.txt")
	if err != nil {
		fmt.Println("Error listing versions after adding Version 3:", err)
		return
	}

	fmt.Println("\nAll Versions for 'testfile.txt' after adding Version 3:")
	for _, v := range versions {
		fmt.Printf("Version ID: %d, Description: %s, Timestamp: %d\n", v.ID(), v.Description(), v.Timestamp())
	}
}

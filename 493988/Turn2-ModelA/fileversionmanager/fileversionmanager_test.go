package fileversionmanager

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFileVersionManager(t *testing.T) {
	fvManager := GetInstance()

	// Test adding versions
	err := fvManager.AddVersion("testfile.txt", []byte("Version 1 data"))
	if err != nil {
		t.Fatal("Error adding version 1:", err)
	}

	err = fvManager.AddVersion("testfile.txt", []byte("Version 2 data"))
	if err != nil {
		t.Fatal("Error adding version 2:", err)
	}

	// Test retrieving versions
	v1, err := fvManager.GetVersion("testfile.txt", 1)
	if err != nil {
		t.Fatal("Error retrieving version 1:", err)
	}
	if string(v1.Data) != "Version 1 data" {
		t.Errorf("Expected Version 1 data, got %s", string(v1.Data))
	}

	v2, err := fvManager.GetVersion("testfile.txt", 2)
	if err != nil {
		t.Fatal("Error retrieving version 2:", err)
	}
	if string(v2.Data) != "Version 2 data" {
		t.Errorf("Expected Version 2 data, got %s", string(v2.Data))
	}

	// Test listing versions
	versions, err := fvManager.ListVersions("testfile.txt")
	if err != nil {
		t.Fatal("Error listing versions:", err)
	}

	if len(versions) != 2 {
		t.Fatalf("Expected 2 versions, got %d", len(versions))
	}

	if string(versions[0].Data) != "Version 1 data" {
		t.Errorf("Expected Version 1 data, got %s", string(versions[0].Data))
	}

	if string(versions[1].Data) != "Version 2 data" {
		t.Errorf("Expected Version 2 data, got %s", string(versions[1].Data))
	}
}

func TestConcurrencyWithRWMutex(t *testing.T) {
	fvManager := GetInstance()

	// Adding some initial versions for read consistency
	err := fvManager.AddVersion("concurrencytest.txt", []byte("Initial data"))
	if err != nil {
		t.Fatal("Error adding initial version:", err)
	}

	var wg sync.WaitGroup
	const readGoroutines = 4
	const writeGoroutines = 2

	// Add the correct number of goroutines to the WaitGroup before starting
	wg.Add(readGoroutines + writeGoroutines)

	// Start reading goroutines
	for i := 0; i < readGoroutines; i++ {
		go func() {
			defer wg.Done() // Ensure Done is called when the goroutine finishes

			// Simulate reading a version continuously
			for j := 0; j < 10; j++ { // Limiting iterations to avoid infinite loops
				_, err := fvManager.GetVersion("concurrencytest.txt", 1)
				if err != nil {
					t.Error("Error reading version:", err)
				}
				time.Sleep(time.Millisecond) // Introduce slight delay to simulate work
			}
		}()
	}

	// Start writing goroutines
	for i := 0; i < writeGoroutines; i++ {
		go func() {
			defer wg.Done() // Ensure Done is called when the goroutine finishes

			// Simulate adding new versions continuously
			for j := 0; j < 5; j++ { // Limiting iterations to avoid infinite loops
				err := fvManager.AddVersion("concurrencytest.txt", []byte(fmt.Sprintf("New version %d", time.Now().Nanosecond())))
				if err != nil {
					t.Error("Error adding version:", err)
				}
				time.Sleep(100 * time.Millisecond) // Simulate write time
			}
		}()
	}

	// Allow some time for operations to be carried out
	wg.Wait() // Wait for all goroutines to finish
}

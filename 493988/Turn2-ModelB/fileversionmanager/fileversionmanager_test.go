package fileversionmanager

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFileVersionManagerConcurrency(t *testing.T) {
	fvManager := GetInstance()
	numGoroutines := 100
	numVersions := 1000
	fileName := "testfile.txt"

	// Add multiple versions concurrently
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			for j := 0; j < numVersions/numGoroutines; j++ {
				fvManager.AddVersion(fileName, []byte(fmt.Sprintf("Version %d-%d", i, j)))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	// Verify the total number of versions added
	versions, err := fvManager.ListVersions(fileName)
	if err != nil {
		t.Fatal("Error listing versions:", err)
	}
	if len(versions) != numGoroutines*numVersions {
		t.Errorf("Expected %d versions, got %d", numGoroutines*numVersions, len(versions))
	}

	// Read multiple versions concurrently
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			for j := 0; j < numVersions/numGoroutines; j++ {
				_, err := fvManager.GetVersion(fileName, i*numVersions+j+1)
				if err != nil {
					t.Errorf("Error getting version %d: %v", i*numVersions+j+1, err)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	// Test ListVersions with a large number of versions
	numVersions = 100000
	for i := 0; i < numVersions; i++ {
		fvManager.AddVersion(fileName, []byte(fmt.Sprintf("Version %d", i)))
	}

	start := time.Now()
	versions, _ = fvManager.ListVersions(fileName)
	duration := time.Since(start)

	if len(versions) != numVersions {
		t.Errorf("Expected %d versions, got %d", numVersions, len(versions))
	}

	t.Logf("Time taken to list %d versions: %v", numVersions, duration)
}

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

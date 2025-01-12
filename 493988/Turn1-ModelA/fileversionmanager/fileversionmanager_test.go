package fileversionmanager

import (
	"testing"
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

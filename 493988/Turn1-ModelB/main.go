package main

import (
	"fmt"
	"sync"
)

// FileVersionManager struct
type FileVersionManager struct {
	mu       sync.Mutex
	versions map[string][]string
}

// GetInstance method
func (fvm *FileVersionManager) GetInstance() *FileVersionManager {
	fvm.mu.Lock()
	defer fvm.mu.Unlock()
	if fvm == nil {
		fvm = &FileVersionManager{versions: make(map[string][]string)}
	}
	return fvm
}

// AddVersion method
func (fvm *FileVersionManager) AddVersion(fileName string, version string) {
	fvm.mu.Lock()
	defer fvm.mu.Unlock()
	if _, ok := fvm.versions[fileName]; !ok {
		fvm.versions[fileName] = []string{version}
	} else {
		fvm.versions[fileName] = append(fvm.versions[fileName], version)
	}
}

// GetVersion method
func (fvm *FileVersionManager) GetVersion(fileName string, versionNumber int) string {
	fvm.mu.Lock()
	defer fvm.mu.Unlock()
	if versions, ok := fvm.versions[fileName]; ok {
		if versionNumber >= 1 && versionNumber <= len(versions) {
			return versions[versionNumber-1]
		}
	}
	return ""
}

// ListVersions method
func (fvm *FileVersionManager) ListVersions(fileName string) []string {
	fvm.mu.Lock()
	defer fvm.mu.Unlock()
	if versions, ok := fvm.versions[fileName]; ok {
		return versions
	}
	return []string{}
}

func main() {
	fvm := FileVersionManager{}.GetInstance()
	fvm.AddVersion("example.txt", "1.0")
	fvm.AddVersion("example.txt", "1.1")
	fvm.AddVersion("example.txt", "1.2")
	fvm.AddVersion("report.pdf", "1.0")

	fmt.Println("Versions of example.txt:")
	for _, version := range fvm.ListVersions("example.txt") {
		fmt.Println(version)
	}

	fmt.Println("\nVersion 2 of example.txt:")
	version2 := fvm.GetVersion("example.txt", 2)
	fmt.Println(version2)

	fmt.Println("\nVersions of report.pdf:")
	for _, version := range fvm.ListVersions("report.pdf") {
		fmt.Println(version)
	}
}

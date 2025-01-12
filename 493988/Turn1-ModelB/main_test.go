package main

import (
	"testing"
)

func TestFileVersionManager(t *testing.T) {
	fvm := FileVersionManager{}.GetInstance()
	fvm.AddVersion("file1.txt", "1.0")
	fvm.AddVersion("file1.txt", "1.1")
	fvm.AddVersion("file2.txt", "2.0")

	testCases := []struct {
		fileName         string
		expectedVersions []string
	}{
		{
			fileName:         "file1.txt",
			expectedVersions: []string{"1.0", "1.1"},
		},
		{
			fileName:         "file2.txt",
			expectedVersions: []string{"2.0"},
		},
		{
			fileName:         "file3.txt",
			expectedVersions: []string{},
		},
	}

	for _, testCase := range testCases {
		versions := fvm.ListVersions(testCase.fileName)
		if len(versions) != len(testCase.expectedVersions) {
			t.Errorf("Expected %d versions for %s, but got %d", len(testCase.expectedVersions), testCase.fileName, len(versions))
		}
		for i, version := range versions {
			if version != testCase.expectedVersions[i] {
				t.Errorf("Expected version %d for %s to be %s, but got %s", i+1, testCase.fileName, testCase.expectedVersions[i], version)
			}
		}
	}
}

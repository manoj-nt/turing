package main

import (
	"os"
	"testing"
)

func openFileWithDefer() {
	f, err := os.Open("testfile.txt")
	if err != nil {
		return
	}
	defer f.Close()

	// Simulate some work with the file
}

func openFileManual() {
	f, err := os.Open("testfile.txt")
	if err != nil {
		return
	}

	// Simulate some work with the file

	f.Close()
}

func BenchmarkOpenFileWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		openFileWithDefer()
	}
}

func BenchmarkOpenFileManual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		openFileManual()
	}
}

package main

import (
	"fmt"
	"io"
	"os"
)

// CustomReader simulates reading from a file.
type CustomReader struct {
	data   []byte
	offset int
}

func NewCustomReader(data string) *CustomReader {
	return &CustomReader{data: []byte(data)}
}

func (r *CustomReader) Read(p []byte) (n int, err error) {
	if r.offset >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(p, r.data[r.offset:])
	r.offset += n
	return n, nil
}

// CustomWriter simulates writing to a file.
type CustomWriter struct {
	file *os.File
}

func NewCustomWriter(filename string) (*CustomWriter, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &CustomWriter{file: file}, nil
}

func (w *CustomWriter) Write(p []byte) (n int, err error) {
	return w.file.Write(p)
}

func (w *CustomWriter) Close() error {
	return w.file.Close()
}

func main() {
	// Simulate reading from a file
	reader := NewCustomReader("This is the data to be read and written to a new file.")

	// Simulate writing to a new file
	writer, err := NewCustomWriter("output.txt")
	if err != nil {
		fmt.Println("Error creating writer:", err)
		return
	}
	defer writer.Close()

	// Buffer to hold data temporarily
	buffer := make([]byte, 8)

	// Read from the custom reader and write to the custom writer
	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading:", err)
			break
		}
		if n == 0 {
			break
		}

		if _, err := writer.Write(buffer[:n]); err != nil {
			fmt.Println("Error writing:", err)
			break
		}
	}

	fmt.Println("Data has been successfully copied to output.txt")
}

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Bounded Context: XmlProcessor

// Entity: Document
type Document struct {
	ID       string    `xml:"id"`
	Title    string    `xml:"title"`
	Sections []Section `xml:"section"`
}

// Value Object: Section
type Section struct {
	Name    string `xml:"name"`
	Content string `xml:"content"`
}

// Aggregate Root
type DocumentRepository interface {
	Save(document *Document) error
	FindById(id string) (*Document, error)
}

type InMemoryDocumentRepository struct {
	documents map[string]*Document
}

func NewInMemoryDocumentRepository() *InMemoryDocumentRepository {
	return &InMemoryDocumentRepository{
		documents: make(map[string]*Document),
	}
}

func (r *InMemoryDocumentRepository) Save(document *Document) error {
	r.documents[document.ID] = document
	return nil
}

func (r *InMemoryDocumentRepository) FindById(id string) (*Document, error) {
	return r.documents[id], nil
}

// Domain Service
type XmlDocumentService struct {
	repository DocumentRepository
}

func NewXmlDocumentService(repository DocumentRepository) *XmlDocumentService {
	return &XmlDocumentService{
		repository: repository,
	}
}

func (s *XmlDocumentService) ImportFromFile(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var document Document
	err = xml.Unmarshal(data, &document)
	if err != nil {
		return fmt.Errorf("error unmarshaling XML: %v", err)
	}

	return s.repository.Save(&document)
}

// Infrastructure Service
type FileReader interface {
	ReadFile(path string) ([]byte, error)
}

type DefaultFileReader struct{}

func (f *DefaultFileReader) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// Usage
func main() {
	filePath := "document.xml"

	repository := NewInMemoryDocumentRepository()
	service := NewXmlDocumentService(repository)

	err := service.ImportFromFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc, err := repository.FindById("doc1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Document Title:", doc.Title)
	for _, section := range doc.Sections {
		fmt.Printf("Section Name: %s, Content: %s\n", section.Name, section.Content)
	}
}

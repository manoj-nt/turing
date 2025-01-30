// tools/cmd/generateconfig/main.go
package main  

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/structs"
	"github.com/xeipuuv/gojsonschema"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go generate ./... <json_file_path>")
		os.Exit(1)
	}

	jsonFile := os.Args[1]
	generateStructFromJSON(jsonFile)
}

func generateStructFromJSON(jsonFile string) {
	// Read the JSON file content
	jsonBytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		os.Exit(1)
	}

	// Load the JSON data as an interface{}
	var jsonData interface{}
	if err := json.Unmarshal(jsonBytes, &jsonData); err != nil {
		fmt.Printf("Error parsing JSON data: %s\n", err)
		os.Exit(1)
	}

	// Generate Go code for the struct
	generatedGoCode := generateStructFile("config", "UserPreferences", jsonData)

	// Write the generated code to the file
	outputFile := filepath.Join("..", "config", "auto_generated.go")
	if err := ioutil.WriteFile(outputFile, []byte(generatedGoCode), 0644); err != nil {
		fmt.Printf("Error writing Go file: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Go structs generated successfully.")
}

func generateStructFile(packageName, structName string, jsonData interface{}) string {
	// Create a gojsonschema Loader from the interface
	loader := gojsonschema.NewGoLoader(jsonData)

	// Validate the JSON data
	result, err := gojsonschema.Validate(loader, loader)
	if err != nil {
		fmt.Printf("Error validating JSON data: %s\n", err)
		os.Exit(1)
	}

	// Check if there are any validation errors
	if result.Valid() {
		fmt.Printf("JSON data is valid\n")
	} else {
		fmt.Printf("JSON data is invalid: %s\n", result.Errors())
		os.Exit(1)
	}

	// Get the struct type from the validated data
	strct := structs.New(jsonData)
	structFields := strct.Fields()

	// Build the Go struct code
	var goCodeBuilder strings.Builder
	goCodeBuilder.WriteString(fmt.Sprintf("// Code generated by generateconfig; DO NOT EDIT.\n\npackage %s\n\ntype %s struct {\n", packageName, structName))

	for _, field := range structFields {
		fieldName := strings.Title(field.Name())
		tag := fmt.Sprintf("`json:\"%s\"`", field.Name())
		fieldType := getFieldType(field)
		goCodeBuilder.WriteString(fmt.Sprintf("\t%s %s %s\n", fieldName, fieldType, tag))
	}

	goCodeBuilder.WriteString("}\n")
	return goCodeBuilder.String()
}

func getFieldType(field *structs.Field) string {
	switch field.Kind() {
	case structs.String:
		return "string"
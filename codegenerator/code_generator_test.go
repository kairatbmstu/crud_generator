package codegenerator_test

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"example.com/ast1/codegenerator"
	"example.com/ast1/model"
)

func TestGenerateEntity(t *testing.T) {
	entity := &model.Entity{
		Name: "Student",
		Fields: []model.Field{
			{Name: "ID", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "Age", Type: "int"},
		},
	}
	directory := "/path/to/directory"

	codegenerator.GenerateEntity(directory, entity)

	filePath := directory + "/student.go"
	if err := checkFileExists(filePath); err != nil {
		t.Errorf("Expected file %s to be created, but got error: %v", filePath, err)
	}

	// Additional tests to check the content of the generated file can be added here
}

func TestGenerateDTO(t *testing.T) {
	entity := &model.Entity{
		Name: "Student",
		Fields: []model.Field{
			{Name: "ID", Type: "int"},
			{Name: "Name", Type: "string"},
			{Name: "Age", Type: "int"},
		},
	}
	directory := "/path/to/directory"

	codegenerator.GenerateDTO(directory, entity)

	filePath := directory + "/student_dto.go"
	if err := checkFileExists(filePath); err != nil {
		t.Errorf("Expected file %s to be created, but got error: %v", filePath, err)
	}

	// Additional tests to check the content of the generated file can be added here
}

// Implement similar tests for other code generation functions...

// Helper function to check if a file exists
func checkFileExists(filePath string) error {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
	if err != nil && strings.Contains(err.Error(), "no such file") {
		return err
	}
	return nil
}

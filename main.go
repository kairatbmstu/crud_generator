package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
)

func main() {
	// Create a new file set.
	fset := token.NewFileSet()

	// Define the root of the AST.
	root := &ast.File{
		Name:  ast.NewIdent("main"),
		Decls: []ast.Decl{},
	}

	// Import declarations
	imports := []*ast.ImportSpec{
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"database/sql\"",
			},
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"fmt\"",
			},
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"log\"",
			},
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"github.com/lib/pq\"",
			},
		},
	}

	root.Imports = imports

	// Student struct declaration
	studentStruct := &ast.TypeSpec{
		Name: &ast.Ident{Name: "Student"},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "ID"}}, Type: &ast.Ident{Name: "int"}},
					{Names: []*ast.Ident{{Name: "Name"}}, Type: &ast.Ident{Name: "string"}},
					{Names: []*ast.Ident{{Name: "Age"}}, Type: &ast.Ident{Name: "int"}},
				},
			},
		},
	}

	gendecl1 := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			studentStruct,
		},
	}

	root.Decls = append(root.Decls, gendecl1)

	ast.Print(fset, root)

	// Format and write the AST code to a Go file.
	file, err := os.Create("generated_ast.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = format.Node(file, fset, root)
	if err != nil {
		fmt.Println("Error writing AST to file:", err)
		return
	}

}

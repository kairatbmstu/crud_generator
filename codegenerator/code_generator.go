package codegenerator

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"os"
	"strings"

	"example.com/ast1/model"
)

type CodeGenerator struct {
}

func GenerateCode(entity *model.Entity) {
	// Create a new file set.
	fset := token.NewFileSet()

	// Define the root of the AST.
	root := &ast.File{
		Name:  ast.NewIdent("model"),
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

	astFields := []*ast.Field{}
	for _, field := range entity.Fields {
		astField := ast.Field{
			Names: []*ast.Ident{{Name: field.Name}}, Type: &ast.Ident{Name: string(field.Type)},
		}
		astFields = append(astFields, &astField)
	}

	// Student struct declaration
	modelStruct := &ast.TypeSpec{
		Name: &ast.Ident{Name: entity.Name},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: astFields,
			},
		},
	}

	gendecl1 := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			modelStruct,
		},
	}

	root.Decls = append(root.Decls, gendecl1)

	ast.Print(fset, root)

	// Format and write the AST code to a Go file.
	file, err := os.Create(strings.ToLower(entity.Name) + ".go")
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

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

func GenerateEntity(directory string, entity *model.Entity) {
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
	file, err := os.Create(directory + "/" + strings.ToLower(entity.Name) + ".go")
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

func GenerateRepository(directory string, entity *model.Entity) {
	// Create a new file set.
	fset := token.NewFileSet()

	// Define the root of the AST.
	root := &ast.File{
		Name:  ast.NewIdent("repository"),
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
				Value: "\"example.com/ast1/test/repository\"",
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
	repositoryStruct := &ast.TypeSpec{
		Name: &ast.Ident{Name: entity.Name + "Repository"},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{{Name: "StudentRepository"}}, Type: &ast.Ident{Name: string("repository.StudentRepository")},
					},
				},
			},
		},
	}

	gendecl1 := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			repositoryStruct,
		},
	}

	root.Decls = append(root.Decls, &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0], imports[1], imports[2], imports[3]}})
	root.Decls = append(root.Decls, gendecl1)

	ast.Print(fset, root)

	// Format and write the AST code to a Go file.
	file, err := os.Create(directory + "/" + strings.ToLower(entity.Name) + "_repo.go")
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

func GenerateService(directory string, entity *model.Entity) {
	// Create a new file set.
	fset := token.NewFileSet()

	// Define the root of the AST.
	root := &ast.File{
		Name:  ast.NewIdent("service"),
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
	repositoryStruct := &ast.TypeSpec{
		Name: &ast.Ident{Name: entity.Name + "Service"},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{{Name: "db"}}, Type: &ast.Ident{Name: string("*sql.DB")},
					},
				},
			},
		},
	}

	gendecl1 := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			repositoryStruct,
		},
	}

	root.Decls = append(root.Decls, gendecl1)

	ast.Print(fset, root)

	// Format and write the AST code to a Go file.
	file, err := os.Create(directory + "/" + strings.ToLower(entity.Name) + "_service.go")
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

func GenerateRestApiHandler(directory string, entity *model.Entity) {
	// Create a new file set.
	fset := token.NewFileSet()

	// Define the root of the AST.
	root := &ast.File{
		Name:  ast.NewIdent("handler"),
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
	repositoryStruct := &ast.TypeSpec{
		Name: &ast.Ident{Name: entity.Name + "Handler"},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{{Name: "db"}}, Type: &ast.Ident{Name: string("*sql.DB")},
					},
				},
			},
		},
	}

	gendecl1 := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			repositoryStruct,
		},
	}

	root.Decls = append(root.Decls, &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0], imports[1], imports[2], imports[3]}})
	root.Decls = append(root.Decls, gendecl1)

	ast.Print(fset, root)

	// Format and write the AST code to a Go file.
	file, err := os.Create(directory + "/" + strings.ToLower(entity.Name) + "_handler.go")
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

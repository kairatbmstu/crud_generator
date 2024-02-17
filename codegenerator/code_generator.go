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
		Name:  ast.NewIdent("entity"),
		Decls: []ast.Decl{},
	}

	// Import declarations
	imports := []*ast.ImportSpec{

		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"github.com/google/uuid\"",
			},
		},
	}

	root.Imports = append(root.Imports, imports...)

	astFields := []*ast.Field{}
	for _, field := range entity.Fields {
		astField := ast.Field{
			Names: []*ast.Ident{{Name: strings.ToUpper(field.Name[:1]) + field.Name[1:]}}, Type: &ast.Ident{Name: string(field.Type)},
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

	root.Decls = append(root.Decls, &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0]}})
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

func GenerateDTO(directory string, entity *model.Entity) {
	// Create a new file set.
	fset := token.NewFileSet()

	// Define the root of the AST.
	root := &ast.File{
		Name:  ast.NewIdent("dto"),
		Decls: []ast.Decl{},
	}

	// Import declarations
	imports := []*ast.ImportSpec{

		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"github.com/google/uuid\"",
			},
		},
	}

	root.Imports = append(root.Imports, imports...)

	astFields := []*ast.Field{}
	for _, field := range entity.Fields {
		astField := ast.Field{
			Names: []*ast.Ident{{Name: strings.ToUpper(field.Name[:1]) + field.Name[1:]}}, Type: &ast.Ident{Name: string(field.Type)},
		}
		astFields = append(astFields, &astField)
	}

	// Student struct declaration
	modelStruct := &ast.TypeSpec{
		Name: &ast.Ident{Name: entity.Name + "DTO"},
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

	root.Decls = append(root.Decls, &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0]}})
	root.Decls = append(root.Decls, gendecl1)

	ast.Print(fset, root)

	// Format and write the AST code to a Go file.
	file, err := os.Create(directory + "/" + strings.ToLower(entity.Name) + "_dto.go")
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

func GenerateMapper(directory string, entity *model.Entity) {
	// Create a new file set.
	fset := token.NewFileSet()

	// Define the root of the AST.
	root := &ast.File{
		Name:  ast.NewIdent("mapper"),
		Decls: []ast.Decl{},
	}

	// Import declarations
	imports := []*ast.ImportSpec{
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"example.com/ast1/test/entity\"",
			},
		},
		{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"example.com/ast1/test/dto\"",
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
		Name: &ast.Ident{Name: entity.Name + "Mapper"},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					// &ast.Field{
					// 	Names: []*ast.Ident{{Name: "db"}}, Type: &ast.Ident{Name: string("*sql.DB")},
					// },
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

	root.Decls = append(root.Decls, &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0], imports[1]}})
	root.Decls = append(root.Decls, gendecl1)

	root.Decls = append(root.Decls, GenerateToDTOMethod(entity))
	root.Decls = append(root.Decls, GenerateToEntityMethod(entity))

	ast.Print(fset, root)

	// Format and write the AST code to a Go file.
	file, err := os.Create(directory + "/" + strings.ToLower(entity.Name) + "_mapper.go")
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

func GenerateToDTOMethod(entity *model.Entity) *ast.FuncDecl {
	// Create a field list for the method parameters
	entityVarName := strings.ToLower(entity.Name)
	params := []*ast.Field{
		{
			Names: []*ast.Ident{ast.NewIdent(entityVarName)},
			Type:  &ast.Ident{Name: "*entity." + entity.Name},
		},
	}

	// Create a field list for the method results
	results := []*ast.Field{
		{
			Type: &ast.Ident{Name: "dto." + entity.Name + "DTO"},
		},
	}

	dtoLocalVarName := strings.ToLower(entity.Name) + "DTO"
	dtoTypeName := entity.Name + "DTO"

	fieldMappings := []*ast.AssignStmt{}
	for _, field := range entity.Fields {
		var fieldMapping = &ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(dtoLocalVarName + "." + field.Name)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{ast.NewIdent(entityVarName + "." + field.Name)},
		}
		fieldMappings = append(fieldMappings, fieldMapping)
	}

	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent("var " + dtoLocalVarName)},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{ast.NewIdent("dto." + dtoTypeName + "{}")},
			},
			// &ast.AssignStmt{
			// 	Lhs: []ast.Expr{ast.NewIdent(dtoLocalVarName + ".Id")},
			// 	Tok: token.ASSIGN,
			// 	Rhs: []ast.Expr{ast.NewIdent(entityVarName + ".Id")},
			// },
			// &ast.AssignStmt{
			// 	Lhs: []ast.Expr{ast.NewIdent("_"), ast.NewIdent("err")},
			// 	Tok: token.DEFINE,
			// 	Rhs: []ast.Expr{
			// 		&ast.CallExpr{
			// 			Fun: &ast.SelectorExpr{
			// 				X:   ast.NewIdent("r.db"),
			// 				Sel: ast.NewIdent("Exec"),
			// 			},
			// 			Args: []ast.Expr{
			// 				&ast.BasicLit{Kind: token.STRING, Value: "\"INSERT INTO students (id, name, age) VALUES ($1, $2, $3)\""},
			// 				&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Id")},
			// 				&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Name")},
			// 				&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Age")},
			// 			},
			// 		},
			// 	},
			// },

		},
	}

	for _, fieldMapping := range fieldMappings {
		body.List = append(body.List, fieldMapping)
	}

	body.List = append(body.List, &ast.ReturnStmt{
		Results: []ast.Expr{ast.NewIdent(dtoLocalVarName)},
	})

	// Create a function declaration for the example method
	exampleMethod := &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent("r")},
					Type:  &ast.StarExpr{X: &ast.Ident{Name: entity.Name + "Mapper"}},
				},
			},
		},
		Name: ast.NewIdent("ToDTO"),
		Type: &ast.FuncType{
			Params:  &ast.FieldList{List: params},
			Results: &ast.FieldList{List: results},
		},
		Body: body,
	}

	return exampleMethod
}

func GenerateToEntityMethod(entity *model.Entity) *ast.FuncDecl {
	// Create a field list for the method parameters
	dtoLocalVarName := strings.ToLower(entity.Name) + "DTO"
	dtoTypeName := "*dto." + entity.Name + "DTO"
	entityLocalVarName := strings.ToLower(entity.Name)
	params := []*ast.Field{
		{
			Names: []*ast.Ident{ast.NewIdent(dtoLocalVarName)},
			Type:  &ast.Ident{Name: dtoTypeName},
		},
	}

	// Create a field list for the method results
	results := []*ast.Field{
		{
			Type: &ast.Ident{Name: "entity." + entity.Name},
		},
	}

	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent("var " + entityLocalVarName)},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{ast.NewIdent("entity." + entity.Name + "{}")},
			},
			// &ast.AssignStmt{
			// 	Lhs: []ast.Expr{ast.NewIdent("_"), ast.NewIdent("err")},
			// 	Tok: token.DEFINE,
			// 	Rhs: []ast.Expr{
			// 		&ast.CallExpr{
			// 			Fun: &ast.SelectorExpr{
			// 				X:   ast.NewIdent("r.db"),
			// 				Sel: ast.NewIdent("Exec"),
			// 			},
			// 			Args: []ast.Expr{
			// 				&ast.BasicLit{Kind: token.STRING, Value: "\"INSERT INTO students (id, name, age) VALUES ($1, $2, $3)\""},
			// 				&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Id")},
			// 				&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Name")},
			// 				&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Age")},
			// 			},
			// 		},
			// 	},
			// },
			&ast.ReturnStmt{
				Results: []ast.Expr{ast.NewIdent(entityLocalVarName)},
			},
		},
	}

	// Create a function declaration for the example method
	exampleMethod := &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent("r")},
					Type:  &ast.StarExpr{X: &ast.Ident{Name: entity.Name + "Mapper"}},
				},
			},
		},
		Name: ast.NewIdent("ToEntity"),
		Type: &ast.FuncType{
			Params:  &ast.FieldList{List: params},
			Results: &ast.FieldList{List: results},
		},
		Body: body,
	}

	return exampleMethod
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
				Value: "\"example.com/ast1/test/entity\"",
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

	root.Decls = append(root.Decls, &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0], imports[1], imports[2]}})
	root.Decls = append(root.Decls, gendecl1)

	root.Decls = append(root.Decls, GenerateCreateMethod(entity))
	root.Decls = append(root.Decls, GenerateUpdateMethod(entity))
	root.Decls = append(root.Decls, GenerateDeleteMethod(entity))

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

func GenerateCreateMethod(entity *model.Entity) *ast.FuncDecl {
	// Create a field list for the method parameters
	params := []*ast.Field{
		{
			Names: []*ast.Ident{ast.NewIdent(strings.ToLower(entity.Name))},
			Type:  &ast.Ident{Name: "*entity." + entity.Name},
		},
	}

	// Create a field list for the method results
	results := []*ast.Field{
		{
			Type: &ast.Ident{Name: "error"},
		},
	}

	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent("_"), ast.NewIdent("err")},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("r.db"),
							Sel: ast.NewIdent("Exec"),
						},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: "\"INSERT INTO students (id, name, age) VALUES ($1, $2, $3)\""},
							&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Id")},
							&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Name")},
							&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Age")},
						},
					},
				},
			},
			&ast.ReturnStmt{
				Results: []ast.Expr{ast.NewIdent("err")},
			},
		},
	}

	// Create a function declaration for the example method
	exampleMethod := &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent("r")},
					Type:  &ast.StarExpr{X: &ast.Ident{Name: entity.Name + "Repository"}},
				},
			},
		},
		Name: ast.NewIdent("Create"),
		Type: &ast.FuncType{
			Params:  &ast.FieldList{List: params},
			Results: &ast.FieldList{List: results},
		},
		Body: body,
	}

	return exampleMethod
}

func GenerateUpdateMethod(entity *model.Entity) *ast.FuncDecl {
	// Create a field list for the method parameters
	params := []*ast.Field{
		{
			Names: []*ast.Ident{ast.NewIdent(strings.ToLower(entity.Name))},
			Type:  &ast.Ident{Name: "*entity." + entity.Name},
		},
	}

	// Create a field list for the method results
	results := []*ast.Field{
		{
			Type: &ast.Ident{Name: "error"},
		},
	}

	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent("_"), ast.NewIdent("err")},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("r.db"),
							Sel: ast.NewIdent("Exec"),
						},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: "\"UPDATE students SET name = $2, age = $3 WHERE id = $1\""},
							&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Id")},
							&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Name")},
							&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Age")},
						},
					},
				},
			},
			&ast.ReturnStmt{
				Results: []ast.Expr{ast.NewIdent("err")},
			},
		},
	}

	// Create a function declaration for the example method
	exampleMethod := &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent("r")},
					Type:  &ast.StarExpr{X: &ast.Ident{Name: entity.Name + "Repository"}},
				},
			},
		},
		Name: ast.NewIdent("Update"),
		Type: &ast.FuncType{
			Params:  &ast.FieldList{List: params},
			Results: &ast.FieldList{List: results},
		},
		Body: body,
	}

	return exampleMethod
}

func GenerateDeleteMethod(entity *model.Entity) *ast.FuncDecl {
	// Create a field list for the method parameters
	params := []*ast.Field{
		{
			Names: []*ast.Ident{ast.NewIdent(strings.ToLower(entity.Name))},
			Type:  &ast.Ident{Name: "*entity." + entity.Name},
		},
	}

	// Create a field list for the method results
	results := []*ast.Field{
		{
			Type: &ast.Ident{Name: "error"},
		},
	}

	body := &ast.BlockStmt{
		List: []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent("_"), ast.NewIdent("err")},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("r.db"),
							Sel: ast.NewIdent("Exec"),
						},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: "\"DELETE students  WHERE id = $1\""},
							&ast.SelectorExpr{X: ast.NewIdent("student"), Sel: ast.NewIdent("Id")},
						},
					},
				},
			},
			&ast.ReturnStmt{
				Results: []ast.Expr{ast.NewIdent("err")},
			},
		},
	}

	// Create a function declaration for the example method
	exampleMethod := &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent("r")},
					Type:  &ast.StarExpr{X: &ast.Ident{Name: entity.Name + "Repository"}},
				},
			},
		},
		Name: ast.NewIdent("Delete"),
		Type: &ast.FuncType{
			Params:  &ast.FieldList{List: params},
			Results: &ast.FieldList{List: results},
		},
		Body: body,
	}

	return exampleMethod
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
		Name: &ast.Ident{Name: entity.Name + "Service"},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					&ast.Field{
						Names: []*ast.Ident{{Name: entity.Name + "Repository"}}, Type: &ast.Ident{Name: string("repository." + entity.Name + "Repository")},
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

	root.Decls = append(root.Decls, &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0]}})
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
				Value: "\"example.com/ast1/test/service\"",
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
					{
						Names: []*ast.Ident{{Name: entity.Name + "Service"}}, Type: &ast.Ident{Name: "service." + entity.Name + "Service"},
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

	root.Decls = append(root.Decls, &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0]}})
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

package example

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet()

	// Package declaration
	pkg := &ast.Package{
		Name: "main",
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

	// StudentRepository struct declaration
	studentRepoStruct := &ast.TypeSpec{
		Name: &ast.Ident{Name: "StudentRepository"},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "db"}}, Type: &ast.SelectorExpr{X: &ast.Ident{Name: "sql"}, Sel: &ast.Ident{Name: "DB"}}},
				},
			},
		},
	}

	// NewStudentRepository function declaration
	newStudentRepoFunc := &ast.FuncDecl{
		Name: &ast.Ident{Name: "NewStudentRepository"},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{{Names: []*ast.Ident{{Name: "dataSourceName"}}, Type: &ast.Ident{Name: "string"}}},
			},
			Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.StarExpr{X: &ast.Ident{Name: "StudentRepository"}}, Names: []*ast.Ident{{Name: ""}}}}},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "db"}, &ast.Ident{Name: "err"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "sql"}, Sel: &ast.Ident{Name: "Open"}},
							Args: []ast.Expr{
								&ast.BasicLit{Kind: token.STRING, Value: "\"postgres\""},
								&ast.Ident{Name: "dataSourceName"},
							},
						},
					},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  &ast.Ident{Name: "err"},
						Op: token.NEQ,
						Y:  &ast.Ident{Name: "nil"},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "nil"}, &ast.Ident{Name: "err"}}},
						},
					},
				},
				&ast.ReturnStmt{Results: []ast.Expr{&ast.UnaryExpr{Op: token.AND, X: &ast.CompositeLit{Type: &ast.Ident{Name: "StudentRepository"}, Elts: []ast.Expr{{&ast.Ident{Name: "db"}}}}}, &ast.Ident{Name: "nil"}}},
			},
		},
	}

	// Main function
	mainFunc := &ast.FuncDecl{
		Name: &ast.Ident{Name: "main"},
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				// Replace the dataSourceName with your PostgreSQL connection string
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "dataSourceName"}, &ast.Ident{Name: "err"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"user=yourusername password=yourpassword dbname=yourdbname sslmode=disable\""}},
				},
				&ast.IfStmt{
					Init: &ast.AssignStmt{
						Lhs: []ast.Expr{&ast.Ident{Name: "repo"}, &ast.Ident{Name: "err"}},
						Tok: token.DEFINE,
						Rhs: []ast.Expr{&ast.CallExpr{
							Fun: &ast.Ident{Name: "NewStudentRepository"},
							Args: []ast.Expr{
								&ast.Ident{Name: "dataSourceName"},
							},
						}},
					},
					Cond: &ast.BinaryExpr{
						X:  &ast.Ident{Name: "err"},
						Op: token.NEQ,
						Y:  &ast.Ident{Name: "nil"},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ExprStmt{X: &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "log"}, Sel: &ast.Ident{Name: "Fatal"}}, Args: []ast.Expr{&ast.Ident{Name: "err"}}}},
						},
					},
				},
				&ast.DeferStmt{Call: &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "repo"}, Sel: &ast.Ident{Name: "Close"}}}},
				// Example usage
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "student"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{&ast.UnaryExpr{
						Op: token.AND,
						X: &ast.CompositeLit{
							Type: &ast.Ident{Name: "Student"},
							Elts: []ast.Expr{
								&ast.KeyValueExpr{Key: &ast.Ident{Name: "Name"}, Value: &ast.BasicLit{Kind: token.STRING, Value: "\"John Doe\""}},
								&ast.KeyValueExpr{Key: &ast.Ident{Name: "Age"}, Value: &ast.BasicLit{Kind: token.INT, Value: "20"}},
							},
						},
					}},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  &ast.Ident{Name: "err"},
						Op: token.NEQ,
						Y:  &ast.Ident{Name: "nil"},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ExprStmt{X: &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "log"}, Sel: &ast.Ident{Name: "Fatal"}}, Args: []ast.Expr{&ast.Ident{Name: "err"}}}},
						},
					},
				},
				// Querying data
				&ast.AssignStmt{
					Lhs: []ast.Expr{&ast.Ident{Name: "studentsByName"}, &ast.Ident{Name: "err"}},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "repo"}, Sel: &ast.Ident{Name: "FindByName"}}, Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"John Doe\""}}}},
				},
				&ast.IfStmt{
					Cond: &ast.BinaryExpr{
						X:  &ast.Ident{Name: "err"},
						Op: token.NEQ,
						Y:  &ast.Ident{Name: "nil"},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.ExprStmt{X: &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "log"}, Sel: &ast.Ident{Name: "Fatal"}}, Args: []ast.Expr{&ast.Ident{Name: "err"}}}},
						},
					},
				},
				&ast.ExprStmt{X: &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "fmt"}, Sel: &ast.Ident{Name: "Println"}}, Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING, Value: "\"Students with name John Doe:\""},
					&ast.Ident{Name: "studentsByName"},
				}}},
			},
		},
	}

	// File
	file := &ast.File{
		Name: &ast.Ident{Name: "main"},
		Decls: []ast.Decl{
			&ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{imports[0], imports[1], imports[2], imports[3]}},
			&ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{studentStruct, studentRepoStruct}},
			&ast.FuncDecl{Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "*StudentRepository"}}}}, Name: &ast.Ident{Name: "Close"}, Type: &ast.FuncType{}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "db"}}, Args: nil}}}}}},
			&ast.FuncDecl{Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "*StudentRepository"}}}}, Name: &ast.Ident{Name: "Create"}, Type: &ast.FuncType{Params: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "s"}}, Type: &ast.StarExpr{X: &ast.Ident{Name: "Student"}}}}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.AssignStmt{Lhs: []ast.Expr{&ast.BlankIdent{}, &ast.Ident{Name: "err"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "db"}}, Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"INSERT INTO students (name, age) VALUES ($1, $2)\""}, &ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "Name"}}, &ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "Age"}}}}}}, &ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "err"}}}}}},
			&ast.FuncDecl{Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "*StudentRepository"}}}}, Name: &ast.Ident{Name: "Update"}, Type: &ast.FuncType{Params: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "id"}}, Type: &ast.Ident{Name: "int"}}, {Names: []*ast.Ident{{Name: "s"}}, Type: &ast.StarExpr{X: &ast.Ident{Name: "Student"}}}}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.AssignStmt{Lhs: []ast.Expr{&ast.BlankIdent{}, &ast.Ident{Name: "err"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "db"}}, Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"UPDATE students SET name=$1, age=$2 WHERE id=$3\""}, &ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "Name"}}, &ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "Age"}}, &ast.Ident{Name: "id"}}}}}, &ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "err"}}}}}},
			&ast.FuncDecl{Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "*StudentRepository"}}}}, Name: &ast.Ident{Name: "Delete"}, Type: &ast.FuncType{Params: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "id"}}, Type: &ast.Ident{Name: "int"}}}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.AssignStmt{Lhs: []ast.Expr{&ast.BlankIdent{}, &ast.Ident{Name: "err"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "db"}}, Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"DELETE FROM students WHERE id=$1\""}, &ast.Ident{Name: "id"}}}}}, &ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "err"}}}}}},
			&ast.FuncDecl{Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "*StudentRepository"}}}}, Name: &ast.Ident{Name: "FindByID"}, Type: &ast.FuncType{Params: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "id"}}, Type: &ast.Ident{Name: "int"}}}}, Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.StarExpr{X: &ast.Ident{Name: "Student"}}}, {Type: &ast.Ident{Name: "error"}}}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.DeclStmt{Decl: &ast.GenDecl{Tok: token.VAR, Specs: []ast.Spec{&ast.ValueSpec{Names: []*ast.Ident{{Name: "s"}}, Values: []ast.Expr{&ast.CompositeLit{Type: &ast.Ident{Name: "Student"}, Elts: []ast.Expr{{&ast.KeyValueExpr{Key: &ast.Ident{Name: "ID"}, Value: &ast.Ident{Name: "id"}}}}}}}}}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.Ident{Name: "err"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "db"}}, Sel: &ast.Ident{Name: "QueryRow"}}, Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"SELECT name, age FROM students WHERE id=$1\""}, &ast.Ident{Name: "id"}}}}}, &ast.IfStmt{Cond: &ast.BinaryExpr{X: &ast.Ident{Name: "err"}, Op: token.NEQ, Y: &ast.Ident{Name: "nil"}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "nil"}, &ast.Ident{Name: "err"}}}}}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "ID"}}}, Tok: token.ASSIGN, Rhs: []ast.Expr{&ast.Ident{Name: "id"}}}, &ast.ReturnStmt{Results: []ast.Expr{&ast.UnaryExpr{Op: token.AND, X: &ast.Ident{Name: "s"}}, &ast.Ident{Name: "nil"}}}}}},
			&ast.FuncDecl{Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "*StudentRepository"}}}}, Name: &ast.Ident{Name: "FindByName"}, Type: &ast.FuncType{Params: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "name"}}, Type: &ast.Ident{Name: "string"}}}}, Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.StarExpr{X: &ast.Ident{Name: "Student"}}}, {Type: &ast.Ident{Name: "error"}}}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.AssignStmt{Lhs: []ast.Expr{&ast.Ident{Name: "rows"}, &ast.Ident{Name: "err"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "db"}}, Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"SELECT id, age FROM students WHERE name=$1\""}, &ast.Ident{Name: "name"}}}}}}, &ast.IfStmt{Cond: &ast.BinaryExpr{X: &ast.Ident{Name: "err"}, Op: token.NEQ, Y: &ast.Ident{Name: "nil"}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "nil"}, &ast.Ident{Name: "err"}}}}}}, &ast.DeferStmt{Call: &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "Close"}}, Args: nil}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.Ident{Name: "students"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CompositeLit{Type: &ast.ArrayType{Elt: &ast.StarExpr{X: &ast.Ident{Name: "Student"}}}, Elts: []ast.Expr{}}}}, &ast.ForStmt{Init: &ast.AssignStmt{Lhs: []ast.Expr{&ast.BlankIdent{}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "Next"}}, Args: nil}}}, Cond: &ast.UnaryExpr{Op: token.NOT, X: &ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "Next"}}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.AssignStmt{Lhs: []ast.Expr{&ast.Ident{Name: "s"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CompositeLit{Type: &ast.Ident{Name: "Student"}, Elts: []ast.Expr{{&ast.KeyValueExpr{Key: &ast.Ident{Name: "Name"}, Value: &ast.Ident{Name: "name"}}}}}}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "ID"}}}, Tok: token.ASSIGN, Rhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "ID"}}}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "Age"}}}, Tok: token.ASSIGN, Rhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "Age"}}}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "Name"}}}, Tok: token.ASSIGN, Rhs: []ast.Expr{&ast.Ident{Name: "name"}}}}, Post: nil}}, &ast.IfStmt{Cond: &ast.BinaryExpr{X: &ast.Ident{Name: "err"}, Op: token.NEQ, Y: &ast.Ident{Name: "nil"}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "nil"}, &ast.Ident{Name: "err"}}}}}}, &ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "students"}, &ast.Ident{Name: "nil"}}}}}},
			&ast.FuncDecl{Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "r"}}, Type: &ast.Ident{Name: "*StudentRepository"}}}}, Name: &ast.Ident{Name: "FindByAge"}, Type: &ast.FuncType{Params: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "age"}}, Type: &ast.Ident{Name: "int"}}}}, Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.StarExpr{X: &ast.Ident{Name: "Student"}}}, {Type: &ast.Ident{Name: "error"}}}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.AssignStmt{Lhs: []ast.Expr{&ast.Ident{Name: "rows"}, &ast.Ident{Name: "err"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "db"}}, Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: "\"SELECT id, name FROM students WHERE age=$1\""}, &ast.Ident{Name: "age"}}}}}}, &ast.IfStmt{Cond: &ast.BinaryExpr{X: &ast.Ident{Name: "err"}, Op: token.NEQ, Y: &ast.Ident{Name: "nil"}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "nil"}, &ast.Ident{Name: "err"}}}}}}, &ast.DeferStmt{Call: &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "Close"}}, Args: nil}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.Ident{Name: "students"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CompositeLit{Type: &ast.ArrayType{Elt: &ast.StarExpr{X: &ast.Ident{Name: "Student"}}}, Elts: []ast.Expr{}}}}, &ast.ForStmt{Init: &ast.AssignStmt{Lhs: []ast.Expr{&ast.BlankIdent{}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "Next"}}, Args: nil}}}, Cond: &ast.UnaryExpr{Op: token.NOT, X: &ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "Next"}}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.AssignStmt{Lhs: []ast.Expr{&ast.Ident{Name: "s"}}, Tok: token.DEFINE, Rhs: []ast.Expr{&ast.CompositeLit{Type: &ast.Ident{Name: "Student"}, Elts: []ast.Expr{{&ast.KeyValueExpr{Key: &ast.Ident{Name: "Age"}, Value: &ast.Ident{Name: "age"}}}}}}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "ID"}}}, Tok: token.ASSIGN, Rhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "ID"}}}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "Name"}}}, Tok: token.ASSIGN, Rhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "rows"}, Sel: &ast.Ident{Name: "Name"}}}}, &ast.AssignStmt{Lhs: []ast.Expr{&ast.SelectorExpr{X: &ast.Ident{Name: "s"}, Sel: &ast.Ident{Name: "Age"}}}, Tok: token.ASSIGN, Rhs: []ast.Expr{&ast.Ident{Name: "age"}}}}, Post: nil}}, &ast.IfStmt{Cond: &ast.BinaryExpr{X: &ast.Ident{Name: "err"}, Op: token.NEQ, Y: &ast.Ident{Name: "nil"}}, Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "nil"}, &ast.Ident{Name: "err"}}}}}}, &ast.ReturnStmt{Results: []ast.Expr{&ast.Ident{Name: "students"}, &ast.Ident{Name: "nil"}}}}}},
			newStudentRepoFunc,
			mainFunc,
		},
	}

	ast.SortImports(fset, file)
	err := format.Node(os.Stdout, fset, file)
	if err != nil {
		fmt.Println(err)
		return
	}
}

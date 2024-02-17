package codegenerator

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"example.com/ast1/model"
)

type EntityManager struct {
	Entity *model.Entity
}

func NewEntityManager(entity *model.Entity) *EntityManager {
	var entityManager = EntityManager{
		Entity: entity,
	}

	return &entityManager
}

func (e EntityManager) TableName() string {
	return e.Entity.Name.Lower() + "s"
}

func (e EntityManager) TableFields() string {
	var result = ""
	for i := 0; i < len(e.Entity.Fields); i++ {
		result += strings.ToLower(e.Entity.Fields[i].Name)
		if i < len(e.Entity.Fields)-1 {
			result += ","
		}
	}
	return result
}

func (e EntityManager) InsertParameters() string {
	var result = ""
	for i := 1; i <= len(e.Entity.Fields); i++ {
		result += "$" + strconv.Itoa(i)
		if i < len(e.Entity.Fields) {
			result += ","
		}
	}
	return result
}

func (e EntityManager) UpdateParameters() string {
	var result = " set "
	for i := 0; i < len(e.Entity.Fields); i++ {
		if e.Entity.Fields[i].ColumnName == "id" {
			continue
		}
		result += e.Entity.Fields[i].ColumnName + " = " + "$" + strconv.Itoa(i+1)
		if i < len(e.Entity.Fields) {
			result += ","
		}
	}
	return result
}

func (e EntityManager) InsertSql() []ast.Expr {
	result := []ast.Expr{
		&ast.BasicLit{Kind: token.STRING, Value: "\"INSERT INTO " + e.TableName() + " (" +
			e.TableFields() + ") VALUES (" + e.InsertParameters() + ")\""},
	}
	for _, field := range e.Entity.Fields {
		result = append(result, &ast.SelectorExpr{X: ast.NewIdent(e.Entity.Name.Lower()), Sel: ast.NewIdent(field.Name)})
	}
	return result
}

func (e EntityManager) UpdateSql() []ast.Expr {
	result := []ast.Expr{
		&ast.BasicLit{Kind: token.STRING, Value: "\"UPDATE  " + e.TableName() + e.UpdateParameters() +
			" WHERE id = $1\""},
	}
	for _, field := range e.Entity.Fields {
		result = append(result, &ast.SelectorExpr{X: ast.NewIdent(e.Entity.Name.Lower()), Sel: ast.NewIdent(field.Name)})
	}
	return result
}

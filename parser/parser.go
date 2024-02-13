package main

import (
	"fmt"

	"example.com/ast1/model"
)

func ParseFile(file string) model.Model {
	return model.Model{}
}

const (
	TokenType_Keyword     = "keyword"
	TokenType_Parenthesis = "parenthesis"
	TokenType_Brace       = "brace"
	TokenType_Identifier  = "identifier"
)

type TokenType string

type Token struct {
	Value     string
	TokenType TokenType
}

func main() {
	fmt.Println("Hello World!")
}

func ParseModel() model.Model {
	return model.Model{}
}

func ParseEntity() []model.Entity {
	return []model.Entity{}
}

func ParseField() []model.Field {
	return []model.Field{}
}

func ParseRelationshipGroup() []model.Relationship {
	return []model.Relationship{}
}

func ParseRelationship() model.Relationship {
	return model.Relationship{}
}

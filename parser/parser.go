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

const (
	Keyword_Entity       = "entity"
	Keyword_Relationship = "relationship"
)

type TokenType string

type Token struct {
	Value     string
	TokenType TokenType
}

func main() {
	fmt.Println("Hello World!")
}

func ParseModel(tokens []Token) model.Model {
	for index, token := range tokens {
		switch token.TokenType {
		case TokenType_Keyword:
			if token.Value == Keyword_Entity {
				ParseEntity(&index, &tokens)
			}

			if token.Value == Keyword_Relationship {
				ParseEntity(&index, &tokens)
			}
		case TokenType_Parenthesis:
		case TokenType_Brace:
		case TokenType_Identifier:
		}
	}
	return model.Model{}
}

func ParseEntity(index *int, tokens *[]Token) []model.Entity {
	return []model.Entity{}
}

func ParseField(index *int, tokens *[]Token) []model.Field {
	return []model.Field{}
}

func ParseRelationshipGroup(index *int, tokens *[]Token) []model.Relationship {
	return []model.Relationship{}
}

func ParseRelationship(index *int, tokens *[]Token) model.Relationship {
	return model.Relationship{}
}

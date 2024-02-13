package main

import (
	"fmt"
	"io/ioutil"
	"strings"

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
	Keyword_Paginate     = "paginate"
)

type TokenType string

type Token struct {
	Value     string
	TokenType TokenType
}

func main() {
	data, err := ioutil.ReadFile("example.jdl")
	if err != nil {
		panic(err)
	}
	jdlText := string(data)

	tokens := strings.FieldsFunc(jdlText, func(r rune) bool {
		return r == ' ' || r == '{' || r == '}'
	})

	for _, token := range tokens {
		fmt.Println(strings.TrimSpace(token))
	}
	//fmt.Println(tokens)
}

func ParseModel(tokens []Token) model.Model {
	var entities = []model.Entity{}
	var relationships = []model.Relationship{}
	var paginates = []model.Paginate{}

	for index, token := range tokens {
		switch token.TokenType {
		case TokenType_Keyword:
			if token.Value == Keyword_Entity {
				entity := ParseEntity(&index, &tokens)
				entities = append(entities, *entity)
			}

			if token.Value == Keyword_Relationship {
				relationship := ParseRelationship(&index, &tokens)
				relationships = append(relationships, *relationship)
			}

			if token.Value == Keyword_Paginate {
				paginate := ParsePaginate(&index, &tokens)
				paginates = append(paginates, *paginate)
			}
		case TokenType_Parenthesis:
		case TokenType_Brace:
		case TokenType_Identifier:
		}
	}
	return model.Model{}
}

func ParseEntity(index *int, tokens *[]Token) *model.Entity {
	*index++

	return &model.Entity{}
}

func ParseField(index *int, tokens *[]Token) *model.Field {
	return &model.Field{}
}

func ParseRelationshipGroup(index *int, tokens *[]Token) []model.Relationship {
	return []model.Relationship{}
}

func ParseRelationship(index *int, tokens *[]Token) *model.Relationship {
	return &model.Relationship{}
}

func ParsePaginate(index *int, tokens *[]Token) *model.Paginate {
	return &model.Paginate{}
}

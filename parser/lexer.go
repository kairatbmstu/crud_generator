package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"example.com/ast1/model"
)

func ParseFile(file string) model.Model {
	return model.Model{}
}

const (
	TokenType_Undifined   = "undefined"
	TokenType_Keyword     = "keyword"
	TokenType_Parenthesis = "parenthesis"
	TokenType_Brace       = "brace"
	TokenType_Identifier  = "identifier"
)

const (
	Entity       = "entity"
	Relationship = "relationship"
	Paginate     = "paginate"
)

type DataType int

const (
	TypeUndefined   DataType = 0
	TypeText        DataType = 1
	TypePunctuation DataType = 2
	TypeWhiteSpace  DataType = 3
)

type TokenType string

type Token struct {
	Value     string
	TokenType TokenType
	DataType  DataType
}

func main() {
	data, err := ioutil.ReadFile("example.jdl")
	if err != nil {
		panic(err)
	}
	jdlText := string(data)

	tokens, err := Tokenize(jdlText)

	if err != nil {
		panic(err)
	}

	err = LexicalAnalysis(tokens)
	if err != nil {
		panic(err)
	}

	model, err := ParseModel(tokens)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
}

func LexicalAnalysis(tokens *[]Token) error {
	for i, _ := range *tokens {
		if (*tokens)[i].DataType == TypePunctuation {
			if (*tokens)[i].Value == "{" || (*tokens)[i].Value == "}" {
				(*tokens)[i].TokenType = TokenType_Parenthesis
			} else if (*tokens)[i].Value == "(" || (*tokens)[i].Value == ")" {
				(*tokens)[i].TokenType = TokenType_Brace
			}
		} else if (*tokens)[i].DataType == TypeText {
			if (*tokens)[i].Value == "entity" || (*tokens)[i].Value == "relationship" ||
				(*tokens)[i].Value == "String" || (*tokens)[i].Value == "UUID" ||
				(*tokens)[i].Value == "Integer" || (*tokens)[i].Value == "" ||
				(*tokens)[i].Value == "paginate" || (*tokens)[i].Value == "pagination" ||
				(*tokens)[i].Value == "with" || (*tokens)[i].Value == "to" {
				(*tokens)[i].TokenType = TokenType_Keyword
			} else {
				(*tokens)[i].TokenType = TokenType_Identifier
			}
		} else {
			return errors.New("Undefined type found : " + (*tokens)[i].Value)
		}
	}

	return nil
}

func Tokenize(text string) (*[]Token, error) {
	tokens := []Token{}

	buffer := []rune{}
	prevDataType := TypeUndefined
	for _, currentChar := range text {
		if currentChar >= '0' && currentChar <= '9' {
			return nil, errors.New("numbers are not allowed")
		}

		if (currentChar >= 'a' && currentChar <= 'z') || (currentChar >= 'A' && currentChar <= 'Z') {
			switch prevDataType {
			case TypeUndefined:
				buffer = []rune{}
				buffer = append(buffer, currentChar)
				break
			case TypeText:
				buffer = append(buffer, currentChar)
				break
			case TypeWhiteSpace:
				buffer = []rune{}
				buffer = append(buffer, currentChar)
				break
			case TypePunctuation:
				if len(buffer) > 0 {
					token := Token{
						DataType: TypePunctuation,
						Value:    string(buffer),
					}
					tokens = append(tokens, token)
					buffer = []rune{}
				}
				buffer = []rune{}
				buffer = append(buffer, currentChar)
				break
			}
			prevDataType = TypeText
		} else if currentChar == '{' || currentChar == '}' || currentChar == '(' || currentChar == ')' {
			switch prevDataType {
			case TypeUndefined:
				buffer = append(buffer, currentChar)
				token := Token{
					DataType: TypePunctuation,
					Value:    string(buffer),
				}
				tokens = append(tokens, token)
				buffer = []rune{}
				break
			case TypeText:
				if len(buffer) > 0 {
					token := Token{
						DataType: TypeText,
						Value:    string(buffer),
					}
					tokens = append(tokens, token)
					buffer = []rune{}
				}
				buffer = append(buffer, currentChar)
				break
			case TypeWhiteSpace:
				buffer = []rune{}
				buffer = append(buffer, currentChar)
				token := Token{
					DataType: TypePunctuation,
					Value:    string(buffer),
				}
				tokens = append(tokens, token)
				buffer = []rune{}
				break
			case TypePunctuation:
				if len(buffer) > 0 {
					token := Token{
						DataType: TypePunctuation,
						Value:    string(buffer),
					}
					tokens = append(tokens, token)
					buffer = []rune{}
				}

				break
			}
			prevDataType = TypePunctuation
		} else if currentChar == ' ' || currentChar == '\n' || currentChar == '\t' {
			switch prevDataType {
			case TypeUndefined:
				break
			case TypeText:
				if len(buffer) > 0 {
					token := Token{
						DataType: TypeText,
						Value:    string(buffer),
					}
					tokens = append(tokens, token)
					buffer = []rune{}
				}
				break
			case TypeWhiteSpace:
				break
			case TypePunctuation:
				if len(buffer) > 0 {
					token := Token{
						DataType: TypePunctuation,
						Value:    string(buffer),
					}
					tokens = append(tokens, token)
				}
				break
			}
			prevDataType = TypeWhiteSpace
		}

	}
	return &tokens, nil
}

func ParseModel(tokens *[]Token) (*model.Model, error) {
	var entities = []model.Entity{}
	var relationships = []model.Relationship{}
	var paginates = []model.Paginate{}

	for index := 0; index < len(*tokens); index++ {
		if (*tokens)[index].TokenType == TokenType_Keyword {
			if (*tokens)[index].Value == Entity {
				entity, err := ParseEntity(&index, tokens)
				if err != nil {
					return nil, err
				}
				entities = append(entities, *entity)
			}

			if (*tokens)[index].Value == Relationship {
				relationship := ParseRelationship(&index, tokens)
				relationships = append(relationships, *relationship)
			}

			if (*tokens)[index].Value == Paginate {
				paginate := ParsePaginate(&index, tokens)
				paginates = append(paginates, *paginate)
			}
		}

	}
	return &model.Model{
		Entities:      entities,
		Relationships: relationships,
	}, nil
}

func ParseEntity(index *int, tokens *[]Token) (*model.Entity, error) {
	*index++
	identifierName := (*tokens)[*index]
	*index++
	openParenthesis := (*tokens)[*index]
	entity := model.Entity{}
	entity.Name = identifierName.Value

	if openParenthesis.Value != "{" {
		return nil, errors.New("Open parenthesis should have been found { but " + openParenthesis.Value + " was found")
	}

	fields, err := ParseFields(index, tokens)

	if err != nil {
		return nil, err
	}

	entity.Fields = *fields

	*index++
	closeParenthesis := (*tokens)[*index]
	if closeParenthesis.Value != "{" {
		return nil, errors.New("{ should have been found but " + closeParenthesis.Value + " was found")
	}

	ParseFields(index, tokens)

	return &entity, nil
}

func ParseFields(index *int, tokens *[]Token) (*[]model.Field, error) {
	fields := []model.Field{}
	for i := *index; i < len((*tokens)); i++ {
		field, err := ParseField(index, tokens)
		if err != nil {
			return nil, err
		}
		fields = append(fields, *field)
		// if (*tokens)[i].Value != "," {
		// 	return nil, errors.New(", should have been found  but " + (*tokens)[i].Value + " was found")
		// }
	}
	return &fields, nil
}

func ParseField(index *int, tokens *[]Token) (*model.Field, error) {
	return &model.Field{}, nil
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

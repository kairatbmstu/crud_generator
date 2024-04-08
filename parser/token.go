package parser

const (
	ILLEGAL = iota
	KEYWORD
	LCBRACE
	RCBRACE
	INTEGER
	STRING
	UUID
	PAGINATE
	WITH
	PAGINATION
	ENTITY
	RELATIONSHIP
	ONETOONE
	ONETOMANY
	MANYTOONE
	MANYTOMANY
	EOF
)

type TokenType int

type Token struct {
	Value     string
	TokenType TokenType
}

package parser

const (
	TokenType_Undefined   = "undefined"
	TokenType_Keyword     = "keyword"
	TokenType_Punctuation = "punctuation"
	TokenType_Identifier  = "identifier"
)

type TokenType string

type Token struct {
	Value     string
	TokenType TokenType
	DataType  DataType
}

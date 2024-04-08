package parser

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Token
	}{
		{
			input: "entity User { id Integer }",
			expected: []Token{
				{Value: "entity", TokenType: TokenType_Keyword, DataType: DataTypeText},
				{Value: "User", TokenType: TokenType_Identifier, DataType: DataTypeText},
				{Value: "{", TokenType: TokenType_Parenthesis, DataType: DataTypePunctuation},
				{Value: "id", TokenType: TokenType_Identifier, DataType: DataTypeText},
				{Value: "Integer", TokenType: TokenType_Keyword, DataType: DataTypeText},
				{Value: "}", TokenType: TokenType_Parenthesis, DataType: DataTypePunctuation},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			tokens, err := Tokenize(tc.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(*tokens, tc.expected) {
				t.Errorf("got tokens %+v, want %+v", *tokens, tc.expected)
			}
		})
	}
}

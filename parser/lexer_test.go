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
				{Value: "entity", TokenType: TokenType_Keyword, DataType: TypeText},
				{Value: "User", TokenType: TokenType_Identifier, DataType: TypeText},
				{Value: "{", TokenType: TokenType_Parenthesis, DataType: TypePunctuation},
				{Value: "id", TokenType: TokenType_Identifier, DataType: TypeText},
				{Value: "Integer", TokenType: TokenType_Keyword, DataType: TypeText},
				{Value: "}", TokenType: TokenType_Parenthesis, DataType: TypePunctuation},
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

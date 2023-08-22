package lexer

import (
	"reslang/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := "=+(){},;!"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.ILLEGAL, "!"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for _, v := range tests {
		token := lexer.NextToken()

		if token.Type != v.expectedType {
			t.Fatalf("Type wrong, expected=%q, got=%q", v.expectedType, token.Type)
		}

		if token.Literal != v.expectedLiteral {
			t.Fatalf("Literal wrong, expected=%q, got=%q", v.expectedLiteral, token.Literal)
		}
	}
}

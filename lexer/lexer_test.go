package lexer

import (
	"reslang/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `
	let five = 5;
	let ten_thousand = 10000;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten_thousand"},
		{token.ASSIGN, "="},
		{token.INT, "10000"},
		{token.SEMICOLON, ";"},

		{token.EOF, ""},
	}

	lexer := New(input)

	for _, v := range tests {
		token := lexer.Parse()

		if token.Type != v.expectedType {
			t.Fatalf("Type wrong, expected=%q, got=%q", v.expectedType, token.Type)
		}

		if token.Literal != v.expectedLiteral {
			t.Fatalf("Literal wrong, expected=%q, got=%q", v.expectedLiteral, token.Literal)
		}
	}
}

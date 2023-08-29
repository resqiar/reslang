package lexer

import (
	"reslang/token"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `
	let one = 1;
	let two = 2;

	let sum = fn(a, b) {
		a + b;
	};

	let result = sum(one, two);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "one"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "two"},
		{token.ASSIGN, "="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "sum"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "sum"},
		{token.LPAREN, "("},
		{token.IDENT, "one"},
		{token.COMMA, ","},
		{token.IDENT, "two"},
		{token.RPAREN, ")"},
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

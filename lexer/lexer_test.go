package lexer

import (
	"os"
	"reslang/token"
	"testing"
)

func TestLexer(t *testing.T) {
	raw, err := os.ReadFile("../index.rsq")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// let one = 1;
		// let two = 2;
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

		// let sum = fn(a, b) {
		// 	return a + b;
		// }
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
		{token.RETURN, "return"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// let result = sum(one, two);
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

		// !-/*5;
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// 5 < 10 > 5;
		{token.INT, "5"},
		{token.LTHAN, "<"},
		{token.INT, "10"},
		{token.GTHAN, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		// if (result < 10) {
		// 	result = sum(result * 100);
		// }
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.LTHAN, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "sum"},
		{token.LPAREN, "("},
		{token.IDENT, "result"},
		{token.ASTERISK, "*"},
		{token.INT, "100"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		// let foreverFalse = false;
		{token.LET, "let"},
		{token.IDENT, "foreverFalse"},
		{token.ASSIGN, "="},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},

		// if(5 < 10) {
		// 	return true;
		// } else if(10 > 5) {
		// 	return false;
		// } else {
		// 	return false;
		// }
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LTHAN, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "10"},
		{token.GTHAN, ">"},
		{token.INT, "5"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.EOF, ""},
	}

	lexer := New(string(raw))

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

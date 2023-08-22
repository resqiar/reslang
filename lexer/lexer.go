package lexer

import "reslang/token"

type Lexer struct {
	input      string
	currentPos int
	readPos    int
	char       byte
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}

	// initialize reading
	lexer.read()

	return lexer
}

func (l *Lexer) read() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}

	l.currentPos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.char {
	case '=':
		t = newToken(token.ASSIGN, l.char)
	case '+':
		t = newToken(token.PLUS, l.char)
	case '(':
		t = newToken(token.LPAREN, l.char)
	case ')':
		t = newToken(token.RPAREN, l.char)
	case '{':
		t = newToken(token.LBRACE, l.char)
	case '}':
		t = newToken(token.RBRACE, l.char)
	case ',':
		t = newToken(token.COMMA, l.char)
	case ';':
		t = newToken(token.SEMICOLON, l.char)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		t = newToken(token.ILLEGAL, l.char)
	}

	l.read()
	return t
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

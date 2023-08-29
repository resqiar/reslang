package lexer

import (
	"reslang/token"
)

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

func (l *Lexer) readKeyword() string {
	lastPos := l.currentPos

	// fast forward until find a letter which is not azAZ_
	for isLetter(l.char) {
		l.read()
	}

	return l.input[lastPos:l.currentPos]
}

func (l *Lexer) readNumber() string {
	lastPos := l.currentPos

	// fast forward until find a letter which is not 0-9
	for isNumber(l.char) {
		l.read()
	}

	return l.input[lastPos:l.currentPos]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.read()
	}
}

func (l *Lexer) Parse() token.Token {
	var t token.Token

	// skip whitespace, tab, enter, etc
	l.skipWhitespace()

	switch l.char {
	case '=':
		t = newToken(token.ASSIGN, l.char)
	case '+':
		t = newToken(token.PLUS, l.char)
	case '-':
		t = newToken(token.MINUS, l.char)
	case '!':
		t = newToken(token.BANG, l.char)
	case '*':
		t = newToken(token.ASTERISK, l.char)
	case '/':
		t = newToken(token.SLASH, l.char)

	case '(':
		t = newToken(token.LPAREN, l.char)
	case ')':
		t = newToken(token.RPAREN, l.char)
	case '{':
		t = newToken(token.LBRACE, l.char)
	case '}':
		t = newToken(token.RBRACE, l.char)
	case '<':
		t = newToken(token.LTHAN, l.char)
	case '>':
		t = newToken(token.GTHAN, l.char)
	case ',':
		t = newToken(token.COMMA, l.char)
	case ';':
		t = newToken(token.SEMICOLON, l.char)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.char) {
			// read identifiers until it finds non-match azAZ_ character
			t.Literal = l.readKeyword()

			// lookup the type of the keyword
			t.Type = token.LookupKeyword(t.Literal)

			// return here is necessary to skip the l.read() below
			// we already did the l.read() repeatly in LookupKeyword.
			return t
		} else if isNumber(l.char) {
			// read identifiers until it finds non-numeric character
			t.Literal = l.readNumber()
			t.Type = token.INT

			return t
		} else {
			// else, return ILLEGAL characters
			t = newToken(token.ILLEGAL, l.char)
		}
	}

	l.read()
	return t
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	// return true if char is within a-zA-Z_
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char == '_')
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

package token

const (
	// contains words that unknown for the lexer
	ILLEGAL = "ILLEGAL"
	// "End of File" - tell parser to stop
	EOF = "EOF"

	// identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

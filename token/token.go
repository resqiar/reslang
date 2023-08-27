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
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
}

func LookupKeyword(keyword string) TokenType {
	// if keyword is expected, then return it as KEYWORD
	if kw, ok := keywords[keyword]; ok {
		return kw
	}

	// otherwise, it will recognized as a IDENTIFIER
	return IDENT
}

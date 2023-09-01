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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LTHAN  = "<"
	GTHAN  = ">"

	EQUAL      = "=="
	NOTEQUAL   = "!="
	LTHANEQUAL = "<="
	GTHANEQUAL = ">="

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
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
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupKeyword(keyword string) TokenType {
	// if keyword is expected, then return it as KEYWORD
	if kw, ok := keywords[keyword]; ok {
		return kw
	}

	// otherwise, it will recognized as a IDENTIFIER
	return IDENT
}

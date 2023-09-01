package parser

import (
	"reslang/ast"
	"reslang/lexer"
	"reslang/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: l,
	}

	// read two times so currentToken and peekToken is set
	parser.next()
	parser.next()

	return parser
}

func (p *Parser) next() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.Parse()
}

// still no idea what is this going to do
func (p *Parser) Parse() *ast.Program {
	return nil
}

package parser

import (
	"fmt"
	"reslang/ast"
	"reslang/lexer"
	"reslang/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token

	// error handling
	errors []string
}

func New(l *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:  l,
		errors: []string{},
	}

	// read twice so both currentToken and peekToken is set
	parser.next()
	parser.next()

	return parser
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) AddError(message string) {
	p.errors = append(p.errors, message)
}

func (p *Parser) next() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.Parse()
}

func (p *Parser) Parse() *ast.Program {
	// initialize the root of the AST
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// loop through the lexer statement until EOF
	for !p.currentTokenIs(token.EOF) {
		statement := p.parseStatement()

		if statement != nil {
			// if there is a statement, push the statement into the array of statements
			program.Statements = append(program.Statements, statement)
		}

		p.next()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}

	// if next peek token is not IDENT, then return nil
	if !p.assertPeek(token.IDENT) {
		return nil
	}

	// construct Identifier ast-node
	statement.Ident = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	// if next peek token is not ASSIGN, then return nil
	if !p.assertPeek(token.ASSIGN) {
		return nil
	}

	// for now, skip until we find semicolon
	// next time, we need to parse the Value (expression) statement
	for !p.currentTokenIs(token.SEMICOLON) {
		p.next()
	}

	return statement
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) assertPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		// advance forward
		p.next()
		return true
	} else {
		p.AddError(fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type))
		return false
	}
}

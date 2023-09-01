package ast

import "reslang/token"

// Every Node in AST should implement
// TokenLiteral method, which returns value of the associated
// token literal. ONLY used for debugging and testing.
type Node interface {
	TokenLiteral() string
}

// Statement and Expression can implement a DUMMY method called
// statementNode or expressionNode. This is not necessary, but it can helps
// us tell the Go compiler if something bad happens, i guess?

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Implementation of a Node.
// This Program node is going to be the root of the AST.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// The reason why LetStatement have 3 fields
// is that "let x = 3;" is actually can be divided into that 3 parts,
// Token for "let"
// Ident for "x"
// Value for Expression that comes after "="
type LetStatement struct {
	Token token.Token
	Ident *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

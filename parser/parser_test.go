package parser

import (
	"reslang/ast"
	"reslang/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let firstNumber = 5;
		let secondNumber = 15;
		let sum = firstNumber + secondNumber;

		let foo = 99999;
	`

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.Parse()
	if program == nil {
		t.Fatalf("Program returned nil")
	}

	if len(program.Statements) != 4 {
		t.Fatalf("Program.Statements does not contains 4 statements. Got = %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdent string
	}{
		{"firstNumber"},
		{"secondNumber"},
		{"sum"},
		{"foo"},
	}

	for i, v := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, v.expectedIdent) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got = %q", s.TokenLiteral())
		return false
	}

	letStm, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not *ast.Statement. got = %T", letStm)
		return false
	}

	if letStm.Ident.Value != name {
		t.Errorf("letStm.Ident.Value is not %s. got = %s", name, letStm.Ident.Value)
		return false
	}

	if letStm.Ident.TokenLiteral() != name {
		t.Errorf("letStm.Ident.TokenLiteral() is not %s. got = %s", name, letStm.Ident.TokenLiteral())
		return false
	}

	return true
}

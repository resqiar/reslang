package parser

import (
	"log"
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

	// check parsing errors
	errors := parser.Errors()
	if len(errors) != 0 {
		t.Errorf("Parser has %d associated error(s)", len(errors))
		for _, message := range errors {
			t.Errorf("Parser Error: %q", message)
		}
		t.FailNow() // immediate fail
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

func TestReturnStatements(t *testing.T) {
	input := `
		return 18000;
		return "string";
		return 20;
	`
	expectedStatements := 3

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.Parse()
	if program == nil {
		t.Fatalf("Program returned nil")
	}

	if len(program.Statements) != expectedStatements {
		t.Fatalf("Program Statements produce = %d, expected = %d", len(program.Statements), expectedStatements)
	}

	// check parsing errors
	errors := parser.Errors()
	if len(errors) != 0 {
		t.Errorf("Parser has %d associated error(s)", len(errors))
		for _, message := range errors {
			t.Errorf("Parser Error: %q", message)
		}
		t.FailNow() // immediate fail
	}

	for _, statement := range program.Statements {
		log.Println(statement)
		returnStatement, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement is not *ast.ReturnStatement, got = %q", returnStatement)
			continue
		}

		if returnStatement.Token.Literal != "return" {
			t.Errorf("returnStatement.Token.Literal is not 'return', got = %q", returnStatement.Token.Literal)
		}
	}
}

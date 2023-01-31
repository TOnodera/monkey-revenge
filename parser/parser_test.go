package parser

import (
	"main/ast"
	"main/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {

	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram()がnilを返しました")
	}

	if len(program.Statemens) != 3 {
		t.Fatalf("program.Statementsが含むステートメントが期待値３と一致しません。実際の値:%d", len(program.Statemens))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statemens[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteralがletと一致しません。実際の値: %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("sが*ast.LetStatementと一致しません。実際の値: %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Valueが%sと一致しません。実際の値: %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.TokenLiteral()が%sと一致しません。実際の値: %s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true

}

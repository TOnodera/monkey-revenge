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
	checkParseErrors(t, p)

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

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("パーサーで%d個のエラーが発生しました。", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statemens) != 3 {
		t.Fatalf("program.Statementsの個数が期待値3と一致しません。期待値:3 実際の値: %d", len(program.Statemens))
	}

	for _, stmt := range program.Statemens {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmtがast.ReturnStatement1と一致しません。実際の値: %T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteralがreturnになっていません。期待値: return 実際の値: %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifirerExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statemens) != 1 {
		t.Fatalf("入力されたプログラムのStatementsの個数が期待値と一致しません。期待値:1 実際の値:%d", len(program.Statemens))
	}

	stmt, ok := program.Statemens[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0]はast.ExpressionStatementと一致しません。")
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("ident.Valueは%sではなく%sでした。", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLIteralは%sではなく%sでした。", "foobar", ident.TokenLiteral())
	}

}

func TestItengerLiteralExpression(t *testing.T) {
	input := "5"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statemens) != 1 {
		t.Fatalf("program.Statementsの数が期待値と一致しません。期待値:1 , 実際の値: %d", len(program.Statemens))
	}

	stmt, ok := program.Statemens[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0]が型ast.ExpressionStatementと一致しません。実際の値: %T", program.Statemens[0])
	}

	literal, ok = stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("式がast.IntgerLiteralと一致しません。実際の値: %T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Valueが期待値と一致しません。期待値: 5,実際の値: %d", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLIteralが%sと一致しません。", "5", literal.TokenLIteral())
	}
}

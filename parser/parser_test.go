package parser

import (
	"fmt"
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
	checkParserErrors(t, p)

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

func checkParserErrors(t *testing.T, p *Parser) {
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
	checkParserErrors(t, p)

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
	checkParserErrors(t, p)

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
	checkParserErrors(t, p)

	if len(program.Statemens) != 1 {
		t.Fatalf("program.Statementsの数が期待値と一致しません。期待値:1 , 実際の値: %d", len(program.Statemens))
	}

	stmt, ok := program.Statemens[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0]が型ast.ExpressionStatementと一致しません。実際の値: %T", program.Statemens[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("式がast.IntgerLiteralと一致しません。実際の値: %T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Valueが期待値と一致しません。期待値: 5,実際の値: %d", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLIteralが%sと一致しません。実際の値: %s", "5", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statemens) != 1 {
			t.Fatalf("program.Statementsの数が期待値と一致しません。期待値:1, 実際の値: %d",
				len(program.Statemens))
		}

		stmt, ok := program.Statemens[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0]は型ast.ExpressionStatemtと一致しません。実際の値: %T",
				program.Statemens[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmtはast.PrefixExpressionではありません。実際の値: %T",
				stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operatorは%sではありません。実際の値: %s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("ilはast.IntegerLiteral型ではありません。実際の値: %T", il)
	}

	if integ.Value != value {
		t.Errorf("integ.Valueは%dではありません。実際の値: %d", value, integ.Value)
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteralは%dではありません。実際の値: %s", value, integ.TokenLiteral())
	}

	return true
}

func TestParseingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5+5;", 5, "+", 5},
		{"5-5;", 5, "-", 5},
		{"5*5;", 5, "*", 5},
		{"5/5;", 5, "/", 5},
		{"5>5;", 5, ">", 5},
		{"5<5;", 5, "<", 5},
		{"5==5;", 5, "==", 5},
		{"5!=5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statemens) != 1 {
			t.Fatalf("program.Statementsの数が期待値と一致しません。期待値: 1, 実際の値: %d", len(program.Statemens))
		}

		stmt, ok := program.Statemens[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0]はast.ExpressionStatement型ではありません。実際の型: %T", program.Statemens[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("expはInfixExpression型ではありません。実際の型: %T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operatorが期待値と一致しません。期待値: %s, 実際の値: %s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

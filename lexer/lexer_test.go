package lexer

import (
	"main/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `
	let five = 5;
	let ten = 10;

	let add = fn(x,y){
		x + y;
	};

	let result = add(five,ten);
	!-/*5;
	5 < 10 > 5;
		
	if (5 < 10) {
		return true;
	} else {
		return false;
	}
	
	10 == 10;
	10 != 9;
	`
	tests := []struct {
		extectedType    token.TokenType
		extectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.extectedType {
			t.Fatalf("テスト%dの期待値の型(%q)と実際の型(%q)が一致しません", i, tt.extectedType, tok.Type)
		}

		if tok.Literal != tt.extectedLiteral {
			t.Fatalf("テスト%dの期待値(%q)と実際の値(%q)が一致しません", i, tt.extectedLiteral, tok.Literal)
		}

	}

}

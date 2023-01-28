package lexer

import (
	"main/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		extectedType    token.TokenType
		extectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
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

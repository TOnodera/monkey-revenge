package ast

import (
	"main/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statemens: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{
						Type: token.IDENT, Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type: token.IDENT, Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar=anotherVar;" {
		t.Errorf("program.String()に誤りがあります。実際の値: %q", program.String())
	}
}
package ast_test

import (
	"testing"
	tok "github.com/Appleby43/blakescript/token"
)

type token tok.Token

func TestSimpleStatement(t *testing.T) {
	inpt := []token {
		token{Type: tok.Let, Literal: "let"},
		token{Type: tok.Id, Literal: "x"},
		token{Type: tok.Assign, Literal: "="},
		token{Type: tok.Int, Literal: "69"},
		token{Type: tok.Semicolon, Literal: ";"},
	}

	
}
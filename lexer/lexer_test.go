package lexer

import (
	"testing"

	"github.com/Appleby43/blakescript/assert"
	"github.com/Appleby43/blakescript/token"
)

func TestNextToken(t *testing.T) {
	inpt := "=+(){},;"

	expected := [8]token.TokenType {
		token.Assign,
		token.Plus,
		token.OpenParen,
		token.ClosedParen,
		token.OpenBrace,
		token.ClosedBrace,
		token.Comma,
		token.Semicolon,
	}

	lexer := New(inpt)

	for _, tokenType := range expected {
		next := lexer.NextToken()
		assert.IntEquals(int(next.Type), int(tokenType), t)
	}
	
}
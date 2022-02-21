package lexer

import (
	"fmt"
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

func TestComplexCode(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	}

	let result = add(five, ten);
	`

	expectedResults := []struct {
		type_ token.TokenType
		literal string
		} {
			{token.Let, "let"},
			{token.Id, "five"},
			{token.Assign, "="},
			{token.Int, "5"},
			{token.Semicolon, ";"},
			{token.Let, "let"},
			{token.Id, "ten"},
			{token.Assign, "="},
			{token.Int, "10"},
			{token.Semicolon, ";"},
			{token.Let, "let"},
			{token.Id, "add"},
			{token.Assign, "="},
			{token.Function, "fn"},
			{token.OpenParen, "("},
			{token.Id, "x"},
			{token.Comma, ","},
			{token.Id, "y"},
			{token.ClosedParen, ")"},
			{token.OpenBrace, "{"},
			{token.Id, "x"},
			{token.Plus, "+"},
			{token.Id, "y"},
			{token.Semicolon, ";"},
			{token.Let, "let"},
			{token.Id, "result"},
			{token.Assign, "="},
			{token.Id, "add"},
			{token.OpenParen, "("},
			{token.Id, "five"},
			{token.Comma, ","},
			{token.Id, "ten"},
			{token.ClosedParen, ")"},
			{token.Semicolon, ";"},
	}

	lexer := New(input)

	for _, expected := range expectedResults {
		actual := lexer.NextToken();  
		if assert.IntEquals(int(actual.Type), int(expected.type_), t) {
			fmt.Printf("expected %s, got %s with val %s\n", expected.literal, actual.Type.String(), actual.Literal)
		}
	}
}
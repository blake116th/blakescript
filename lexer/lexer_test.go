package lexer

import (
	"fmt"
	"testing"

	"github.com/Appleby43/blakescript/assert"
	"github.com/Appleby43/blakescript/token"
)

type tokenPair struct {
	type_ token.TokenType
	literal string
}

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

	expectedResults := []tokenPair {
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
			{token.ClosedBrace, "}"},
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

	generalLexTest(t, input, expectedResults)
}

func TestOperators(t *testing.T) {
	input := `
	- 5 + 5 / 123 * foo

	!(113 < 115)

	6 > four
	`

	expectedResults := []tokenPair {
			{token.Minus, "-"},
			{token.Int, "5"},
			{token.Plus, "+"},
			{token.Int, "5"},
			{token.Divide, "/"},
			{token.Int, "123"},
			{token.Times, "*"},
			{token.Id, "foo"},
			{token.Not, "!"},
			{token.OpenParen, "("},
			{token.Int, "113"},
			{token.LessThan, "<"},
			{token.Int, "115"},
			{token.ClosedParen, ")"},
			{token.Int, "6"},
			{token.GreaterThan, ">"},
			{token.Id, "four"},
	}

	generalLexTest(t, input, expectedResults)
}

func TestMoreKeywords(t *testing.T) {
	input := `
	if (true) { return false } else return 0
	`

	expectedResults := []tokenPair {
			{token.If, "if"},
			{token.OpenParen, "("},
			{token.True, "true"},
			{token.ClosedParen, ")"},
			{token.OpenBrace, "{"},
			{token.Return, "return"},
			{token.False, "false"},
			{token.ClosedBrace, "}"},
			{token.Else, "else"},
			{token.Return, "return"},
			{token.Int, "0"},
	}

	generalLexTest(t, input, expectedResults)
}

func generalLexTest(t *testing.T, input string, expectedResults []tokenPair) {
	lexer := New(input)

	for _, expected := range expectedResults {
		actual := lexer.NextToken();  
		if assert.IntEquals(int(actual.Type), int(expected.type_), t) {
			fmt.Printf("expected %s, got %s with val %s\n", expected.literal, actual.Type.String(), actual.Literal)
		}

		assert.StringEquals(actual.Literal, expected.literal, t)
	}
}
package lexer

import (
	"fmt"
	"testing"

	"github.com/Appleby43/blakescript/assert"
	"github.com/Appleby43/blakescript/token"
)

type tokenPair struct {
	type_   token.TokenType
	literal string
	lineNo  int
}

func TestNextToken(t *testing.T) {
	inpt := "=+(){},;"

	expected := [8]token.TokenType{
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
	input :=
		`let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	}

	let result = add(five, ten);
	`

	expectedResults := []tokenPair{
		{token.Let, "let", 1},
		{token.Id, "five", 1},
		{token.Assign, "=", 1},
		{token.Int, "5", 1},
		{token.Semicolon, ";", 1},
		{token.Let, "let", 2},
		{token.Id, "ten", 2},
		{token.Assign, "=", 2},
		{token.Int, "10", 2},
		{token.Semicolon, ";", 2},
		{token.Let, "let", 3},
		{token.Id, "add", 3},
		{token.Assign, "=", 3},
		{token.Function, "fn", 3},
		{token.OpenParen, "(", 3},
		{token.Id, "x", 3},
		{token.Comma, ",", 3},
		{token.Id, "y", 3},
		{token.ClosedParen, ")", 3},
		{token.OpenBrace, "{", 3},
		{token.Id, "x", 4},
		{token.Plus, "+", 4},
		{token.Id, "y", 4},
		{token.Semicolon, ";", 4},
		{token.ClosedBrace, "}", 5},
		{token.Let, "let", 7},
		{token.Id, "result", 7},
		{token.Assign, "=", 7},
		{token.Id, "add", 7},
		{token.OpenParen, "(", 7},
		{token.Id, "five", 7},
		{token.Comma, ",", 7},
		{token.Id, "ten", 7},
		{token.ClosedParen, ")", 7},
		{token.Semicolon, ";", 7},
	}

	generalLexTest(t, input, expectedResults)
}

func TestOperators(t *testing.T) {
	input := `
	- 5 + 5 / 123 * foo

	!(113 < 115)

	6 > four
	`

	expectedResults := []tokenPair{
		{token.Minus, "-", 2},
		{token.Int, "5", 2},
		{token.Plus, "+", 2},
		{token.Int, "5", 2},
		{token.Divide, "/", 2},
		{token.Int, "123", 2},
		{token.Times, "*", 2},
		{token.Id, "foo", 2},
		{token.Not, "!", 4},
		{token.OpenParen, "(", 4},
		{token.Int, "113", 4},
		{token.LessThan, "<", 4},
		{token.Int, "115", 4},
		{token.ClosedParen, ")", 4},
		{token.Int, "6", 6},
		{token.GreaterThan, ">", 6},
		{token.Id, "four", 6},
	}

	generalLexTest(t, input, expectedResults)
}

func TestMoreKeywords(t *testing.T) {
	input := `
	if (true) { return false } else return 0
	`

	expectedResults := []tokenPair{
		{token.If, "if", 2},
		{token.OpenParen, "(", 2},
		{token.True, "true", 2},
		{token.ClosedParen, ")", 2},
		{token.OpenBrace, "{", 2},
		{token.Return, "return", 2},
		{token.False, "false", 2},
		{token.ClosedBrace, "}", 2},
		{token.Else, "else", 2},
		{token.Return, "return", 2},
		{token.Int, "0", 2},
	}

	generalLexTest(t, input, expectedResults)
}

func TestEqualAndNotEqual(t *testing.T) {
	input := `
	left = 3 == 2 

	left != right !
	`

	expectedResults := []tokenPair{
		{token.Id, "left", 2},
		{token.Assign, "=", 2},
		{token.Int, "3", 2},
		{token.Equals, "==", 2},
		{token.Int, "2", 2},
		{token.Id, "left", 4},
		{token.NotEquals, "!=", 4},
		{token.Id, "right", 4},
		{token.Not, "!", 4},
	}

	generalLexTest(t, input, expectedResults)
}

func generalLexTest(t *testing.T, input string, expectedResults []tokenPair) {
	lexer := New(input)

	for _, expected := range expectedResults {
		actual := lexer.NextToken()
		if assert.IntEquals(int(actual.Type), int(expected.type_), t) {
			fmt.Printf("expected %s, got %s with val %s\n", expected.literal, actual.Type.String(), actual.Literal)
		}

		assert.StringEquals(actual.Literal, expected.literal, t)
		assert.IntEquals(actual.LineNumber, expected.lineNo, t)
	}
}

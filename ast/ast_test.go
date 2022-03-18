package ast_test

import (
	"testing"

	"github.com/Appleby43/blakescript/assert"
	"github.com/Appleby43/blakescript/ast"
	tkn "github.com/Appleby43/blakescript/token"
)

func TestSimpleStatement(t *testing.T) {
	inpt := []tkn.Token {
		{Type: tkn.Let, Literal: "let"},
		{Type: tkn.Id, Literal: "a"},
		{Type: tkn.Assign, Literal: "="},
		{Type: tkn.Int, Literal: "69"},
		{Type: tkn.Semicolon, Literal: ";"},
		{Type: tkn.Let, Literal: "let"},
		{Type: tkn.Id, Literal: "b"},
		{Type: tkn.Assign, Literal: "="},
		{Type: tkn.Int, Literal: "420"},
		{Type: tkn.Semicolon, Literal: ";"},
	}

	expected := ast.CodeBlock{
		Children: []ast.Statement{
			&ast.LetStatement{
				Id: "a",
				Expression: &ast.IntLiteral{
					Value: 69,
			}},
			&ast.LetStatement{
				Id: "b",
				Expression: &ast.IntLiteral{
					Value: 420,
			}},
		},
	}

	tree := ast.Parse(inpt)

	assert.CodeBlockEquals(tree, expected, t)
}
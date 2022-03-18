package ast

import "github.com/Appleby43/blakescript/token"

func Parse(tokens []token.Token) CodeBlock {
	return CodeBlock{
		Children: make([]Statement, 0),
	}
}
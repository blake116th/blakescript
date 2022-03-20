package ast

import (
	"fmt"

	"github.com/Appleby43/blakescript/token"
)

type UnexpectedTokenError struct {
	tkn token.Token
}

func (u *UnexpectedTokenError) Error() string {
	return "Unexpected Token \"" + u.tkn.Literal + "\""
}

type PrematureEndOfFileError struct {
	lastTkn token.Token
}

func (u *PrematureEndOfFileError) Error() string {
	return "File ended prematurely while parsing token " + u.lastTkn.Literal	
}

type ExpectedExpressionError struct {
	lastTkn token.Token
}

func (u *ExpectedExpressionError) Error() string {
	return fmt.Sprintf("Expected an expression on line %v, and was unable to parse token \"%v\" into an expression.", u.lastTkn.LineNumber, u.lastTkn.Literal)	
}
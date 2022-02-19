package token

import "log"

type TokenType int

const (
	Illegal TokenType = iota
	EOF

	Id

	//Literals
	Int

	//operators
	Assign
	Plus

	//Special characters
	Comma
	Semicolon

	OpenParen
	ClosedParen

	OpenBrace
	ClosedBrace

	//keywords
	Function
	Let
)

func (t TokenType) String() string {
	switch t {
	case Illegal:
		return "Illegal"
	case EOF:
		return "End of File"
	case Id:
		return "Identifier"
	case Int:
		return "Integer"
	case Assign:
		return "Assignment"
	case Plus:
		return "Plus"
	case Comma:
		return "Comma"
	case Semicolon:
		return "Semicolon"
	case OpenParen:
		return "Opening Parenthesis"
	case ClosedParen:
		return "Closing Parenthesis"
	case OpenBrace:
		return "Opening Bracket"
	case ClosedBrace:
		return "Closing Bracket"
	case Function:
		return "Function Declaration"
	case Let:
		return "Variable Declaration"
	default:
		log.Fatal("Could not find a string representation for a token")
		return ""
	}
}

type Token struct {
	Type TokenType
	Literal string
}


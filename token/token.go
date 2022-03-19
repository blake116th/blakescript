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
	Minus
	Not
	Times
	Divide

	LessThan
	GreaterThan
	Equals
	NotEquals

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
	If
	Else
	Return
	True
	False
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
	case Minus:
		return "Minus"
	case Not:
		return "Not"
	case Times:
		return "Times"
	case Comma:
		return "Comma"
	case Divide:
		return "Divide"
	case LessThan:
		return "Less Than"
	case GreaterThan:
		return "Greater Than"
	case Equals:
		return "Equals"
	case NotEquals:
		return "Not Equals"
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
	case If:
		return "If statement"
	case Else:
		return "Else"
	case Return:
		return "Return"
	case True:
		return "True"
	case False:
		return "False"
	default:
		log.Fatal("Could not find a string representation for a token")
		return ""
	}
}

type Token struct {
	Type TokenType
	Literal string
	LineNumber int
}


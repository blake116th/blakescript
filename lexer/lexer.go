package lexer

import "github.com/Appleby43/blakescript/token"


type Lexer struct {
	input string
	position int
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		position: 0,
	}
	return l
}

func (l *Lexer) NextToken() token.Token {
	char := l.input[l.position]

	l.position++
	if (l.position > len(l.input)) {
		return makeToken(token.EOF, '0')
	}

	switch char {
	case '=':
		return makeToken(token.Assign, char)
	case ';':
		return makeToken(token.Semicolon, char)
	case '(':
		return makeToken(token.OpenParen, char)
	case ')':
		return makeToken(token.ClosedParen, char)
	case ',':
		return makeToken(token.Comma, char)
	case '+':
		return makeToken(token.Plus, char)
	case '{':
		return makeToken(token.OpenBrace, char)
	case '}':
		return makeToken(token.ClosedBrace, char)
	default:
		return makeToken(token.Illegal, char)
	}
}

func makeToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
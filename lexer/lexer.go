package lexer

import (
	"strings"

	"github.com/Appleby43/blakescript/token"
)

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

var keywords = map[string] token.TokenType {
	"fn" : token.Function,
	"let" : token.Let,
}

func (l *Lexer) NextToken() token.Token {
	char, atEnd := l.advanceChar()

	//skip whitespace
	if (isWhitespace(char)) {
		return l.NextToken()
	}

	if atEnd {
		return makeToken(token.EOF, string(char))
	}

	switch char {
	case '=':
		return makeToken(token.Assign, string(char))
	case ';':
		return makeToken(token.Semicolon, string(char))
	case '(':
		return makeToken(token.OpenParen, string(char))
	case ')':
		return makeToken(token.ClosedParen, string(char))
	case ',':
		return makeToken(token.Comma, string(char))
	case '+':
		return makeToken(token.Plus, string(char))
	case '{':
		return makeToken(token.OpenBrace, string(char))
	case '}':
		return makeToken(token.ClosedBrace, string(char))
	default:
		if isLetter(char) {
			return l.parseWord(char)
		}
		if isDigit(char) {
			return l.parseNumber(char)
		}
		return makeToken(token.Illegal, string(char))
	}
}

func (l *Lexer) parseWord(currentChar byte) token.Token {
	var sb strings.Builder
	sb.WriteByte(currentChar)

	for {
		nextChar, atEnd := l.peekChar()

		if !isLetter(nextChar) || atEnd {
			break
		}

		sb.WriteByte(nextChar)
		l.position++
	}

	word := sb.String()
	
	if tokenType, contains := keywords[word]; contains {
		return makeToken(tokenType, word)
	} else {
		return makeToken(token.Id, word)
	}
}

func (l *Lexer) parseNumber(currentChar byte) token.Token {
	var sb strings.Builder
	sb.WriteByte(currentChar)

	for {
		nextChar, atEnd := l.peekChar()

		if !isDigit(nextChar) || atEnd {
			break
		}

		sb.WriteByte(nextChar)
		l.position++
	}

	return makeToken(token.Int, sb.String())
}

//returns true if it's the last character
func (l *Lexer) advanceChar() (byte, bool){
	retVal, atEnd := l.peekChar()

	if !atEnd {
		l.position++
	}
	return byte(retVal), atEnd
}

// peeks at the current character without incrementing the index
//returns true if at end of file
func (l *Lexer) peekChar() (byte, bool){
	if (l.position > len(l.input) - 1) {
		return '0', true
	}

	return l.input[l.position], false
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func makeToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
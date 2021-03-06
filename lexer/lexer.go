package lexer

import (
	"strings"

	"github.com/Appleby43/blakescript/token"
)

type Lexer struct {
	input string
	position int
	lineNo int
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		position: 0,
		lineNo: 1,
	}
	return l
}

var keywords = map[string] token.TokenType {
	"fn" : token.Function,
	"let" : token.Let,
	"true" : token.True,
	"false" : token.False,
	"if" : token.If,
	"else" : token.Else,
	"return" : token.Return,
}

//certain tokens can be definitively parsed into tokens if they are only used in one context
//ex: ( can only be interpreted as an open parentheses, while = could be assignment, or ==, or !=
var unambiguousCharTokens = map[byte] token.TokenType {
	';' : token.Semicolon,
	'(' : token.OpenParen,
	')' : token.ClosedParen,
	',' : token.Comma,
	'+' : token.Plus,
	'-' : token.Minus,
	'*' : token.Times,
	'/' : token.Divide,
	'<' : token.LessThan,
	'>' : token.GreaterThan,
	'{' : token.OpenBrace,
	'}' : token.ClosedBrace,
}

func (l *Lexer) NextToken() token.Token {
	char, atEnd := l.advanceChar()

	//skip whitespace and newlines
	if isWhitespace(char) {
		return l.NextToken()
	}

	if isNewLine(char) {
		l.lineNo++
		return l.NextToken()
	}

	if atEnd {
		return l.makeToken(token.EOF, string(char))
	}

	tokenType, unambiguous := unambiguousCharTokens[char]
	if unambiguous {
		return l.makeToken(tokenType, string(char))
	}

	if char == '=' {
		nextChar, atEnd := l.peekChar()
		if !atEnd && nextChar == '='  {
			l.advanceChar()
			return l.makeToken(token.Equals, "==")
		}
		return l.makeToken(token.Assign, "=")
	}

	if char == '!' {
		nextChar, atEnd := l.peekChar()
		if !atEnd && nextChar == '='  {
			l.advanceChar()
			return l.makeToken(token.NotEquals, "!=")
		}
		return l.makeToken(token.Not, "!")
	}

	if isLetter(char) {
		return l.parseWord(char)
	}
	
	if isDigit(char) {
		return l.parseNumber(char)
	}

	return l.makeToken(token.Illegal, string(char))
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
		return l.makeToken(tokenType, word)
	} else {
		return l.makeToken(token.Id, word)
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

	return l.makeToken(token.Int, sb.String())
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
	if l.AtEnd() {
		return '0', true
	}

	return l.input[l.position], false
}

func (l *Lexer) AtEnd() bool {
	return l.position > len(l.input) - 1
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\r'
}

func isNewLine(char byte) bool {
	return char == '\n'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) makeToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal, LineNumber: l.lineNo}
}
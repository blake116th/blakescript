package ast

import (
	"strconv"

	"github.com/Appleby43/blakescript/token"
)

type parser struct {
	tokens []token.Token
	index int
}


func Parse(tokens []token.Token) (CodeBlock, error) {
	retVal := CodeBlock{
		Children: make([]Statement, 0),
	}

	p := parser{tokens: tokens}

	for {
		tkn := p.current();
		var err error
		switch tkn.Type {
			case token.Let:
				err = p.parseLet(&retVal)
			case token.Semicolon:
				// ignore
			default:
				return retVal, &UnexpectedTokenError{tkn: tkn}
		}

		if err != nil {
			return retVal, err
		}

		if !p.hasNext() {
			break;
		}
		p.next() //increment index
	}

	return retVal, nil	
}

func (p *parser) hasNext() bool {
	return p.index <= len(p.tokens) - 2
}

func (p *parser) next() token.Token {
	p.index++
	return p.tokens[p.index]
}

func (p *parser) current() token.Token {
	return p.tokens[p.index]
}

//effectively asserts the type of the next token. increments such that 
//next token then becomes the current token
func (p *parser) expectNext(expectedType token.TokenType) (next token.Token, err error) {
	if !p.hasNext() {
		return token.Token{}, &PrematureEndOfFileError{lastTkn: p.current()}
	}

	if p.next().Type != expectedType {
		return token.Token{}, &UnexpectedTokenError{tkn: p.current()}
	}

	return p.current(), nil
}

func (p *parser) parseLet(codeBlock *CodeBlock) error {
	tkn, err := p.expectNext(token.Id)

	if err != nil {
		return err
	}

	id := tkn.Literal

	_, err = p.expectNext(token.Assign)

	if err != nil {
		return err
	}

	expr, err := p.parseExpression()

	if err != nil {
		return err
	}

	codeBlock.Children = append(codeBlock.Children, &LetStatement{Id: id, Expression: expr})

	return nil
}

//attempts to parse an expression starting at p's current index
func (p *parser) parseExpression() (Expression, error) {
	if !p.hasNext() {
		return nil, &PrematureEndOfFileError{lastTkn: p.current()}
	}

	switch p.next().Type {
	case token.Int:
		val, err := strconv.Atoi(p.current().Literal)

		if err != nil {
			return nil, err
		}

		return &IntLiteral{Value: val}, nil

	case token.Id:
		return &IdentifierExpression{Id: p.current().Literal}, nil
	default:
		return nil, &ExpectedExpressionError{p.current()}
	}
}
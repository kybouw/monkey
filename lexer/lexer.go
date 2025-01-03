package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (where ch is)
	readPosition int  // current reading position in input (after ch, the thing to be read next)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // fill first char into examination
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// if next read is out of bounds, ch is null
		l.ch = 0
	} else {
		// else we can bump up to the next char in input
		l.ch = l.input[l.readPosition]
	}
	// shift positions in input up
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

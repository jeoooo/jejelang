package lexer

import (
	"unicode"
)

type Lexer struct {
    input        string
    position     int // current position in input (points to current char)
    readPosition int // current reading position in input (after current char)
    ch           byte // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition++
}

func (l *Lexer) NextToken() Token {
    var tok Token

    l.skipWhitespace()

    switch l.ch {
    case '=':
        tok = newToken(ASSIGN, l.ch)
    case '+':
        tok = newToken(PLUS, l.ch)
    case ',':
        tok = newToken(COMMA, l.ch)
    case ';':
        tok = newToken(SEMICOLON, l.ch)
    case '(':
        tok = newToken(LPAREN, l.ch)
    case ')':
        tok = newToken(RPAREN, l.ch)
    case '{':
        tok = newToken(LBRACE, l.ch)
    case '}':
        tok = newToken(RBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = EOF
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = LookupIdent(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Literal = l.readNumber()
            tok.Type = INT
            return tok
        } else {
            tok = newToken(ILLEGAL, l.ch)
        }
    }

    l.readChar()
    return tok
}

func newToken(tokenType TokenType, ch byte) Token {
    return Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func isLetter(ch byte) bool {
    return unicode.IsLetter(rune(ch))
}

func isDigit(ch byte) bool {
    return unicode.IsDigit(rune(ch))
}

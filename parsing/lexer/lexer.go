package lexer

import (
	par "arit/parsing"
	"io"
	"strings"
	"time"
	"unicode"
)

type Lexer struct {
	pos    par.Position
	reader *strings.Reader

	old rune
	ch  rune
	err error
}

func New(s string) *Lexer {
	return &Lexer{
		pos: par.Position{
			Line:   1,
			Column: 0,
		},
		reader: strings.NewReader(s),
	}
}

func (l *Lexer) unread() {
	if l.pos.IsAtStart() {
		return
	}
	l.err = l.reader.UnreadRune()

	l.pos.Line = max(1, l.pos.Line-1)
	l.pos.Column = max(0, l.pos.Column-1)
}

func (l *Lexer) next() {
	l.ch, _, l.err = l.reader.ReadRune()
	if l.ch == '\n' {
		l.pos.Line++
		l.pos.Column = 0
	}

	l.pos.Column++
}

func (l *Lexer) Lex() par.Token {
	ch := make(chan par.Token, 1)
	go func() {
		ch <- l.lex()
	}()
	select {
	case x := <-ch:
		return x
	case <-time.After(3 * time.Second):
		panic("lexer timed out")
	}
}

func (l *Lexer) lex() par.Token {

	for {
		l.next()
		if l.err != nil {
			if l.err == io.EOF {
				return par.Token{T: par.EOF, Pos: l.pos, S: ""}
			}
			return par.Token{T: par.ILLEGAL, Pos: l.pos, S: "ILLEGAL"}
		}

		switch l.ch {
		case ' ':
			continue
		case '.':
			return par.Token{T: par.DOT, Pos: l.pos, S: "."}
		case '"':
			return par.Token{T: par.DOUBLEQUOTE, Pos: l.pos, S: "\""}
		case '\'':
			return par.Token{T: par.SINGLEQUOTE, Pos: l.pos, S: "'"}
		case '(':
			return par.Token{T: par.LPAREN, Pos: l.pos, S: "("}
		case ')':
			return par.Token{T: par.RPAREN, Pos: l.pos, S: ")"}
		case '=':
			return par.Token{T: par.EQ, Pos: l.pos, S: "="}
		case '+':
			return par.Token{T: par.PLUS, Pos: l.pos, S: "+"}
		case '-':
			return par.Token{T: par.MINUS, Pos: l.pos, S: "-"}
		case '*':
			return par.Token{T: par.MULTI, Pos: l.pos, S: "*"}
		case '\\':
			return par.Token{T: par.BACKSLASH, Pos: l.pos, S: "\\"}
		case '/':
			return par.Token{T: par.SLASH, Pos: l.pos, S: "/"}
		default:
			if unicode.IsDigit(l.ch) {
				return l.number()
			} else if unicode.IsLetter(l.ch) {
				return l.ident()
			}

		}
	}
}

func (l *Lexer) number() par.Token {
	tok := par.Token{T: par.NUMBER, Pos: l.pos, S: string(l.ch)}
	for {
		l.next()
		if l.ch == '\n' {
			return tok
		}
		if unicode.IsDigit(l.ch) {
			tok.S += string(l.ch)
			continue
		}
		l.unread()

		return tok
	}
}

func (l *Lexer) ident() par.Token {
	tok := par.Token{T: par.IDENT, Pos: l.pos, S: string(l.ch)}
	for {
		l.next()
		if l.ch == '\n' {
			return tok
		}
		if unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
			tok.S += string(l.ch)
			continue
		}
		l.unread()

		return tok
	}
}

func (l *Lexer) string() par.Token {
	tok := par.Token{T: par.IDENT, Pos: l.pos, S: string(l.ch)}
	for {
		l.next()
		if l.ch == '"' {
			return tok
		}
		if unicode.IsLetter(l.ch) || unicode.IsDigit(l.ch) {
			tok.S += string(l.ch)
			continue
		}

		l.unread()

		return tok
	}
}

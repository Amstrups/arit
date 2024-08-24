package parsing

import (
	e "errors"
	"fmt"
)

var InvalidArgument = e.New("lists of unequal lengths")

type TokenCode int

const (
	EOF = iota
	ILLEGAL

	IDENT
	NUMBER

	DOT
	DOUBLEQUOTE
  SINGLEQUOTE

	LPAREN
	RPAREN

	EQ
	PLUS
	MINUS
	MULTI

	BACKSLASH
  SLASH
)

type Token struct {
	T   TokenCode
	Pos Position
	S   string
}

func IsEqual(value []Token, expected []TokenCode) bool {
	if len(value) != len(expected) {
		return false
	}
	for i := range value {
		if value[i].T != expected[i] {
			return false
		}

	}
	return true
}

func Pretty(pos Position, t TokenCode, s string) string {
	sString := fmt.Sprintf("'%s'", s)
	return fmt.Sprintf("[%-8s %12s - %d", pos.String(), sString, t)
}

func PrettyT(tok Token) string {
	return Pretty(tok.Pos, tok.T, tok.S)
}

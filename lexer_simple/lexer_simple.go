package lexersimple

import (
	"io"
	"strings"
	"unicode"
)

type Position struct {
  Line, Column int
}

type GLexer interface {
  Lex() (Position, Token, string)
}


type SLexer struct { 
  pos Position
  reader *strings.Reader
}

func NewSLexer(s string) *SLexer {
  return &SLexer{
    pos:Position{
      Line: 1,
      Column: 0,
    }, 
    reader: strings.NewReader(s),
  }
}

func (l *SLexer) stepBack() {
  err := l.reader.UnreadRune()

  if err != nil {
    panic(err)
  }
  
  l.pos.Column--
}

func (l *SLexer) newLine() { 
  l.pos.Line++
  l.pos.Column = 0
}

func (l *SLexer) Lex() (Position, Token, string) {
  //  var lit string
  for { 
    r, _, err := l.reader.ReadRune()
    if err != nil {
      if err == io.EOF {
        return l.pos, EOF, ""
      }
      panic(err)
    }

    l.pos.Column++ 

    switch r { 
    case '\n':
      l.newLine()
    case '+':
      return l.pos, ADD, "+"
    case '-':
      return l.pos, SUB, "-"
    case '*':
      return l.pos, MUL, "*"
    case '/':
      return l.pos, DIV, "/"
    case '(':
      return l.pos, LPAREN, "("
    case ')':
      return l.pos, RPAREN, ")"
    default:
      if (unicode.IsDigit(r)) { 
        tok, s := l.lexInt(r)
        return l.pos, tok, s
      }
      if (unicode.IsLetter(r)) {
        return l.pos, IDENT, l.lexIdent(r)
      }

      return l.pos, ILLEGAL, string(r) 
    }
  }
}


func (l *SLexer) lexInt(prefix rune) (Token, string) {
  lit := string(prefix) 
  var tok Token = INT
  for {
    r, _, err := l.reader.ReadRune()
    if err != nil { 
      if err == io.EOF {
        return tok,lit
      }
      panic(err)
    }

    if unicode.IsDigit(r) {
      lit = lit + string(r)
    } else if r == '.' {
      tok = FLOAT
      lit = lit + "."
    } else {
      l.stepBack()
      return tok,lit 
    }
  }
}

func (l *SLexer) lexIdent(prefix rune) string {
  lit := string(prefix) 
  for {
    r, _, err := l.reader.ReadRune()
    if err != nil { 
      if err == io.EOF {
        return lit
      }
      panic(err)
    }

    if unicode.IsLetter(r) {
      lit = lit + string(r)
    } else {
      l.stepBack()
      return lit 
    }
  }
}

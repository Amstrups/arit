package lexersimple

import (
  "bufio"
  "fmt"
  "io"
)

type Lexer struct { 
  pos Position
  reader *bufio.Reader
}

func (l *Lexer) newLine() {
  l.pos.Line++
  l.pos.Column = 0
}

func NewLexer(reader io.Reader) *Lexer {
  return &Lexer{
    pos:Position{
      Line: 1,
      Column: 0,
    }, 
    reader: bufio.NewReader(reader),
  }
}


func (l *Lexer) Lex() (Position, Token, string) {
  //l2 := func (b []byte) int { 
  //  return 2    
  //} 
  fmt.Println("lexing so hard")
  return Position{0,0}, 0,"s"
}


package parsersimple

import (
  "arit/lexer_simple"
  "fmt"
)
type Lexed struct {
  pos lexersimple.Position
  tok lexersimple.Token
  s string
}

type buffer struct {


}

func (b *buffer) pop() 


type Parser struct { 
  ch chan Lexed 
}

func newParser() *Parser {
  return &Parser{make(chan Lexed)}
}

func (p *Parser) Parse() {
  for i := range p.ch {

  }
}

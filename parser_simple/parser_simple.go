package parsersimple

import (
	l "arit/lexer_simple"
  "fmt"
)


type EnrichedToken struct {
  pos l.Position
  tok l.Token
  s string
}


type Parser struct { 
  lexer l.GLexer
  exps []Exp 
  stack *stack 
  lookahead *EnrichedToken
}

func NewParser(l l.GLexer) *Parser {
  return &Parser{
    lexer: l,
    exps: []Exp{},
    stack: NewStack(0),
  }
}

func (p *Parser) reduce() {
  switch p.stack.peek() {
    case 0:

  }
}

func (p *Parser) shift(state int) {  
    pos, t, s := p.lexer.Lex()
    p.lookahead = &EnrichedToken{pos,t,s}
    p.stack.push(state)
}

func (p *Parser) parse() {
  
}

func (p *Parser) goTo() { 
  switch p.stack.peek()  { 
  case 0:
    p.stack.push(1)
  case 3:
    p.stack.push(9)
  case 4:
    p.stack.push(10)
  case 5:
    p.stack.push(11)
  case 6:
    p.stack.push(12)
  case 7:
    p.stack.push(13)
  case 8:
    p.stack.push(14)
  default:
    fmt.Println("default")
  }
}

func (p *Parser) Parse() Exp {
  p.parse()
  if (len(p.exps) > 1) {
    panic("fuck")
  }
  return p.exps[0]
}

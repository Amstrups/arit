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

type state struct {
  current int 
  next int
}

type Parser struct { 
  lexer l.GLexer
  exps []Exp 
  states []int
  lookahead *EnrichedToken
}

func NewParser(l l.GLexer) *Parser {
  return &Parser{
    lexer: l,
    exps: []Exp{},
    states: []int{0}, 
  }
}

func (p *Parser) popState() {
  p.states = p.states[:len(p.states)-2] 
}

func (p *Parser) reduce() {
  
}

func (p *Parser) shift() {  
    pos, t, s := p.lexer.Lex()
    p.lookahead = &EnrichedToken{pos,t,s}
}

func (p *Parser) parse() {
  
}

func (p *Parser) goTo() { 
  switch state := p.states[len(p.states)-1]; state { 
  case 0:
    p.states = append(p.states, 1)
  case 3:
    p.states = append(p.states, 9)
  case 4:
    p.states = append(p.states, 10)
  case 5:
    p.states = append(p.states, 11)
  case 6:
    p.states = append(p.states, 12)
  case 7:
    p.states = append(p.states, 13)
  case 8:
    p.states = append(p.states, 14)
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

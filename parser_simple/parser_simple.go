package parsersimple

import (
	l "arit/lexer_simple"
  "fmt"
)

/*
-- (E)BNF syntax --
1: expr      ::= sum
2: expr      ::= product
3: expr      ::= value


4: sum       ::= sum addsub prodvalue
5: sum       ::= prodvalue addsub prodvalue
6: sum       ::= sum '-' parensum
7: sum       ::= prodvalue '-' parensum

8: addsub    ::= '+' | '-'
9: muldiv    ::= '*' | '/'

10: parensum  ::= '(' sum ')'

11: prodvalue ::= value
12: prodvalue ::= product

13: product   ::= prodterm muldiv value
14: product   ::= prodterm muldiv parensum
15: product   ::= prodterm '/' parenprod

16: prodterm  ::= prodvalue
17: prodterm  ::= parensum

18: parenprod ::= '(' product ')'

19: value     ::= number

-- Grammophone syntax --
Ref: https://mdaines.github.io/grammophone/#/ 
1:  expr      -> sum .
2:  expr      -> product .
3:  expr      -> value .
    
    
4:  sum       -> sum addsub prodvalue .
5:  sum       -> prodvalue addsub prodvalue .
6:  sum       -> sum minus parensum .
7:  sum       -> prodvalue minus parensum .
    
8:   addsub    -> plus | minus . 
9:   muldiv    -> mul | div .
  
10:  parensum  -> lp sum rp . 
  
11:  prodvalue -> value .
12:  prodvalue -> product .
  
13:  product   -> prodterm muldiv value .
14:  product   -> prodterm muldiv parensum .
15:  product   -> prodterm div parenprod .

16:  prodterm  -> prodvalue .
17:  prodterm  -> parensum .
     
18:  parenprod -> lp product rp .
  
19:  value     -> number .
*/

type EnrichedToken struct {
  pos l.Position
  tok l.Token
  s string
}


type Parser struct { 
  lexer l.GLexer
  exps []expr 
  stack *stack 
  lookahead *EnrichedToken
}

func NewParser(l l.GLexer) *Parser {
  return &Parser{
    lexer: l,
    exps: []expr{},
    stack: NewStack(0),
  }
}

func (p *Parser) reduceToNonTerm(name string, pops int, use []int) {
  if len(p.exps) < pops {
    panic("cant pop that much")
  }
  es := p.exps[len(p.exps)-pops:] 
  et := []expr{}
  for _,e := range use {
    et = append(et, es[e])
  }
  p.exps = p.exps[:len(p.exps)-pops]
  p.exps = append(p.exps, &nonterm{name, et})
}

func (p *Parser) reduce(rule int) {
  switch rule {
    case 0:
      fmt.Println("rule 0")
    default:
      fmt.Println("basicly all other rules")
  }
}

func (p *Parser) shift(state int) {  
    pos, t, s := p.lexer.Lex()
    p.lookahead = &EnrichedToken{pos,t,s}
    p.exps = append(p.exps, &term{l.Name(t), s})
    p.stack.push(state)
}

func (p *Parser) parse() {
/*  while (p.lookahead.tok != l.EOF) {
    
} */
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

func (p *Parser) Parse() expr {
  p.parse()
  if (len(p.exps) > 1) {
    panic("fuck")
  }
  return p.exps[0]
}

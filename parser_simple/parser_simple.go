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

9: parensum  ::= '(' sum ')'

10: prodvalue ::= value
11: prodvalue ::= product

12: product   ::= prodterm muldiv value
13: product   ::= prodterm muldiv parensum
14: product   ::= prodterm '/' parenprod

15: prodterm  ::= prodvalue
16: prodterm  ::= parensum

17: parenprod ::= '(' product ')'

18: value     ::= number

-- Grammophone syntax --
Ref: https://mdaines.github.io/grammophone/#/ 

1: expr      -> sum .
2: expr      -> product .
3: expr      -> value .


4: sum       -> sum addsub prodvalue .
5: sum       -> prodvalue addsub prodvalue .
6: sum       -> sum minus parensum .
7: sum       -> prodvalue minus parensum .

8: addsub    -> plus | minus . 

9: parensum  -> lp sum rp . 

10: prodvalue -> value .
11: prodvalue -> product .

12: product   -> prodterm muldiv value .
13: product   -> prodterm muldiv parensum .
14: product   -> prodterm div parenprod .

15: prodterm  -> prodvalue .
16: prodterm  -> parensum .

17: parenprod -> lp product rp .

18: value     -> number .
*/

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
/*
0: E ::= number 
1: E ::= ( E )
2: E ::= E ADD E
3: E ::= E MIN E
4: E ::= E MUL E
5: E ::= E DIV E
6: E ::= MIN E

*/
func (p *Parser) reduce(rule int) {
  switch {
    case 0:


  }
}

func (p *Parser) shift(state int) {  
    pos, t, s := p.lexer.Lex()
    p.lookahead = &EnrichedToken{pos,t,s}
    p.stack.push(state)
}

func (p *Parser) parse() {
  while (p.lookahead.tok != l.EOF) {
    
  }
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

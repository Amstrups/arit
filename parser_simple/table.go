package parsersimple

import l "arit/lexer_simple"

func (p *Parser) l0345678() {
  switch p.lookahead.tok {
  case l.INT: 
    p.shift(2)
  case l.LPAREN:
    p.shift(3)
  case l.SUB:
    p.shift(4)
  default:
    panic("fuck")
  }
}

func (p *Parser) l1() {
  switch p.lookahead.tok {
  case l.SUB: 
    p.shift(5)
  case l.MUL:
    p.shift(6)
  case l.DIV:
    p.shift(7)
  case l.ADD:
    p.shift(8)
  case l.EOF:
    return
  default:
    panic("fuck")
  }
}

func (p *Parser) l2() {
  //switch p.lookahead.tok {
  //case
}


func (p *Parser) l9() {
  switch p.lookahead.tok {
  case l.RPAREN:
    p.shift(15)
  case l.SUB:
    p.shift(5)
  case l.MUL:
    p.shift(6)
  case l.DIV:
    p.shift(7)
  case l.ADD:
    p.shift(8)
  case l.EOF:
    return
  default:
    panic("fuck")
  }
}



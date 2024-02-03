package parser

import "fmt"

// Matching on UB+LB - Not my worst idea actually
const (
  // Terminals
  TERM_LB = iota 
  INT
  FLOAT
  TERM_UB

  // Binary Operation
  OP_LB
  ADD
  SUB
  DIV
  MUL
  POW
  LT
  GT
  EQ
  NEQ
  ASSIGN
  OP_UB

  PAREN_LB
  LPAREN
  RPAREN
  LSQUARE
  RSQUARE
  LBRACK
  RBRACK
  PAREN_UB

  PUNCH_LB
  COLON
  SCOLON
  DOT
  COMMA
  PUNCH_UB

  LOGIC_LB
  AND
  OR
  NEG
  LOGIC_UB

  VARS_LB
  VAR
  VARS_UB
)

type Token struct {
  tokent int 
  val string
  pos,line int 
  ext int
}

func (t Token) str() string { 
  return t.val
}

func SimpleToken(tt int, val string, pos int) Token {
  return Token{tt, val, pos, 1, 1} 
}

func RToken(tt int, val string, t Token) Token {
  return Token{tt, val, t.pos, t.line, t.ext}
}

type TokenTRange struct { 
  lb,ub int
} 

// SingleValueRange
func SVR(ts []int) []TokenTRange { 
  ranges := []TokenTRange{}
  for _,t := range ts { 
    ranges = append(ranges, TokenTRange{t-1, t+1})
  } 
  return ranges
} 
// Single-SingleValueRange
func SSVR(t int) TokenTRange { 
  return TokenTRange{t-1,t+1}
}

func PrettyPrint(ts []Token) {
  fmt.Println("Pretty printing: ")
  for _, s := range ts {
    fmt.Print(s.str(), " ")
  }
  fmt.Println()
}

func (t Token) isEqualType(ot Token) bool { 
  return t.tokent == ot.tokent
}


func (t Token) isEqual(ot TokenTRange) bool { 
  return t.tokent > ot.lb && t.tokent < ot.ub

} 

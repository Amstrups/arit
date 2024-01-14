package parser

import "fmt"

const (
  TERM_LB = iota 
  INT
  FLOAT
  TERM_UB

  OP_LB
  ADD
  SUB
  DIV
  MUL
  OP_UB

  PAREN_LB
  LPAREN
  RPAREN
  LSQUARE
  RSQUARE
  LBRACK
  RBRACK
  PAREN_UB

  EQ_LB
  LT
  GT
  EQ
  NEQ
  EQ_UB

  PUNCH_LB
  COLON
  SCOLON
  DOT
  COMMA
  PUNCH_UB

  LOGIC_LB
  AND
  OR
  LOGIC_UB

  VARS_LB
  VAR
  VARS_UB
)


type Token struct {
  tokent int 
  val string
  pos,line int 
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


func PrettyPrint(ts []Token) {
  fmt.Println("Pretty printing: ")
  for _, s := range ts {
    fmt.Println(s.val)
  }
}

func (t Token) isEqualType(ot Token) bool { 
  return t.tokent == ot.tokent
}

func (t Token) isEqual(ot TokenTRange) bool { 
  return t.tokent > ot.lb && t.tokent < ot.ub

} 

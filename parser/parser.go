package parser

import (
	"fmt"
)
var TERM_RANGE = TokenTRange{TERM_LB, TERM_UB}
var OP_RANGE = TokenTRange{OP_LB, OP_UB}
var PAREN_RANGE = TokenTRange{PAREN_LB, PAREN_UB} 
var EQ_RANGE = TokenTRange{EQ_LB, EQ_UB} 

// yes, wannabe ocaml
func lookheadMatch(pos, ahead int, input []Token, pattern []TokenTRange) bool {
  if (pos+ahead >= len(input) || ahead < len(pattern)) { 
    return false
  }

  tokens := input[pos:pos+ahead]

  eq := true
  fmt.Println("new compare")
  fmt.Println("pattern: ", pattern)
  for i,t := range tokens { 
    fmt.Println(t.tokent, "is equal to ", pattern[i])
    eq = eq && t.isEqual(pattern[i])
  }
  
  return eq 
}

func findNext(ts []Token, next int) int { 
  for i,t := range ts {
     if (t.tokent == next) { 
       return i
     }
  } 
  return -1
} 

func preprocess(input []Token) []Token { 
  o := []Token{}
  pos := 0
  
  for {
    if (pos >= len(input)) { 
      break
    } 
    t := input[pos]
    inc := 1
    switch { 
    case lookheadMatch(pos, 3, input, SVR([]int{INT, DOT, INT})):
      t2,t3 := input[pos+1],input[pos+2]
      floatVal := t.val+t2.val+t3.val
      t = Token{FLOAT, floatVal, t.pos, t.line}
      inc = 3 

    case lookheadMatch(pos, 2, input, SVR([]int{COLON, EQ})):
      t = Token{ASSIGN, ":=", t.pos, t.line}
      inc = 2 
    
    case lookheadMatch(pos, 2, input, SVR([]int{EQ, EQ})):
      t = Token{ASSIGN, "==", t.pos, t.line}
      inc = 2 
    }

    o = append(o, t)
    pos += inc 

  }

  return o
} 

func Parser(input []Token) []Exp {
  fmt.Println("original: ", input)
  input = preprocess(input)
  fmt.Println("processed: ", input)
  exps := []Exp{}
  pos := 0
  for {
    if (pos >= len(input)) {
      break 
    } 
    switch { 
    case lookheadMatch(pos, 3, input, []TokenTRange{TERM_RANGE,OP_RANGE,TERM_RANGE}):
      fmt.Println("making op", pos, len(input))
      exps = append(exps, NewOpExp(input[pos:pos+3]))
      pos += 2
    case input[pos].tokent == INT:
      exps = append(exps, NewIntExp(input[pos]))
    case input[pos].tokent == FLOAT:
      exps = append(exps, NewFloatExp(input[pos]))
    default:
      println("fuck", pos)
      fmt.Println(exps)
      return exps
    }
    pos += 1
  }
  fmt.Println(exps)
  return exps

}

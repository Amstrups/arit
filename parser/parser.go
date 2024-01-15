package parser

import (
	"fmt"
)
var TERM_RANGE = TokenTRange{TERM_LB, TERM_UB}
var OP_RANGE = TokenTRange{OP_LB, OP_UB}
var PAREN_RANGE = TokenTRange{PAREN_LB, PAREN_UB} 
var EQ_RANGE = TokenTRange{EQ_LB, EQ_UB} 

func lookheadMatch(input []Token, pattern []TokenTRange) bool {
  if (len(input) < len(pattern)) { 
    return false 
  }
  eq := true
  fmt.Println("new compare")
  fmt.Println("pattern: ", pattern)
  for i,t := range input { 
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
    if pos+3 <= len(input) && lookheadMatch(input[pos:pos+3], SVR([]int{INT, DOT, INT})) {
      t2,t3 := input[pos+1],input[pos+2]
      floatVal := t.val+t2.val+t3.val
      o = append(o, Token{FLOAT, floatVal, t.pos, t.line}) 
      pos += 3 
    } else if pos+2 <= len(input) && lookheadMatch(input[pos:pos+2], SVR([]int{COLON, EQ})) { 
      o = append(o, Token{ASSIGN, ":=", t.pos, t.line})
      pos += 2 
    } else {
      o = append(o, t)
      pos += 1
    }
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
      // yes, wannabe ocaml
    case lookheadMatch(input[pos:pos+3], []TokenTRange{TERM_RANGE,OP_RANGE,TERM_RANGE}):
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

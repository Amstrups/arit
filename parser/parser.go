package parser

import (
  "arit/util"
  "fmt"
  "log"
)
var TERM_RANGE = TokenTRange{TERM_LB, TERM_UB}
var OP_RANGE = TokenTRange{OP_LB, OP_UB}
var PAREN_RANGE = TokenTRange{PAREN_LB, PAREN_UB} 
var PUNCH_RANGE = TokenTRange{PUNCH_LB, PUNCH_UB} 


// yes, wannabe ocaml - namde "lookaheadMatch"
func look(pos, ahead int, input []Token, pattern []TokenTRange) bool {
  if (pos+ahead > len(input) || (ahead < len(pattern))) { 
    return false
  }
  
  tokens := input[pos:pos+ahead]

  eq := true
  for i,t := range tokens { 
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

func preprocess(input []Token) ([]Token, error) { 
  o := []Token{}
  pos := 0

  // might as well check for paren balance while processing
  ps := make(util.IntStack, 0) 

  for {
    if (pos >= len(input)){ 
      break
    } 
    t := input[pos]
    inc := 1
    switch { 
    case t.tokent == LPAREN:
      ps = ps.Push(len(o))

    case t.tokent == RPAREN:
      ps_, p, err :=  ps.Pop()
      ps = ps_ // complains ps is unused otherwise
      if (err != nil) {
        return o, err
      }
      o[p].ext = pos-p
    case look(pos, 2, input, []TokenTRange{OP_RANGE,SSVR(SUB),TERM_RANGE}):
      t2 := input[pos+1]
      t.tokent = t2.tokent
      t.val = "-"+t2.val
      t.ext = -1
      inc = 2

    case look(pos, 3, input, SVR([]int{INT, DOT, INT})):
      t2,t3 := input[pos+1],input[pos+2]
      floatVal := t.val+t2.val+t3.val
      t = RToken(FLOAT, floatVal, t)
      inc = 3 

    case look(pos, 2, input, SVR([]int{COLON, EQ})):
      t = RToken(ASSIGN, ":=", t)
      inc = 2 

    case look(pos, 2, input, SVR([]int{EQ, EQ})):
      t = RToken(ASSIGN, "==", t)
      inc = 2 
    }

    o = append(o, t)
    pos += inc 
  }

  return o, nil
} 

func Parser(input []Token) []Exp {
  fmt.Println("original: ", input)
  input, err := preprocess(input)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("processed: ", input)
  exps := []Exp{}
  pos := 0
  for {
    if (pos >= len(input)) {
      break 
    } 
    switch { 
    case look(pos, 3, input, []TokenTRange{TERM_RANGE,OP_RANGE,TERM_RANGE}):
      fmt.Println("making op", pos, len(input))
      exps = append(exps, NewOpExp(input[pos:pos+3]))
      pos += 2

    case input[pos].tokent == INT:
      fmt.Println("just making int")
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

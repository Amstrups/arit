package parser

import "fmt"

func lookheadMatch(input []Token, pattern []TokenTRange) bool {
  if (len(input) < len(pattern)) { 
    return false 
  }
  eq := true
  fmt.Println("new compare")
  for i,t := range input { 
    fmt.Println(t.tokent, "is equal to ", pattern[i])
    eq = eq && t.isEqual(pattern[i])
  }
  
  return eq 
}

func findNext(input []Token, next int)  { 

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
    } else {
      o = append(o, t)
      pos += 1
    }
  }

  return o
} 

func Parser(input []Token) []Exp {
  fmt.Println("original: ", input)
  TERM_RANGE := TokenTRange{TERM_LB, TERM_UB}
  OP_RANGE := TokenTRange{OP_LB, OP_UB}
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
      pos += 3
      println("found addition")

    case input[pos].tokent == 1:
      exps = append(exps, NewIntExp(input[pos]))
      pos += 1
    case input[pos].tokent == 2:
      exps = append(exps, NewFloatExp(input[pos]))
      pos += 1
    default:
      println("fuck", pos)
      fmt.Println(exps)
      return exps
    }
  }
  fmt.Println(exps)
  return exps

}

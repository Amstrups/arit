package parser

import "fmt"

const (
  INT = 1
  FLOAT = 2

  ADD = 3 // +
  SUB = 4 // -
  DIV = 5 // /
  MUL = 6 // *

  LPAREN = 7
  RPAREN = 8
  LSQUARE = 9
  RSQUARE = 10
  LBRACK = 11 
  RBRACK = 12

  LT = 13 // <
  GT = 14 // >
  EQ = 15 // == 
  NEQ = 16 // !=

  COLON = 17 // :
  SCOLON = 18 // ;
  DOT = 19 // .   

  AND = 20 // &&
  OR = 21 // ||

  VAR = 22 // variables
)

type Token struct {
  tokent int 
  val string
  pos,line int 
}

func PrettyPrint(ts []Token) {
  for _, s := range ts { 
    fmt.Println(s.val)
  }
}

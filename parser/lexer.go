package parser

import (
	"fmt"
	"log"
)

func byteIsNumber(b byte) bool {
	return b >= 48 && b <= 57 // thats a number
}

func byteIsVar(b byte) bool {
	return b >= 97 && b <= 122 // thats a letter
}

func findPos(input []byte, pos int, byteTypeCheck func(byte) bool) int {
	tail := pos
	for {
		if tail < len(input) && byteTypeCheck(input[tail]) {
			tail += 1
		} else {
			break
		}
	}
	return max(pos, tail-1)
}

func Lexer(input []byte) []Token {
  tokens := []Token{}
   
  pos := 0

  for {
    if pos >= len(input) {
      break
    }
    switch {
    case byteIsNumber(input[pos]):
      intEnd := findPos(input, pos, byteIsNumber)
      tokens = append(tokens, Token{INT, string(input[pos : intEnd+1]), pos, 1})
      pos = intEnd
    case byteIsVar(input[pos]):
      varEnd := findPos(input, pos, byteIsVar)
      tokens = append(tokens, Token{VAR, string(input[pos : varEnd+1]), pos, 1})
      pos = varEnd
    case input[pos] == 43:
      tokens = append(tokens, Token{ADD, "+", pos, 1})
    case input[pos] == 45:
      tokens = append(tokens, Token{SUB, "-", pos, 1})
    case input[pos] == 47:
      tokens = append(tokens, Token{DIV, "/", pos, 1})
    case input[pos] == 42:
      tokens = append(tokens, Token{MUL, "*", pos, 1})
    case input[pos] == 40:
      tokens = append(tokens, Token{LPAREN, "(", pos, 1})
    case input[pos] == 41:
      tokens = append(tokens, Token{RPAREN, ")", pos, 1})
    case input[pos] == 91:
      tokens = append(tokens, Token{LSQUARE, "[", pos, 1})
    case input[pos] == 93:
      tokens = append(tokens, Token{RSQUARE, "]", pos, 1})
    case input[pos] == 123:
      tokens = append(tokens, Token{LBRACK, "{", pos, 1})
    case input[pos] == 125:
      tokens = append(tokens, Token{RBRACK, "}", pos, 1})
    case input[pos] == 46:
      tokens = append(tokens, Token{DOT, ".", pos, 1})
    case input[pos] == 58:
      tokens = append(tokens, Token{COLON, ":", pos, 1})
    case input[pos] == 59:
      tokens = append(tokens, Token{SCOLON, ";", pos, 1})
    case input[pos] == 38:
      tokens = append(tokens, Token{AND, "&", pos, 1})
    case input[pos] == 124:
      tokens = append(tokens, Token{OR, "|", pos, 1})
    case input[pos] == 33:
      tokens = append(tokens, Token{NEG, "!", pos, 1})
    case input[pos] == 94:
      tokens = append(tokens, Token{POW, "^", pos, 1})
    case input[pos]== 32: 
      fmt.Println("found space")
    default:
      log.Fatal("unknown char: ", input[pos])
    }
    pos += 1

	}
	PrettyPrint(tokens)
	return tokens
}

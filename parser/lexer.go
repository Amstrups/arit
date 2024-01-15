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
      tokens = append(tokens, SimpleToken(INT, string(input[pos : intEnd+1]), pos))
      pos = intEnd
    case byteIsVar(input[pos]):
      varEnd := findPos(input, pos, byteIsVar)
      tokens = append(tokens, SimpleToken(VAR, string(input[pos : varEnd+1]), pos))
      pos = varEnd
    case input[pos] == 43:
      tokens = append(tokens, SimpleToken(ADD, "+", pos))
    case input[pos] == 45:
      tokens = append(tokens, SimpleToken(SUB, "-", pos))
    case input[pos] == 47:
      tokens = append(tokens, SimpleToken(DIV, "/", pos))
    case input[pos] == 42:
      tokens = append(tokens, SimpleToken(MUL, "*", pos))
    case input[pos] == 40:
      tokens = append(tokens, SimpleToken(LPAREN, "(", pos))
    case input[pos] == 41:
      tokens = append(tokens, SimpleToken(RPAREN, ")", pos))
    case input[pos] == 91:
      tokens = append(tokens, SimpleToken(LSQUARE, "[", pos))
    case input[pos] == 93:
      tokens = append(tokens, SimpleToken(RSQUARE, "]", pos))
    case input[pos] == 123:
      tokens = append(tokens, SimpleToken(LBRACK, "{", pos))
    case input[pos] == 125:
      tokens = append(tokens, SimpleToken(RBRACK, "}", pos))
    case input[pos] == 46:
      tokens = append(tokens, SimpleToken(DOT, ".", pos))
    case input[pos] == 58:
      tokens = append(tokens, SimpleToken(COLON, ":", pos))
    case input[pos] == 59:
      tokens = append(tokens, SimpleToken(SCOLON, ";", pos))
    case input[pos] == 38:
      tokens = append(tokens, SimpleToken(AND, "&", pos))
    case input[pos] == 124:
      tokens = append(tokens, SimpleToken(OR, "|", pos))
    case input[pos] == 33:
      tokens = append(tokens, SimpleToken(NEG, "!", pos))
    case input[pos] == 94:
      tokens = append(tokens, SimpleToken(POW, "^", pos))
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

package main

import (
	lexer "arit/lexer_simple"
  parser "arit/parser_simple"
	"fmt"
)

func main() { 
  l := lexer.NewSLexer("+-*/\n111.111")
  p := parser.NewParser(l)
  stmt := p.Parse()
  fmt.Println(stmt)
}

/* Scanner example

    var s scanner.Scanner
    input := "x := 5 + 3\nif x > 5 {\n    fmt.Println(\"x is greater than 5\")\n}" 
    s.Init(strings.NewReader(input))
    for {
        tok := s.Scan()
        if tok == scanner.EOF {
            break
        }
        fmt.Printf("Token: %s, Value: %s\n", s.TokenText(), string(tok))
    }
*/

/* Args example 

  fmt.Println(l)
  return 
  argsWithProg := os.Args 

  flag.Parse()
  fmt.Println(argsWithProg[1])
*/
